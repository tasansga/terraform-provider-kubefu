package kubefu

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	schemaApi "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	k8yaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

const (
	manifestFieldManagerDefault = "kubefu"
)

func resourceManifest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceManifestCreateOrUpdate,
		ReadContext:   resourceManifestRead,
		UpdateContext: resourceManifestCreateOrUpdate,
		DeleteContext: resourceManifestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceManifestImport,
		},
		Schema: map[string]*schema.Schema{
			"manifest": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kubernetes manifest to apply (YAML or JSON)",
			},
			"field_manager": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     manifestFieldManagerDefault,
				Description: "Field manager name used for server-side apply",
			},
			"force": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Force apply in case of conflicts",
			},
		},
	}
}

func resourceManifestCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return diags
	}
	manifest := d.Get("manifest").(string)
	fieldManager := d.Get("field_manager").(string)
	force := d.Get("force").(bool)

	objs, err := decodeManifests(manifest)
	if err != nil {
		return diag.FromErr(err)
	}
	for _, u := range objs {
		if err := validateManifest(u); err != nil {
			return diag.FromErr(err)
		}
	}

	var oldObjs []*unstructured.Unstructured
	if d.Id() != "" {
		if parsed, err := parseManifestIDs(d.Id()); err == nil {
			oldObjs = parsed
			inheritRecordedNamespaces(objs, oldObjs)
		}
	}

	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return diag.FromErr(err)
	}

	appliedObjs := make([]*unstructured.Unstructured, 0, len(objs))
	for _, u := range objs {
		applied, err := applyManifest(ctx, client, mapper, cfg, u, fieldManager, force)
		if err != nil {
			if len(oldObjs) > 0 {
				retained := retainUntouchedObjects(appliedObjs, oldObjs)
				d.SetId(manifestIDs(retained))
				return diag.FromErr(err)
			}
			if len(appliedObjs) > 0 {
				d.SetId(manifestIDs(appliedObjs))
			}
			return diag.FromErr(err)
		}
		if applied != nil {
			appliedObjs = append(appliedObjs, applied)
		} else {
			appliedObjs = append(appliedObjs, u)
		}
	}

	// On update, delete any previously managed objects that were removed from the manifest.
	if len(oldObjs) > 0 {
		var unpruned []*unstructured.Unstructured
		for _, oldU := range prunedObjects(oldObjs, appliedObjs) {
			if err := deleteManifest(ctx, client, mapper, cfg, oldU); err != nil {
				unpruned = append(unpruned, oldU)
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Failed to prune removed manifest object %s: %s", manifestID(oldU), err),
				})
			}
		}
		if len(unpruned) > 0 {
			allTracked := append(appliedObjs, unpruned...)
			d.SetId(manifestIDs(allTracked))
			return diags
		}
	}

	d.SetId(manifestIDs(appliedObjs))
	return nil
}

func resourceManifestRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return diags
	}

	var objs []*unstructured.Unstructured
	var idObjs []*unstructured.Unstructured
	if d.Id() != "" {
		idObjs, _ = parseManifestIDs(d.Id())
	}

	if manifest := strings.TrimSpace(d.Get("manifest").(string)); manifest != "" {
		parsed, err := decodeManifests(manifest)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, u := range parsed {
			if err := validateManifest(u); err != nil {
				return diag.FromErr(err)
			}
		}
		if len(idObjs) > 0 {
			inheritRecordedNamespaces(parsed, idObjs)
		}
		objs = parsed
	} else if len(idObjs) > 0 {
		objs = idObjs
	} else {
		return diag.Errorf("manifest is required unless importing by ID")
	}

	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return diag.FromErr(err)
	}

	var lives []*unstructured.Unstructured
	notFoundCount := 0
	for _, u := range objs {
		live, err := getManifest(ctx, client, mapper, cfg, u)
		if err != nil {
			if apierrors.IsNotFound(err) {
				notFoundCount++
				continue
			}
			return diag.FromErr(err)
		}
		lives = append(lives, live)
	}

	if notFoundCount == len(objs) {
		d.SetId("")
		return nil
	}

	if err := setManifestFromObjects(d, lives); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(manifestIDs(lives))
	return nil
}

func resourceManifestDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return diags
	}

	// Prioritize objects from d.Id() to delete exactly what was created in its recorded namespace.
	var objs []*unstructured.Unstructured
	if d.Id() != "" {
		parsed, err := parseManifestIDs(d.Id())
		if err == nil {
			objs = parsed
		}
	}
	if len(objs) == 0 {
		manifest := strings.TrimSpace(d.Get("manifest").(string))
		if manifest != "" {
			parsed, err := decodeManifests(manifest)
			if err != nil {
				return diag.FromErr(err)
			}
			for _, u := range parsed {
				if err := validateManifest(u); err != nil {
					return diag.FromErr(err)
				}
			}
			objs = parsed
		} else {
			return diag.Errorf("manifest is required unless importing by ID")
		}
	}

	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return diag.FromErr(err)
	}

	// Delete in reverse order so dependents/workloads are deleted before dependencies/CRDs.
	for _, u := range reverseObjects(objs) {
		err = deleteManifest(ctx, client, mapper, cfg, u)
		if err != nil && !apierrors.IsNotFound(err) && !meta.IsNoMatchError(err) {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceManifestImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return nil, fmt.Errorf("%s", diags[0].Summary)
	}
	objs, err := parseManifestIDs(d.Id())
	if err != nil {
		return nil, err
	}
	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return nil, err
	}
	lives := make([]*unstructured.Unstructured, 0, len(objs))
	for _, u := range objs {
		live, err := getManifest(ctx, client, mapper, cfg, u)
		if err != nil {
			return nil, err
		}
		lives = append(lives, live)
	}
	if err := setManifestFromObjects(d, lives); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func providerConfigFromMeta(meta any) (*providerConfig, diag.Diagnostics) {
	cfg, ok := meta.(*providerConfig)
	if !ok || cfg == nil {
		return nil, diag.Errorf("invalid provider configuration")
	}
	return cfg, nil
}

func decodeManifests(manifest string) ([]*unstructured.Unstructured, error) {
	trimmed := strings.TrimSpace(manifest)
	if trimmed == "" {
		return nil, fmt.Errorf("manifest must not be empty")
	}
	decoder := k8yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(trimmed)), 4096)
	var objs []*unstructured.Unstructured
	for {
		obj := map[string]any{}
		if err := decoder.Decode(&obj); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("decode manifest: %w", err)
		}
		if len(obj) == 0 {
			continue
		}
		objs = append(objs, &unstructured.Unstructured{Object: obj})
	}
	if len(objs) == 0 {
		return nil, fmt.Errorf("manifest contains no resources")
	}
	return objs, nil
}

func validateManifest(u *unstructured.Unstructured) error {
	if u.GetAPIVersion() == "" {
		return fmt.Errorf("manifest is missing apiVersion")
	}
	if u.GetKind() == "" {
		return fmt.Errorf("manifest is missing kind")
	}
	if u.GetName() == "" {
		return fmt.Errorf("manifest is missing metadata.name")
	}
	return nil
}

func buildKubeClients(cfg *providerConfig) (dynamic.Interface, meta.RESTMapper, error) {
	if cfg != nil {
		if cfg.dynamicClient != nil && cfg.restMapper != nil {
			return cfg.dynamicClient, cfg.restMapper, nil
		}
		if cfg.dynamicClient != nil || cfg.restMapper != nil {
			return nil, nil, fmt.Errorf("both dynamicClient and restMapper must be provided together for client injection")
		}
	}
	restCfg, err := buildRESTConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	client, err := dynamic.NewForConfig(restCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("create dynamic client: %w", err)
	}
	disco, err := discovery.NewDiscoveryClientForConfig(restCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("create discovery client: %w", err)
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(disco))
	return client, mapper, nil
}

func buildRESTConfig(cfg *providerConfig) (*rest.Config, error) {
	if cfg == nil {
		return nil, fmt.Errorf("provider configuration is missing")
	}
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	if cfg.KubeConfigPath != "" {
		loader.ExplicitPath = cfg.KubeConfigPath
	}
	overrides := &clientcmd.ConfigOverrides{}
	if cfg.KubeContext != "" {
		overrides.CurrentContext = cfg.KubeContext
	}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	restCfg, err := clientConfig.ClientConfig()
	if err == nil {
		return restCfg, nil
	}
	if cfg.KubeConfigPath == "" {
		if inCluster, inErr := rest.InClusterConfig(); inErr == nil {
			return inCluster, nil
		}
	}
	return nil, fmt.Errorf("load kubeconfig: %w", err)
}

func applyManifest(ctx context.Context, client dynamic.Interface, mapper meta.RESTMapper, cfg *providerConfig, u *unstructured.Unstructured, fieldManager string, force bool) (*unstructured.Unstructured, error) {
	mapping, err := mapper.RESTMapping(u.GroupVersionKind().GroupKind(), u.GroupVersionKind().Version)
	if err != nil {
		return nil, fmt.Errorf("map resource: %w", err)
	}
	ns, err := resolveNamespace(cfg, u, mapping)
	if err != nil {
		return nil, err
	}
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace && u.GetNamespace() == "" {
		u.SetNamespace(ns)
	}
	resource := client.Resource(mapping.Resource)
	var target dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		target = resource.Namespace(ns)
	} else {
		target = resource
	}
	payload, err := json.Marshal(u.Object)
	if err != nil {
		return nil, fmt.Errorf("marshal manifest: %w", err)
	}
	if fieldManager == "" {
		fieldManager = manifestFieldManagerDefault
	}
	options := metav1.PatchOptions{FieldManager: fieldManager}
	if force {
		forceApply := true
		options.Force = &forceApply
	}
	return target.Patch(ctx, u.GetName(), types.ApplyPatchType, payload, options)
}

func getManifest(ctx context.Context, client dynamic.Interface, mapper meta.RESTMapper, cfg *providerConfig, u *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	mapping, err := mapper.RESTMapping(u.GroupVersionKind().GroupKind(), u.GroupVersionKind().Version)
	if err != nil {
		if meta.IsNoMatchError(err) {
			return nil, &apierrors.StatusError{ErrStatus: metav1.Status{Reason: metav1.StatusReasonNotFound}}
		}
		return nil, fmt.Errorf("map resource: %w", err)
	}
	ns, err := resolveNamespace(cfg, u, mapping)
	if err != nil {
		return nil, err
	}
	resource := client.Resource(mapping.Resource)
	var target dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		target = resource.Namespace(ns)
	} else {
		target = resource
	}
	return target.Get(ctx, u.GetName(), metav1.GetOptions{})
}

func deleteManifest(ctx context.Context, client dynamic.Interface, mapper meta.RESTMapper, cfg *providerConfig, u *unstructured.Unstructured) error {
	mapping, err := mapper.RESTMapping(u.GroupVersionKind().GroupKind(), u.GroupVersionKind().Version)
	if err != nil {
		if meta.IsNoMatchError(err) {
			return nil
		}
		return fmt.Errorf("map resource: %w", err)
	}
	ns, err := resolveNamespace(cfg, u, mapping)
	if err != nil {
		return err
	}
	resource := client.Resource(mapping.Resource)
	var target dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		target = resource.Namespace(ns)
	} else {
		target = resource
	}
	err = target.Delete(ctx, u.GetName(), metav1.DeleteOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	return nil
}

func resolveNamespace(cfg *providerConfig, u *unstructured.Unstructured, mapping *meta.RESTMapping) (string, error) {
	if mapping.Scope.Name() != meta.RESTScopeNameNamespace {
		return "", nil
	}
	ns := strings.TrimSpace(u.GetNamespace())
	if ns != "" {
		return ns, nil
	}
	if cfg == nil || cfg.dynamicClient != nil {
		return "default", nil
	}
	if fromCfg, err := kubeconfigNamespace(cfg); err == nil && fromCfg != "" {
		return fromCfg, nil
	}
	return "default", nil
}

func kubeconfigNamespace(cfg *providerConfig) (string, error) {
	if cfg.KubeConfigPath == "" {
		if data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
			if inClusterNS := strings.TrimSpace(string(data)); inClusterNS != "" {
				return inClusterNS, nil
			}
		}
	}
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	if cfg.KubeConfigPath != "" {
		loader.ExplicitPath = cfg.KubeConfigPath
	}
	overrides := &clientcmd.ConfigOverrides{}
	if cfg.KubeContext != "" {
		overrides.CurrentContext = cfg.KubeContext
	}
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	ns, _, err := clientConfig.Namespace()
	if err != nil {
		return "", err
	}
	return ns, nil
}

func manifestID(u *unstructured.Unstructured) string {
	if u == nil {
		return ""
	}
	apiver := u.GetAPIVersion()
	kind := u.GetKind()
	name := u.GetName()
	if apiver == "" || kind == "" || name == "" {
		return ""
	}
	ns := u.GetNamespace()
	if ns == "" {
		ns = "cluster"
	}
	return fmt.Sprintf("%s/%s/%s/%s", apiver, kind, ns, name)
}

func manifestIDs(objs []*unstructured.Unstructured) string {
	if len(objs) == 0 {
		return ""
	}
	if len(objs) == 1 {
		return manifestID(objs[0])
	}
	ids := make([]string, 0, len(objs))
	for _, u := range objs {
		if id := manifestID(u); id != "" {
			ids = append(ids, id)
		}
	}
	return strings.Join(ids, ";")
}

func reverseObjects(objs []*unstructured.Unstructured) []*unstructured.Unstructured {
	if len(objs) <= 1 {
		return objs
	}
	reversed := make([]*unstructured.Unstructured, len(objs))
	for i, u := range objs {
		reversed[len(objs)-1-i] = u
	}
	return reversed
}

type objectKey struct {
	gk        schemaApi.GroupKind
	namespace string
	name      string
}

func objectKeyOf(u *unstructured.Unstructured) (objectKey, bool) {
	if u == nil || u.GetName() == "" || u.GetKind() == "" {
		return objectKey{}, false
	}
	return objectKey{
		gk:        u.GroupVersionKind().GroupKind(),
		namespace: u.GetNamespace(),
		name:      u.GetName(),
	}, true
}

func prunedObjects(oldObjs, newObjs []*unstructured.Unstructured) []*unstructured.Unstructured {
	newKeys := make(map[objectKey]struct{}, len(newObjs))
	for _, u := range newObjs {
		if key, ok := objectKeyOf(u); ok {
			newKeys[key] = struct{}{}
		}
	}
	var pruned []*unstructured.Unstructured
	for _, oldU := range reverseObjects(oldObjs) {
		if key, ok := objectKeyOf(oldU); ok {
			if _, exists := newKeys[key]; !exists {
				pruned = append(pruned, oldU)
			}
		}
	}
	return pruned
}

func inheritRecordedNamespaces(objs, recordedObjs []*unstructured.Unstructured) {
	for _, u := range objs {
		if u == nil || u.GetName() == "" || u.GetNamespace() != "" {
			continue
		}
		for _, rec := range recordedObjs {
			if rec == nil || rec.GetName() == "" || rec.GetNamespace() == "" {
				continue
			}
			if rec.GroupVersionKind().GroupKind() == u.GroupVersionKind().GroupKind() && rec.GetName() == u.GetName() {
				u.SetNamespace(rec.GetNamespace())
				break
			}
		}
	}
}

func retainUntouchedObjects(appliedObjs, oldObjs []*unstructured.Unstructured) []*unstructured.Unstructured {
	retained := make([]*unstructured.Unstructured, len(appliedObjs), len(appliedObjs)+len(oldObjs))
	copy(retained, appliedObjs)

	appliedKeys := make(map[objectKey]struct{}, len(appliedObjs))
	for _, u := range appliedObjs {
		if key, ok := objectKeyOf(u); ok {
			appliedKeys[key] = struct{}{}
		}
	}

	for _, oldU := range oldObjs {
		if key, ok := objectKeyOf(oldU); ok {
			if _, exists := appliedKeys[key]; !exists {
				retained = append(retained, oldU)
			}
		}
	}
	return retained
}

func parseManifestID(id string) (*unstructured.Unstructured, error) {
	parts := strings.Split(id, "/")
	var apiver, kind, namespace, name string
	switch len(parts) {
	case 4:
		apiver = strings.TrimSpace(parts[0])
		kind = strings.TrimSpace(parts[1])
		namespace = strings.TrimSpace(parts[2])
		name = strings.TrimSpace(parts[3])
	case 5:
		apiver = strings.TrimSpace(parts[0]) + "/" + strings.TrimSpace(parts[1])
		kind = strings.TrimSpace(parts[2])
		namespace = strings.TrimSpace(parts[3])
		name = strings.TrimSpace(parts[4])
	default:
		return nil, fmt.Errorf("invalid manifest ID %q (expected [group/]version/kind/namespace/name)", id)
	}
	if apiver == "" || kind == "" || name == "" {
		return nil, fmt.Errorf("invalid manifest ID %q (missing apiVersion/kind/name)", id)
	}
	u := &unstructured.Unstructured{}
	u.SetAPIVersion(apiver)
	u.SetKind(kind)
	u.SetName(name)
	if namespace != "" && namespace != "cluster" {
		u.SetNamespace(namespace)
	}
	return u, nil
}

func parseManifestIDs(id string) ([]*unstructured.Unstructured, error) {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, fmt.Errorf("manifest ID must not be empty")
	}
	rawIDs := strings.FieldsFunc(trimmed, func(r rune) bool {
		return r == ';' || r == '\n'
	})
	if len(rawIDs) == 0 {
		return nil, fmt.Errorf("invalid manifest ID %q", id)
	}
	var objs []*unstructured.Unstructured
	for _, rawID := range rawIDs {
		tok := strings.TrimSpace(rawID)
		if tok == "" {
			continue
		}
		u, err := parseManifestID(tok)
		if err != nil {
			return nil, err
		}
		objs = append(objs, u)
	}
	if len(objs) == 0 {
		return nil, fmt.Errorf("invalid manifest ID %q", id)
	}
	return objs, nil
}

func setManifestFromObjects(d *schema.ResourceData, objs []*unstructured.Unstructured) error {
	if len(objs) == 0 {
		return nil
	}
	manifestStr, _ := d.Get("manifest").(string)
	isJSON := len(objs) == 1 && manifestLooksLikeJSON(manifestStr)

	var docs []string
	for _, obj := range objs {
		copyObj := runtime.DeepCopyJSON(obj.Object)
		normalizeManifestObject(copyObj)
		sorted := sortManifestValue(copyObj)
		jsonPayload, err := json.Marshal(sorted)
		if err != nil {
			return fmt.Errorf("marshal live manifest: %w", err)
		}
		if isJSON {
			docs = append(docs, string(jsonPayload))
		} else {
			yamlPayload, err := yaml.JSONToYAML(jsonPayload)
			if err != nil {
				return fmt.Errorf("marshal live manifest yaml: %w", err)
			}
			docs = append(docs, string(yamlPayload))
		}
	}

	var payload string
	if isJSON || len(docs) == 1 {
		payload = docs[0]
	} else {
		var sb strings.Builder
		for i, doc := range docs {
			if i > 0 {
				sb.WriteString("---\n")
			}
			sb.WriteString(doc)
			if !strings.HasSuffix(doc, "\n") {
				sb.WriteString("\n")
			}
		}
		payload = sb.String()
	}

	if err := d.Set("manifest", payload); err != nil {
		return fmt.Errorf("set manifest: %w", err)
	}
	return nil
}

func normalizeManifestObject(obj map[string]interface{}) {
	if obj == nil {
		return
	}
	// Drop status entirely (server-managed).
	delete(obj, "status")

	kind, _ := obj["kind"].(string)

	metadata, _ := obj["metadata"].(map[string]interface{})
	if metadata != nil {
		delete(metadata, "creationTimestamp")
		delete(metadata, "managedFields")
		delete(metadata, "resourceVersion")
		delete(metadata, "uid")
		delete(metadata, "generation")
		delete(metadata, "selfLink")

		if annotations, ok := metadata["annotations"].(map[string]interface{}); ok {
			delete(annotations, "kubectl.kubernetes.io/last-applied-configuration")
			if len(annotations) == 0 {
				delete(metadata, "annotations")
			}
		}

		if kind == "Namespace" {
			if labels, ok := metadata["labels"].(map[string]interface{}); ok {
				delete(labels, "kubernetes.io/metadata.name")
				if len(labels) == 0 {
					delete(metadata, "labels")
				}
			}
		}
	}

	if kind == "Namespace" {
		if spec, ok := obj["spec"].(map[string]interface{}); ok {
			delete(spec, "finalizers")
			if len(spec) == 0 {
				delete(obj, "spec")
			}
		}
	}
}

func sortManifestValue(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		sorted := make(map[string]interface{}, len(v))
		for _, k := range keys {
			sorted[k] = sortManifestValue(v[k])
		}
		return sorted
	case []interface{}:
		sorted := make([]interface{}, len(v))
		for i := range v {
			sorted[i] = sortManifestValue(v[i])
		}
		return sorted
	default:
		return v
	}
}

func manifestLooksLikeJSON(manifest string) bool {
	trimmed := strings.TrimSpace(manifest)
	if trimmed == "" {
		return false
	}
	switch trimmed[0] {
	case '{', '[':
		return true
	default:
		return false
	}
}

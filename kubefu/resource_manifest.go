package kubefu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
	k8yaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
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

	u, err := decodeManifest(manifest)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := validateManifest(u); err != nil {
		return diag.FromErr(err)
	}

	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return diag.FromErr(err)
	}

	applied, err := applyManifest(ctx, client, mapper, cfg, u, fieldManager, force)
	if err != nil {
		return diag.FromErr(err)
	}

	id := manifestID(applied)
	if id == "" {
		id = manifestID(u)
	}
	d.SetId(id)
	return nil
}

func resourceManifestRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return diags
	}

	var u *unstructured.Unstructured
	if manifest := strings.TrimSpace(d.Get("manifest").(string)); manifest != "" {
		parsed, err := decodeManifest(manifest)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := validateManifest(parsed); err != nil {
			return diag.FromErr(err)
		}
		u = parsed
	} else if d.Id() != "" {
		parsed, err := parseManifestID(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		u = parsed
	} else {
		return diag.Errorf("manifest is required unless importing by ID")
	}

	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return diag.FromErr(err)
	}

	live, err := getManifest(ctx, client, mapper, cfg, u)
	if err != nil {
		if errors.IsNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if d.Id() == "" {
		d.SetId(manifestID(u))
	}
	if err := setManifestFromObject(d, live); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceManifestDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return diags
	}
	manifest := strings.TrimSpace(d.Get("manifest").(string))
	var u *unstructured.Unstructured
	if manifest != "" {
		parsed, err := decodeManifest(manifest)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := validateManifest(parsed); err != nil {
			return diag.FromErr(err)
		}
		u = parsed
	} else if d.Id() != "" {
		parsed, err := parseManifestID(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		u = parsed
	} else {
		return diag.Errorf("manifest is required unless importing by ID")
	}

	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return diag.FromErr(err)
	}

	err = deleteManifest(ctx, client, mapper, cfg, u)
	if err != nil && !errors.IsNotFound(err) {
		return diag.FromErr(err)
	}
	return nil
}

func resourceManifestImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	cfg, diags := providerConfigFromMeta(m)
	if diags.HasError() {
		return nil, fmt.Errorf("%s", diags[0].Summary)
	}
	u, err := parseManifestID(d.Id())
	if err != nil {
		return nil, err
	}
	client, mapper, err := buildKubeClients(cfg)
	if err != nil {
		return nil, err
	}
	live, err := getManifest(ctx, client, mapper, cfg, u)
	if err != nil {
		return nil, err
	}
	if err := setManifestFromObject(d, live); err != nil {
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

func decodeManifest(manifest string) (*unstructured.Unstructured, error) {
	trimmed := strings.TrimSpace(manifest)
	if trimmed == "" {
		return nil, fmt.Errorf("manifest must not be empty")
	}
	decoder := k8yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(trimmed)), 4096)
	obj := map[string]any{}
	if err := decoder.Decode(&obj); err != nil {
		return nil, fmt.Errorf("decode manifest: %w", err)
	}
	return &unstructured.Unstructured{Object: obj}, nil
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
	return target.Delete(ctx, u.GetName(), metav1.DeleteOptions{})
}

func resolveNamespace(cfg *providerConfig, u *unstructured.Unstructured, mapping *meta.RESTMapping) (string, error) {
	if mapping.Scope.Name() != meta.RESTScopeNameNamespace {
		return "", nil
	}
	ns := strings.TrimSpace(u.GetNamespace())
	if ns != "" {
		return ns, nil
	}
	if cfg == nil {
		return "default", nil
	}
	if fromCfg, err := kubeconfigNamespace(cfg); err == nil && fromCfg != "" {
		return fromCfg, nil
	}
	return "default", nil
}

func kubeconfigNamespace(cfg *providerConfig) (string, error) {
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

func parseManifestID(id string) (*unstructured.Unstructured, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid manifest ID %q (expected apiVersion/kind/namespace/name)", id)
	}
	apiver := strings.TrimSpace(parts[0])
	kind := strings.TrimSpace(parts[1])
	namespace := strings.TrimSpace(parts[2])
	name := strings.TrimSpace(parts[3])
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

func setManifestFromObject(d *schema.ResourceData, obj *unstructured.Unstructured) error {
	if obj == nil {
		return nil
	}
	copyObj := runtime.DeepCopyJSON(obj.Object)
	normalizeManifestObject(copyObj)
	sorted := sortManifestValue(copyObj)
	jsonPayload, err := json.Marshal(sorted)
	if err != nil {
		return fmt.Errorf("marshal live manifest: %w", err)
	}
	payload := string(jsonPayload)
	if !manifestLooksLikeJSON(d.Get("manifest").(string)) {
		yamlPayload, err := yaml.JSONToYAML(jsonPayload)
		if err != nil {
			return fmt.Errorf("marshal live manifest yaml: %w", err)
		}
		payload = string(yamlPayload)
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

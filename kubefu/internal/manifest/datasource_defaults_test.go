package manifest

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/yaml"
)

type testRenderModeConfig struct {
	mode string
}

func (c testRenderModeConfig) ManifestRenderMode() string {
	return c.mode
}

func TestSetDataSourceManifestWithObjectPathsUnwrapsNestedObjects(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"template": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					MaxItems:   1,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"metadata": {
							Type:     schema.TypeMap,
							Optional: true,
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"template": []interface{}{
					map[string]interface{}{
						"metadata": map[string]interface{}{
							"name": "example",
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, []string{"spec", "spec.template"}); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	spec, ok := manifest["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map, got %T", manifest["spec"])
	}
	template, ok := spec["template"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.template map, got %T", spec["template"])
	}
	meta, ok := template["metadata"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.template.metadata map, got %T", template["metadata"])
	}
	if meta["name"] != "example" {
		t.Fatalf("expected metadata.name example, got %v", meta["name"])
	}
}

func TestSetDataSourceManifestWithObjectPathsLowerCamelizesObjectKeys(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"template": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					MaxItems:   1,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"metadata": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"annotations": {
									Type:     schema.TypeMap,
									Optional: true,
								},
							}},
						},
						"spec": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"service_account_name": {
									Type:     schema.TypeString,
									Optional: true,
								},
							}},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"template": []interface{}{
					map[string]interface{}{
						"metadata": []interface{}{
							map[string]interface{}{
								"annotations": map[string]interface{}{
									"example_key": "value",
								},
							},
						},
						"spec": []interface{}{
							map[string]interface{}{
								"service_account_name": "default",
							},
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	objectPaths := []string{
		"spec",
		"spec.template",
		"spec.template.metadata",
		"spec.template.spec",
	}
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, objectPaths); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	spec, ok := manifest["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map, got %T", manifest["spec"])
	}
	template, ok := spec["template"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.template map, got %T", spec["template"])
	}
	meta, ok := template["metadata"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.template.metadata map, got %T", template["metadata"])
	}
	annotations, ok := meta["annotations"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected annotations map, got %T", meta["annotations"])
	}
	if annotations["example_key"] != "value" {
		t.Fatalf("expected annotations.example_key value, got %v", annotations["example_key"])
	}
	specNested, ok := template["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.template.spec map, got %T", template["spec"])
	}
	if specNested["serviceAccountName"] != "default" {
		t.Fatalf("expected serviceAccountName default, got %v", specNested["serviceAccountName"])
	}
	if _, ok := specNested["service_account_name"]; ok {
		t.Fatalf("unexpected service_account_name key in manifest")
	}
}

func TestSetDataSourceManifestWithObjectPathsCompactPrunesImplicitEmptyAndZeroValues(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"api_version":          {Type: schema.TypeString, Computed: true},
		"kind":                 {Type: schema.TypeString, Computed: true},
		"name":                 {Type: schema.TypeString, Optional: true, Computed: true},
		"enabled":              {Type: schema.TypeBool, Optional: true, Computed: true},
		"replicas":             {Type: schema.TypeInt, Optional: true, Computed: true},
		"tags":                 {Type: schema.TypeList, Optional: true, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
		"metadata":             {Type: schema.TypeMap, Optional: true, Computed: true},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"api_version": "example/v1",
		"kind":        "Example",
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := d.Set("name", ""); err != nil {
		t.Fatalf("set name: %v", err)
	}
	if err := d.Set("enabled", false); err != nil {
		t.Fatalf("set enabled: %v", err)
	}
	if err := d.Set("replicas", 0); err != nil {
		t.Fatalf("set replicas: %v", err)
	}
	if err := d.Set("tags", []interface{}{}); err != nil {
		t.Fatalf("set tags: %v", err)
	}
	if err := d.Set("metadata", map[string]interface{}{}); err != nil {
		t.Fatalf("set metadata: %v", err)
	}
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"name", "enabled", "replicas", "tags", "metadata"}, nil); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	if _, ok := manifest["name"]; ok {
		t.Fatalf("expected implicit name to be pruned in compact mode")
	}
	if _, ok := manifest["enabled"]; ok {
		t.Fatalf("expected implicit enabled=false to be pruned in compact mode")
	}
	if _, ok := manifest["replicas"]; ok {
		t.Fatalf("expected implicit replicas=0 to be pruned in compact mode")
	}
	if _, ok := manifest["tags"]; ok {
		t.Fatalf("expected implicit tags=[] to be pruned in compact mode")
	}
	if _, ok := manifest["metadata"]; ok {
		t.Fatalf("expected implicit metadata={} to be pruned in compact mode")
	}
}

func TestPruneManifestValueCompactKeepsExplicitEmptyAndZeroValues(t *testing.T) {
	explicit := map[string]struct{}{
		"name":     {},
		"enabled":  {},
		"replicas": {},
		"tags":     {},
		"metadata": {},
	}
	tests := []struct {
		path  string
		input interface{}
	}{
		{path: "name", input: ""},
		{path: "enabled", input: false},
		{path: "replicas", input: 0},
		{path: "tags", input: []interface{}{}},
		{path: "metadata", input: map[string]interface{}{}},
	}
	for _, tt := range tests {
		ctx := &manifestContext{
			explicitPaths: explicit,
			renderMode:    RenderModeCompact,
		}
		out, keep := pruneManifestValue(tt.input, tt.path, ctx)
		if !keep {
			t.Fatalf("expected explicit value at %s to be kept in compact mode", tt.path)
		}
		if !reflect.DeepEqual(out, tt.input) {
			t.Fatalf("expected explicit value at %s to be preserved, got %T (%v)", tt.path, out, out)
		}
	}
}

func TestPruneManifestValueCanonicalKeepsEmptyAndZeroValues(t *testing.T) {
	cases := []interface{}{
		"",
		false,
		0,
		[]interface{}{},
		map[string]interface{}{},
	}
	for _, input := range cases {
		ctx := &manifestContext{
			renderMode: RenderModeCanonical,
		}
		out, keep := pruneManifestValue(input, "root", ctx)
		if !keep {
			t.Fatalf("expected canonical mode to keep value %T (%v)", input, input)
		}
		if reflect.DeepEqual(out, input) == false {
			t.Fatalf("expected canonical mode to preserve value %T (%v), got %T (%v)", input, input, out, out)
		}
	}
}

func TestPruneManifestValueCompactKeepsNestedExplicitValuesOnObjectPaths(t *testing.T) {
	explicit := map[string]struct{}{
		"spec.template.spec.service_account_name": {},
	}
	objectPaths := map[string]struct{}{
		"spec":               {},
		"spec.template":      {},
		"spec.template.spec": {},
	}
	value := map[string]interface{}{
		"serviceAccountName": "",
	}
	ctx := &manifestContext{
		explicitPaths: explicit,
		objectPaths:   objectPaths,
		renderMode:    RenderModeCompact,
	}
	out, keep := pruneManifestValue(value, "spec.template.spec", ctx)
	if !keep {
		t.Fatalf("expected nested explicit value to be kept on object path")
	}
	result, ok := out.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map output, got %T", out)
	}
	v, ok := result["serviceAccountName"]
	if !ok || v != "" {
		t.Fatalf("expected serviceAccountName to be preserved, got %v", result["serviceAccountName"])
	}
}

func TestPruneManifestValueCompactKeepsNestedExplicitValuesOnEscapedObjectPaths(t *testing.T) {
	explicit := map[string]struct{}{
		"spec.provider_.enabled": {},
	}
	objectPaths := map[string]struct{}{
		"spec":           {},
		"spec.provider_": {},
	}
	value := map[string]interface{}{
		"provider": map[string]interface{}{
			"enabled": false,
		},
	}
	ctx := &manifestContext{
		explicitPaths: explicit,
		objectPaths:   objectPaths,
		renderMode:    RenderModeCompact,
	}
	out, keep := pruneManifestValue(value, "spec", ctx)
	if !keep {
		t.Fatalf("expected nested explicit escaped value to be kept on object path")
	}
	result, ok := out.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map output, got %T", out)
	}
	provider, ok := result["provider"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected provider map output, got %T", result["provider"])
	}
	v, ok := provider["enabled"]
	if !ok || v != false {
		t.Fatalf("expected provider.enabled=false to be preserved, got %v", provider["enabled"])
	}
}

func TestSetDataSourceManifestWithObjectPathsValuesYAMLRendersIntoValues(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"values_yaml": {
					Type:     schema.TypeString,
					Optional: true,
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"values_yaml": "cluster:\n  name: demo\nflag: false\ncount: 0\n",
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, []string{"spec"}); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	spec, ok := manifest["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map, got %T", manifest["spec"])
	}
	values, ok := spec["values"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.values map, got %T", spec["values"])
	}
	cluster, ok := values["cluster"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.values.cluster map, got %T", values["cluster"])
	}
	if cluster["name"] != "demo" {
		t.Fatalf("expected spec.values.cluster.name=demo, got %v", cluster["name"])
	}
	if values["flag"] != false {
		t.Fatalf("expected spec.values.flag=false, got %v", values["flag"])
	}
	count, ok := values["count"].(float64)
	if !ok || count != 0 {
		t.Fatalf("expected spec.values.count=0, got %T (%v)", values["count"], values["count"])
	}
	if _, ok := spec["valuesYaml"]; ok {
		t.Fatalf("unexpected spec.valuesYaml key in manifest")
	}
}

func TestSetDataSourceManifestWithObjectPathsValuesYAMLOmittedDoesNotRenderValues(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"interval": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"values_yaml": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"interval": "5m",
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := d.Set("spec", []interface{}{map[string]interface{}{"interval": "5m", "values_yaml": ""}}); err != nil {
		t.Fatalf("set spec: %v", err)
	}
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, []string{"spec"}); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	spec, ok := manifest["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map, got %T", manifest["spec"])
	}
	if _, ok := spec["values"]; ok {
		t.Fatalf("expected spec.values to be omitted when values_yaml is unset")
	}
	if spec["interval"] != "5m" {
		t.Fatalf("expected spec.interval=5m, got %v", spec["interval"])
	}
}

func TestSetDataSourceManifestWithObjectPathsValuesYAMLCanonicalRendersParsedObject(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"values_yaml": {
					Type:     schema.TypeString,
					Optional: true,
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"values_yaml": "cluster:\n  name: demo\n",
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := SetDataSourceManifestWithObjectPathsForMeta(d, testRenderModeConfig{mode: RenderModeCanonical}, []string{"spec"}, []string{"spec"}); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	spec, ok := manifest["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map, got %T", manifest["spec"])
	}
	values, ok := spec["values"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec.values map, got %T", spec["values"])
	}
	cluster, ok := values["cluster"].(map[string]interface{})
	if !ok || cluster["name"] != "demo" {
		t.Fatalf("expected parsed values object, got %v", values)
	}
	if _, ok := values["Value"]; ok {
		t.Fatalf("unexpected struct wrapper key in canonical output")
	}
}

func TestSetDataSourceManifestWithObjectPathsValuesYAMLRejectsNonObject(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"values_yaml": {
					Type:     schema.TypeString,
					Optional: true,
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	cases := []struct {
		name string
		yaml string
	}{
		{name: "scalar", yaml: "42"},
		{name: "list", yaml: "- a\n- b\n"},
	}
	for _, tc := range cases {
		d := schema.TestResourceDataRaw(t, testSchema, map[string]interface{}{
			"spec": []interface{}{
				map[string]interface{}{
					"values_yaml": tc.yaml,
				},
			},
		})
		err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, []string{"spec"})
		if err == nil {
			t.Fatalf("expected error for %s values_yaml shape", tc.name)
		}
		if !strings.Contains(err.Error(), "expected YAML object") {
			t.Fatalf("expected YAML object error for %s, got: %v", tc.name, err)
		}
	}
}

func TestSetDataSourceManifestWithObjectPathsFlattensKustomizationGeneratorWrappers(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"api_version": {Type: schema.TypeString, Computed: true},
		"kind":        {Type: schema.TypeString, Computed: true},
		"config_map_generator": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"generator_args": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					MaxItems:   1,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kv_pair_sources": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"files": {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							}},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"api_version": "kustomize.config.k8s.io/v1beta1",
		"kind":        "Kustomization",
		"config_map_generator": []interface{}{
			map[string]interface{}{
				"generator_args": []interface{}{
					map[string]interface{}{
						"name":      "network-config",
						"namespace": "foo-messaging",
						"kv_pair_sources": []interface{}{
							map[string]interface{}{
								"files": []interface{}{"bar.prototext"},
							},
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCompact},
		[]string{"config_map_generator"},
		[]string{"config_map_generator.generator_args", "config_map_generator.generator_args.kv_pair_sources"},
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	list, ok := manifest["configMapGenerator"].([]interface{})
	if !ok || len(list) != 1 {
		t.Fatalf("expected one configMapGenerator entry, got %T (%v)", manifest["configMapGenerator"], manifest["configMapGenerator"])
	}
	entry, ok := list[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected configMapGenerator entry map, got %T", list[0])
	}
	if entry["name"] != "network-config" {
		t.Fatalf("expected name=network-config, got %v", entry["name"])
	}
	if entry["namespace"] != "foo-messaging" {
		t.Fatalf("expected namespace=foo-messaging, got %v", entry["namespace"])
	}
	files, ok := entry["files"].([]interface{})
	if !ok || len(files) != 1 || files[0] != "bar.prototext" {
		t.Fatalf("expected files=[bar.prototext], got %v", entry["files"])
	}
	if _, ok := entry["generator_args"]; ok {
		t.Fatalf("unexpected generator_args key in rendered output")
	}
	if _, ok := entry["generatorArgs"]; ok {
		t.Fatalf("unexpected generatorArgs key in rendered output")
	}
	if _, ok := entry["kv_pair_sources"]; ok {
		t.Fatalf("unexpected kv_pair_sources key in rendered output")
	}
	if _, ok := entry["kvPairSources"]; ok {
		t.Fatalf("unexpected kvPairSources key in rendered output")
	}
}

func TestSetDataSourceManifestWithObjectPathsFlattensConfigMapArgsWrappers(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"api_version": {Type: schema.TypeString, Computed: true},
		"kind":        {Type: schema.TypeString, Computed: true},
		"generator_args": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"kv_pair_sources": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					MaxItems:   1,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"files": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"api_version": "kustomize.config.k8s.io/v1beta1",
		"kind":        "ConfigMapArgs",
		"generator_args": []interface{}{
			map[string]interface{}{
				"name": "network-config",
				"kv_pair_sources": []interface{}{
					map[string]interface{}{
						"files": []interface{}{"bar.prototext"},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCompact},
		[]string{"generator_args"},
		[]string{"generator_args", "generator_args.kv_pair_sources"},
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	if manifest["name"] != "network-config" {
		t.Fatalf("expected name=network-config, got %v", manifest["name"])
	}
	files, ok := manifest["files"].([]interface{})
	if !ok || len(files) != 1 || files[0] != "bar.prototext" {
		t.Fatalf("expected files=[bar.prototext], got %v", manifest["files"])
	}
	if _, ok := manifest["generator_args"]; ok {
		t.Fatalf("unexpected generator_args key in rendered output")
	}
	if _, ok := manifest["generatorArgs"]; ok {
		t.Fatalf("unexpected generatorArgs key in rendered output")
	}
	if _, ok := manifest["kv_pair_sources"]; ok {
		t.Fatalf("unexpected kv_pair_sources key in rendered output")
	}
	if _, ok := manifest["kvPairSources"]; ok {
		t.Fatalf("unexpected kvPairSources key in rendered output")
	}
}

func TestSetDataSourceManifestWithListOfObjectsPreservesArrayAndCamelCasesKeys(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"subjects": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"api_group": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"kind": {
					Type:     schema.TypeString,
					Required: true,
				},
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
			}},
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"api_groups": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"resources": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"subjects": []interface{}{
			map[string]interface{}{
				"api_group": "rbac.authorization.k8s.io",
				"kind":      "Group",
				"name":      "k3s-admins",
			},
		},
		"rules": []interface{}{
			map[string]interface{}{
				"api_groups": []interface{}{"apps"},
				"resources":  []interface{}{"pods"},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{} // neither subjects nor rules are MaxItems: 1
	allObjectPaths := []string{"subjects", "rules"}

	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCompact},
		[]string{"subjects", "rules"},
		singleObjectPaths,
		allObjectPaths,
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}

	// Verify subjects remains a list of length 1 (not unwrapped into a map)
	subjects, ok := manifest["subjects"].([]interface{})
	if !ok {
		t.Fatalf("expected subjects to be a slice, got %T: %v", manifest["subjects"], manifest["subjects"])
	}
	if len(subjects) != 1 {
		t.Fatalf("expected 1 subject, got %d", len(subjects))
	}
	subject0, ok := subjects[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected subject[0] to be map, got %T", subjects[0])
	}
	if subject0["apiGroup"] != "rbac.authorization.k8s.io" {
		t.Fatalf("expected apiGroup to be rbac.authorization.k8s.io, got %v", subject0["apiGroup"])
	}
	if _, ok := subject0["api_group"]; ok {
		t.Fatalf("unexpected snake_case api_group in subject[0]")
	}

	// Verify rules remains a list of length 1
	rules, ok := manifest["rules"].([]interface{})
	if !ok {
		t.Fatalf("expected rules to be a slice, got %T: %v", manifest["rules"], manifest["rules"])
	}
	if len(rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(rules))
	}
	rule0, ok := rules[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected rule[0] to be map, got %T", rules[0])
	}
	if _, ok := rule0["apiGroups"]; !ok {
		t.Fatalf("expected apiGroups in rule[0]")
	}
	if _, ok := rule0["api_groups"]; ok {
		t.Fatalf("unexpected snake_case api_groups in rule[0]")
	}
	apiGroups := rule0["apiGroups"].([]interface{})
	if len(apiGroups) != 1 || apiGroups[0] != "apps" {
		t.Fatalf("expected apiGroups to be ['apps'], got %v", apiGroups)
	}
}

func TestSetDataSourceManifestWithEmptyStringInListPreservesElement(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"rules": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"api_groups": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem:       &schema.Schema{Type: schema.TypeString},
				},
				"resources": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"verbs": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"rules": []interface{}{
			map[string]interface{}{
				"api_groups": []interface{}{""},
				"resources":  []interface{}{"pods"},
				"verbs":      []interface{}{"get"},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := d.Set("rules", raw["rules"]); err != nil {
		t.Fatalf("set rules: %v", err)
	}
	singleObjectPaths := []string{}
	allObjectPaths := []string{"rules"}

	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCompact},
		[]string{"rules"},
		singleObjectPaths,
		allObjectPaths,
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}

	rules, ok := manifest["rules"].([]interface{})
	if !ok || len(rules) != 1 {
		t.Fatalf("expected rules slice of len 1, got %T (%v)", manifest["rules"], manifest["rules"])
	}
	rule0, ok := rules[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected rule[0] to be map, got %T", rules[0])
	}
	apiGroups, ok := rule0["apiGroups"].([]interface{})
	if !ok {
		t.Fatalf("expected apiGroups in rule[0] to be slice, got %T (%v)", rule0["apiGroups"], rule0["apiGroups"])
	}
	if len(apiGroups) != 1 || apiGroups[0] != "" {
		t.Fatalf("expected apiGroups to be [\"\"], got %v", apiGroups)
	}
}

func TestSetDataSourceManifestWithNilStringInListPreservesElement(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"rules": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"api_groups": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem:       &schema.Schema{Type: schema.TypeString},
				},
				"resources": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"verbs": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"rules": []interface{}{
			map[string]interface{}{
				"api_groups": []interface{}{nil},
				"resources":  []interface{}{"pods"},
				"verbs":      []interface{}{"get"},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{}
	allObjectPaths := []string{"rules"}

	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCompact},
		[]string{"rules"},
		singleObjectPaths,
		allObjectPaths,
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}

	rules, ok := manifest["rules"].([]interface{})
	if !ok || len(rules) != 1 {
		t.Fatalf("expected rules slice of len 1, got %T (%v)", manifest["rules"], manifest["rules"])
	}
	rule0, ok := rules[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected rule[0] to be map, got %T", rules[0])
	}
	apiGroups, ok := rule0["apiGroups"].([]interface{})
	if !ok {
		t.Fatalf("expected apiGroups in rule[0] to be slice, got %T (%v)", rule0["apiGroups"], rule0["apiGroups"])
	}
	if len(apiGroups) != 1 || apiGroups[0] != "" {
		t.Fatalf("expected apiGroups to be [\"\"], got %v", apiGroups)
	}
}

func TestSetDataSourceManifestWithNilInObjectListDoesNotCoerceToString(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"rules": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"resources": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"verbs": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"rules": []interface{}{
			map[string]interface{}{
				"resources": []interface{}{"pods"},
				"verbs":     []interface{}{"get"},
			},
			nil,
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{}
	allObjectPaths := []string{"rules"}

	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCompact},
		[]string{"rules"},
		singleObjectPaths,
		allObjectPaths,
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}

	rules, ok := manifest["rules"].([]interface{})
	if !ok {
		t.Fatalf("expected rules slice, got %T (%v)", manifest["rules"], manifest["rules"])
	}
	for i, r := range rules {
		if _, isString := r.(string); isString {
			t.Fatalf("rules[%d] was unexpectedly coerced to string: %v", i, r)
		}
	}
	if len(rules) != 1 {
		t.Fatalf("expected rules slice of len 1 (nil omitted), got len %d (%v)", len(rules), rules)
	}
}

func TestSetDataSourceManifestCanonicalPreservesNilStringInList(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"rules": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"api_groups": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem:       &schema.Schema{Type: schema.TypeString},
				},
				"resources": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"verbs": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"rules": []interface{}{
			map[string]interface{}{
				"api_groups": []interface{}{nil},
				"resources":  []interface{}{"pods"},
				"verbs":      []interface{}{"get"},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{}
	allObjectPaths := []string{"rules"}

	if err := SetDataSourceManifestWithObjectPathsForMeta(
		d,
		testRenderModeConfig{mode: RenderModeCanonical},
		[]string{"rules"},
		singleObjectPaths,
		allObjectPaths,
	); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}

	rules, ok := manifest["rules"].([]interface{})
	if !ok || len(rules) != 1 {
		t.Fatalf("expected rules slice of len 1, got %T (%v)", manifest["rules"], manifest["rules"])
	}
	rule0, ok := rules[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected rule[0] to be map, got %T", rules[0])
	}
	apiGroups, ok := rule0["apiGroups"].([]interface{})
	if !ok {
		t.Fatalf("expected apiGroups in rule[0] to be slice, got %T (%v)", rule0["apiGroups"], rule0["apiGroups"])
	}
	if len(apiGroups) != 1 || apiGroups[0] != "" {
		t.Fatalf("expected apiGroups to be [\"\"], got %v", apiGroups)
	}
}

func TestSetDataSourceManifestSingleObjectPathNeverRendersAsListWhenEmpty(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"containers": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"liveness_probe": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"http_get": {
									Type:       schema.TypeList,
									Optional:   true,
									ConfigMode: schema.SchemaConfigModeAttr,
									MaxItems:   1,
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"path": {
											Type:     schema.TypeString,
											Optional: true,
										},
									}},
								},
							}},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"containers": []interface{}{
					map[string]interface{}{
						"name":           "c1",
						"liveness_probe": []interface{}{},
					},
					map[string]interface{}{
						"name": "c2",
						"liveness_probe": []interface{}{
							map[string]interface{}{
								"http_get": []interface{}{
									map[string]interface{}{
										"path": "/health",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{
		"spec",
		"spec.containers.liveness_probe",
		"spec.containers.liveness_probe.http_get",
	}
	allObjectPaths := []string{
		"spec",
		"spec.containers",
		"spec.containers.liveness_probe",
		"spec.containers.liveness_probe.http_get",
	}
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	spec, ok := manifest["spec"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected spec map, got %T", manifest["spec"])
	}
	containers, ok := spec["containers"].([]interface{})
	if !ok || len(containers) != 2 {
		t.Fatalf("expected 2 containers, got %T (%v)", spec["containers"], spec["containers"])
	}
	c1, ok := containers[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected container 0 map, got %T", containers[0])
	}
	if probe, exists := c1["livenessProbe"]; exists {
		t.Fatalf("expected container 0 not to have livenessProbe, but got %v (%T)", probe, probe)
	}
	c2, ok := containers[1].(map[string]interface{})
	if !ok {
		t.Fatalf("expected container 1 map, got %T", containers[1])
	}
	if _, exists := c2["livenessProbe"]; !exists {
		t.Fatalf("expected container 1 to have livenessProbe")
	}
}

func TestSetDataSourceManifestCoercesNumericIntOrStringToInt(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"ports": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"target_port": {
							Type:     schema.TypeString,
							Optional: true,
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port":        80,
						"target_port": "8080",
					},
					map[string]interface{}{
						"port":        443,
						"target_port": "https",
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{"spec"}
	allObjectPaths := []string{"spec", "spec.ports"}

	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	jsonPayload := d.Get("kubefu_manifest_json").(string)
	var manifest map[string]interface{}
	if err := json.Unmarshal([]byte(jsonPayload), &manifest); err != nil {
		t.Fatalf("unmarshal json: %v", err)
	}
	spec := manifest["spec"].(map[string]interface{})
	ports := spec["ports"].([]interface{})
	p1 := ports[0].(map[string]interface{})
	p2 := ports[1].(map[string]interface{})

	// Assert integer in JSON unmarshalled (float64 in json.Unmarshal)
	if tp, ok := p1["targetPort"].(float64); !ok || int(tp) != 8080 {
		t.Fatalf("expected targetPort in port 0 to be integer 8080, got %T (%v)", p1["targetPort"], p1["targetPort"])
	}
	if !strings.Contains(jsonPayload, `"targetPort":8080`) {
		t.Fatalf("expected json to contain '\"targetPort\":8080', got: %s", jsonPayload)
	}

	if tp, ok := p2["targetPort"].(string); !ok || tp != "https" {
		t.Fatalf("expected targetPort in port 1 to be string 'https', got %T (%v)", p2["targetPort"], p2["targetPort"])
	}
	if !strings.Contains(jsonPayload, `"targetPort":"https"`) {
		t.Fatalf("expected json to contain '\"targetPort\":\"https\"', got: %s", jsonPayload)
	}

	yamlPayload := d.Get("kubefu_manifest_yaml").(string)
	if !strings.Contains(yamlPayload, "targetPort: 8080") {
		t.Fatalf("expected yaml to contain 'targetPort: 8080', got: %s", yamlPayload)
	}
	if strings.Contains(yamlPayload, "targetPort: \"8080\"") {
		t.Fatalf("yaml should not contain 'targetPort: \"8080\"', got: %s", yamlPayload)
	}
	if !strings.Contains(yamlPayload, "targetPort: https") {
		t.Fatalf("expected yaml to contain 'targetPort: https', got: %s", yamlPayload)
	}
}

func TestSetDataSourceManifestCoercesServicePortToInt(t *testing.T) {
	t.Run("service_port snake_case", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"spec": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				MaxItems:   1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"backend": {
						Type:       schema.TypeList,
						Optional:   true,
						ConfigMode: schema.SchemaConfigModeAttr,
						MaxItems:   1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"service_port": {
								Type:     schema.TypeString,
								Optional: true,
							},
						}},
					},
				}},
			},
			"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
			"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
		}
		raw := map[string]interface{}{
			"spec": []interface{}{
				map[string]interface{}{
					"backend": []interface{}{
						map[string]interface{}{
							"service_port": "80",
						},
					},
				},
			},
		}
		d := schema.TestResourceDataRaw(t, testSchema, raw)
		singleObjectPaths := []string{"spec", "spec.backend"}
		allObjectPaths := []string{"spec", "spec.backend"}

		if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
			t.Fatalf("set manifest: %v", err)
		}

		jsonPayload := d.Get("kubefu_manifest_json").(string)
		var manifest map[string]interface{}
		if err := json.Unmarshal([]byte(jsonPayload), &manifest); err != nil {
			t.Fatalf("unmarshal json: %v", err)
		}
		spec := manifest["spec"].(map[string]interface{})
		backend := spec["backend"].(map[string]interface{})

		if sp, ok := backend["servicePort"].(float64); !ok || int(sp) != 80 {
			t.Fatalf("expected service_port in backend to be integer 80, got %T (%v)", backend["servicePort"], backend["servicePort"])
		}
		if !strings.Contains(jsonPayload, `"servicePort":80`) {
			t.Fatalf("expected json to contain '\"servicePort\":80', got: %s", jsonPayload)
		}
		if strings.Contains(jsonPayload, `"servicePort":"80"`) {
			t.Fatalf("json should not contain '\"servicePort\":\"80\"', got: %s", jsonPayload)
		}

		yamlPayload := d.Get("kubefu_manifest_yaml").(string)
		if !strings.Contains(yamlPayload, "servicePort: 80") {
			t.Fatalf("expected yaml to contain 'servicePort: 80', got: %s", yamlPayload)
		}
		if strings.Contains(yamlPayload, "servicePort: \"80\"") {
			t.Fatalf("yaml should not contain 'servicePort: \"80\"', got: %s", yamlPayload)
		}
	})

	t.Run("servicePort camelCase", func(t *testing.T) {
		testSchema := map[string]*schema.Schema{
			"spec": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				MaxItems:   1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"backend": {
						Type:       schema.TypeList,
						Optional:   true,
						ConfigMode: schema.SchemaConfigModeAttr,
						MaxItems:   1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"servicePort": {
								Type:     schema.TypeString,
								Optional: true,
							},
						}},
					},
				}},
			},
			"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
			"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
		}
		raw := map[string]interface{}{
			"spec": []interface{}{
				map[string]interface{}{
					"backend": []interface{}{
						map[string]interface{}{
							"servicePort": "443",
						},
					},
				},
			},
		}
		d := schema.TestResourceDataRaw(t, testSchema, raw)
		singleObjectPaths := []string{"spec", "spec.backend"}
		allObjectPaths := []string{"spec", "spec.backend"}

		if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
			t.Fatalf("set manifest: %v", err)
		}

		jsonPayload := d.Get("kubefu_manifest_json").(string)
		var manifest map[string]interface{}
		if err := json.Unmarshal([]byte(jsonPayload), &manifest); err != nil {
			t.Fatalf("unmarshal json: %v", err)
		}
		spec := manifest["spec"].(map[string]interface{})
		backend := spec["backend"].(map[string]interface{})

		if sp, ok := backend["servicePort"].(float64); !ok || int(sp) != 443 {
			t.Fatalf("expected servicePort in backend to be integer 443, got %T (%v)", backend["servicePort"], backend["servicePort"])
		}
		if !strings.Contains(jsonPayload, `"servicePort":443`) {
			t.Fatalf("expected json to contain '\"servicePort\":443', got: %s", jsonPayload)
		}
		if strings.Contains(jsonPayload, `"servicePort":"443"`) {
			t.Fatalf("json should not contain '\"servicePort\":\"443\"', got: %s", jsonPayload)
		}

		yamlPayload := d.Get("kubefu_manifest_yaml").(string)
		if !strings.Contains(yamlPayload, "servicePort: 443") {
			t.Fatalf("expected yaml to contain 'servicePort: 443', got: %s", yamlPayload)
		}
		if strings.Contains(yamlPayload, "servicePort: \"443\"") {
			t.Fatalf("yaml should not contain 'servicePort: \"443\"', got: %s", yamlPayload)
		}
	})
}

func TestSetDataSourceManifestCoercesProbePortToInt(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"containers": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"liveness_probe": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"http_get": {
									Type:       schema.TypeList,
									Optional:   true,
									ConfigMode: schema.SchemaConfigModeAttr,
									MaxItems:   1,
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"path": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"port": {
											Type:     schema.TypeString,
											Optional: true,
										},
									}},
								},
								"tcp_socket": {
									Type:       schema.TypeList,
									Optional:   true,
									ConfigMode: schema.SchemaConfigModeAttr,
									MaxItems:   1,
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"port": {
											Type:     schema.TypeString,
											Optional: true,
										},
									}},
								},
							}},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"containers": []interface{}{
					map[string]interface{}{
						"name": "c1",
						"liveness_probe": []interface{}{
							map[string]interface{}{
								"http_get": []interface{}{
									map[string]interface{}{
										"path": "/health",
										"port": "8080",
									},
								},
							},
						},
					},
					map[string]interface{}{
						"name": "c2",
						"liveness_probe": []interface{}{
							map[string]interface{}{
								"tcp_socket": []interface{}{
									map[string]interface{}{
										"port": "9090",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{
		"spec",
		"spec.containers.liveness_probe",
		"spec.containers.liveness_probe.http_get",
		"spec.containers.liveness_probe.tcp_socket",
	}
	allObjectPaths := []string{
		"spec",
		"spec.containers",
		"spec.containers.liveness_probe",
		"spec.containers.liveness_probe.http_get",
		"spec.containers.liveness_probe.tcp_socket",
	}

	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	jsonPayload := d.Get("kubefu_manifest_json").(string)
	var manifest map[string]interface{}
	if err := json.Unmarshal([]byte(jsonPayload), &manifest); err != nil {
		t.Fatalf("unmarshal json: %v", err)
	}
	spec := manifest["spec"].(map[string]interface{})
	containers := spec["containers"].([]interface{})
	c1 := containers[0].(map[string]interface{})
	lp1 := c1["livenessProbe"].(map[string]interface{})
	httpGet := lp1["httpGet"].(map[string]interface{})

	c2 := containers[1].(map[string]interface{})
	lp2 := c2["livenessProbe"].(map[string]interface{})
	tcpSocket := lp2["tcpSocket"].(map[string]interface{})

	// Check http_get port integer coercion
	if p, ok := httpGet["port"].(float64); !ok || int(p) != 8080 {
		t.Fatalf("expected httpGet.port to be integer 8080, got %T (%v)", httpGet["port"], httpGet["port"])
	}
	if !strings.Contains(jsonPayload, `"port":8080`) {
		t.Fatalf("expected json to contain '\"port\":8080', got: %s", jsonPayload)
	}
	if strings.Contains(jsonPayload, `"port":"8080"`) {
		t.Fatalf("json should not contain '\"port\":\"8080\"', got: %s", jsonPayload)
	}

	// Check tcp_socket port integer coercion
	if p, ok := tcpSocket["port"].(float64); !ok || int(p) != 9090 {
		t.Fatalf("expected tcpSocket.port to be integer 9090, got %T (%v)", tcpSocket["port"], tcpSocket["port"])
	}
	if !strings.Contains(jsonPayload, `"port":9090`) {
		t.Fatalf("expected json to contain '\"port\":9090', got: %s", jsonPayload)
	}
	if strings.Contains(jsonPayload, `"port":"9090"`) {
		t.Fatalf("json should not contain '\"port\":\"9090\"', got: %s", jsonPayload)
	}

	yamlPayload := d.Get("kubefu_manifest_yaml").(string)
	if !strings.Contains(yamlPayload, "port: 8080") {
		t.Fatalf("expected yaml to contain 'port: 8080', got: %s", yamlPayload)
	}
	if strings.Contains(yamlPayload, "port: \"8080\"") {
		t.Fatalf("yaml should not contain 'port: \"8080\"', got: %s", yamlPayload)
	}
	if !strings.Contains(yamlPayload, "port: 9090") {
		t.Fatalf("expected yaml to contain 'port: 9090', got: %s", yamlPayload)
	}
	if strings.Contains(yamlPayload, "port: \"9090\"") {
		t.Fatalf("yaml should not contain 'port: \"9090\"', got: %s", yamlPayload)
	}
}

func TestSetDataSourceManifestCoercesNetworkPolicyPortToInt(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"ingress": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"ports": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"port": {
									Type:     schema.TypeString,
									Optional: true,
								},
							}},
						},
					}},
				},
				"egress": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"ports": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"port": {
									Type:     schema.TypeString,
									Optional: true,
								},
							}},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"ingress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"port": "80",
							},
						},
					},
				},
				"egress": []interface{}{
					map[string]interface{}{
						"ports": []interface{}{
							map[string]interface{}{
								"port": "443",
							},
						},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{
		"spec",
	}
	allObjectPaths := []string{
		"spec",
		"spec.ingress",
		"spec.ingress.ports",
		"spec.egress",
		"spec.egress.ports",
	}

	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	jsonPayload := d.Get("kubefu_manifest_json").(string)
	var manifest map[string]interface{}
	if err := json.Unmarshal([]byte(jsonPayload), &manifest); err != nil {
		t.Fatalf("unmarshal json: %v", err)
	}
	spec := manifest["spec"].(map[string]interface{})
	ingressList := spec["ingress"].([]interface{})
	ingress0 := ingressList[0].(map[string]interface{})
	inPorts := ingress0["ports"].([]interface{})
	inPort0 := inPorts[0].(map[string]interface{})

	egressList := spec["egress"].([]interface{})
	egress0 := egressList[0].(map[string]interface{})
	egPorts := egress0["ports"].([]interface{})
	egPort0 := egPorts[0].(map[string]interface{})

	if p, ok := inPort0["port"].(float64); !ok || int(p) != 80 {
		t.Fatalf("expected ingress port to be integer 80, got %T (%v)", inPort0["port"], inPort0["port"])
	}
	if !strings.Contains(jsonPayload, `"port":80`) {
		t.Fatalf("expected json to contain '\"port\":80', got: %s", jsonPayload)
	}
	if strings.Contains(jsonPayload, `"port":"80"`) {
		t.Fatalf("json should not contain '\"port\":\"80\"', got: %s", jsonPayload)
	}

	if p, ok := egPort0["port"].(float64); !ok || int(p) != 443 {
		t.Fatalf("expected egress port to be integer 443, got %T (%v)", egPort0["port"], egPort0["port"])
	}
	if !strings.Contains(jsonPayload, `"port":443`) {
		t.Fatalf("expected json to contain '\"port\":443', got: %s", jsonPayload)
	}
	if strings.Contains(jsonPayload, `"port":"443"`) {
		t.Fatalf("json should not contain '\"port\":\"443\"', got: %s", jsonPayload)
	}

	yamlPayload := d.Get("kubefu_manifest_yaml").(string)
	if !strings.Contains(yamlPayload, "port: 80") {
		t.Fatalf("expected yaml to contain 'port: 80', got: %s", yamlPayload)
	}
	if strings.Contains(yamlPayload, "port: \"80\"") {
		t.Fatalf("yaml should not contain 'port: \"80\"', got: %s", yamlPayload)
	}
	if !strings.Contains(yamlPayload, "port: 443") {
		t.Fatalf("expected yaml to contain 'port: 443', got: %s", yamlPayload)
	}
	if strings.Contains(yamlPayload, "port: \"443\"") {
		t.Fatalf("yaml should not contain 'port: \"443\"', got: %s", yamlPayload)
	}
}

func TestSetDataSourceManifestSiblingListScopingNoPollution(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"spec": {
			Type:       schema.TypeList,
			Optional:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			MaxItems:   1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"containers": {
					Type:       schema.TypeList,
					Optional:   true,
					ConfigMode: schema.SchemaConfigModeAttr,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"args": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					}},
				},
			}},
		},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"spec": []interface{}{
			map[string]interface{}{
				"containers": []interface{}{
					map[string]interface{}{
						"name": "c1",
						"args": []interface{}{"--mode=sse"},
					},
					map[string]interface{}{
						"name": "c2",
						"args": []interface{}{},
					},
				},
			},
		},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	singleObjectPaths := []string{"spec"}
	allObjectPaths := []string{"spec", "spec.containers"}

	if err := SetDataSourceManifestWithObjectPaths(d, []string{"spec"}, singleObjectPaths, allObjectPaths); err != nil {
		t.Fatalf("set manifest: %v", err)
	}

	yamlPayload := d.Get("kubefu_manifest_yaml").(string)
	if strings.Contains(yamlPayload, "args: []") {
		t.Fatalf("expected yaml NOT to contain 'args: []' on sibling container, got:\n%s", yamlPayload)
	}
}

func TestStripIndices(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"metadata", "metadata"},
		{"spec.template.spec.containers", "spec.template.spec.containers"},
		{"spec.containers.0.ports.1.container_port", "spec.containers.ports.container_port"},
		{"spec.containers.0", "spec.containers"},
		{"0.1.2", ""},
		{"a.0.b.1", "a.b"},
		{"0", ""},
	}

	for _, tt := range tests {
		actual := stripIndices(tt.input)
		if actual != tt.expected {
			t.Errorf("stripIndices(%q) = %q, want %q", tt.input, actual, tt.expected)
		}
	}
}

func TestIsIntOrStringPath(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		// Probe paths
		{"spec.containers.0.liveness_probe.http_get.port", true},
		{"spec.containers.0.readiness_probe.tcp_socket.port", true},
		{"spec.containers.0.startup_probe.httpget.port", true},
		{"spec.containers.0.startup_probe.tcpsocket.port", true},
		{"spec.containers.0.lifecycle.post_start.http_get.port", true},
		{"spec.template.spec.containers.0.liveness_probe.port", true},
		// NetworkPolicy paths
		{"spec.ingress.0.ports.0.port", true},
		{"spec.egress.0.ports.0.port", true},
		{"spec.network_policy.port", true},
		{"spec.networkpolicy.port", true},
		// IntOrString target / availability fields
		{"spec.ports.0.target_port", true},
		{"spec.ports.0.targetPort", true},
		{"spec.rules.0.http.paths.0.backend.service_port", true},
		{"spec.rules.0.http.paths.0.backend.servicePort", true},
		{"spec.strategy.rolling_update.max_unavailable", true},
		{"spec.strategy.rolling_update.maxUnavailable", true},
		{"spec.strategy.rolling_update.max_surge", true},
		{"spec.strategy.rolling_update.maxSurge", true},
		{"spec.min_available", true},
		{"spec.minAvailable", true},
		// Trailing array index segments
		{"spec.target_port.0", true},
		{"spec.ports.0", false},
		// False positives that should NOT match
		{"spec.ports.0.port", false},
		{"spec.containers.0.ports.0.container_port", false},
		{"spec.containers.0.ports.0.host_port", false},
		{"spec.replicas", false},
		{"spec.ports", false},
	}

	for _, tt := range tests {
		actual := isIntOrStringPath(tt.path)
		if actual != tt.expected {
			t.Errorf("isIntOrStringPath(%q) = %v, want %v", tt.path, actual, tt.expected)
		}
	}
}

func TestCollectExplicitManifestPathsFromValue_PreservesExplicitEmptyStrings(t *testing.T) {
	m := map[string]interface{}{
		"rules": []interface{}{
			map[string]interface{}{
				"api_groups": []interface{}{""},
			},
		},
	}
	paths := make(map[string]struct{})
	for k, v := range m {
		collectExplicitManifestPathsFromValue(v, k, paths, false, nil)
	}

	if _, ok := paths["rules.0.api_groups"]; !ok {
		t.Errorf("expected paths to contain 'rules.0.api_groups', got: %v", paths)
	}
	if _, ok := paths["rules.0.api_groups.0"]; !ok {
		t.Errorf("expected paths to contain 'rules.0.api_groups.0', got: %v", paths)
	}
}

func TestCollectExplicitManifestPaths_StringsAlwaysExplicit(t *testing.T) {
	// 1. In raw HCL cty.Value traversal, even empty string with parentExplicit = false returns true and is in paths
	paths := make(map[string]struct{})
	ret := collectExplicitManifestPaths(cty.StringVal(""), "empty_str", paths, false, nil)
	if !ret {
		t.Errorf("expected collectExplicitManifestPaths(cty.StringVal(\"\"), parentExplicit=false) to return true, got false")
	}
	if _, ok := paths["empty_str"]; !ok {
		t.Errorf("expected paths to contain 'empty_str', got %v", paths)
	}

	// Also test whitespace string
	paths = make(map[string]struct{})
	ret = collectExplicitManifestPaths(cty.StringVal("   "), "whitespace_str", paths, false, nil)
	if !ret {
		t.Errorf("expected collectExplicitManifestPaths(cty.StringVal(\"   \"), parentExplicit=false) to return true, got false")
	}
	if _, ok := paths["whitespace_str"]; !ok {
		t.Errorf("expected paths to contain 'whitespace_str', got %v", paths)
	}

	// 2. cty.StringVal("") with parentExplicit = true (e.g. within a list) returns true and is in paths
	paths = make(map[string]struct{})
	ret = collectExplicitManifestPaths(cty.StringVal(""), "rules.0.api_groups.0", paths, true, nil)
	if !ret {
		t.Errorf("expected collectExplicitManifestPaths(cty.StringVal(\"\"), parentExplicit=true) to return true, got false")
	}
	if _, ok := paths["rules.0.api_groups.0"]; !ok {
		t.Errorf("expected paths to contain 'rules.0.api_groups.0', got %v", paths)
	}

	// 3. Non-empty string cty.StringVal("foo") returns true and is in paths
	paths = make(map[string]struct{})
	ret = collectExplicitManifestPaths(cty.StringVal("foo"), "foo_str", paths, false, nil)
	if !ret {
		t.Errorf("expected collectExplicitManifestPaths(cty.StringVal(\"foo\"), parentExplicit=false) to return true, got false")
	}
	if _, ok := paths["foo_str"]; !ok {
		t.Errorf("expected paths to contain 'foo_str', got %v", paths)
	}
}

func TestCollectExplicitManifestPaths_PreservesExplicitEmptyStringsInMap(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		val := cty.MapVal(map[string]cty.Value{
			"empty_val": cty.StringVal(""),
		})
		paths := make(map[string]struct{})
		ret := collectExplicitManifestPaths(val, "data", paths, false, nil)
		if !ret {
			t.Errorf("expected collectExplicitManifestPaths to return true, got false")
		}
		if _, ok := paths["data.empty_val"]; !ok {
			t.Errorf("expected paths to contain 'data.empty_val', got: %v", paths)
		}
	})

	t.Run("object", func(t *testing.T) {
		val := cty.ObjectVal(map[string]cty.Value{
			"empty_val": cty.StringVal(""),
		})
		paths := make(map[string]struct{})
		ret := collectExplicitManifestPaths(val, "data", paths, false, nil)
		if !ret {
			t.Errorf("expected collectExplicitManifestPaths to return true, got false")
		}
		if _, ok := paths["data.empty_val"]; !ok {
			t.Errorf("expected paths to contain 'data.empty_val', got: %v", paths)
		}
	})
}

func TestCollectExplicitManifestPathsFromValue_EmptyNestedCollectionWithParentExplicit(t *testing.T) {
	paths := make(map[string]struct{})
	val := []interface{}{}
	ret := collectExplicitManifestPathsFromValue(val, "spec.args", paths, true, nil)
	if !ret {
		t.Errorf("expected collectExplicitManifestPathsFromValue to return true for empty nested collection with parentExplicit=true, got false")
	}
	if _, ok := paths["spec.args"]; !ok {
		t.Errorf("expected paths to contain 'spec.args', got: %v", paths)
	}
}

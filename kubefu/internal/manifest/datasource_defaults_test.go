package manifest

import (
	"reflect"
	"strings"
	"testing"

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
		out, keep := pruneManifestValue(tt.input, tt.path, explicit, map[string]struct{}{}, RenderModeCompact)
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
		out, keep := pruneManifestValue(input, "root", map[string]struct{}{}, map[string]struct{}{}, RenderModeCanonical)
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
	out, keep := pruneManifestValue(value, "spec.template.spec", explicit, objectPaths, RenderModeCompact)
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
	out, keep := pruneManifestValue(value, "spec", explicit, objectPaths, RenderModeCompact)
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

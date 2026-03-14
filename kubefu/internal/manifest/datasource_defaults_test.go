package manifest

import (
	"reflect"
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

func TestSetDataSourceManifestWithObjectPathsCompactPrunesEmptyAndZeroValues(t *testing.T) {
	testSchema := map[string]*schema.Schema{
		"api_version":          {Type: schema.TypeString, Computed: true},
		"kind":                 {Type: schema.TypeString, Computed: true},
		"name":                 {Type: schema.TypeString, Optional: true},
		"enabled":              {Type: schema.TypeBool, Optional: true},
		"replicas":             {Type: schema.TypeInt, Optional: true},
		"tags":                 {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}},
		"metadata":             {Type: schema.TypeMap, Optional: true},
		"kubefu_manifest_json": {Type: schema.TypeString, Computed: true},
		"kubefu_manifest_yaml": {Type: schema.TypeString, Computed: true},
	}
	raw := map[string]interface{}{
		"api_version": "example/v1",
		"kind":        "Example",
		"name":        "",
		"enabled":     false,
		"replicas":    0,
		"tags":        []interface{}{},
		"metadata":    map[string]interface{}{},
	}
	d := schema.TestResourceDataRaw(t, testSchema, raw)
	if err := SetDataSourceManifestWithObjectPaths(d, []string{"name", "enabled", "replicas", "tags", "metadata"}, nil); err != nil {
		t.Fatalf("set manifest: %v", err)
	}
	payload := d.Get("kubefu_manifest_yaml").(string)
	var manifest map[string]interface{}
	if err := yaml.Unmarshal([]byte(payload), &manifest); err != nil {
		t.Fatalf("unmarshal yaml: %v", err)
	}
	if manifest["apiVersion"] != "example/v1" {
		t.Fatalf("expected apiVersion to remain, got %v", manifest["apiVersion"])
	}
	if manifest["kind"] != "Example" {
		t.Fatalf("expected kind to remain, got %v", manifest["kind"])
	}
	if _, ok := manifest["name"]; ok {
		t.Fatalf("expected name to be pruned in compact mode")
	}
	if _, ok := manifest["enabled"]; ok {
		t.Fatalf("expected enabled=false to be pruned in compact mode")
	}
	if _, ok := manifest["replicas"]; ok {
		t.Fatalf("expected replicas=0 to be pruned in compact mode")
	}
	if _, ok := manifest["tags"]; ok {
		t.Fatalf("expected tags=[] to be pruned in compact mode")
	}
	if _, ok := manifest["metadata"]; ok {
		t.Fatalf("expected metadata={} to be pruned in compact mode")
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
		out, keep := pruneManifestValue(input, RenderModeCanonical)
		if !keep {
			t.Fatalf("expected canonical mode to keep value %T (%v)", input, input)
		}
		if reflect.DeepEqual(out, input) == false {
			t.Fatalf("expected canonical mode to preserve value %T (%v), got %T (%v)", input, input, out, out)
		}
	}
}

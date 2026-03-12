package manifest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/yaml"
)

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

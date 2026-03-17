package resourcegen

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDefinitionAsDataSource(t *testing.T) {
	def := Definition{
		Kind:           "Test",
		Description:    "Test data source",
		DefinitionName: "io.k8s.api.core.v1.Test",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
	result, err := def.AsDataSource("genpkg", "k8s")
	if err != nil {
		t.Fatalf("AsDataSource() error: %v", err)
	}
	if result.Name != "datasource_k8s_test_v1.go" {
		t.Fatalf("unexpected file name %q", result.Name)
	}
	if !strings.Contains(string(result.Content), "func dataSourceK8sCoreTestV1() *schema.Resource") {
		t.Fatalf("generated data source missing func signature: %s", result.Content)
	}
	if !strings.Contains(string(result.Content), `"Test data source"`) {
		t.Fatalf("description missing from generated source: %s", result.Content)
	}
}

func TestDefinitionAsDataSourceAddsTypeMetaFieldsWhenMissing(t *testing.T) {
	def := Definition{
		Kind:        "Kustomization",
		Group:       "kustomize.config.k8s.io",
		Version:     "v1beta1",
		Description: "Generated data source for io.k8s.api.apps.v1.Kustomization",
		Schema: map[string]*schema.Schema{
			"config_map_generator": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
				}},
			},
		},
	}
	result, err := def.AsDataSource("generated", "kustomize")
	if err != nil {
		t.Fatalf("AsDataSource() error: %v", err)
	}
	content := string(result.Content)
	if !strings.Contains(content, `"api_version": {`) {
		t.Fatalf("generated data source missing api_version field: %s", content)
	}
	if !strings.Contains(content, `"kind": {`) {
		t.Fatalf("generated data source missing kind field: %s", content)
	}
	if !strings.Contains(content, `SetDataSourceDefaults(d, "kustomize.config.k8s.io/v1beta1", "Kustomization", "kustomize.config.k8s.io/v1beta1/Kustomization")`) {
		t.Fatalf("generated read function missing expected defaults: %s", content)
	}
}

func TestDefinitionAsDataSourceHelmReleaseUsesValuesYAML(t *testing.T) {
	def := Definition{
		Kind:        "HelmRelease",
		Group:       "helm.toolkit.fluxcd.io",
		Version:     "v2beta2",
		Description: "HelmRelease schema",
		Schema: map[string]*schema.Schema{
			"spec": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"values": {
						Type:     schema.TypeMap,
						Optional: true,
						Computed: true,
					},
				}},
			},
		},
	}
	result, err := def.AsDataSource("generated", "flux")
	if err != nil {
		t.Fatalf("AsDataSource() error: %v", err)
	}
	content := string(result.Content)
	if strings.Contains(content, `"values": {`) {
		t.Fatalf("generated data source should not include values map for HelmRelease: %s", content)
	}
	if !strings.Contains(content, `"values_yaml": {`) {
		t.Fatalf("generated data source missing values_yaml for HelmRelease: %s", content)
	}
	if !strings.Contains(content, "Values holds the values for this Helm release as YAML.") {
		t.Fatalf("generated values_yaml field missing expected description: %s", content)
	}
}

func TestDefinitionAsDataSourceDoesNotMutateInputSchema(t *testing.T) {
	def := Definition{
		Kind:    "HelmRelease",
		Group:   "helm.toolkit.fluxcd.io",
		Version: "v2beta2",
		Schema: map[string]*schema.Schema{
			"spec": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"values": {
						Type:     schema.TypeMap,
						Optional: true,
					},
				}},
			},
		},
	}
	_, err := def.AsDataSource("generated", "flux")
	if err != nil {
		t.Fatalf("AsDataSource() error: %v", err)
	}
	specField := def.Schema["spec"]
	if specField == nil {
		t.Fatalf("expected spec field on original schema")
	}
	specElem, ok := specField.Elem.(*schema.Resource)
	if !ok || specElem == nil {
		t.Fatalf("expected spec elem resource on original schema")
	}
	if _, ok := specElem.Schema["values"]; !ok {
		t.Fatalf("expected original schema to retain spec.values")
	}
	if _, ok := specElem.Schema["values_yaml"]; ok {
		t.Fatalf("expected original schema not to gain spec.values_yaml")
	}
}

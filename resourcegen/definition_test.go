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

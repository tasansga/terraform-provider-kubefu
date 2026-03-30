package resource

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDataSourceRenderOmitsCompatibilityVarWhenNoVersions(t *testing.T) {
	ds := &DataSource{
		PackageName:        "generated",
		FuncName:           "dataSourceExample",
		Description:        "example",
		Schema:             map[string]*schema.Schema{},
		APIVersion:         "v1",
		Kind:               "Example",
		ID:                 "example/v1/Example",
		CompatibleVersions: nil,
	}

	rendered, err := ds.Render()
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if strings.Contains(rendered, "CompatibleVersions") {
		t.Fatalf("expected no compatibility variable when versions are empty")
	}
}

func TestDataSourceRenderWritesCompatibilityVarWhenVersionsProvided(t *testing.T) {
	ds := &DataSource{
		PackageName: "generated",
		FuncName:    "dataSourceExample",
		Description: "example",
		Schema:      map[string]*schema.Schema{},
		APIVersion:  "v1",
		Kind:        "Example",
		ID:          "example/v1/Example",
		CompatibleVersions: []string{
			"v1.28.0",
		},
	}

	rendered, err := ds.Render()
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if !strings.Contains(rendered, "var dataSourceExampleCompatibleVersions = []string{") {
		t.Fatalf("expected compatibility variable when versions are provided")
	}
	if !strings.Contains(rendered, "\"v1.28.0\"") {
		t.Fatalf("expected rendered compatibility version entry")
	}
}

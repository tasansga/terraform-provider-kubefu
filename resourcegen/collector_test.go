package resourcegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefinitionCollectorPrefersStableVersion(t *testing.T) {
	collector := newDefinitionCollector()
	collector.add(Definition{Kind: "Test", Group: "example.io", Version: "v1beta1"}, "testprov", "1.0")
	collector.add(Definition{Kind: "Test", Group: "example.io", Version: "v1"}, "testprov", "1.1")
	defs := collector.list()
	if len(defs) != 2 {
		t.Fatalf("expected 2 definitions, got %d", len(defs))
	}
	versions := map[string]bool{
		defs[0].Version: true,
		defs[1].Version: true,
	}
	if !versions["v1"] || !versions["v1beta1"] {
		t.Fatalf("missing expected versions %v", versions)
	}
}

func TestDefinitionsFromCRD(t *testing.T) {
	crd := customResourceDefinition{
		Kind: "CustomResourceDefinition",
		Spec: struct {
			Group string `json:"group" yaml:"group"`
			Names struct {
				Kind string `json:"kind" yaml:"kind"`
			} `json:"names" yaml:"names"`
			Versions   []crdVersion     `json:"versions" yaml:"versions"`
			Version    string           `json:"version" yaml:"version"`
			Validation crdVersionSchema `json:"validation" yaml:"validation"`
		}{
			Group: "example.io",
			Names: struct {
				Kind string `json:"kind" yaml:"kind"`
			}{
				Kind: "Test",
			},
			Versions: []crdVersion{
				{
					Name: "v1beta1",
					Schema: crdVersionSchema{
						OpenAPIV3Schema: definition{
							Type: "object",
							Properties: map[string]property{
								"spec": {
									Type: "object",
									Properties: map[string]property{
										"foo": {Type: "string"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	defs := definitionsFromCRD(crd)
	if len(defs) != 1 {
		t.Fatalf("expected one definition, got %d", len(defs))
	}
	if defs[0].Kind != "Test" {
		t.Fatalf("unexpected kind: %s", defs[0].Kind)
	}
	if defs[0].Group != "example.io" {
		t.Fatalf("unexpected group: %s", defs[0].Group)
	}
	if _, ok := defs[0].Schema["spec"]; !ok {
		t.Fatalf("expected spec field")
	}
}

func TestGenerateFromSchemasWritesFiles(t *testing.T) {
	root := t.TempDir()
	providerDir := filepath.Join(root, "testprov")
	if err := os.Mkdir(providerDir, 0o755); err != nil {
		t.Fatalf("mkdir testprov: %v", err)
	}
	spec := `
{
  "paths": {
    "/api/v1/tests": {
      "get": {
        "x-kubernetes-group-version-kind": {
          "group": "example.io",
          "version": "v1",
          "kind": "Test"
        }
      }
    }
  },
  "definitions": {
    "io.k8s.api.example.io.v1.Test": {
      "description": "Test",
      "type": "object",
      "properties": {
        "metadata": {
          "description": "metadata",
          "type": "object",
          "properties": {
            "name": {"type": "string"}
          },
          "required": ["name"]
        }
      },
      "required": ["metadata"]
    }
  }
}
`
	if err := os.WriteFile(filepath.Join(providerDir, "schema.json"), []byte(spec), 0o644); err != nil {
		t.Fatalf("write spec: %v", err)
	}
	out := t.TempDir()
	if err := GenerateFromSchemas(root, out, "genpkg"); err != nil {
		t.Fatalf("generate: %v", err)
	}
	target := out
	entries, err := os.ReadDir(target)
	if err != nil {
		t.Fatalf("read output dir: %v", err)
	}
	found := false
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "datasource_") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected datasource file in %s", target)
	}
	dataSourcesContent, err := os.ReadFile(filepath.Join(target, "datasources.go"))
	if err != nil {
		t.Fatalf("read datasources file: %v", err)
	}
	if !strings.Contains(string(dataSourcesContent), `"kubefu_testprov_example_io_test_v1"`) {
		t.Fatalf("datasources file missing entries: %s", dataSourcesContent)
	}
}

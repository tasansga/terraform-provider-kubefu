package resourcegen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDataSources(t *testing.T) {
	root := t.TempDir()
	providerDir := filepath.Join(root, "testprov")
	if err := os.Mkdir(providerDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
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
	dataSources, err := LoadDataSources(root)
	if err != nil {
		t.Fatalf("load data sources: %v", err)
	}
	expected := "kubefu_testprov_test_v1"
	if _, ok := dataSources[expected]; !ok {
		t.Fatalf("missing data source %s", expected)
	}
}

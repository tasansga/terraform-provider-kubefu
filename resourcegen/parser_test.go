package resourcegen

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestParseRootResourcesBuildsSchema(t *testing.T) {
	jsonData := `
{
  "paths": {
    "/api/v1/configmaps": {
      "get": {
        "x-kubernetes-group-version-kind": {
          "group": "",
          "version": "v1",
          "kind": "ConfigMap"
        }
      }
    }
  },
  "definitions": {
    "io.k8s.api.core.v1.ConfigMap": {
      "description": "ConfigMap",
      "type": "object",
      "properties": {
        "metadata": {
          "description": "object metadata",
          "type": "object",
          "properties": {
            "name": {
              "description": "resource name",
              "type": "string"
            }
          }
        },
        "data": {
          "description": "data map",
          "type": "object"
        },
        "items": {
          "description": "string list",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "spec": {
          "$ref": "#/definitions/io.k8s.api.core.v1.ConfigMapSpec"
        }
      },
      "required": [
        "metadata"
      ]
    },
    "io.k8s.api.core.v1.ConfigMapSpec": {
      "description": "spec",
      "type": "object",
      "properties": {
        "replicas": {
          "description": "replica count",
          "type": "integer"
        }
      },
      "required": []
    }
  }
}
`
	entries, err := ParseRootResources([]byte(jsonData))
	if err != nil {
		t.Fatalf("unexpected error parsing entry: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	entry := entries[0]
	if entry.Path != "/api/v1/configmaps" {
		t.Fatalf("unexpected path %q", entry.Path)
	}
	expectedDef := "io.k8s.api.core.v1.ConfigMap"
	if entry.DefinitionName != expectedDef {
		t.Fatalf("expected definition %q, got %q", expectedDef, entry.DefinitionName)
	}
	if entry.DefinitionDescription != "ConfigMap" {
		t.Fatalf("unexpected description %q", entry.DefinitionDescription)
	}
	metadata := entry.Schema["metadata"]
	if metadata == nil {
		t.Fatalf("metadata schema missing")
	}
	if metadata.Type != schema.TypeMap {
		t.Fatalf("expected metadata to be a map, got %d", metadata.Type)
	}
	if !metadata.Required {
		t.Fatalf("expected metadata to be required")
	}
	if metadata.Elem != nil {
		t.Fatalf("expected metadata Elem to be nil for map")
	}
	data := entry.Schema["data"]
	if !data.Optional {
		t.Fatalf("expected data to be optional")
	}
	if data.Type != schema.TypeMap {
		t.Fatalf("expected data to be map, got %d", data.Type)
	}
	spec := entry.Schema["spec"]
	if spec == nil {
		t.Fatalf("spec schema missing")
	}
	if spec.Type != schema.TypeMap {
		t.Fatalf("expected spec to be map, got %d", spec.Type)
	}
	if spec.Elem != nil {
		t.Fatalf("expected spec Elem to be nil for map")
	}
	items := entry.Schema["items"]
	if items == nil {
		t.Fatalf("items schema missing")
	}
	if items.Type != schema.TypeList {
		t.Fatalf("expected items to be list, got %d", items.Type)
	}
	if elem, ok := items.Elem.(*schema.Schema); !ok || elem.Type != schema.TypeString {
		t.Fatalf("expected items list of strings")
	}
}

func TestParseRootResourcesDefinitionsOnly(t *testing.T) {
	doc := `{
  "definitions": {
    "io.k8s.api.apps.v1.Kustomization": {
      "description": "Local kustomize config",
      "type": "object",
      "properties": {
        "apiVersion": { "type": "string" },
        "kind": { "type": "string" },
        "metadata": { "type": "object" }
      },
      "x-kubernetes-group-version-kind": [
        { "group": "kustomize.config.k8s.io", "version": "v1beta1", "kind": "Kustomization" }
      ]
    }
  }
}`
	entries, err := ParseRootResources([]byte(doc))
	if err != nil {
		t.Fatalf("ParseRootResources returned error: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].Kind != "Kustomization" || entries[0].Group != "kustomize.config.k8s.io" || entries[0].Version != "v1beta1" {
		t.Fatalf("unexpected entry: %+v", entries[0])
	}
	if entries[0].Schema == nil || entries[0].Schema["metadata"] == nil {
		t.Fatalf("expected schema to include metadata")
	}
}

func TestSchemaBuilderNormalizesPropertyNames(t *testing.T) {
	builder := &schemaBuilder{}
	def := definition{
		Required: []string{"imagePullSecrets"},
		Properties: map[string]property{
			"imagePullSecrets": {
				Description: "test",
				Type:        "array",
				Items: &property{
					Type: "string",
				},
			},
		},
	}
	schemaMap := builder.buildProperties(def)
	field := schemaMap["image_pull_secrets"]
	if field == nil {
		t.Fatalf("expected normalized property key")
	}
	if !field.Required {
		t.Fatalf("expected normalized property to be required")
	}
	if _, ok := schemaMap["imagePullSecrets"]; ok {
		t.Fatalf("did not expect original camelCase key")
	}
}

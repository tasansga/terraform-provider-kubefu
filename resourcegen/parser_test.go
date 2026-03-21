package resourcegen

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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
        },
        "quantity": {
          "$ref": "#/definitions/io.k8s.apimachinery.pkg.api.resource.Quantity"
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
    },
    "io.k8s.apimachinery.pkg.api.resource.Quantity": {
      "description": "quantity",
      "type": "string"
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
	if metadata.Type != schema.TypeList {
		t.Fatalf("expected metadata to be a list, got %d", metadata.Type)
	}
	if !metadata.Required {
		t.Fatalf("expected metadata to be required")
	}
	if metadata.MaxItems != 1 {
		t.Fatalf("expected metadata MaxItems=1, got %d", metadata.MaxItems)
	}
	metaElem, ok := metadata.Elem.(*schema.Resource)
	if !ok || metaElem == nil {
		t.Fatalf("expected metadata Elem resource")
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
	if spec.Type != schema.TypeList {
		t.Fatalf("expected spec to be list, got %d", spec.Type)
	}
	if spec.MaxItems != 1 {
		t.Fatalf("expected spec MaxItems=1, got %d", spec.MaxItems)
	}
	specElem, ok := spec.Elem.(*schema.Resource)
	if !ok || specElem == nil {
		t.Fatalf("expected spec Elem resource")
	}
	if specField := specElem.Schema["replicas"]; specField == nil || specField.Type != schema.TypeInt {
		t.Fatalf("expected spec.replicas to be int")
	}
	quantity := entry.Schema["quantity"]
	if quantity == nil {
		t.Fatalf("quantity schema missing")
	}
	if quantity.Type != schema.TypeString {
		t.Fatalf("expected quantity to be string, got %d", quantity.Type)
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

func TestSchemaBuilderRefToObjectWithoutPropertiesUsesMap(t *testing.T) {
	defs := map[string]definition{
		"example.Parent": {
			Type: "object",
			Properties: map[string]property{
				"opaque": {
					Ref: "#/definitions/example.Opaque",
				},
			},
		},
		"example.Opaque": {
			Type:       "object",
			Properties: map[string]property{},
		},
	}
	builder := &schemaBuilder{definitions: defs, cache: make(map[string]map[string]*schema.Schema)}
	schemaMap := builder.buildDefinition("example.Parent", defs["example.Parent"])
	field := schemaMap["opaque"]
	if field == nil {
		t.Fatalf("expected opaque schema to exist")
	}
	if field.Type != schema.TypeMap {
		t.Fatalf("expected opaque to be map, got %d", field.Type)
	}
}

func TestSchemaBuilderRefToArrayDefinitionPopulatesElem(t *testing.T) {
	defs := map[string]definition{
		"example.Parent": {
			Type: "object",
			Properties: map[string]property{
				"events": {
					Ref: "#/definitions/example.EventList",
				},
			},
		},
		"example.EventList": {
			Type: "array",
			Items: &property{
				Type: "string",
			},
		},
	}
	builder := &schemaBuilder{definitions: defs, cache: make(map[string]map[string]*schema.Schema)}
	schemaMap := builder.buildDefinition("example.Parent", defs["example.Parent"])
	field := schemaMap["events"]
	if field == nil {
		t.Fatalf("expected events schema to exist")
	}
	if field.Type != schema.TypeList {
		t.Fatalf("expected events to be list, got %d", field.Type)
	}
	elem, ok := field.Elem.(*schema.Schema)
	if !ok || elem.Type != schema.TypeString {
		t.Fatalf("expected events elem to be string schema")
	}
}

func TestSchemaBuilderUnresolvedRefUsesMap(t *testing.T) {
	defs := map[string]definition{
		"example.Parent": {
			Type: "object",
			Properties: map[string]property{
				"opaque": {
					Ref: "#/definitions/example.Missing",
				},
			},
		},
	}
	builder := &schemaBuilder{definitions: defs, cache: make(map[string]map[string]*schema.Schema)}
	schemaMap := builder.buildDefinition("example.Parent", defs["example.Parent"])
	field := schemaMap["opaque"]
	if field == nil {
		t.Fatalf("expected opaque schema to exist")
	}
	if field.Type != schema.TypeMap {
		t.Fatalf("expected unresolved ref to be map, got %d", field.Type)
	}
}

func TestNormalizedPropertyNameSanitizesSpecialChars(t *testing.T) {
	if got := normalizedPropertyName("x-kubernetes-list-type"); got != "x_kubernetes_list_type" {
		t.Fatalf("expected x_kubernetes_list_type, got %q", got)
	}
	if got := normalizedPropertyName("$schema"); got != "schema" {
		t.Fatalf("expected schema, got %q", got)
	}
}

func TestSchemaBuilderPreserveUnknownFieldsUsesMap(t *testing.T) {
	defs := map[string]definition{
		"example.Parent": {
			Type: "object",
			Properties: map[string]property{
				"values": {
					XKubernetesPreserveUnknownFields: true,
				},
			},
		},
	}
	builder := &schemaBuilder{definitions: defs, cache: make(map[string]map[string]*schema.Schema)}
	schemaMap := builder.buildDefinition("example.Parent", defs["example.Parent"])
	field := schemaMap["values"]
	if field == nil {
		t.Fatalf("expected values schema to exist")
	}
	if field.Type != schema.TypeMap {
		t.Fatalf("expected values to be map, got %d", field.Type)
	}
}

func TestParseRootResourcesSupportsAPIsWithoutDottedGroup(t *testing.T) {
	doc := `{
  "paths": {
    "/apis/apps/v1/namespaces/{namespace}/deployments": {
      "get": {
        "x-kubernetes-group-version-kind": {
          "group": "apps",
          "version": "v1",
          "kind": "Deployment"
        }
      }
    },
    "/apis/apps/v1/namespaces/{namespace}/daemonsets": {
      "get": {
        "x-kubernetes-group-version-kind": {
          "group": "apps",
          "version": "v1",
          "kind": "DaemonSet"
        }
      }
    }
  },
  "definitions": {
    "io.k8s.api.apps.v1.Deployment": {
      "description": "Deployment",
      "type": "object",
      "properties": {
        "metadata": { "type": "object" }
      }
    },
    "io.k8s.api.apps.v1.DaemonSet": {
      "description": "DaemonSet",
      "type": "object",
      "properties": {
        "metadata": { "type": "object" }
      }
    }
  }
}`
	entries, err := ParseRootResources([]byte(doc))
	if err != nil {
		t.Fatalf("ParseRootResources returned error: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	if !hasGVK(entries, "apps", "v1", "Deployment") {
		t.Fatalf("expected apps/v1 Deployment entry")
	}
	if !hasGVK(entries, "apps", "v1", "DaemonSet") {
		t.Fatalf("expected apps/v1 DaemonSet entry")
	}
}

func TestParseRootResourcesK8sBaselineKindsPresent(t *testing.T) {
	path, err := latestK8sSchemaPath()
	if err != nil {
		t.Fatalf("find latest k8s schema: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read k8s schema %s: %v", path, err)
	}
	entries, err := ParseRootResources(data)
	if err != nil {
		t.Fatalf("parse k8s schema %s: %v", path, err)
	}
	required := []struct {
		group   string
		version string
		kind    string
	}{
		{group: "", version: "v1", kind: "ConfigMap"},
		{group: "", version: "v1", kind: "Service"},
		{group: "apps", version: "v1", kind: "Deployment"},
		{group: "apps", version: "v1", kind: "DaemonSet"},
	}
	for _, req := range required {
		if !hasGVK(entries, req.group, req.version, req.kind) {
			t.Fatalf("missing required kind %q/%q/%q in %s", req.group, req.version, req.kind, path)
		}
	}
}

func hasGVK(entries []ResourceEntry, group, version, kind string) bool {
	for _, entry := range entries {
		if entry.Group == group && entry.Version == version && entry.Kind == kind {
			return true
		}
	}
	return false
}

func latestK8sSchemaPath() (string, error) {
	matches, err := filepath.Glob(filepath.Join("..", "schemas", "k8s", "kubernetes-api-*.json"))
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", os.ErrNotExist
	}
	sort.Slice(matches, func(i, j int) bool {
		ai, aj := parseK8sSchemaMinor(matches[i]), parseK8sSchemaMinor(matches[j])
		if ai == aj {
			return matches[i] < matches[j]
		}
		return ai < aj
	})
	return matches[len(matches)-1], nil
}

func parseK8sSchemaMinor(path string) int {
	base := filepath.Base(path)
	const prefix = "kubernetes-api-1."
	const suffix = ".json"
	if !strings.HasPrefix(base, prefix) || !strings.HasSuffix(base, suffix) {
		return -1
	}
	minorPart := strings.TrimSuffix(strings.TrimPrefix(base, prefix), suffix)
	minor, err := strconv.Atoi(minorPart)
	if err != nil {
		return -1
	}
	return minor
}

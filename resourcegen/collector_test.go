package resourcegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDefinitionCollectorPrefersStableVersion(t *testing.T) {
	collector := newDefinitionCollector()
	collector.add(Definition{Kind: "Test", Group: "example.io", Version: "v1beta1"}, "testprov", "1.0")
	collector.add(Definition{Kind: "Test", Group: "example.io", Version: "v1"}, "testprov", "1.1")
	defs := collector.list()
	if len(defs) != 2 {
		t.Fatalf("expected 2 merged definitions, got %d", len(defs))
	}
	versions := map[string]bool{}
	for _, def := range defs {
		versions[def.Version] = true
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

func TestCollectDefinitionsReturnsMergedOnly(t *testing.T) {
	root := t.TempDir()
	providerDir := filepath.Join(root, "testprov")
	if err := os.Mkdir(providerDir, 0o755); err != nil {
		t.Fatalf("mkdir testprov: %v", err)
	}
	specTemplate := `
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
        "%s": {"type": "string"}
      }
    }
  }
}
`
	if err := os.WriteFile(filepath.Join(providerDir, "schema-1.10.0.json"), []byte(fmt.Sprintf(specTemplate, "field_v1_10")), 0o644); err != nil {
		t.Fatalf("write v1.10 spec: %v", err)
	}
	if err := os.WriteFile(filepath.Join(providerDir, "schema-1.11.0.json"), []byte(fmt.Sprintf(specTemplate, "field_v1_11")), 0o644); err != nil {
		t.Fatalf("write v1.11 spec: %v", err)
	}

	mergedOnly, err := CollectDefinitions(root)
	if err != nil {
		t.Fatalf("CollectDefinitions: %v", err)
	}
	defs := mergedOnly["testprov"]
	if len(defs) != 1 {
		t.Fatalf("expected merged-only output to contain one definition, got %d", len(defs))
	}
	if !strings.Contains(strings.Join(defs[0].ProviderVersions, ","), "v1.10.0") || !strings.Contains(strings.Join(defs[0].ProviderVersions, ","), "v1.11.0") {
		t.Fatalf("expected merged definition to carry both source versions, got %v", defs[0].ProviderVersions)
	}
}

func TestDefinitionCollectorMergesSchemaForSameGVK(t *testing.T) {
	collector := newDefinitionCollector()
	collector.add(Definition{
		Kind:    "ConfigMap",
		Version: "v1",
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
		},
	}, "k8s", "1.10.0")
	collector.add(Definition{
		Kind:    "ConfigMap",
		Version: "v1",
		Schema: map[string]*schema.Schema{
			"immutable": {
				Type:     schema.TypeBool,
				Required: true,
				Computed: false,
			},
		},
	}, "k8s", "1.18.0")
	defs := collector.list()
	if len(defs) != 1 {
		t.Fatalf("expected one merged definition, got %d", len(defs))
	}
	cm := defs[0]
	if !strings.Contains(strings.Join(cm.ProviderVersions, ","), "v1.10.0") || !strings.Contains(strings.Join(cm.ProviderVersions, ","), "v1.18.0") {
		t.Fatalf("expected merged provider versions, got %v", cm.ProviderVersions)
	}
	if cm.Schema["data"] == nil {
		t.Fatalf("expected data field in merged schema")
	}
	immutable := cm.Schema["immutable"]
	if immutable == nil {
		t.Fatalf("expected immutable field in merged schema")
	}
	if immutable.Required {
		t.Fatalf("expected merged immutable field to be relaxed (not required)")
	}
	if !immutable.Optional || !immutable.Computed {
		t.Fatalf("expected merged immutable field to be optional+computed")
	}
}

func TestDefinitionCollectorMergeRelaxesRemovedRequiredField(t *testing.T) {
	collector := newDefinitionCollector()
	collector.add(Definition{
		Kind:    "Widget",
		Group:   "example.io",
		Version: "v1",
		Schema: map[string]*schema.Schema{
			"legacy_required": {
				Type:     schema.TypeString,
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, "example", "1.0.0")
	collector.add(Definition{
		Kind:    "Widget",
		Group:   "example.io",
		Version: "v1",
		Schema: map[string]*schema.Schema{
			"new_optional": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}, "example", "1.1.0")
	defs := collector.list()
	var merged *Definition
	for i := range defs {
		if defs[i].Group == "example.io" && defs[i].Kind == "Widget" {
			merged = &defs[i]
			break
		}
	}
	if merged == nil {
		t.Fatalf("expected merged definition")
	}
	legacy := merged.Schema["legacy_required"]
	if legacy == nil {
		t.Fatalf("expected legacy_required to remain present in merged schema")
	}
	if legacy.Required {
		t.Fatalf("expected removed field to be relaxed from required")
	}
	if !legacy.Optional || !legacy.Computed {
		t.Fatalf("expected removed field to be optional+computed")
	}
}

func TestMergeSchemaFieldTypeConflictPrefersNonMapType(t *testing.T) {
	existing := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	incoming := &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
		Computed: true,
	}
	merged := mergeSchemaField(existing, incoming)
	if merged.Type != schema.TypeList {
		t.Fatalf("expected merged type to prefer non-map list type, got %v", merged.Type)
	}
	if !merged.Optional || !merged.Computed || merged.Required {
		t.Fatalf("expected merged field to be optional+computed and not required")
	}
}

func TestMergeSchemaFieldTypeConflictUsesIncomingNonMapWhenExistingIsMap(t *testing.T) {
	existing := &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
		Computed: true,
	}
	incoming := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	merged := mergeSchemaField(existing, incoming)
	if merged.Type != schema.TypeList {
		t.Fatalf("expected merged type to prefer incoming non-map list type, got %v", merged.Type)
	}
	if !merged.Optional || !merged.Computed || merged.Required {
		t.Fatalf("expected merged field to be optional+computed and not required")
	}
}

func TestMergeSchemaFieldAdoptsIncomingElemWhenExistingElemNil(t *testing.T) {
	existing := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem:     nil,
	}
	incoming := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"field": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
	merged := mergeSchemaField(existing, incoming)
	elem, ok := merged.Elem.(*schema.Resource)
	if !ok || elem == nil {
		t.Fatalf("expected merged Elem to adopt incoming resource schema, got %T", merged.Elem)
	}
	if _, ok := elem.Schema["field"]; !ok {
		t.Fatalf("expected merged Elem schema to include incoming field")
	}
}

func TestMergeSchemaMapExistingNilPreservesFieldRequiredness(t *testing.T) {
	incoming := map[string]*schema.Schema{
		"required_field": {
			Type:     schema.TypeString,
			Required: true,
			Optional: false,
			Computed: false,
		},
	}
	merged := mergeSchemaMap(nil, incoming)
	field := merged["required_field"]
	if field == nil {
		t.Fatalf("expected required_field in merged schema")
	}
	if !field.Required || field.Optional || field.Computed {
		t.Fatalf("expected first-seen required field flags to be preserved, got required=%v optional=%v computed=%v", field.Required, field.Optional, field.Computed)
	}
}

func TestMergeSchemaFieldTypeConflictNonMapIsOrderIndependent(t *testing.T) {
	listSchema := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	}
	setSchema := &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Computed: true,
	}
	mergedListSet := mergeSchemaField(listSchema, setSchema)
	mergedSetList := mergeSchemaField(setSchema, listSchema)
	if mergedListSet.Type != mergedSetList.Type {
		t.Fatalf("expected deterministic non-map conflict type selection, got %v vs %v", mergedListSet.Type, mergedSetList.Type)
	}
}

func TestMergeSchemaFieldElemConflictResourceVsSchemaPrefersIncomingNonMapSchema(t *testing.T) {
	existing := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"obj": {Type: schema.TypeString, Optional: true, Computed: true},
			},
		},
	}
	incoming := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
	merged := mergeSchemaField(existing, incoming)
	elem, ok := merged.Elem.(*schema.Schema)
	if !ok || elem == nil {
		t.Fatalf("expected merged Elem to prefer incoming schema, got %T", merged.Elem)
	}
	if elem.Type != schema.TypeString {
		t.Fatalf("expected merged Elem type string, got %v", elem.Type)
	}
}

func TestMergeSchemaFieldElemConflictSchemaMapVsResourcePrefersIncomingResource(t *testing.T) {
	existing := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeMap,
		},
	}
	incoming := &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"obj": {Type: schema.TypeString, Optional: true, Computed: true},
			},
		},
	}
	merged := mergeSchemaField(existing, incoming)
	elem, ok := merged.Elem.(*schema.Resource)
	if !ok || elem == nil {
		t.Fatalf("expected merged Elem to prefer incoming resource schema, got %T", merged.Elem)
	}
	if _, ok := elem.Schema["obj"]; !ok {
		t.Fatalf("expected merged Elem resource schema to include incoming fields")
	}
}

func TestMergeSchemaFieldPreservesComputedOnlySemantics(t *testing.T) {
	existing := &schema.Schema{
		Type:     schema.TypeString,
		Required: false,
		Optional: false,
		Computed: true,
	}
	incoming := &schema.Schema{
		Type:     schema.TypeString,
		Required: false,
		Optional: false,
		Computed: true,
	}
	merged := mergeSchemaField(existing, incoming)
	if merged.Required {
		t.Fatalf("expected merged field to stay non-required")
	}
	if merged.Optional {
		t.Fatalf("expected computed-only merged field to remain non-optional")
	}
	if !merged.Computed {
		t.Fatalf("expected computed-only merged field to remain computed")
	}
}

func TestSortSchemaInputFilesOrdersBySemverThenPath(t *testing.T) {
	files := []schemaInputFile{
		{path: "/tmp/b-v1.10.0.json", providerVersion: "v1.10.0"},
		{path: "/tmp/a-v1.10.0.json", providerVersion: "v1.10.0"},
		{path: "/tmp/c-v1.9.0.json", providerVersion: "v1.9.0"},
		{path: "/tmp/d-no-version.json", providerVersion: ""},
	}
	sortSchemaInputFiles(files)
	var got []string
	for _, f := range files {
		got = append(got, f.path)
	}
	want := []string{
		"/tmp/d-no-version.json",
		"/tmp/c-v1.9.0.json",
		"/tmp/a-v1.10.0.json",
		"/tmp/b-v1.10.0.json",
	}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("unexpected sorted order: got %v want %v", got, want)
	}
}

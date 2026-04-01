package resourcegen

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	manifestJSONFieldName = "kubefu_manifest_json"
	manifestYAMLFieldName = "kubefu_manifest_yaml"
)

type PrepareDataSourceSchemaOptions struct {
	EnsureRootMetadataName bool
}

type PreparedDataSourceSchema struct {
	Schema              map[string]*schema.Schema
	ManifestKeys        []string
	ManifestObjectPaths []string
}

func PrepareDataSourceSchema(def Definition, input map[string]*schema.Schema, opts PrepareDataSourceSchemaOptions) PreparedDataSourceSchema {
	schemaMap := input
	if schemaMap == nil {
		schemaMap = map[string]*schema.Schema{}
	}
	schemaMap = copySchemaMap(schemaMap)
	applyDefinitionSchemaOverrides(def, schemaMap)
	if opts.EnsureRootMetadataName {
		ensureRootMetadataName(schemaMap)
	}
	schemaMap[manifestJSONFieldName] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Rendered manifest (canonical JSON) for this data source.",
		Computed:    true,
	}
	schemaMap[manifestYAMLFieldName] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Rendered manifest (canonical YAML) for this data source.",
		Computed:    true,
	}
	ensureComputedString(schemaMap, "api_version", "APIVersion defines the versioned schema of this representation of an object.")
	ensureComputedString(schemaMap, "kind", "Kind is a string value representing the REST resource this object represents.")

	manifestKeys := make([]string, 0, len(schemaMap))
	for key := range schemaMap {
		if key == "api_version" || key == "kind" || key == manifestJSONFieldName || key == manifestYAMLFieldName {
			continue
		}
		manifestKeys = append(manifestKeys, key)
	}
	sort.Strings(manifestKeys)
	return PreparedDataSourceSchema{
		Schema:              schemaMap,
		ManifestKeys:        manifestKeys,
		ManifestObjectPaths: objectPathsForSchema(schemaMap),
	}
}

func ensureRootMetadataName(schemaMap map[string]*schema.Schema) {
	if schemaMap == nil {
		return
	}
	metadata, ok := schemaMap["metadata"]
	if !ok || metadata == nil {
		schemaMap["metadata"] = minimalMetadataSchema()
		return
	}
	if metadata.Type != schema.TypeList {
		return
	}
	elem, ok := metadata.Elem.(*schema.Resource)
	if !ok || elem == nil {
		return
	}
	if elem.Schema == nil {
		elem.Schema = map[string]*schema.Schema{}
	}
	if _, exists := elem.Schema["name"]; exists {
		return
	}
	elem.Schema["name"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Name of the resource.",
		Optional:    true,
		Computed:    true,
	}
}

func minimalMetadataSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Standard object's metadata.",
		Optional:    true,
		Computed:    true,
		MaxItems:    1,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the resource.",
				Optional:    true,
				Computed:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "Namespace for namespaced resources.",
				Optional:    true,
				Computed:    true,
			},
			"labels": {
				Type:        schema.TypeMap,
				Description: "Map of string keys and values for organizing resources.",
				Optional:    true,
				Computed:    true,
			},
			"annotations": {
				Type:        schema.TypeMap,
				Description: "Map of string keys and values for storing arbitrary metadata.",
				Optional:    true,
				Computed:    true,
			},
		}},
	}
}

package resourcegen

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tasansga/terraform-provider-kubefu/resourcegen/resource"
)

// Definition represents a generic Terraform-ready schema that can originate from
// any Kubernetes API object, including CRDs.
type Definition struct {
	// Kind specifies the resource kind (e.g. ConfigMap, MyCustomResource).
	Kind string
	// Group identifies the API group, e.g. cert-manager.io.
	Group string
	// Version holds the API version string (v1, v1beta1, etc.).
	Version string
	// Description stores a human-readable description of the definition.
	Description string
	// DefinitionName is an optional identifier (for example the OpenAPI definition name).
	DefinitionName string
	// Schema describes the Terraform attributes for the resource.
	Schema map[string]*schema.Schema
	// Provider names the parent ecosystem (k8s, flux, cert-manager).
	Provider string
	// ProviderVersions captures which parent versions include this definition.
	ProviderVersions []string
}

// FileData contains the generated file information.
type FileData struct {
	// Name is the target file name (e.g. datasource_config_map_v1.go).
	Name string
	// Content holds the generated source code.
	Content []byte
}

// AsDataSource renders this definition as a Terraform data source file. The
// generated data source mirrors the schema stored in Definition.Schema and is
// written into a file named datasource_<provider>_<kind>_<version>.go.
func (d Definition) AsDataSource(pkgName, provider string) (FileData, error) {
	schemaMap := d.Schema
	if schemaMap == nil {
		schemaMap = map[string]*schema.Schema{}
	}
	providerName := providerSegment(provider)
	funcName := d.DataSourceFuncName(providerName)
	description := d.descriptionOrDefault()
	apiVersion := apiVersionFor(d.Group, d.Version)
	manifestJSONField := "kubefu_manifest_json"
	manifestYAMLField := "kubefu_manifest_yaml"
	schemaMap = copySchemaMap(schemaMap)
	schemaMap[manifestJSONField] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Rendered manifest (canonical JSON) for this data source.",
		Computed:    true,
	}
	schemaMap[manifestYAMLField] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Rendered manifest (canonical YAML) for this data source.",
		Computed:    true,
	}
	manifestKeys := make([]string, 0, len(schemaMap))
	for key := range schemaMap {
		if key == "api_version" || key == "kind" || key == manifestJSONField || key == manifestYAMLField {
			continue
		}
		manifestKeys = append(manifestKeys, key)
	}
	sort.Strings(manifestKeys)
	ds := &resource.DataSource{
		PackageName:        pkgName,
		FuncName:           funcName,
		Description:        description,
		Schema:             schemaMap,
		ProviderName:       d.Provider,
		CompatibleVersions: d.ProviderVersions,
		APIVersion:         apiVersion,
		Kind:               d.Kind,
		ID:                 d.ID(),
		ManifestKeys:       manifestKeys,
	}
	source, err := ds.Render()
	if err != nil {
		return FileData{}, err
	}
	return FileData{
		Name:    fmt.Sprintf("datasource_%s_%s_%s.go", providerName, toSnakeCase(d.Kind), versionSegment(d.Version)),
		Content: []byte(source),
	}, nil
}

// AsSchemaResource turns the definition into a Terraform schema.Resource that can
// be registered directly in a provider.
func (d Definition) AsSchemaResource() *schema.Resource {
	schemaMap := d.Schema
	if schemaMap == nil {
		schemaMap = map[string]*schema.Schema{}
	}
	description := d.descriptionOrDefault()
	id := d.ID()
	return &schema.Resource{
		ReadContext: func(_ context.Context, rd *schema.ResourceData, _ any) diag.Diagnostics {
			rd.SetId(id)
			return diag.Diagnostics{}
		},
		Description: description,
		Schema:      schemaMap,
	}
}

// ID returns a stable identifier for this definition that can be used for Terraform state.
func (d Definition) ID() string {
	return definitionID(d)
}

func (d Definition) descriptionOrDefault() string {
	if d.Description != "" {
		return d.Description
	}
	if d.DefinitionName != "" {
		return fmt.Sprintf("Generated data source for %s", d.DefinitionName)
	}
	return fmt.Sprintf("Generated data source for %s", d.Kind)
}

func definitionID(def Definition) string {
	group := def.Group
	if group == "" {
		group = "core"
	}
	version := def.Version
	if version == "" {
		version = "v1"
	}
	kind := def.Kind
	if kind == "" {
		kind = "resource"
	}
	return fmt.Sprintf("%s/%s/%s", group, version, kind)
}

func definitionKey(def Definition) string {
	group := def.Group
	if group == "" {
		group = "core"
	}
	kind := def.Kind
	if kind == "" {
		kind = "resource"
	}
	return fmt.Sprintf("%s/%s", group, kind)
}

func apiVersionFor(group, version string) string {
	if group == "" || group == "core" {
		return version
	}
	return fmt.Sprintf("%s/%s", group, version)
}

// DataSourceKey returns the tfplugindocs data source name exposed by Terraform.
func (d Definition) DataSourceKey(provider string) string {
	return dataSourceKey(provider, d.Group, d.Kind, d.Version)
}

// DataSourceFuncName returns the generated factory function name for this definition.
func (d Definition) DataSourceFuncName(provider string) string {
	return fmt.Sprintf("dataSource%s%s%s%s", providerCamel(provider), groupCamel(d.Group), d.Kind, versionCamel(d.Version))
}


func copySchemaMap(src map[string]*schema.Schema) map[string]*schema.Schema {
	if src == nil {
		return map[string]*schema.Schema{}
	}
	dst := make(map[string]*schema.Schema, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

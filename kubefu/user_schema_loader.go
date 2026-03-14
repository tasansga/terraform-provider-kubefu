package kubefu

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
	"github.com/tasansga/terraform-provider-kubefu/resourcegen"
)

const (
	userSchemaProviderDefault = "user"
)

func loadUserSchemas(paths []string) (map[string]*schema.Resource, diag.Diagnostics) {
	result := map[string]*schema.Resource{}
	if len(paths) == 0 {
		return result, nil
	}

	var diags diag.Diagnostics
	for _, raw := range paths {
		path := strings.TrimSpace(raw)
		if path == "" {
			continue
		}
		info, err := os.Stat(path)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to load schema path",
				Detail:   fmt.Sprintf("path %q: %s", path, err),
			})
			continue
		}
		if info.IsDir() {
			provider := normalizeProviderSegment(filepath.Base(path))
			if err := walkSchemaDir(path, provider, result, &diags); err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Failed to load schema directory",
					Detail:   err.Error(),
				})
			}
			continue
		}
		if err := loadSchemaFile(path, userSchemaProviderDefault, result, &diags); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to load schema file",
				Detail:   err.Error(),
			})
		}
	}

	return result, diags
}

func walkSchemaDir(dir, provider string, result map[string]*schema.Resource, diags *diag.Diagnostics) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if err := loadSchemaFile(path, provider, result, diags); err != nil {
			if diags != nil {
				*diags = append(*diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Failed to load schema file",
					Detail:   err.Error(),
				})
			}
			return nil
		}
		return nil
	})
}

func loadSchemaFile(path, provider string, result map[string]*schema.Resource, diags *diag.Diagnostics) error {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".json" && ext != ".yaml" && ext != ".yml" {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read schema file %s: %w", path, err)
	}
	var entries []resourcegen.ResourceEntry
	switch ext {
	case ".json":
		entries, err = resourcegen.ParseRootResources(data)
	case ".yaml", ".yml":
		entries, err = resourcegen.ParseCRDResources(data)
	}
	if err != nil {
		return fmt.Errorf("parse schema file %s: %w", path, err)
	}
	if (ext == ".yaml" || ext == ".yml") && len(entries) == 0 && diags != nil {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "No CRD definitions found",
			Detail:   fmt.Sprintf("schema file %q did not contain any CustomResourceDefinition objects", path),
		})
	}
	for _, entry := range entries {
		if entry.Kind == "" {
			continue
		}
		resource, key := buildUserSchemaDataSource(entry, provider)
		if resource == nil || key == "" {
			continue
		}
		if _, exists := result[key]; exists {
			if diags != nil {
				*diags = append(*diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Duplicate user schema ignored",
					Detail:   fmt.Sprintf("data source %q already exists (path %s)", key, path),
				})
			}
			continue
		}
		result[key] = resource
	}
	return nil
}

func buildUserSchemaDataSource(entry resourcegen.ResourceEntry, provider string) (*schema.Resource, string) {
	def := resourcegen.Definition{
		Group:   entry.Group,
		Version: entry.Version,
		Kind:    entry.Kind,
	}
	key := def.DataSourceKey(provider)
	if key == "" {
		return nil, ""
	}
	schemaMap := copySchemaMap(entry.Schema)
	ensureComputedString(schemaMap, "api_version", "APIVersion defines the versioned schema of this representation of an object.")
	ensureComputedString(schemaMap, "kind", "Kind is a string value representing the REST resource this object represents.")
	schemaMap["kubefu_manifest_json"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Rendered manifest (canonical JSON) for this data source.",
		Computed:    true,
	}
	schemaMap["kubefu_manifest_yaml"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Rendered manifest (canonical YAML) for this data source.",
		Computed:    true,
	}
	manifestKeys := make([]string, 0, len(schemaMap))
	for name := range schemaMap {
		if name == "api_version" || name == "kind" || name == "kubefu_manifest_json" || name == "kubefu_manifest_yaml" {
			continue
		}
		manifestKeys = append(manifestKeys, name)
	}
	manifestObjectPaths := objectPathsForSchema(schemaMap)
	apiVersion := apiVersionFor(entry.Group, entry.Version)
	id := def.ID()
	description := entry.DefinitionDescription
	if strings.TrimSpace(description) == "" {
		description = fmt.Sprintf("Generated data source for %s", entry.Kind)
	}
	return &schema.Resource{
		ReadContext: func(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			if err := manifest.SetDataSourceDefaults(d, apiVersion, entry.Kind, id); err != nil {
				return diag.FromErr(err)
			}
			if err := manifest.SetDataSourceManifestWithObjectPathsForMeta(d, m, manifestKeys, manifestObjectPaths); err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Description: description,
		Schema:      schemaMap,
	}, key
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

func ensureComputedString(schemaMap map[string]*schema.Schema, key, description string) {
	if schemaMap == nil {
		return
	}
	if existing, ok := schemaMap[key]; ok && existing != nil {
		clone := *existing
		clone.Required = false
		clone.Optional = false
		clone.Computed = true
		if clone.Description == "" && description != "" {
			clone.Description = description
		}
		schemaMap[key] = &clone
		return
	}
	schemaMap[key] = &schema.Schema{
		Type:        schema.TypeString,
		Description: description,
		Computed:    true,
	}
}

func objectPathsForSchema(schemaMap map[string]*schema.Schema) []string {
	if len(schemaMap) == 0 {
		return nil
	}
	paths := make(map[string]struct{})
	collectObjectPaths(schemaMap, "", paths)
	if len(paths) == 0 {
		return nil
	}
	keys := make([]string, 0, len(paths))
	for key := range paths {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func collectObjectPaths(schemaMap map[string]*schema.Schema, prefix string, paths map[string]struct{}) {
	for name, entry := range schemaMap {
		if entry == nil {
			continue
		}
		path := name
		if prefix != "" {
			path = prefix + "." + name
		}
		elem, ok := entry.Elem.(*schema.Resource)
		if entry.Type == schema.TypeList && ok {
			if entry.MaxItems == 1 {
				paths[path] = struct{}{}
			}
			if elem != nil {
				collectObjectPaths(elem.Schema, path, paths)
			}
		}
	}
}

func normalizeProviderSegment(provider string) string {
	name := sanitizeProviderInput(provider)
	name = toSnakeCase(name)
	if name == "" {
		return userSchemaProviderDefault
	}
	return name
}

func sanitizeProviderInput(input string) string {
	if input == "" {
		return ""
	}
	var b strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			continue
		}
		b.WriteRune('_')
	}
	return b.String()
}

func toSnakeCase(input string) string {
	runes := []rune(input)
	var b []rune
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 && shouldInsertUnderscore(runes, i) {
				b = append(b, '_')
			}
			b = append(b, unicode.ToLower(r))
			continue
		}
		b = append(b, r)
	}
	return string(b)
}

func shouldInsertUnderscore(runes []rune, i int) bool {
	if i == 0 {
		return false
	}
	prev := runes[i-1]
	nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
	return unicode.IsLower(prev) || nextLower
}

func apiVersionFor(group, version string) string {
	if group == "" || group == "core" {
		return version
	}
	return fmt.Sprintf("%s/%s", group, version)
}

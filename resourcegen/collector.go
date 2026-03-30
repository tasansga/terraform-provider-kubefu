package resourcegen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/mod/semver"
	"gopkg.in/yaml.v3"

	"github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

// CollectDefinitions walks the given schema root and returns the parsed
// definitions grouped by provider (subdirectory or the root). Each provider
// directory may contain Kubernetes OpenAPI JSON specs or CRD YAML manifests.
func CollectDefinitions(root string) (map[string][]Definition, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("stat schema root %s: %w", root, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", root)
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("read schema root %s: %w", root, err)
	}
	hasSubdirs := false
	for _, entry := range entries {
		if entry.IsDir() {
			hasSubdirs = true
			break
		}
	}
	results := make(map[string][]Definition)
	if hasSubdirs {
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			providerPath := filepath.Join(root, entry.Name())
			providerName := providerSegment(entry.Name())
			defs, err := collectProviderDefinitions(providerPath, providerName)
			if err != nil {
				return nil, fmt.Errorf("collect definitions for provider %s: %w", entry.Name(), err)
			}
			if len(defs) == 0 {
				continue
			}
			results[providerName] = defs
		}
	} else {
		name := filepath.Base(root)
		providerName := providerSegment(name)
		defs, err := collectProviderDefinitions(root, providerName)
		if err != nil {
			return nil, err
		}
		if len(defs) > 0 {
			results[providerName] = defs
		}
	}
	return results, nil
}

func collectProviderDefinitions(dir, provider string) ([]Definition, error) {
	collector := newDefinitionCollector()
	var files []schemaInputFile
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		files = append(files, schemaInputFile{
			path:            path,
			ext:             strings.ToLower(filepath.Ext(d.Name())),
			providerVersion: version.FromFilename(d.Name()),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sortSchemaInputFiles(files)
	for _, file := range files {
		switch file.ext {
		case ".json":
			if err := processOpenAPISpec(file.path, provider, file.providerVersion, collector); err != nil {
				return nil, fmt.Errorf("process OpenAPI spec %s: %w", file.path, err)
			}
		case ".yaml", ".yml":
			if err := processCRDFile(file.path, provider, file.providerVersion, collector); err != nil {
				return nil, fmt.Errorf("process CRD file %s: %w", file.path, err)
			}
		default:
			// ignore other files
		}
	}
	defs := collector.list()
	sort.Slice(defs, func(i, j int) bool {
		return definitionID(defs[i]) < definitionID(defs[j])
	})
	return defs, nil
}

type schemaInputFile struct {
	path            string
	ext             string
	providerVersion string
}

func sortSchemaInputFiles(files []schemaInputFile) {
	sort.Slice(files, func(i, j int) bool {
		left := version.Canonical(files[i].providerVersion)
		right := version.Canonical(files[j].providerVersion)
		if left == "" && right != "" {
			return true
		}
		if right == "" && left != "" {
			return false
		}
		if left != right {
			return semver.Compare(left, right) < 0
		}
		return files[i].path < files[j].path
	})
}

func processOpenAPISpec(path, provider, providerVersion string, collector *definitionCollector) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read spec: %w", err)
	}
	return processOpenAPISpecBytes(data, provider, providerVersion, collector)
}

func processOpenAPISpecBytes(data []byte, provider, providerVersion string, collector *definitionCollector) error {
	entries, err := ParseRootResources(data)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		collector.add(Definition{
			Kind:           entry.Kind,
			Group:          entry.Group,
			Version:        entry.Version,
			Description:    entry.DefinitionDescription,
			DefinitionName: entry.DefinitionName,
			Schema:         entry.Schema,
		}, provider, providerVersion)
	}
	return nil
}

func processCRDFile(path, provider, providerVersion string, collector *definitionCollector) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open CRD file: %w", err)
	}
	defer f.Close()
	return decodeCRD(f, provider, providerVersion, collector)
}

func definitionsFromCRD(crd customResourceDefinition) []Definition {
	if crd.Kind != "CustomResourceDefinition" {
		return nil
	}
	group := crd.Spec.Group
	kind := crd.Spec.Names.Kind
	if kind == "" {
		return nil
	}
	var versions []crdVersion
	if len(crd.Spec.Versions) > 0 {
		versions = crd.Spec.Versions
	} else if crd.Spec.Version != "" {
		versions = []crdVersion{
			{
				Name: crd.Spec.Version,
				Schema: crdVersionSchema{
					OpenAPIV3Schema: crd.Spec.Validation.OpenAPIV3Schema,
				},
			},
		}
	}
	var definitions []Definition
	for _, v := range versions {
		if v.Schema.OpenAPIV3Schema.Type == "" && len(v.Schema.OpenAPIV3Schema.Properties) == 0 {
			continue
		}
		name := fmt.Sprintf("crd.%s.%s.%s", group, v.Name, kind)
		schemaMap := buildTerraformSchema(name, v.Schema.OpenAPIV3Schema, v.Schema.OpenAPIV3Schema.Definitions)
		definitions = append(definitions, Definition{
			Kind:           kind,
			Group:          group,
			Version:        v.Name,
			Description:    v.Schema.OpenAPIV3Schema.Description,
			DefinitionName: name,
			Schema:         schemaMap,
		})
	}
	return definitions
}

type customResourceDefinition struct {
	Kind string `json:"kind" yaml:"kind"`
	Spec struct {
		Group string `json:"group" yaml:"group"`
		Names struct {
			Kind string `json:"kind" yaml:"kind"`
		} `json:"names" yaml:"names"`
		Versions   []crdVersion     `json:"versions" yaml:"versions"`
		Version    string           `json:"version" yaml:"version"`
		Validation crdVersionSchema `json:"validation" yaml:"validation"`
	} `json:"spec" yaml:"spec"`
}

type crdVersion struct {
	Name   string           `json:"name" yaml:"name"`
	Schema crdVersionSchema `json:"schema" yaml:"schema"`
}

type crdVersionSchema struct {
	OpenAPIV3Schema definition `json:"openAPIV3Schema" yaml:"openAPIV3Schema"`
}

type definitionCollector struct {
	merged map[string]Definition
}

func newDefinitionCollector() *definitionCollector {
	return &definitionCollector{
		merged: make(map[string]Definition),
	}
}

func (c *definitionCollector) add(def Definition, provider, providerVersion string) {
	if def.Kind == "" {
		return
	}
	key := definitionID(def)
	existing, ok := c.merged[key]
	if ok {
		def = mergeDefinitions(existing, def)
	}
	if provider != "" {
		def.Provider = provider
	}
	if canonical := version.Canonical(providerVersion); canonical != "" {
		def.ProviderVersions = append(def.ProviderVersions, canonical)
	}
	c.merged[key] = def
}

func mergeDefinitions(existing, incoming Definition) Definition {
	merged := existing
	if merged.Kind == "" {
		merged.Kind = incoming.Kind
	}
	if merged.Group == "" {
		merged.Group = incoming.Group
	}
	if merged.Version == "" {
		merged.Version = incoming.Version
	}
	if merged.Description == "" {
		merged.Description = incoming.Description
	}
	if merged.DefinitionName == "" {
		merged.DefinitionName = incoming.DefinitionName
	}
	merged.Schema = mergeSchemaMap(merged.Schema, incoming.Schema)
	return merged
}

func mergeSchemaMap(existing, incoming map[string]*schema.Schema) map[string]*schema.Schema {
	if existing == nil && incoming == nil {
		return nil
	}
	if existing == nil {
		out := make(map[string]*schema.Schema, len(incoming))
		for k, v := range incoming {
			out[k] = copySchema(v)
		}
		return out
	}
	out := make(map[string]*schema.Schema, len(existing))
	for k, v := range existing {
		out[k] = copySchema(v)
	}
	for k, incomingField := range incoming {
		existingField, ok := out[k]
		if !ok {
			out[k] = relaxSchemaForCompatibility(copySchema(incomingField))
			continue
		}
		out[k] = mergeSchemaField(existingField, incomingField)
	}
	// If a field disappears in a newer schema snapshot, keep it in merged output
	// but relax it so merged/default validation remains compatible with versions
	// where that field is absent.
	for k, existingField := range out {
		if _, ok := incoming[k]; ok {
			continue
		}
		out[k] = relaxSchemaForCompatibility(existingField)
	}
	return out
}

func mergeSchemaField(existing, incoming *schema.Schema) *schema.Schema {
	if existing == nil {
		return copySchema(incoming)
	}
	if incoming == nil {
		return copySchema(existing)
	}
	if existing.Type != incoming.Type {
		merged := chooseConflictSchema(existing, incoming)
		merged.Description = firstNonEmpty(existing.Description, incoming.Description)
		merged.Required = false
		merged.Optional = true
		merged.Computed = true
		merged.MinItems = 0
		merged.MaxItems = 0
		return merged
	}
	merged := copySchema(existing)
	merged.Description = firstNonEmpty(existing.Description, incoming.Description)
	merged.Required = existing.Required && incoming.Required
	merged.Optional = !merged.Required && (existing.Optional || incoming.Optional)
	merged.Computed = existing.Computed || incoming.Computed
	if !merged.Required && !merged.Optional && !merged.Computed {
		merged.Computed = true
	}
	if existing.MaxItems == 0 || incoming.MaxItems == 0 {
		merged.MaxItems = 0
	} else if existing.MaxItems > incoming.MaxItems {
		merged.MaxItems = existing.MaxItems
	} else {
		merged.MaxItems = incoming.MaxItems
	}
	if existing.MinItems == 0 || incoming.MinItems == 0 {
		merged.MinItems = 0
	} else if existing.MinItems < incoming.MinItems {
		merged.MinItems = existing.MinItems
	} else {
		merged.MinItems = incoming.MinItems
	}
	if existing.Elem == nil && incoming.Elem != nil {
		switch incomingElem := incoming.Elem.(type) {
		case *schema.Resource:
			if incomingElem != nil {
				merged.Elem = &schema.Resource{Schema: copySchemaMap(incomingElem.Schema)}
			}
		case *schema.Schema:
			merged.Elem = copySchema(incomingElem)
		default:
			merged.Elem = incoming.Elem
		}
		return merged
	}
	switch existingElem := existing.Elem.(type) {
	case *schema.Resource:
		incomingElem, ok := incoming.Elem.(*schema.Resource)
		if ok && existingElem != nil && incomingElem != nil {
			merged.Elem = &schema.Resource{
				Schema: mergeSchemaMap(existingElem.Schema, incomingElem.Schema),
			}
		}
		incomingElemSchema, ok := incoming.Elem.(*schema.Schema)
		if ok && existingElem != nil && incomingElemSchema != nil {
			if incomingElemSchema.Type != schema.TypeMap {
				chosen := copySchema(incomingElemSchema)
				chosen.Optional = true
				chosen.Required = false
				chosen.Computed = true
				merged.Elem = chosen
			} else {
				merged.Elem = &schema.Resource{Schema: copySchemaMap(existingElem.Schema)}
			}
		}
	case *schema.Schema:
		incomingElem, ok := incoming.Elem.(*schema.Schema)
		if ok && existingElem != nil && incomingElem != nil && existingElem.Type != incomingElem.Type {
			chosen := chooseConflictElem(existingElem, incomingElem)
			chosen.Optional = true
			chosen.Required = false
			chosen.Computed = true
			merged.Elem = chosen
		}
		incomingElemResource, ok := incoming.Elem.(*schema.Resource)
		if ok && existingElem != nil && incomingElemResource != nil {
			if existingElem.Type == schema.TypeMap {
				merged.Elem = &schema.Resource{Schema: copySchemaMap(incomingElemResource.Schema)}
			} else {
				chosen := copySchema(existingElem)
				chosen.Optional = true
				chosen.Required = false
				chosen.Computed = true
				merged.Elem = chosen
			}
		}
	}
	return merged
}

func chooseConflictSchema(existing, incoming *schema.Schema) *schema.Schema {
	if existing != nil && incoming != nil {
		if existing.Type == schema.TypeMap && incoming.Type != schema.TypeMap {
			return copySchema(incoming)
		}
		if incoming.Type == schema.TypeMap && existing.Type != schema.TypeMap {
			return copySchema(existing)
		}
		if conflictTypeRank(existing.Type) >= conflictTypeRank(incoming.Type) {
			return copySchema(existing)
		}
		return copySchema(incoming)
	}
	return copySchema(existing)
}

func chooseConflictElem(existing, incoming *schema.Schema) *schema.Schema {
	if existing != nil && incoming != nil {
		if existing.Type == schema.TypeMap && incoming.Type != schema.TypeMap {
			return copySchema(incoming)
		}
		if incoming.Type == schema.TypeMap && existing.Type != schema.TypeMap {
			return copySchema(existing)
		}
		if conflictTypeRank(existing.Type) >= conflictTypeRank(incoming.Type) {
			return copySchema(existing)
		}
		return copySchema(incoming)
	}
	return copySchema(existing)
}

func conflictTypeRank(valueType schema.ValueType) int {
	switch valueType {
	case schema.TypeSet:
		return 60
	case schema.TypeList:
		return 50
	case schema.TypeString:
		return 40
	case schema.TypeFloat:
		return 30
	case schema.TypeInt:
		return 20
	case schema.TypeBool:
		return 10
	case schema.TypeMap:
		return 0
	default:
		return -1
	}
}

func relaxSchemaForCompatibility(sch *schema.Schema) *schema.Schema {
	if sch == nil {
		return nil
	}
	out := copySchema(sch)
	out.Required = false
	out.Optional = true
	out.Computed = true
	switch elem := out.Elem.(type) {
	case *schema.Resource:
		if elem == nil {
			return out
		}
		relaxed := make(map[string]*schema.Schema, len(elem.Schema))
		for k, v := range elem.Schema {
			relaxed[k] = relaxSchemaForCompatibility(v)
		}
		out.Elem = &schema.Resource{Schema: relaxed}
	}
	return out
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func (c *definitionCollector) list() []Definition {
	var defs []Definition
	for _, def := range c.merged {
		def.ProviderVersions = version.NormalizeList(def.ProviderVersions)
		defs = append(defs, def)
	}
	return defs
}

func decodeCRD(r io.Reader, provider, providerVersion string, collector *definitionCollector) error {
	decoder := yaml.NewDecoder(r)
	for {
		var crd customResourceDefinition
		if err := decoder.Decode(&crd); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("decode CRD document: %w", err)
		}
		defs := definitionsFromCRD(crd)
		for _, def := range defs {
			collector.add(def, provider, providerVersion)
		}
	}
	return nil
}

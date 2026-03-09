package resourcegen

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	namePattern        = regexp.MustCompile(`\{name\}`)
	doubleParamPattern = regexp.MustCompile(`/\{[^}]+\}/\{[^}]+\}`)
	suffixPattern      = regexp.MustCompile(`/(status|scale|proxy|bind|exec|attach|log|portforward|rename|approval)$`)
	apisVersionPattern = regexp.MustCompile(`/apis/[^/]+\.[^/]+/(v\d+[\w\.]*)/`)
	apiVersionPattern  = regexp.MustCompile(`/api/v\d+/`)
	errMissingGVK      = errors.New("missing x-kubernetes-group-version-kind")
)

// ResourceEntry describes a Kubernetes root resource that can be generated.
type ResourceEntry struct {
	Path                  string
	Group                 string
	Version               string
	Kind                  string
	DefinitionName        string
	DefinitionDescription string
	Schema                map[string]*schema.Schema
}

// ParseRootResources returns the Kubernetes root resources described in the
// OpenAPI document, applying the jq filter from the DESIGN note and returning
// resource metadata augmented with a Terraform schema for the definition.
func ParseRootResources(data []byte) ([]ResourceEntry, error) {
	var doc struct {
		Paths       map[string]json.RawMessage `json:"paths"`
		Definitions map[string]definition      `json:"definitions"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("unmarshal OpenAPI document: %w", err)
	}
	if len(doc.Paths) == 0 {
		return parseDefinitionsOnly(doc.Definitions)
	}
	lookup := buildDefinitionLookup(doc.Definitions)
	var result []ResourceEntry
	for path, raw := range doc.Paths {
		if !isRootPath(path) {
			continue
		}
		entry, err := parsePath(path, raw, doc.Definitions, lookup)
		if err != nil {
			if errors.Is(err, errMissingGVK) {
				continue
			}
			return nil, err
		}
		result = append(result, entry)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Path < result[j].Path
	})
	return result, nil
}

func parseDefinitionsOnly(definitions map[string]definition) ([]ResourceEntry, error) {
	if len(definitions) == 0 {
		return nil, fmt.Errorf("missing definitions section")
	}
	seen := map[string]struct{}{}
	var result []ResourceEntry
	for name, def := range definitions {
		for i := range def.XGroupVersionKind {
			gvk := def.XGroupVersionKind[i]
			key := gvkKey(&groupVersionKind{
				Group:   gvk.Group,
				Version: gvk.Version,
				Kind:    gvk.Kind,
			})
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			result = append(result, ResourceEntry{
				Group:                 gvk.Group,
				Version:               gvk.Version,
				Kind:                  gvk.Kind,
				DefinitionName:        name,
				DefinitionDescription: def.Description,
				Schema:                buildTerraformSchema(name, def, definitions),
			})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Group == result[j].Group {
			if result[i].Version == result[j].Version {
				return result[i].Kind < result[j].Kind
			}
			return result[i].Version < result[j].Version
		}
		return result[i].Group < result[j].Group
	})
	return result, nil
}

var (
	methodOrder = []func(*pathItem) *operation{
		func(p *pathItem) *operation { return p.Get },
		func(p *pathItem) *operation { return p.Post },
		func(p *pathItem) *operation { return p.Put },
		func(p *pathItem) *operation { return p.Patch },
		func(p *pathItem) *operation { return p.Delete },
		func(p *pathItem) *operation { return p.Options },
		func(p *pathItem) *operation { return p.Head },
		func(p *pathItem) *operation { return p.Connect },
	}
)

func parsePath(path string, raw json.RawMessage, definitions map[string]definition, lookup map[string]string) (ResourceEntry, error) {
	var item pathItem
	if err := json.Unmarshal(raw, &item); err != nil {
		return ResourceEntry{}, fmt.Errorf("parse path %s: %w", path, err)
	}
	gvk := item.groupVersionKind()
	if gvk == nil {
		return ResourceEntry{}, fmt.Errorf("path %s: %w", path, errMissingGVK)
	}
	defName := resolveDefinitionName(definitions, gvk, lookup)
	def, ok := definitions[defName]
	if !ok {
		return ResourceEntry{}, fmt.Errorf("definition %s not found for path %s", defName, path)
	}
	return ResourceEntry{
		Path:                  path,
		Group:                 gvk.Group,
		Version:               gvk.Version,
		Kind:                  gvk.Kind,
		DefinitionName:        defName,
		DefinitionDescription: def.Description,
		Schema:                buildTerraformSchema(defName, def, definitions),
	}, nil
}

type pathItem struct {
	Get     *operation `json:"get"`
	Post    *operation `json:"post"`
	Put     *operation `json:"put"`
	Patch   *operation `json:"patch"`
	Delete  *operation `json:"delete"`
	Options *operation `json:"options"`
	Head    *operation `json:"head"`
	Connect *operation `json:"connect"`
}

func (p pathItem) groupVersionKind() *groupVersionKind {
	for _, getter := range methodOrder {
		if op := getter(&p); op != nil && op.XGroupVersionKind != nil {
			return op.XGroupVersionKind
		}
	}
	return nil
}

type operation struct {
	XGroupVersionKind *groupVersionKind `json:"x-kubernetes-group-version-kind"`
}

type groupVersionKind struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

func formatDefinitionName(gvk *groupVersionKind) string {
	group := gvk.Group
	if group == "" {
		group = "core"
	}
	return fmt.Sprintf("io.k8s.api.%s.%s.%s", group, gvk.Version, gvk.Kind)
}

func resolveDefinitionName(definitions map[string]definition, gvk *groupVersionKind, lookup map[string]string) string {
	key := gvkKey(gvk)
	if name, ok := lookup[key]; ok {
		return name
	}
	if name, ok := findDefinitionNameByGVK(definitions, gvk); ok {
		return name
	}
	if name, ok := findDefinitionBySuffix(definitions, gvk); ok {
		return name
	}
	return formatDefinitionName(gvk)
}

func gvkKey(gvk *groupVersionKind) string {
	group := gvk.Group
	if group == "" {
		group = "core"
	}
	return fmt.Sprintf("%s/%s/%s", group, gvk.Version, gvk.Kind)
}

func findDefinitionNameByGVK(definitions map[string]definition, gvk *groupVersionKind) (string, bool) {
	for name, def := range definitions {
		for i := range def.XGroupVersionKind {
			if gvkEqual(gvk, &def.XGroupVersionKind[i]) {
				return name, true
			}
		}
	}
	return "", false
}

func gvkEqual(a, b *groupVersionKind) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Group == b.Group && a.Version == b.Version && a.Kind == b.Kind
}

func findDefinitionBySuffix(definitions map[string]definition, gvk *groupVersionKind) (string, bool) {
	suffix := fmt.Sprintf(".%s.%s", gvk.Version, gvk.Kind)
	var candidate string
	for name := range definitions {
		if !strings.HasSuffix(name, suffix) {
			continue
		}
		if containsGroupSegment(name, gvk.Group) {
			return name, true
		}
		if candidate == "" {
			candidate = name
		}
	}
	if candidate != "" {
		return candidate, true
	}
	return "", false
}

func containsGroupSegment(name, group string) bool {
	if group == "" {
		return false
	}
	if strings.Contains(name, group) {
		return true
	}
	for _, segment := range strings.Split(group, ".") {
		if segment == "" {
			continue
		}
		if strings.Contains(name, segment) {
			return true
		}
	}
	return false
}

type definition struct {
	Description       string                `json:"description"`
	Type              string                `json:"type"`
	Properties        map[string]property   `json:"properties"`
	Required          []string              `json:"required"`
	XGroupVersionKind []groupVersionKind    `json:"x-kubernetes-group-version-kind"`
	Definitions       map[string]definition `json:"definitions"`
}

type property struct {
	Description string              `json:"description"`
	Type        string              `json:"type"`
	Format      string              `json:"format"`
	Items       *property           `json:"items"`
	Ref         string              `json:"$ref"`
	Properties  map[string]property `json:"properties"`
}

var jsonToSchemaType = map[string]schema.ValueType{
	"string":  schema.TypeString,
	"integer": schema.TypeInt,
	"number":  schema.TypeFloat,
	"boolean": schema.TypeBool,
	"array":   schema.TypeList,
	"object":  schema.TypeMap,
}

func buildTerraformSchema(defName string, def definition, definitions map[string]definition) map[string]*schema.Schema {
	builder := &schemaBuilder{
		definitions: definitions,
		cache:       make(map[string]map[string]*schema.Schema),
	}
	return builder.buildDefinition(defName, def)
}

type schemaBuilder struct {
	definitions map[string]definition
	cache       map[string]map[string]*schema.Schema
}

func (b *schemaBuilder) buildDefinition(name string, def definition) map[string]*schema.Schema {
	if sch, ok := b.cache[name]; ok {
		return sch
	}
	placeholder := make(map[string]*schema.Schema)
	b.cache[name] = placeholder
	schemaMap := b.buildProperties(def)
	b.cache[name] = schemaMap
	return schemaMap
}

func (b *schemaBuilder) buildProperties(def definition) map[string]*schema.Schema {
	required := make(map[string]bool, len(def.Required))
	for _, prop := range def.Required {
		required[normalizedPropertyName(prop)] = true
	}
	schemaMap := make(map[string]*schema.Schema, len(def.Properties))
	for name, prop := range def.Properties {
		attr := normalizedPropertyName(name)
		if attr == "" {
			attr = name
		}
		schemaMap[attr] = b.buildSchemaForProperty(name, prop, required[attr])
	}
	return schemaMap
}

func (b *schemaBuilder) buildSchemaForProperty(name string, prop property, isRequired bool) *schema.Schema {
	sch := &schema.Schema{
		Description: prop.Description,
		Optional:    !isRequired,
		Required:    isRequired,
		Computed:    !isRequired,
	}
	if prop.Type == "array" {
		sch.Type = schema.TypeList
		sch.Elem = b.buildElemFromProperty(prop.Items)
		return sch
	}
	if prop.Ref != "" {
		sch.Type = schema.TypeMap
		// Map of unknown/complex object. Avoid invalid TypeMap Elem *Resource.
		return sch
	}
	sch.Type = mapPropertyToSchemaType(prop)
	if prop.Type == "object" {
		// Object properties render as a TypeMap with free-form values.
	}
	return sch
}

func (b *schemaBuilder) buildElemFromProperty(p *property) interface{} {
	if p == nil {
		return &schema.Resource{Schema: map[string]*schema.Schema{}}
	}
	if p.Ref != "" {
		return b.buildElemFromRef(p.Ref)
	}
	switch p.Type {
	case "object":
		if len(p.Properties) > 0 {
			return &schema.Resource{Schema: b.buildInlineProperties(p.Properties)}
		}
		return &schema.Resource{Schema: map[string]*schema.Schema{}}
	case "array":
		return &schema.Resource{Schema: map[string]*schema.Schema{}}
	default:
		return &schema.Schema{Type: mapPropertyToSchemaType(*p)}
	}
}

func (b *schemaBuilder) buildElemFromRef(ref string) interface{} {
	name, def, ok := b.lookupDefinition(ref)
	if !ok {
		return &schema.Resource{Schema: map[string]*schema.Schema{}}
	}
	return &schema.Resource{Schema: b.buildDefinition(name, def)}
}

func (b *schemaBuilder) lookupDefinition(ref string) (string, definition, bool) {
	const prefix = "#/definitions/"
	if !strings.HasPrefix(ref, prefix) {
		return "", definition{}, false
	}
	name := strings.TrimPrefix(ref, prefix)
	def, ok := b.definitions[name]
	return name, def, ok
}

func (b *schemaBuilder) buildInlineProperties(props map[string]property) map[string]*schema.Schema {
	schemaMap := make(map[string]*schema.Schema, len(props))
	for name, prop := range props {
		attr := normalizedPropertyName(name)
		if attr == "" {
			attr = name
		}
		schemaMap[attr] = b.buildSchemaForProperty(name, prop, false)
	}
	return schemaMap
}

func normalizedPropertyName(name string) string {
	if name == "" {
		return ""
	}
	normalized := toSnakeCase(name)
	if normalized == "" {
		return name
	}
	if isReservedTerraformName(normalized) {
		return normalized + "_"
	}
	return normalized
}

func isReservedTerraformName(name string) bool {
	switch name {
	case "count", "for_each", "depends_on", "provider", "providers", "lifecycle", "connection", "provisioner":
		return true
	default:
		return false
	}
}

func mapPropertyToSchemaType(p property) schema.ValueType {
	if valueType, ok := jsonToSchemaType[p.Type]; ok {
		return valueType
	}
	if p.Ref != "" {
		return schema.TypeMap
	}
	return schema.TypeString
}

func buildDefinitionLookup(definitions map[string]definition) map[string]string {
	lookup := make(map[string]string, len(definitions))
	for name, def := range definitions {
		for _, gvk := range def.XGroupVersionKind {
			key := gvkKey(&gvk)
			if key != "" {
				lookup[key] = name
			}
		}
	}
	return lookup
}

func isRootPath(path string) bool {
	if !(strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/apis/")) {
		return false
	}
	if namePattern.MatchString(path) ||
		doubleParamPattern.MatchString(path) ||
		strings.Contains(path, "/watch") ||
		strings.Contains(path, "/proxy") ||
		strings.Contains(path, "/$/") ||
		suffixPattern.MatchString(path) {
		return false
	}
	if strings.HasSuffix(path, "/") {
		return false
	}
	if !apisVersionPattern.MatchString(path) && !apiVersionPattern.MatchString(path) {
		return false
	}
	return namespacesOK(path)
}

func namespacesOK(path string) bool {
	const marker = "/namespaces/"
	idx := strings.Index(path, marker)
	if idx == -1 {
		return true
	}
	after := path[idx+len(marker):]
	if strings.HasPrefix(after, "{namespace}/") {
		return true
	}
	if after != "" && after[0] != '{' {
		return true
	}
	return false
}

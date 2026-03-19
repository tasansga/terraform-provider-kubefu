package manifest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/yaml"
)

const (
	RenderModeCompact   = "compact"
	RenderModeCanonical = "canonical"
)

type renderModeProvider interface {
	ManifestRenderMode() string
}

type manifestLiteralValue struct {
	Value interface{}
}

// SetDataSourceDefaults applies default apiVersion/kind/id to the datasource.
func SetDataSourceDefaults(d *schema.ResourceData, apiVersion, kind, id string) error {
	if apiVersion != "" {
		if v, ok := d.GetOk("api_version"); !ok || strings.TrimSpace(v.(string)) == "" {
			if err := d.Set("api_version", apiVersion); err != nil {
				return fmt.Errorf("set api_version: %w", err)
			}
		}
	}
	if kind != "" {
		if v, ok := d.GetOk("kind"); !ok || strings.TrimSpace(v.(string)) == "" {
			if err := d.Set("kind", kind); err != nil {
				return fmt.Errorf("set kind: %w", err)
			}
		}
	}
	if d.Id() == "" {
		d.SetId(id)
	}
	return nil
}

// SetDataSourceManifest renders kubefu_manifest_json/yaml from known schema keys.
func SetDataSourceManifest(d *schema.ResourceData, keys []string) error {
	return SetDataSourceManifestWithObjectPaths(d, keys, nil)
}

// SetDataSourceManifestWithObjectKeys is kept for backward compatibility.
func SetDataSourceManifestWithObjectKeys(d *schema.ResourceData, keys []string, objectKeys []string) error {
	return SetDataSourceManifestWithObjectPaths(d, keys, objectKeys)
}

// SetDataSourceManifestWithObjectPaths renders kubefu_manifest_json/yaml from known schema keys,
// treating paths in objectPaths as single-object attributes.
func SetDataSourceManifestWithObjectPaths(d *schema.ResourceData, keys []string, objectPaths []string) error {
	return setDataSourceManifestWithObjectPathsAndMode(d, keys, objectPaths, RenderModeCompact)
}

// SetDataSourceManifestWithObjectPathsForMeta renders kubefu_manifest_json/yaml from known schema keys,
// selecting render mode from provider metadata when available.
func SetDataSourceManifestWithObjectPathsForMeta(d *schema.ResourceData, meta any, keys []string, objectPaths []string) error {
	mode := RenderModeCompact
	if cfg, ok := meta.(renderModeProvider); ok {
		mode = normalizeRenderMode(cfg.ManifestRenderMode())
	}
	return setDataSourceManifestWithObjectPathsAndMode(d, keys, objectPaths, mode)
}

func setDataSourceManifestWithObjectPathsAndMode(d *schema.ResourceData, keys []string, objectPaths []string, mode string) error {
	manifest := make(map[string]interface{})
	objectPathSet := make(map[string]struct{}, len(objectPaths))
	for _, path := range objectPaths {
		if path == "" {
			continue
		}
		objectPathSet[path] = struct{}{}
	}
	explicitPaths := explicitManifestPaths(d, keys)
	if v, ok := d.GetOk("api_version"); ok {
		if s := strings.TrimSpace(v.(string)); s != "" {
			manifest["apiVersion"] = s
		}
	}
	if v, ok := d.GetOk("kind"); ok {
		if s := strings.TrimSpace(v.(string)); s != "" {
			manifest["kind"] = s
		}
	}
	for _, key := range keys {
		if key == "" {
			continue
		}
		v, ok := d.GetOkExists(key)
		if !ok {
			continue
		}
		normalized, err := normalizeManifestValue(v, key, objectPathSet, explicitPaths)
		if err != nil {
			return fmt.Errorf("normalize manifest value %q: %w", key, err)
		}
		pruned, keep := pruneManifestValue(normalized, key, explicitPaths, objectPathSet, mode)
		if keep {
			manifest[toLowerCamel(key)] = pruned
		}
	}
	if len(manifest) == 0 {
		return nil
	}
	sorted := sortManifestValue(manifest)
	jsonPayload, err := json.Marshal(sorted)
	if err != nil {
		return fmt.Errorf("marshal manifest json: %w", err)
	}
	if err := d.Set("kubefu_manifest_json", string(jsonPayload)); err != nil {
		return fmt.Errorf("set kubefu_manifest_json: %w", err)
	}
	yamlPayload, err := yaml.JSONToYAML(jsonPayload)
	if err != nil {
		return fmt.Errorf("marshal manifest yaml: %w", err)
	}
	if err := d.Set("kubefu_manifest_yaml", string(yamlPayload)); err != nil {
		return fmt.Errorf("set kubefu_manifest_yaml: %w", err)
	}
	return nil
}

func normalizeRenderMode(mode string) string {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case RenderModeCanonical:
		return RenderModeCanonical
	default:
		return RenderModeCompact
	}
}

func pruneManifestValue(value interface{}, path string, explicitPaths map[string]struct{}, objectPaths map[string]struct{}, mode string) (interface{}, bool) {
	if normalizeRenderMode(mode) == RenderModeCanonical {
		switch v := value.(type) {
		case manifestLiteralValue:
			return v.Value, v.Value != nil
		case map[string]interface{}:
			out := make(map[string]interface{}, len(v))
			_, isObjectPath := objectPaths[path]
			for key, item := range v {
				childKey := key
				if isObjectPath {
					childKey = resolveObjectPathChildKey(path, key, explicitPaths, objectPaths)
				}
				childPath := childKey
				if path != "" {
					childPath = path + "." + childKey
				}
				next, keep := pruneManifestValue(item, childPath, explicitPaths, objectPaths, mode)
				if keep {
					out[key] = next
				}
			}
			return out, true
		case []interface{}:
			out := make([]interface{}, 0, len(v))
			for _, item := range v {
				next, keep := pruneManifestValue(item, path, explicitPaths, objectPaths, mode)
				if keep {
					if next == nil {
						continue
					}
					out = append(out, next)
				}
			}
			return out, true
		default:
			return value, value != nil
		}
	}
	_, explicit := explicitPaths[path]
	switch v := value.(type) {
	case manifestLiteralValue:
		if v.Value == nil {
			return nil, explicit
		}
		return v.Value, true
	}
	switch v := value.(type) {
	case nil:
		return nil, explicit
	case string:
		return v, explicit || strings.TrimSpace(v) != ""
	case bool:
		return v, explicit || v
	case []interface{}:
		pruned := make([]interface{}, 0, len(v))
		for _, item := range v {
			next, keep := pruneManifestValue(item, path, explicitPaths, objectPaths, mode)
			if keep {
				if next == nil {
					continue
				}
				pruned = append(pruned, next)
			}
		}
		if len(pruned) == 0 {
			if explicit {
				return []interface{}{}, true
			}
			return nil, false
		}
		return pruned, true
	case map[string]interface{}:
		pruned := make(map[string]interface{}, len(v))
		_, isObjectPath := objectPaths[path]
		for key, item := range v {
			childKey := key
			if isObjectPath {
				childKey = resolveObjectPathChildKey(path, key, explicitPaths, objectPaths)
			}
			childPath := childKey
			if path != "" {
				childPath = path + "." + childKey
			}
			next, keep := pruneManifestValue(item, childPath, explicitPaths, objectPaths, mode)
			if keep {
				pruned[key] = next
			}
		}
		if len(pruned) == 0 {
			if explicit {
				return map[string]interface{}{}, true
			}
			return nil, false
		}
		return pruned, true
	default:
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return value, explicit || rv.Int() != 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return value, explicit || rv.Uint() != 0
		case reflect.Float32, reflect.Float64:
			return value, explicit || rv.Float() != 0
		}
		return value, true
	}
}

func explicitManifestPaths(d *schema.ResourceData, keys []string) map[string]struct{} {
	paths := make(map[string]struct{})
	for _, key := range keys {
		if key == "" {
			continue
		}
		rawValue, rawDiags := d.GetRawConfigAt(cty.Path{cty.GetAttrStep{Name: key}})
		if hasDiagErrors(rawDiags) {
			continue
		}
		collectExplicitManifestPaths(rawValue, key, paths, false)
	}
	return paths
}

func hasDiagErrors(diags diag.Diagnostics) bool {
	for _, d := range diags {
		if d.Severity == diag.Error {
			return true
		}
	}
	return false
}

func collectExplicitManifestPaths(value cty.Value, path string, paths map[string]struct{}, parentExplicit bool) bool {
	if path == "" || !value.IsKnown() || value.IsNull() {
		return false
	}

	t := value.Type()
	switch {
	case t.IsObjectType() || t.IsMapType():
		anyExplicit := false
		it := value.ElementIterator()
		for it.Next() {
			key, next := it.Element()
			childPath := key.AsString()
			if path != "" {
				childPath = path + "." + childPath
			}
			if collectExplicitManifestPaths(next, childPath, paths, parentExplicit) {
				anyExplicit = true
			}
		}
		if anyExplicit || parentExplicit || isTopLevelManifestPath(path) {
			paths[path] = struct{}{}
			return true
		}
		return false
	case t.IsListType() || t.IsSetType() || t.IsTupleType():
		anyExplicit := false
		hasElements := false
		it := value.ElementIterator()
		for it.Next() {
			hasElements = true
			_, next := it.Element()
			if collectExplicitManifestPaths(next, path, paths, true) {
				anyExplicit = true
			}
		}
		// Keep non-empty collections explicit. Empty nested collections are treated
		// as implicit unless they are top-level attributes.
		if anyExplicit || hasElements || isTopLevelManifestPath(path) {
			paths[path] = struct{}{}
			return true
		}
		return false
	default:
		paths[path] = struct{}{}
		return true
	}
}

func isTopLevelManifestPath(path string) bool {
	return path != "" && !strings.Contains(path, ".")
}

func lowerCamelToSnake(value string) string {
	if value == "" {
		return value
	}
	var b strings.Builder
	b.Grow(len(value) + 4)
	for i, r := range value {
		if i > 0 && r >= 'A' && r <= 'Z' {
			b.WriteByte('_')
		}
		if r >= 'A' && r <= 'Z' {
			b.WriteRune(r + ('a' - 'A'))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func resolveObjectPathChildKey(parentPath, renderedKey string, explicitPaths map[string]struct{}, objectPaths map[string]struct{}) string {
	base := lowerCamelToSnake(renderedKey)
	if parentPath == "" {
		return base
	}
	basePath := parentPath + "." + base
	if _, ok := explicitPaths[basePath]; ok {
		return base
	}
	if _, ok := objectPaths[basePath]; ok {
		return base
	}
	escaped := base + "_"
	escapedPath := parentPath + "." + escaped
	if _, ok := explicitPaths[escapedPath]; ok {
		return escaped
	}
	if _, ok := objectPaths[escapedPath]; ok {
		return escaped
	}
	return base
}

func normalizeManifestValue(value interface{}, path string, objectPaths map[string]struct{}, explicitPaths map[string]struct{}) (interface{}, error) {
	switch v := value.(type) {
	case []interface{}:
		if _, ok := objectPaths[path]; ok && len(v) == 1 {
			if m, ok := v[0].(map[string]interface{}); ok {
				return normalizeManifestValue(m, path, objectPaths, explicitPaths)
			}
			if v[0] == nil {
				return map[string]interface{}{}, nil
			}
		}
		normalized := make([]interface{}, len(v))
		for i := range v {
			next, err := normalizeManifestValue(v[i], path, objectPaths, explicitPaths)
			if err != nil {
				return nil, err
			}
			normalized[i] = next
		}
		return normalized, nil
	case map[string]interface{}:
		normalized := make(map[string]interface{}, len(v))
		_, isObjectPath := objectPaths[path]
		for k, child := range v {
			childPath := k
			if path != "" {
				childPath = path + "." + k
			}
			outKey := k
			if isObjectPath {
				outKey = toLowerCamel(k)
				if k == "values_yaml" {
					outKey = "values"
				}
			}
			next, err := normalizeManifestValue(child, childPath, objectPaths, explicitPaths)
			if err != nil {
				return nil, err
			}
			normalized[outKey] = next
		}
		return normalized, nil
	case string:
		if pathHasValuesYAMLSuffix(path) {
			var parsed interface{}
			if strings.TrimSpace(v) == "" {
				if _, explicit := explicitPaths[path]; !explicit {
					return nil, nil
				}
				return manifestLiteralValue{Value: map[string]interface{}{}}, nil
			}
			if err := yaml.Unmarshal([]byte(v), &parsed); err != nil {
				return nil, fmt.Errorf("parse values_yaml as YAML: %w", err)
			}
			obj, ok := parsed.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("parse values_yaml as YAML: expected YAML object at %q, got %T", path, parsed)
			}
			return manifestLiteralValue{Value: obj}, nil
		}
		return v, nil
	default:
		return v, nil
	}
}

func pathHasValuesYAMLSuffix(path string) bool {
	if path == "values_yaml" {
		return true
	}
	return strings.HasSuffix(path, ".values_yaml")
}

func toLowerCamel(value string) string {
	if value == "" {
		return value
	}
	parts := strings.Split(value, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		if i == 0 {
			parts[i] = strings.ToLower(part[:1]) + part[1:]
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, "")
}

func sortManifestValue(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		sorted := make(map[string]interface{}, len(v))
		for _, k := range keys {
			sorted[k] = sortManifestValue(v[k])
		}
		return sorted
	case []interface{}:
		sorted := make([]interface{}, len(v))
		for i := range v {
			sorted[i] = sortManifestValue(v[i])
		}
		return sorted
	default:
		return v
	}
}

package manifest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

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
		v, ok := d.GetOk(key)
		if !ok {
			continue
		}
		normalized := normalizeManifestValue(v, key, objectPathSet)
		pruned, keep := pruneManifestValue(normalized, mode)
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

func pruneManifestValue(value interface{}, mode string) (interface{}, bool) {
	if normalizeRenderMode(mode) == RenderModeCanonical {
		return value, value != nil
	}
	switch v := value.(type) {
	case nil:
		return nil, false
	case string:
		return v, strings.TrimSpace(v) != ""
	case bool:
		return v, v
	case []interface{}:
		pruned := make([]interface{}, 0, len(v))
		for _, item := range v {
			next, keep := pruneManifestValue(item, mode)
			if keep {
				pruned = append(pruned, next)
			}
		}
		if len(pruned) == 0 {
			return nil, false
		}
		return pruned, true
	case map[string]interface{}:
		pruned := make(map[string]interface{}, len(v))
		for key, item := range v {
			next, keep := pruneManifestValue(item, mode)
			if keep {
				pruned[key] = next
			}
		}
		if len(pruned) == 0 {
			return nil, false
		}
		return pruned, true
	default:
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return value, rv.Int() != 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return value, rv.Uint() != 0
		case reflect.Float32, reflect.Float64:
			return value, rv.Float() != 0
		}
		return value, true
	}
}

func normalizeManifestValue(value interface{}, path string, objectPaths map[string]struct{}) interface{} {
	switch v := value.(type) {
	case []interface{}:
		if _, ok := objectPaths[path]; ok && len(v) == 1 {
			if m, ok := v[0].(map[string]interface{}); ok {
				return normalizeManifestValue(m, path, objectPaths)
			}
		}
		normalized := make([]interface{}, len(v))
		for i := range v {
			normalized[i] = normalizeManifestValue(v[i], path, objectPaths)
		}
		return normalized
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
			}
			normalized[outKey] = normalizeManifestValue(child, childPath, objectPaths)
		}
		return normalized
	default:
		return v
	}
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

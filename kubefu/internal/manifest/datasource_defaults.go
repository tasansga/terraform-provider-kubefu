package manifest

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/yaml"
)

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
	manifest := make(map[string]interface{})
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
		manifest[toLowerCamel(key)] = v
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

package kubefu

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProviderRegisteredUserSchemasDataSource(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "foo.yaml")
	if err := os.WriteFile(path, []byte(sampleCRD), 0o644); err != nil {
		t.Fatalf("write CRD: %v", err)
	}
	t.Setenv("KUBEFU_SCHEMA_PATHS", path)

	p := Provider()
	ds, ok := p.DataSourcesMap[registeredUserSchemasDataSourceName]
	if !ok {
		t.Fatalf("expected %s data source", registeredUserSchemasDataSourceName)
	}

	d := ds.Data(nil)
	diags := ds.ReadContext(context.Background(), d, nil)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %+v", diags)
	}

	raw := d.Get("names")
	names, ok := raw.([]interface{})
	if !ok {
		t.Fatalf("expected names list, got %T", raw)
	}
	if len(names) != 1 {
		t.Fatalf("expected one registered user schema, got %d", len(names))
	}
	if names[0] != "kubefu_user_crd_custom_test_tst_queue_v1beta1" {
		t.Fatalf("unexpected data source name: %v", names[0])
	}
}

func TestProviderConfigureReportsRegisteredUserSchemas(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "foo.yaml")
	if err := os.WriteFile(path, []byte(sampleCRD), 0o644); err != nil {
		t.Fatalf("write CRD: %v", err)
	}
	t.Setenv("KUBEFU_SCHEMA_PATHS", path)

	p := Provider()
	cfg := schema.TestResourceDataRaw(t, p.Schema, map[string]interface{}{
		"schema_paths": []interface{}{path},
	})
	_, diags := p.ConfigureContextFunc(context.Background(), cfg)
	if len(diags) == 0 {
		t.Fatalf("expected diagnostics")
	}

	found := false
	for _, d := range diags {
		if d.Summary != "User schemas registered" {
			continue
		}
		if !strings.Contains(d.Detail, "kubefu_user_crd_custom_test_tst_queue_v1beta1") {
			t.Fatalf("diagnostic missing registered key: %q", d.Detail)
		}
		found = true
	}
	if !found {
		t.Fatalf("expected User schemas registered diagnostic, got: %+v", diags)
	}
}

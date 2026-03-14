package kubefu

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const sampleCRD = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: tstqueues.crd.custom.test
spec:
  group: crd.custom.test
  names:
    kind: TstQueue
    plural: tstqueues
  scope: Namespaced
  versions:
    - name: v1beta1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            apiVersion:
              type: string
            kind:
              type: string
            metadata:
              type: object
            spec:
              type: object
              properties:
                topicName:
                  type: string
                partitions:
                  type: integer
`

func TestLoadUserSchemasFromCRD(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "ktopic.yaml")
	if err := os.WriteFile(path, []byte(sampleCRD), 0o644); err != nil {
		t.Fatalf("write CRD: %v", err)
	}
	dataSources, diags := loadUserSchemas([]string{path})
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %+v", diags)
	}
	if len(dataSources) != 1 {
		t.Fatalf("expected 1 data source, got %d", len(dataSources))
	}
	var key string
	for k := range dataSources {
		key = k
	}
	if !strings.Contains(key, "tst_queue") {
		t.Fatalf("expected data source key to include kind, got %q", key)
	}
	ds := dataSources[key]
	if ds == nil || ds.Schema == nil {
		t.Fatalf("expected data source schema")
	}
	if field, ok := ds.Schema["api_version"]; !ok || !field.Computed {
		t.Fatalf("api_version should be computed")
	}
	if field, ok := ds.Schema["kind"]; !ok || !field.Computed {
		t.Fatalf("kind should be computed")
	}
	if _, ok := ds.Schema["kubefu_manifest_json"]; !ok {
		t.Fatalf("expected kubefu_manifest_json output")
	}
	if _, ok := ds.Schema["spec"]; !ok {
		t.Fatalf("expected spec field in schema")
	}
}

func TestLoadUserSchemasFromDirContinuesOnError(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "good.yaml")
	bad := filepath.Join(dir, "bad.yaml")
	if err := os.WriteFile(good, []byte(sampleCRD), 0o644); err != nil {
		t.Fatalf("write good CRD: %v", err)
	}
	if err := os.WriteFile(bad, []byte("not a crd"), 0o644); err != nil {
		t.Fatalf("write bad CRD: %v", err)
	}

	dataSources, diags := loadUserSchemas([]string{dir})
	if len(dataSources) != 1 {
		t.Fatalf("expected 1 data source, got %d", len(dataSources))
	}
	if len(diags) == 0 {
		t.Fatalf("expected diagnostics for bad file")
	}
}

package kubefu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestDecodeManifests_SingleDocumentYAML(t *testing.T) {
	manifest := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
  namespace: default
data:
  key: value
`
	objs, err := decodeManifests(manifest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(objs) != 1 {
		t.Fatalf("expected 1 object, got %d", len(objs))
	}
	if objs[0].GetAPIVersion() != "v1" {
		t.Errorf("expected apiVersion 'v1', got %q", objs[0].GetAPIVersion())
	}
	if objs[0].GetKind() != "ConfigMap" {
		t.Errorf("expected kind 'ConfigMap', got %q", objs[0].GetKind())
	}
	if objs[0].GetName() != "my-config" {
		t.Errorf("expected name 'my-config', got %q", objs[0].GetName())
	}
	if objs[0].GetNamespace() != "default" {
		t.Errorf("expected namespace 'default', got %q", objs[0].GetNamespace())
	}
}

func TestDecodeManifests_SingleDocumentJSON(t *testing.T) {
	manifest := `{"apiVersion":"v1","kind":"ServiceAccount","metadata":{"name":"my-sa","namespace":"kube-system"}}`
	objs, err := decodeManifests(manifest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(objs) != 1 {
		t.Fatalf("expected 1 object, got %d", len(objs))
	}
	if objs[0].GetKind() != "ServiceAccount" {
		t.Errorf("expected kind 'ServiceAccount', got %q", objs[0].GetKind())
	}
	if objs[0].GetName() != "my-sa" {
		t.Errorf("expected name 'my-sa', got %q", objs[0].GetName())
	}
}

func TestDecodeManifests_MultiDocumentYAML(t *testing.T) {
	manifest := `
apiVersion: v1
kind: Namespace
metadata:
  name: test-ns
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: test-sa
  namespace: test-ns
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-deploy
  namespace: test-ns
`
	objs, err := decodeManifests(manifest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(objs) != 3 {
		t.Fatalf("expected 3 objects, got %d", len(objs))
	}

	// Verify order
	expected := []struct {
		apiVersion string
		kind       string
		name       string
		namespace  string
	}{
		{"v1", "Namespace", "test-ns", ""},
		{"v1", "ServiceAccount", "test-sa", "test-ns"},
		{"apps/v1", "Deployment", "test-deploy", "test-ns"},
	}

	for i, exp := range expected {
		if objs[i].GetAPIVersion() != exp.apiVersion {
			t.Errorf("[%d] expected apiVersion %q, got %q", i, exp.apiVersion, objs[i].GetAPIVersion())
		}
		if objs[i].GetKind() != exp.kind {
			t.Errorf("[%d] expected kind %q, got %q", i, exp.kind, objs[i].GetKind())
		}
		if objs[i].GetName() != exp.name {
			t.Errorf("[%d] expected name %q, got %q", i, exp.name, objs[i].GetName())
		}
		if objs[i].GetNamespace() != exp.namespace {
			t.Errorf("[%d] expected namespace %q, got %q", i, exp.namespace, objs[i].GetNamespace())
		}
	}
}

func TestDecodeManifests_EmptyAndComments(t *testing.T) {
	manifest := `
# Leading comment
---
# Comment before first object

apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-1
  namespace: default
---
# Empty document between separators
---

---
# Another comment
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-2
  namespace: default
---
# Trailing comment
---
`
	objs, err := decodeManifests(manifest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(objs) != 2 {
		t.Fatalf("expected 2 objects, got %d", len(objs))
	}
	if objs[0].GetName() != "cm-1" {
		t.Errorf("expected first object 'cm-1', got %q", objs[0].GetName())
	}
	if objs[1].GetName() != "cm-2" {
		t.Errorf("expected second object 'cm-2', got %q", objs[1].GetName())
	}
}

func TestDecodeManifests_EmptyManifest(t *testing.T) {
	tests := []struct {
		name     string
		manifest string
	}{
		{"empty string", ""},
		{"whitespace only", "   \n\t\n  "},
		{"comments only", "# just a comment\n# another comment\n"},
		{"separator only", "---\n---\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			objs, err := decodeManifests(tc.manifest)
			if err == nil {
				t.Fatalf("expected error for %q, got %d objects", tc.name, len(objs))
			}
		})
	}
}

func TestManifestIDsAndParseManifestIDs(t *testing.T) {
	u1 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]any{
				"name": "my-ns",
			},
		},
	}
	u2 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":      "my-cm",
				"namespace": "my-ns",
			},
		},
	}
	u3 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]any{
				"name":      "my-deploy",
				"namespace": "my-ns",
			},
		},
	}
	u4 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]any{
				"name": "my-role",
			},
		},
	}

	// Single object ID (core group)
	id1 := manifestIDs([]*unstructured.Unstructured{u1})
	expectedID1 := "v1/Namespace/cluster/my-ns"
	if id1 != expectedID1 {
		t.Errorf("expected single ID %q, got %q", expectedID1, id1)
	}

	parsed1, err := parseManifestIDs(id1)
	if err != nil {
		t.Fatalf("parseManifestIDs failed: %v", err)
	}
	if len(parsed1) != 1 {
		t.Fatalf("expected 1 object, got %d", len(parsed1))
	}
	if parsed1[0].GetName() != "my-ns" || parsed1[0].GetKind() != "Namespace" || parsed1[0].GetAPIVersion() != "v1" {
		t.Errorf("parsed object mismatch: %+v", parsed1[0])
	}

	// Single object ID (grouped apiVersion with namespace)
	id3 := manifestIDs([]*unstructured.Unstructured{u3})
	expectedID3 := "apps/v1/Deployment/my-ns/my-deploy"
	if id3 != expectedID3 {
		t.Errorf("expected grouped ID %q, got %q", expectedID3, id3)
	}

	parsed3, err := parseManifestIDs(id3)
	if err != nil {
		t.Fatalf("parseManifestIDs failed: %v", err)
	}
	if len(parsed3) != 1 {
		t.Fatalf("expected 1 object, got %d", len(parsed3))
	}
	if parsed3[0].GetAPIVersion() != "apps/v1" || parsed3[0].GetKind() != "Deployment" || parsed3[0].GetNamespace() != "my-ns" || parsed3[0].GetName() != "my-deploy" {
		t.Errorf("parsed object mismatch: %+v", parsed3[0])
	}

	// Single object ID (grouped apiVersion cluster scoped)
	id4 := manifestIDs([]*unstructured.Unstructured{u4})
	expectedID4 := "rbac.authorization.k8s.io/v1/ClusterRole/cluster/my-role"
	if id4 != expectedID4 {
		t.Errorf("expected cluster role ID %q, got %q", expectedID4, id4)
	}

	parsed4, err := parseManifestIDs(id4)
	if err != nil {
		t.Fatalf("parseManifestIDs failed: %v", err)
	}
	if len(parsed4) != 1 {
		t.Fatalf("expected 1 object, got %d", len(parsed4))
	}
	if parsed4[0].GetAPIVersion() != "rbac.authorization.k8s.io/v1" || parsed4[0].GetKind() != "ClusterRole" || parsed4[0].GetNamespace() != "" || parsed4[0].GetName() != "my-role" {
		t.Errorf("parsed object mismatch: %+v", parsed4[0])
	}

	// Multiple objects ID (semicolon delimited with mixed core and grouped resources)
	compositeID := manifestIDs([]*unstructured.Unstructured{u1, u2, u3})
	expectedComposite := "v1/Namespace/cluster/my-ns;v1/ConfigMap/my-ns/my-cm;apps/v1/Deployment/my-ns/my-deploy"
	if compositeID != expectedComposite {
		t.Errorf("expected composite ID %q, got %q", expectedComposite, compositeID)
	}

	parsedComposite, err := parseManifestIDs(compositeID)
	if err != nil {
		t.Fatalf("parseManifestIDs failed: %v", err)
	}
	if len(parsedComposite) != 3 {
		t.Fatalf("expected 3 objects, got %d", len(parsedComposite))
	}
	if parsedComposite[0].GetName() != "my-ns" || parsedComposite[1].GetName() != "my-cm" || parsedComposite[2].GetName() != "my-deploy" {
		t.Errorf("parsed objects mismatch: %+v, %+v, %+v", parsedComposite[0], parsedComposite[1], parsedComposite[2])
	}
	if parsedComposite[2].GetAPIVersion() != "apps/v1" {
		t.Errorf("expected apiVersion 'apps/v1', got %q", parsedComposite[2].GetAPIVersion())
	}

	// Newline delimited parsing support
	newlineID := "v1/Namespace/cluster/my-ns\nv1/ConfigMap/my-ns/my-cm\napps/v1/Deployment/my-ns/my-deploy"
	parsedNewline, err := parseManifestIDs(newlineID)
	if err != nil {
		t.Fatalf("parseManifestIDs newline failed: %v", err)
	}
	if len(parsedNewline) != 3 {
		t.Fatalf("expected 3 objects, got %d", len(parsedNewline))
	}
	if parsedNewline[0].GetName() != "my-ns" || parsedNewline[1].GetName() != "my-cm" || parsedNewline[2].GetName() != "my-deploy" {
		t.Errorf("parsed objects mismatch: %+v, %+v, %+v", parsedNewline[0], parsedNewline[1], parsedNewline[2])
	}

	// Invalid IDs
	if _, err := parseManifestIDs(""); err == nil {
		t.Errorf("expected error for empty ID")
	}
	if _, err := parseManifestIDs("invalid-id"); err == nil {
		t.Errorf("expected error for invalid ID")
	}
	if _, err := parseManifestIDs("v1/ConfigMap/my-ns"); err == nil {
		t.Errorf("expected error for 3-part ID")
	}
	if _, err := parseManifestIDs("a/b/c/d/e/f"); err == nil {
		t.Errorf("expected error for 6-part ID")
	}

	// Empty and nil manifestIDs
	if id := manifestIDs(nil); id != "" {
		t.Errorf("expected empty string for nil slice, got %q", id)
	}
	if id := manifestIDs([]*unstructured.Unstructured{}); id != "" {
		t.Errorf("expected empty string for empty slice, got %q", id)
	}
	if id := manifestID(nil); id != "" {
		t.Errorf("expected empty string for nil object, got %q", id)
	}
}

func TestReverseObjects(t *testing.T) {
	if res := reverseObjects(nil); res != nil {
		t.Errorf("expected nil for nil slice, got %v", res)
	}
	if res := reverseObjects([]*unstructured.Unstructured{}); len(res) != 0 {
		t.Errorf("expected empty slice for empty slice, got %v", res)
	}

	u1 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "Namespace", "metadata": map[string]any{"name": "ns1"}}}
	u2 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ServiceAccount", "metadata": map[string]any{"name": "sa2"}}}
	u3 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "apps/v1", "kind": "Deployment", "metadata": map[string]any{"name": "dep3"}}}

	single := reverseObjects([]*unstructured.Unstructured{u1})
	if len(single) != 1 || single[0].GetName() != "ns1" {
		t.Errorf("unexpected single element result: %v", single)
	}

	multi := reverseObjects([]*unstructured.Unstructured{u1, u2, u3})
	if len(multi) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(multi))
	}
	if multi[0].GetName() != "dep3" || multi[1].GetName() != "sa2" || multi[2].GetName() != "ns1" {
		t.Errorf("expected [dep3, sa2, ns1], got [%s, %s, %s]", multi[0].GetName(), multi[1].GetName(), multi[2].GetName())
	}
}

func TestPrunedObjects(t *testing.T) {
	u1 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "Namespace", "metadata": map[string]any{"name": "ns1"}}}
	u2 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ServiceAccount", "metadata": map[string]any{"name": "sa2", "namespace": "ns1"}}}
	u3 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "apps/v1", "kind": "Deployment", "metadata": map[string]any{"name": "dep3", "namespace": "ns1"}}}

	// Removing u3 from [u1, u2, u3] -> should return [u3]
	pruned := prunedObjects([]*unstructured.Unstructured{u1, u2, u3}, []*unstructured.Unstructured{u1, u2})
	if len(pruned) != 1 || pruned[0].GetName() != "dep3" {
		t.Fatalf("expected [dep3], got %v", pruned)
	}

	// Removing u2 and u3 -> should return [u3, u2] (reverse order)
	pruned2 := prunedObjects([]*unstructured.Unstructured{u1, u2, u3}, []*unstructured.Unstructured{u1})
	if len(pruned2) != 2 || pruned2[0].GetName() != "dep3" || pruned2[1].GetName() != "sa2" {
		t.Fatalf("expected [dep3, sa2], got %v", pruned2)
	}

	// No objects removed
	prunedNone := prunedObjects([]*unstructured.Unstructured{u1, u2}, []*unstructured.Unstructured{u1, u2})
	if len(prunedNone) != 0 {
		t.Fatalf("expected 0 pruned objects, got %d", len(prunedNone))
	}
}

func TestInheritRecordedNamespaces(t *testing.T) {
	// u1 has no namespace specified in YAML
	u1 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm1"}}}
	// u2 already has explicit namespace in YAML
	u2 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm2", "namespace": "custom-ns"}}}
	// u3 is a new object not in recorded state
	u3 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm3"}}}

	recorded := []*unstructured.Unstructured{
		{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm1", "namespace": "recorded-ns"}}},
		{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm2", "namespace": "old-ns"}}},
	}

	inheritRecordedNamespaces([]*unstructured.Unstructured{u1, u2, u3}, recorded)

	if u1.GetNamespace() != "recorded-ns" {
		t.Errorf("expected u1 to inherit 'recorded-ns', got %q", u1.GetNamespace())
	}
	if u2.GetNamespace() != "custom-ns" {
		t.Errorf("expected u2 to keep 'custom-ns', got %q", u2.GetNamespace())
	}
	if u3.GetNamespace() != "" {
		t.Errorf("expected u3 to remain empty, got %q", u3.GetNamespace())
	}
}

func TestRetainUntouchedObjects(t *testing.T) {
	old1 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm1", "namespace": "ns"}}}
	old2 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm2", "namespace": "ns"}}}
	old3 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm3", "namespace": "ns"}}}

	// applied only replaced cm1, cm2 failed to apply, cm3 was not yet attempted
	applied1 := &unstructured.Unstructured{Object: map[string]any{"apiVersion": "v1", "kind": "ConfigMap", "metadata": map[string]any{"name": "cm1", "namespace": "ns"}}}

	retained := retainUntouchedObjects([]*unstructured.Unstructured{applied1}, []*unstructured.Unstructured{old1, old2, old3})
	if len(retained) != 3 {
		t.Fatalf("expected 3 retained objects, got %d", len(retained))
	}
	if retained[0].GetName() != "cm1" || retained[1].GetName() != "cm2" || retained[2].GetName() != "cm3" {
		t.Errorf("unexpected retained list: %v", retained)
	}
}

func TestValidateManifest(t *testing.T) {
	tests := []struct {
		name    string
		obj     *unstructured.Unstructured
		wantErr bool
	}{
		{
			name: "valid",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "v1",
					"kind":       "ConfigMap",
					"metadata": map[string]any{
						"name": "cm1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing apiVersion",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"kind": "ConfigMap",
					"metadata": map[string]any{
						"name": "cm1",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing kind",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "v1",
					"metadata": map[string]any{
						"name": "cm1",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "missing name",
			obj: &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "v1",
					"kind":       "ConfigMap",
					"metadata":   map[string]any{},
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateManifest(tc.obj)
			if (err != nil) != tc.wantErr {
				t.Fatalf("validateManifest() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestDecodeManifests_InvalidYAML(t *testing.T) {
	invalidYAML := `
apiVersion: v1
kind: ConfigMap
metadata: [unbalanced
`
	_, err := decodeManifests(invalidYAML)
	if err == nil {
		t.Fatalf("expected error for invalid YAML")
	}
}

func TestSetManifestFromObjects(t *testing.T) {
	resource := resourceManifest()
	schemaMap := resource.Schema

	t.Run("single object yaml", func(t *testing.T) {
		d := schemaMapToResourceData(t, schemaMap, map[string]any{
			"manifest": "original yaml",
		})
		u := &unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "v1",
				"kind":       "ConfigMap",
				"metadata": map[string]any{
					"name":      "cm1",
					"namespace": "default",
				},
				"data": map[string]any{
					"k": "v",
				},
			},
		}
		err := setManifestFromObjects(d, []*unstructured.Unstructured{u})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got := d.Get("manifest").(string)
		decoded, err := decodeManifests(got)
		if err != nil {
			t.Fatalf("failed to decode result: %v", err)
		}
		if len(decoded) != 1 || decoded[0].GetName() != "cm1" {
			t.Errorf("unexpected decoded manifest: %+v", decoded)
		}
	})

	t.Run("multi-document yaml", func(t *testing.T) {
		d := schemaMapToResourceData(t, schemaMap, map[string]any{
			"manifest": "original yaml",
		})
		u1 := &unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "v1",
				"kind":       "Namespace",
				"metadata": map[string]any{
					"name": "ns1",
				},
			},
		}
		u2 := &unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "v1",
				"kind":       "ConfigMap",
				"metadata": map[string]any{
					"name":      "cm2",
					"namespace": "ns1",
				},
			},
		}
		err := setManifestFromObjects(d, []*unstructured.Unstructured{u1, u2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		got := d.Get("manifest").(string)
		decoded, err := decodeManifests(got)
		if err != nil {
			t.Fatalf("failed to decode result: %v", err)
		}
		if len(decoded) != 2 {
			t.Fatalf("expected 2 decoded objects, got %d", len(decoded))
		}
		if decoded[0].GetName() != "ns1" || decoded[1].GetName() != "cm2" {
			t.Errorf("unexpected decoded manifest order: %+v", decoded)
		}
	})
}

func schemaMapToResourceData(t *testing.T, s map[string]*schema.Schema, raw map[string]any) *schema.ResourceData {
	t.Helper()
	res := &schema.Resource{Schema: s}
	d := res.TestResourceData()
	for k, v := range raw {
		if err := d.Set(k, v); err != nil {
			t.Fatalf("failed to set %s: %v", k, err)
		}
	}
	return d
}

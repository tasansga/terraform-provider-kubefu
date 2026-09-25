package kubefu

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	meta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	schemaApi "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/fake"
	clienttesting "k8s.io/client-go/testing"
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

func TestPrunedObjects_APIVersionUpgradeDoesNotPrune(t *testing.T) {
	oldObjs := []*unstructured.Unstructured{
		{
			Object: map[string]any{
				"apiVersion": "batch/v1beta1",
				"kind":       "CronJob",
				"metadata": map[string]any{
					"name":      "my-job",
					"namespace": "default",
				},
			},
		},
	}
	newObjs := []*unstructured.Unstructured{
		{
			Object: map[string]any{
				"apiVersion": "batch/v1",
				"kind":       "CronJob",
				"metadata": map[string]any{
					"name":      "my-job",
					"namespace": "default",
				},
			},
		},
	}

	pruned := prunedObjects(oldObjs, newObjs)
	if len(pruned) != 0 {
		t.Fatalf("expected 0 pruned objects on API version upgrade, got %d (%v)", len(pruned), pruned)
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

func TestObjectKeyOf(t *testing.T) {
	makeObj := func(apiVersion, kind, namespace, name string) *unstructured.Unstructured {
		obj := map[string]any{}
		if apiVersion != "" {
			obj["apiVersion"] = apiVersion
		}
		if kind != "" {
			obj["kind"] = kind
		}
		md := map[string]any{}
		if namespace != "" {
			md["namespace"] = namespace
		}
		if name != "" {
			md["name"] = name
		}
		if len(md) > 0 {
			obj["metadata"] = md
		}
		return &unstructured.Unstructured{Object: obj}
	}

	tests := []struct {
		name      string
		objA      *unstructured.Unstructured
		objB      *unstructured.Unstructured
		expectOkA bool
		expectOkB bool
		wantEqual bool
	}{
		{
			name:      "version-agnostic equality: batch/v1beta1 vs batch/v1 CronJob",
			objA:      makeObj("batch/v1beta1", "CronJob", "default", "my-job"),
			objB:      makeObj("batch/v1", "CronJob", "default", "my-job"),
			expectOkA: true,
			expectOkB: true,
			wantEqual: true,
		},
		{
			name:      "group mismatch: extensions/v1beta1 vs networking.k8s.io/v1 Ingress",
			objA:      makeObj("extensions/v1beta1", "Ingress", "default", "my-ingress"),
			objB:      makeObj("networking.k8s.io/v1", "Ingress", "default", "my-ingress"),
			expectOkA: true,
			expectOkB: true,
			wantEqual: false,
		},
		{
			name:      "kind mismatch",
			objA:      makeObj("v1", "ConfigMap", "default", "cm1"),
			objB:      makeObj("v1", "Secret", "default", "cm1"),
			expectOkA: true,
			expectOkB: true,
			wantEqual: false,
		},
		{
			name:      "namespace mismatch",
			objA:      makeObj("v1", "ConfigMap", "default", "cm1"),
			objB:      makeObj("v1", "ConfigMap", "other", "cm1"),
			expectOkA: true,
			expectOkB: true,
			wantEqual: false,
		},
		{
			name:      "name mismatch",
			objA:      makeObj("v1", "ConfigMap", "default", "cm1"),
			objB:      makeObj("v1", "ConfigMap", "default", "cm2"),
			expectOkA: true,
			expectOkB: true,
			wantEqual: false,
		},
		{
			name:      "nil safety: nil object",
			objA:      nil,
			objB:      makeObj("v1", "ConfigMap", "default", "cm1"),
			expectOkA: false,
			expectOkB: true,
			wantEqual: false,
		},
		{
			name:      "malformed: empty kind",
			objA:      makeObj("v1", "", "default", "cm1"),
			objB:      makeObj("v1", "ConfigMap", "default", "cm1"),
			expectOkA: false,
			expectOkB: true,
			wantEqual: false,
		},
		{
			name:      "malformed: empty name",
			objA:      makeObj("v1", "ConfigMap", "default", ""),
			objB:      makeObj("v1", "ConfigMap", "default", "cm1"),
			expectOkA: false,
			expectOkB: true,
			wantEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyA, okA := objectKeyOf(tt.objA)
			if okA != tt.expectOkA {
				t.Fatalf("objectKeyOf(objA) ok = %v, want %v", okA, tt.expectOkA)
			}
			keyB, okB := objectKeyOf(tt.objB)
			if okB != tt.expectOkB {
				t.Fatalf("objectKeyOf(objB) ok = %v, want %v", okB, tt.expectOkB)
			}
			isEqual := okA && okB && keyA == keyB
			if isEqual != tt.wantEqual {
				t.Errorf("objectKeyOf comparison = %v, want %v (keyA=%+v, keyB=%+v)", isEqual, tt.wantEqual, keyA, keyB)
			}
		})
	}
}

func TestInheritRecordedNamespaces_NilSafety(t *testing.T) {
	cm := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name": "cm1",
			},
		},
	}
	// Should not panic when nil objects are present
	inheritRecordedNamespaces([]*unstructured.Unstructured{nil, cm}, []*unstructured.Unstructured{nil, cm})
}

func TestObjectKeyOf_RejectsMalformedOrUnnamedObjects(t *testing.T) {
	if _, ok := objectKeyOf(nil); ok {
		t.Errorf("expected objectKeyOf(nil) to return ok=false")
	}

	missingKind := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"metadata": map[string]any{
				"name": "cm1",
			},
		},
	}
	if _, ok := objectKeyOf(missingKind); ok {
		t.Errorf("expected objectKeyOf with empty kind to return ok=false")
	}

	missingName := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata":   map[string]any{},
		},
	}
	if _, ok := objectKeyOf(missingName); ok {
		t.Errorf("expected objectKeyOf with empty name to return ok=false")
	}

	valid := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name": "cm1",
			},
		},
	}
	key, ok := objectKeyOf(valid)
	if !ok {
		t.Fatalf("expected objectKeyOf with valid kind and name to return ok=true")
	}
	if key.name != "cm1" || key.gk.Kind != "ConfigMap" {
		t.Errorf("unexpected key: %+v", key)
	}
}

func TestInheritRecordedNamespaces_IgnoresUnnamedObjects(t *testing.T) {
	unnamedObj := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata":   map[string]any{},
		},
	}
	recordedUnnamed := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"namespace": "recorded-ns",
			},
		},
	}

	inheritRecordedNamespaces([]*unstructured.Unstructured{unnamedObj}, []*unstructured.Unstructured{recordedUnnamed})

	if unnamedObj.GetNamespace() != "" {
		t.Errorf("expected unnamed object not to inherit namespace, got %q", unnamedObj.GetNamespace())
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

func setupDriftTestFixture(t *testing.T, liveObjs ...runtime.Object) (*schema.ResourceData, *providerConfig, *unstructured.Unstructured, *unstructured.Unstructured) {
	t.Helper()

	u1 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":      "cm1",
				"namespace": "default",
			},
			"data": map[string]any{
				"k1": "v1",
			},
		},
	}
	u2 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":      "cm2",
				"namespace": "default",
			},
			"data": map[string]any{
				"k2": "v2",
			},
		},
	}

	manifest := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
data:
  k1: v1
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm2
  namespace: default
data:
  k2: v2
`

	scheme := runtime.NewScheme()
	effectiveLive := liveObjs
	if len(effectiveLive) == 0 {
		effectiveLive = []runtime.Object{u1}
	}
	fakeClient := fake.NewSimpleDynamicClient(scheme, effectiveLive...)
	fakeClient.PrependReactor("patch", "configmaps", func(action clienttesting.Action) (handled bool, ret runtime.Object, err error) {
		patchAction := action.(clienttesting.PatchAction)
		var obj unstructured.Unstructured
		if err := json.Unmarshal(patchAction.GetPatch(), &obj.Object); err != nil {
			return true, nil, err
		}
		return true, &obj, nil
	})

	mapper := meta.NewDefaultRESTMapper([]schemaApi.GroupVersion{
		{Group: "", Version: "v1"},
	})
	mapper.Add(schemaApi.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}, meta.RESTScopeNamespace)

	cfg := &providerConfig{
		dynamicClient: fakeClient,
		restMapper:    mapper,
	}

	resource := resourceManifest()
	initialID := manifestIDs([]*unstructured.Unstructured{u1, u2})
	d := schemaMapToResourceData(t, resource.Schema, map[string]any{
		"manifest": manifest,
	})
	d.SetId(initialID)

	return d, cfg, u1, u2
}

func setupDriftTestCase(t *testing.T, manifest string, id string, liveObjs ...*unstructured.Unstructured) (*schema.ResourceData, *providerConfig) {
	t.Helper()
	scheme := runtime.NewScheme()
	fakeClient := fake.NewSimpleDynamicClient(scheme)
	mapper := meta.NewDefaultRESTMapper([]schemaApi.GroupVersion{
		{Group: "", Version: "v1"},
		{Group: "apps", Version: "v1"},
	})
	mapper.Add(schemaApi.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}, meta.RESTScopeNamespace)
	mapper.Add(schemaApi.GroupVersionKind{Group: "", Version: "v1", Kind: "Service"}, meta.RESTScopeNamespace)
	mapper.Add(schemaApi.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}, meta.RESTScopeNamespace)

	for _, u := range liveObjs {
		if u == nil {
			continue
		}
		mapping, err := mapper.RESTMapping(u.GroupVersionKind().GroupKind(), u.GroupVersionKind().Version)
		if err != nil {
			t.Fatalf("setupDriftTestCase mapper error: %v", err)
		}
		ns := u.GetNamespace()
		if ns == "" && mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			ns = "default"
		}
		var target dynamic.ResourceInterface
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			target = fakeClient.Resource(mapping.Resource).Namespace(ns)
		} else {
			target = fakeClient.Resource(mapping.Resource)
		}
		if _, err := target.Create(context.Background(), u, metav1.CreateOptions{}); err != nil {
			t.Fatalf("setupDriftTestCase seed error: %v", err)
		}
	}

	cfg := &providerConfig{
		dynamicClient: fakeClient,
		restMapper:    mapper,
	}

	resource := resourceManifest()
	d := schemaMapToResourceData(t, resource.Schema, map[string]any{
		"manifest": manifest,
	})
	if id == "" && len(liveObjs) > 0 {
		id = manifestIDs(liveObjs)
	}
	d.SetId(id)

	return d, cfg
}

func TestResourceManifestRead_PartialNotFoundDrift(t *testing.T) {
	d, cfg, u1, _ := setupDriftTestFixture(t)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	// 1. Live ID must be updated in d.Id() to reflect true cluster state (cm1 only)
	expectedLiveID := manifestIDs([]*unstructured.Unstructured{u1})
	if d.Id() != expectedLiveID {
		t.Errorf("expected d.Id() to be %q, got %q", expectedLiveID, d.Id())
	}

	// 2. d.Get("manifest") must be updated with live objects (only cm1)
	newManifest := d.Get("manifest").(string)
	objs, err := decodeManifests(newManifest)
	if err != nil {
		t.Fatalf("failed to decode updated manifest: %v", err)
	}
	if len(objs) != 1 {
		t.Fatalf("expected 1 live object in manifest, got %d", len(objs))
	}
	if objs[0].GetName() != "cm1" {
		t.Errorf("expected live object to be 'cm1', got %q", objs[0].GetName())
	}
}

func TestResourceManifestRead_DriftLifecycle_AcceptedOutOfBandDeletion(t *testing.T) {
	d, cfg, u1, _ := setupDriftTestFixture(t)

	// First read: drift detected (cm2 missing from cluster)
	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("first resourceManifestRead failed: %v", diags)
	}

	expectedLiveID := manifestIDs([]*unstructured.Unstructured{u1})
	if d.Id() != expectedLiveID {
		t.Fatalf("expected 1st read d.Id() to update to live ID %q, got %q", expectedLiveID, d.Id())
	}

	// Subsequent plan: user accepted out-of-band deletion and updated HCL config to only cm1.
	// State now has manifest=cm1 and id=cm1.
	acceptedManifest := d.Get("manifest").(string)
	resource := resourceManifest()
	d2 := schemaMapToResourceData(t, resource.Schema, map[string]any{
		"manifest": acceptedManifest,
	})
	d2.SetId(expectedLiveID)

	diags2 := resourceManifestRead(context.Background(), d2, cfg)
	if diags2.HasError() {
		t.Fatalf("second resourceManifestRead failed: %v", diags2)
	}

	// Clean state: id matches live cluster (cm1) without any trapped cm2
	if d2.Id() != expectedLiveID {
		t.Fatalf("expected 2nd read d.Id() to remain %q, got %q", expectedLiveID, d2.Id())
	}
}

func TestResourceManifestRead_DriftLifecycle_RecreateMissingResource(t *testing.T) {
	d, cfg, u1, u2 := setupDriftTestFixture(t)

	resource := resourceManifest()
	initialID := manifestIDs([]*unstructured.Unstructured{u1, u2})
	manifest := d.Get("manifest").(string)

	// Drift detected on read: cm2 missing
	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	expectedLiveID := manifestIDs([]*unstructured.Unstructured{u1})
	if d.Id() != expectedLiveID {
		t.Fatalf("expected d.Id() to update to live ID %q, got %q", expectedLiveID, d.Id())
	}

	// State manifest now has cm1, but HCL still has cm1 and cm2.
	// OpenTofu detects state != HCL and plans an Update.
	// When update runs with original HCL manifest:
	dUpdate := schemaMapToResourceData(t, resource.Schema, map[string]any{
		"manifest": manifest,
	})
	dUpdate.SetId(expectedLiveID)

	diagsUpdate := resourceManifestCreateOrUpdate(context.Background(), dUpdate, cfg)
	if diagsUpdate.HasError() {
		t.Fatalf("resourceManifestCreateUpdate failed: %v", diagsUpdate)
	}

	if dUpdate.Id() != initialID {
		t.Fatalf("expected dUpdate.Id() to restore %q, got %q", initialID, dUpdate.Id())
	}
}

func TestBuildKubeClients_AsymmetricInjectionErrors(t *testing.T) {
	fakeClient := fake.NewSimpleDynamicClient(runtime.NewScheme())
	fakeMapper := meta.NewDefaultRESTMapper([]schemaApi.GroupVersion{
		{Group: "", Version: "v1"},
	})

	expectedErr := "both dynamicClient and restMapper must be provided together for client injection"

	t.Run("dynamicClient only returns error", func(t *testing.T) {
		cfg := &providerConfig{dynamicClient: fakeClient, restMapper: nil}
		client, mapper, err := buildKubeClients(cfg)
		if err == nil || !strings.Contains(err.Error(), expectedErr) {
			t.Fatalf("expected error containing %q, got err=%v", expectedErr, err)
		}
		if client != nil || mapper != nil {
			t.Errorf("expected nil clients on error, got client=%v, mapper=%v", client, mapper)
		}
	})

	t.Run("restMapper only returns error", func(t *testing.T) {
		cfg := &providerConfig{dynamicClient: nil, restMapper: fakeMapper}
		client, mapper, err := buildKubeClients(cfg)
		if err == nil || !strings.Contains(err.Error(), expectedErr) {
			t.Fatalf("expected error containing %q, got err=%v", expectedErr, err)
		}
		if client != nil || mapper != nil {
			t.Errorf("expected nil clients on error, got client=%v, mapper=%v", client, mapper)
		}
	})

	t.Run("both non-nil returns clients without error", func(t *testing.T) {
		cfg := &providerConfig{dynamicClient: fakeClient, restMapper: fakeMapper}
		client, mapper, err := buildKubeClients(cfg)
		if err != nil {
			t.Fatalf("unexpected error when both clients provided: %v", err)
		}
		if client != fakeClient || mapper != fakeMapper {
			t.Errorf("expected returned clients to match injected clients")
		}
	})

	t.Run("nil cfg returns error", func(t *testing.T) {
		client, mapper, err := buildKubeClients(nil)
		if err == nil || !strings.Contains(err.Error(), "provider configuration is missing") {
			t.Fatalf("expected missing configuration error, got err=%v", err)
		}
		if client != nil || mapper != nil {
			t.Errorf("expected nil clients on error, got client=%v, mapper=%v", client, mapper)
		}
	})

	t.Run("retries client construction after transient failure without caching error", func(t *testing.T) {
		cfg := &providerConfig{
			KubeConfigPath: "/nonexistent/invalid/kubeconfig/path",
		}
		client, mapper, err := buildKubeClients(cfg)
		if err == nil {
			t.Fatalf("expected error on first call with invalid kubeconfig path")
		}
		if client != nil || mapper != nil {
			t.Errorf("expected nil clients on error, got client=%v, mapper=%v", client, mapper)
		}

		// Create a valid temporary kubeconfig
		tmpDir := t.TempDir()
		kubeconfigPath := tmpDir + "/kubeconfig"
		kubeconfigContent := `apiVersion: v1
clusters:
- cluster:
    server: https://127.0.0.1:6443
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
current-context: test-context
kind: Config
preferences: {}
users:
- name: test-user
  user:
    token: fake-token
`
		if err := os.WriteFile(kubeconfigPath, []byte(kubeconfigContent), 0600); err != nil {
			t.Fatalf("failed to write temporary kubeconfig: %v", err)
		}

		// Update path to valid kubeconfig - retry should now succeed
		cfg.KubeConfigPath = kubeconfigPath
		client2, mapper2, err2 := buildKubeClients(cfg)
		if err2 != nil {
			t.Fatalf("expected retry to succeed after transient error, got err=%v", err2)
		}
		if client2 == nil || mapper2 == nil {
			t.Fatalf("expected non-nil client and mapper on retry")
		}

		// Subsequent call should return cached instances
		client3, mapper3, err3 := buildKubeClients(cfg)
		if err3 != nil {
			t.Fatalf("unexpected error on subsequent call: %v", err3)
		}
		if client3 != client2 || mapper3 != mapper2 {
			t.Errorf("expected cached clients to be returned")
		}
	})
}

func TestResolveNamespace_InjectedClientsDefaultsHermetically(t *testing.T) {
	fakeClient := fake.NewSimpleDynamicClient(runtime.NewScheme())
	cfg := &providerConfig{dynamicClient: fakeClient}

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name": "no-namespace-cm",
			},
		},
	}

	namespacedMapping := &meta.RESTMapping{
		Scope: meta.RESTScopeNamespace,
	}

	ns, err := resolveNamespace(cfg, u, namespacedMapping)
	if err != nil {
		t.Fatalf("unexpected error from resolveNamespace: %v", err)
	}
	if ns != "default" {
		t.Errorf("expected namespace 'default' when mock client is injected, got %q", ns)
	}
}

func TestResourceManifestCreateOrUpdate_PartialApplyFailurePreservesResolvedNamespace(t *testing.T) {
	manifest := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm2
`

	u2 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":      "cm2",
				"namespace": "default",
			},
		},
	}
	// cm1 is newly added without explicit namespace and is not yet in oldObjs,
	// so inheritRecordedNamespaces does not pre-populate its namespace.
	oldObjs := []*unstructured.Unstructured{u2}

	scheme := runtime.NewScheme()
	fakeClient := fake.NewSimpleDynamicClient(scheme)
	fakeClient.PrependReactor("patch", "configmaps", func(action clienttesting.Action) (handled bool, ret runtime.Object, err error) {
		patchAction := action.(clienttesting.PatchAction)
		if patchAction.GetName() == "cm2" {
			return true, nil, fmt.Errorf("simulated patch failure for cm2")
		}
		var obj unstructured.Unstructured
		if err := json.Unmarshal(patchAction.GetPatch(), &obj.Object); err != nil {
			return true, nil, err
		}
		return true, &obj, nil
	})

	mapper := meta.NewDefaultRESTMapper([]schemaApi.GroupVersion{
		{Group: "", Version: "v1"},
	})
	mapper.Add(schemaApi.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}, meta.RESTScopeNamespace)

	cfg := &providerConfig{
		dynamicClient: fakeClient,
		restMapper:    mapper,
	}

	resource := resourceManifest()
	d := schemaMapToResourceData(t, resource.Schema, map[string]any{
		"manifest": manifest,
	})
	d.SetId(manifestIDs(oldObjs))

	diags := resourceManifestCreateOrUpdate(context.Background(), d, cfg)
	if !diags.HasError() {
		t.Fatalf("expected create/update to fail, got success")
	}

	ids, err := parseManifestIDs(d.Id())
	if err != nil {
		t.Fatalf("unexpected error parsing d.Id(): %v", err)
	}
	if len(ids) != 2 {
		t.Fatalf("expected 2 objects in d.Id(), got %d: %q", len(ids), d.Id())
	}
	for _, idObj := range ids {
		if idObj.GetNamespace() == "" || idObj.GetNamespace() == "cluster" {
			t.Errorf("expected non-empty namespace, got %q for %s", idObj.GetNamespace(), idObj.GetName())
		}
	}
}

func TestManifestDiffSuppress(t *testing.T) {
	tests := []struct {
		name     string
		oldVal   string
		newVal   string
		id       string
		suppress bool
	}{
		{
			name:     "identical strings",
			oldVal:   "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: cm\n",
			newVal:   "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: cm\n",
			suppress: true,
		},
		{
			name:     "whitespace differences",
			oldVal:   "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: cm\n",
			newVal:   "\n\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: cm\n\n",
			suppress: true,
		},
		{
			name:     "yaml comments",
			oldVal:   "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: cm\n",
			newVal:   "# Comment here\napiVersion: v1\n# Another comment\nkind: ConfigMap\nmetadata:\n  name: cm\n",
			suppress: true,
		},
		{
			name: "key ordering differences",
			oldVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
  namespace: default
data:
  a: "1"
  b: "2"
`,
			newVal: `
kind: ConfigMap
apiVersion: v1
metadata:
  namespace: default
  name: test
data:
  b: "2"
  a: "1"
`,
			suppress: true,
		},
		{
			name: "server generated metadata (status, uid, resourceVersion, creationTimestamp)",
			oldVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
  namespace: default
  uid: "12345-67890"
  resourceVersion: "9999"
  creationTimestamp: "2026-09-25T00:00:00Z"
status:
  phase: Active
data:
  key: value
`,
			newVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
  namespace: default
data:
  key: value
`,
			suppress: true,
		},
		{
			name: "actual data change must not be suppressed",
			oldVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
data:
  key: old-value
`,
			newVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
data:
  key: new-value
`,
			suppress: false,
		},
		{
			name: "field deletion in newVal must not be suppressed",
			oldVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
  namespace: default
data:
  a: "1"
  b: "2"
`,
			newVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test
  namespace: default
data:
  a: "1"
`,
			suppress: false,
		},
		{
			name: "different number of objects must not be suppressed",
			oldVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test1
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test2
`,
			newVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test1
`,
			suppress: false,
		},
		{
			name: "multi-document reordering is suppressed",
			oldVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test1
  namespace: default
data:
  key: val1
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test2
  namespace: default
data:
  key: val2
`,
			newVal: `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test2
  namespace: default
data:
  key: val2
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test1
  namespace: default
data:
  key: val1
`,
			suppress: true,
		},
		{
			name: "duplicate object replacing another object must not be suppressed",
			oldVal: `apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm2
  namespace: default
`,
			newVal: `apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
`,
			suppress: false,
		},
		{
			name: "namespace omitted in newVal is suppressed when matching recorded id namespace",
			oldVal: `apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
data:
  key: value
`,
			newVal: `apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
data:
  key: value
`,
			id:       "v1/ConfigMap/default/cm1",
			suppress: true,
		},
		{
			name: "field deletion in newVal (e.g. clusterIP) must not be suppressed",
			oldVal: `apiVersion: v1
kind: Service
metadata:
  name: my-svc
  namespace: default
spec:
  clusterIP: 10.96.0.1
  ports:
  - port: 80
`,
			newVal: `apiVersion: v1
kind: Service
metadata:
  name: my-svc
  namespace: default
spec:
  ports:
  - port: 80
`,
			suppress: false,
		},
		{
			name: "numeric type representation difference (JSON float64 vs YAML int64) is suppressed",
			oldVal: `{"apiVersion": "apps/v1", "kind": "Deployment", "metadata": {"name": "d", "namespace": "default"}, "spec": {"replicas": 1.0}}`,
			newVal: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: d
  namespace: default
spec:
  replicas: 1
`,
			suppress: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var d *schema.ResourceData
			if tc.id != "" {
				res := resourceManifest()
				d = schemaMapToResourceData(t, res.Schema, map[string]any{
					"manifest": tc.oldVal,
				})
				d.SetId(tc.id)
			}
			got := manifestDiffSuppress("manifest", tc.oldVal, tc.newVal, d)
			if got != tc.suppress {
				t.Errorf("manifestDiffSuppress() = %v, want %v", got, tc.suppress)
			}
		})
	}
}

func TestResourceManifestRead_PreservesOriginalManifestWhenAllFound(t *testing.T) {
	originalManifest := `# Custom comment preserved
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: default
data:
  key: value
`
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":              "test-cm",
				"namespace":         "default",
				"uid":               "abc-123",
				"resourceVersion":   "100",
				"creationTimestamp": "2026-09-25T00:00:00Z",
			},
			"data": map[string]any{
				"key": "value",
			},
		},
	}
	d, cfg := setupDriftTestCase(t, originalManifest, "", u)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	// Verify that the manifest string in state was NOT overwritten by the live cluster object
	gotManifest := d.Get("manifest").(string)
	if gotManifest != originalManifest {
		t.Errorf("expected original manifest to be preserved byte-for-byte, got:\n%s", gotManifest)
	}
}

func TestResourceManifestRead_UpdatesManifestWhenPartialNotFound(t *testing.T) {
	manifest := `apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm1
  namespace: default
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm2
  namespace: default
`
	// Seed only test-cm1 in the cluster, test-cm2 is missing (simulating out-of-band deletion)
	u1 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":      "test-cm1",
				"namespace": "default",
			},
		},
	}
	d, cfg := setupDriftTestCase(t, manifest, "v1/ConfigMap/default/test-cm1;v1/ConfigMap/default/test-cm2", u1)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	// Because test-cm2 was missing, setManifestFromObjects MUST have been called to reflect only test-cm1,
	// triggering Terraform drift so cm2 can be recreated.
	gotManifest := d.Get("manifest").(string)
	if strings.Contains(gotManifest, "test-cm2") {
		t.Errorf("expected missing cm2 to be removed from state manifest to trigger drift recreation, got:\n%s", gotManifest)
	}
	if !strings.Contains(gotManifest, "test-cm1") {
		t.Errorf("expected surviving cm1 to remain in state manifest, got:\n%s", gotManifest)
	}
}

func TestResourceManifestRead_DetectsSemanticDrift(t *testing.T) {
	desiredManifest := `apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: default
data:
  key: desired-value
`
	// Seed the object in the cluster with drifted data
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":              "test-cm",
				"namespace":         "default",
				"uid":               "abc-123",
				"resourceVersion":   "101",
				"creationTimestamp": "2026-09-25T00:00:00Z",
			},
			"data": map[string]any{
				"key": "cluster-drifted-value",
			},
		},
	}
	d, cfg := setupDriftTestCase(t, desiredManifest, "", u)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	gotManifest := d.Get("manifest").(string)
	if !strings.Contains(gotManifest, "cluster-drifted-value") {
		t.Errorf("expected state manifest to be updated with cluster-drifted-value, got:\n%s", gotManifest)
	}
}

func TestResourceManifestRead_NoSemanticDriftPreservesConfigString(t *testing.T) {
	originalManifest := `# custom comment
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: default
data:
  key: same-value
`
	// Seed the object in the cluster with server metadata and status
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":              "test-cm",
				"namespace":         "default",
				"uid":               "abc-123",
				"resourceVersion":   "100",
				"creationTimestamp": "2026-09-25T00:00:00Z",
			},
			"data": map[string]any{
				"key": "same-value",
			},
			"status": map[string]any{
				"phase": "Active",
			},
		},
	}
	d, cfg := setupDriftTestCase(t, originalManifest, "", u)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	gotManifest := d.Get("manifest").(string)
	if gotManifest != originalManifest {
		t.Errorf("expected original manifest to be preserved byte-for-byte, got:\n%s", gotManifest)
	}
}

func TestResourceManifestRead_NoSemanticDriftWhenNamespaceOmittedInConfig(t *testing.T) {
	originalManifest := `apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
data:
  key: value
`
	// Seed the object in the cluster with namespace: "default" and data.key: "value"
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":              "cm1",
				"namespace":         "default",
				"uid":               "abc-123",
				"resourceVersion":   "100",
				"creationTimestamp": "2026-09-25T00:00:00Z",
			},
			"data": map[string]any{
				"key": "value",
			},
		},
	}
	d, cfg := setupDriftTestCase(t, originalManifest, "v1/ConfigMap/default/cm1", u)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	gotManifest := d.Get("manifest").(string)
	if gotManifest != originalManifest {
		t.Errorf("expected original manifest to be preserved byte-for-byte, got:\n%s", gotManifest)
	}
}

func TestResourceManifestRead_ServerDefaultsIgnoredInDriftCheck(t *testing.T) {
	stateManifest := `apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: default
spec:
  ports:
  - port: 80
`
	// Seed the object in the cluster with ports and server defaults
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]any{
				"name":              "my-service",
				"namespace":         "default",
				"uid":               "abc-123",
				"resourceVersion":   "100",
				"creationTimestamp": "2026-09-25T00:00:00Z",
			},
			"spec": map[string]any{
				"clusterIP":       "10.96.0.1",
				"clusterIPs":      []any{"10.96.0.1"},
				"sessionAffinity": "None",
				"type":            "ClusterIP",
				"ports": []any{
					map[string]any{
						"port": int64(80),
					},
				},
			},
		},
	}
	d, cfg := setupDriftTestCase(t, stateManifest, "", u)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	gotManifest := d.Get("manifest").(string)
	if gotManifest != stateManifest {
		t.Errorf("expected original manifest to be preserved byte-for-byte, got:\n%s", gotManifest)
	}
}

func TestResourceManifestRead_DuplicateKeyInStateTriggersDrift(t *testing.T) {
	// Duplicate ConfigMap key (same name/namespace/GVK) in state
	stateManifest := `# comment to verify drift rewrite
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
data:
  key: value
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm1
  namespace: default
data:
  key: value
`
	u1 := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name":      "cm1",
				"namespace": "default",
			},
			"data": map[string]any{
				"key": "value",
			},
		},
	}
	d, cfg := setupDriftTestCase(t, stateManifest, "v1/ConfigMap/default/cm1;v1/ConfigMap/default/cm1", u1)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	// Because duplicate key was present in state, drift must be detected and manifest rewritten (comment removed)
	gotManifest := d.Get("manifest").(string)
	if gotManifest == stateManifest {
		t.Errorf("expected drift to be detected and state manifest overwritten, but it was unchanged")
	}
	if strings.Contains(gotManifest, "# comment to verify drift rewrite") {
		t.Errorf("expected custom comment to be stripped upon drift rewrite, got:\n%s", gotManifest)
	}
}

func TestIsSubsetOrEqual_NumericComparison(t *testing.T) {
	if !isSubsetOrEqual(1, 1) {
		t.Errorf("expected isSubsetOrEqual(1, 1) to be true")
	}
	if isSubsetOrEqual(1, 2) {
		t.Errorf("expected isSubsetOrEqual(1, 2) to be false")
	}
	if !isSubsetOrEqual(int64(10), float64(10.0)) {
		t.Errorf("expected isSubsetOrEqual(int64(10), float64(10.0)) to be true")
	}
	if isSubsetOrEqual(int64(10), float64(20.0)) {
		t.Errorf("expected isSubsetOrEqual(int64(10), float64(20.0)) to be false")
	}
	if isSubsetOrEqual(float32(10.0), float64(10.0)) {
		t.Errorf("expected isSubsetOrEqual with float32 to be false")
	}
}

func TestIsSemanticEqual(t *testing.T) {
	if !isSemanticEqual(map[string]any{"replicas": float64(1)}, map[string]any{"replicas": int64(1)}) {
		t.Errorf("expected isSemanticEqual to be true for float64(1) and int64(1)")
	}
	if isSemanticEqual(map[string]any{"replicas": float64(1)}, map[string]any{"replicas": int64(2)}) {
		t.Errorf("expected isSemanticEqual to be false for float64(1) and int64(2)")
	}
	if isSemanticEqual(map[string]any{"replicas": float64(1), "extra": "field"}, map[string]any{"replicas": int64(1)}) {
		t.Errorf("expected isSemanticEqual to be false when subset is not superset")
	}
}

func TestResourceManifestRead_DetectsNumericDrift(t *testing.T) {
	stateManifest := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: dep1
  namespace: default
spec:
  replicas: 1
`
	uLive := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]any{
				"name":      "dep1",
				"namespace": "default",
			},
			"spec": map[string]any{
				"replicas": int64(5),
			},
		},
	}
	d, cfg := setupDriftTestCase(t, stateManifest, "apps/v1/Deployment/default/dep1", uLive)

	diags := resourceManifestRead(context.Background(), d, cfg)
	if diags.HasError() {
		t.Fatalf("resourceManifestRead failed: %v", diags)
	}

	gotManifest := d.Get("manifest").(string)
	objs, err := decodeManifests(gotManifest)
	if err != nil {
		t.Fatalf("failed to decode updated manifest: %v", err)
	}
	if len(objs) != 1 {
		t.Fatalf("expected 1 object, got %d", len(objs))
	}
	replicasRaw, found, err := unstructured.NestedFieldNoCopy(objs[0].Object, "spec", "replicas")
	if err != nil || !found {
		t.Fatalf("failed to get replicas from updated manifest: found=%v, err=%v", found, err)
	}
	replicas, ok := toFloat64(replicasRaw)
	if !ok || replicas != 5 {
		t.Errorf("expected replicas to be 5, got %v", replicasRaw)
	}
}

func TestIsSubsetOrEqual_RecursionDepthLimit(t *testing.T) {
	buildDeepMap := func(depth int) map[string]any {
		root := make(map[string]any)
		curr := root
		for i := 0; i < depth; i++ {
			next := make(map[string]any)
			curr["child"] = next
			curr = next
		}
		return root
	}

	m1 := buildDeepMap(50)
	m2 := buildDeepMap(50)
	if !isSubsetOrEqual(m1, m2) {
		t.Errorf("expected depth 50 to return true")
	}

	deep1 := buildDeepMap(105)
	deep2 := buildDeepMap(105)
	if isSubsetOrEqual(deep1, deep2) {
		t.Errorf("expected depth 105 to exceed maxSubsetDepth and return false")
	}
}

func TestIsSubsetOrEqual_LargeIntegerPrecision(t *testing.T) {
	n1 := int64(1<<60 + 1)
	n2 := int64(1<<60 + 2)
	if isSubsetOrEqual(n1, n2) {
		t.Errorf("expected isSubsetOrEqual(%d, %d) to be false due to 64-bit precision, got true", n1, n2)
	}
	if !isSubsetOrEqual(n1, n1) {
		t.Errorf("expected isSubsetOrEqual(%d, %d) to be true, got false", n1, n1)
	}

	u1 := uint64(1<<63 + 1)
	u2 := uint64(1<<63 + 2)
	if isSubsetOrEqual(u1, u2) {
		t.Errorf("expected isSubsetOrEqual(%d, %d) to be false due to 64-bit precision, got true", u1, u2)
	}
	if !isSubsetOrEqual(u1, u1) {
		t.Errorf("expected isSubsetOrEqual(%d, %d) to be true, got false", u1, u1)
	}
}

func TestBuildKubeClients_Concurrent(t *testing.T) {
	tmpDir := t.TempDir()
	kubeconfigPath := tmpDir + "/kubeconfig"
	kubeconfigContent := `apiVersion: v1
clusters:
- cluster:
    server: https://127.0.0.1:6443
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
current-context: test-context
kind: Config
preferences: {}
users:
- name: test-user
  user:
    token: fake-token
`
	if err := os.WriteFile(kubeconfigPath, []byte(kubeconfigContent), 0600); err != nil {
		t.Fatalf("failed to write temporary kubeconfig: %v", err)
	}

	cfg := &providerConfig{
		KubeConfigPath: kubeconfigPath,
	}

	type clientResult struct {
		client dynamic.Interface
		mapper meta.RESTMapper
	}
	results := make([]clientResult, 20)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			client, mapper, err := buildKubeClients(cfg)
			if err != nil {
				t.Errorf("unexpected error from buildKubeClients: %v", err)
				return
			}
			if client == nil || mapper == nil {
				t.Errorf("expected non-nil client and mapper")
				return
			}
			results[idx] = clientResult{client: client, mapper: mapper}
		}(i)
	}
	wg.Wait()

	if results[0].client == nil || results[0].mapper == nil {
		t.Fatalf("goroutine 0 failed to obtain clients")
	}
	firstClient := results[0].client
	firstMapper := results[0].mapper
	for i := 1; i < 20; i++ {
		if results[i].client != firstClient {
			t.Errorf("goroutine %d got client %p, want %p", i, results[i].client, firstClient)
		}
		if results[i].mapper != firstMapper {
			t.Errorf("goroutine %d got mapper %p, want %p", i, results[i].mapper, firstMapper)
		}
	}
}

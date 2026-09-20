# kubefu_manifest

Apply a Kubernetes manifest (YAML or JSON) to the target cluster using server-side apply.

Notes:
- Multi-document YAML manifests (separated by `---`) are fully supported. Documents are applied in stream order on create/update, and deleted in **reverse order** on destroy (so workloads/dependents are deleted before CRDs/dependencies).
- On read, the provider normalizes the manifest (drops `status`, sorts keys) which may reformat the input. Multi-document manifests are re-serialized with `---` separators.
- If `metadata.namespace` is omitted, the namespace is resolved from the kubeconfig context, falling back to `default`. Cluster-scoped resources ignore namespace.
- Destroy deletes the live objects in the cluster (ignoring resources that are already NotFound).

## Example Usage

### Single Document

```terraform
resource "kubefu_manifest" "example" {
  manifest = <<-YAML
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: example
      namespace: default
    data:
      hello: world
  YAML
}
```

### Multi-Document YAML

```terraform
resource "kubefu_manifest" "bundle" {
  manifest = <<-YAML
    apiVersion: v1
    kind: Namespace
    metadata:
      name: app-demo
    ---
    apiVersion: v1
    kind: ServiceAccount
    metadata:
      name: demo-sa
      namespace: app-demo
    ---
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: demo-config
      namespace: app-demo
    data:
      env: production
  YAML
}
```

## Argument Reference

- `manifest` (Required) Kubernetes manifest to apply. Accepts YAML or JSON (including multi-document YAML).
- `field_manager` (Optional) Field manager name used for server-side apply. Defaults to `kubefu`.
- `force` (Optional) Force apply in case of conflicts. Defaults to `false`.

## Import

Use the following import ID format for a single resource:

`apiVersion/kind/namespace/name`

For cluster-scoped resources, use `cluster` as the namespace value.

For multi-document manifests, join each resource's ID with a semicolon `;` (or newline):

`apiVersion/kind/namespace/name;apiVersion/kind/namespace/name`

Example:

```bash
# Single resource
terraform import kubefu_manifest.example v1/ConfigMap/default/example

# Multi-document bundle
terraform import kubefu_manifest.bundle "v1/Namespace/cluster/app-demo;v1/ServiceAccount/app-demo/demo-sa;v1/ConfigMap/app-demo/demo-config"
```

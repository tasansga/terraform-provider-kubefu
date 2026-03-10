# kubefu Provider

The `kubefu` provider exposes Kubernetes API schemas as local, type-checked data sources and provides a single `kubefu_manifest` resource to apply manifests to a cluster. Generated data sources do not call a Kubernetes API; they validate and render manifests locally for use with `kubefu_manifest` or other delivery workflows.

Notes:
- Data sources are schema-driven and do not reach out to a cluster.
- `kubefu_manifest` uses server-side apply and requires kubeconfig access.
- Schema version fields are lists used only to emit compatibility warnings; they do not change which schemas are available.
- For Kustomize, versions correspond to the `kyaml/vX.Y.Z` schema tags.
- User-supplied schemas can be loaded via `schema_paths` (local CRD YAML or OpenAPI JSON). Terraform needs them at schema load time, so `KUBEFU_SCHEMA_PATHS` must be set and must match `schema_paths`.

## Example Usage

```terraform
provider "kubefu" {
  kubeconfig_path    = "~/.kube/config"
  kubeconfig_context = "dev"
}

data "kubefu_k8s_core_config_map_v1" "example" {
  metadata = {
    name      = "example"
    namespace = "default"
  }
  data = {
    hello = "world"
  }
}

resource "kubefu_manifest" "example" {
  manifest = data.kubefu_k8s_core_config_map_v1.example.kubefu_manifest_yaml
}
```

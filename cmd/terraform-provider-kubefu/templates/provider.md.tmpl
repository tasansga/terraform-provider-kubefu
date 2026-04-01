# kubefu Provider

The `kubefu` provider exposes Kubernetes API schemas as local, type-checked data sources and provides a single `kubefu_manifest` resource to apply manifests to a cluster. Generated data sources do not call a Kubernetes API; they validate and render manifests locally for use with `kubefu_manifest` or other delivery workflows.

Notes:
- Data sources are schema-driven and do not reach out to a cluster.
- `kubefu_manifest` uses server-side apply and requires kubeconfig access.
- Schema version fields are lists used only to emit compatibility warnings; they do not change which schemas are available.
- For Kustomize, versions correspond to the `kyaml/vX.Y.Z` schema tags.
- User-supplied schemas can be loaded via `schema_paths` (local CRD YAML or OpenAPI JSON). Terraform needs them at schema load time, so `KUBEFU_SCHEMA_PATHS` must be set and must match `schema_paths` exactly. Using absolute paths for both is recommended.
- Rendered manifests default to `compact` mode (omit empty/zero values). Set `manifest_render_mode = "canonical"` to preserve all values.

## Example Usage

```terraform
provider "kubefu" {
  kubeconfig_path    = "~/.kube/config"
  kubeconfig_context = "dev"
  manifest_render_mode = "compact"
}

data "kubefu_k8s_apps_deployment_v1" "example" {
  metadata {
    name      = "web"
    namespace = "default"
  }

  spec {
    replicas = 2

    selector {
      match_labels = {
        app = "web"
      }
    }

    template {
      metadata {
        labels = {
          app = "web"
        }
      }

      spec {
        container {
          name  = "web"
          image = "nginx:1.27"

          env {
            name = "DB_PASSWORD"
            value_from {
              secret_key_ref {
                name = "db-secrets"
                key  = "password"
              }
            }
          }

          liveness_probe {
            http_get {
              path = "/healthz"
              port = 8080
            }
          }
        }
      }
    }
  }
}

resource "kubefu_manifest" "example" {
  manifest = data.kubefu_k8s_apps_deployment_v1.example.kubefu_manifest_yaml
}
```

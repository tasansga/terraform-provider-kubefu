# terraform-provider-kubefu

`kubefu` is a Terraform provider that exposes Kubernetes API schemas as local, type-checked data sources and provides a single `kubefu_manifest` resource to apply manifests to a cluster. It is designed for GitOps-style workflows where you want schema validation without coupling to a specific controller.

Links:
- [kubefu homepage on GitHub](https://github.com/tasansga/terraform-provider-kubefu)
- [kubefu in the terraform registry](https://registry.terraform.io/providers/tasansga/kubefu/latest)
- [kubefu in the OpenTofu registry](https://search.opentofu.org/provider/tasansga/kubefu/latest)

## What it provides

- **Generated data sources** for Kubernetes, Flux, cert-manager, Prometheus Operator, Gateway API, and External Secrets Operator schemas.
- **`kubefu_manifest` resource** that applies YAML or JSON to a target cluster using server-side apply.
- **Rendered manifest outputs** on every data source: `kubefu_manifest_json` and `kubefu_manifest_yaml`.
- **Compact manifest rendering by default**: rendered manifests omit empty/zero values (set `manifest_render_mode = "canonical"` to keep all values).

The generated data sources do **not** call a Kubernetes API. They exist to validate and structure data locally. Use with `kubefu_manifest` to deploy to a cluster or with a git/github provider to write to git.

## Provider configuration

Optional version settings are used to emit compatibility warnings for generated data sources:

- `k8s_version`
- `flux_version`
- `cert_manager_version`
- `prometheus_operator_version`
- `gateway_api_version`
- `external_secrets_version`

Kubernetes access for `kubefu_manifest` is configured via:

- `kubeconfig_path` (defaults to `KUBECONFIG`)
- `kubeconfig_context` (optional)

## Example: Apply a manifest

```hcl
provider "kubefu" {
  kubeconfig_path = "~/.kube/config"
  kubeconfig_context = "dev"
}

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

## Example: Use a data source manifest

```hcl
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


## Development

### Schema downloads

Schemas are downloaded and stored under `schemas/`. Use the downloader to keep them up to date:

```bash
go run ./cmd/resourcegen download-schema
```

Or specify a subset:

```bash
go run ./cmd/resourcegen download-schema k8s flux cert-manager
```

### Regenerating data sources

Generated data sources live in `kubefu/generated` and are created from the schema directory:

```bash
go run ./cmd/resourcegen generate schemas
```

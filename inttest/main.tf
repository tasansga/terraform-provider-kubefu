terraform {
  required_providers {
    kubefu = {
      source  = "tasansga/kubefu"
      version = "0.1.0-test"
    }
  }
}

provider "kubefu" {
  schema_paths = [
    abspath("${path.module}/schemas/foo.yaml"),
    abspath("${path.module}/schemas/bar.yaml"),
  ]
}

variable "apply_manifests" {
  type        = bool
  default     = false
  description = "Apply kubefu_manifest resources to a live cluster when true."
}

data "kubefu_k8s_core_namespace_v1" "inttest" {
  metadata {
    name = "kubefu-inttest"
  }
}

data "kubefu_k8s_core_config_map_v1" "inttest" {
  metadata {
    name      = "kubefu-inttest"
    namespace = data.kubefu_k8s_core_namespace_v1.inttest.metadata[0].name
  }
  data = {
    hello = "world"
  }
}

data "kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1" "inttest" {
  metadata = {
    name      = "variables"
    namespace = data.kubefu_k8s_core_namespace_v1.inttest.metadata[0].name
  }

  spec {
    interval = "10m"
    path     = "./workloads/cisco-msp-golden-config/foo-messaging/variables"
    prune    = true

    source_ref {
      kind      = "GitRepository"
      name      = "cisco-msp-golden-config"
      namespace = "flux-system"
    }
  }
}

data "kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1" "inttest_explicit_empty" {
  metadata = {
    name      = "variables-explicit-empty"
    namespace = data.kubefu_k8s_core_namespace_v1.inttest.metadata[0].name
  }

  spec {
    interval = "10m"
    path     = "./workloads/cisco-msp-golden-config/foo-messaging/variables"
    prune    = true

    common_metadata {
      annotations = {}
    }

    source_ref {
      kind      = "GitRepository"
      name      = "cisco-msp-golden-config"
      namespace = "flux-system"
    }
  }
}

resource "kubefu_manifest" "inttest_namespace" {
  count    = var.apply_manifests ? 1 : 0
  manifest = data.kubefu_k8s_core_namespace_v1.inttest.kubefu_manifest_yaml
}

resource "kubefu_manifest" "inttest_config_map" {
  count      = var.apply_manifests ? 1 : 0
  manifest   = data.kubefu_k8s_core_config_map_v1.inttest.kubefu_manifest_yaml
  depends_on = [kubefu_manifest.inttest_namespace]
}

output "namespace_yaml" {
  value = data.kubefu_k8s_core_namespace_v1.inttest.kubefu_manifest_yaml
}

output "config_map_yaml" {
  value = data.kubefu_k8s_core_config_map_v1.inttest.kubefu_manifest_yaml
}

output "flux_kustomization_yaml" {
  value = data.kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1.inttest.kubefu_manifest_yaml
}

output "flux_kustomization_explicit_empty_yaml" {
  value = data.kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1.inttest_explicit_empty.kubefu_manifest_yaml
}

data "kubefu_user_crd_custom_test_tst_queue_v1beta1" "inttest_user" {}

output "user_schema_yaml" {
  value = data.kubefu_user_crd_custom_test_tst_queue_v1beta1.inttest_user.kubefu_manifest_yaml
}

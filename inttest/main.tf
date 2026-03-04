terraform {
  required_providers {
    kubefu = {
      source  = "tasansga/kubefu"
      version = "0.1.0-test"
    }
  }
}

provider "kubefu" {}

data "kubefu_k8s_core_namespace_v1" "inttest" {
  metadata = {
    name = "kubefu-inttest"
  }
}

data "kubefu_k8s_core_config_map_v1" "inttest" {
  metadata = {
    name      = "kubefu-inttest"
    namespace = data.kubefu_k8s_core_namespace_v1.inttest.metadata["name"]
  }
  data = {
    hello = "world"
  }
}

resource "kubefu_manifest" "inttest_namespace" {
  manifest = data.kubefu_k8s_core_namespace_v1.inttest.kubefu_manifest_yaml
}

resource "kubefu_manifest" "inttest_config_map" {
  manifest   = data.kubefu_k8s_core_config_map_v1.inttest.kubefu_manifest_yaml
  depends_on = [kubefu_manifest.inttest_namespace]
}

output "namespace_yaml" {
  value = data.kubefu_k8s_core_namespace_v1.inttest.kubefu_manifest_yaml
}

output "config_map_yaml" {
  value = data.kubefu_k8s_core_config_map_v1.inttest.kubefu_manifest_yaml
}

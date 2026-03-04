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

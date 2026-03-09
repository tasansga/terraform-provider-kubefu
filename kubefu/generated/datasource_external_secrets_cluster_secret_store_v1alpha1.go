package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1Read,
		Description: "ClusterSecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object.\nServers should convert recognized schemas to the latest internal value, and\nmay reject unrecognized values.\nMore info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents.\nServers may infer this from the endpoint the client submits requests to.\nCannot be updated.\nIn CamelCase.\nMore info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kubefu_manifest_json": {
				Type:        schema.TypeString,
				Description: "Rendered manifest (canonical JSON) for this data source.",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kubefu_manifest_yaml": {
				Type:        schema.TypeString,
				Description: "Rendered manifest (canonical YAML) for this data source.",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"metadata": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"spec": {
				Type:        schema.TypeMap,
				Description: "SecretStoreSpec defines the desired state of SecretStore.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeMap,
				Description: "SecretStoreStatus defines the observed state of the SecretStore.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "external-secrets.io/v1alpha1", "ClusterSecretStore", "external-secrets.io/v1alpha1/ClusterSecretStore"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "spec", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1CompatibleVersions = []string{
	"v0.7.0",
	"v0.7.1",
	"v0.7.2",
	"v0.7.3",
	"v0.8.0",
	"v0.8.1",
	"v0.8.2",
	"v0.8.3",
	"v0.8.4",
	"v0.8.5",
	"v0.8.6",
	"v0.8.7",
	"v0.8.8",
	"v0.8.9",
	"v0.8.10",
	"v0.8.11",
	"v0.8.12",
	"v0.8.13",
	"v0.8.14",
	"v0.8.15",
	"v0.8.16",
	"v0.8.17",
	"v0.9.0",
	"v0.9.1",
	"v0.9.2",
	"v0.9.3",
	"v0.9.4",
	"v0.9.5",
	"v0.9.6",
	"v0.9.7",
	"v0.9.8",
	"v0.9.9",
	"v0.9.10",
	"v0.9.11",
	"v0.9.12",
	"v0.9.13",
	"v0.9.14",
	"v0.9.15",
	"v0.9.16",
	"v0.9.17",
	"v0.9.18",
	"v0.9.19",
	"v0.9.20",
	"v0.10.0",
	"v0.10.1",
	"v0.10.2",
	"v0.10.3",
	"v0.10.4",
	"v0.10.5",
	"v0.10.6",
	"v0.10.7",
	"v0.11.0",
	"v0.12.0",
	"v0.12.1",
	"v0.13.0",
	"v0.14.0",
	"v0.14.1",
	"v0.14.2",
	"v0.14.3",
	"v0.14.4",
	"v0.15.0",
	"v0.15.1",
}

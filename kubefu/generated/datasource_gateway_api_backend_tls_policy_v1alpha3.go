package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3Read,
		Description: "BackendTLSPolicy provides a way to configure how a Gateway\nconnects to a Backend via TLS.",
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
				Description: "Spec defines the desired state of BackendTLSPolicy.",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"status": {
				Type:        schema.TypeMap,
				Description: "Status defines the current state of BackendTLSPolicy.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "gateway.networking.k8s.io/v1alpha3", "BackendTLSPolicy", "gateway.networking.k8s.io/v1alpha3/BackendTLSPolicy"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "spec", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3CompatibleVersions = []string{
	"v1.4.0",
	"v1.4.1",
}

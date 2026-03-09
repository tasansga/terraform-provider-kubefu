package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1Read,
		Description: "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
		Schema: map[string]*schema.Schema{
			"address_type": {
				Type:        schema.TypeString,
				Description: "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. Default is IP",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"endpoints": {
				Type:        schema.TypeList,
				Description: "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
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
				Description: "Standard object's metadata.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"ports": {
				Type:        schema.TypeList,
				Description: "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates \"all ports\". Each slice may include a maximum of 100 ports.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
		},
	}
}



func dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "discovery.k8s.io/v1alpha1", "EndpointSlice", "discovery.k8s.io/v1alpha1/EndpointSlice"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"address_type", "endpoints", "metadata", "ports"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1CompatibleVersions = []string{
	"v1.16.0",
}

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sResourceK8sIoResourceSliceV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sResourceK8sIoResourceSliceV1Alpha2Read,
		Description: "ResourceSlice provides information about available resources on individual nodes.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"driver_name": {
				Type:        schema.TypeString,
				Description: "DriverName identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name.",
				Optional:    false,
				Required:    true,
				Computed:    false,
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
				Description: "Standard object metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"named_resources": {
				Type:        schema.TypeMap,
				Description: "NamedResources describes available resources using the named resources model.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"node_name": {
				Type:        schema.TypeString,
				Description: "NodeName identifies the node which provides the resources if they are local to a node.\n\nA field selector can be used to list only ResourceSlice objects with a certain node name.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sResourceK8sIoResourceSliceV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "resource.k8s.io/v1alpha2", "ResourceSlice", "resource.k8s.io/v1alpha2/ResourceSlice"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"driver_name", "metadata", "named_resources", "node_name"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sResourceK8sIoResourceSliceV1Alpha2CompatibleVersions = []string{
	"v1.30.0",
}

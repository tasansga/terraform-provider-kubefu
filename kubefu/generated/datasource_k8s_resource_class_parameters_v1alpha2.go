package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2Read,
		Description: "ResourceClassParameters defines resource requests for a ResourceClass in an in-tree format understood by Kubernetes.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"filters": {
				Type:        schema.TypeList,
				Description: "Filters describes additional contraints that must be met when using the class.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
			"generated_from": {
				Type:        schema.TypeMap,
				Description: "If this object was created from some other resource, then this links back to that resource. This field is used to find the in-tree representation of the class parameters when the parameter reference of the class refers to some unknown type.",
				Optional:    true,
				Required:    false,
				Computed:    true,
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
			"vendor_parameters": {
				Type:        schema.TypeList,
				Description: "VendorParameters are arbitrary setup parameters for all claims using this class. They are ignored while allocating the claim. There must not be more than one entry per driver.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
		},
	}
}



func dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "resource.k8s.io/v1alpha2", "ResourceClassParameters", "resource.k8s.io/v1alpha2/ResourceClassParameters"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"filters", "generated_from", "metadata", "vendor_parameters"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2CompatibleVersions = []string{
	"v1.30.0",
}

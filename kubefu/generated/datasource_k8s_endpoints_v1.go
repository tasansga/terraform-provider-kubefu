package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sCoreEndpointsV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sCoreEndpointsV1Read,
		Description: "Endpoints is a collection of endpoints that implement the actual service. Example:\n  Name: \"mysvc\",\n  Subsets: [\n    {\n      Addresses: [{\"ip\": \"10.10.1.1\"}, {\"ip\": \"10.10.2.2\"}],\n      Ports: [{\"name\": \"a\", \"port\": 8675}, {\"name\": \"b\", \"port\": 309}]\n    },\n    {\n      Addresses: [{\"ip\": \"10.10.3.3\"}],\n      Ports: [{\"name\": \"a\", \"port\": 93}, {\"name\": \"b\", \"port\": 76}]\n    },\n ]",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
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
				Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"subsets": {
				Type:        schema.TypeList,
				Description: "The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
		},
	}
}



func dataSourceK8sCoreEndpointsV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "v1", "Endpoints", "core/v1/Endpoints"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "subsets"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sCoreEndpointsV1CompatibleVersions = []string{
	"v1.7.0",
	"v1.8.0",
	"v1.9.0",
	"v1.10.0",
	"v1.11.0",
	"v1.12.0",
	"v1.13.0",
	"v1.14.0",
	"v1.15.0",
	"v1.16.0",
	"v1.17.0",
	"v1.18.0",
	"v1.19.0",
	"v1.20.0",
	"v1.21.0",
	"v1.22.0",
	"v1.23.0",
	"v1.24.0",
	"v1.25.0",
	"v1.26.0",
	"v1.27.0",
	"v1.28.0",
	"v1.29.0",
	"v1.30.0",
	"v1.31.0",
	"v1.32.0",
	"v1.33.0",
	"v1.34.0",
	"v1.35.0",
}

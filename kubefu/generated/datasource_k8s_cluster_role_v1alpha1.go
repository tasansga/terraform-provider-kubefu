package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

func dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1Read,
		Description: "ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.",
		Schema: map[string]*schema.Schema{
			"aggregation_rule": {
				Type:        schema.TypeMap,
				Description: "AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
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
				Description: "Standard object's metadata.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"rules": {
				Type:        schema.TypeList,
				Description: "Rules holds all the PolicyRules for this ClusterRole",
				Optional:    false,
				Required:    true,
				Computed:    false,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
		},
	}
}



func dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "rbac.authorization.k8s.io/v1alpha1", "ClusterRole", "rbac.authorization.k8s.io/v1alpha1/ClusterRole"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"aggregation_rule", "metadata", "rules"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1CompatibleVersions = []string{
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
}

func dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1IsCompatibleWith(version string) bool {
	return versionpkg.IsCompatibleWith(version, dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1CompatibleVersions)
}

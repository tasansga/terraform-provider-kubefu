package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1Read,
		Description: "ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.",
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
				Description: "Standard object's metadata.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"role_ref": {
				Type:        schema.TypeMap,
				Description: "RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"subjects": {
				Type:        schema.TypeList,
				Description: "Subjects holds references to the objects the role applies to.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
		},
	}
}



func dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "rbac.authorization.k8s.io/v1beta1", "ClusterRoleBinding", "rbac.authorization.k8s.io/v1beta1/ClusterRoleBinding"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "role_ref", "subjects"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1CompatibleVersions = []string{
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
}

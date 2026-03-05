package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

func dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1Read,
		Description: "PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "description is an arbitrary string that usually provides guidelines on when this priority class should be used.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"global_default": {
				Type:        schema.TypeBool,
				Description: "globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
				Optional:    true,
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
			"value": {
				Type:        schema.TypeInt,
				Description: "The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
		},
	}
}



func dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "scheduling.k8s.io/v1beta1", "PriorityClass", "scheduling.k8s.io/v1beta1/PriorityClass"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"description", "global_default", "metadata", "value"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1CompatibleVersions = []string{
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

func dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1IsCompatibleWith(version string) bool {
	return versionpkg.IsCompatibleWith(version, dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1CompatibleVersions)
}

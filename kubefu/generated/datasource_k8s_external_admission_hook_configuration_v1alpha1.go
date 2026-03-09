package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1Read,
		Description: "ExternalAdmissionHookConfiguration describes the configuration of initializers.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"external_admission_hooks": {
				Type:        schema.TypeList,
				Description: "ExternalAdmissionHooks is a list of external admission webhooks and the affected resources and operations.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
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
				Description: "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "admissionregistration.k8s.io/v1alpha1", "ExternalAdmissionHookConfiguration", "admissionregistration.k8s.io/v1alpha1/ExternalAdmissionHookConfiguration"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"external_admission_hooks", "metadata"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1CompatibleVersions = []string{
	"v1.7.0",
	"v1.8.0",
}

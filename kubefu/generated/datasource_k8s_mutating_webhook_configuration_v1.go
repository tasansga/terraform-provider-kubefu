package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Read,
		Description: "MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
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
				Description: "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"webhooks": {
				Type:        schema.TypeList,
				Description: "Webhooks is a list of webhooks and the affected resources and operations.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
		},
	}
}



func dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", "admissionregistration.k8s.io/v1/MutatingWebhookConfiguration"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "webhooks"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1CompatibleVersions = []string{
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

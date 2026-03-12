package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1Read,
		Description: "PrometheusRule defines recording and alerting rules for a Prometheus instance",
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
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"spec": {
				Type:        schema.TypeList,
				Description: "Specification of desired alerting rule definitions for Prometheus.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"groups": {
						Type:        schema.TypeList,
						Description: "Content of Prometheus rule file",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"interval": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"partial_response_strategy": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"rules": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"alert": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"annotations": {
										Type:        schema.TypeMap,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"expr": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"for": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"labels": {
										Type:        schema.TypeMap,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"record": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
				}},
			},
		},
	}
}



func dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "monitoring.coreos.com/v1", "PrometheusRule", "monitoring.coreos.com/v1/PrometheusRule"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec"}, []string{"spec"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1CompatibleVersions = []string{
	"v0.54.0",
	"v0.54.1",
	"v0.55.0",
	"v0.55.1",
	"v0.56.0",
	"v0.56.1",
	"v0.56.2",
	"v0.56.3",
	"v0.57.0",
	"v0.58.0",
	"v0.59.0",
	"v0.59.1",
	"v0.59.2",
	"v0.60.0",
	"v0.60.1",
	"v0.61.0",
	"v0.61.1",
	"v0.62.0",
	"v0.63.0",
	"v0.64.0",
	"v0.64.1",
	"v0.65.0",
	"v0.65.1",
	"v0.65.2",
	"v0.66.0",
	"v0.67.0",
	"v0.67.1",
	"v0.68.0",
	"v0.69.0",
	"v0.69.1",
	"v0.70.0",
	"v0.71.0",
	"v0.71.1",
	"v0.71.2",
	"v0.72.0",
	"v0.73.0",
	"v0.73.1",
	"v0.73.2",
	"v0.74.0",
	"v0.75.0",
	"v0.75.1",
	"v0.75.2",
	"v0.76.0",
	"v0.76.1",
	"v0.76.2",
	"v0.77.0",
	"v0.77.1",
	"v0.77.2",
	"v0.78.0",
	"v0.78.1",
	"v0.78.2",
	"v0.79.0",
	"v0.79.1",
	"v0.79.2",
	"v0.80.0",
	"v0.80.1",
	"v0.81.0",
	"v0.82.0",
	"v0.82.1",
	"v0.82.2",
	"v0.83.0",
	"v0.84.0",
	"v0.84.1",
	"v0.85.0",
	"v0.86.0",
	"v0.86.1",
	"v0.86.2",
	"v0.87.0",
	"v0.87.1",
	"v0.88.0",
	"v0.88.1",
	"v0.89.0",
}

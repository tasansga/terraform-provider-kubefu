package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1Read,
		Description: "ScrapeConfig defines a namespaced Prometheus scrape_config to be aggregated across multiple namespaces into the Prometheus configuration.",
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
				Description: "ScrapeConfigSpec is a specification of the desired configuration for a scrape configuration.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"file_sd_configs": {
						Type:        schema.TypeList,
						Description: "FileSDConfigs defines a list of file service discovery configurations.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"files": {
								Type:        schema.TypeList,
								Description: "List of files to be used for file discovery. Recommendation: use absolute paths. While relative paths work, the prometheus-operator project makes no guarantees about the working directory where the configuration file is stored. Files must be mounted using Prometheus.ConfigMaps or Prometheus.Secrets.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"refresh_interval": {
								Type:        schema.TypeString,
								Description: "RefreshInterval configures the refresh interval at which Prometheus will reload the content of the files.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"honor_labels": {
						Type:        schema.TypeBool,
						Description: "HonorLabels chooses the metric's labels on collisions with target labels.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"honor_timestamps": {
						Type:        schema.TypeBool,
						Description: "HonorTimestamps controls whether Prometheus respects the timestamps present in scraped data.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"http_sd_configs": {
						Type:        schema.TypeList,
						Description: "HTTPSDConfigs defines a list of HTTP service discovery configurations.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"refresh_interval": {
								Type:        schema.TypeString,
								Description: "RefreshInterval configures the refresh interval at which Prometheus will re-query the endpoint to update the target list.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"url": {
								Type:        schema.TypeString,
								Description: "URL from which the targets are fetched.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"metrics_path": {
						Type:        schema.TypeString,
						Description: "MetricsPath HTTP path to scrape for metrics. If empty, Prometheus uses the default value (e.g. /metrics).",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"relabelings": {
						Type:        schema.TypeList,
						Description: "RelabelConfigs defines how to rewrite the target's labels before scraping. Prometheus Operator automatically adds relabelings for a few standard Kubernetes fields. The original scrape job's name is available via the `__tmp_prometheus_job_name` label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"action": {
								Type:        schema.TypeString,
								Description: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"modulus": {
								Type:        schema.TypeInt,
								Description: "Modulus to take of the hash of the source label values.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"regex": {
								Type:        schema.TypeString,
								Description: "Regular expression against which the extracted value is matched. Default is '(.*)'",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"replacement": {
								Type:        schema.TypeString,
								Description: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"separator": {
								Type:        schema.TypeString,
								Description: "Separator placed between concatenated source label values. default is ';'.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"source_labels": {
								Type:        schema.TypeList,
								Description: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"target_label": {
								Type:        schema.TypeString,
								Description: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"static_configs": {
						Type:        schema.TypeList,
						Description: "StaticConfigs defines a list of static targets with a common label set.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"labels": {
								Type:        schema.TypeMap,
								Description: "Labels assigned to all metrics scraped from the targets.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"targets": {
								Type:        schema.TypeList,
								Description: "List of targets for this static configuration.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
						}},
					},
				}},
			},
		},
	}
}



func dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "monitoring.coreos.com/v1alpha1", "ScrapeConfig", "monitoring.coreos.com/v1alpha1/ScrapeConfig"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec"}, []string{"spec"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1CompatibleVersions = []string{
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

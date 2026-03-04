package kubefu

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/tasansga/terraform-provider-kubefu/kubefu/generated"
)

type providerConfig struct {
	K8sVersion                string
	FluxVersion               string
	CertManagerVersion        string
	PrometheusOperatorVersion string
	GatewayAPIVersion         string
	ExternalSecretsVersion    string
	KubeConfigPath            string
	KubeContext               string
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"k8s_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional Kubernetes schema version to target",
			},
			"flux_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional Flux schema version to target",
			},
			"cert_manager_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional cert-manager schema version to target",
			},
			"prometheus_operator_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional Prometheus Operator schema version to target",
			},
			"gateway_api_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional Gateway API schema version to target",
			},
			"external_secrets_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional External Secrets Operator schema version to target",
			},
			"kubeconfig_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KUBECONFIG", ""),
				Description: "Path to the kubeconfig file; defaults to the KUBECONFIG environment variable",
			},
			"kubeconfig_context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional kubeconfig context to use",
			},
		},
		ResourcesMap: resources(),
	}

	p.DataSourcesMap = generated.DataSources(generated.Versions{})
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		cfg := &providerConfig{
			K8sVersion:                d.Get("k8s_version").(string),
			FluxVersion:               d.Get("flux_version").(string),
			CertManagerVersion:        d.Get("cert_manager_version").(string),
			PrometheusOperatorVersion: d.Get("prometheus_operator_version").(string),
			GatewayAPIVersion:         d.Get("gateway_api_version").(string),
			ExternalSecretsVersion:    d.Get("external_secrets_version").(string),
			KubeConfigPath:            d.Get("kubeconfig_path").(string),
			KubeContext:               d.Get("kubeconfig_context").(string),
		}
		p.DataSourcesMap = generated.DataSources(generated.Versions{
			K8sVersion:                cfg.K8sVersion,
			FluxVersion:               cfg.FluxVersion,
			CertManagerVersion:        cfg.CertManagerVersion,
			PrometheusOperatorVersion: cfg.PrometheusOperatorVersion,
			GatewayAPIVersion:         cfg.GatewayAPIVersion,
			ExternalSecretsVersion:    cfg.ExternalSecretsVersion,
		})
		return cfg, warnIfClusterVersionMismatch(ctx, cfg)
	}

	return p
}

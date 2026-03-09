package kubefu

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/tasansga/terraform-provider-kubefu/kubefu/generated"
)

type providerConfig struct {
	K8sVersions                []string
	FluxVersions               []string
	CertManagerVersions        []string
	PrometheusOperatorVersions []string
	GatewayAPIVersions         []string
	ExternalSecretsVersions    []string
	KustomizeVersions          []string
	KubeConfigPath            string
	KubeContext               string
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"k8s_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Kubernetes schema versions to target",
			},
			"flux_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Flux schema versions to target",
			},
			"cert_manager_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of cert-manager schema versions to target",
			},
			"prometheus_operator_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Prometheus Operator schema versions to target",
			},
			"gateway_api_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Gateway API schema versions to target",
			},
			"external_secrets_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of External Secrets Operator schema versions to target",
			},
			"kustomize_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Kustomize schema versions to target",
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
			K8sVersions:                getStringList(d, "k8s_version"),
			FluxVersions:               getStringList(d, "flux_version"),
			CertManagerVersions:        getStringList(d, "cert_manager_version"),
			PrometheusOperatorVersions: getStringList(d, "prometheus_operator_version"),
			GatewayAPIVersions:         getStringList(d, "gateway_api_version"),
			ExternalSecretsVersions:    getStringList(d, "external_secrets_version"),
			KustomizeVersions:          getStringList(d, "kustomize_version"),
			KubeConfigPath:            d.Get("kubeconfig_path").(string),
			KubeContext:               d.Get("kubeconfig_context").(string),
		}
		p.DataSourcesMap = generated.DataSources(generated.Versions{
			K8sVersions:                cfg.K8sVersions,
			FluxVersions:               cfg.FluxVersions,
			CertManagerVersions:        cfg.CertManagerVersions,
			PrometheusOperatorVersions: cfg.PrometheusOperatorVersions,
			GatewayAPIVersions:         cfg.GatewayAPIVersions,
			ExternalSecretsVersions:    cfg.ExternalSecretsVersions,
			KustomizeVersions:          cfg.KustomizeVersions,
		})
		return cfg, warnIfClusterVersionMismatch(ctx, cfg)
	}

	return p
}

func getStringList(d *schema.ResourceData, key string) []string {
	raw := d.Get(key)
	if raw == nil {
		return nil
	}
	items, ok := raw.([]interface{})
	if !ok {
		return nil
	}
	values := make([]string, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		if s, ok := item.(string); ok && s != "" {
			values = append(values, s)
		}
	}
	return values
}

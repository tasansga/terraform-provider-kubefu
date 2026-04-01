package kubefu

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/tasansga/terraform-provider-kubefu/kubefu/generated"
	"github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

type providerConfig struct {
	K8sVersions                []string
	FluxVersions               []string
	CertManagerVersions        []string
	PrometheusOperatorVersions []string
	GatewayAPIVersions         []string
	ExternalSecretsVersions    []string
	KustomizeVersions          []string
	KarpenterAWSVersions       []string
	KarpenterCoreVersions      []string
	KubeConfigPath             string
	KubeContext                string
	SchemaPaths                []string
	ManifestRenderModeValue    string
}

func (c *providerConfig) ManifestRenderMode() string {
	if c == nil {
		return manifest.RenderModeCompact
	}
	return c.ManifestRenderModeValue
}

func Provider() *schema.Provider {
	initSchemaPaths := parseSchemaPathsEnv(os.Getenv("KUBEFU_SCHEMA_PATHS"))
	userDataSources, userInitDiags := loadUserSchemas(initSchemaPaths)
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
			"karpenter_aws_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Karpenter AWS schema versions to target",
			},
			"karpenter_core_version": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Optional list of Karpenter core schema versions to target",
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
			"schema_paths": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				DefaultFunc: schemaPathsDefaultFunc,
				Description: "Optional list of local schema files or directories to load (CRD YAML or OpenAPI JSON). Must match KUBEFU_SCHEMA_PATHS exactly; absolute paths are recommended.",
			},
			"manifest_render_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      manifest.RenderModeCompact,
				ValidateFunc: validation.StringInSlice([]string{manifest.RenderModeCompact, manifest.RenderModeCanonical}, false),
				Description:  "Manifest rendering mode for kubefu_manifest_json/kubefu_manifest_yaml. Use \"compact\" to omit empty/zero values, or \"canonical\" to preserve all values.",
			},
		},
		ResourcesMap: resources(),
	}

	dataSources := generated.DataSources(generated.Versions{})
	registeredUserDataSourceNames := make([]string, 0, len(userDataSources))
	for key, resource := range userDataSources {
		if _, exists := dataSources[key]; exists {
			userInitDiags = append(userInitDiags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "User schema ignored",
				Detail:   fmt.Sprintf("data source %q already exists in built-in schemas", key),
			})
			continue
		}
		dataSources[key] = resource
		registeredUserDataSourceNames = append(registeredUserDataSourceNames, key)
	}
	if ds := userSchemasRegisteredDiagnostic(registeredUserDataSourceNames); ds != nil {
		userInitDiags = append(userInitDiags, *ds)
	}
	dataSources[registeredUserSchemasDataSourceName] = dataSourceRegisteredUserSchemas(registeredUserDataSourceNames)
	p.DataSourcesMap = dataSources
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		configSchemaPaths := getStringList(d, "schema_paths")
		if len(initSchemaPaths) == 0 && len(configSchemaPaths) > 0 {
			userInitDiags = append(userInitDiags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "KUBEFU_SCHEMA_PATHS must be set",
				Detail:   "Terraform requires data sources to be registered before the provider schema is served. Set KUBEFU_SCHEMA_PATHS and keep schema_paths identical to avoid schema mismatches.",
			})
		} else if len(initSchemaPaths) > 0 && !sameStringSet(configSchemaPaths, initSchemaPaths) {
			userInitDiags = append(userInitDiags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "schema_paths must match KUBEFU_SCHEMA_PATHS",
				Detail:   "Terraform requires data sources to be registered before the provider schema is served. Set KUBEFU_SCHEMA_PATHS and schema_paths to the same values.",
			})
		}
		cfg := &providerConfig{
			K8sVersions:                getStringList(d, "k8s_version"),
			FluxVersions:               getStringList(d, "flux_version"),
			CertManagerVersions:        getStringList(d, "cert_manager_version"),
			PrometheusOperatorVersions: getStringList(d, "prometheus_operator_version"),
			GatewayAPIVersions:         getStringList(d, "gateway_api_version"),
			ExternalSecretsVersions:    getStringList(d, "external_secrets_version"),
			KustomizeVersions:          getStringList(d, "kustomize_version"),
			KarpenterAWSVersions:       getStringList(d, "karpenter_aws_version"),
			KarpenterCoreVersions:      getStringList(d, "karpenter_core_version"),
			KubeConfigPath:             d.Get("kubeconfig_path").(string),
			KubeContext:                d.Get("kubeconfig_context").(string),
			SchemaPaths:                configSchemaPaths,
			ManifestRenderModeValue:    d.Get("manifest_render_mode").(string),
		}
		versioned := generated.DataSources(generated.Versions{
			K8sVersions:                cfg.K8sVersions,
			FluxVersions:               cfg.FluxVersions,
			CertManagerVersions:        cfg.CertManagerVersions,
			PrometheusOperatorVersions: cfg.PrometheusOperatorVersions,
			GatewayAPIVersions:         cfg.GatewayAPIVersions,
			ExternalSecretsVersions:    cfg.ExternalSecretsVersions,
			KustomizeVersions:          cfg.KustomizeVersions,
			KarpenterAWSVersions:       cfg.KarpenterAWSVersions,
			KarpenterCoreVersions:      cfg.KarpenterCoreVersions,
		})
		for key, resource := range userDataSources {
			if _, exists := versioned[key]; exists {
				continue
			}
			versioned[key] = resource
		}
		versioned[registeredUserSchemasDataSourceName] = dataSourceRegisteredUserSchemas(registeredUserDataSourceNames)
		p.DataSourcesMap = versioned
		diags := append(userInitDiags, warnIfClusterVersionMismatch(ctx, cfg)...)
		return cfg, diags
	}

	return p
}

func getStringList(d *schema.ResourceData, key string) []string {
	raw := d.Get(key)
	if raw == nil {
		return nil
	}
	if value, ok := raw.(string); ok {
		return parseSchemaPathsEnv(value)
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

func schemaPathsDefaultFunc() (any, error) {
	value := os.Getenv("KUBEFU_SCHEMA_PATHS")
	if value == "" {
		return []interface{}{}, nil
	}
	parts := parseSchemaPathsEnv(value)
	result := make([]interface{}, 0, len(parts))
	for _, part := range parts {
		if part == "" {
			continue
		}
		result = append(result, part)
	}
	return result, nil
}

func parseSchemaPathsEnv(value string) []string {
	if value == "" {
		return nil
	}
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == ',' || r == ';' || r == rune(os.PathListSeparator)
	})
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			values = append(values, trimmed)
		}
	}
	return values
}

func normalizeStringSet(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(values))
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		normalized = append(normalized, value)
	}
	sort.Strings(normalized)
	return normalized
}

func sameStringSet(left []string, right []string) bool {
	ln := normalizeStringSet(left)
	rn := normalizeStringSet(right)
	if len(ln) != len(rn) {
		return false
	}
	for i := range ln {
		if ln[i] != rn[i] {
			return false
		}
	}
	return true
}

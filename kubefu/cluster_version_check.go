package kubefu

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"k8s.io/client-go/discovery"
)

const defaultDiscoveryTimeout = 5 * time.Second

func warnIfClusterVersionMismatch(ctx context.Context, cfg *providerConfig) diag.Diagnostics {
	if cfg == nil || strings.TrimSpace(cfg.K8sVersion) == "" {
		return nil
	}
	restCfg, err := buildRESTConfig(cfg)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to determine Kubernetes cluster version",
			Detail:   err.Error(),
		}}
	}
	shallow := *restCfg
	shallow.Timeout = defaultDiscoveryTimeout
	restCfg = &shallow
	disco, err := discovery.NewDiscoveryClientForConfig(restCfg)
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to determine Kubernetes cluster version",
			Detail:   err.Error(),
		}}
	}
	version, err := disco.ServerVersion()
	if err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Unable to determine Kubernetes cluster version",
			Detail:   err.Error(),
		}}
	}
	serverVersion := strings.TrimSpace(version.GitVersion)
	configured := strings.TrimSpace(cfg.K8sVersion)
	if serverVersion != "" && configured != "" && serverVersion != configured {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Configured Kubernetes version does not match cluster",
			Detail:   fmt.Sprintf("k8s_version is set to %q, but the cluster reports %q", configured, serverVersion),
		}}
	}
	return nil
}

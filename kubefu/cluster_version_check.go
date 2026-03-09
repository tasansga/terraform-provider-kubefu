package kubefu

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
	"k8s.io/client-go/discovery"
)

const defaultDiscoveryTimeout = 5 * time.Second

func warnIfClusterVersionMismatch(ctx context.Context, cfg *providerConfig) diag.Diagnostics {
	if cfg == nil || len(cfg.K8sVersions) == 0 {
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
	serverVersion := versionpkg.Canonical(version.GitVersion)
	if serverVersion == "" {
		return nil
	}
	configured := versionpkg.NormalizeList(cfg.K8sVersions)
	if len(configured) == 0 {
		return nil
	}
	if incompatible := versionpkg.FilterIncompatible([]string{serverVersion}, configured); len(incompatible) > 0 {
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Configured Kubernetes version does not match cluster",
			Detail:   fmt.Sprintf("k8s_version does not include cluster version %q (configured: %s)", serverVersion, strings.Join(configured, ", ")),
		}}
	}
	return nil
}

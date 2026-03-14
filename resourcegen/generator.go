package resourcegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

type dataSourceEntry struct {
	Key                  string
	FuncName             string
	Provider             string
	HasSupportedVersions bool
}

// GenerateFromSchemas reads the schema directory tree rooted at schemaRoot, builds
// definitions per provider, and writes versioned data source files plus an
// aggregated data source registry into outDir.
func GenerateFromSchemas(schemaRoot, outDir, pkgName string) error {
	providers, providerNames, err := providerDefinitions(schemaRoot)
	if err != nil {
		return err
	}
	if len(providerNames) == 0 {
		return fmt.Errorf("no definitions found under %s", schemaRoot)
	}
	if err := os.RemoveAll(outDir); err != nil {
		return fmt.Errorf("remove output dir %s: %w", outDir, err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create output dir %s: %w", outDir, err)
	}
	var entries []dataSourceEntry
	for _, provider := range providerNames {
		defs := providers[provider]
		if len(defs) == 0 {
			continue
		}
		providerName := providerSegment(provider)
		baseNameCounts := dataSourceBaseNameCounts(defs, providerName)
		for _, def := range defs {
			fileData, err := def.AsDataSource(pkgName, providerName)
			if err != nil {
				return fmt.Errorf("render provider %s kind %s: %w", provider, def.Kind, err)
			}
			fileData.Name = dataSourceFinalFileName(providerName, def, baseNameCounts)
			path := filepath.Join(outDir, fileData.Name)
			if err := os.WriteFile(path, fileData.Content, 0o644); err != nil {
				return fmt.Errorf("write file %s: %w", path, err)
			}
			entries = append(entries, dataSourceEntry{
				Key:                  dataSourceKey(providerName, def.Group, def.Kind, def.Version),
				FuncName:             def.DataSourceFuncName(providerName),
				Provider:             providerName,
				HasSupportedVersions: len(def.ProviderVersions) > 0,
			})
		}
	}
	if len(entries) == 0 {
		return fmt.Errorf("no definitions found under %s", schemaRoot)
	}
	if err := writeDataSourcesFile(outDir, pkgName, entries); err != nil {
		return err
	}
	legacyPath := filepath.Join(outDir, "manifest_helpers.go")
	if err := os.Remove(legacyPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove legacy manifest helpers %s: %w", legacyPath, err)
	}
	return nil
}

func dataSourceKey(providerName, group, kind, version string) string {
	return fmt.Sprintf("kubefu_%s_%s_%s_%s", providerName, groupSegment(group), toSnakeCase(kind), versionSegment(version))
}

func dataSourceBaseNameCounts(defs []Definition, providerName string) map[string]int {
	counts := make(map[string]int, len(defs))
	for _, def := range defs {
		base := dataSourceBaseFileName(providerName, def.Kind, def.Version)
		counts[base]++
	}
	return counts
}

func dataSourceBaseFileName(providerName, kind, version string) string {
	return fmt.Sprintf("datasource_%s_%s_%s.go", providerName, toSnakeCase(kind), versionSegment(version))
}

func dataSourceFinalFileName(providerName string, def Definition, baseCounts map[string]int) string {
	base := dataSourceBaseFileName(providerName, def.Kind, def.Version)
	if baseCounts != nil && baseCounts[base] > 1 {
		return fmt.Sprintf("datasource_%s_%s_%s_%s.go", providerName, groupSegment(def.Group), toSnakeCase(def.Kind), versionSegment(def.Version))
	}
	return base
}

func writeDataSourcesFile(dir, pkgName string, entries []dataSourceEntry) error {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Key < entries[j].Key
	})
	data := providerTemplateData{
		Package: pkgName,
		Entries: entries,
	}
	var buf bytes.Buffer
	if err := dataSourcesTemplate.Execute(&buf, data); err != nil {
		return fmt.Errorf("render data sources template: %w", err)
	}
	return os.WriteFile(filepath.Join(dir, "datasources.go"), buf.Bytes(), 0o644)
}

type providerTemplateData struct {
	Package string
	Entries []dataSourceEntry
}

var dataSourcesTemplate = template.Must(template.New("datasources").Parse(`package {{ .Package }}

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

type Versions struct {
	K8sVersions                []string
	FluxVersions               []string
	CertManagerVersions        []string
	PrometheusOperatorVersions []string
	GatewayAPIVersions         []string
	ExternalSecretsVersions    []string
	KustomizeVersions          []string
	KarpenterAWSVersions       []string
	KarpenterCoreVersions      []string
}

func DataSources(versions Versions) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource, {{ len .Entries }})
	{{- range .Entries }}
	{
		ds := {{ .FuncName }}()
		configured := versions.versionFor("{{ .Provider }}")
		{{- if .HasSupportedVersions }}
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, {{ .FuncName }}CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"{{ .Key }}",
				"{{ .Provider }}",
				strings.Join({{ .FuncName }}CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		{{- end }}
		result["{{ .Key }}"] = ds
	}
	{{- end }}
	return result
}

func (v Versions) versionFor(provider string) []string {
	switch provider {
	case "k8s":
		return v.K8sVersions
	case "flux":
		return v.FluxVersions
	case "cert_manager", "cert-manager":
		return v.CertManagerVersions
	case "prometheus_operator", "prometheus-operator", "prometheus":
		return v.PrometheusOperatorVersions
	case "gateway_api", "gateway-api", "gateway":
		return v.GatewayAPIVersions
	case "external_secrets", "external-secrets", "externalsecrets":
		return v.ExternalSecretsVersions
	case "kustomize":
		return v.KustomizeVersions
	case "karpenter_aws", "karpenter-aws":
		return v.KarpenterAWSVersions
	case "karpenter_core", "karpenter-core":
		return v.KarpenterCoreVersions
	default:
		return nil
	}
}
`))

func providerDefinitions(schemaRoot string) (map[string][]Definition, []string, error) {
	providers, err := CollectDefinitions(schemaRoot)
	if err != nil {
		return nil, nil, err
	}
	providerNames := make([]string, 0, len(providers))
	for provider := range providers {
		providerNames = append(providerNames, provider)
	}
	sort.Strings(providerNames)
	return providers, providerNames, nil
}

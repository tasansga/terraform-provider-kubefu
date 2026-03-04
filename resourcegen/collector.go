package resourcegen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

// CollectDefinitions walks the given schema root and returns the parsed
// definitions grouped by provider (subdirectory or the root). Each provider
// directory may contain Kubernetes OpenAPI JSON specs or CRD YAML manifests.
func CollectDefinitions(root string) (map[string][]Definition, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("stat schema root %s: %w", root, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", root)
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("read schema root %s: %w", root, err)
	}
	hasSubdirs := false
	for _, entry := range entries {
		if entry.IsDir() {
			hasSubdirs = true
			break
		}
	}
	results := make(map[string][]Definition)
	if hasSubdirs {
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			providerPath := filepath.Join(root, entry.Name())
			providerName := providerSegment(entry.Name())
			defs, err := collectProviderDefinitions(providerPath, providerName)
			if err != nil {
				return nil, fmt.Errorf("collect definitions for provider %s: %w", entry.Name(), err)
			}
			if len(defs) == 0 {
				continue
			}
			results[providerName] = defs
		}
	} else {
		name := filepath.Base(root)
		providerName := providerSegment(name)
		defs, err := collectProviderDefinitions(root, providerName)
		if err != nil {
			return nil, err
		}
		if len(defs) > 0 {
			results[providerName] = defs
		}
	}
	return results, nil
}

func collectProviderDefinitions(dir, provider string) ([]Definition, error) {
	collector := newDefinitionCollector()
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		version := version.FromFilename(d.Name())
		switch strings.ToLower(filepath.Ext(d.Name())) {
		case ".json":
			if err := processOpenAPISpec(path, provider, version, collector); err != nil {
				return fmt.Errorf("process OpenAPI spec %s: %w", path, err)
			}
		case ".yaml", ".yml":
			if err := processCRDFile(path, provider, version, collector); err != nil {
				return fmt.Errorf("process CRD file %s: %w", path, err)
			}
		default:
			// ignore other files
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	defs := collector.list()
	sort.Slice(defs, func(i, j int) bool {
		return definitionID(defs[i]) < definitionID(defs[j])
	})
	return defs, nil
}

func processOpenAPISpec(path, provider, providerVersion string, collector *definitionCollector) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read spec: %w", err)
	}
	return processOpenAPISpecBytes(data, provider, providerVersion, collector)
}

func processOpenAPISpecBytes(data []byte, provider, providerVersion string, collector *definitionCollector) error {
	entries, err := ParseRootResources(data)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		collector.add(Definition{
			Kind:           entry.Kind,
			Group:          entry.Group,
			Version:        entry.Version,
			Description:    entry.DefinitionDescription,
			DefinitionName: entry.DefinitionName,
			Schema:         entry.Schema,
		}, provider, providerVersion)
	}
	return nil
}

func processCRDFile(path, provider, providerVersion string, collector *definitionCollector) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open CRD file: %w", err)
	}
	defer f.Close()
	return decodeCRD(f, provider, providerVersion, collector)
}

func definitionsFromCRD(crd customResourceDefinition) []Definition {
	if crd.Kind != "CustomResourceDefinition" {
		return nil
	}
	group := crd.Spec.Group
	kind := crd.Spec.Names.Kind
	if kind == "" {
		return nil
	}
	var versions []crdVersion
	if len(crd.Spec.Versions) > 0 {
		versions = crd.Spec.Versions
	} else if crd.Spec.Version != "" {
		versions = []crdVersion{
			{
				Name: crd.Spec.Version,
				Schema: crdVersionSchema{
					OpenAPIV3Schema: crd.Spec.Validation.OpenAPIV3Schema,
				},
			},
		}
	}
	var definitions []Definition
	for _, v := range versions {
		if v.Schema.OpenAPIV3Schema.Type == "" && len(v.Schema.OpenAPIV3Schema.Properties) == 0 {
			continue
		}
		name := fmt.Sprintf("crd.%s.%s.%s", group, v.Name, kind)
		schemaMap := buildTerraformSchema(name, v.Schema.OpenAPIV3Schema, v.Schema.OpenAPIV3Schema.Definitions)
		definitions = append(definitions, Definition{
			Kind:           kind,
			Group:          group,
			Version:        v.Name,
			Description:    v.Schema.OpenAPIV3Schema.Description,
			DefinitionName: name,
			Schema:         schemaMap,
		})
	}
	return definitions
}

type customResourceDefinition struct {
	Kind string `json:"kind" yaml:"kind"`
	Spec struct {
		Group string `json:"group" yaml:"group"`
		Names struct {
			Kind string `json:"kind" yaml:"kind"`
		} `json:"names" yaml:"names"`
		Versions   []crdVersion     `json:"versions" yaml:"versions"`
		Version    string           `json:"version" yaml:"version"`
		Validation crdVersionSchema `json:"validation" yaml:"validation"`
	} `json:"spec" yaml:"spec"`
}

type crdVersion struct {
	Name   string           `json:"name" yaml:"name"`
	Schema crdVersionSchema `json:"schema" yaml:"schema"`
}

type crdVersionSchema struct {
	OpenAPIV3Schema definition `json:"openAPIV3Schema" yaml:"openAPIV3Schema"`
}

type definitionCollector struct {
	definitions map[string]Definition
}

func newDefinitionCollector() *definitionCollector {
	return &definitionCollector{
		definitions: make(map[string]Definition),
	}
}

func (c *definitionCollector) add(def Definition, provider, providerVersion string) {
	if def.Kind == "" {
		return
	}
	key := definitionID(def)
	existing, ok := c.definitions[key]
	if ok {
		def = existing
	}
	if provider != "" {
		def.Provider = provider
	}
	if canonical := version.Canonical(providerVersion); canonical != "" {
		def.ProviderVersions = append(def.ProviderVersions, canonical)
	}
	c.definitions[key] = def
}

func (c *definitionCollector) list() []Definition {
	var defs []Definition
	for _, def := range c.definitions {
		def.ProviderVersions = version.NormalizeList(def.ProviderVersions)
		defs = append(defs, def)
	}
	return defs
}

func decodeCRD(r io.Reader, provider, providerVersion string, collector *definitionCollector) error {
	decoder := yaml.NewDecoder(r)
	for {
		var crd customResourceDefinition
		if err := decoder.Decode(&crd); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("decode CRD document: %w", err)
		}
		defs := definitionsFromCRD(crd)
		for _, def := range defs {
			collector.add(def, provider, providerVersion)
		}
	}
	return nil
}

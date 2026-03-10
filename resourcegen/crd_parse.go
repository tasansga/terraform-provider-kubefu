package resourcegen

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// ParseCRDResources parses one or more CRD YAML documents and returns ResourceEntry
// values with Terraform schemas derived from the OpenAPI v3 schema.
func ParseCRDResources(data []byte) ([]ResourceEntry, error) {
	if len(bytes.TrimSpace(data)) == 0 {
		return nil, fmt.Errorf("empty CRD input")
	}
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var entries []ResourceEntry
	for {
		var crd customResourceDefinition
		if err := decoder.Decode(&crd); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("decode CRD document: %w", err)
		}
		defs := definitionsFromCRD(crd)
		for _, def := range defs {
			entries = append(entries, ResourceEntry{
				Group:                 def.Group,
				Version:               def.Version,
				Kind:                  def.Kind,
				DefinitionName:        def.DefinitionName,
				DefinitionDescription: def.Description,
				Schema:                def.Schema,
			})
		}
	}
	return entries, nil
}

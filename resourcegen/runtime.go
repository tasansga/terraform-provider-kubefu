package resourcegen

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// LoadDataSources collects definitions from the provided schema root directory
// and returns Terraform data sources keyed by their provider-aware names.
func LoadDataSources(schemaRoot string) (map[string]*schema.Resource, error) {
	definitions, err := CollectDefinitions(schemaRoot)
	if err != nil {
		return nil, err
	}
	return buildDataSources(definitions), nil
}

func buildDataSources(definitions map[string][]Definition) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource)
	for provider, defs := range definitions {
		providerName := providerSegment(provider)
		for _, def := range defs {
			name := fmt.Sprintf("kubefu_%s_%s_%s", providerName, toSnakeCase(def.Kind), versionSegment(def.Version))
			result[name] = def.AsSchemaResource()
		}
	}
	return result
}

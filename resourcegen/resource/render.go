package resource

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func renderSchema(schemaMap map[string]*schema.Schema) string {
	var keys []string
	for name := range schemaMap {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, name := range keys {
		entry := schemaMap[name]
		b.WriteString(renderSchemaEntry(name, entry))
	}
	return b.String()
}

func renderSchemaEntry(name string, entry *schema.Schema) string {
	var b strings.Builder
	fmt.Fprintf(&b, `			"%s": {
				Type:        %s,
				Description: %q,
				Optional:    %t,
				Required:    %t,
				Computed:    %t,
`, name, schemaValueTypeName(entry.Type), entry.Description, entry.Optional, entry.Required, entry.Computed)
	if entry.Elem != nil {
		switch elem := entry.Elem.(type) {
		case *schema.Resource:
			fmt.Fprintf(&b, "				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},\n")
		case *schema.Schema:
			fmt.Fprintf(&b, "				Elem: &schema.Schema{Type: %s},\n", schemaValueTypeName(elem.Type))
		default:
			fmt.Fprintf(&b, "				Elem: %#v,\n", elem)
		}
	}
	fmt.Fprintf(&b, "			},\n")
	return b.String()
}

func schemaValueTypeName(valueType schema.ValueType) string {
	switch valueType {
	case schema.TypeBool:
		return "schema.TypeBool"
	case schema.TypeInt:
		return "schema.TypeInt"
	case schema.TypeFloat:
		return "schema.TypeFloat"
	case schema.TypeString:
		return "schema.TypeString"
	case schema.TypeList:
		return "schema.TypeList"
	case schema.TypeSet:
		return "schema.TypeSet"
	case schema.TypeMap:
		return "schema.TypeMap"
	default:
		return "schema.TypeString"
	}
}

func escapeString(input string) string {
	input = strings.ReplaceAll(input, `"`, `\"`)
	input = strings.ReplaceAll(input, "\n", `\n`)
	input = strings.ReplaceAll(input, "\r", `\r`)
	return input
}

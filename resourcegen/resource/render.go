package resource

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func renderSchema(schemaMap map[string]*schema.Schema) string {
	return renderSchemaWithIndent(schemaMap, "\t\t\t")
}

func renderSchemaWithIndent(schemaMap map[string]*schema.Schema, indent string) string {
	var keys []string
	for name := range schemaMap {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, name := range keys {
		entry := schemaMap[name]
		b.WriteString(renderSchemaEntry(name, entry, indent))
	}
	return b.String()
}

func renderSchemaEntry(name string, entry *schema.Schema, indent string) string {
	var b strings.Builder
	fieldIndent := indent + "\t"
	fmt.Fprintf(&b, `%s"%s": {
%sType:        %s,
%sDescription: %q,
%sOptional:    %t,
%sRequired:    %t,
%sComputed:    %t,
`, indent, name, fieldIndent, schemaValueTypeName(entry.Type), fieldIndent, entry.Description, fieldIndent, entry.Optional, fieldIndent, entry.Required, fieldIndent, entry.Computed)
	if entry.ConfigMode != 0 {
		fmt.Fprintf(&b, "%sConfigMode:  %s,\n", fieldIndent, schemaConfigModeName(entry.ConfigMode))
	}
	if entry.MinItems > 0 {
		fmt.Fprintf(&b, "%sMinItems:    %d,\n", fieldIndent, entry.MinItems)
	}
	if entry.MaxItems > 0 {
		fmt.Fprintf(&b, "%sMaxItems:    %d,\n", fieldIndent, entry.MaxItems)
	}
	if entry.Elem != nil {
		switch elem := entry.Elem.(type) {
		case *schema.Resource:
			fmt.Fprintf(&b, "%sElem: &schema.Resource{Schema: map[string]*schema.Schema{\n", fieldIndent)
			b.WriteString(renderSchemaWithIndent(elem.Schema, fieldIndent+"\t"))
			fmt.Fprintf(&b, "%s}},\n", fieldIndent)
		case *schema.Schema:
			fmt.Fprintf(&b, "%sElem: &schema.Schema{Type: %s},\n", fieldIndent, schemaValueTypeName(elem.Type))
		default:
			fmt.Fprintf(&b, "%sElem: %#v,\n", fieldIndent, elem)
		}
	}
	fmt.Fprintf(&b, "%s},\n", indent)
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

func schemaConfigModeName(mode schema.SchemaConfigMode) string {
	switch mode {
	case schema.SchemaConfigModeAttr:
		return "schema.SchemaConfigModeAttr"
	case schema.SchemaConfigModeBlock:
		return "schema.SchemaConfigModeBlock"
	default:
		return "schema.SchemaConfigModeAuto"
	}
}

func escapeString(input string) string {
	input = strings.ReplaceAll(input, `"`, `\"`)
	input = strings.ReplaceAll(input, "\n", `\n`)
	input = strings.ReplaceAll(input, "\r", `\r`)
	return input
}

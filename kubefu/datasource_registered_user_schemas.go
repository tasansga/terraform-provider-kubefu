package kubefu

import (
	"context"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const registeredUserSchemasDataSourceName = "kubefu_registered_user_schemas"

func dataSourceRegisteredUserSchemas(names []string) *schema.Resource {
	registered := append([]string(nil), names...)
	sort.Strings(registered)

	return &schema.Resource{
		ReadContext: func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
			items := make([]interface{}, 0, len(registered))
			for _, name := range registered {
				items = append(items, name)
			}
			if err := d.Set("names", items); err != nil {
				return diag.FromErr(err)
			}
			d.SetId("registered-user-schemas")
			return nil
		},
		Description: "Lists Terraform data source names registered from user-supplied schemas.",
		Schema: map[string]*schema.Schema{
			"names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Registered data source names from user-supplied schemas.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func userSchemasRegisteredDiagnostic(names []string) *diag.Diagnostic {
	if len(names) == 0 {
		return nil
	}
	registered := append([]string(nil), names...)
	sort.Strings(registered)

	return &diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "User schemas registered",
		Detail:   "Registered user data sources:\n- " + strings.Join(registered, "\n- "),
	}
}

package resource

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource describes enough of a Terraform resource to emit a Go source file
// similar to ./example_resource_group.go.
type Resource struct {
	// PackageName will be used as the package statement in the generated file.
	PackageName string
	// FuncName is the name of the exported factory function (e.g. "resourceConfigMap").
	FuncName string
	// Description describes the resource and is written into the schema.Resource metadata.
	Description string
	// Schema carries the Terraform attribute schema that will be rendered verbatim.
	Schema map[string]*schema.Schema
}

// WriteFile emits the resource as a Go source file to the provided path.
func (r *Resource) WriteFile(path string) error {
	data, err := r.Render()
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", path, err)
	}
	return nil
}

// Render builds the resource source code and returns it as a string.
func (r *Resource) Render() (string, error) {
	var buf bytes.Buffer
	if err := r.writeHeader(&buf); err != nil {
		return "", err
	}
	if err := r.writeResource(&buf); err != nil {
		return "", err
	}
	if err := r.writeStubs(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *Resource) writeHeader(buf *bytes.Buffer) error {
	_, err := buf.WriteString(fmt.Sprintf(`package %s

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

`, r.PackageName))
	return err
}

func (r *Resource) writeResource(buf *bytes.Buffer) error {
	_, err := buf.WriteString(fmt.Sprintf(`func %s() *schema.Resource {
	return &schema.Resource{
		CreateContext: %sCreate,
		ReadContext:   %sRead,
		UpdateContext: %sUpdate,
		DeleteContext: %sDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
				_ = d.Set("id", d.Id())
				return schema.ImportStatePassthroughContext(ctx, d, m)
			},
		},
		Description: "%s",
		Schema: map[string]*schema.Schema{
%s		},
	}
}


`, r.FuncName, r.FuncName, r.FuncName, r.FuncName, r.FuncName, escapeString(r.Description), renderSchema(r.Schema)))
	return err
}

func (r *Resource) writeStubs(buf *bytes.Buffer) error {
	stubs := fmt.Sprintf(`
func %sCreate(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return diag.Diagnostics{}
}

func %sRead(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return diag.Diagnostics{}
}

func %sUpdate(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return diag.Diagnostics{}
}

func %sDelete(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	return diag.Diagnostics{}
}
`, r.FuncName, r.FuncName, r.FuncName, r.FuncName)
	_, err := buf.WriteString(stubs)
	return err
}

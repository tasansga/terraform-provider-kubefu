package resource

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSource describes enough of a Terraform data source to emit a Go source file
// similar to ./example_datasource_group.go.
type DataSource struct {
	PackageName         string
	FuncName            string
	Description         string
	Schema              map[string]*schema.Schema
	ProviderName        string
	CompatibleVersions  []string
	APIVersion          string
	Kind                string
	ID                  string
	ManifestKeys        []string
	ManifestObjectPaths []string
}

// WriteFile emits the data source as a Go source file to the provided path.
func (d *DataSource) WriteFile(path string) error {
	data, err := d.Render()
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", path, err)
	}
	return nil
}

// Render builds the data source source code and returns it as a string.
func (d *DataSource) Render() (string, error) {
	var buf bytes.Buffer
	if err := d.writeHeader(&buf); err != nil {
		return "", err
	}
	if err := d.writeDataSource(&buf); err != nil {
		return "", err
	}
	if err := d.writeStubs(&buf); err != nil {
		return "", err
	}
	if err := d.writeCompatibility(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (d *DataSource) writeHeader(buf *bytes.Buffer) error {
	_, err := buf.WriteString(fmt.Sprintf(`package %s

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

`, d.PackageName))
	return err
}

func (d *DataSource) writeDataSource(buf *bytes.Buffer) error {
	_, err := buf.WriteString(fmt.Sprintf(`func %s() *schema.Resource {
	return &schema.Resource{
		ReadContext: %sRead,
		Description: "%s",
		Schema: map[string]*schema.Schema{
%s		},
	}
}


`, d.FuncName, d.FuncName, escapeString(d.Description), renderSchema(d.Schema)))
	return err
}

func (d *DataSource) writeStubs(buf *bytes.Buffer) error {
	stubs := fmt.Sprintf(`
func %sRead(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, %q, %q, %q); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{%s}, []string{%s}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
`, d.FuncName, d.APIVersion, d.Kind, d.ID, renderStringSlice(d.ManifestKeys), renderStringSlice(d.ManifestObjectPaths))
	_, err := buf.WriteString(stubs)
	return err
}

func (d *DataSource) writeCompatibility(buf *bytes.Buffer) error {
	fmt.Fprintf(buf, "var %sCompatibleVersions = []string{\n", d.FuncName)
	for _, version := range d.CompatibleVersions {
		fmt.Fprintf(buf, "\t%q,\n", version)
	}
	fmt.Fprintf(buf, "}\n\n")
	return nil
}

func renderStringSlice(values []string) string {
	if len(values) == 0 {
		return ""
	}
	var b bytes.Buffer
	for i, v := range values {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%q", v)
	}
	return b.String()
}

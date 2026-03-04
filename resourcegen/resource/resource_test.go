package resource

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceWriteFile(t *testing.T) {
	dir := t.TempDir()
	res := &Resource{
		PackageName: "genpkg",
		FuncName:    "resourceFoo",
		Description: "Foo resource",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
	path := filepath.Join(dir, "resource_foo.go")
	if err := res.WriteFile(path); err != nil {
		t.Fatalf("WriteFile() error: %v", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read generated file: %v", err)
	}
	if !strings.Contains(string(content), "func resourceFoo() *schema.Resource") {
		t.Fatalf("generated file missing func signature: %s", content)
	}
}

func TestDataSourceWriteFile(t *testing.T) {
	dir := t.TempDir()
	ds := &DataSource{
		PackageName: "genpkg",
		FuncName:    "dataSourceBar",
		Description: "Bar data source",
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
	path := filepath.Join(dir, "datasource_bar.go")
	if err := ds.WriteFile(path); err != nil {
		t.Fatalf("WriteFile() error: %v", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read generated file: %v", err)
	}
	if !strings.Contains(string(content), "func dataSourceBar() *schema.Resource") {
		t.Fatalf("generated file missing func signature: %s", content)
	}
}

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1Read,
		Description: "Generated data source for io.k8s.api.apps.v1.SecretArgs",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object.",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"generator_args": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"behavior": {
						Type:        schema.TypeString,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"kv_pair_sources": {
						Type:        schema.TypeList,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"env": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"envs": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"files": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"literals": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
						}},
					},
					"name": {
						Type:        schema.TypeString,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"namespace": {
						Type:        schema.TypeString,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents.",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kubefu_manifest_json": {
				Type:        schema.TypeString,
				Description: "Rendered manifest (canonical JSON) for this data source.",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kubefu_manifest_yaml": {
				Type:        schema.TypeString,
				Description: "Rendered manifest (canonical YAML) for this data source.",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "kustomize.config.k8s.io/v1beta1", "SecretArgs", "kustomize.config.k8s.io/v1beta1/SecretArgs"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"generator_args", "type"}, []string{"generator_args", "generator_args.kv_pair_sources"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1CompatibleVersions = []string{
	"v0.3.1",
	"v0.3.2",
	"v0.3.3",
	"v0.3.4",
	"v0.4.0",
	"v0.4.1",
	"v0.4.2",
	"v0.5.0",
	"v0.6.0",
	"v0.6.1",
	"v0.7.0",
	"v0.7.1",
	"v0.8.0",
	"v0.8.1",
	"v0.9.0",
	"v0.9.1",
	"v0.9.2",
	"v0.9.3",
	"v0.9.4",
	"v0.9.5",
	"v0.10.0",
	"v0.10.1",
	"v0.10.2",
	"v0.10.3",
	"v0.10.4",
	"v0.10.5",
	"v0.10.6",
	"v0.10.7",
	"v0.10.8",
	"v0.10.9",
	"v0.10.10",
	"v0.10.11",
	"v0.10.12",
	"v0.10.13",
	"v0.10.14",
	"v0.10.15",
	"v0.10.16",
	"v0.10.17",
	"v0.10.18",
	"v0.10.19",
	"v0.10.20",
	"v0.10.21",
	"v0.11.0",
	"v0.11.1",
	"v0.12.0",
	"v0.13.0",
	"v0.13.1",
	"v0.13.2",
	"v0.13.3",
	"v0.13.4",
	"v0.13.5",
	"v0.13.6",
	"v0.13.7",
	"v0.13.8",
	"v0.13.9",
	"v0.13.10",
	"v0.14.0",
	"v0.14.1",
	"v0.14.2",
	"v0.14.3",
	"v0.15.0",
	"v0.16.0",
	"v0.17.0",
	"v0.17.1",
	"v0.17.2",
	"v0.18.0",
	"v0.18.1",
	"v0.19.0",
	"v0.20.0",
	"v0.20.1",
	"v0.21.0",
	"v0.21.1",
}

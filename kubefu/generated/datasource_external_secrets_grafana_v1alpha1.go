package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1Read,
		Description: "Grafana represents a generator for Grafana service account tokens.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object.\nServers should convert recognized schemas to the latest internal value, and\nmay reject unrecognized values.\nMore info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents.\nServers may infer this from the endpoint the client submits requests to.\nCannot be updated.\nIn CamelCase.\nMore info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
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
			"metadata": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"spec": {
				Type:        schema.TypeList,
				Description: "GrafanaSpec controls the behavior of the grafana generator.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"auth": {
						Type:        schema.TypeList,
						Description: "Auth is the authentication configuration to authenticate\nagainst the Grafana instance.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"basic": {
								Type:        schema.TypeList,
								Description: "Basic auth credentials used to authenticate against the Grafana instance.\nNote: you need a token which has elevated permissions to create service accounts.\nSee here for the documentation on basic roles offered by Grafana:\nhttps://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/rbac-fixed-basic-role-definitions/",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"password": {
										Type:        schema.TypeList,
										Description: "A basic auth password used to authenticate against the Grafana instance.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "The key where the token is found.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "The name of the Secret resource being referred to.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"username": {
										Type:        schema.TypeString,
										Description: "A basic auth username used to authenticate against the Grafana instance.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"token": {
								Type:        schema.TypeList,
								Description: "A service account token used to authenticate against the Grafana instance.\nNote: you need a token which has elevated permissions to create service accounts.\nSee here for the documentation on basic roles offered by Grafana:\nhttps://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/rbac-fixed-basic-role-definitions/",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Description: "The key where the token is found.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "The name of the Secret resource being referred to.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"service_account": {
						Type:        schema.TypeList,
						Description: "ServiceAccount is the configuration for the service account that\nis supposed to be generated by the generator.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the service account that will be created by ESO.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"role": {
								Type:        schema.TypeString,
								Description: "Role is the role of the service account.\nSee here for the documentation on basic roles offered by Grafana:\nhttps://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/rbac-fixed-basic-role-definitions/",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL is the URL of the Grafana instance.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "Grafana", "generators.external-secrets.io/v1alpha1/Grafana"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.auth", "spec.auth.basic", "spec.auth.basic.password", "spec.auth.token", "spec.service_account"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1CompatibleVersions = []string{
	"v0.14.0",
	"v0.14.1",
	"v0.14.2",
	"v0.14.3",
	"v0.14.4",
	"v0.15.0",
	"v0.15.1",
	"v0.16.0",
	"v0.16.1",
	"v0.16.2",
	"v0.17.0",
	"v0.18.0",
	"v0.18.1",
	"v0.18.2",
	"v0.19.0",
	"v0.19.1",
	"v0.19.2",
	"v0.20.1",
	"v0.20.2",
	"v0.20.3",
	"v0.20.4",
	"v1.0.0",
	"v1.1.0",
	"v1.1.1",
	"v1.2.0",
	"v1.2.1",
	"v1.3.1",
	"v1.3.2",
	"v2.0.0",
	"v2.0.1",
	"v2.1.0",
	"v2.2.0",
}

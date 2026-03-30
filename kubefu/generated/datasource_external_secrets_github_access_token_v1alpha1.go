package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1Read,
		Description: "GithubAccessToken generates ghs_ accessToken",
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
				Description: "GithubAccessTokenSpec defines the desired state to generate a GitHub access token.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"app_id": {
						Type:        schema.TypeString,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"auth": {
						Type:        schema.TypeList,
						Description: "Auth configures how ESO authenticates with a Github instance.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"privat_key": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"secret_ref": {
										Type:        schema.TypeList,
										Description: "A reference to a specific 'key' within a Secret resource,\nIn some instances, `key` is a required field.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be\ndefaulted, in others it may be required.",
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
											"namespace": {
												Type:        schema.TypeString,
												Description: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults\nto the namespace of the referent.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
								}},
							},
							"private_key": {
								Type:        schema.TypeList,
								Description: "GithubSecretRef references a secret containing GitHub credentials.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"secret_ref": {
										Type:        schema.TypeList,
										Description: "A reference to a specific 'key' within a Secret resource,\nIn some instances, `key` is a required field.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be\ndefaulted, in others it may be required.",
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
											"namespace": {
												Type:        schema.TypeString,
												Description: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults\nto the namespace of the referent.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
								}},
							},
						}},
					},
					"install_id": {
						Type:        schema.TypeString,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"permissions": {
						Type:        schema.TypeMap,
						Description: "Map of permissions the token will have. If omitted, defaults to all permissions the GitHub App has.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"repositories": {
						Type:        schema.TypeList,
						Description: "List of repositories the token will have access to. If omitted, defaults to all repositories the GitHub App\nis installed to.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL configures the Github instance URL. Defaults to https://github.com/.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "GithubAccessToken", "generators.external-secrets.io/v1alpha1/GithubAccessToken"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.auth", "spec.auth.privat_key", "spec.auth.privat_key.secret_ref", "spec.auth.private_key", "spec.auth.private_key.secret_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1CompatibleVersions = []string{
	"v0.8.15",
	"v0.8.16",
	"v0.8.17",
	"v0.9.15",
	"v0.9.16",
	"v0.9.17",
	"v0.9.18",
	"v0.9.19",
	"v0.9.20",
	"v0.10.0",
	"v0.10.1",
	"v0.10.2",
	"v0.10.3",
	"v0.10.4",
	"v0.10.5",
	"v0.10.6",
	"v0.10.7",
	"v0.11.0",
	"v0.12.0",
	"v0.12.1",
	"v0.13.0",
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
}

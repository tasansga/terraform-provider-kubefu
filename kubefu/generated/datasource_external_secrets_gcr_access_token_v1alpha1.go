package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1Read,
		Description: "GCRAccessToken generates an GCP access token\nthat can be used to authenticate with GCR.",
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
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"auth": {
						Type:        schema.TypeList,
						Description: "Auth defines the means for authenticating with GCP",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"secret_ref": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"secret_access_key_secret_ref": {
										Type:        schema.TypeList,
										Description: "The SecretAccessKey is used for authentication",
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
							"workload_identity": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"cluster_location": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"cluster_name": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"cluster_project_id": {
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"service_account_ref": {
										Type:        schema.TypeList,
										Description: "A reference to a ServiceAccount resource.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"audiences": {
												Type:        schema.TypeList,
												Description: "Audience specifies the `aud` claim for the service account token\nIf the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity\nthen this audiences will be appended to the list",
												Optional:    true,
												Required:    false,
												Computed:    true,
												Elem: &schema.Schema{Type: schema.TypeString},
											},
											"name": {
												Type:        schema.TypeString,
												Description: "The name of the ServiceAccount resource being referred to.",
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
					"project_id": {
						Type:        schema.TypeString,
						Description: "ProjectID defines which project to use to authenticate with",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "GCRAccessToken", "generators.external-secrets.io/v1alpha1/GCRAccessToken"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.auth", "spec.auth.secret_ref", "spec.auth.secret_ref.secret_access_key_secret_ref", "spec.auth.workload_identity", "spec.auth.workload_identity.service_account_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1CompatibleVersions = []string{
	"v0.7.0",
	"v0.7.1",
	"v0.7.2",
	"v0.7.3",
	"v0.8.0",
	"v0.8.1",
	"v0.8.2",
	"v0.8.3",
	"v0.8.4",
	"v0.8.5",
	"v0.8.6",
	"v0.8.7",
	"v0.8.8",
	"v0.8.9",
	"v0.8.10",
	"v0.8.11",
	"v0.8.12",
	"v0.8.13",
	"v0.8.14",
	"v0.8.15",
	"v0.8.16",
	"v0.8.17",
	"v0.9.0",
	"v0.9.1",
	"v0.9.2",
	"v0.9.3",
	"v0.9.4",
	"v0.9.5",
	"v0.9.6",
	"v0.9.7",
	"v0.9.8",
	"v0.9.9",
	"v0.9.10",
	"v0.9.11",
	"v0.9.12",
	"v0.9.13",
	"v0.9.14",
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

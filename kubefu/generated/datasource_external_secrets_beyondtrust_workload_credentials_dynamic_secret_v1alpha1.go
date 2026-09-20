package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoBeyondtrustWorkloadCredentialsDynamicSecretV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoBeyondtrustWorkloadCredentialsDynamicSecretV1Alpha1Read,
		Description: "BeyondtrustWorkloadCredentialsDynamicSecret represents a generator that requests dynamic credentials from BeyondTrust Workload Credentials.\nThis generator calls the BeyondTrust Workload Credentials API to generate fresh, temporary credentials\n(such as AWS STS credentials) each time an ExternalSecret is refreshed.\nDynamic secret definitions must be created in BeyondTrust Workload Credentials before they can be referenced.\nFor complete documentation, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api",
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
				Description: "BeyondtrustWorkloadCredentialsDynamicSecretSpec defines the desired spec for BeyondtrustWorkloadCredentials dynamic generator.\nThis generator enables obtaining temporary, short-lived credentials from BeyondTrust Workload Credentials.\nFor more information, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"controller": {
						Type:        schema.TypeString,
						Description: "Controller selects the controller that should handle this generator.\nLeave empty to use the default controller.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"provider_": {
						Type:        schema.TypeList,
						Description: "Provider contains the BeyondtrustWorkloadCredentials provider configuration including authentication,\nserver connection details, and the folder path to the dynamic secret definition.\nThe folderPath should point to a dynamic secret definition that has been created in\nBeyondTrust Workload Credentials (e.g., \"production/aws-temp\").\nFor setup details, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"auth": {
								Type:        schema.TypeList,
								Description: "Auth configures how the Operator authenticates with the BeyondTrust Workload Credentials API.\nCurrently supports API key authentication via Kubernetes secret reference.\nFor authentication setup, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api#authentication",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"apikey": {
										Type:        schema.TypeList,
										Description: "APIKey configures API token authentication for BeyondTrust Workload Credentials.\nThe token is retrieved from a Kubernetes secret and used as a Bearer token for API requests.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"token": {
												Type:        schema.TypeList,
												Description: "Token references the Kubernetes secret containing the BeyondTrust Workload Credentials API token.\nThe secret should contain the API key used to authenticate with BeyondTrust Workload Credentials.\nCreate an API token in your BeyondTrust Workload Credentials console and store it in a Kubernetes secret.\nFor details on creating API tokens, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api#authentication",
												Optional:    true,
												Required:    false,
												Computed:    true,
												MaxItems:    1,
												Elem: &schema.Resource{Schema: map[string]*schema.Schema{
													"key": {
														Type:        schema.TypeString,
														Description: "A key in the referenced Secret.\nSome instances of this field may be defaulted, in others it may be required.",
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
														Description: "The namespace of the Secret resource being referred to.\nIgnored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
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
							"ca_bundle": {
								Type:        schema.TypeString,
								Description: "CABundle is a base64-encoded CA certificate used to validate the BeyondTrust Workload Credentials API TLS certificate.\nUse this when your BeyondTrust instance uses a self-signed certificate or internal CA.\nIf not set, the system's trusted root certificates are used.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"ca_provider": {
								Type:        schema.TypeList,
								Description: "CAProvider points to a Secret or ConfigMap containing a PEM-encoded CA certificate.\nThis is used to validate the BeyondTrust Workload Credentials API TLS certificate.\nUse this as an alternative to CABundle when you want to reference an existing Kubernetes resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Description: "The key where the CA certificate can be found in the Secret or ConfigMap.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "The name of the object located at the provider type.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"namespace": {
										Type:        schema.TypeString,
										Description: "The namespace the Provider type is in.\nCan only be defined when used in a ClusterSecretStore.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"folder_path": {
								Type:        schema.TypeString,
								Description: "FolderPath specifies the default folder path for secret retrieval.\nSecrets will be fetched from this folder unless overridden in the ExternalSecret spec.\nExample: \"production/database\" or \"dev/api-keys\"\nLeave empty to retrieve secrets from the root folder.\nFor folder organization, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api#folders",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"server": {
								Type:        schema.TypeList,
								Description: "Server configures the BeyondTrust Workload Credentials server connection details.\nIncludes the API URL and Site ID for your BeyondTrust instance.\nFor API reference, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"api_url": {
										Type:        schema.TypeString,
										Description: "APIURL is the base URL of your BeyondTrust Workload Credentials API server.\nThis should be the full URL to your BeyondTrust instance.\nExample: https://api.beyondtrust.io/siie\nFor more information, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api#base-url",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"site_id": {
										Type:        schema.TypeString,
										Description: "SiteID is your BeyondTrust Workload Credentials site identifier (UUID format).\nThis identifier is unique to your BeyondTrust Workload Credentials instance.\nYou can find your Site ID in the BeyondTrust Workload Credentials admin console.\nExample: a1b2c3d4-e5f6-4890-abcd-ef1234567890\nFor more information, see: https://docs.beyondtrust.com/bt-docs/docs/secrets-api",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"retry_settings": {
						Type:        schema.TypeList,
						Description: "RetrySettings configures exponential backoff for failed API requests.\nIf not specified, uses the default retry settings.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"max_retries": {
								Type:        schema.TypeInt,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"retry_interval": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoBeyondtrustWorkloadCredentialsDynamicSecretV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "BeyondtrustWorkloadCredentialsDynamicSecret", "generators.external-secrets.io/v1alpha1/BeyondtrustWorkloadCredentialsDynamicSecret"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.provider_", "spec.provider_.auth", "spec.provider_.auth.apikey", "spec.provider_.auth.apikey.token", "spec.provider_.ca_provider", "spec.provider_.server", "spec.retry_settings"}, []string{"spec", "spec.provider_", "spec.provider_.auth", "spec.provider_.auth.apikey", "spec.provider_.auth.apikey.token", "spec.provider_.ca_provider", "spec.provider_.server", "spec.retry_settings"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoBeyondtrustWorkloadCredentialsDynamicSecretV1Alpha1CompatibleVersions = []string{
	"v2.7.0",
	"v2.8.0",
	"v2.9.0",
	"v2.10.0",
}

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1Read,
		Description: "STSSessionToken uses the GetSessionToken API to retrieve an authorization token.\nThe authorization token is valid for 12 hours.\nThe authorizationToken returned is a base64 encoded string that can be decoded.\nFor more information, see GetSessionToken (https://docs.aws.amazon.com/STS/latest/APIReference/API_GetSessionToken.html).",
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
						Description: "Auth defines how to authenticate with AWS",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"jwt": {
								Type:        schema.TypeList,
								Description: "Authenticate against AWS using service account tokens.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
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
							"secret_ref": {
								Type:        schema.TypeList,
								Description: "AWSAuthSecretRef holds secret references for AWS credentials\nboth AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"access_key_id_secret_ref": {
										Type:        schema.TypeList,
										Description: "The AccessKeyID is used for authentication",
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
									"session_token_secret_ref": {
										Type:        schema.TypeList,
										Description: "The SessionToken used for authentication\nThis must be defined if AccessKeyID and SecretAccessKey are temporary credentials\nsee: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
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
					"region": {
						Type:        schema.TypeString,
						Description: "Region specifies the region to operate in.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"request_parameters": {
						Type:        schema.TypeList,
						Description: "RequestParameters contains parameters that can be passed to the STS service.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"serial_number": {
								Type:        schema.TypeString,
								Description: "SerialNumber is the identification number of the MFA device that is associated with the IAM user who is making\nthe GetSessionToken call.\nPossible values: hardware device (such as GAHT12345678) or an Amazon Resource Name (ARN) for a virtual device\n(such as arn:aws:iam::123456789012:mfa/user)",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"session_duration": {
								Type:        schema.TypeInt,
								Description: "SessionDuration The duration, in seconds, that the credentials should remain valid. Acceptable durations for\nIAM user sessions range from 900 seconds (15 minutes) to 129,600 seconds (36 hours), with 43,200 seconds\n(12 hours) as the default.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"token_code": {
								Type:        schema.TypeString,
								Description: "TokenCode is the value provided by the MFA device, if MFA is required.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"role": {
						Type:        schema.TypeString,
						Description: "You can assume a role before making calls to the\ndesired AWS service.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "STSSessionToken", "generators.external-secrets.io/v1alpha1/STSSessionToken"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec"}, []string{"spec", "spec.auth", "spec.auth.jwt", "spec.auth.jwt.service_account_ref", "spec.auth.secret_ref", "spec.auth.secret_ref.access_key_id_secret_ref", "spec.auth.secret_ref.secret_access_key_secret_ref", "spec.auth.secret_ref.session_token_secret_ref", "spec.request_parameters"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1CompatibleVersions = []string{
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

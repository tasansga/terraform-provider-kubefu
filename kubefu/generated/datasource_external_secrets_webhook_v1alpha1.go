package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1Read,
		Description: "Webhook connects to a third party API server to handle the secrets generation\nconfiguration parameters in spec.\nYou can specify the server, the token, and additional body parameters.\nSee documentation for the full API specification for requests and responses.",
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
				Description: "WebhookSpec controls the behavior of the external generator. Any body parameters should be passed to the server through the parameters field.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"body": {
						Type:        schema.TypeString,
						Description: "Body",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"ca_bundle": {
						Type:        schema.TypeString,
						Description: "PEM encoded CA bundle used to validate webhook server certificate. Only used\nif the Server URL is using HTTPS protocol. This parameter is ignored for\nplain HTTP protocol connection. If not set the system root certificates\nare used to validate the TLS connection.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"ca_provider": {
						Type:        schema.TypeList,
						Description: "The provider for the CA bundle to use to validate webhook server certificate.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"key": {
								Type:        schema.TypeString,
								Description: "The key the value inside of the provider type to use, only used with \"Secret\" type",
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
								Description: "The namespace the Provider type is in.",
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
					"headers": {
						Type:        schema.TypeMap,
						Description: "Headers",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"method": {
						Type:        schema.TypeString,
						Description: "Webhook Method",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"result": {
						Type:        schema.TypeList,
						Description: "Result formatting",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"json_path": {
								Type:        schema.TypeString,
								Description: "Json path of return value",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"secrets": {
						Type:        schema.TypeList,
						Description: "Secrets to fill in templates\nThese secrets will be passed to the templating function as key value pairs under the given name",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name of this secret in templates",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"secret_ref": {
								Type:        schema.TypeList,
								Description: "Secret ref to fill in credentials",
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
					"timeout": {
						Type:        schema.TypeString,
						Description: "Timeout",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"url": {
						Type:        schema.TypeString,
						Description: "Webhook url to call",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "Webhook", "generators.external-secrets.io/v1alpha1/Webhook"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.ca_provider", "spec.result", "spec.secrets.secret_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1CompatibleVersions = []string{
	"v0.8.13",
	"v0.8.14",
	"v0.8.15",
	"v0.8.16",
	"v0.8.17",
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

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3Read,
		Description: "Provider is the Schema for the providers API",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
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
				Description: "ProviderSpec defines the desired state of the Provider.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"address": {
						Type:        schema.TypeString,
						Description: "Address specifies the endpoint, in a generic sense, to where alerts are sent. What kind of endpoint depends on the specific Provider type being used. For the generic Provider, for example, this is an HTTP/S address. For other Provider types this could be a project ID or a namespace.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"cert_secret_ref": {
						Type:        schema.TypeList,
						Description: "CertSecretRef specifies the Secret containing a PEM-encoded CA certificate (in the `ca.crt` key). \n Note: Support for the `caFile` key has been deprecated.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"channel": {
						Type:        schema.TypeString,
						Description: "Channel specifies the destination channel where events should be posted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"commit_status_expr": {
						Type:        schema.TypeString,
						Description: "CommitStatusExpr is a CEL expression that evaluates to a string value\nthat can be used to generate a custom commit status message for use\nwith eligible Provider types (github, gitlab, gitea, bitbucketserver,\nbitbucket, azuredevops). Supported variables are: event, provider,\nand alert.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval at which to reconcile the Provider with its Secret references. Deprecated and not used in v1beta3.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"proxy": {
						Type:        schema.TypeString,
						Description: "Proxy the HTTP/S address of the proxy server.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"proxy_secret_ref": {
						Type:        schema.TypeList,
						Description: "ProxySecretRef specifies the Secret containing the proxy configuration\nfor this Provider. The Secret should contain an 'address' key with the\nHTTP/S address of the proxy server. Optional 'username' and 'password'\nkeys can be provided for proxy authentication.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"secret_ref": {
						Type:        schema.TypeList,
						Description: "SecretRef specifies the Secret containing the authentication credentials for this Provider.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"service_account_name": {
						Type:        schema.TypeString,
						Description: "ServiceAccountName is the name of the service account used to\nauthenticate with services from cloud providers. An error is thrown if a\nstatic credential is also defined inside the Secret referenced by the\nSecretRef.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"suspend": {
						Type:        schema.TypeBool,
						Description: "Suspend tells the controller to suspend subsequent events handling for this Provider.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"timeout": {
						Type:        schema.TypeString,
						Description: "Timeout for sending alerts to the Provider.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"type": {
						Type:        schema.TypeString,
						Description: "Type specifies which Provider implementation to use.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"username": {
						Type:        schema.TypeString,
						Description: "Username specifies the name under which events are posted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "notification.toolkit.fluxcd.io/v1beta3", "Provider", "notification.toolkit.fluxcd.io/v1beta3/Provider"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.cert_secret_ref", "spec.proxy_secret_ref", "spec.secret_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3CompatibleVersions = []string{
	"v2.2.0",
	"v2.2.1",
	"v2.2.2",
	"v2.2.3",
	"v2.3.0",
	"v2.4.0",
	"v2.5.0",
	"v2.5.1",
	"v2.6.0",
	"v2.6.1",
	"v2.6.2",
	"v2.6.3",
	"v2.6.4",
	"v2.7.0",
	"v2.7.1",
	"v2.7.2",
	"v2.7.3",
	"v2.7.4",
	"v2.7.5",
}

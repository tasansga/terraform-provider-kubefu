package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1Read,
		Description: "CloudsmithAccessToken generates Cloudsmith access token using OIDC authentication",
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
				Description: "CloudsmithAccessTokenSpec defines the configuration for generating a Cloudsmith access token using OIDC authentication.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"api_url": {
						Type:        schema.TypeString,
						Description: "APIURL configures the Cloudsmith API URL. Defaults to https://api.cloudsmith.io.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"org_slug": {
						Type:        schema.TypeString,
						Description: "OrgSlug is the organization slug in Cloudsmith",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"service_account_ref": {
						Type:        schema.TypeList,
						Description: "Name of the service account you are federating with",
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
								Description: "Namespace of the resource being referred to.\nIgnored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"service_slug": {
						Type:        schema.TypeString,
						Description: "ServiceSlug is the service slug in Cloudsmith for OIDC authentication",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "CloudsmithAccessToken", "generators.external-secrets.io/v1alpha1/CloudsmithAccessToken"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.service_account_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1CompatibleVersions = []string{
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

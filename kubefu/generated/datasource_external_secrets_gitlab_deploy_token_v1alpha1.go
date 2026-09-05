package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoGitlabDeployTokenV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoGitlabDeployTokenV1Alpha1Read,
		Description: "GitlabDeployToken generates a GitLab deploy token.",
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
				Description: "GitlabDeployTokenSpec defines the desired state to generate a GitLab deploy token.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"auth": {
						Type:        schema.TypeList,
						Description: "Auth configures how ESO authenticates with the GitLab API.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"token": {
								Type:        schema.TypeList,
								Description: "Token references a secret containing a GitLab access token (personal, group, or\nproject) with the api scope and at least the Maintainer role on the target.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"secret_ref": {
										Type:        schema.TypeList,
										Description: "SecretKeySelector is a reference to a specific 'key' within a Secret resource.\nIn some instances, `key` is a required field.",
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
					"expires_at": {
						Type:        schema.TypeString,
						Description: "ExpiresAt is an optional expiry for the deploy token. If omitted the token does\nnot expire on the GitLab side and is revoked only when the generator state is\ncleaned up (on regeneration or when the consuming ExternalSecret is deleted).",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"group_id": {
						Type:        schema.TypeString,
						Description: "GroupID is the numeric ID or unescaped path (e.g. parent/group) of the group to\ncreate the deploy token in. The generator URL-escapes paths before calling the\nGitLab API, so do not pre-encode. Mutually exclusive with projectID.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"name": {
						Type:        schema.TypeString,
						Description: "Name of the deploy token.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"project_id": {
						Type:        schema.TypeString,
						Description: "ProjectID is the numeric ID or unescaped path (e.g. group/project) of the\nproject to create the deploy token in. The generator URL-escapes paths before\ncalling the GitLab API, so do not pre-encode. Mutually exclusive with groupID.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"scopes": {
						Type:        schema.TypeList,
						Description: "Scopes granted to the deploy token. At least one scope is required.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL configures the GitLab instance URL. Defaults to https://gitlab.com.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"username": {
						Type:        schema.TypeString,
						Description: "Username is an optional username for the deploy token. GitLab defaults it to\ngitlab+deploy-token-{n} when omitted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoGitlabDeployTokenV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "GitlabDeployToken", "generators.external-secrets.io/v1alpha1/GitlabDeployToken"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec", "spec.auth", "spec.auth.token", "spec.auth.token.secret_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoGitlabDeployTokenV1Alpha1CompatibleVersions = []string{
	"v2.8.0",
	"v2.9.0",
	"v2.10.0",
}

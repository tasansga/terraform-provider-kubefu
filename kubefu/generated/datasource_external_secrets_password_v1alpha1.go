package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1Read,
		Description: "Password generates a random password based on the configuration parameters in spec. You can specify the length, characterset and other attributes.",
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
				Description: "PasswordSpec controls the behavior of the password generator.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allow_repeat": {
						Type:        schema.TypeBool,
						Description: "set AllowRepeat to true to allow repeating characters.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"digits": {
						Type:        schema.TypeInt,
						Description: "Digits specifies the number of digits in the generated password. If omitted it defaults to 25% of the length of the password",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"encoding": {
						Type:        schema.TypeString,
						Description: "Encoding specifies the encoding of the generated password.\nValid values are:\n- \"raw\" (default): no encoding\n- \"base64\": standard base64 encoding\n- \"base64url\": base64url encoding\n- \"base32\": base32 encoding\n- \"hex\": hexadecimal encoding",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"length": {
						Type:        schema.TypeInt,
						Description: "Length of the password to be generated. Defaults to 24",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"no_upper": {
						Type:        schema.TypeBool,
						Description: "Set NoUpper to disable uppercase characters",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"secret_keys": {
						Type:        schema.TypeList,
						Description: "SecretKeys defines the keys that will be populated with generated passwords.\nDefaults to \"password\" when not set.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"symbol_characters": {
						Type:        schema.TypeString,
						Description: "SymbolCharacters specifies the special characters that should be used in the generated password.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"symbols": {
						Type:        schema.TypeInt,
						Description: "Symbols specifies the number of symbol characters in the generated password. If omitted it defaults to 25% of the length of the password",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "generators.external-secrets.io/v1alpha1", "Password", "generators.external-secrets.io/v1alpha1/Password"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"spec"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1CompatibleVersions = []string{
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

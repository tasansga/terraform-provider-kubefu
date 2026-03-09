package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1Read,
		Description: "SelfSubjectReview contains the user information that the kube-apiserver has about the user making this request. When using impersonation, users will receive the user info of the user being impersonated.  If impersonation or request header authentication is used, any extra keys will have their case ignored and returned as lowercase.",
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
				Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeMap,
				Description: "Status is filled in by the server with the user attributes.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "authentication.k8s.io/v1beta1", "SelfSubjectReview", "authentication.k8s.io/v1beta1/SelfSubjectReview"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1CompatibleVersions = []string{
	"v1.27.0",
	"v1.28.0",
	"v1.29.0",
	"v1.30.0",
	"v1.31.0",
	"v1.32.0",
}

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1Read,
		Description: "ImageRepository is the Schema for the imagerepositories API",
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
				Type:        schema.TypeMap,
				Description: "ImageRepositorySpec defines the parameters for scanning an image repository, e.g., `fluxcd/flux`.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeMap,
				Description: "ImageRepositoryStatus defines the observed state of ImageRepository",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "image.toolkit.fluxcd.io/v1beta1", "ImageRepository", "image.toolkit.fluxcd.io/v1beta1/ImageRepository"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "spec", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1CompatibleVersions = []string{
	"v0.16.0",
	"v0.16.1",
	"v0.16.2",
	"v0.17.0",
	"v0.17.1",
	"v0.17.2",
	"v0.18.0",
	"v0.18.1",
	"v0.18.2",
	"v0.18.3",
	"v0.19.0",
	"v0.19.1",
	"v0.20.0",
	"v0.20.1",
	"v0.21.0",
	"v0.21.1",
	"v0.22.0",
	"v0.22.1",
	"v0.23.0",
	"v0.24.0",
	"v0.24.1",
	"v0.25.0",
	"v0.25.1",
	"v0.25.2",
	"v0.25.3",
	"v0.26.0",
	"v0.26.1",
	"v0.26.2",
	"v0.26.3",
	"v0.27.0",
	"v0.27.1",
	"v0.27.2",
	"v0.27.3",
	"v0.27.4",
	"v0.28.0",
	"v0.28.1",
	"v0.28.2",
	"v0.28.3",
	"v0.28.4",
	"v0.28.5",
	"v0.29.0",
	"v0.29.1",
	"v0.29.2",
	"v0.29.3",
	"v0.29.4",
	"v0.29.5",
	"v0.30.0",
	"v0.30.1",
	"v0.30.2",
	"v0.31.0",
	"v0.31.1",
	"v0.31.2",
	"v0.31.3",
	"v0.31.4",
	"v0.31.5",
	"v0.32.0",
	"v0.33.0",
	"v0.34.0",
	"v0.35.0",
	"v0.36.0",
	"v0.37.0",
	"v0.38.0",
	"v0.38.1",
	"v0.38.2",
	"v0.38.3",
	"v0.39.0",
	"v0.40.0",
	"v0.40.1",
	"v0.40.2",
	"v0.41.0",
	"v0.41.1",
	"v0.41.2",
	"v2.0.0",
	"v2.0.1",
	"v2.1.0",
	"v2.1.1",
	"v2.1.2",
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
}

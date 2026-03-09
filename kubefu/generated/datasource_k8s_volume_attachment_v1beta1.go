package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1Read,
		Description: "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.\n\nVolumeAttachment objects are non-namespaced.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
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
				Description: "Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"spec": {
				Type:        schema.TypeMap,
				Description: "Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"status": {
				Type:        schema.TypeMap,
				Description: "Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "storage.k8s.io/v1beta1", "VolumeAttachment", "storage.k8s.io/v1beta1/VolumeAttachment"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"metadata", "spec", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1CompatibleVersions = []string{
	"v1.10.0",
	"v1.11.0",
	"v1.12.0",
	"v1.13.0",
	"v1.14.0",
	"v1.15.0",
	"v1.16.0",
	"v1.17.0",
	"v1.18.0",
	"v1.19.0",
	"v1.20.0",
	"v1.21.0",
}

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sStorageK8sIoStorageClassV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sStorageK8sIoStorageClassV1Read,
		Description: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.\n\nStorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
		Schema: map[string]*schema.Schema{
			"allow_volume_expansion": {
				Type:        schema.TypeBool,
				Description: "AllowVolumeExpansion shows whether the storage class allow volume expand",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
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
				Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"mount_options": {
				Type:        schema.TypeList,
				Description: "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. [\"ro\", \"soft\"]. Not validated - mount of the PVs will simply fail if one is invalid.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"parameters": {
				Type:        schema.TypeMap,
				Description: "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"provisioner_": {
				Type:        schema.TypeString,
				Description: "Provisioner indicates the type of the provisioner.",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"reclaim_policy": {
				Type:        schema.TypeString,
				Description: "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"volume_binding_mode": {
				Type:        schema.TypeString,
				Description: "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is alpha-level and is only honored by servers that enable the VolumeScheduling feature.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sStorageK8sIoStorageClassV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "storage.k8s.io/v1", "StorageClass", "storage.k8s.io/v1/StorageClass"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"allow_volume_expansion", "metadata", "mount_options", "parameters", "provisioner_", "reclaim_policy", "volume_binding_mode"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sStorageK8sIoStorageClassV1CompatibleVersions = []string{
	"v1.7.0",
	"v1.8.0",
	"v1.9.0",
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
	"v1.22.0",
	"v1.23.0",
	"v1.24.0",
	"v1.25.0",
	"v1.26.0",
	"v1.27.0",
	"v1.28.0",
	"v1.29.0",
	"v1.30.0",
	"v1.31.0",
	"v1.32.0",
	"v1.33.0",
	"v1.34.0",
	"v1.35.0",
}

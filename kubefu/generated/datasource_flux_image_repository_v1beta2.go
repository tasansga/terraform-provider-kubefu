package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2Read,
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
				Type:        schema.TypeList,
				Description: "ImageRepositorySpec defines the parameters for scanning an image repository, e.g., `fluxcd/flux`.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"access_from": {
						Type:        schema.TypeList,
						Description: "AccessFrom defines an ACL for allowing cross-namespace references to the ImageRepository object based on the caller's namespace labels.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"namespace_selectors": {
								Type:        schema.TypeList,
								Description: "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"match_labels": {
										Type:        schema.TypeMap,
										Description: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"cert_secret_ref": {
						Type:        schema.TypeList,
						Description: "CertSecretRef can be given the name of a secret containing either or both of \n - a PEM-encoded client certificate (`certFile`) and private key (`keyFile`); - a PEM-encoded CA certificate (`caFile`) \n and whichever are supplied, will be used for connecting to the registry. The client cert and key are useful if you are authenticating with a certificate; the CA cert is useful if you are using a self-signed server certificate.",
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
					"exclusion_list": {
						Type:        schema.TypeList,
						Description: "ExclusionList is a list of regex strings used to exclude certain tags from being stored in the database.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"image": {
						Type:        schema.TypeString,
						Description: "Image is the name of the image repository",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval is the length of time to wait between scans of the image repository.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"provider_": {
						Type:        schema.TypeString,
						Description: "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'. When not specified, defaults to 'generic'.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"secret_ref": {
						Type:        schema.TypeList,
						Description: "SecretRef can be given the name of a secret containing credentials to use for the image registry. The secret should be created with `kubectl create secret docker-registry`, or the equivalent.",
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
						Description: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate the image pull if the service account has attached pull secrets.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"suspend": {
						Type:        schema.TypeBool,
						Description: "This flag tells the controller to suspend subsequent image scans. It does not apply to already started scans. Defaults to false.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"timeout": {
						Type:        schema.TypeString,
						Description: "Timeout for image scanning. Defaults to 'Interval' duration.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "ImageRepositoryStatus defines the observed state of ImageRepository",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"canonical_image_name": {
						Type:        schema.TypeString,
						Description: "CanonicalName is the name of the image repository with all the implied bits made explicit; e.g., `docker.io/library/alpine` rather than `alpine`.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"conditions": {
						Type:        schema.TypeList,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"last_transition_time": {
								Type:        schema.TypeString,
								Description: "lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"message": {
								Type:        schema.TypeString,
								Description: "message is a human readable message indicating details about the transition. This may be an empty string.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"observed_generation": {
								Type:        schema.TypeInt,
								Description: "observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"reason": {
								Type:        schema.TypeString,
								Description: "reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"status": {
								Type:        schema.TypeString,
								Description: "status of the condition, one of True, False, Unknown.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"type": {
								Type:        schema.TypeString,
								Description: "type of condition in CamelCase or in foo.example.com/CamelCase. --- Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"last_handled_reconcile_at": {
						Type:        schema.TypeString,
						Description: "LastHandledReconcileAt holds the value of the most recent reconcile request value, so a change of the annotation value can be detected.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"last_scan_result": {
						Type:        schema.TypeList,
						Description: "LastScanResult contains the number of fetched tags.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"latest_tags": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"scan_time": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"tag_count": {
								Type:        schema.TypeInt,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"observed_exclusion_list": {
						Type:        schema.TypeList,
						Description: "ObservedExclusionList is a list of observed exclusion list. It reflects the exclusion rules used for the observed scan result in spec.lastScanResult.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "ObservedGeneration is the last reconciled generation.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "image.toolkit.fluxcd.io/v1beta2", "ImageRepository", "image.toolkit.fluxcd.io/v1beta2/ImageRepository"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "spec.access_from", "spec.cert_secret_ref", "spec.secret_ref", "status", "status.last_scan_result"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2CompatibleVersions = []string{
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
	"v2.7.0",
	"v2.7.1",
	"v2.7.2",
	"v2.7.3",
	"v2.7.4",
	"v2.7.5",
}

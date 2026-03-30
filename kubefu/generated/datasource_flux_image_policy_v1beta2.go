package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2Read,
		Description: "ImagePolicy is the Schema for the imagepolicies API",
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
				Description: "ImagePolicySpec defines the parameters for calculating the ImagePolicy.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"digest_reflection_policy": {
						Type:        schema.TypeString,
						Description: "DigestReflectionPolicy governs the setting of the `.status.latestRef.digest` field.\n\nNever: The digest field will always be set to the empty string.\n\nIfNotPresent: The digest field will be set to the digest of the elected\nlatest image if the field is empty and the image did not change.\n\nAlways: The digest field will always be set to the digest of the elected\nlatest image.\n\nDefault: Never.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"filter_tags": {
						Type:        schema.TypeList,
						Description: "FilterTags enables filtering for only a subset of tags based on a set of rules. If no rules are provided, all the tags from the repository will be ordered and compared.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"extract": {
								Type:        schema.TypeString,
								Description: "Extract allows a capture group to be extracted from the specified regular expression pattern, useful before tag evaluation.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"pattern": {
								Type:        schema.TypeString,
								Description: "Pattern specifies a regular expression pattern used to filter for image tags.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"image_repository_ref": {
						Type:        schema.TypeList,
						Description: "ImageRepositoryRef points at the object specifying the image being scanned",
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
							"namespace": {
								Type:        schema.TypeString,
								Description: "Namespace of the referent, when not specified it acts as LocalObjectReference.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval is the length of time to wait between\nrefreshing the digest of the latest tag when the\nreflection policy is set to \"Always\".\n\nDefaults to 10m.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"policy": {
						Type:        schema.TypeList,
						Description: "Policy gives the particulars of the policy to be followed in selecting the most recent image",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"alphabetical": {
								Type:        schema.TypeList,
								Description: "Alphabetical set of rules to use for alphabetical ordering of the tags.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"order": {
										Type:        schema.TypeString,
										Description: "Order specifies the sorting order of the tags. Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"numerical": {
								Type:        schema.TypeList,
								Description: "Numerical set of rules to use for numerical ordering of the tags.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"order": {
										Type:        schema.TypeString,
										Description: "Order specifies the sorting order of the tags. Given the integer values from 0 to 9 as tags, ascending order would select 9, and descending order would select 0.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"semver": {
								Type:        schema.TypeList,
								Description: "SemVer gives a semantic version range to check against the tags available.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"range": {
										Type:        schema.TypeString,
										Description: "Range gives a semver range for the image tag; the highest version within the range that's a tag yields the latest image.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"suspend": {
						Type:        schema.TypeBool,
						Description: "This flag tells the controller to suspend subsequent policy reconciliations.\nIt does not apply to already started reconciliations. Defaults to false.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "ImagePolicyStatus defines the observed state of ImagePolicy",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
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
						Description: "LastHandledReconcileAt holds the value of the most recent\nreconcile request value, so a change of the annotation value\ncan be detected.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"latest_image": {
						Type:        schema.TypeString,
						Description: "LatestImage gives the first in the list of images scanned by the image repository, when filtered and ordered according to the policy.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"latest_ref": {
						Type:        schema.TypeList,
						Description: "LatestRef gives the first in the list of images scanned by\nthe image repository, when filtered and ordered according\nto the policy.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"digest": {
								Type:        schema.TypeString,
								Description: "Digest is the image's digest.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the bare image's name.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"tag": {
								Type:        schema.TypeString,
								Description: "Tag is the image's tag.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_previous_image": {
						Type:        schema.TypeString,
						Description: "ObservedPreviousImage is the observed previous LatestImage. It is used to keep track of the previous and current images.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_previous_ref": {
						Type:        schema.TypeList,
						Description: "ObservedPreviousRef is the observed previous LatestRef. It is used\nto keep track of the previous and current images.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"digest": {
								Type:        schema.TypeString,
								Description: "Digest is the image's digest.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the bare image's name.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"tag": {
								Type:        schema.TypeString,
								Description: "Tag is the image's tag.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
				}},
			},
		},
	}
}



func dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "image.toolkit.fluxcd.io/v1beta2", "ImagePolicy", "image.toolkit.fluxcd.io/v1beta2/ImagePolicy"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.filter_tags", "spec.image_repository_ref", "spec.policy", "spec.policy.alphabetical", "spec.policy.numerical", "spec.policy.semver", "status", "status.latest_ref", "status.observed_previous_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2CompatibleVersions = []string{
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

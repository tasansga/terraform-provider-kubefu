package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1Read,
		Description: "ArtifactGenerator is the Schema for the artifactgenerators API.",
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
				Description: "ArtifactGeneratorSpec defines the desired state of ArtifactGenerator.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"artifacts": {
						Type:        schema.TypeList,
						Description: "OutputArtifacts is a list of output artifacts to be generated.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"copy": {
								Type:        schema.TypeList,
								Description: "Copy defines a list of copy operations to perform from the sources to the generated artifact.\nThe copy operations are performed in the order they are listed with existing files\nbeing overwritten by later copy operations.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"exclude": {
										Type:        schema.TypeList,
										Description: "Exclude specifies a list of glob patterns to exclude\nfiles and dirs matched by the 'From' field.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"from": {
										Type:        schema.TypeString,
										Description: "From specifies the source (by alias) and the glob pattern to match files.\nThe format is \"@<alias>/<glob-pattern>\".",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"strategy": {
										Type:        schema.TypeString,
										Description: "Strategy specifies the copy strategy to use.\n'Overwrite' will overwrite existing files in the destination.\n'Merge' is for merging YAML files using Helm values merge strategy.\nIf not specified, defaults to 'Overwrite'.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"to": {
										Type:        schema.TypeString,
										Description: "To specifies the destination path within the artifact.\nThe format is \"@artifact/path\", the alias \"artifact\"\nrefers to the root path of the generated artifact.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the generated artifact.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"origin_revision": {
								Type:        schema.TypeString,
								Description: "OriginRevision is used to set the 'org.opencontainers.image.revision'\nannotation on the generated artifact metadata.\nIf specified, it must point to an existing source alias in the format \"@<alias>\".\nIf the referenced source has an origin revision (e.g. a Git commit SHA),\nit will be used to set the annotation on the generated artifact.\nIf the referenced source does not have an origin revision, the field is ignored.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"revision": {
								Type:        schema.TypeString,
								Description: "Revision is the revision of the generated artifact.\nIf specified, it must point to an existing source alias in the format \"@<alias>\".\nIf not specified, the revision is automatically set to the digest of the artifact content.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"sources": {
						Type:        schema.TypeList,
						Description: "Sources is a list of references to the Flux source-controller\nresources that will be used to generate the artifact.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"alias": {
								Type:        schema.TypeString,
								Description: "Alias of the source within the ArtifactGenerator context.\nThe alias must be unique per ArtifactGenerator, and must consist\nof lower case alphanumeric characters, underscores, and hyphens.\nIt must start and end with an alphanumeric character.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind of the source.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the source.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"namespace": {
								Type:        schema.TypeString,
								Description: "Namespace of the source.\nIf not provided, defaults to the same namespace as the ArtifactGenerator.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "ArtifactGeneratorStatus defines the observed state of ArtifactGenerator.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"conditions": {
						Type:        schema.TypeList,
						Description: "Conditions holds the conditions for the ArtifactGenerator.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"last_transition_time": {
								Type:        schema.TypeString,
								Description: "lastTransitionTime is the last time the condition transitioned from one status to another.\nThis should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"message": {
								Type:        schema.TypeString,
								Description: "message is a human readable message indicating details about the transition.\nThis may be an empty string.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"observed_generation": {
								Type:        schema.TypeInt,
								Description: "observedGeneration represents the .metadata.generation that the condition was set based upon.\nFor instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date\nwith respect to the current state of the instance.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"reason": {
								Type:        schema.TypeString,
								Description: "reason contains a programmatic identifier indicating the reason for the condition's last transition.\nProducers of specific condition types may define expected values and meanings for this field,\nand whether the values are considered a guaranteed API.\nThe value should be a CamelCase string.\nThis field may not be empty.",
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
								Description: "type of condition in CamelCase or in foo.example.com/CamelCase.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"inventory": {
						Type:        schema.TypeList,
						Description: "Inventory contains the list of generated ExternalArtifact references.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"digest": {
								Type:        schema.TypeString,
								Description: "Digest of the referent artifact.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"filename": {
								Type:        schema.TypeString,
								Description: "Filename is the name of the artifact file.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the referent artifact.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"namespace": {
								Type:        schema.TypeString,
								Description: "Namespace of the referent artifact.",
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
					"observed_sources_digest": {
						Type:        schema.TypeString,
						Description: "ObservedSourcesDigest is a hash representing the current state of\nall the sources referenced by the ArtifactGenerator.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "source.extensions.fluxcd.io/v1beta1", "ArtifactGenerator", "source.extensions.fluxcd.io/v1beta1/ArtifactGenerator"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1CompatibleVersions = []string{
	"v2.7.0",
	"v2.7.1",
	"v2.7.2",
	"v2.7.3",
	"v2.7.4",
	"v2.7.5",
}

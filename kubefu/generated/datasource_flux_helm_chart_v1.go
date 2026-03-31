package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxSourceToolkitFluxcdIoHelmChartV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Read,
		Description: "HelmChart is the Schema for the helmcharts API.",
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
				Description: "HelmChartSpec specifies the desired state of a Helm chart.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"chart": {
						Type:        schema.TypeString,
						Description: "Chart is the name or path the Helm chart is available at in the\nSourceRef.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"ignore_missing_values_files": {
						Type:        schema.TypeBool,
						Description: "IgnoreMissingValuesFiles controls whether to silently ignore missing values\nfiles rather than failing.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval at which the HelmChart SourceRef is checked for updates.\nThis interval is approximate and may be subject to jitter to ensure\nefficient use of resources.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"reconcile_strategy": {
						Type:        schema.TypeString,
						Description: "ReconcileStrategy determines what enables the creation of a new artifact.\nValid values are ('ChartVersion', 'Revision').\nSee the documentation of the values for an explanation on their behavior.\nDefaults to ChartVersion when omitted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"source_ref": {
						Type:        schema.TypeList,
						Description: "SourceRef is the reference to the Source the chart is available at.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_version": {
								Type:        schema.TypeString,
								Description: "APIVersion of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind of the referent, valid values are ('HelmRepository', 'GitRepository',\n'Bucket').",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"suspend": {
						Type:        schema.TypeBool,
						Description: "Suspend tells the controller to suspend the reconciliation of this\nsource.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"values_files": {
						Type:        schema.TypeList,
						Description: "ValuesFiles is an alternative list of values files to use as the chart\nvalues (values.yaml is not included by default), expected to be a\nrelative path in the SourceRef.\nValues files are merged in the order of this list with the last file\noverriding the first. Ignored when omitted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"verify": {
						Type:        schema.TypeList,
						Description: "Verify contains the secret name containing the trusted public keys\nused to verify the signature and specifies which provider to use to check\nwhether OCI image is authentic.\nThis field is only supported when using HelmRepository source with spec.type 'oci'.\nChart dependencies, which are not bundled in the umbrella chart artifact, are not verified.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"match_oidc_identity": {
								Type:        schema.TypeList,
								Description: "MatchOIDCIdentity specifies the identity matching criteria to use\nwhile verifying an OCI artifact which was signed using Cosign keyless\nsigning. The artifact's identity is deemed to be verified if any of the\nspecified matchers match against the identity.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"issuer": {
										Type:        schema.TypeString,
										Description: "Issuer specifies the regex pattern to match against to verify\nthe OIDC issuer in the Fulcio certificate. The pattern must be a\nvalid Go regular expression.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"subject": {
										Type:        schema.TypeString,
										Description: "Subject specifies the regex pattern to match against to verify\nthe identity subject in the Fulcio certificate. The pattern must\nbe a valid Go regular expression.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"provider_": {
								Type:        schema.TypeString,
								Description: "Provider specifies the technology used to sign the OCI Artifact.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"secret_ref": {
								Type:        schema.TypeList,
								Description: "SecretRef specifies the Kubernetes Secret containing the\ntrusted public keys.",
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
						}},
					},
					"version": {
						Type:        schema.TypeString,
						Description: "Version is the chart version semver expression, ignored for charts from\nGitRepository and Bucket sources. Defaults to latest when omitted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "HelmChartStatus records the observed state of the HelmChart.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"artifact": {
						Type:        schema.TypeList,
						Description: "Artifact represents the output of the last successful reconciliation.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"digest": {
								Type:        schema.TypeString,
								Description: "Digest is the digest of the file in the form of '<algorithm>:<checksum>'.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"last_update_time": {
								Type:        schema.TypeString,
								Description: "LastUpdateTime is the timestamp corresponding to the last update of the\nArtifact.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"metadata": {
								Type:        schema.TypeMap,
								Description: "Metadata holds upstream information such as OCI annotations.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"path": {
								Type:        schema.TypeString,
								Description: "Path is the relative file path of the Artifact. It can be used to locate\nthe file in the root of the Artifact storage on the local file system of\nthe controller managing the Source.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"revision": {
								Type:        schema.TypeString,
								Description: "Revision is a human-readable identifier traceable in the origin source\nsystem. It can be a Git commit SHA, Git tag, a Helm chart version, etc.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"size": {
								Type:        schema.TypeInt,
								Description: "Size is the number of bytes in the file.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"url": {
								Type:        schema.TypeString,
								Description: "URL is the HTTP address of the Artifact as exposed by the controller\nmanaging the Source. It can be used to retrieve the Artifact for\nconsumption, e.g. by another controller applying the Artifact contents.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"conditions": {
						Type:        schema.TypeList,
						Description: "Conditions holds the conditions for the HelmChart.",
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
								Description: "type of condition in CamelCase or in foo.example.com/CamelCase.\n---\nMany .condition.type values are consistent across resources like Available, but because arbitrary conditions can be\nuseful (see .node.status.conditions), the ability to deconflict is important.\nThe regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
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
					"observed_chart_name": {
						Type:        schema.TypeString,
						Description: "ObservedChartName is the last observed chart name as specified by the\nresolved chart reference.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "ObservedGeneration is the last observed generation of the HelmChart\nobject.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_source_artifact_revision": {
						Type:        schema.TypeString,
						Description: "ObservedSourceArtifactRevision is the last observed Artifact.Revision\nof the HelmChartSpec.SourceRef.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_values_files": {
						Type:        schema.TypeList,
						Description: "ObservedValuesFiles are the observed value files of the last successful\nreconciliation.\nIt matches the chart in the last successfully reconciled artifact.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL is the dynamic fetch link for the latest Artifact.\nIt is provided on a \"best effort\" basis, and using the precise\nBucketStatus.Artifact data is recommended.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "source.toolkit.fluxcd.io/v1", "HelmChart", "source.toolkit.fluxcd.io/v1/HelmChart"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.source_ref", "spec.verify", "spec.verify.secret_ref", "status", "status.artifact"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxSourceToolkitFluxcdIoHelmChartV1CompatibleVersions = []string{
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
	"v2.8.0",
	"v2.8.1",
	"v2.8.2",
	"v2.8.3",
}

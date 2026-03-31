package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Read,
		Description: "ImageUpdateAutomation is the Schema for the imageupdateautomations API",
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
				Description: "ImageUpdateAutomationSpec defines the desired state of ImageUpdateAutomation",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"git": {
						Type:        schema.TypeList,
						Description: "GitSpec contains all the git-specific definitions. This is\ntechnically optional, but in practice mandatory until there are\nother kinds of source allowed.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"checkout": {
								Type:        schema.TypeList,
								Description: "Checkout gives the parameters for cloning the git repository,\nready to make changes. If not present, the `spec.ref` field from the\nreferenced `GitRepository` or its default will be used.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"ref": {
										Type:        schema.TypeList,
										Description: "Reference gives a branch, tag or commit to clone from the Git\nrepository.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"branch": {
												Type:        schema.TypeString,
												Description: "Branch to check out, defaults to 'master' if no other field is defined.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"commit": {
												Type:        schema.TypeString,
												Description: "Commit SHA to check out, takes precedence over all reference fields.\n\nThis can be combined with Branch to shallow clone the branch, in which\nthe commit is expected to exist.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.\n\nIt must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description\nExamples: \"refs/heads/main\", \"refs/tags/v0.1.0\", \"refs/pull/420/head\", \"refs/merge-requests/1/head\"",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"semver": {
												Type:        schema.TypeString,
												Description: "SemVer tag expression to check out, takes precedence over Tag.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"tag": {
												Type:        schema.TypeString,
												Description: "Tag to check out, takes precedence over Branch.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
								}},
							},
							"commit": {
								Type:        schema.TypeList,
								Description: "Commit specifies how to commit to the git repository.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"author": {
										Type:        schema.TypeList,
										Description: "Author gives the email and optionally the name to use as the\nauthor of commits.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"email": {
												Type:        schema.TypeString,
												Description: "Email gives the email to provide when making a commit.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "Name gives the name to provide when making a commit.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"message_template": {
										Type:        schema.TypeString,
										Description: "MessageTemplate provides a template for the commit message,\ninto which will be interpolated the details of the change made.\nNote: The `Updated` template field has been removed. Use `Changed` instead.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"message_template_values": {
										Type:        schema.TypeMap,
										Description: "MessageTemplateValues provides additional values to be available to the\ntemplating rendering.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"signing_key": {
										Type:        schema.TypeList,
										Description: "SigningKey provides the option to sign commits with a GPG key",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"secret_ref": {
												Type:        schema.TypeList,
												Description: "SecretRef holds the name to a secret that contains a 'git.asc' key\ncorresponding to the ASCII Armored file containing the GPG signing\nkeypair as the value. It must be in the same namespace as the\nImageUpdateAutomation.",
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
								}},
							},
							"push": {
								Type:        schema.TypeList,
								Description: "Push specifies how and where to push commits made by the\nautomation. If missing, commits are pushed (back) to\n`.spec.checkout.branch` or its default.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"branch": {
										Type:        schema.TypeString,
										Description: "Branch specifies that commits should be pushed to the branch\nnamed. The branch is created using `.spec.checkout.branch` as the\nstarting point, if it doesn't already exist.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"options": {
										Type:        schema.TypeMap,
										Description: "Options specifies the push options that are sent to the Git\nserver when performing a push operation. For details, see:\nhttps://git-scm.com/docs/git-push#Documentation/git-push.txt---push-optionltoptiongt",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"refspec": {
										Type:        schema.TypeString,
										Description: "Refspec specifies the Git Refspec to use for a push operation.\nIf both Branch and Refspec are provided, then the commit is pushed\nto the branch and also using the specified refspec.\nFor more details about Git Refspecs, see:\nhttps://git-scm.com/book/en/v2/Git-Internals-The-Refspec",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval gives an lower bound for how often the automation\nrun should be attempted.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"policy_selector": {
						Type:        schema.TypeList,
						Description: "PolicySelector allows to filter applied policies based on labels.\nBy default includes all policies in namespace.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"match_expressions": {
								Type:        schema.TypeList,
								Description: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Description: "key is the label key that the selector applies to.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"operator": {
										Type:        schema.TypeString,
										Description: "operator represents a key's relationship to a set of values.\nValid operators are In, NotIn, Exists and DoesNotExist.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"values": {
										Type:        schema.TypeList,
										Description: "values is an array of string values. If the operator is In or NotIn,\nthe values array must be non-empty. If the operator is Exists or DoesNotExist,\nthe values array must be empty. This array is replaced during a strategic\nmerge patch.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
								}},
							},
							"match_labels": {
								Type:        schema.TypeMap,
								Description: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels\nmap is equivalent to an element of matchExpressions, whose key field is \"key\", the\noperator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"source_ref": {
						Type:        schema.TypeList,
						Description: "SourceRef refers to the resource giving access details\nto a git repository.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_version": {
								Type:        schema.TypeString,
								Description: "API version of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind of the referent.",
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
							"namespace": {
								Type:        schema.TypeString,
								Description: "Namespace of the referent, defaults to the namespace of the Kubernetes resource object that contains the reference.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"suspend": {
						Type:        schema.TypeBool,
						Description: "Suspend tells the controller to not run this automation, until\nit is unset (or set to false). Defaults to false.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"update": {
						Type:        schema.TypeList,
						Description: "Update gives the specification for how to update the files in\nthe repository. This can be left empty, to use the default\nvalue.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"path": {
								Type:        schema.TypeString,
								Description: "Path to the directory containing the manifests to be updated.\nDefaults to 'None', which translates to the root path\nof the GitRepositoryRef.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"strategy": {
								Type:        schema.TypeString,
								Description: "Strategy names the strategy to be used.",
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
				Description: "ImageUpdateAutomationStatus defines the observed state of ImageUpdateAutomation",
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
					"last_automation_run_time": {
						Type:        schema.TypeString,
						Description: "LastAutomationRunTime records the last time the controller ran\nthis automation through to completion (even if no updates were\nmade).",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"last_handled_reconcile_at": {
						Type:        schema.TypeString,
						Description: "LastHandledReconcileAt holds the value of the most recent\nreconcile request value, so a change of the annotation value\ncan be detected.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"last_push_commit": {
						Type:        schema.TypeString,
						Description: "LastPushCommit records the SHA1 of the last commit made by the\ncontroller, for this automation object",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"last_push_time": {
						Type:        schema.TypeString,
						Description: "LastPushTime records the time of the last pushed change.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_policies": {
						Type:        schema.TypeMap,
						Description: "ObservedPolicies is the list of observed ImagePolicies that were\nconsidered by the ImageUpdateAutomation update process.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_source_revision": {
						Type:        schema.TypeString,
						Description: "ObservedPolicies []ObservedPolicy `json:\"observedPolicies,omitempty\"`\nObservedSourceRevision is the last observed source revision. This can be\nused to determine if the source has been updated since last observation.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "image.toolkit.fluxcd.io/v1", "ImageUpdateAutomation", "image.toolkit.fluxcd.io/v1/ImageUpdateAutomation"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.git", "spec.git.checkout", "spec.git.checkout.ref", "spec.git.commit", "spec.git.commit.author", "spec.git.commit.signing_key", "spec.git.commit.signing_key.secret_ref", "spec.git.push", "spec.policy_selector", "spec.source_ref", "spec.update", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1CompatibleVersions = []string{
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

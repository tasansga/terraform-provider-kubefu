package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Read,
		Description: "OCIRepository is the Schema for the ocirepositories API",
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
				Description: "OCIRepositorySpec defines the desired state of OCIRepository",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"cert_secret_ref": {
						Type:        schema.TypeList,
						Description: "CertSecretRef can be given the name of a Secret containing\neither or both of\n\n- a PEM-encoded client certificate (`tls.crt`) and private\nkey (`tls.key`);\n- a PEM-encoded CA certificate (`ca.crt`)\n\nand whichever are supplied, will be used for connecting to the\nregistry. The client cert and key are useful if you are\nauthenticating with a certificate; the CA cert is useful if\nyou are using a self-signed server certificate. The Secret must\nbe of type `Opaque` or `kubernetes.io/tls`.",
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
					"ignore": {
						Type:        schema.TypeString,
						Description: "Ignore overrides the set of excluded patterns in the .sourceignore format\n(which is the same as .gitignore). If not provided, a default will be used,\nconsult the documentation for your version to find out what those are.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"insecure": {
						Type:        schema.TypeBool,
						Description: "Insecure allows connecting to a non-TLS HTTP container registry.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval at which the OCIRepository URL is checked for updates.\nThis interval is approximate and may be subject to jitter to ensure\nefficient use of resources.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"layer_selector": {
						Type:        schema.TypeList,
						Description: "LayerSelector specifies which layer should be extracted from the OCI artifact.\nWhen not specified, the first layer found in the artifact is selected.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"media_type": {
								Type:        schema.TypeString,
								Description: "MediaType specifies the OCI media type of the layer\nwhich should be extracted from the OCI Artifact. The\nfirst layer matching this type is selected.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"operation": {
								Type:        schema.TypeString,
								Description: "Operation specifies how the selected layer should be processed.\nBy default, the layer compressed content is extracted to storage.\nWhen the operation is set to 'copy', the layer compressed content\nis persisted to storage as it is.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"provider_": {
						Type:        schema.TypeString,
						Description: "The provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.\nWhen not specified, defaults to 'generic'.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"proxy_secret_ref": {
						Type:        schema.TypeList,
						Description: "ProxySecretRef specifies the Secret containing the proxy configuration\nto use while communicating with the container registry.",
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
					"ref": {
						Type:        schema.TypeList,
						Description: "The OCI reference to pull and monitor for changes,\ndefaults to the latest tag.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"digest": {
								Type:        schema.TypeString,
								Description: "Digest is the image digest to pull, takes precedence over SemVer.\nThe value should be in the format 'sha256:<HASH>'.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"semver": {
								Type:        schema.TypeString,
								Description: "SemVer is the range of tags to pull selecting the latest within\nthe range, takes precedence over Tag.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"semver_filter": {
								Type:        schema.TypeString,
								Description: "SemverFilter is a regex pattern to filter the tags within the SemVer range.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"tag": {
								Type:        schema.TypeString,
								Description: "Tag is the image tag to pull, defaults to latest.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"secret_ref": {
						Type:        schema.TypeList,
						Description: "SecretRef contains the secret name containing the registry login\ncredentials to resolve image metadata.\nThe secret must be of type kubernetes.io/dockerconfigjson.",
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
						Description: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate\nthe image pull if the service account has attached pull secrets. For more information:\nhttps://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"suspend": {
						Type:        schema.TypeBool,
						Description: "This flag tells the controller to suspend the reconciliation of this source.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"timeout": {
						Type:        schema.TypeString,
						Description: "The timeout for remote OCI Repository operations like pulling, defaults to 60s.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL is a reference to an OCI artifact repository hosted\non a remote container registry.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"verify": {
						Type:        schema.TypeList,
						Description: "Verify contains the secret name containing the trusted public keys\nused to verify the signature and specifies which provider to use to check\nwhether OCI image is authentic.",
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
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "OCIRepositoryStatus defines the observed state of OCIRepository",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"artifact": {
						Type:        schema.TypeList,
						Description: "Artifact represents the output of the last successful OCI Repository sync.",
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
						Description: "Conditions holds the conditions for the OCIRepository.",
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
					"last_handled_reconcile_at": {
						Type:        schema.TypeString,
						Description: "LastHandledReconcileAt holds the value of the most recent\nreconcile request value, so a change of the annotation value\ncan be detected.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "ObservedGeneration is the last observed generation.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_ignore": {
						Type:        schema.TypeString,
						Description: "ObservedIgnore is the observed exclusion patterns used for constructing\nthe source artifact.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"observed_layer_selector": {
						Type:        schema.TypeList,
						Description: "ObservedLayerSelector is the observed layer selector used for constructing\nthe source artifact.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"media_type": {
								Type:        schema.TypeString,
								Description: "MediaType specifies the OCI media type of the layer\nwhich should be extracted from the OCI Artifact. The\nfirst layer matching this type is selected.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"operation": {
								Type:        schema.TypeString,
								Description: "Operation specifies how the selected layer should be processed.\nBy default, the layer compressed content is extracted to storage.\nWhen the operation is set to 'copy', the layer compressed content\nis persisted to storage as it is.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL is the download link for the artifact output of the last OCI Repository sync.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "source.toolkit.fluxcd.io/v1", "OCIRepository", "source.toolkit.fluxcd.io/v1/OCIRepository"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "spec.cert_secret_ref", "spec.layer_selector", "spec.proxy_secret_ref", "spec.ref", "spec.secret_ref", "spec.verify", "spec.verify.secret_ref", "status", "status.artifact", "status.observed_layer_selector"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1CompatibleVersions = []string{
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

package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Read,
		Description: "HelmRepository is the Schema for the helmrepositories API.",
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
				Description: "HelmRepositorySpec specifies the required configuration to produce an\nArtifact for a Helm repository index YAML.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"access_from": {
						Type:        schema.TypeList,
						Description: "AccessFrom specifies an Access Control List for allowing cross-namespace\nreferences to this object.\nNOTE: Not implemented, provisional as of https://github.com/fluxcd/flux2/pull/2092",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"namespace_selectors": {
								Type:        schema.TypeList,
								Description: "NamespaceSelectors is the list of namespace selectors to which this ACL applies.\nItems in this list are evaluated using a logical OR operation.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"match_labels": {
										Type:        schema.TypeMap,
										Description: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels\nmap is equivalent to an element of matchExpressions, whose key field is \"key\", the\noperator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
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
						Description: "CertSecretRef can be given the name of a Secret containing\neither or both of\n\n\n- a PEM-encoded client certificate (`tls.crt`) and private\nkey (`tls.key`);\n- a PEM-encoded CA certificate (`ca.crt`)\n\n\nand whichever are supplied, will be used for connecting to the\nregistry. The client cert and key are useful if you are\nauthenticating with a certificate; the CA cert is useful if\nyou are using a self-signed server certificate. The Secret must\nbe of type `Opaque` or `kubernetes.io/tls`.\n\n\nIt takes precedence over the values specified in the Secret referred\nto by `.spec.secretRef`.",
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
					"insecure": {
						Type:        schema.TypeBool,
						Description: "Insecure allows connecting to a non-TLS HTTP container registry.\nThis field is only taken into account if the .spec.type field is set to 'oci'.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "Interval at which the HelmRepository URL is checked for updates.\nThis interval is approximate and may be subject to jitter to ensure\nefficient use of resources.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"pass_credentials": {
						Type:        schema.TypeBool,
						Description: "PassCredentials allows the credentials from the SecretRef to be passed\non to a host that does not match the host as defined in URL.\nThis may be required if the host of the advertised chart URLs in the\nindex differ from the defined URL.\nEnabling this should be done with caution, as it can potentially result\nin credentials getting stolen in a MITM-attack.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"provider_": {
						Type:        schema.TypeString,
						Description: "Provider used for authentication, can be 'aws', 'azure', 'gcp' or 'generic'.\nThis field is optional, and only taken into account if the .spec.type field is set to 'oci'.\nWhen not specified, defaults to 'generic'.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"secret_ref": {
						Type:        schema.TypeList,
						Description: "SecretRef specifies the Secret containing authentication credentials\nfor the HelmRepository.\nFor HTTP/S basic auth the secret must contain 'username' and 'password'\nfields.\nSupport for TLS auth using the 'certFile' and 'keyFile', and/or 'caFile'\nkeys is deprecated. Please use `.spec.certSecretRef` instead.",
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
					"suspend": {
						Type:        schema.TypeBool,
						Description: "Suspend tells the controller to suspend the reconciliation of this\nHelmRepository.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"timeout": {
						Type:        schema.TypeString,
						Description: "Timeout is used for the index fetch operation for an HTTPS helm repository,\nand for remote OCI Repository operations like pulling for an OCI helm\nchart by the associated HelmChart.\nIts default value is 60s.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"type": {
						Type:        schema.TypeString,
						Description: "Type of the HelmRepository.\nWhen this field is set to  \"oci\", the URL field value must be prefixed with \"oci://\".",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL of the Helm repository, a valid URL contains at least a protocol and\nhost.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "HelmRepositoryStatus records the observed state of the HelmRepository.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"artifact": {
						Type:        schema.TypeList,
						Description: "Artifact represents the last successful HelmRepository reconciliation.",
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
						Description: "Conditions holds the conditions for the HelmRepository.",
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
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "ObservedGeneration is the last observed generation of the HelmRepository\nobject.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL is the dynamic fetch link for the latest Artifact.\nIt is provided on a \"best effort\" basis, and using the precise\nHelmRepositoryStatus.Artifact data is recommended.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "source.toolkit.fluxcd.io/v1", "HelmRepository", "source.toolkit.fluxcd.io/v1/HelmRepository"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.access_from", "spec.cert_secret_ref", "spec.secret_ref", "status", "status.artifact"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1CompatibleVersions = []string{
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

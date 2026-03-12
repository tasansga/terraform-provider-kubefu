package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2Read,
		Description: "OCIRepository is the Schema for the ocirepositories API",
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
				Description: "OCIRepositorySpec defines the desired state of OCIRepository",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"cert_secret_ref": {
						Type:        schema.TypeList,
						Description: "CertSecretRef can be given the name of a secret containing either or both of \n  - a PEM-encoded client certificate (`certFile`) and private  key (`keyFile`);  - a PEM-encoded CA certificate (`caFile`) \n  and whichever are supplied, will be used for connecting to the  registry. The client cert and key are useful if you are  authenticating with a certificate; the CA cert is useful if  you are using a self-signed server certificate.",
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
						Description: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"interval": {
						Type:        schema.TypeString,
						Description: "The interval at which to check for image updates.",
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
					"ref": {
						Type:        schema.TypeList,
						Description: "The OCI reference to pull and monitor for changes, defaults to the latest tag.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"digest": {
								Type:        schema.TypeString,
								Description: "Digest is the image digest to pull, takes precedence over SemVer. The value should be in the format 'sha256:<HASH>'.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"semver": {
								Type:        schema.TypeString,
								Description: "SemVer is the range of tags to pull selecting the latest within the range, takes precedence over Tag.",
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
						Description: "SecretRef contains the secret name containing the registry login credentials to resolve image metadata. The secret must be of type kubernetes.io/dockerconfigjson.",
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
						Description: "ServiceAccountName is the name of the Kubernetes ServiceAccount used to authenticate the image pull if the service account has attached pull secrets. For more information: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#add-imagepullsecrets-to-a-service-account",
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
						Description: "URL is a reference to an OCI artifact repository hosted on a remote container registry.",
						Optional:    true,
						Required:    false,
						Computed:    true,
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
							"checksum": {
								Type:        schema.TypeString,
								Description: "Checksum is the SHA256 checksum of the Artifact file.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"last_update_time": {
								Type:        schema.TypeString,
								Description: "LastUpdateTime is the timestamp corresponding to the last update of the Artifact.",
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
								Description: "Path is the relative file path of the Artifact. It can be used to locate the file in the root of the Artifact storage on the local file system of the controller managing the Source.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"revision": {
								Type:        schema.TypeString,
								Description: "Revision is a human-readable identifier traceable in the origin source system. It can be a Git commit SHA, Git tag, a Helm chart version, etc.",
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
								Description: "URL is the HTTP address of the Artifact as exposed by the controller managing the Source. It can be used to retrieve the Artifact for consumption, e.g. by another controller applying the Artifact contents.",
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
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "ObservedGeneration is the last observed generation.",
						Optional:    true,
						Required:    false,
						Computed:    true,
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



func dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "source.toolkit.fluxcd.io/v1beta2", "OCIRepository", "source.toolkit.fluxcd.io/v1beta2/OCIRepository"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "spec.cert_secret_ref", "spec.ref", "spec.secret_ref", "status", "status.artifact"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2CompatibleVersions = []string{
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
	"v2.7.0",
	"v2.7.1",
	"v2.7.2",
	"v2.7.3",
	"v2.7.4",
	"v2.7.5",
}

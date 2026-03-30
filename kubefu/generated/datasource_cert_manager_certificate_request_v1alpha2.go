package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2Read,
		Description: "A CertificateRequest is used to request a signed certificate from one of the configured issuers. \n All fields within the CertificateRequest's `spec` are immutable after creation. A CertificateRequest will either succeed or fail, as denoted by its `status.state` field. \n A CertificateRequest is a 'one-shot' resource, meaning it represents a single point in time request for a certificate and cannot be re-used.",
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
				Description: "Desired state of the CertificateRequest resource.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"csr": {
						Type:        schema.TypeString,
						Description: "The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"duration": {
						Type:        schema.TypeString,
						Description: "The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"extra": {
						Type:        schema.TypeMap,
						Description: "Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"groups": {
						Type:        schema.TypeList,
						Description: "Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"is_ca": {
						Type:        schema.TypeBool,
						Description: "IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the `cert sign` usage to the list of `usages`.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"issuer_ref": {
						Type:        schema.TypeList,
						Description: "IssuerRef is a reference to the issuer for this CertificateRequest.  If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to 'cert-manager.io' if empty.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"group": {
								Type:        schema.TypeString,
								Description: "Group of the resource being referred to.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind of the resource being referred to.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the resource being referred to.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"uid": {
						Type:        schema.TypeString,
						Description: "UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"usages": {
						Type:        schema.TypeList,
						Description: "Usages is the set of x509 usages that are requested for the certificate. Defaults to `digital signature` and `key encipherment` if not specified.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"username": {
						Type:        schema.TypeString,
						Description: "Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "Status of the CertificateRequest. This is set and managed automatically.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"ca": {
						Type:        schema.TypeString,
						Description: "The PEM encoded x509 certificate of the signer, also known as the CA (Certificate Authority). This is set on a best-effort basis by different issuers. If not set, the CA is assumed to be unknown/not available.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"certificate": {
						Type:        schema.TypeString,
						Description: "The PEM encoded x509 certificate resulting from the certificate signing request. If not set, the CertificateRequest has either not been completed or has failed. More information on failure can be found by checking the `conditions` field.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"conditions": {
						Type:        schema.TypeList,
						Description: "List of status conditions to indicate the status of a CertificateRequest. Known condition types are `Ready` and `InvalidRequest`.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"last_transition_time": {
								Type:        schema.TypeString,
								Description: "LastTransitionTime is the timestamp corresponding to the last status change of this condition.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"message": {
								Type:        schema.TypeString,
								Description: "Message is a human readable description of the details of the last transition, complementing reason.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"reason": {
								Type:        schema.TypeString,
								Description: "Reason is a brief machine readable explanation for the condition's last transition.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"status": {
								Type:        schema.TypeString,
								Description: "Status of the condition, one of ('True', 'False', 'Unknown').",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"type": {
								Type:        schema.TypeString,
								Description: "Type of the condition, known values are ('Ready', 'InvalidRequest').",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"failure_time": {
						Type:        schema.TypeString,
						Description: "FailureTime stores the time that this CertificateRequest failed. This is used to influence garbage collection and back-off.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "cert-manager.io/v1alpha2", "CertificateRequest", "cert-manager.io/v1alpha2/CertificateRequest"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.issuer_ref", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2CompatibleVersions = []string{
	"v0.16.0",
	"v1.0.0",
	"v1.1.0",
	"v1.2.0",
	"v1.3.0",
	"v1.4.0",
	"v1.5.0",
	"v1.6.0",
}

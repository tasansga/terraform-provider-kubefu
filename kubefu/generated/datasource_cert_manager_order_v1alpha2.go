package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2Read,
		Description: "Order is a type to represent an Order with an ACME server",
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
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"spec": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"common_name": {
						Type:        schema.TypeString,
						Description: "CommonName is the common name as specified on the DER encoded CSR. If specified, this value must also be present in `dnsNames`. This field must match the corresponding field on the DER encoded CSR.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"csr": {
						Type:        schema.TypeString,
						Description: "Certificate signing request bytes in DER encoding. This will be used when finalizing the order. This field must be set on the order.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"dns_names": {
						Type:        schema.TypeList,
						Description: "DNSNames is a list of DNS names that should be included as part of the Order validation process. This field must match the corresponding field on the DER encoded CSR.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"issuer_ref": {
						Type:        schema.TypeList,
						Description: "IssuerRef references a properly configured ACME-type Issuer which should be used to create this Order. If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Order will be marked as failed.",
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
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"authorizations": {
						Type:        schema.TypeList,
						Description: "Authorizations contains data returned from the ACME server on what authorizations must be completed in order to validate the DNS names specified on the Order.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"challenges": {
								Type:        schema.TypeList,
								Description: "Challenges specifies the challenge types offered by the ACME server. One of these challenge types will be selected when validating the DNS name and an appropriate Challenge resource will be created to perform the ACME challenge process.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"token": {
										Type:        schema.TypeString,
										Description: "Token is the token that must be presented for this challenge. This is used to compute the 'key' that must also be presented.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Type is the type of challenge being offered, e.g. 'http-01', 'dns-01', 'tls-sni-01', etc. This is the raw value retrieved from the ACME server. Only 'http-01' and 'dns-01' are supported by cert-manager, other values will be ignored.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"url": {
										Type:        schema.TypeString,
										Description: "URL is the URL of this challenge. It can be used to retrieve additional metadata about the Challenge from the ACME server.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"identifier": {
								Type:        schema.TypeString,
								Description: "Identifier is the DNS name to be validated as part of this authorization",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"initial_state": {
								Type:        schema.TypeString,
								Description: "InitialState is the initial state of the ACME authorization when first fetched from the ACME server. If an Authorization is already 'valid', the Order controller will not create a Challenge resource for the authorization. This will occur when working with an ACME server that enables 'authz reuse' (such as Let's Encrypt's production endpoint). If not set and 'identifier' is set, the state is assumed to be pending and a Challenge will be created.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"url": {
								Type:        schema.TypeString,
								Description: "URL is the URL of the Authorization that must be completed",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"wildcard": {
								Type:        schema.TypeBool,
								Description: "Wildcard will be true if this authorization is for a wildcard DNS name. If this is true, the identifier will be the *non-wildcard* version of the DNS name. For example, if '*.example.com' is the DNS name being validated, this field will be 'true' and the 'identifier' field will be 'example.com'.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"certificate": {
						Type:        schema.TypeString,
						Description: "Certificate is a copy of the PEM encoded certificate for this Order. This field will be populated after the order has been successfully finalized with the ACME server, and the order has transitioned to the 'valid' state.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"failure_time": {
						Type:        schema.TypeString,
						Description: "FailureTime stores the time that this order failed. This is used to influence garbage collection and back-off.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"finalize_url": {
						Type:        schema.TypeString,
						Description: "FinalizeURL of the Order. This is used to obtain certificates for this order once it has been completed.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"reason": {
						Type:        schema.TypeString,
						Description: "Reason optionally provides more information about a why the order is in the current state.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"state": {
						Type:        schema.TypeString,
						Description: "State contains the current state of this Order resource. States 'success' and 'expired' are 'final'",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"url": {
						Type:        schema.TypeString,
						Description: "URL of the Order. This will initially be empty when the resource is first created. The Order controller will populate this field when the Order is first processed. This field will be immutable after it is initially set.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "acme.cert-manager.io/v1alpha2", "Order", "acme.cert-manager.io/v1alpha2/Order"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "spec.issuer_ref", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2CompatibleVersions = []string{
	"v0.16.0",
	"v1.0.0",
	"v1.1.0",
	"v1.2.0",
	"v1.3.0",
	"v1.4.0",
	"v1.5.0",
	"v1.6.0",
}

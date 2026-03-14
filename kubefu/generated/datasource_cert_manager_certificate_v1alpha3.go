package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceCertManagerCertManagerIoCertificateV1Alpha3() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertManagerCertManagerIoCertificateV1Alpha3Read,
		Description: "Certificate is a type to represent a Certificate from ACME",
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
				Description: "CertificateSpec defines the desired state of Certificate. A valid Certificate requires at least one of a CommonName, DNSName, or URISAN to be valid.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"common_name": {
						Type:        schema.TypeString,
						Description: "CommonName is a common name to be used on the Certificate. The CommonName should have a length of 64 characters or fewer to avoid generating invalid CSRs. This value is ignored by TLS clients when any subject alt name is set. This is x509 behaviour: https://tools.ietf.org/html/rfc6125#section-6.4.4",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"dns_names": {
						Type:        schema.TypeList,
						Description: "DNSNames is a list of subject alt names to be used on the Certificate.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"duration": {
						Type:        schema.TypeString,
						Description: "Certificate default Duration",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"email_sa_ns": {
						Type:        schema.TypeList,
						Description: "EmailSANs is a list of Email Subject Alternative Names to be set on this Certificate.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"ip_addresses": {
						Type:        schema.TypeList,
						Description: "IPAddresses is a list of IP addresses to be used on the Certificate",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"is_ca": {
						Type:        schema.TypeBool,
						Description: "IsCA will mark this Certificate as valid for signing. This implies that the 'cert sign' usage is set",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"issuer_ref": {
						Type:        schema.TypeList,
						Description: "IssuerRef is a reference to the issuer for this certificate. If the 'kind' field is not set, or set to 'Issuer', an Issuer resource with the given name in the same namespace as the Certificate will be used. If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the provided name will be used. The 'name' field in this stanza is required at all times.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"group": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"key_algorithm": {
						Type:        schema.TypeString,
						Description: "KeyAlgorithm is the private key algorithm of the corresponding private key for this certificate. If provided, allowed values are either \"rsa\" or \"ecdsa\" If KeyAlgorithm is specified and KeySize is not provided, key size of 256 will be used for \"ecdsa\" key algorithm and key size of 2048 will be used for \"rsa\" key algorithm.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"key_encoding": {
						Type:        schema.TypeString,
						Description: "KeyEncoding is the private key cryptography standards (PKCS) for this certificate's private key to be encoded in. If provided, allowed values are \"pkcs1\" and \"pkcs8\" standing for PKCS#1 and PKCS#8, respectively. If KeyEncoding is not specified, then PKCS#1 will be used by default.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"key_size": {
						Type:        schema.TypeInt,
						Description: "KeySize is the key bit size of the corresponding private key for this certificate. If provided, value must be between 2048 and 8192 inclusive when KeyAlgorithm is empty or is set to \"rsa\", and value must be one of (256, 384, 521) when KeyAlgorithm is set to \"ecdsa\".",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"keystores": {
						Type:        schema.TypeList,
						Description: "Keystores configures additional keystore output formats stored in the `secretName` Secret resource.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"jks": {
								Type:        schema.TypeList,
								Description: "JKS configures options for storing a JKS keystore in the `spec.secretName` Secret resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"create": {
										Type:        schema.TypeBool,
										Description: "Create enables JKS keystore creation for the Certificate. If true, a file named `keystore.jks` will be created in the target Secret resource, encrypted using the password stored in `passwordSecretRef`. The keystore file will only be updated upon re-issuance.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"password_secret_ref": {
										Type:        schema.TypeList,
										Description: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the JKS keystore.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "The key of the secret to select from. Must be a valid secret key.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
								}},
							},
							"pkcs12": {
								Type:        schema.TypeList,
								Description: "PKCS12 configures options for storing a PKCS12 keystore in the `spec.secretName` Secret resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"create": {
										Type:        schema.TypeBool,
										Description: "Create enables PKCS12 keystore creation for the Certificate. If true, a file named `keystore.p12` will be created in the target Secret resource, encrypted using the password stored in `passwordSecretRef`. The keystore file will only be updated upon re-issuance.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"password_secret_ref": {
										Type:        schema.TypeList,
										Description: "PasswordSecretRef is a reference to a key in a Secret resource containing the password used to encrypt the PKCS12 keystore.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "The key of the secret to select from. Must be a valid secret key.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
					"private_key": {
						Type:        schema.TypeList,
						Description: "Options to control private keys used for the Certificate.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"rotation_policy": {
								Type:        schema.TypeString,
								Description: "RotationPolicy controls how private keys should be regenerated when a re-issuance is being processed. If set to Never, a private key will only be generated if one does not already exist in the target `spec.secretName`. If one does exists but it does not have the correct algorithm or size, a warning will be raised to await user intervention. If set to Always, a private key matching the specified requirements will be generated whenever a re-issuance occurs. Default is 'Never' for backward compatibility.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"renew_before": {
						Type:        schema.TypeString,
						Description: "Certificate renew before expiration duration",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"secret_name": {
						Type:        schema.TypeString,
						Description: "SecretName is the name of the secret resource to store this secret in",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"subject": {
						Type:        schema.TypeList,
						Description: "Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"countries": {
								Type:        schema.TypeList,
								Description: "Countries to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"localities": {
								Type:        schema.TypeList,
								Description: "Cities to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"organizational_units": {
								Type:        schema.TypeList,
								Description: "Organizational Units to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"organizations": {
								Type:        schema.TypeList,
								Description: "Organizations to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"postal_codes": {
								Type:        schema.TypeList,
								Description: "Postal codes to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"provinces": {
								Type:        schema.TypeList,
								Description: "State/Provinces to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
							"serial_number": {
								Type:        schema.TypeString,
								Description: "Serial number to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"street_addresses": {
								Type:        schema.TypeList,
								Description: "Street addresses to be used on the Certificate.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Schema{Type: schema.TypeString},
							},
						}},
					},
					"uri_sa_ns": {
						Type:        schema.TypeList,
						Description: "URISANs is a list of URI Subject Alternative Names to be set on this Certificate.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"usages": {
						Type:        schema.TypeList,
						Description: "Usages is the set of x509 actions that are enabled for a given key. Defaults are ('digital signature', 'key encipherment') if empty",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "CertificateStatus defines the observed state of Certificate",
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
								Description: "Type of the condition, currently ('Ready').",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"last_failure_time": {
						Type:        schema.TypeString,
						Description: "",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"next_private_key_secret_name": {
						Type:        schema.TypeString,
						Description: "The name of the Secret resource containing the private key to be used for the next certificate iteration. The keymanager controller will automatically set this field if the `Issuing` condition is set to `True`. It will automatically unset this field when the Issuing condition is not set or False.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"not_after": {
						Type:        schema.TypeString,
						Description: "The expiration time of the certificate stored in the secret named by this resource in spec.secretName.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"revision": {
						Type:        schema.TypeInt,
						Description: "The current 'revision' of the certificate as issued. \n When a CertificateRequest resource is created, it will have the `cert-manager.io/certificate-revision` set to one greater than the current value of this field. \n Upon issuance, this field will be set to the value of the annotation on the CertificateRequest resource used to issue the certificate. \n Persisting the value on the CertificateRequest resource allows the certificates controller to know whether a request is part of an old issuance or if it is part of the ongoing revision's issuance by checking if the revision value in the annotation is greater than this field.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceCertManagerCertManagerIoCertificateV1Alpha3Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "cert-manager.io/v1alpha3", "Certificate", "cert-manager.io/v1alpha3/Certificate"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.issuer_ref", "spec.keystores", "spec.keystores.jks", "spec.keystores.jks.password_secret_ref", "spec.keystores.pkcs12", "spec.keystores.pkcs12.password_secret_ref", "spec.private_key", "spec.subject", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceCertManagerCertManagerIoCertificateV1Alpha3CompatibleVersions = []string{
	"v0.15.0",
	"v0.16.0",
	"v1.0.0",
	"v1.1.0",
	"v1.2.0",
	"v1.3.0",
	"v1.4.0",
	"v1.5.0",
	"v1.6.0",
}

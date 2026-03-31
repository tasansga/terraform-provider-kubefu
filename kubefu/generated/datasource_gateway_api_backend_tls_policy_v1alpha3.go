package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3Read,
		Description: "BackendTLSPolicy provides a way to configure how a Gateway\nconnects to a Backend via TLS.",
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
				Description: "Spec defines the desired state of BackendTLSPolicy.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"options": {
						Type:        schema.TypeMap,
						Description: "Options are a list of key/value pairs to enable extended TLS\nconfiguration for each implementation. For example, configuring the\nminimum TLS version or supported cipher suites.\n\nA set of common keys MAY be defined by the API in the future. To avoid\nany ambiguity, implementation-specific definitions MUST use\ndomain-prefixed names, such as `example.com/my-custom-option`.\nUn-prefixed names are reserved for key names defined by Gateway API.\n\nSupport: Implementation-specific",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"target_refs": {
						Type:        schema.TypeList,
						Description: "TargetRefs identifies an API object to apply the policy to.\nOnly Services have Extended support. Implementations MAY support\nadditional objects, with Implementation Specific support.\nNote that this config applies to the entire referenced resource\nby default, but this default may change in the future to provide\na more granular application of the policy.\n\nTargetRefs must be _distinct_. This means either that:\n\n* They select different targets. If this is the case, then targetRef\n  entries are distinct. In terms of fields, this means that the\n  multi-part key defined by `group`, `kind`, and `name` must\n  be unique across all targetRef entries in the BackendTLSPolicy.\n* They select different sectionNames in the same target.\n\nWhen more than one BackendTLSPolicy selects the same target and\nsectionName, implementations MUST determine precedence using the\nfollowing criteria, continuing on ties:\n\n* The older policy by creation timestamp takes precedence. For\n  example, a policy with a creation timestamp of \"2021-07-15\n  01:02:03\" MUST be given precedence over a policy with a\n  creation timestamp of \"2021-07-15 01:02:04\".\n* The policy appearing first in alphabetical order by {name}.\n  For example, a policy named `bar` is given precedence over a\n  policy named `baz`.\n\nFor any BackendTLSPolicy that does not take precedence, the\nimplementation MUST ensure the `Accepted` Condition is set to\n`status: False`, with Reason `Conflicted`.\n\nSupport: Extended for Kubernetes Service\n\nSupport: Implementation-specific for any other resource",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"group": {
								Type:        schema.TypeString,
								Description: "Group is the group of the target resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind is kind of the target resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the target resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"section_name": {
								Type:        schema.TypeString,
								Description: "SectionName is the name of a section within the target resource. When\nunspecified, this targetRef targets the entire resource. In the following\nresources, SectionName is interpreted as the following:\n\n* Gateway: Listener name\n* HTTPRoute: HTTPRouteRule name\n* Service: Port name\n\nIf a SectionName is specified, but does not exist on the targeted object,\nthe Policy must fail to attach, and the policy implementation should record\na `ResolvedRefs` or similar Condition in the Policy's status.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"validation": {
						Type:        schema.TypeList,
						Description: "Validation contains backend TLS validation configuration.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"ca_certificate_refs": {
								Type:        schema.TypeList,
								Description: "CACertificateRefs contains one or more references to Kubernetes objects that\ncontain a PEM-encoded TLS CA certificate bundle, which is used to\nvalidate a TLS handshake between the Gateway and backend Pod.\n\nIf CACertificateRefs is empty or unspecified, then WellKnownCACertificates must be\nspecified. Only one of CACertificateRefs or WellKnownCACertificates may be specified,\nnot both. If CACertificateRefs is empty or unspecified, the configuration for\nWellKnownCACertificates MUST be honored instead if supported by the implementation.\n\nA CACertificateRef is invalid if:\n\n* It refers to a resource that cannot be resolved (e.g., the referenced resource\n  does not exist) or is misconfigured (e.g., a ConfigMap does not contain a key\n  named `ca.crt`). In this case, the Reason must be set to `InvalidCACertificateRef`\n  and the Message of the Condition must indicate which reference is invalid and why.\n\n* It refers to an unknown or unsupported kind of resource. In this case, the Reason\n  must be set to `InvalidKind` and the Message of the Condition must explain which\n  kind of resource is unknown or unsupported.\n\n* It refers to a resource in another namespace. This may change in future\n  spec updates.\n\nImplementations MAY choose to perform further validation of the certificate\ncontent (e.g., checking expiry or enforcing specific formats). In such cases,\nan implementation-specific Reason and Message must be set for the invalid reference.\n\nIn all cases, the implementation MUST ensure the `ResolvedRefs` Condition on\nthe BackendTLSPolicy is set to `status: False`, with a Reason and Message\nthat indicate the cause of the error. Connections using an invalid\nCACertificateRef MUST fail, and the client MUST receive an HTTP 5xx error\nresponse. If ALL CACertificateRefs are invalid, the implementation MUST also\nensure the `Accepted` Condition on the BackendTLSPolicy is set to\n`status: False`, with a Reason `NoValidCACertificate`.\n\nA single CACertificateRef to a Kubernetes ConfigMap kind has \"Core\" support.\nImplementations MAY choose to support attaching multiple certificates to\na backend, but this behavior is implementation-specific.\n\nSupport: Core - An optional single reference to a Kubernetes ConfigMap,\nwith the CA certificate in a key named `ca.crt`.\n\nSupport: Implementation-specific - More than one reference, other kinds\nof resources, or a single reference that includes multiple certificates.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"group": {
										Type:        schema.TypeString,
										Description: "Group is the group of the referent. For example, \"gateway.networking.k8s.io\".\nWhen unspecified or empty string, core API group is inferred.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"kind": {
										Type:        schema.TypeString,
										Description: "Kind is kind of the referent. For example \"HTTPRoute\" or \"Service\".",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name is the name of the referent.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"hostname": {
								Type:        schema.TypeString,
								Description: "Hostname is used for two purposes in the connection between Gateways and\nbackends:\n\n1. Hostname MUST be used as the SNI to connect to the backend (RFC 6066).\n2. Hostname MUST be used for authentication and MUST match the certificate\n   served by the matching backend, unless SubjectAltNames is specified.\n3. If SubjectAltNames are specified, Hostname can be used for certificate selection\n   but MUST NOT be used for authentication. If you want to use the value\n   of the Hostname field for authentication, you MUST add it to the SubjectAltNames list.\n\nSupport: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"subject_alt_names": {
								Type:        schema.TypeList,
								Description: "SubjectAltNames contains one or more Subject Alternative Names.\nWhen specified the certificate served from the backend MUST\nhave at least one Subject Alternate Name matching one of the specified SubjectAltNames.\n\nSupport: Extended",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"hostname": {
										Type:        schema.TypeString,
										Description: "Hostname contains Subject Alternative Name specified in DNS name format.\nRequired when Type is set to Hostname, ignored otherwise.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "Type determines the format of the Subject Alternative Name. Always required.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"uri": {
										Type:        schema.TypeString,
										Description: "URI contains Subject Alternative Name specified in a full URI format.\nIt MUST include both a scheme (e.g., \"http\" or \"ftp\") and a scheme-specific-part.\nCommon values include SPIFFE IDs like \"spiffe://mycluster.example.com/ns/myns/sa/svc1sa\".\nRequired when Type is set to URI, ignored otherwise.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"well_known_ca_certificates": {
								Type:        schema.TypeString,
								Description: "WellKnownCACertificates specifies whether system CA certificates may be used in\nthe TLS handshake between the gateway and backend pod.\n\nIf WellKnownCACertificates is unspecified or empty (\"\"), then CACertificateRefs\nmust be specified with at least one entry for a valid configuration. Only one of\nCACertificateRefs or WellKnownCACertificates may be specified, not both.\nIf an implementation does not support the WellKnownCACertificates field, or\nthe supplied value is not recognized, the implementation MUST ensure the\n`Accepted` Condition on the BackendTLSPolicy is set to `status: False`, with\na Reason `Invalid`.\n\nSupport: Implementation-specific",
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
				Description: "Status defines the current state of BackendTLSPolicy.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"ancestors": {
						Type:        schema.TypeList,
						Description: "Ancestors is a list of ancestor resources (usually Gateways) that are\nassociated with the policy, and the status of the policy with respect to\neach ancestor. When this policy attaches to a parent, the controller that\nmanages the parent and the ancestors MUST add an entry to this list when\nthe controller first sees the policy and SHOULD update the entry as\nappropriate when the relevant ancestor is modified.\n\nNote that choosing the relevant ancestor is left to the Policy designers;\nan important part of Policy design is designing the right object level at\nwhich to namespace this status.\n\nNote also that implementations MUST ONLY populate ancestor status for\nthe Ancestor resources they are responsible for. Implementations MUST\nuse the ControllerName field to uniquely identify the entries in this list\nthat they are responsible for.\n\nNote that to achieve this, the list of PolicyAncestorStatus structs\nMUST be treated as a map with a composite key, made up of the AncestorRef\nand ControllerName fields combined.\n\nA maximum of 16 ancestors will be represented in this list. An empty list\nmeans the Policy is not relevant for any ancestors.\n\nIf this slice is full, implementations MUST NOT add further entries.\nInstead they MUST consider the policy unimplementable and signal that\non any related resources such as the ancestor that would be referenced\nhere. For example, if this list was full on BackendTLSPolicy, no\nadditional Gateways would be able to reference the Service targeted by\nthe BackendTLSPolicy.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"ancestor_ref": {
								Type:        schema.TypeList,
								Description: "AncestorRef corresponds with a ParentRef in the spec that this\nPolicyAncestorStatus struct describes the status of.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"group": {
										Type:        schema.TypeString,
										Description: "Group is the group of the referent.\nWhen unspecified, \"gateway.networking.k8s.io\" is inferred.\nTo set the core API group (such as for a \"Service\" kind referent),\nGroup must be explicitly set to \"\" (empty string).\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"kind": {
										Type:        schema.TypeString,
										Description: "Kind is kind of the referent.\n\nThere are two kinds of parent resources with \"Core\" support:\n\n* Gateway (Gateway conformance profile)\n* Service (Mesh conformance profile, ClusterIP Services only)\n\nSupport for other resources is Implementation-Specific.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name is the name of the referent.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"namespace": {
										Type:        schema.TypeString,
										Description: "Namespace is the namespace of the referent. When unspecified, this refers\nto the local namespace of the Route.\n\nNote that there are specific rules for ParentRefs which cross namespace\nboundaries. Cross-namespace references are only valid if they are explicitly\nallowed by something in the namespace they are referring to. For example:\nGateway has the AllowedRoutes field, and ReferenceGrant provides a\ngeneric way to enable any other kind of cross-namespace reference.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeInt,
										Description: "Port is the network port this Route targets. It can be interpreted\ndifferently based on the type of parent resource.\n\nWhen the parent resource is a Gateway, this targets all listeners\nlistening on the specified port that also support this kind of Route(and\nselect this Route). It's not recommended to set `Port` unless the\nnetworking behaviors specified in a Route must apply to a specific port\nas opposed to a listener(s) whose port(s) may be changed. When both Port\nand SectionName are specified, the name and port of the selected listener\nmust match both specified values.\n\nImplementations MAY choose to support other parent resources.\nImplementations supporting other types of parent resources MUST clearly\ndocument how/if Port is interpreted.\n\nFor the purpose of status, an attachment is considered successful as\nlong as the parent resource accepts it partially. For example, Gateway\nlisteners can restrict which Routes can attach to them by Route kind,\nnamespace, or hostname. If 1 of 2 Gateway listeners accept attachment\nfrom the referencing Route, the Route MUST be considered successfully\nattached. If no Gateway listeners accept attachment from this Route,\nthe Route MUST be considered detached from the Gateway.\n\nSupport: Extended",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"section_name": {
										Type:        schema.TypeString,
										Description: "SectionName is the name of a section within the target resource. In the\nfollowing resources, SectionName is interpreted as the following:\n\n* Gateway: Listener name. When both Port (experimental) and SectionName\nare specified, the name and port of the selected listener must match\nboth specified values.\n* Service: Port name. When both Port (experimental) and SectionName\nare specified, the name and port of the selected listener must match\nboth specified values.\n\nImplementations MAY choose to support attaching Routes to other resources.\nIf that is the case, they MUST clearly document how SectionName is\ninterpreted.\n\nWhen unspecified (empty string), this will reference the entire resource.\nFor the purpose of status, an attachment is considered successful if at\nleast one section in the parent resource accepts it. For example, Gateway\nlisteners can restrict which Routes can attach to them by Route kind,\nnamespace, or hostname. If 1 of 2 Gateway listeners accept attachment from\nthe referencing Route, the Route MUST be considered successfully\nattached. If no Gateway listeners accept attachment from this Route, the\nRoute MUST be considered detached from the Gateway.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"conditions": {
								Type:        schema.TypeList,
								Description: "Conditions describes the status of the Policy with respect to the given Ancestor.",
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
							"controller_name": {
								Type:        schema.TypeString,
								Description: "ControllerName is a domain/path string that indicates the name of the\ncontroller that wrote this status. This corresponds with the\ncontrollerName field on GatewayClass.\n\nExample: \"example.net/gateway-controller\".\n\nThe format of this field is DOMAIN \"/\" PATH, where DOMAIN and PATH are\nvalid Kubernetes names\n(https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).\n\nControllers MUST populate this field when writing status. Controllers should ensure that\nentries to status populated with their ControllerName are cleaned up when they are no\nlonger necessary.",
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



func dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "gateway.networking.k8s.io/v1alpha3", "BackendTLSPolicy", "gateway.networking.k8s.io/v1alpha3/BackendTLSPolicy"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.validation", "status", "status.ancestors.ancestor_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3CompatibleVersions = []string{
	"v1.4.0",
	"v1.4.1",
	"v1.5.0",
	"v1.5.1",
}

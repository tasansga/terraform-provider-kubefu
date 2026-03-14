package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1Read,
		Description: "Gateway represents an instance of a service-traffic handling infrastructure by binding Listeners to a set of IP addresses.",
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
				Description: "Spec defines the desired state of Gateway.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"addresses": {
						Type:        schema.TypeList,
						Description: "Addresses requested for this Gateway. This is optional and behavior can depend on the implementation. If a value is set in the spec and the requested address is invalid or unavailable, the implementation MUST indicate this in the associated entry in GatewayStatus.Addresses. \n The Addresses field represents a request for the address(es) on the \"outside of the Gateway\", that traffic bound for this Gateway will use. This could be the IP address or hostname of an external load balancer or other networking infrastructure, or some other address that traffic will be sent to. \n The .listener.hostname field is used to route traffic that has already arrived at the Gateway to the correct in-cluster destination. \n If no Addresses are specified, the implementation MAY schedule the Gateway in an implementation-specific manner, assigning an appropriate set of Addresses. \n The implementation MUST bind all Listeners to every GatewayAddress that it assigns to the Gateway and add a corresponding entry in GatewayStatus.Addresses. \n Support: Extended",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"type": {
								Type:        schema.TypeString,
								Description: "Type of the address.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"value": {
								Type:        schema.TypeString,
								Description: "Value of the address. The validity of the values will depend on the type and support by the controller. \n Examples: `1.2.3.4`, `128::1`, `my-ip-address`.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"gateway_class_name": {
						Type:        schema.TypeString,
						Description: "GatewayClassName used for this Gateway. This is the name of a GatewayClass resource.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"listeners": {
						Type:        schema.TypeList,
						Description: "Listeners associated with this Gateway. Listeners define logical endpoints that are bound on this Gateway's addresses. At least one Listener MUST be specified. \n Each listener in a Gateway must have a unique combination of Hostname, Port, and Protocol. \n An implementation MAY group Listeners by Port and then collapse each group of Listeners into a single Listener if the implementation determines that the Listeners in the group are \"compatible\". An implementation MAY also group together and collapse compatible Listeners belonging to different Gateways. \n For example, an implementation might consider Listeners to be compatible with each other if all of the following conditions are met: \n 1. Either each Listener within the group specifies the \"HTTP\"    Protocol or each Listener within the group specifies either    the \"HTTPS\" or \"TLS\" Protocol. \n 2. Each Listener within the group specifies a Hostname that is unique    within the group. \n 3. As a special case, one Listener within a group may omit Hostname,    in which case this Listener matches when no other Listener    matches. \n If the implementation does collapse compatible Listeners, the hostname provided in the incoming client request MUST be matched to a Listener to find the correct set of Routes. The incoming hostname MUST be matched using the Hostname field for each Listener in order of most to least specific. That is, exact matches must be processed before wildcard matches. \n If this field specifies multiple Listeners that have the same Port value but are not compatible, the implementation must raise a \"Conflicted\" condition in the Listener status. \n Support: Core",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"allowed_routes": {
								Type:        schema.TypeList,
								Description: "AllowedRoutes defines the types of routes that MAY be attached to a Listener and the trusted namespaces where those Route resources MAY be present. \n Although a client request may match multiple route rules, only one rule may ultimately receive the request. Matching precedence MUST be determined in order of the following criteria: \n * The most specific match as defined by the Route type. * The oldest Route based on creation timestamp. For example, a Route with   a creation timestamp of \"2020-09-08 01:02:03\" is given precedence over   a Route with a creation timestamp of \"2020-09-08 01:02:04\". * If everything else is equivalent, the Route appearing first in   alphabetical order (namespace/name) should be given precedence. For   example, foo/bar is given precedence over foo/baz. \n All valid rules within a Route attached to this Listener should be implemented. Invalid Route rules can be ignored (sometimes that will mean the full Route). If a Route rule transitions from valid to invalid, support for that Route rule should be dropped to ensure consistency. For example, even if a filter specified by a Route rule is invalid, the rest of the rules within that Route should still be supported. \n Support: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"kinds": {
										Type:        schema.TypeList,
										Description: "Kinds specifies the groups and kinds of Routes that are allowed to bind to this Gateway Listener. When unspecified or empty, the kinds of Routes selected are determined using the Listener protocol. \n A RouteGroupKind MUST correspond to kinds of Routes that are compatible with the application protocol specified in the Listener's Protocol field. If an implementation does not support or recognize this resource type, it MUST set the \"ResolvedRefs\" condition to False for this Listener with the \"InvalidRouteKinds\" reason. \n Support: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"group": {
												Type:        schema.TypeString,
												Description: "Group is the group of the Route.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"kind": {
												Type:        schema.TypeString,
												Description: "Kind is the kind of the Route.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"namespaces": {
										Type:        schema.TypeList,
										Description: "Namespaces indicates namespaces from which Routes may be attached to this Listener. This is restricted to the namespace of this Gateway by default. \n Support: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"from": {
												Type:        schema.TypeString,
												Description: "From indicates where Routes will be selected for this Gateway. Possible values are: * All: Routes in all namespaces may be used by this Gateway. * Selector: Routes in namespaces selected by the selector may be used by   this Gateway. * Same: Only Routes in the same namespace may be used by this Gateway. \n Support: Core",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"selector": {
												Type:        schema.TypeList,
												Description: "Selector must be specified when From is set to \"Selector\". In that case, only Routes in Namespaces matching this Selector will be selected by this Gateway. This field is ignored for other values of \"From\". \n Support: Core",
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
																Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Optional:    true,
																Required:    false,
																Computed:    true,
															},
															"values": {
																Type:        schema.TypeList,
																Description: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																Optional:    true,
																Required:    false,
																Computed:    true,
																Elem: &schema.Schema{Type: schema.TypeString},
															},
														}},
													},
													"match_labels": {
														Type:        schema.TypeMap,
														Description: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
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
							"hostname": {
								Type:        schema.TypeString,
								Description: "Hostname specifies the virtual hostname to match for protocol types that define this concept. When unspecified, all hostnames are matched. This field is ignored for protocols that don't require hostname based matching. \n Implementations MUST apply Hostname matching appropriately for each of the following protocols: \n * TLS: The Listener Hostname MUST match the SNI. * HTTP: The Listener Hostname MUST match the Host header of the request. * HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP   protocol layers as described above. If an implementation does not   ensure that both the SNI and Host header match the Listener hostname,   it MUST clearly document that. \n For HTTPRoute and TLSRoute resources, there is an interaction with the `spec.hostnames` array. When both listener and route specify hostnames, there MUST be an intersection between the values for a Route to be accepted. For more information, refer to the Route specific Hostnames documentation. \n Hostnames that are prefixed with a wildcard label (`*.`) are interpreted as a suffix match. That means that a match for `*.example.com` would match both `test.example.com`, and `foo.test.example.com`, but not `example.com`. \n Support: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the Listener. This name MUST be unique within a Gateway. \n Support: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"port": {
								Type:        schema.TypeInt,
								Description: "Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules. \n Support: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"protocol": {
								Type:        schema.TypeString,
								Description: "Protocol specifies the network protocol this listener expects to receive. \n Support: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"tls": {
								Type:        schema.TypeList,
								Description: "TLS is the TLS configuration for the Listener. This field is required if the Protocol field is \"HTTPS\" or \"TLS\". It is invalid to set this field if the Protocol field is \"HTTP\", \"TCP\", or \"UDP\". \n The association of SNIs to Certificate defined in GatewayTLSConfig is defined based on the Hostname field for this listener. \n The GatewayClass MUST use the longest matching SNI out of all available certificates for any TLS handshake. \n Support: Core",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"certificate_refs": {
										Type:        schema.TypeList,
										Description: "CertificateRefs contains a series of references to Kubernetes objects that contains TLS certificates and private keys. These certificates are used to establish a TLS handshake for requests that match the hostname of the associated listener. \n A single CertificateRef to a Kubernetes Secret has \"Core\" support. Implementations MAY choose to support attaching multiple certificates to a Listener, but this behavior is implementation-specific. \n References to a resource in different namespace are invalid UNLESS there is a ReferenceGrant in the target namespace that allows the certificate to be attached. If a ReferenceGrant does not allow this reference, the \"ResolvedRefs\" condition MUST be set to False for this listener with the \"InvalidCertificateRef\" reason. \n This field is required to have at least one element when the mode is set to \"Terminate\" (default) and is optional otherwise. \n CertificateRefs can reference to standard Kubernetes resources, i.e. Secret, or implementation-specific custom resources. \n Support: Core - A single reference to a Kubernetes Secret of type kubernetes.io/tls \n Support: Implementation-specific (More than one reference or other resource types)",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"group": {
												Type:        schema.TypeString,
												Description: "Group is the group of the referent. For example, \"networking.k8s.io\". When unspecified (empty string), core API group is inferred.",
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
											"namespace": {
												Type:        schema.TypeString,
												Description: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred. \n Note that when a namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. \n Support: Core",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"mode": {
										Type:        schema.TypeString,
										Description: "Mode defines the TLS behavior for the TLS session initiated by the client. There are two possible modes: \n - Terminate: The TLS session between the downstream client   and the Gateway is terminated at the Gateway. This mode requires   certificateRefs to be set and contain at least one element. - Passthrough: The TLS session is NOT terminated by the Gateway. This   implies that the Gateway can't decipher the TLS stream except for   the ClientHello message of the TLS protocol.   CertificateRefs field is ignored in this mode. \n Support: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"options": {
										Type:        schema.TypeMap,
										Description: "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites. \n A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as `example.com/my-custom-option`. Un-prefixed names are reserved for key names defined by Gateway API. \n Support: Implementation-specific",
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
				Description: "Status defines the current state of Gateway.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"addresses": {
						Type:        schema.TypeList,
						Description: "Addresses lists the IP addresses that have actually been bound to the Gateway. These addresses may differ from the addresses in the Spec, e.g. if the Gateway automatically assigns an address from a reserved pool.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"type": {
								Type:        schema.TypeString,
								Description: "Type of the address.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"value": {
								Type:        schema.TypeString,
								Description: "Value of the address. The validity of the values will depend on the type and support by the controller. \n Examples: `1.2.3.4`, `128::1`, `my-ip-address`.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"conditions": {
						Type:        schema.TypeList,
						Description: "Conditions describe the current conditions of the Gateway. \n Implementations should prefer to express Gateway conditions using the `GatewayConditionType` and `GatewayConditionReason` constants so that operators and tools can converge on a common vocabulary to describe Gateway state. \n Known condition types are: \n * \"Scheduled\" * \"Ready\"",
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
					"listeners": {
						Type:        schema.TypeList,
						Description: "Listeners provide status for each unique listener port defined in the Spec.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"attached_routes": {
								Type:        schema.TypeInt,
								Description: "AttachedRoutes represents the total number of Routes that have been successfully attached to this Listener.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"conditions": {
								Type:        schema.TypeList,
								Description: "Conditions describe the current condition of this listener.",
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
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the Listener that this status corresponds to.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"supported_kinds": {
								Type:        schema.TypeList,
								Description: "SupportedKinds is the list indicating the Kinds supported by this listener. This MUST represent the kinds an implementation supports for that Listener configuration. \n If kinds are specified in Spec that are not supported, they MUST NOT appear in this list and an implementation MUST set the \"ResolvedRefs\" condition to \"False\" with the \"InvalidRouteKinds\" reason. If both valid and invalid Route kinds are specified, the implementation MUST reference the valid Route kinds that have been specified.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"group": {
										Type:        schema.TypeString,
										Description: "Group is the group of the Route.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"kind": {
										Type:        schema.TypeString,
										Description: "Kind is the kind of the Route.",
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
		},
	}
}



func dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "gateway.networking.k8s.io/v1beta1", "Gateway", "gateway.networking.k8s.io/v1beta1/Gateway"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.listeners.allowed_routes", "spec.listeners.allowed_routes.namespaces", "spec.listeners.allowed_routes.namespaces.selector", "spec.listeners.tls", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1CompatibleVersions = []string{
	"v0.5.0",
	"v0.5.1",
	"v0.6.0",
	"v0.6.1",
	"v0.6.2",
	"v0.7.0",
	"v0.7.1",
	"v0.8.0",
	"v0.8.1",
	"v1.0.0",
	"v1.1.0",
	"v1.1.1",
	"v1.2.0",
	"v1.2.1",
	"v1.3.0",
	"v1.4.0",
	"v1.4.1",
}

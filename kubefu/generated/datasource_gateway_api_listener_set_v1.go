package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceGatewayApiGatewayNetworkingK8sIoListenerSetV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayApiGatewayNetworkingK8sIoListenerSetV1Read,
		Description: "ListenerSet defines a set of additional listeners to attach to an existing Gateway.\nThis resource provides a mechanism to merge multiple listeners into a single Gateway.\n\nThe parent Gateway must explicitly allow ListenerSet attachment through its\nAllowedListeners configuration. By default, Gateways do not allow ListenerSet\nattachment.\n\nRoutes can attach to a ListenerSet by specifying it as a parentRef, and can\noptionally target specific listeners using the sectionName field.\n\nPolicy Attachment:\n- Policies that attach to a ListenerSet apply to all listeners defined in that resource\n- Policies do not impact listeners in the parent Gateway\n- Different ListenerSets attached to the same Gateway can have different policies\n- If an implementation cannot apply a policy to specific listeners, it should reject the policy\n\nReferenceGrant Semantics:\n- ReferenceGrants applied to a Gateway are not inherited by child ListenerSets\n- ReferenceGrants applied to a ListenerSet do not grant permission to the parent Gateway's listeners\n- A ListenerSet can reference secrets/backends in its own namespace without a ReferenceGrant\n\nGateway Integration:\n  - The parent Gateway's status will include \"AttachedListenerSets\"\n    which is the count of ListenerSets that have successfully attached to a Gateway\n    A ListenerSet is successfully attached to a Gateway when all the following conditions are met:\n  - The ListenerSet is selected by the Gateway's AllowedListeners field\n  - The ListenerSet has a valid ParentRef selecting the Gateway\n  - The ListenerSet's status has the condition \"Accepted: true\"",
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
				Description: "Spec defines the desired state of ListenerSet.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"listeners": {
						Type:        schema.TypeList,
						Description: "Listeners associated with this ListenerSet. Listeners define\nlogical endpoints that are bound on this referenced parent Gateway's addresses.\n\nListeners in a `Gateway` and their attached `ListenerSets` are concatenated\nas a list when programming the underlying infrastructure. Each listener\nname does not need to be unique across the Gateway and ListenerSets.\nSee ListenerEntry.Name for more details.\n\nImplementations MUST treat the parent Gateway as having the merged\nlist of all listeners from itself and attached ListenerSets using\nthe following precedence:\n\n1. \"parent\" Gateway\n2. ListenerSet ordered by creation time (oldest first)\n3. ListenerSet ordered alphabetically by \"{namespace}/{name}\".\n\nAn implementation MAY reject listeners by setting the ListenerEntryStatus\n`Accepted` condition to False with the Reason `TooManyListeners`\n\nIf a listener has a conflict, this will be reported in the\nStatus.ListenerEntryStatus setting the `Conflicted` condition to True.\n\nImplementations SHOULD be cautious about what information from the\nparent or siblings are reported to avoid accidentally leaking\nsensitive information that the child would not otherwise have access\nto. This can include contents of secrets etc.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"allowed_routes": {
								Type:        schema.TypeList,
								Description: "AllowedRoutes defines the types of routes that MAY be attached to a\nListener and the trusted namespaces where those Route resources MAY be\npresent.\n\nAlthough a client request may match multiple route rules, only one rule\nmay ultimately receive the request. Matching precedence MUST be\ndetermined in order of the following criteria:\n\n* The most specific match as defined by the Route type.\n* The oldest Route based on creation timestamp. For example, a Route with\n  a creation timestamp of \"2020-09-08 01:02:03\" is given precedence over\n  a Route with a creation timestamp of \"2020-09-08 01:02:04\".\n* If everything else is equivalent, the Route appearing first in\n  alphabetical order (namespace/name) should be given precedence. For\n  example, foo/bar is given precedence over foo/baz.\n\nAll valid rules within a Route attached to this Listener should be\nimplemented. Invalid Route rules can be ignored (sometimes that will mean\nthe full Route). If a Route rule transitions from valid to invalid,\nsupport for that Route rule should be dropped to ensure consistency. For\nexample, even if a filter specified by a Route rule is invalid, the rest\nof the rules within that Route should still be supported.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"kinds": {
										Type:        schema.TypeList,
										Description: "Kinds specifies the groups and kinds of Routes that are allowed to bind\nto this Gateway Listener. When unspecified or empty, the kinds of Routes\nselected are determined using the Listener protocol.\n\nA RouteGroupKind MUST correspond to kinds of Routes that are compatible\nwith the application protocol specified in the Listener's Protocol field.\nIf an implementation does not support or recognize this resource type, it\nMUST set the \"ResolvedRefs\" condition to False for this Listener with the\n\"InvalidRouteKinds\" reason.\n\nSupport: Core",
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
										Description: "Namespaces indicates namespaces from which Routes may be attached to this\nListener. This is restricted to the namespace of this Gateway by default.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"from": {
												Type:        schema.TypeString,
												Description: "From indicates where Routes will be selected for this Gateway. Possible\nvalues are:\n\n* All: Routes in all namespaces may be used by this Gateway.\n* Selector: Routes in namespaces selected by the selector may be used by\n  this Gateway.\n* Same: Only Routes in the same namespace may be used by this Gateway.\n\nSupport: Core",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"selector": {
												Type:        schema.TypeList,
												Description: "Selector must be specified when From is set to \"Selector\". In that case,\nonly Routes in Namespaces matching this Selector will be selected by this\nGateway. This field is ignored for other values of \"From\".\n\nSupport: Core",
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
										}},
									},
								}},
							},
							"hostname": {
								Type:        schema.TypeString,
								Description: "Hostname specifies the virtual hostname to match for protocol types that\ndefine this concept. When unspecified, all hostnames are matched. This\nfield is ignored for protocols that don't require hostname based\nmatching.\n\nImplementations MUST apply Hostname matching appropriately for each of\nthe following protocols:\n\n* TLS: The Listener Hostname MUST match the SNI.\n* HTTP: The Listener Hostname MUST match the Host header of the request.\n* HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP\n  protocol layers as described above. If an implementation does not\n  ensure that both the SNI and Host header match the Listener hostname,\n  it MUST clearly document that.\n\nFor HTTPRoute and TLSRoute resources, there is an interaction with the\n`spec.hostnames` array. When both listener and route specify hostnames,\nthere MUST be an intersection between the values for a Route to be\naccepted. For more information, refer to the Route specific Hostnames\ndocumentation.\n\nHostnames that are prefixed with a wildcard label (`*.`) are interpreted\nas a suffix match. That means that a match for `*.example.com` would match\nboth `test.example.com`, and `foo.test.example.com`, but not `example.com`.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the Listener. This name MUST be unique within a\nListenerSet.\n\nName is not required to be unique across a Gateway and ListenerSets.\nRoutes can attach to a Listener by having a ListenerSet as a parentRef\nand setting the SectionName",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"port": {
								Type:        schema.TypeInt,
								Description: "Port is the network port. Multiple listeners may use the\nsame port, subject to the Listener compatibility rules.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"protocol": {
								Type:        schema.TypeString,
								Description: "Protocol specifies the network protocol this listener expects to receive.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"tls": {
								Type:        schema.TypeList,
								Description: "TLS is the TLS configuration for the Listener. This field is required if\nthe Protocol field is \"HTTPS\" or \"TLS\". It is invalid to set this field\nif the Protocol field is \"HTTP\", \"TCP\", or \"UDP\".\n\nThe association of SNIs to Certificate defined in ListenerTLSConfig is\ndefined based on the Hostname field for this listener.\n\nThe GatewayClass MUST use the longest matching SNI out of all\navailable certificates for any TLS handshake.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"certificate_refs": {
										Type:        schema.TypeList,
										Description: "CertificateRefs contains a series of references to Kubernetes objects that\ncontains TLS certificates and private keys. These certificates are used to\nestablish a TLS handshake for requests that match the hostname of the\nassociated listener.\n\nA single CertificateRef to a Kubernetes Secret has \"Core\" support.\nImplementations MAY choose to support attaching multiple certificates to\na Listener, but this behavior is implementation-specific.\n\nReferences to a resource in different namespace are invalid UNLESS there\nis a ReferenceGrant in the target namespace that allows the certificate\nto be attached. If a ReferenceGrant does not allow this reference, the\n\"ResolvedRefs\" condition MUST be set to False for this listener with the\n\"RefNotPermitted\" reason.\n\nThis field is required to have at least one element when the mode is set\nto \"Terminate\" (default) and is optional otherwise.\n\nCertificateRefs can reference to standard Kubernetes resources, i.e.\nSecret, or implementation-specific custom resources.\n\nSupport: Core - A single reference to a Kubernetes Secret of type kubernetes.io/tls\n\nSupport: Implementation-specific (More than one reference or other resource types)",
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
												Description: "Kind is kind of the referent. For example \"Secret\".",
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
												Description: "Namespace is the namespace of the referenced object. When unspecified, the local\nnamespace is inferred.\n\nNote that when a namespace different than the local namespace is specified,\na ReferenceGrant object is required in the referent namespace to allow that\nnamespace's owner to accept the reference. See the ReferenceGrant\ndocumentation for details.\n\nSupport: Core",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"mode": {
										Type:        schema.TypeString,
										Description: "Mode defines the TLS behavior for the TLS session initiated by the client.\nThere are two possible modes:\n\n- Terminate: The TLS session between the downstream client and the\n  Gateway is terminated at the Gateway. This mode requires certificates\n  to be specified in some way, such as populating the certificateRefs\n  field.\n- Passthrough: The TLS session is NOT terminated by the Gateway. This\n  implies that the Gateway can't decipher the TLS stream except for\n  the ClientHello message of the TLS protocol. The certificateRefs field\n  is ignored in this mode.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"options": {
										Type:        schema.TypeMap,
										Description: "Options are a list of key/value pairs to enable extended TLS\nconfiguration for each implementation. For example, configuring the\nminimum TLS version or supported cipher suites.\n\nA set of common keys MAY be defined by the API in the future. To avoid\nany ambiguity, implementation-specific definitions MUST use\ndomain-prefixed names, such as `example.com/my-custom-option`.\nUn-prefixed names are reserved for key names defined by Gateway API.\n\nSupport: Implementation-specific",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"parent_ref": {
						Type:        schema.TypeList,
						Description: "ParentRef references the Gateway that the listeners are attached to.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"group": {
								Type:        schema.TypeString,
								Description: "Group is the group of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind is kind of the referent. For example \"Gateway\".",
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
								Description: "Namespace is the namespace of the referent.  If not present,\nthe namespace of the referent is assumed to be the same as\nthe namespace of the referring object.",
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
				Description: "Status defines the current state of ListenerSet.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"conditions": {
						Type:        schema.TypeList,
						Description: "Conditions describe the current conditions of the ListenerSet.\n\nImplementations MUST express ListenerSet conditions using the\n`ListenerSetConditionType` and `ListenerSetConditionReason`\nconstants so that operators and tools can converge on a common\nvocabulary to describe ListenerSet state.\n\nKnown condition types are:\n\n* \"Accepted\"\n* \"Programmed\"",
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
					"listeners": {
						Type:        schema.TypeList,
						Description: "Listeners provide status for each unique listener port defined in the Spec.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"attached_routes": {
								Type:        schema.TypeInt,
								Description: "AttachedRoutes represents the total number of Routes that have been\nsuccessfully attached to this Listener.\n\nSuccessful attachment of a Route to a Listener is based solely on the\ncombination of the AllowedRoutes field on the corresponding Listener\nand the Route's ParentRefs field. A Route is successfully attached to\na Listener when it is selected by the Listener's AllowedRoutes field\nAND the Route has a valid ParentRef selecting the whole Gateway\nresource or a specific Listener as a parent resource (more detail on\nattachment semantics can be found in the documentation on the various\nRoute kinds ParentRefs fields). Listener status does not impact\nsuccessful attachment, i.e. the AttachedRoutes field count MUST be set\nfor Listeners, even if the Accepted condition of an individual Listener is set\nto \"False\". The AttachedRoutes number represents the number of Routes with\nthe Accepted condition set to \"True\" that have been attached to this Listener.\nRoutes with any other value for the Accepted condition MUST NOT be included\nin this count.\n\nUses for this field include troubleshooting Route attachment and\nmeasuring blast radius/impact of changes to a Listener.",
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
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the Listener that this status corresponds to.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"supported_kinds": {
								Type:        schema.TypeList,
								Description: "SupportedKinds is the list indicating the Kinds supported by this\nlistener. This MUST represent the kinds supported by an implementation for\nthat Listener configuration.\n\nIf kinds are specified in Spec that are not supported, they MUST NOT\nappear in this list and an implementation MUST set the \"ResolvedRefs\"\ncondition to \"False\" with the \"InvalidRouteKinds\" reason. If both valid\nand invalid Route kinds are specified, the implementation MUST\nreference the valid Route kinds that have been specified.",
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



func dataSourceGatewayApiGatewayNetworkingK8sIoListenerSetV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "gateway.networking.k8s.io/v1", "ListenerSet", "gateway.networking.k8s.io/v1/ListenerSet"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.listeners.allowed_routes", "spec.listeners.allowed_routes.namespaces", "spec.listeners.allowed_routes.namespaces.selector", "spec.listeners.tls", "spec.parent_ref", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceGatewayApiGatewayNetworkingK8sIoListenerSetV1CompatibleVersions = []string{
	"v1.5.0",
	"v1.5.1",
}

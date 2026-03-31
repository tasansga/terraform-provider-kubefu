package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceGatewayApiGatewayNetworkingK8sIoTLSRouteV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayApiGatewayNetworkingK8sIoTLSRouteV1Alpha2Read,
		Description: "The TLSRoute resource is similar to TCPRoute, but can be configured\nto match against TLS-specific metadata. This allows more flexibility\nin matching streams for a given TLS listener.",
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
				Description: "Spec defines the desired state of TLSRoute.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"hostnames": {
						Type:        schema.TypeList,
						Description: "Hostnames defines a set of SNI names that should match against the\nSNI attribute of TLS ClientHello message in TLS handshake. This matches\nthe RFC 1123 definition of a hostname with 2 notable exceptions:\n\n1. IPs are not allowed in SNI names per RFC 6066.\n2. A hostname may be prefixed with a wildcard label (`*.`). The wildcard\n   label must appear by itself as the first label.\n\nIf a hostname is specified by both the Listener and TLSRoute, there\nmust be at least one intersecting hostname for the TLSRoute to be\nattached to the Listener. For example:\n\n* A Listener with `test.example.com` as the hostname matches TLSRoutes\n  that have either not specified any hostnames, or have specified at\n  least one of `test.example.com` or `*.example.com`.\n* A Listener with `*.example.com` as the hostname matches TLSRoutes\n  that have either not specified any hostnames or have specified at least\n  one hostname that matches the Listener hostname. For example,\n  `test.example.com` and `*.example.com` would both match. On the other\n  hand, `example.com` and `test.example.net` would not match.\n\nIf both the Listener and TLSRoute have specified hostnames, any\nTLSRoute hostnames that do not match the Listener hostname MUST be\nignored. For example, if a Listener specified `*.example.com`, and the\nTLSRoute specified `test.example.com` and `test.example.net`,\n`test.example.net` must not be considered for a match.\n\nIf both the Listener and TLSRoute have specified hostnames, and none\nmatch with the criteria above, then the TLSRoute is not accepted. The\nimplementation must raise an 'Accepted' Condition with a status of\n`False` in the corresponding RouteParentStatus.\n\nSupport: Core",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"parent_refs": {
						Type:        schema.TypeList,
						Description: "ParentRefs references the resources (usually Gateways) that a Route wants\nto be attached to. Note that the referenced parent resource needs to\nallow this for the attachment to be complete. For Gateways, that means\nthe Gateway needs to allow attachment from Routes of this kind and\nnamespace. For Services, that means the Service must either be in the same\nnamespace for a \"producer\" route, or the mesh implementation must support\nand allow \"consumer\" routes for the referenced Service. ReferenceGrant is\nnot applicable for governing ParentRefs to Services - it is not possible to\ncreate a \"producer\" route for a Service in a different namespace from the\nRoute.\n\nThere are two kinds of parent resources with \"Core\" support:\n\n* Gateway (Gateway conformance profile)\n* Service (Mesh conformance profile, ClusterIP Services only)\n\nThis API may be extended in the future to support additional kinds of parent\nresources.\n\nParentRefs must be _distinct_. This means either that:\n\n* They select different objects.  If this is the case, then parentRef\n  entries are distinct. In terms of fields, this means that the\n  multi-part key defined by `group`, `kind`, `namespace`, and `name` must\n  be unique across all parentRef entries in the Route.\n* They do not select different objects, but for each optional field used,\n  each ParentRef that selects the same object must set the same set of\n  optional fields to different values. If one ParentRef sets a\n  combination of optional fields, all must set the same combination.\n\nSome examples:\n\n* If one ParentRef sets `sectionName`, all ParentRefs referencing the\n  same object must also set `sectionName`.\n* If one ParentRef sets `port`, all ParentRefs referencing the same\n  object must also set `port`.\n* If one ParentRef sets `sectionName` and `port`, all ParentRefs\n  referencing the same object must also set `sectionName` and `port`.\n\nIt is possible to separately reference multiple distinct objects that may\nbe collapsed by an implementation. For example, some implementations may\nchoose to merge compatible Gateway Listeners together. If that is the\ncase, the list of routes attached to those resources should also be\nmerged.\n\nNote that for ParentRefs that cross namespace boundaries, there are specific\nrules. Cross-namespace references are only valid if they are explicitly\nallowed by something in the namespace they are referring to. For example,\nGateway has the AllowedRoutes field, and ReferenceGrant provides a\ngeneric way to enable other kinds of cross-namespace reference.",
						Optional:    true,
						Required:    false,
						Computed:    true,
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
					"rules": {
						Type:        schema.TypeList,
						Description: "Rules are a list of TLS matchers and actions.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"backend_refs": {
								Type:        schema.TypeList,
								Description: "BackendRefs defines the backend(s) where matching requests should be\nsent. If unspecified or invalid (refers to a nonexistent resource or\na Service with no endpoints), the rule performs no forwarding; if no\nfilters are specified that would result in a response being sent, the\nunderlying implementation must actively reject request attempts to this\nbackend, by rejecting the connection. Request rejections must respect\nweight; if an invalid backend is requested to have 80% of requests, then\n80% of requests must be rejected instead.\n\nSupport: Core for Kubernetes Service\n\nSupport: Extended for Kubernetes ServiceImport\n\nSupport: Implementation-specific for any other resource\n\nSupport for weight: Extended",
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
										Description: "Kind is the Kubernetes resource kind of the referent. For example\n\"Service\".\n\nDefaults to \"Service\" when not specified.\n\nExternalName services can refer to CNAME DNS records that may live\noutside of the cluster and as such are difficult to reason about in\nterms of conformance. They also may not be safe to forward to (see\nCVE-2021-25740 for more information). Implementations SHOULD NOT\nsupport ExternalName Services.\n\nSupport: Core (Services with a type other than ExternalName)\n\nSupport: Implementation-specific (Services with type ExternalName)",
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
										Description: "Namespace is the namespace of the backend. When unspecified, the local\nnamespace is inferred.\n\nNote that when a namespace different than the local namespace is specified,\na ReferenceGrant object is required in the referent namespace to allow that\nnamespace's owner to accept the reference. See the ReferenceGrant\ndocumentation for details.\n\nSupport: Core",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeInt,
										Description: "Port specifies the destination port number to use for this resource.\nPort is required when the referent is a Kubernetes Service. In this\ncase, the port number is the service port number, not the target port.\nFor other resources, destination port might be derived from the referent\nresource or this field.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Weight specifies the proportion of requests forwarded to the referenced\nbackend. This is computed as weight/(sum of all weights in this\nBackendRefs list). For non-zero values, there may be some epsilon from\nthe exact proportion defined here depending on the precision an\nimplementation supports. Weight is not a percentage and the sum of\nweights does not need to equal 100.\n\nIf only one backend is specified and it has a weight greater than 0, 100%\nof the traffic is forwarded to that backend. If weight is set to 0, no\ntraffic should be forwarded for this entry. If unspecified, weight\ndefaults to 1.\n\nSupport for this field varies based on the context where used.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of the route rule. This name MUST be unique within a Route if it is set.",
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
				Description: "Status defines the current state of TLSRoute.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"parents": {
						Type:        schema.TypeList,
						Description: "Parents is a list of parent resources (usually Gateways) that are\nassociated with the route, and the status of the route with respect to\neach parent. When this route attaches to a parent, the controller that\nmanages the parent must add an entry to this list when the controller\nfirst sees the route and should update the entry as appropriate when the\nroute or gateway is modified.\n\nNote that parent references that cannot be resolved by an implementation\nof this API will not be added to this list. Implementations of this API\ncan only populate Route status for the Gateways/parent resources they are\nresponsible for.\n\nA maximum of 32 Gateways will be represented in this list. An empty list\nmeans the route has not been attached to any Gateway.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"conditions": {
								Type:        schema.TypeList,
								Description: "Conditions describes the status of the route with respect to the Gateway.\nNote that the route's availability is also subject to the Gateway's own\nstatus conditions and listener status.\n\nIf the Route's ParentRef specifies an existing Gateway that supports\nRoutes of this kind AND that Gateway's controller has sufficient access,\nthen that Gateway's controller MUST set the \"Accepted\" condition on the\nRoute, to indicate whether the route has been accepted or rejected by the\nGateway, and why.\n\nA Route MUST be considered \"Accepted\" if at least one of the Route's\nrules is implemented by the Gateway.\n\nThere are a number of cases where the \"Accepted\" condition may not be set\ndue to lack of controller visibility, that includes when:\n\n* The Route refers to a nonexistent parent.\n* The Route is of a type that the controller does not support.\n* The Route is in a namespace to which the controller does not have access.",
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
							"parent_ref": {
								Type:        schema.TypeList,
								Description: "ParentRef corresponds with a ParentRef in the spec that this\nRouteParentStatus struct describes the status of.",
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
						}},
					},
				}},
			},
		},
	}
}



func dataSourceGatewayApiGatewayNetworkingK8sIoTLSRouteV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "gateway.networking.k8s.io/v1alpha2", "TLSRoute", "gateway.networking.k8s.io/v1alpha2/TLSRoute"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "status", "status.parents.parent_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceGatewayApiGatewayNetworkingK8sIoTLSRouteV1Alpha2CompatibleVersions = []string{
	"v1.5.0",
	"v1.5.1",
}

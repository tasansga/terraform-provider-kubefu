package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1Read,
		Description: "FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of inbound API requests with similar attributes and is identified by a pair of strings: the name of the FlowSchema and a \"flow distinguisher\".",
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
				Type:        schema.TypeList,
				Description: "`metadata` is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"annotations": {
						Type:        schema.TypeMap,
						Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"cluster_name": {
						Type:        schema.TypeString,
						Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"creation_timestamp": {
						Type:        schema.TypeString,
						Description: "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"deletion_grace_period_seconds": {
						Type:        schema.TypeInt,
						Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"deletion_timestamp": {
						Type:        schema.TypeString,
						Description: "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"finalizers": {
						Type:        schema.TypeList,
						Description: "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order.  Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"generate_name": {
						Type:        schema.TypeString,
						Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"generation": {
						Type:        schema.TypeInt,
						Description: "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"labels": {
						Type:        schema.TypeMap,
						Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"managed_fields": {
						Type:        schema.TypeList,
						Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_version": {
								Type:        schema.TypeString,
								Description: "APIVersion defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"fields_type": {
								Type:        schema.TypeString,
								Description: "FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"fields_v1": {
								Type:        schema.TypeMap,
								Description: "FieldsV1 holds the first JSON version format as described in the \"FieldsV1\" type.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"manager": {
								Type:        schema.TypeString,
								Description: "Manager is an identifier of the workflow managing these fields.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"operation": {
								Type:        schema.TypeString,
								Description: "Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"subresource": {
								Type:        schema.TypeString,
								Description: "Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"time": {
								Type:        schema.TypeString,
								Description: "Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"name": {
						Type:        schema.TypeString,
						Description: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"namespace": {
						Type:        schema.TypeString,
						Description: "Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"owner_references": {
						Type:        schema.TypeList,
						Description: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_version": {
								Type:        schema.TypeString,
								Description: "API version of the referent.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"block_owner_deletion": {
								Type:        schema.TypeBool,
								Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"controller": {
								Type:        schema.TypeBool,
								Description: "If true, this reference points to the managing controller.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"uid": {
								Type:        schema.TypeString,
								Description: "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
					"resource_version": {
						Type:        schema.TypeString,
						Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"self_link": {
						Type:        schema.TypeString,
						Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.\n\nDEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"uid": {
						Type:        schema.TypeString,
						Description: "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"spec": {
				Type:        schema.TypeList,
				Description: "`spec` is the specification of the desired behavior of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"distinguisher_method": {
						Type:        schema.TypeList,
						Description: "`distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies that the distinguisher is disabled and thus will always be the empty string.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"type": {
								Type:        schema.TypeString,
								Description: "`type` is the type of flow distinguisher method The supported types are \"ByUser\" and \"ByNamespace\". Required.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
					"matching_precedence": {
						Type:        schema.TypeInt,
						Description: "`matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"priority_level_configuration": {
						Type:        schema.TypeList,
						Description: "`priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.",
						Optional:    false,
						Required:    true,
						Computed:    false,
						MinItems:    1,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "`name` is the name of the priority level configuration being referenced Required.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
					"rules": {
						Type:        schema.TypeList,
						Description: "`rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"non_resource_rules": {
								Type:        schema.TypeList,
								Description: "`nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"non_resource_ur_ls": {
										Type:        schema.TypeList,
										Description: "`nonResourceURLs` is a set of url prefixes that a user should have access to and may not be empty. For example:\n  - \"/healthz\" is legal\n  - \"/hea*\" is illegal\n  - \"/hea\" is legal but matches nothing\n  - \"/hea/*\" also matches nothing\n  - \"/healthz/*\" matches all per-component health checks.\n\"*\" matches all non-resource urls. if it is present, it must be the only entry. Required.",
										Optional:    false,
										Required:    true,
										Computed:    false,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"verbs": {
										Type:        schema.TypeList,
										Description: "`verbs` is a list of matching verbs and may not be empty. \"*\" matches all verbs. If it is present, it must be the only entry. Required.",
										Optional:    false,
										Required:    true,
										Computed:    false,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
								}},
							},
							"resource_rules": {
								Type:        schema.TypeList,
								Description: "`resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of `resourceRules` and `nonResourceRules` has to be non-empty.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"api_groups": {
										Type:        schema.TypeList,
										Description: "`apiGroups` is a list of matching API groups and may not be empty. \"*\" matches all API groups and, if present, must be the only entry. Required.",
										Optional:    false,
										Required:    true,
										Computed:    false,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"cluster_scope": {
										Type:        schema.TypeBool,
										Description: "`clusterScope` indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the `namespaces` field must contain a non-empty list.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"namespaces": {
										Type:        schema.TypeList,
										Description: "`namespaces` is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains \"*\".  Note that \"*\" matches any specified namespace but does not match a request that _does not specify_ a namespace (see the `clusterScope` field for that). This list may be empty, but only if `clusterScope` is true.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"resources": {
										Type:        schema.TypeList,
										Description: "`resources` is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ \"services\", \"nodes/status\" ].  This list may not be empty. \"*\" matches all resources and, if present, must be the only entry. Required.",
										Optional:    false,
										Required:    true,
										Computed:    false,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"verbs": {
										Type:        schema.TypeList,
										Description: "`verbs` is a list of matching verbs and may not be empty. \"*\" matches all verbs and, if present, must be the only entry. Required.",
										Optional:    false,
										Required:    true,
										Computed:    false,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
								}},
							},
							"subjects": {
								Type:        schema.TypeList,
								Description: "subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.",
								Optional:    false,
								Required:    true,
								Computed:    false,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"group": {
										Type:        schema.TypeList,
										Description: "`group` matches based on user group name.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Description: "name is the user group that matches, or \"*\" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.",
												Optional:    false,
												Required:    true,
												Computed:    false,
											},
										}},
									},
									"kind": {
										Type:        schema.TypeString,
										Description: "Required",
										Optional:    false,
										Required:    true,
										Computed:    false,
									},
									"service_account": {
										Type:        schema.TypeList,
										Description: "`serviceAccount` matches ServiceAccounts.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Description: "`name` is the name of matching ServiceAccount objects, or \"*\" to match regardless of name. Required.",
												Optional:    false,
												Required:    true,
												Computed:    false,
											},
											"namespace": {
												Type:        schema.TypeString,
												Description: "`namespace` is the namespace of matching ServiceAccount objects. Required.",
												Optional:    false,
												Required:    true,
												Computed:    false,
											},
										}},
									},
									"user": {
										Type:        schema.TypeList,
										Description: "`user` matches based on username.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Description: "`name` is the username that matches, or \"*\" to match all usernames. Required.",
												Optional:    false,
												Required:    true,
												Computed:    false,
											},
										}},
									},
								}},
							},
						}},
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "`status` is the current status of a FlowSchema. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"conditions": {
						Type:        schema.TypeList,
						Description: "`conditions` is a list of the current states of FlowSchema.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"last_transition_time": {
								Type:        schema.TypeString,
								Description: "`lastTransitionTime` is the last time the condition transitioned from one status to another.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"message": {
								Type:        schema.TypeString,
								Description: "`message` is a human-readable message indicating details about last transition.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"reason": {
								Type:        schema.TypeString,
								Description: "`reason` is a unique, one-word, CamelCase reason for the condition's last transition.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"status": {
								Type:        schema.TypeString,
								Description: "`status` is the status of the condition. Can be True, False, Unknown. Required.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"type": {
								Type:        schema.TypeString,
								Description: "`type` is the type of the condition. Required.",
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



func dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "flowcontrol.apiserver.k8s.io/v1beta1", "FlowSchema", "flowcontrol.apiserver.k8s.io/v1beta1/FlowSchema"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"metadata", "spec", "spec.distinguisher_method", "spec.priority_level_configuration", "spec.rules.subjects.group", "spec.rules.subjects.service_account", "spec.rules.subjects.user", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1CompatibleVersions = []string{
	"v1.20.0",
	"v1.21.0",
	"v1.22.0",
	"v1.23.0",
	"v1.24.0",
	"v1.25.0",
}

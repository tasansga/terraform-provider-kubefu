package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sCoreConfigMapV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sCoreConfigMapV1Read,
		Description: "ConfigMap holds configuration data for pods to consume.",
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"binary_data": {
				Type:        schema.TypeMap,
				Description: "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"data": {
				Type:        schema.TypeMap,
				Description: "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
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
				Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
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
						Description: "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
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
						Description: "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"finalizers": {
						Type:        schema.TypeList,
						Description: "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
					"generate_name": {
						Type:        schema.TypeString,
						Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency",
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
					"initializers": {
						Type:        schema.TypeList,
						Description: "An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.\n\nWhen an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"pending": {
								Type:        schema.TypeList,
								Description: "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
								Optional:    false,
								Required:    true,
								Computed:    false,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "name of the process that is responsible for initializing this object.",
										Optional:    false,
										Required:    true,
										Computed:    false,
									},
								}},
							},
							"result": {
								Type:        schema.TypeList,
								Description: "If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"api_version": {
										Type:        schema.TypeString,
										Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"code": {
										Type:        schema.TypeInt,
										Description: "Suggested HTTP return code for this status, 0 if not set.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"details": {
										Type:        schema.TypeList,
										Description: "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"causes": {
												Type:        schema.TypeList,
												Description: "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
												Optional:    true,
												Required:    false,
												Computed:    true,
												Elem: &schema.Resource{Schema: map[string]*schema.Schema{
													"field": {
														Type:        schema.TypeString,
														Description: "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
														Optional:    true,
														Required:    false,
														Computed:    true,
													},
													"message": {
														Type:        schema.TypeString,
														Description: "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
														Optional:    true,
														Required:    false,
														Computed:    true,
													},
													"reason": {
														Type:        schema.TypeString,
														Description: "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
														Optional:    true,
														Required:    false,
														Computed:    true,
													},
												}},
											},
											"group": {
												Type:        schema.TypeString,
												Description: "The group attribute of the resource associated with the status StatusReason.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"kind": {
												Type:        schema.TypeString,
												Description: "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"retry_after_seconds": {
												Type:        schema.TypeInt,
												Description: "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"uid": {
												Type:        schema.TypeString,
												Description: "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"kind": {
										Type:        schema.TypeString,
										Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"message": {
										Type:        schema.TypeString,
										Description: "A human-readable description of the status of this operation.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"metadata": {
										Type:        schema.TypeList,
										Description: "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"continue": {
												Type:        schema.TypeString,
												Description: "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"resource_version": {
												Type:        schema.TypeString,
												Description: "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"self_link": {
												Type:        schema.TypeString,
												Description: "selfLink is a URL representing this object. Populated by the system. Read-only.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"reason": {
										Type:        schema.TypeString,
										Description: "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
						}},
					},
					"labels": {
						Type:        schema.TypeMap,
						Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
						Optional:    true,
						Required:    false,
						Computed:    true,
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
						Description: "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
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
								Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
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
						Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"self_link": {
						Type:        schema.TypeString,
						Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.",
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
		},
	}
}



func dataSourceK8sCoreConfigMapV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "v1", "ConfigMap", "core/v1/ConfigMap"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"binary_data", "data", "metadata"}, []string{"metadata", "metadata.initializers", "metadata.initializers.result", "metadata.initializers.result.details", "metadata.initializers.result.metadata"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sCoreConfigMapV1CompatibleVersions = []string{
	"v1.7.0",
	"v1.8.0",
	"v1.9.0",
	"v1.10.0",
	"v1.11.0",
	"v1.12.0",
	"v1.13.0",
	"v1.14.0",
	"v1.15.0",
	"v1.16.0",
	"v1.17.0",
	"v1.18.0",
	"v1.19.0",
	"v1.20.0",
	"v1.21.0",
	"v1.22.0",
	"v1.23.0",
	"v1.24.0",
	"v1.25.0",
	"v1.26.0",
	"v1.27.0",
	"v1.28.0",
	"v1.29.0",
	"v1.30.0",
	"v1.31.0",
	"v1.32.0",
	"v1.33.0",
	"v1.34.0",
	"v1.35.0",
}

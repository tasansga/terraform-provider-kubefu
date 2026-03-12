package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sResourceK8sIoResourceClaimV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sResourceK8sIoResourceClaimV1Alpha2Read,
		Description: "ResourceClaim describes which resources are needed by a resource consumer. Its status tracks whether the resource has been allocated and what the resulting attributes are.\n\nThis is an alpha type and requires enabling the DynamicResourceAllocation feature gate.",
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
				Description: "Standard object metadata",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"annotations": {
						Type:        schema.TypeMap,
						Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
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
						Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will return a 409.\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
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
						Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
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
								Description: "Time is the timestamp of when the ManagedFields entry was added. The timestamp will also be updated if a field is added, the manager changes any of the owned fields value or removes a field. The timestamp does not update when a field is removed from the entry because another manager took it over.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"name": {
						Type:        schema.TypeString,
						Description: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"namespace": {
						Type:        schema.TypeString,
						Description: "Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces",
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
								Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
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
								Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"uid": {
								Type:        schema.TypeString,
								Description: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
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
						Description: "Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"uid": {
						Type:        schema.TypeString,
						Description: "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"spec": {
				Type:        schema.TypeList,
				Description: "Spec describes the desired attributes of a resource that then needs to be allocated. It can only be set once when creating the ResourceClaim.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allocation_mode": {
						Type:        schema.TypeString,
						Description: "Allocation can start immediately or when a Pod wants to use the resource. \"WaitForFirstConsumer\" is the default.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"parameters_ref": {
						Type:        schema.TypeList,
						Description: "ParametersRef references a separate object with arbitrary parameters that will be used by the driver when allocating a resource for the claim.\n\nThe object must be in the same namespace as the ResourceClaim.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_group": {
								Type:        schema.TypeString,
								Description: "APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata, for example \"ConfigMap\".",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of resource being referenced.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
					"resource_class_name": {
						Type:        schema.TypeString,
						Description: "ResourceClassName references the driver and additional parameters via the name of a ResourceClass that was created as part of the driver deployment.",
						Optional:    false,
						Required:    true,
						Computed:    false,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "Status describes whether the resource is available and with which attributes.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allocation": {
						Type:        schema.TypeList,
						Description: "Allocation is set by the resource driver once a resource or set of resources has been allocated successfully. If this is not specified, the resources have not been allocated yet.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"available_on_nodes": {
								Type:        schema.TypeList,
								Description: "This field will get set by the resource driver after it has allocated the resource to inform the scheduler where it can schedule Pods using the ResourceClaim.\n\nSetting this field is optional. If null, the resource is available everywhere.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"node_selector_terms": {
										Type:        schema.TypeList,
										Description: "Required. A list of node selector terms. The terms are ORed.",
										Optional:    false,
										Required:    true,
										Computed:    false,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"match_expressions": {
												Type:        schema.TypeList,
												Description: "A list of node selector requirements by node's labels.",
												Optional:    true,
												Required:    false,
												Computed:    true,
												Elem: &schema.Resource{Schema: map[string]*schema.Schema{
													"key": {
														Type:        schema.TypeString,
														Description: "The label key that the selector applies to.",
														Optional:    false,
														Required:    true,
														Computed:    false,
													},
													"operator": {
														Type:        schema.TypeString,
														Description: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														Optional:    false,
														Required:    true,
														Computed:    false,
													},
													"values": {
														Type:        schema.TypeList,
														Description: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														Optional:    true,
														Required:    false,
														Computed:    true,
														Elem: &schema.Schema{Type: schema.TypeString},
													},
												}},
											},
											"match_fields": {
												Type:        schema.TypeList,
												Description: "A list of node selector requirements by node's fields.",
												Optional:    true,
												Required:    false,
												Computed:    true,
												Elem: &schema.Resource{Schema: map[string]*schema.Schema{
													"key": {
														Type:        schema.TypeString,
														Description: "The label key that the selector applies to.",
														Optional:    false,
														Required:    true,
														Computed:    false,
													},
													"operator": {
														Type:        schema.TypeString,
														Description: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														Optional:    false,
														Required:    true,
														Computed:    false,
													},
													"values": {
														Type:        schema.TypeList,
														Description: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														Optional:    true,
														Required:    false,
														Computed:    true,
														Elem: &schema.Schema{Type: schema.TypeString},
													},
												}},
											},
										}},
									},
								}},
							},
							"resource_handles": {
								Type:        schema.TypeList,
								Description: "ResourceHandles contain the state associated with an allocation that should be maintained throughout the lifetime of a claim. Each ResourceHandle contains data that should be passed to a specific kubelet plugin once it lands on a node. This data is returned by the driver after a successful allocation and is opaque to Kubernetes. Driver documentation may explain to users how to interpret this data if needed.\n\nSetting this field is optional. It has a maximum size of 32 entries. If null (or empty), it is assumed this allocation will be processed by a single kubelet plugin with no ResourceHandle data attached. The name of the kubelet plugin invoked will match the DriverName set in the ResourceClaimStatus this AllocationResult is embedded in.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"data": {
										Type:        schema.TypeString,
										Description: "Data contains the opaque data associated with this ResourceHandle. It is set by the controller component of the resource driver whose name matches the DriverName set in the ResourceClaimStatus this ResourceHandle is embedded in. It is set at allocation time and is intended for processing by the kubelet plugin whose name matches the DriverName set in this ResourceHandle.\n\nThe maximum size of this field is 16KiB. This may get increased in the future, but not reduced.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"driver_name": {
										Type:        schema.TypeString,
										Description: "DriverName specifies the name of the resource driver whose kubelet plugin should be invoked to process this ResourceHandle's data once it lands on a node. This may differ from the DriverName set in ResourceClaimStatus this ResourceHandle is embedded in.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"shareable": {
								Type:        schema.TypeBool,
								Description: "Shareable determines whether the resource supports more than one consumer at a time.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"deallocation_requested": {
						Type:        schema.TypeBool,
						Description: "DeallocationRequested indicates that a ResourceClaim is to be deallocated.\n\nThe driver then must deallocate this claim and reset the field together with clearing the Allocation field.\n\nWhile DeallocationRequested is set, no new consumers may be added to ReservedFor.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"driver_name": {
						Type:        schema.TypeString,
						Description: "DriverName is a copy of the driver name from the ResourceClass at the time when allocation started.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"reserved_for": {
						Type:        schema.TypeList,
						Description: "ReservedFor indicates which entities are currently allowed to use the claim. A Pod which references a ResourceClaim which is not reserved for that Pod will not be started.\n\nThere can be at most 32 such reservations. This may get increased in the future, but not reduced.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_group": {
								Type:        schema.TypeString,
								Description: "APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of resource being referenced.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"resource": {
								Type:        schema.TypeString,
								Description: "Resource is the type of resource being referenced, for example \"pods\".",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"uid": {
								Type:        schema.TypeString,
								Description: "UID identifies exactly one incarnation of the resource.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
				}},
			},
		},
	}
}



func dataSourceK8sResourceK8sIoResourceClaimV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "resource.k8s.io/v1alpha2", "ResourceClaim", "resource.k8s.io/v1alpha2/ResourceClaim"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"metadata", "spec", "spec.parameters_ref", "status", "status.allocation", "status.allocation.available_on_nodes"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sResourceK8sIoResourceClaimV1Alpha2CompatibleVersions = []string{
	"v1.27.0",
	"v1.28.0",
	"v1.29.0",
	"v1.30.0",
}

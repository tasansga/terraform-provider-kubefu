package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sLifecycleK8sIoEvictionRequestV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sLifecycleK8sIoEvictionRequestV1Alpha1Read,
		Description: "EvictionRequest defines a request that should ideally result in a graceful eviction of a .spec.target (e.g. termination of a pod).\n\nThe evictionrequest-controller observes intents of all EvictionRequests and transforms them into Evictions.\n  - .spec.requester is set as a label on the Eviction for easier lookup.\n  - Each target can have a set of responders assigned to it. Eviction objects are observed by\n    these responders, who implement the eviction logic and update the Eviction's status with\n    progress.\n\nThere is many-to-many relationship between EvictionRequests and Evictions in general. And many-to-one if the target is a  pod.\n\nIf all requesters withdraw their eviction intent for a common target, the eviction will be canceled. Deleting an EvictionRequest also counts as a withdrawal. Once all EvictionRequest of a target are removed, the corresponding Evictions are eventually garbage collected.",
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
				Description: "metadata is the standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.",
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
				Description: "spec defines the eviction request specification. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"intent": {
						Type:        schema.TypeString,
						Description: "intent specifies the action that should be taken for the specified target.\n\n- Eviction means that the requester is interested in the eviction of the target. - Withdrawn means that the requester is no longer interested in the eviction of the target.\n  If all requesters' intents are withdrawn for a common target, the eviction will be canceled.\n  Cancellation consequences:\n  - Inactive responders will never run.\n  - Active responders are expected to cancel the eviction.\n  - Completed or Interrupted responders should not take any action.",
						Optional:    false,
						Required:    true,
						Computed:    false,
					},
					"requester": {
						Type:        schema.TypeString,
						Description: "requester allows you to identify the entity, that requested the eviction of the target.\n\nIt must be a valid domain-prefixed key (such as \"acme.io/foo\"). Domain names *.k8s.io and *.kubernetes.io are reserved. This field is required and immutable.",
						Optional:    false,
						Required:    true,
						Computed:    false,
					},
					"target": {
						Type:        schema.TypeList,
						Description: "target contains a reference to an object (e.g. a pod) that should be evicted. This field is required and immutable.",
						Optional:    false,
						Required:    true,
						Computed:    false,
						MinItems:    1,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"pod": {
								Type:        schema.TypeList,
								Description: "pod references a pod that is subject to eviction/termination. Pods that are part of a PodGroup (.spec.schedulingGroup is set) are not supported.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "name of the target. This field is required.",
										Optional:    false,
										Required:    true,
										Computed:    false,
									},
									"uid": {
										Type:        schema.TypeString,
										Description: "uid of the target. It can be found in .metadata.uid of the target and is a lowercase UUID in 8-4-4-4-12 format. This field is required.",
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
			"status": {
				Type:        schema.TypeList,
				Description: "status represents the most recently observed status of the eviction request. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"conditions": {
						Type:        schema.TypeList,
						Description: "conditions contain information about the eviction request.\n\nEvictionRequest specific conditions are: TargetEvicted or Failed (managed by evictionrequest-controller). - Failed means that the eviction request is no longer being processed\n  by any eviction responder. This can happen if the request is canceled or if no responder\n  managed to evict the target (e.g. terminate or delete a pod).\n- TargetEvicted means that the target has been evicted (e.g. a pod has been terminated or deleted).\n\nThese conditions can be reset if the eviction was unsuccessful and a new Eviction intent has been submitted.\n\nThe maximum length of the conditions list is 100.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"last_transition_time": {
								Type:        schema.TypeString,
								Description: "lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"message": {
								Type:        schema.TypeString,
								Description: "message is a human readable message indicating details about the transition. This may be an empty string.",
								Optional:    false,
								Required:    true,
								Computed:    false,
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
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"status": {
								Type:        schema.TypeString,
								Description: "status of the condition, one of True, False, Unknown.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"type": {
								Type:        schema.TypeString,
								Description: "type of condition in CamelCase or in foo.example.com/CamelCase.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
					"observed_generation": {
						Type:        schema.TypeInt,
						Description: "observedGeneration is EvictionRequest's .metadata.generation observed by the evictionrequest-controller. The observed generation value cannot be negative and can only be incremented. The minimum value is 1. This field is managed by evictionrequest-controller.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceK8sLifecycleK8sIoEvictionRequestV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "lifecycle.k8s.io/v1alpha1", "EvictionRequest", "lifecycle.k8s.io/v1alpha1/EvictionRequest"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"metadata", "spec", "spec.target", "spec.target.pod", "status"}, []string{"metadata", "metadata.managed_fields", "metadata.owner_references", "spec", "spec.target", "spec.target.pod", "status", "status.conditions"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sLifecycleK8sIoEvictionRequestV1Alpha1CompatibleVersions = []string{
	"v1.37.0",
}

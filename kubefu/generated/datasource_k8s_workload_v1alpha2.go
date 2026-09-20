package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sSchedulingK8sIoWorkloadV1Alpha2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sSchedulingK8sIoWorkloadV1Alpha2Read,
		Description: "Workload allows for expressing scheduling constraints that should be used when managing the lifecycle of workloads from the scheduling perspective, including scheduling, preemption, eviction and other phases. Workload API enablement is toggled by the GenericWorkload feature gate.",
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
				Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
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
				Description: "Spec defines the desired behavior of a Workload.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"controller_ref": {
						Type:        schema.TypeList,
						Description: "ControllerRef is an optional reference to the controlling object, such as a Deployment or Job. This field is intended for use by tools like CLIs to provide a link back to the original workload definition. This field is immutable.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_group": {
								Type:        schema.TypeString,
								Description: "APIGroup is the group for the resource being referenced. If APIGroup is empty, the specified Kind must be in the core API group. For any other third-party types, setting APIGroup is required. It must be a DNS subdomain.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "Kind is the type of resource being referenced. It must be a path segment name.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is the name of resource being referenced. It must be a path segment name.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
						}},
					},
					"pod_group_templates": {
						Type:        schema.TypeList,
						Description: "PodGroupTemplates is the list of templates that make up the Workload. The maximum number of templates is 8. This field is immutable.",
						Optional:    false,
						Required:    true,
						Computed:    false,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"disruption_mode": {
								Type:        schema.TypeString,
								Description: "DisruptionMode defines the mode in which a given PodGroup can be disrupted. One of Pod, PodGroup. This field is available only when the WorkloadAwarePreemption feature gate is enabled.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "Name is a unique identifier for the PodGroupTemplate within the Workload. It must be a DNS label. This field is immutable.",
								Optional:    false,
								Required:    true,
								Computed:    false,
							},
							"priority": {
								Type:        schema.TypeInt,
								Description: "Priority is the value of priority of pod groups created from this template. Various system components use this field to find the priority of the pod group. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. This field is available only when the WorkloadAwarePreemption feature gate is enabled.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"priority_class_name": {
								Type:        schema.TypeString,
								Description: "PriorityClassName indicates the priority that should be considered when scheduling a pod group created from this template. If no priority class is specified, admission control can set this to the global default priority class if it exists. Otherwise, pod groups created from this template will have the priority set to zero. This field is available only when the WorkloadAwarePreemption feature gate is enabled.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"resource_claims": {
								Type:        schema.TypeList,
								Description: "ResourceClaims defines which ResourceClaims may be shared among Pods in the group. Pods consume the devices allocated to a PodGroup's claim by defining a claim in its own Spec.ResourceClaims that matches the PodGroup's claim exactly. The claim must have the same name and refer to the same ResourceClaim or ResourceClaimTemplate.\n\nThis is an alpha-level field and requires that the DRAWorkloadResourceClaims feature gate is enabled.\n\nThis field is immutable.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Name uniquely identifies this resource claim inside the PodGroup. This must be a DNS_LABEL.",
										Optional:    false,
										Required:    true,
										Computed:    false,
									},
									"resource_claim_name": {
										Type:        schema.TypeString,
										Description: "ResourceClaimName is the name of a ResourceClaim object in the same namespace as this PodGroup. The ResourceClaim will be reserved for the PodGroup instead of its individual pods.\n\nExactly one of ResourceClaimName and ResourceClaimTemplateName must be set.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"resource_claim_template_name": {
										Type:        schema.TypeString,
										Description: "ResourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace as this PodGroup.\n\nThe template will be used to create a new ResourceClaim, which will be bound to this PodGroup. When this PodGroup is deleted, the ResourceClaim will also be deleted. The PodGroup name and resource name, along with a generated component, will be used to form a unique name for the ResourceClaim, which will be recorded in podgroup.status.resourceClaimStatuses.\n\nThis field is immutable and no changes will be made to the corresponding ResourceClaim by the control plane after creating the ResourceClaim.\n\nExactly one of ResourceClaimName and ResourceClaimTemplateName must be set.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"scheduling_constraints": {
								Type:        schema.TypeList,
								Description: "SchedulingConstraints defines optional scheduling constraints (e.g. topology) for this PodGroupTemplate. This field is only available when the TopologyAwareWorkloadScheduling feature gate is enabled.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"topology": {
										Type:        schema.TypeList,
										Description: "Topology defines the topology constraints for the pod group. Currently only a single topology constraint can be specified. This may change in the future.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "Key specifies the key of the node label representing the topology domain. All pods within the PodGroup must be colocated within the same domain instance. Different PodGroups can land on different domain instances even if they derive from the same PodGroupTemplate. Examples: \"topology.kubernetes.io/rack\"",
												Optional:    false,
												Required:    true,
												Computed:    false,
											},
										}},
									},
								}},
							},
							"scheduling_policy": {
								Type:        schema.TypeList,
								Description: "SchedulingPolicy defines the scheduling policy for this PodGroupTemplate.",
								Optional:    false,
								Required:    true,
								Computed:    false,
								MinItems:    1,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"basic": {
										Type:        schema.TypeMap,
										Description: "Basic specifies that the pods in this group should be scheduled using standard Kubernetes scheduling behavior.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"gang": {
										Type:        schema.TypeList,
										Description: "Gang specifies that the pods in this group should be scheduled using all-or-nothing semantics.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"min_count": {
												Type:        schema.TypeInt,
												Description: "MinCount is the minimum number of pods that must be schedulable or scheduled at the same time for the scheduler to admit the entire group. It must be a positive integer.",
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
		},
	}
}



func dataSourceK8sSchedulingK8sIoWorkloadV1Alpha2Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "scheduling.k8s.io/v1alpha2", "Workload", "scheduling.k8s.io/v1alpha2/Workload"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec"}, []string{"metadata", "spec", "spec.controller_ref", "spec.pod_group_templates.scheduling_constraints", "spec.pod_group_templates.scheduling_policy", "spec.pod_group_templates.scheduling_policy.gang"}, []string{"metadata", "metadata.managed_fields", "metadata.owner_references", "spec", "spec.controller_ref", "spec.pod_group_templates", "spec.pod_group_templates.resource_claims", "spec.pod_group_templates.scheduling_constraints", "spec.pod_group_templates.scheduling_constraints.topology", "spec.pod_group_templates.scheduling_policy", "spec.pod_group_templates.scheduling_policy.gang"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sSchedulingK8sIoWorkloadV1Alpha2CompatibleVersions = []string{
	"v1.36.0",
}

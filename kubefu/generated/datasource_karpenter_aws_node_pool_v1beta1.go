package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1Read,
		Description: "NodePool is the Schema for the NodePools API",
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
				Description: "NodePoolSpec is the top level nodepool specification. Nodepools\nlaunch nodes in response to pods that are unschedulable. A single nodepool\nis capable of managing a diverse set of nodes. Node properties are determined\nfrom a combination of nodepool and pod scheduling constraints.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"disruption": {
						Type:        schema.TypeList,
						Description: "Disruption contains the parameters that relate to Karpenter's disruption logic",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"budgets": {
								Type:        schema.TypeList,
								Description: "Budgets is a list of Budgets.\nIf there are multiple active budgets, Karpenter uses\nthe most restrictive value. If left undefined,\nthis will default to one budget with a value to 10%.",
								Optional:    true,
								Required:    false,
								Computed:    true,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"duration": {
										Type:        schema.TypeString,
										Description: "Duration determines how long a Budget is active since each Schedule hit.\nOnly minutes and hours are accepted, as cron does not work in seconds.\nIf omitted, the budget is always active.\nThis is required if Schedule is set.\nThis regex has an optional 0s at the end since the duration.String() always adds\na 0s at the end.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"nodes": {
										Type:        schema.TypeString,
										Description: "Nodes dictates the maximum number of NodeClaims owned by this NodePool\nthat can be terminating at once. This is calculated by counting nodes that\nhave a deletion timestamp set, or are actively being deleted by Karpenter.\nThis field is required when specifying a budget.\nThis cannot be of type intstr.IntOrString since kubebuilder doesn't support pattern\nchecking for int nodes for IntOrString nodes.\nRef: https://github.com/kubernetes-sigs/controller-tools/blob/55efe4be40394a288216dab63156b0a64fb82929/pkg/crd/markers/validation.go#L379-L388",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"schedule": {
										Type:        schema.TypeString,
										Description: "Schedule specifies when a budget begins being active, following\nthe upstream cronjob syntax. If omitted, the budget is always active.\nTimezones are not supported.\nThis field is required if Duration is set.",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"consolidate_after": {
								Type:        schema.TypeString,
								Description: "ConsolidateAfter is the duration the controller will wait\nbefore attempting to terminate nodes that are underutilized.\nRefer to ConsolidationPolicy for how underutilization is considered.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"consolidation_policy": {
								Type:        schema.TypeString,
								Description: "ConsolidationPolicy describes which nodes Karpenter can disrupt through its consolidation\nalgorithm. This policy defaults to \"WhenUnderutilized\" if not specified",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"expire_after": {
								Type:        schema.TypeString,
								Description: "ExpireAfter is the duration the controller will wait\nbefore terminating a node, measured from when the node is created. This\nis useful to implement features like eventually consistent node upgrade,\nmemory leak protection, and disruption testing.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"limits": {
						Type:        schema.TypeMap,
						Description: "Limits define a set of bounds for provisioning capacity.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"template": {
						Type:        schema.TypeList,
						Description: "Template contains the template of possibilities for the provisioning logic to launch a NodeClaim with.\nNodeClaims launched from this NodePool will often be further constrained than the template specifies.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"metadata": {
								Type:        schema.TypeList,
								Description: "",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"annotations": {
										Type:        schema.TypeMap,
										Description: "Annotations is an unstructured key value map stored with a resource that may be\nset by external tools to store and retrieve arbitrary metadata. They are not\nqueryable and should be preserved when modifying objects.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
									"labels": {
										Type:        schema.TypeMap,
										Description: "Map of string keys and values that can be used to organize and categorize\n(scope and select) objects. May match selectors of replication controllers\nand services.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										Optional:    true,
										Required:    false,
										Computed:    true,
									},
								}},
							},
							"spec": {
								Type:        schema.TypeList,
								Description: "NodeClaimSpec describes the desired state of the NodeClaim",
								Optional:    true,
								Required:    false,
								Computed:    true,
								MaxItems:    1,
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"kubelet": {
										Type:        schema.TypeList,
										Description: "Kubelet defines args to be used when configuring kubelet on provisioned nodes.\nThey are a subset of the upstream types, recognizing not all options may be supported.\nWherever possible, the types and names should reflect the upstream kubelet types.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"cluster_dns": {
												Type:        schema.TypeList,
												Description: "clusterDNS is a list of IP addresses for the cluster DNS server.\nNote that not all providers may use all addresses.",
												Optional:    true,
												Required:    false,
												Computed:    true,
												Elem: &schema.Schema{Type: schema.TypeString},
											},
											"cpu_cfs_quota": {
												Type:        schema.TypeBool,
												Description: "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"eviction_hard": {
												Type:        schema.TypeMap,
												Description: "EvictionHard is the map of signal names to quantities that define hard eviction thresholds",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"eviction_max_pod_grace_period": {
												Type:        schema.TypeInt,
												Description: "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods in\nresponse to soft eviction thresholds being met.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"eviction_soft": {
												Type:        schema.TypeMap,
												Description: "EvictionSoft is the map of signal names to quantities that define soft eviction thresholds",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"eviction_soft_grace_period": {
												Type:        schema.TypeMap,
												Description: "EvictionSoftGracePeriod is the map of signal names to quantities that define grace periods for each eviction signal",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"image_gc_high_threshold_percent": {
												Type:        schema.TypeInt,
												Description: "ImageGCHighThresholdPercent is the percent of disk usage after which image\ngarbage collection is always run. The percent is calculated by dividing this\nfield value by 100, so this field must be between 0 and 100, inclusive.\nWhen specified, the value must be greater than ImageGCLowThresholdPercent.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"image_gc_low_threshold_percent": {
												Type:        schema.TypeInt,
												Description: "ImageGCLowThresholdPercent is the percent of disk usage before which image\ngarbage collection is never run. Lowest disk usage to garbage collect to.\nThe percent is calculated by dividing this field value by 100,\nso the field value must be between 0 and 100, inclusive.\nWhen specified, the value must be less than imageGCHighThresholdPercent",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"kube_reserved": {
												Type:        schema.TypeMap,
												Description: "KubeReserved contains resources reserved for Kubernetes system components.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"max_pods": {
												Type:        schema.TypeInt,
												Description: "MaxPods is an override for the maximum number of pods that can run on\na worker node instance.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"pods_per_core": {
												Type:        schema.TypeInt,
												Description: "PodsPerCore is an override for the number of pods that can run on a worker node\ninstance based on the number of cpu cores. This value cannot exceed MaxPods, so, if\nMaxPods is a lower value, that value will be used.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"system_reserved": {
												Type:        schema.TypeMap,
												Description: "SystemReserved contains resources reserved for OS system daemons and kernel memory.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"node_class_ref": {
										Type:        schema.TypeList,
										Description: "NodeClassRef is a reference to an object that defines provider specific configuration",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"api_version": {
												Type:        schema.TypeString,
												Description: "API version of the referent",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"kind": {
												Type:        schema.TypeString,
												Description: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds\"",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"name": {
												Type:        schema.TypeString,
												Description: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"requirements": {
										Type:        schema.TypeList,
										Description: "Requirements are layered with GetLabels and applied to every node.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Description: "The label key that the selector applies to.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"min_values": {
												Type:        schema.TypeInt,
												Description: "This field is ALPHA and can be dropped or replaced at any time\nMinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"operator": {
												Type:        schema.TypeString,
												Description: "Represents a key's relationship to a set of values.\nValid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"values": {
												Type:        schema.TypeList,
												Description: "An array of string values. If the operator is In or NotIn,\nthe values array must be non-empty. If the operator is Exists or DoesNotExist,\nthe values array must be empty. If the operator is Gt or Lt, the values\narray must have a single element, which will be interpreted as an integer.\nThis array is replaced during a strategic merge patch.",
												Optional:    true,
												Required:    false,
												Computed:    true,
												Elem: &schema.Schema{Type: schema.TypeString},
											},
										}},
									},
									"resources": {
										Type:        schema.TypeList,
										Description: "Resources models the resource requirements for the NodeClaim to launch",
										Optional:    true,
										Required:    false,
										Computed:    true,
										MaxItems:    1,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"requests": {
												Type:        schema.TypeMap,
												Description: "Requests describes the minimum required resources for the NodeClaim to launch",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"startup_taints": {
										Type:        schema.TypeList,
										Description: "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automatically\nwithin a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used by\ndaemonsets to allow initialization and enforce startup ordering.  StartupTaints are ignored for provisioning\npurposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"effect": {
												Type:        schema.TypeString,
												Description: "Required. The effect of the taint on pods\nthat do not tolerate the taint.\nValid effects are NoSchedule, PreferNoSchedule and NoExecute.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"key": {
												Type:        schema.TypeString,
												Description: "Required. The taint key to be applied to a node.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"time_added": {
												Type:        schema.TypeString,
												Description: "TimeAdded represents the time at which the taint was added.\nIt is only written for NoExecute taints.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"value": {
												Type:        schema.TypeString,
												Description: "The taint value corresponding to the taint key.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
										}},
									},
									"taints": {
										Type:        schema.TypeList,
										Description: "Taints will be applied to the NodeClaim's node.",
										Optional:    true,
										Required:    false,
										Computed:    true,
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"effect": {
												Type:        schema.TypeString,
												Description: "Required. The effect of the taint on pods\nthat do not tolerate the taint.\nValid effects are NoSchedule, PreferNoSchedule and NoExecute.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"key": {
												Type:        schema.TypeString,
												Description: "Required. The taint key to be applied to a node.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"time_added": {
												Type:        schema.TypeString,
												Description: "TimeAdded represents the time at which the taint was added.\nIt is only written for NoExecute taints.",
												Optional:    true,
												Required:    false,
												Computed:    true,
											},
											"value": {
												Type:        schema.TypeString,
												Description: "The taint value corresponding to the taint key.",
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
					"weight": {
						Type:        schema.TypeInt,
						Description: "Weight is the priority given to the nodepool during scheduling. A higher\nnumerical weight indicates that this nodepool will be ordered\nahead of other nodepools with lower weights. A nodepool with no weight\nwill be treated as if it is a nodepool with a weight of 0.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "NodePoolStatus defines the observed state of NodePool",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"resources": {
						Type:        schema.TypeMap,
						Description: "Resources is the list of resources that have been provisioned.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "karpenter.sh/v1beta1", "NodePool", "karpenter.sh/v1beta1/NodePool"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.disruption", "spec.template", "spec.template.metadata", "spec.template.spec", "spec.template.spec.kubelet", "spec.template.spec.node_class_ref", "spec.template.spec.resources", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1CompatibleVersions = []string{
	"v0.37.0",
	"v0.37.1",
	"v0.37.2",
	"v0.37.3",
	"v0.37.4",
	"v0.37.5",
	"v0.37.6",
	"v0.37.7",
	"v0.37.8",
	"v1.0.0",
	"v1.0.1",
	"v1.0.2",
	"v1.0.3",
	"v1.0.4",
	"v1.0.5",
	"v1.0.6",
	"v1.0.7",
	"v1.0.8",
	"v1.0.9",
	"v1.0.10",
	"v1.0.11",
	"v1.0.12",
}

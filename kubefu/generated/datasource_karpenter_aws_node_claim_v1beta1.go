package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1Read,
		Description: "NodeClaim is the Schema for the NodeClaims API",
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
				Description: "NodeClaimSpec describes the desired state of the NodeClaim",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
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
			"status": {
				Type:        schema.TypeList,
				Description: "NodeClaimStatus defines the observed state of NodeClaim",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allocatable": {
						Type:        schema.TypeMap,
						Description: "Allocatable is the estimated allocatable capacity of the node",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"capacity": {
						Type:        schema.TypeMap,
						Description: "Capacity is the estimated full capacity of the node",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"conditions": {
						Type:        schema.TypeList,
						Description: "Conditions contains signals for health and readiness",
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
								Description: "type of condition in CamelCase or in foo.example.com/CamelCase.\n---\nMany .condition.type values are consistent across resources like Available, but because arbitrary conditions can be\nuseful (see .node.status.conditions), the ability to deconflict is important.\nThe regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"image_id": {
						Type:        schema.TypeString,
						Description: "ImageID is an identifier for the image that runs on the node",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"node_name": {
						Type:        schema.TypeString,
						Description: "NodeName is the name of the corresponding node object",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"provider_id": {
						Type:        schema.TypeString,
						Description: "ProviderID of the corresponding node object",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "karpenter.sh/v1beta1", "NodeClaim", "karpenter.sh/v1beta1/NodeClaim"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "spec.kubelet", "spec.node_class_ref", "spec.resources", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1CompatibleVersions = []string{
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

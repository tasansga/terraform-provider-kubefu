package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceKarpenterCoreAutoscalingXK8sIoCapacityBufferV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKarpenterCoreAutoscalingXK8sIoCapacityBufferV1Beta1Read,
		Description: "CapacityBuffer is the configuration that an autoscaler can use to provision buffer capacity within a cluster.\nThis buffer is represented by placeholder pods that trigger the Cluster Autoscaler to scale up nodes in advance,\nensuring that there is always spare capacity available to handle sudden workload spikes or to speed up scaling events.",
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
				Description: "spec defines the desired characteristics of the buffer.",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"limits": {
						Type:        schema.TypeMap,
						Description: "limits specifies resource constraints that limit the number of chunks created for this buffer\nbased on total resource requests (e.g., CPU, memory). If there are no other\nlimitations for the number of chunks (i.e., `replicas` or `percentage` are not set),\nthis will be used to create as many chunks as fit into these limits.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"percentage": {
						Type:        schema.TypeInt,
						Description: "percentage defines the desired buffer capacity as a percentage of the\n`scalableRef`'s current replicas. This is only applicable if `scalableRef` is set.\nThe absolute number of replicas is calculated from the percentage by rounding up to a minimum of 1.\nFor example, if `scalableRef` has 10 replicas and `percentage` is 20, 2 buffer chunks will be created.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"pod_template_ref": {
						Type:        schema.TypeList,
						Description: "podTemplateRef is a reference to a PodTemplate resource in the same namespace\nthat declares the shape of a single chunk of the buffer. The pods created\nfrom this template will be used as placeholder pods for the buffer capacity.\nExactly one of `podTemplateRef`, `scalableRef` should be specified.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "name of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"provisioning_strategy": {
						Type:        schema.TypeString,
						Description: "provisioningStrategy defines how the buffer is utilized.\n\"buffer.x-k8s.io/active-capacity\" is the default strategy, where the buffer actively scales up the cluster by creating placeholder pods.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"replicas": {
						Type:        schema.TypeInt,
						Description: "replicas defines the desired number of buffer chunks to provision.\nIf neither `replicas` nor `percentage` is set, as many chunks as fit within\ndefined resource limits (if any) will be created. If both are set, the minimum\nof the two will be used.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"scalable_ref": {
						Type:        schema.TypeList,
						Description: "scalableRef is a reference to an object of a kind that has a scale subresource\nand specifies its label selector field. This allows the CapacityBuffer to\nmanage the buffer by scaling an existing scalable resource.\nExactly one of `podTemplateRef`, `scalableRef` should be specified.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"api_group": {
								Type:        schema.TypeString,
								Description: "apiGroup is the API group of the referent.\nEmpty string for the core API group.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"kind": {
								Type:        schema.TypeString,
								Description: "kind is the kind of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
							"name": {
								Type:        schema.TypeString,
								Description: "name is the name of the referent.",
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
				Description: "status represents the current state of the buffer and its readiness for autoprovisioning.",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"conditions": {
						Type:        schema.TypeList,
						Description: "conditions provide a standard mechanism for reporting the buffer's state.\nThe \"Ready\" condition indicates if the buffer is successfully provisioned\nand active. Other conditions may report on various aspects of the buffer's\nhealth and provisioning process.",
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
					"pod_template_generation": {
						Type:        schema.TypeInt,
						Description: "podTemplateGeneration is the observed generation of the PodTemplate, used\nto determine if the status is up-to-date with the desired `spec.podTemplateRef`.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"pod_template_ref": {
						Type:        schema.TypeList,
						Description: "podTemplateRef is the observed reference to the PodTemplate that was used\nto provision the buffer. If this field is not set, and the `conditions`\nindicate an error, it provides details about the error state.",
						Optional:    true,
						Required:    false,
						Computed:    true,
						MaxItems:    1,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Description: "name of the referent.",
								Optional:    true,
								Required:    false,
								Computed:    true,
							},
						}},
					},
					"provisioning_strategy": {
						Type:        schema.TypeString,
						Description: "provisioningStrategy defines how the buffer should be utilized.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"replicas": {
						Type:        schema.TypeInt,
						Description: "replicas is the actual number of buffer chunks currently provisioned.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
		},
	}
}



func dataSourceKarpenterCoreAutoscalingXK8sIoCapacityBufferV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "autoscaling.x-k8s.io/v1beta1", "CapacityBuffer", "autoscaling.x-k8s.io/v1beta1/CapacityBuffer"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPathsForMeta(d, m, []string{"metadata", "spec", "status"}, []string{"spec", "spec.pod_template_ref", "spec.scalable_ref", "status", "status.pod_template_ref"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceKarpenterCoreAutoscalingXK8sIoCapacityBufferV1Beta1CompatibleVersions = []string{
	"v1.14.0",
	"v1.14.1",
}

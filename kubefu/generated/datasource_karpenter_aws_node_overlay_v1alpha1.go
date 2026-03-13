package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1Read,
		Description: "Generated data source for crd.karpenter.sh.v1alpha1.NodeOverlay",
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
				Description: "",
				Optional:    false,
				Required:    true,
				Computed:    false,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"capacity": {
						Type:        schema.TypeMap,
						Description: "Capacity adds extended resources only, and does not replace any existing resources.\nThese extended resources are appended to the node's existing resource list.\nNote: This field does not modify or override standard resources like cpu, memory, ephemeral-storage, or pods.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"price": {
						Type:        schema.TypeString,
						Description: "Price specifies amount for an instance types that match the specified labels. Users can override prices using a signed float representing the price override",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"price_adjustment": {
						Type:        schema.TypeString,
						Description: "PriceAdjustment specifies the price change for matching instance types. Accepts either:\n- A fixed price modifier (e.g., -0.5, 1.2)\n- A percentage modifier (e.g., +10% for increase, -15% for decrees)",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
					"requirements": {
						Type:        schema.TypeList,
						Description: "Requirements constrain when this NodeOverlay is applied during scheduling simulations.\nThese requirements can match:\n- Well-known labels (e.g., node.kubernetes.io/instance-type, karpenter.sh/nodepool)\n- Custom labels from NodePool's spec.template.labels",
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
					"weight": {
						Type:        schema.TypeInt,
						Description: "Weight defines the priority of this NodeOverlay when overriding node attributes.\nNodeOverlays with higher numerical weights take precedence over those with lower weights.\nIf no weight is specified, the NodeOverlay is treated as having a weight of 0.\nWhen multiple NodeOverlays have identical weights, they are merged in alphabetical order.",
						Optional:    true,
						Required:    false,
						Computed:    true,
					},
				}},
			},
			"status": {
				Type:        schema.TypeList,
				Description: "NodeOverlayStatus defines the observed state of NodeOverlay",
				Optional:    true,
				Required:    false,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
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
								Description: "type of condition in CamelCase or in foo.example.com/CamelCase.",
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



func dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "karpenter.sh/v1alpha1", "NodeOverlay", "karpenter.sh/v1alpha1/NodeOverlay"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifestWithObjectPaths(d, []string{"metadata", "spec", "status"}, []string{"spec", "status"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1CompatibleVersions = []string{
	"v1.7.0",
	"v1.7.1",
	"v1.7.2",
	"v1.7.3",
	"v1.7.4",
	"v1.8.0",
	"v1.8.1",
	"v1.8.2",
	"v1.8.3",
	"v1.8.4",
	"v1.8.5",
	"v1.8.6",
	"v1.9.0",
}

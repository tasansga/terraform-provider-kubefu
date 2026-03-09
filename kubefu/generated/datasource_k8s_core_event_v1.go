package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
)

func dataSourceK8sCoreEventV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sCoreEventV1Read,
		Description: "Event is a report of an event somewhere in the cluster.",
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Description: "What action was taken/failed regarding to the Regarding object.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"count_": {
				Type:        schema.TypeInt,
				Description: "The number of times this event has occurred.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"event_time": {
				Type:        schema.TypeMap,
				Description: "Time when this Event was first observed.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"first_timestamp": {
				Type:        schema.TypeMap,
				Description: "The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"involved_object": {
				Type:        schema.TypeMap,
				Description: "The object that this event is about.",
				Optional:    false,
				Required:    true,
				Computed:    false,
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
			"last_timestamp": {
				Type:        schema.TypeMap,
				Description: "The time at which the most recent occurrence of this event was recorded.",
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
				Type:        schema.TypeMap,
				Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"reason": {
				Type:        schema.TypeString,
				Description: "This should be a short, machine understandable string that gives the reason for the transition into the object's current status.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"related": {
				Type:        schema.TypeMap,
				Description: "Optional secondary object for more complex actions.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reporting_component": {
				Type:        schema.TypeString,
				Description: "Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reporting_instance": {
				Type:        schema.TypeString,
				Description: "ID of the controller instance, e.g. `kubelet-xyzf`.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"series": {
				Type:        schema.TypeMap,
				Description: "Data about the Event series this event represents or nil if it's a singleton Event.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"source": {
				Type:        schema.TypeMap,
				Description: "The component reporting this event. Should be a short machine understandable string.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of this event (Normal, Warning), new types could be added in the future",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sCoreEventV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "v1", "Event", "core/v1/Event"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"action", "count_", "event_time", "first_timestamp", "involved_object", "last_timestamp", "message", "metadata", "reason", "related", "reporting_component", "reporting_instance", "series", "source", "type"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sCoreEventV1CompatibleVersions = []string{
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

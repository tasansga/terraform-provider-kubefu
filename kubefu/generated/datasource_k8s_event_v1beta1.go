package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

func dataSourceK8sEventsK8sIoEventV1Beta1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sEventsK8sIoEventV1Beta1Read,
		Description: "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.",
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Description: "What action was taken/failed regarding to the regarding object.",
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
			"deprecated_count": {
				Type:        schema.TypeInt,
				Description: "Deprecated field assuring backward compatibility with core.v1 Event type",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"deprecated_first_timestamp": {
				Type:        schema.TypeMap,
				Description: "Deprecated field assuring backward compatibility with core.v1 Event type",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"deprecated_last_timestamp": {
				Type:        schema.TypeMap,
				Description: "Deprecated field assuring backward compatibility with core.v1 Event type",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"deprecated_source": {
				Type:        schema.TypeMap,
				Description: "Deprecated field assuring backward compatibility with core.v1 Event type",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"event_time": {
				Type:        schema.TypeMap,
				Description: "Required. Time when this Event was first observed.",
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
			"metadata": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"note": {
				Type:        schema.TypeString,
				Description: "Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reason": {
				Type:        schema.TypeString,
				Description: "Why the action was taken.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"regarding": {
				Type:        schema.TypeMap,
				Description: "The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"related": {
				Type:        schema.TypeMap,
				Description: "Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reporting_controller": {
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
			"type": {
				Type:        schema.TypeString,
				Description: "Type of this event (Normal, Warning), new types could be added in the future.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sEventsK8sIoEventV1Beta1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "events.k8s.io/v1beta1", "Event", "events.k8s.io/v1beta1/Event"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"action", "deprecated_count", "deprecated_first_timestamp", "deprecated_last_timestamp", "deprecated_source", "event_time", "metadata", "note", "reason", "regarding", "related", "reporting_controller", "reporting_instance", "series", "type"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sEventsK8sIoEventV1Beta1CompatibleVersions = []string{
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
}

func dataSourceK8sEventsK8sIoEventV1Beta1IsCompatibleWith(version string) bool {
	return versionpkg.IsCompatibleWith(version, dataSourceK8sEventsK8sIoEventV1Beta1CompatibleVersions)
}

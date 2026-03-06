package generated

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	manifestpkg "github.com/tasansga/terraform-provider-kubefu/kubefu/internal/manifest"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

func dataSourceK8sEventsK8sIoEventV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceK8sEventsK8sIoEventV1Read,
		Description: "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.",
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Description: "action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"api_version": {
				Type:        schema.TypeString,
				Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Optional:    false,
				Required:    false,
				Computed:    true,
			},
			"deprecated_count": {
				Type:        schema.TypeInt,
				Description: "deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"deprecated_first_timestamp": {
				Type:        schema.TypeMap,
				Description: "deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"deprecated_last_timestamp": {
				Type:        schema.TypeMap,
				Description: "deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"deprecated_source": {
				Type:        schema.TypeMap,
				Description: "deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"event_time": {
				Type:        schema.TypeMap,
				Description: "eventTime is the time when this Event was first observed. It is required.",
				Optional:    false,
				Required:    true,
				Computed:    false,
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
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"note": {
				Type:        schema.TypeString,
				Description: "note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reason": {
				Type:        schema.TypeString,
				Description: "reason is why the action was taken. It is human-readable. This field can have at most 128 characters.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"regarding": {
				Type:        schema.TypeMap,
				Description: "regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"related": {
				Type:        schema.TypeMap,
				Description: "related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reporting_controller": {
				Type:        schema.TypeString,
				Description: "reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"reporting_instance": {
				Type:        schema.TypeString,
				Description: "reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"series": {
				Type:        schema.TypeMap,
				Description: "series is data about the Event series this event represents or nil if it's a singleton Event.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.",
				Optional:    true,
				Required:    false,
				Computed:    true,
			},
		},
	}
}



func dataSourceK8sEventsK8sIoEventV1Read(_ context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := manifestpkg.SetDataSourceDefaults(d, "events.k8s.io/v1", "Event", "events.k8s.io/v1/Event"); err != nil {
		return diag.FromErr(err)
	}
	if err := manifestpkg.SetDataSourceManifest(d, []string{"action", "deprecated_count", "deprecated_first_timestamp", "deprecated_last_timestamp", "deprecated_source", "event_time", "metadata", "note", "reason", "regarding", "related", "reporting_controller", "reporting_instance", "series", "type"}); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
var dataSourceK8sEventsK8sIoEventV1CompatibleVersions = []string{
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

func dataSourceK8sEventsK8sIoEventV1IsCompatibleWith(version string) bool {
	return versionpkg.IsCompatibleWith(version, dataSourceK8sEventsK8sIoEventV1CompatibleVersions)
}

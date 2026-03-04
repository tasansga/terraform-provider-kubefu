package kubefu

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"kubefu_manifest": resourceManifest(),
	}
}

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	kubefu "github.com/tasansga/terraform-provider-kubefu/kubefu"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kubefu.Provider,
	})
}

package main

import (
	"github.com/chrisUsick/terraform-provider-ping/ping"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return ping.ProviderFactory(nil)
		},
	})
}

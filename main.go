package main

import (
	"github.com/botlink/terraform-provider-deis/deis"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: deis.Provider,
	})
}

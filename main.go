package main

//go:generate ./gen.sh

import (
	"github.com/PetrusHahol/terraform-provider-contrail/contrail/resources"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return resources.Provider()
		},
	})
}

package resources

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccNetworkBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckRequiredEnvVars(t) },
		Providers: testAccProviders,
		CheckDestroy: func(s *terraform.State) error {
			return testAccCheckDestroy(s, "contrail_virtual_network", "virtual-network")
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNetworking_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
				),
			},
		},
	})
}

const testAccNetworking_basic = `
resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name"
}`

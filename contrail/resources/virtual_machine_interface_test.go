package resources

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccPortBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckRequiredEnvVars(t) },
		Providers: testAccProviders,
		CheckDestroy: func(s *terraform.State) error {
			if err := testAccCheckDestroy(s, "contrail_virtual_network", "virtual-network"); err != nil {
				return err
			}
			if err := testAccCheckDestroy(s, "contrail_virtual_machine_interface", "virtual-machine-interface"); err != nil {
				return err
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPort_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_machine_interface.port_test", "virtual-machine-interface"),
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
				),
			},
		},
	})
}

const testAccPort_basic = `

resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name"
}

resource "contrail_virtual_machine_interface" "port_test" {
 	 name  = "port_test"
	 display_name = "test_display_name"
  	 parent_uuid = "c4e1d6c4-a040-4b5c-9da0-48f1da4ce432"
     virtual_network_refs {
    	to = contrail_virtual_network.network_test.id
     }
}
`

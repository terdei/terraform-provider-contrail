package resources

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

// step #2 should be failed before implementation of reading and updating for references.
// config will not contain virtual machine id once VM's will be tested.
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
			{
				Config: testAccPort_with_additional_network_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_machine_interface.port_test", "virtual-machine-interface"),
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckExists("contrail_virtual_network.network_test2", "virtual-network"),
				),
			},
		},
	})
}

const testAccPort_basic = testAccNetworking_basic + `

resource "contrail_virtual_machine_interface" "port_test" {
 	 name  = "port_test"
	 display_name = "test_display_name"
  	 parent_uuid = "824d5eb9-5f9c-407d-b957-2eec0ec8f269"
     virtual_network_refs {
    	to = contrail_virtual_network.network_test.id
     }
}
`

const testAccPort_with_additional_network_basic = testAccNetworking_basic + `

resource "contrail_virtual_network" "network_test2" {
	name = "test_name2"
	display_name = "test_display_name2"
}

resource "contrail_virtual_machine_interface" "port_test" {
 	 name  = "port_test"
	 display_name = "test_display_name"
  	 parent_uuid = "824d5eb9-5f9c-407d-b957-2eec0ec8f269"
     virtual_network_refs {
    	to = contrail_virtual_network.network_test2.id
     }
	depends_on = [ contrail_virtual_machine_interface.network_test2.id ]
}

`

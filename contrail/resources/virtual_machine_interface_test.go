package resources

import (
	"fmt"
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
					testAccCheckExists("contrail_virtual_machine_interface.port_test1", "virtual-machine-interface"),
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckRefExists("contrail_virtual_machine_interface.port_test1", "virtual-machine-interface",
						"contrail_virtual_network.network_test", "virtual-network", "virtual_network_refs"),
				),
			},
			{
				Config: testAccPortWithAdditionalNetworkBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_machine_interface.port_test1", "virtual-machine-interface"),
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckExists("contrail_virtual_network.network_test2", "virtual-network"),
					testAccCheckRefExists("contrail_virtual_machine_interface.port_test1", "virtual-machine-interface",
						"contrail_virtual_network.network_test2", "virtual-network", "virtual_network_refs"),
				),
			},
		},
	})
}

var testAccPort_basic = testAccNetworking_basic + fmt.Sprintf(`
resource "contrail_virtual_machine_interface" "port_test1" {
 	 name  = "port_test1"
	 display_name = "test_display_name1"
  	 parent_uuid = "%s"
     virtual_network_refs {
    	to = contrail_virtual_network.network_test.id
     }
} 
`, OS_VM_ID)

var testAccPortWithAdditionalNetworkBasic = testAccNetworking_basic + fmt.Sprintf(`

resource "contrail_virtual_network" "network_test2" {
	name = "test_name2"
	display_name = "test_display_name2"
    parent_uuid = "%s"
}

resource "contrail_virtual_machine_interface" "port_test1" {
 	 name  = "port_test1"
	 display_name = "test_display_name1"
  	 parent_uuid = "%s"
     virtual_network_refs {
    	to = contrail_virtual_network.network_test2.id
     }
	 depends_on = [ contrail_virtual_network.network_test2 ]
}
`, OS_PROJECT_ID, OS_VM_ID)

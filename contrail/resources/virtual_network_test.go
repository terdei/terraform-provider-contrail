package resources

import (
	"fmt"
	"github.com/Juniper/contrail-go-api"
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

func TestAccNetworkRefsBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckRequiredEnvVars(t) },
		Providers: testAccProviders,
		CheckDestroy: func(s *terraform.State) error {
			return testAccCheckDestroy(s, "contrail_virtual_network", "virtual-network")
		},
		Steps: []resource.TestStep{
			{
				Config: testTagRef_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckExists("contrail_tag.tag_test", "tag"),
					testAccCheckTagRefExists("contrail_virtual_network.network_test",
						"virtual-network",
						"contrail_tag.tag_test",
						"tag"),
				),
			},
			{
				Config: testTagRefWithoutRef_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckExists("contrail_tag.tag_test", "tag"),
					testAccCheckRefDeleted("contrail_virtual_network.network_test",
						"virtual-network",
						"contrail_tag.tag_test"),
				),
			},
		},
	})
}

func testAccCheckTagRefExists(key string, clientKey string, refKey string, refClientKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[key]
		if !ok {
			return fmt.Errorf("Not found: %s", key)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*contrail.Client)
		base, err := client.FindByUuid(clientKey, rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error creating %s : %s", clientKey, err)
		}

		rs, ok = s.RootModule().Resources[refKey]
		if !ok {
			return fmt.Errorf("Not found: %s", refKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		refBase, err := client.FindByUuid(refClientKey, rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error creating %s : %s", refClientKey, err)
		}
		tagUUID := base.(*VirtualNetwork).tag_refs[0].Uuid
		if tagUUID != refBase.GetUuid() {
			return fmt.Errorf("Error creating reference from %s object to %s object", key, refKey)
		}

		return nil
	}
}

func testAccCheckRefDeleted(key string, clientKey string, refKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[key]
		if !ok {
			return fmt.Errorf("Not found: %s", key)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*contrail.Client)
		base, err := client.FindByUuid(clientKey, rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error creating %s : %s", clientKey, err)
		}

		if len(base.(*VirtualNetwork).tag_refs) != 0 {
			return fmt.Errorf("Error deleting reference from %s object to %s object", key, refKey)
		}

		return nil
	}
}

const testAccNetworking_basic = `
resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name"
}`

const testTagRef_basic = `

resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name"
}

resource "contrail_virtual_network_refs" "network_ref" {
	uuid = contrail_virtual_network.network_test.id
	tag_refs {
	  to = contrail_tag.tag_test.id
	}
	depends_on = [ contrail_tag.tag_test ]
}

resource "contrail_tag" "tag_test" {
    name = "test_tag"
    tag_value = "value"
    tag_type_name = "type"
	display_name = "type=value"
}

`

const testTagRefWithoutRef_basic = `
resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name"
}

resource "contrail_tag" "tag_test" {
    name = "test_tag"
    tag_value = "value"
    tag_type_name = "type"
	display_name = "type=value"
}
`

package resources

import (
	"fmt"
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
			if err := testAccCheckDestroy(s, "contrail_virtual_network", "virtual-network"); err != nil {
				return err
			}
			if err := testAccCheckDestroy(s, "contrail_tag", "tag"); err != nil {
				return err
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testTagRef_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckExists("contrail_tag.tag_test", "tag"),
					testAccCheckRefExists("contrail_virtual_network.network_test",
						"virtual-network", "contrail_tag.tag_test", "tag", "tag_refs"),
				),
			},
			{
				Config: testTagRefWithoutRef_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_virtual_network.network_test", "virtual-network"),
					testAccCheckExists("contrail_tag.tag_test", "tag"),
					testAccCheckRefDeleted("contrail_virtual_network.network_test",
						"virtual-network",
						"tag_refs"),
				),
			},
		},
	})
}

var testAccNetworking_basic = fmt.Sprintf(`
resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name" 
  parent_uuid = "%s"
}`, OS_PROJECT_ID)

var testTagRef_basic = testAccNetworking_basic + `

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

var testTagRefWithoutRef_basic = testAccNetworking_basic + `

resource "contrail_tag" "tag_test" {
    name = "test_tag"
    tag_value = "value"
    tag_type_name = "type"
	display_name = "type=value"
}
`

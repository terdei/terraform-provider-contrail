//
// Automatically generated. DO NOT EDIT.
// (Struct Type [aka CType])
//

package resources

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

var _ = spew.Dump // Avoid import errors if not used

func SetDhcpOptionsListTypeFromMap(object *DhcpOptionsListType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetDhcpOptionsListTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mDhcpOptionObj := vmap["dhcp_option"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mDhcpOptionObj) {
		log.Printf("Setting dhcp_option  DhcpOption <<%T>> => %#v", mDhcpOptionObj, mDhcpOptionObj)
		for _, v := range mDhcpOptionObj.([]interface{}) {
			mDhcpOption := new(DhcpOptionType)
			SetDhcpOptionTypeFromMap(mDhcpOption, d, m, v)
			object.AddDhcpOption(mDhcpOption)
		}
	}

	log.Printf("FINISHED DhcpOptionsListType object: %#v", object)
}

func TakeDhcpOptionsListTypeAsMap(object *DhcpOptionsListType) map[string]interface{} {
	omap := make(map[string]interface{})

	var dhcp_option_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.DhcpOption {
		dhcp_option_map = append(dhcp_option_map, TakeDhcpOptionTypeAsMap(&v))
	}
	omap["dhcp_option"] = dhcp_option_map

	return omap
}

func ResourceDhcpOptionsListTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dhcp_option": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: DhcpOptionType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDhcpOptionType(),
		},
	}
}

func ResourceDhcpOptionsListType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceDhcpOptionsListTypeSchema(),
	}
}

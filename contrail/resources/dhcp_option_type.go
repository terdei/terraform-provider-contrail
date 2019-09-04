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

func SetDhcpOptionTypeFromMap(object *DhcpOptionType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetDhcpOptionTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mDhcpOptionNameObj := vmap["dhcp_option_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDhcpOptionNameObj) {
		log.Printf("Setting dhcp_option_name  DhcpOptionName <<%T>> => %#v", mDhcpOptionNameObj, mDhcpOptionNameObj)
		mDhcpOptionName := mDhcpOptionNameObj.(string)
		object.DhcpOptionName = mDhcpOptionName
	}

	mDhcpOptionValueObj := vmap["dhcp_option_value"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDhcpOptionValueObj) {
		log.Printf("Setting dhcp_option_value  DhcpOptionValue <<%T>> => %#v", mDhcpOptionValueObj, mDhcpOptionValueObj)
		mDhcpOptionValue := mDhcpOptionValueObj.(string)
		object.DhcpOptionValue = mDhcpOptionValue
	}

	mDhcpOptionValueBytesObj := vmap["dhcp_option_value_bytes"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDhcpOptionValueBytesObj) {
		log.Printf("Setting dhcp_option_value_bytes  DhcpOptionValueBytes <<%T>> => %#v", mDhcpOptionValueBytesObj, mDhcpOptionValueBytesObj)
		mDhcpOptionValueBytes := mDhcpOptionValueBytesObj.(string)
		object.DhcpOptionValueBytes = mDhcpOptionValueBytes
	}

	log.Printf("FINISHED DhcpOptionType object: %#v", object)
}

func TakeDhcpOptionTypeAsMap(object *DhcpOptionType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["dhcp_option_name"] = object.DhcpOptionName
	omap["dhcp_option_value"] = object.DhcpOptionValue
	omap["dhcp_option_value_bytes"] = object.DhcpOptionValueBytes

	return omap
}

func ResourceDhcpOptionTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dhcp_option_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"dhcp_option_value": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"dhcp_option_value_bytes": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceDhcpOptionType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceDhcpOptionTypeSchema(),
	}
}

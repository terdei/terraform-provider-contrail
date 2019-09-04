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

func SetMacAddressesTypeFromMap(object *MacAddressesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetMacAddressesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mMacAddressObj := vmap["mac_address"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mMacAddressObj) {
		log.Printf("Setting mac_address  MacAddress <<%T>> => %#v", mMacAddressObj, mMacAddressObj)
		for _, v := range mMacAddressObj.([]interface{}) {
			mMacAddress := v.(string)
			object.AddMacAddress(mMacAddress)
		}
	}

	log.Printf("FINISHED MacAddressesType object: %#v", object)
}

func TakeMacAddressesTypeAsMap(object *MacAddressesType) map[string]interface{} {
	omap := make(map[string]interface{})

	mac_address_arr := make([]string, len(object.MacAddress))
	for _, v := range object.MacAddress {
		mac_address_arr = append(mac_address_arr, v)
	}
	omap["mac_address"] = mac_address_arr

	return omap
}

func ResourceMacAddressesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mac_address": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceMacAddressesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceMacAddressesTypeSchema(),
	}
}

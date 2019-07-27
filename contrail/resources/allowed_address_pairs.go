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

func SetAllowedAddressPairsFromMap(object *AllowedAddressPairs, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAllowedAddressPairsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAllowedAddressPairObj := vmap["allowed_address_pair"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mAllowedAddressPairObj) {
		log.Printf("Setting allowed_address_pair  AllowedAddressPair <<%T>> => %#v", mAllowedAddressPairObj, mAllowedAddressPairObj)
		for _, v := range mAllowedAddressPairObj.([]interface{}) {
			mAllowedAddressPair := new(AllowedAddressPair)
			SetAllowedAddressPairFromMap(mAllowedAddressPair, d, m, v)
			object.AddAllowedAddressPair(mAllowedAddressPair)
		}
	}

	log.Printf("FINISHED AllowedAddressPairs object: %#v", object)
}

func TakeAllowedAddressPairsAsMap(object *AllowedAddressPairs) map[string]interface{} {
	omap := make(map[string]interface{})

	var allowed_address_pair_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.AllowedAddressPair {
		allowed_address_pair_map = append(allowed_address_pair_map, TakeAllowedAddressPairAsMap(&v))
	}
	omap["allowed_address_pair"] = allowed_address_pair_map

	return omap
}

func ResourceAllowedAddressPairsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allowed_address_pair": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AllowedAddressPair
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAllowedAddressPair(),
		},
	}
}

func ResourceAllowedAddressPairs() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAllowedAddressPairsSchema(),
	}
}

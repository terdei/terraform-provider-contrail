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

func SetBridgeDomainMembershipTypeFromMap(object *BridgeDomainMembershipType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetBridgeDomainMembershipTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mVlanTagObj := vmap["vlan_tag"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVlanTagObj) {
		log.Printf("Setting vlan_tag  VlanTag <<%T>> => %#v", mVlanTagObj, mVlanTagObj)
		mVlanTag := mVlanTagObj.(int)
		object.VlanTag = mVlanTag
	}

	log.Printf("FINISHED BridgeDomainMembershipType object: %#v", object)
}

func TakeBridgeDomainMembershipTypeAsMap(object *BridgeDomainMembershipType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["vlan_tag"] = object.VlanTag

	return omap
}

func ResourceBridgeDomainMembershipTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vlan_tag": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceBridgeDomainMembershipType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceBridgeDomainMembershipTypeSchema(),
	}
}

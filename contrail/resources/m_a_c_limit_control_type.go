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

func SetMACLimitControlTypeFromMap(object *MACLimitControlType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetMACLimitControlTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mMacLimitObj := vmap["mac_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMacLimitObj) {
		log.Printf("Setting mac_limit  MacLimit <<%T>> => %#v", mMacLimitObj, mMacLimitObj)
		mMacLimit := mMacLimitObj.(int)
		object.MacLimit = mMacLimit
	}

	mMacLimitActionObj := vmap["mac_limit_action"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMacLimitActionObj) {
		log.Printf("Setting mac_limit_action  MacLimitAction <<%T>> => %#v", mMacLimitActionObj, mMacLimitActionObj)
		mMacLimitAction := mMacLimitActionObj.(string)
		object.MacLimitAction = mMacLimitAction
	}

	log.Printf("FINISHED MACLimitControlType object: %#v", object)
}

func TakeMACLimitControlTypeAsMap(object *MACLimitControlType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["mac_limit"] = object.MacLimit
	omap["mac_limit_action"] = object.MacLimitAction

	return omap
}

func ResourceMACLimitControlTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mac_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"mac_limit_action": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceMACLimitControlType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceMACLimitControlTypeSchema(),
	}
}

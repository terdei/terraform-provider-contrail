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

func SetMACMoveLimitControlTypeFromMap(object *MACMoveLimitControlType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetMACMoveLimitControlTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mMacMoveLimitObj := vmap["mac_move_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMacMoveLimitObj) {
		log.Printf("Setting mac_move_limit  MacMoveLimit <<%T>> => %#v", mMacMoveLimitObj, mMacMoveLimitObj)
		mMacMoveLimit := mMacMoveLimitObj.(int)
		object.MacMoveLimit = mMacMoveLimit
	}

	mMacMoveTimeWindowObj := vmap["mac_move_time_window"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMacMoveTimeWindowObj) {
		log.Printf("Setting mac_move_time_window  MacMoveTimeWindow <<%T>> => %#v", mMacMoveTimeWindowObj, mMacMoveTimeWindowObj)
		mMacMoveTimeWindow := mMacMoveTimeWindowObj.(int)
		object.MacMoveTimeWindow = mMacMoveTimeWindow
	}

	mMacMoveLimitActionObj := vmap["mac_move_limit_action"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMacMoveLimitActionObj) {
		log.Printf("Setting mac_move_limit_action  MacMoveLimitAction <<%T>> => %#v", mMacMoveLimitActionObj, mMacMoveLimitActionObj)
		mMacMoveLimitAction := mMacMoveLimitActionObj.(string)
		object.MacMoveLimitAction = mMacMoveLimitAction
	}

	log.Printf("FINISHED MACMoveLimitControlType object: %#v", object)
}

func TakeMACMoveLimitControlTypeAsMap(object *MACMoveLimitControlType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["mac_move_limit"] = object.MacMoveLimit
	omap["mac_move_time_window"] = object.MacMoveTimeWindow
	omap["mac_move_limit_action"] = object.MacMoveLimitAction

	return omap
}

func ResourceMACMoveLimitControlTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mac_move_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"mac_move_time_window": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"mac_move_limit_action": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceMACMoveLimitControlType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceMACMoveLimitControlTypeSchema(),
	}
}

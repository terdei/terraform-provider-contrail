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

func SetPortTypeFromMap(object *PortType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPortTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mStartPortObj := vmap["start_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mStartPortObj) {
		log.Printf("Setting start_port  StartPort <<%T>> => %#v", mStartPortObj, mStartPortObj)
		mStartPort := mStartPortObj.(int)
		object.StartPort = mStartPort
	}

	mEndPortObj := vmap["end_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEndPortObj) {
		log.Printf("Setting end_port  EndPort <<%T>> => %#v", mEndPortObj, mEndPortObj)
		mEndPort := mEndPortObj.(int)
		object.EndPort = mEndPort
	}

	log.Printf("FINISHED PortType object: %#v", object)
}

func TakePortTypeAsMap(object *PortType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["start_port"] = object.StartPort
	omap["end_port"] = object.EndPort

	return omap
}

func ResourcePortTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"end_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourcePortType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePortTypeSchema(),
	}
}

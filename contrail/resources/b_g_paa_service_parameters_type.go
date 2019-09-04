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

func SetBGPaaServiceParametersTypeFromMap(object *BGPaaServiceParametersType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetBGPaaServiceParametersTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPortStartObj := vmap["port_start"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPortStartObj) {
		log.Printf("Setting port_start  PortStart <<%T>> => %#v", mPortStartObj, mPortStartObj)
		mPortStart := mPortStartObj.(int)
		object.PortStart = mPortStart
	}

	mPortEndObj := vmap["port_end"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPortEndObj) {
		log.Printf("Setting port_end  PortEnd <<%T>> => %#v", mPortEndObj, mPortEndObj)
		mPortEnd := mPortEndObj.(int)
		object.PortEnd = mPortEnd
	}

	log.Printf("FINISHED BGPaaServiceParametersType object: %#v", object)
}

func TakeBGPaaServiceParametersTypeAsMap(object *BGPaaServiceParametersType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["port_start"] = object.PortStart
	omap["port_end"] = object.PortEnd

	return omap
}

func ResourceBGPaaServiceParametersTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port_start": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"port_end": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceBGPaaServiceParametersType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceBGPaaServiceParametersTypeSchema(),
	}
}

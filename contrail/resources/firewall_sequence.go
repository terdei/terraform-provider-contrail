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

func SetFirewallSequenceFromMap(object *FirewallSequence, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFirewallSequenceFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSequenceObj := vmap["sequence"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSequenceObj) {
		log.Printf("Setting sequence  Sequence <<%T>> => %#v", mSequenceObj, mSequenceObj)
		mSequence := mSequenceObj.(string)
		object.Sequence = mSequence
	}

	log.Printf("FINISHED FirewallSequence object: %#v", object)
}

func TakeFirewallSequenceAsMap(object *FirewallSequence) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["sequence"] = object.Sequence

	return omap
}

func ResourceFirewallSequenceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sequence": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceFirewallSequence() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFirewallSequenceSchema(),
	}
}

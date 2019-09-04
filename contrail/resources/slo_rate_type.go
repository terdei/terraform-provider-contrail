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

func SetSloRateTypeFromMap(object *SloRateType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSloRateTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRateObj := vmap["rate"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRateObj) {
		log.Printf("Setting rate  Rate <<%T>> => %#v", mRateObj, mRateObj)
		mRate := mRateObj.(int)
		object.Rate = mRate
	}

	log.Printf("FINISHED SloRateType object: %#v", object)
}

func TakeSloRateTypeAsMap(object *SloRateType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["rate"] = object.Rate

	return omap
}

func ResourceSloRateTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rate": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceSloRateType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSloRateTypeSchema(),
	}
}

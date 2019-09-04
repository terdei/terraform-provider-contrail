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

func SetAllocationPoolTypeFromMap(object *AllocationPoolType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAllocationPoolTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mStartObj := vmap["start"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mStartObj) {
		log.Printf("Setting start  Start <<%T>> => %#v", mStartObj, mStartObj)
		mStart := mStartObj.(string)
		object.Start = mStart
	}

	mEndObj := vmap["end"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEndObj) {
		log.Printf("Setting end  End <<%T>> => %#v", mEndObj, mEndObj)
		mEnd := mEndObj.(string)
		object.End = mEnd
	}

	mVrouterSpecificPoolObj := vmap["vrouter_specific_pool"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVrouterSpecificPoolObj) {
		log.Printf("Setting vrouter_specific_pool  VrouterSpecificPool <<%T>> => %#v", mVrouterSpecificPoolObj, mVrouterSpecificPoolObj)
		mVrouterSpecificPool := mVrouterSpecificPoolObj.(bool)
		object.VrouterSpecificPool = mVrouterSpecificPool
	}

	log.Printf("FINISHED AllocationPoolType object: %#v", object)
}

func TakeAllocationPoolTypeAsMap(object *AllocationPoolType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["start"] = object.Start
	omap["end"] = object.End
	omap["vrouter_specific_pool"] = object.VrouterSpecificPool

	return omap
}

func ResourceAllocationPoolTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"end": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"vrouter_specific_pool": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceAllocationPoolType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAllocationPoolTypeSchema(),
	}
}

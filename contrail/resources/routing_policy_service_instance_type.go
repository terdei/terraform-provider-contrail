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

func SetRoutingPolicyServiceInstanceTypeFromMap(object *RoutingPolicyServiceInstanceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRoutingPolicyServiceInstanceTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mLeftSequenceObj := vmap["left_sequence"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLeftSequenceObj) {
		log.Printf("Setting left_sequence  LeftSequence <<%T>> => %#v", mLeftSequenceObj, mLeftSequenceObj)
		mLeftSequence := mLeftSequenceObj.(string)
		object.LeftSequence = mLeftSequence
	}

	mRightSequenceObj := vmap["right_sequence"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRightSequenceObj) {
		log.Printf("Setting right_sequence  RightSequence <<%T>> => %#v", mRightSequenceObj, mRightSequenceObj)
		mRightSequence := mRightSequenceObj.(string)
		object.RightSequence = mRightSequence
	}

	log.Printf("FINISHED RoutingPolicyServiceInstanceType object: %#v", object)
}

func TakeRoutingPolicyServiceInstanceTypeAsMap(object *RoutingPolicyServiceInstanceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["left_sequence"] = object.LeftSequence
	omap["right_sequence"] = object.RightSequence

	return omap
}

func ResourceRoutingPolicyServiceInstanceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"left_sequence": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"right_sequence": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceRoutingPolicyServiceInstanceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRoutingPolicyServiceInstanceTypeSchema(),
	}
}

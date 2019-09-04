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

func SetSequenceTypeFromMap(object *SequenceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSequenceTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mMajorObj := vmap["major"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMajorObj) {
		log.Printf("Setting major  Major <<%T>> => %#v", mMajorObj, mMajorObj)
		mMajor := mMajorObj.(int)
		object.Major = mMajor
	}

	mMinorObj := vmap["minor"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMinorObj) {
		log.Printf("Setting minor  Minor <<%T>> => %#v", mMinorObj, mMinorObj)
		mMinor := mMinorObj.(int)
		object.Minor = mMinor
	}

	log.Printf("FINISHED SequenceType object: %#v", object)
}

func TakeSequenceTypeAsMap(object *SequenceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["major"] = object.Major
	omap["minor"] = object.Minor

	return omap
}

func ResourceSequenceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"major": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"minor": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceSequenceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSequenceTypeSchema(),
	}
}

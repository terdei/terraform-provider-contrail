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

func SetProviderDetailsFromMap(object *ProviderDetails, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetProviderDetailsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSegmentationIdObj := vmap["segmentation_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSegmentationIdObj) {
		log.Printf("Setting segmentation_id  SegmentationId <<%T>> => %#v", mSegmentationIdObj, mSegmentationIdObj)
		mSegmentationId := mSegmentationIdObj.(int)
		object.SegmentationId = mSegmentationId
	}

	mPhysicalNetworkObj := vmap["physical_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPhysicalNetworkObj) {
		log.Printf("Setting physical_network  PhysicalNetwork <<%T>> => %#v", mPhysicalNetworkObj, mPhysicalNetworkObj)
		mPhysicalNetwork := mPhysicalNetworkObj.(string)
		object.PhysicalNetwork = mPhysicalNetwork
	}

	log.Printf("FINISHED ProviderDetails object: %#v", object)
}

func TakeProviderDetailsAsMap(object *ProviderDetails) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["segmentation_id"] = object.SegmentationId
	omap["physical_network"] = object.PhysicalNetwork

	return omap
}

func ResourceProviderDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"segmentation_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"physical_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceProviderDetails() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceProviderDetailsSchema(),
	}
}

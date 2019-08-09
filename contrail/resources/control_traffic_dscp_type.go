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

func SetControlTrafficDscpTypeFromMap(object *ControlTrafficDscpType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetControlTrafficDscpTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mControlObj := vmap["control"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mControlObj) {
		log.Printf("Setting control  Control <<%T>> => %#v", mControlObj, mControlObj)
		mControl := mControlObj.(int)
		object.Control = mControl
	}

	mAnalyticsObj := vmap["analytics"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAnalyticsObj) {
		log.Printf("Setting analytics  Analytics <<%T>> => %#v", mAnalyticsObj, mAnalyticsObj)
		mAnalytics := mAnalyticsObj.(int)
		object.Analytics = mAnalytics
	}

	mDnsObj := vmap["dns"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDnsObj) {
		log.Printf("Setting dns  Dns <<%T>> => %#v", mDnsObj, mDnsObj)
		mDns := mDnsObj.(int)
		object.Dns = mDns
	}

	log.Printf("FINISHED ControlTrafficDscpType object: %#v", object)
}

func TakeControlTrafficDscpTypeAsMap(object *ControlTrafficDscpType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["control"] = object.Control
	omap["analytics"] = object.Analytics
	omap["dns"] = object.Dns

	return omap
}

func ResourceControlTrafficDscpTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"control": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"analytics": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"dns": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceControlTrafficDscpType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceControlTrafficDscpTypeSchema(),
	}
}

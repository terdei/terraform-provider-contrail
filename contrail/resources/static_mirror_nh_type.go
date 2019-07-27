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

func SetStaticMirrorNhTypeFromMap(object *StaticMirrorNhType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetStaticMirrorNhTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mVtepDstIpAddressObj := vmap["vtep_dst_ip_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVtepDstIpAddressObj) {
		log.Printf("Setting vtep_dst_ip_address  VtepDstIpAddress <<%T>> => %#v", mVtepDstIpAddressObj, mVtepDstIpAddressObj)
		mVtepDstIpAddress := mVtepDstIpAddressObj.(string)
		object.VtepDstIpAddress = mVtepDstIpAddress
	}

	mVtepDstMacAddressObj := vmap["vtep_dst_mac_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVtepDstMacAddressObj) {
		log.Printf("Setting vtep_dst_mac_address  VtepDstMacAddress <<%T>> => %#v", mVtepDstMacAddressObj, mVtepDstMacAddressObj)
		mVtepDstMacAddress := mVtepDstMacAddressObj.(string)
		object.VtepDstMacAddress = mVtepDstMacAddress
	}

	mVniObj := vmap["vni"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVniObj) {
		log.Printf("Setting vni  Vni <<%T>> => %#v", mVniObj, mVniObj)
		mVni := mVniObj.(int)
		object.Vni = mVni
	}

	log.Printf("FINISHED StaticMirrorNhType object: %#v", object)
}

func TakeStaticMirrorNhTypeAsMap(object *StaticMirrorNhType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["vtep_dst_ip_address"] = object.VtepDstIpAddress
	omap["vtep_dst_mac_address"] = object.VtepDstMacAddress
	omap["vni"] = object.Vni

	return omap
}

func ResourceStaticMirrorNhTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vtep_dst_ip_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"vtep_dst_mac_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"vni": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceStaticMirrorNhType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceStaticMirrorNhTypeSchema(),
	}
}

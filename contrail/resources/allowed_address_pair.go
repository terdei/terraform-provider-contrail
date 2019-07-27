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

func SetAllowedAddressPairFromMap(object *AllowedAddressPair, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAllowedAddressPairFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mIpObj := vmap["ip"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mIpObj) {
		log.Printf("Setting ip  Ip <<%T>> => %#v", mIpObj, mIpObj)
		mIp := new(SubnetType)
		SetSubnetTypeFromMap(mIp, d, m, mIpObj)
		object.Ip = mIp
	}

	mMacObj := vmap["mac"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMacObj) {
		log.Printf("Setting mac  Mac <<%T>> => %#v", mMacObj, mMacObj)
		mMac := mMacObj.(string)
		object.Mac = mMac
	}

	mAddressModeObj := vmap["address_mode"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAddressModeObj) {
		log.Printf("Setting address_mode  AddressMode <<%T>> => %#v", mAddressModeObj, mAddressModeObj)
		mAddressMode := mAddressModeObj.(string)
		object.AddressMode = mAddressMode
	}

	log.Printf("FINISHED AllowedAddressPair object: %#v", object)
}

func TakeAllowedAddressPairAsMap(object *AllowedAddressPair) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Ip != nil {
		omap["ip"] = TakeSubnetTypeAsMap(object.Ip)
	}
	omap["mac"] = object.Mac
	omap["address_mode"] = object.AddressMode

	return omap
}

func ResourceAllowedAddressPairSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
		"mac": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"address_mode": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceAllowedAddressPair() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAllowedAddressPairSchema(),
	}
}

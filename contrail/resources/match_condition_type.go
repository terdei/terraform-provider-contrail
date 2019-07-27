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

func SetMatchConditionTypeFromMap(object *MatchConditionType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetMatchConditionTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mProtocolObj := vmap["protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolObj) {
		log.Printf("Setting protocol  Protocol <<%T>> => %#v", mProtocolObj, mProtocolObj)
		mProtocol := mProtocolObj.(string)
		object.Protocol = mProtocol
	}

	mSrcAddressObj := vmap["src_address"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSrcAddressObj) {
		log.Printf("Setting src_address  SrcAddress <<%T>> => %#v", mSrcAddressObj, mSrcAddressObj)
		mSrcAddress := new(AddressType)
		SetAddressTypeFromMap(mSrcAddress, d, m, mSrcAddressObj)
		object.SrcAddress = mSrcAddress
	}

	mSrcPortObj := vmap["src_port"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSrcPortObj) {
		log.Printf("Setting src_port  SrcPort <<%T>> => %#v", mSrcPortObj, mSrcPortObj)
		mSrcPort := new(PortType)
		SetPortTypeFromMap(mSrcPort, d, m, mSrcPortObj)
		object.SrcPort = mSrcPort
	}

	mDstAddressObj := vmap["dst_address"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mDstAddressObj) {
		log.Printf("Setting dst_address  DstAddress <<%T>> => %#v", mDstAddressObj, mDstAddressObj)
		mDstAddress := new(AddressType)
		SetAddressTypeFromMap(mDstAddress, d, m, mDstAddressObj)
		object.DstAddress = mDstAddress
	}

	mDstPortObj := vmap["dst_port"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mDstPortObj) {
		log.Printf("Setting dst_port  DstPort <<%T>> => %#v", mDstPortObj, mDstPortObj)
		mDstPort := new(PortType)
		SetPortTypeFromMap(mDstPort, d, m, mDstPortObj)
		object.DstPort = mDstPort
	}

	mEthertypeObj := vmap["ethertype"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEthertypeObj) {
		log.Printf("Setting ethertype  Ethertype <<%T>> => %#v", mEthertypeObj, mEthertypeObj)
		mEthertype := mEthertypeObj.(string)
		object.Ethertype = mEthertype
	}

	log.Printf("FINISHED MatchConditionType object: %#v", object)
}

func TakeMatchConditionTypeAsMap(object *MatchConditionType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["protocol"] = object.Protocol
	if object.SrcAddress != nil {
		omap["src_address"] = TakeAddressTypeAsMap(object.SrcAddress)
	}
	if object.SrcPort != nil {
		omap["src_port"] = TakePortTypeAsMap(object.SrcPort)
	}
	if object.DstAddress != nil {
		omap["dst_address"] = TakeAddressTypeAsMap(object.DstAddress)
	}
	if object.DstPort != nil {
		omap["dst_port"] = TakePortTypeAsMap(object.DstPort)
	}
	omap["ethertype"] = object.Ethertype

	return omap
}

func ResourceMatchConditionTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"src_address": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: AddressType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceAddressType(),
		},
		"src_port": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: PortType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortType(),
		},
		"dst_address": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: AddressType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceAddressType(),
		},
		"dst_port": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: PortType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortType(),
		},
		"ethertype": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceMatchConditionType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceMatchConditionTypeSchema(),
	}
}

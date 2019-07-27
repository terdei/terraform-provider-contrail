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

func SetVirtualNetworkTypeFromMap(object *VirtualNetworkType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualNetworkTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAllowTransitObj := vmap["allow_transit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAllowTransitObj) {
		log.Printf("Setting allow_transit  AllowTransit <<%T>> => %#v", mAllowTransitObj, mAllowTransitObj)
		mAllowTransit := mAllowTransitObj.(bool)
		object.AllowTransit = mAllowTransit
	}

	mNetworkIdObj := vmap["network_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNetworkIdObj) {
		log.Printf("Setting network_id  NetworkId <<%T>> => %#v", mNetworkIdObj, mNetworkIdObj)
		mNetworkId := mNetworkIdObj.(int)
		object.NetworkId = mNetworkId
	}

	mVxlanNetworkIdentifierObj := vmap["vxlan_network_identifier"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVxlanNetworkIdentifierObj) {
		log.Printf("Setting vxlan_network_identifier  VxlanNetworkIdentifier <<%T>> => %#v", mVxlanNetworkIdentifierObj, mVxlanNetworkIdentifierObj)
		mVxlanNetworkIdentifier := mVxlanNetworkIdentifierObj.(int)
		object.VxlanNetworkIdentifier = mVxlanNetworkIdentifier
	}

	mForwardingModeObj := vmap["forwarding_mode"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mForwardingModeObj) {
		log.Printf("Setting forwarding_mode  ForwardingMode <<%T>> => %#v", mForwardingModeObj, mForwardingModeObj)
		mForwardingMode := mForwardingModeObj.(string)
		object.ForwardingMode = mForwardingMode
	}

	mRpfObj := vmap["rpf"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRpfObj) {
		log.Printf("Setting rpf  Rpf <<%T>> => %#v", mRpfObj, mRpfObj)
		mRpf := mRpfObj.(string)
		object.Rpf = mRpf
	}

	mMirrorDestinationObj := vmap["mirror_destination"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMirrorDestinationObj) {
		log.Printf("Setting mirror_destination  MirrorDestination <<%T>> => %#v", mMirrorDestinationObj, mMirrorDestinationObj)
		mMirrorDestination := mMirrorDestinationObj.(bool)
		object.MirrorDestination = mMirrorDestination
	}

	log.Printf("FINISHED VirtualNetworkType object: %#v", object)
}

func TakeVirtualNetworkTypeAsMap(object *VirtualNetworkType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["allow_transit"] = object.AllowTransit
	omap["network_id"] = object.NetworkId
	omap["vxlan_network_identifier"] = object.VxlanNetworkIdentifier
	omap["forwarding_mode"] = object.ForwardingMode
	omap["rpf"] = object.Rpf
	omap["mirror_destination"] = object.MirrorDestination

	return omap
}

func ResourceVirtualNetworkTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_transit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"network_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"vxlan_network_identifier": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"forwarding_mode": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"rpf": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"mirror_destination": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceVirtualNetworkType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualNetworkTypeSchema(),
	}
}

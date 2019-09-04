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

func SetPolicyBasedForwardingRuleTypeFromMap(object *PolicyBasedForwardingRuleType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPolicyBasedForwardingRuleTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mDirectionObj := vmap["direction"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDirectionObj) {
		log.Printf("Setting direction  Direction <<%T>> => %#v", mDirectionObj, mDirectionObj)
		mDirection := mDirectionObj.(string)
		object.Direction = mDirection
	}

	mVlanTagObj := vmap["vlan_tag"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVlanTagObj) {
		log.Printf("Setting vlan_tag  VlanTag <<%T>> => %#v", mVlanTagObj, mVlanTagObj)
		mVlanTag := mVlanTagObj.(int)
		object.VlanTag = mVlanTag
	}

	mSrcMacObj := vmap["src_mac"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSrcMacObj) {
		log.Printf("Setting src_mac  SrcMac <<%T>> => %#v", mSrcMacObj, mSrcMacObj)
		mSrcMac := mSrcMacObj.(string)
		object.SrcMac = mSrcMac
	}

	mDstMacObj := vmap["dst_mac"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDstMacObj) {
		log.Printf("Setting dst_mac  DstMac <<%T>> => %#v", mDstMacObj, mDstMacObj)
		mDstMac := mDstMacObj.(string)
		object.DstMac = mDstMac
	}

	mMplsLabelObj := vmap["mpls_label"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMplsLabelObj) {
		log.Printf("Setting mpls_label  MplsLabel <<%T>> => %#v", mMplsLabelObj, mMplsLabelObj)
		mMplsLabel := mMplsLabelObj.(int)
		object.MplsLabel = mMplsLabel
	}

	mServiceChainAddressObj := vmap["service_chain_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceChainAddressObj) {
		log.Printf("Setting service_chain_address  ServiceChainAddress <<%T>> => %#v", mServiceChainAddressObj, mServiceChainAddressObj)
		mServiceChainAddress := mServiceChainAddressObj.(string)
		object.ServiceChainAddress = mServiceChainAddress
	}

	mIpv6ServiceChainAddressObj := vmap["ipv6_service_chain_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpv6ServiceChainAddressObj) {
		log.Printf("Setting ipv6_service_chain_address  Ipv6ServiceChainAddress <<%T>> => %#v", mIpv6ServiceChainAddressObj, mIpv6ServiceChainAddressObj)
		mIpv6ServiceChainAddress := mIpv6ServiceChainAddressObj.(string)
		object.Ipv6ServiceChainAddress = mIpv6ServiceChainAddress
	}

	mProtocolObj := vmap["protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolObj) {
		log.Printf("Setting protocol  Protocol <<%T>> => %#v", mProtocolObj, mProtocolObj)
		mProtocol := mProtocolObj.(string)
		object.Protocol = mProtocol
	}

	log.Printf("FINISHED PolicyBasedForwardingRuleType object: %#v", object)
}

func TakePolicyBasedForwardingRuleTypeAsMap(object *PolicyBasedForwardingRuleType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["direction"] = object.Direction
	omap["vlan_tag"] = object.VlanTag
	omap["src_mac"] = object.SrcMac
	omap["dst_mac"] = object.DstMac
	omap["mpls_label"] = object.MplsLabel
	omap["service_chain_address"] = object.ServiceChainAddress
	omap["ipv6_service_chain_address"] = object.Ipv6ServiceChainAddress
	omap["protocol"] = object.Protocol

	return omap
}

func ResourcePolicyBasedForwardingRuleTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"direction": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"vlan_tag": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"src_mac": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"dst_mac": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"mpls_label": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"service_chain_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ipv6_service_chain_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourcePolicyBasedForwardingRuleType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePolicyBasedForwardingRuleTypeSchema(),
	}
}

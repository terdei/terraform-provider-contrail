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

func SetAddressTypeFromMap(object *AddressType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAddressTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSubnetObj := vmap["subnet"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSubnetObj) {
		log.Printf("Setting subnet  Subnet <<%T>> => %#v", mSubnetObj, mSubnetObj)
		mSubnet := new(SubnetType)
		SetSubnetTypeFromMap(mSubnet, d, m, mSubnetObj)
		object.Subnet = mSubnet
	}

	mVirtualNetworkObj := vmap["virtual_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualNetworkObj) {
		log.Printf("Setting virtual_network  VirtualNetwork <<%T>> => %#v", mVirtualNetworkObj, mVirtualNetworkObj)
		mVirtualNetwork := mVirtualNetworkObj.(string)
		object.VirtualNetwork = mVirtualNetwork
	}

	mSecurityGroupObj := vmap["security_group"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSecurityGroupObj) {
		log.Printf("Setting security_group  SecurityGroup <<%T>> => %#v", mSecurityGroupObj, mSecurityGroupObj)
		mSecurityGroup := mSecurityGroupObj.(string)
		object.SecurityGroup = mSecurityGroup
	}

	mNetworkPolicyObj := vmap["network_policy"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNetworkPolicyObj) {
		log.Printf("Setting network_policy  NetworkPolicy <<%T>> => %#v", mNetworkPolicyObj, mNetworkPolicyObj)
		mNetworkPolicy := mNetworkPolicyObj.(string)
		object.NetworkPolicy = mNetworkPolicy
	}

	mSubnetListObj := vmap["subnet_list"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSubnetListObj) {
		log.Printf("Setting subnet_list  SubnetList <<%T>> => %#v", mSubnetListObj, mSubnetListObj)
		for _, v := range mSubnetListObj.([]interface{}) {
			mSubnetList := new(SubnetType)
			SetSubnetTypeFromMap(mSubnetList, d, m, v)
			object.AddSubnetList(mSubnetList)
		}
	}

	log.Printf("FINISHED AddressType object: %#v", object)
}

func TakeAddressTypeAsMap(object *AddressType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Subnet != nil {
		omap["subnet"] = TakeSubnetTypeAsMap(object.Subnet)
	}
	omap["virtual_network"] = object.VirtualNetwork
	omap["security_group"] = object.SecurityGroup
	omap["network_policy"] = object.NetworkPolicy
	var subnet_list_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.SubnetList {
		subnet_list_map = append(subnet_list_map, TakeSubnetTypeAsMap(&v))
	}
	omap["subnet_list"] = subnet_list_map

	return omap
}

func ResourceAddressTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
		"virtual_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"security_group": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"network_policy": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"subnet_list": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
	}
}

func ResourceAddressType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAddressTypeSchema(),
	}
}

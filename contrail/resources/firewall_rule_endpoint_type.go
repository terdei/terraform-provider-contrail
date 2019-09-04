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

func SetFirewallRuleEndpointTypeFromMap(object *FirewallRuleEndpointType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFirewallRuleEndpointTypeFromMAP] map = %#v", val)

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

	mAddressGroupObj := vmap["address_group"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAddressGroupObj) {
		log.Printf("Setting address_group  AddressGroup <<%T>> => %#v", mAddressGroupObj, mAddressGroupObj)
		mAddressGroup := mAddressGroupObj.(string)
		object.AddressGroup = mAddressGroup
	}

	mTagsObj := vmap["tags"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mTagsObj) {
		log.Printf("Setting tags  Tags <<%T>> => %#v", mTagsObj, mTagsObj)
		for _, v := range mTagsObj.([]interface{}) {
			mTags := v.(string)
			object.AddTags(mTags)
		}
	}

	mTagIdsObj := vmap["tag_ids"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mTagIdsObj) {
		log.Printf("Setting tag_ids  TagIds <<%T>> => %#v", mTagIdsObj, mTagIdsObj)
		for _, v := range mTagIdsObj.([]interface{}) {
			mTagIds := v.(int)
			object.AddTagIds(mTagIds)
		}
	}

	mAnyObj := vmap["any"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAnyObj) {
		log.Printf("Setting any  Any <<%T>> => %#v", mAnyObj, mAnyObj)
		mAny := mAnyObj.(bool)
		object.Any = mAny
	}

	log.Printf("FINISHED FirewallRuleEndpointType object: %#v", object)
}

func TakeFirewallRuleEndpointTypeAsMap(object *FirewallRuleEndpointType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Subnet != nil {
		omap["subnet"] = TakeSubnetTypeAsMap(object.Subnet)
	}
	omap["virtual_network"] = object.VirtualNetwork
	omap["address_group"] = object.AddressGroup
	tags_arr := make([]string, len(object.Tags))
	for _, v := range object.Tags {
		tags_arr = append(tags_arr, v)
	}
	omap["tags"] = tags_arr
	tag_ids_arr := make([]int, len(object.TagIds))
	for _, v := range object.TagIds {
		tag_ids_arr = append(tag_ids_arr, v)
	}
	omap["tag_ids"] = tag_ids_arr
	omap["any"] = object.Any

	return omap
}

func ResourceFirewallRuleEndpointTypeSchema() map[string]*schema.Schema {
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
		"address_group": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"tags": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_ids": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"any": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceFirewallRuleEndpointType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFirewallRuleEndpointTypeSchema(),
	}
}

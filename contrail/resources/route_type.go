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

func SetRouteTypeFromMap(object *RouteType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRouteTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPrefixObj := vmap["prefix"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPrefixObj) {
		log.Printf("Setting prefix  Prefix <<%T>> => %#v", mPrefixObj, mPrefixObj)
		mPrefix := mPrefixObj.(string)
		object.Prefix = mPrefix
	}

	mNextHopObj := vmap["next_hop"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNextHopObj) {
		log.Printf("Setting next_hop  NextHop <<%T>> => %#v", mNextHopObj, mNextHopObj)
		mNextHop := mNextHopObj.(string)
		object.NextHop = mNextHop
	}

	mNextHopTypeObj := vmap["next_hop_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNextHopTypeObj) {
		log.Printf("Setting next_hop_type  NextHopType <<%T>> => %#v", mNextHopTypeObj, mNextHopTypeObj)
		mNextHopType := mNextHopTypeObj.(string)
		object.NextHopType = mNextHopType
	}

	mCommunityAttributesObj := vmap["community_attributes"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mCommunityAttributesObj) {
		log.Printf("Setting community_attributes  CommunityAttributes <<%T>> => %#v", mCommunityAttributesObj, mCommunityAttributesObj)
		mCommunityAttributes := new(CommunityAttributes)
		SetCommunityAttributesFromMap(mCommunityAttributes, d, m, mCommunityAttributesObj)
		object.CommunityAttributes = mCommunityAttributes
	}

	log.Printf("FINISHED RouteType object: %#v", object)
}

func TakeRouteTypeAsMap(object *RouteType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["prefix"] = object.Prefix
	omap["next_hop"] = object.NextHop
	omap["next_hop_type"] = object.NextHopType
	if object.CommunityAttributes != nil {
		omap["community_attributes"] = TakeCommunityAttributesAsMap(object.CommunityAttributes)
	}

	return omap
}

func ResourceRouteTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prefix": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"next_hop": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"next_hop_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"community_attributes": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: CommunityAttributes
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceCommunityAttributes(),
		},
	}
}

func ResourceRouteType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRouteTypeSchema(),
	}
}

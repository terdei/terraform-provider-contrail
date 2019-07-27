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

func SetCommunityAttributesFromMap(object *CommunityAttributes, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetCommunityAttributesFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mCommunityAttributeObj := vmap["community_attribute"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mCommunityAttributeObj) {
		log.Printf("Setting community_attribute  CommunityAttribute <<%T>> => %#v", mCommunityAttributeObj, mCommunityAttributeObj)
		for _, v := range mCommunityAttributeObj.([]interface{}) {
			mCommunityAttribute := v.(string)
			object.AddCommunityAttribute(mCommunityAttribute)
		}
	}

	log.Printf("FINISHED CommunityAttributes object: %#v", object)
}

func TakeCommunityAttributesAsMap(object *CommunityAttributes) map[string]interface{} {
	omap := make(map[string]interface{})

	community_attribute_arr := make([]string, len(object.CommunityAttribute))
	for _, v := range object.CommunityAttribute {
		community_attribute_arr = append(community_attribute_arr, v)
	}
	omap["community_attribute"] = community_attribute_arr

	return omap
}

func ResourceCommunityAttributesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"community_attribute": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceCommunityAttributes() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceCommunityAttributesSchema(),
	}
}

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

func SetFirewallRuleMatchTagsTypeIdListFromMap(object *FirewallRuleMatchTagsTypeIdList, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFirewallRuleMatchTagsTypeIdListFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mTagTypeObj := vmap["tag_type"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mTagTypeObj) {
		log.Printf("Setting tag_type  TagType <<%T>> => %#v", mTagTypeObj, mTagTypeObj)
		for _, v := range mTagTypeObj.([]interface{}) {
			mTagType := v.(int)
			object.AddTagType(mTagType)
		}
	}

	log.Printf("FINISHED FirewallRuleMatchTagsTypeIdList object: %#v", object)
}

func TakeFirewallRuleMatchTagsTypeIdListAsMap(object *FirewallRuleMatchTagsTypeIdList) map[string]interface{} {
	omap := make(map[string]interface{})

	tag_type_arr := make([]int, len(object.TagType))
	for _, v := range object.TagType {
		tag_type_arr = append(tag_type_arr, v)
	}
	omap["tag_type"] = tag_type_arr

	return omap
}

func ResourceFirewallRuleMatchTagsTypeIdListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tag_type": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:integer
			Required: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
	}
}

func ResourceFirewallRuleMatchTagsTypeIdList() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFirewallRuleMatchTagsTypeIdListSchema(),
	}
}

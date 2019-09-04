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

func SetFirewallRuleMatchTagsTypeFromMap(object *FirewallRuleMatchTagsType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFirewallRuleMatchTagsTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mTagListObj := vmap["tag_list"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mTagListObj) {
		log.Printf("Setting tag_list  TagList <<%T>> => %#v", mTagListObj, mTagListObj)
		for _, v := range mTagListObj.([]interface{}) {
			mTagList := v.(string)
			object.AddTagList(mTagList)
		}
	}

	log.Printf("FINISHED FirewallRuleMatchTagsType object: %#v", object)
}

func TakeFirewallRuleMatchTagsTypeAsMap(object *FirewallRuleMatchTagsType) map[string]interface{} {
	omap := make(map[string]interface{})

	tag_list_arr := make([]string, len(object.TagList))
	for _, v := range object.TagList {
		tag_list_arr = append(tag_list_arr, v)
	}
	omap["tag_list"] = tag_list_arr

	return omap
}

func ResourceFirewallRuleMatchTagsTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tag_list": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Required: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceFirewallRuleMatchTagsType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFirewallRuleMatchTagsTypeSchema(),
	}
}

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

func SetUveKeysTypeFromMap(object *UveKeysType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetUveKeysTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mUveKeyObj := vmap["uve_key"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mUveKeyObj) {
		log.Printf("Setting uve_key  UveKey <<%T>> => %#v", mUveKeyObj, mUveKeyObj)
		for _, v := range mUveKeyObj.([]interface{}) {
			mUveKey := v.(string)
			object.AddUveKey(mUveKey)
		}
	}

	log.Printf("FINISHED UveKeysType object: %#v", object)
}

func TakeUveKeysTypeAsMap(object *UveKeysType) map[string]interface{} {
	omap := make(map[string]interface{})

	uve_key_arr := make([]string, len(object.UveKey))
	for _, v := range object.UveKey {
		uve_key_arr = append(uve_key_arr, v)
	}
	omap["uve_key"] = uve_key_arr

	return omap
}

func ResourceUveKeysTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uve_key": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Required: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceUveKeysType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceUveKeysTypeSchema(),
	}
}

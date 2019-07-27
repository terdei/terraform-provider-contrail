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

func SetKeyValuePairFromMap(object *KeyValuePair, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetKeyValuePairFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mKeyObj := vmap["key"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mKeyObj) {
		log.Printf("Setting key  Key <<%T>> => %#v", mKeyObj, mKeyObj)
		mKey := mKeyObj.(string)
		object.Key = mKey
	}

	mValueObj := vmap["value"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mValueObj) {
		log.Printf("Setting value  Value <<%T>> => %#v", mValueObj, mValueObj)
		mValue := mValueObj.(string)
		object.Value = mValue
	}

	log.Printf("FINISHED KeyValuePair object: %#v", object)
}

func TakeKeyValuePairAsMap(object *KeyValuePair) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["key"] = object.Key
	omap["value"] = object.Value

	return omap
}

func ResourceKeyValuePairSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"value": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceKeyValuePair() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceKeyValuePairSchema(),
	}
}

func SetKeyValuePairsFromMap(object *KeyValuePairs, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetKeyValuePairsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mKeyValuePairObj := vmap["key_value_pair"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mKeyValuePairObj) {
		log.Printf("Setting key_value_pair  KeyValuePair <<%T>> => %#v", mKeyValuePairObj, mKeyValuePairObj)
		for _, v := range mKeyValuePairObj.([]interface{}) {
			mKeyValuePair := new(KeyValuePair)
			SetKeyValuePairFromMap(mKeyValuePair, d, m, v)
			object.AddKeyValuePair(mKeyValuePair)
		}
	}

	log.Printf("FINISHED KeyValuePairs object: %#v", object)
}

func TakeKeyValuePairsAsMap(object *KeyValuePairs) map[string]interface{} {
	omap := make(map[string]interface{})

	var key_value_pair_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.KeyValuePair {
		key_value_pair_map = append(key_value_pair_map, TakeKeyValuePairAsMap(&v))
	}
	omap["key_value_pair"] = key_value_pair_map

	return omap
}

func ResourceKeyValuePairsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key_value_pair": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: KeyValuePair
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceKeyValuePair(),
		},
	}
}

func ResourceKeyValuePairs() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceKeyValuePairsSchema(),
	}
}

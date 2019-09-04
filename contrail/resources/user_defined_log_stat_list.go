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

func SetUserDefinedLogStatFromMap(object *UserDefinedLogStat, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetUserDefinedLogStatFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mNameObj := vmap["name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNameObj) {
		log.Printf("Setting name  Name <<%T>> => %#v", mNameObj, mNameObj)
		mName := mNameObj.(string)
		object.Name = mName
	}

	mPatternObj := vmap["pattern"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPatternObj) {
		log.Printf("Setting pattern  Pattern <<%T>> => %#v", mPatternObj, mPatternObj)
		mPattern := mPatternObj.(string)
		object.Pattern = mPattern
	}

	log.Printf("FINISHED UserDefinedLogStat object: %#v", object)
}

func TakeUserDefinedLogStatAsMap(object *UserDefinedLogStat) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["name"] = object.Name
	omap["pattern"] = object.Pattern

	return omap
}

func ResourceUserDefinedLogStatSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"pattern": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceUserDefinedLogStat() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceUserDefinedLogStatSchema(),
	}
}

func SetUserDefinedLogStatListFromMap(object *UserDefinedLogStatList, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetUserDefinedLogStatListFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mStatlistObj := vmap["statlist"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mStatlistObj) {
		log.Printf("Setting statlist  Statlist <<%T>> => %#v", mStatlistObj, mStatlistObj)
		for _, v := range mStatlistObj.([]interface{}) {
			mStatlist := new(UserDefinedLogStat)
			SetUserDefinedLogStatFromMap(mStatlist, d, m, v)
			object.AddStatlist(mStatlist)
		}
	}

	log.Printf("FINISHED UserDefinedLogStatList object: %#v", object)
}

func TakeUserDefinedLogStatListAsMap(object *UserDefinedLogStatList) map[string]interface{} {
	omap := make(map[string]interface{})

	var statlist_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Statlist {
		statlist_map = append(statlist_map, TakeUserDefinedLogStatAsMap(&v))
	}
	omap["statlist"] = statlist_map

	return omap
}

func ResourceUserDefinedLogStatListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"statlist": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: UserDefinedLogStat
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceUserDefinedLogStat(),
		},
	}
}

func ResourceUserDefinedLogStatList() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceUserDefinedLogStatListSchema(),
	}
}

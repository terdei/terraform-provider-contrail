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

func SetPluginPropertyFromMap(object *PluginProperty, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPluginPropertyFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPropertyObj := vmap["property"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPropertyObj) {
		log.Printf("Setting property  Property <<%T>> => %#v", mPropertyObj, mPropertyObj)
		mProperty := mPropertyObj.(string)
		object.Property = mProperty
	}

	mValueObj := vmap["value"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mValueObj) {
		log.Printf("Setting value  Value <<%T>> => %#v", mValueObj, mValueObj)
		mValue := mValueObj.(string)
		object.Value = mValue
	}

	log.Printf("FINISHED PluginProperty object: %#v", object)
}

func TakePluginPropertyAsMap(object *PluginProperty) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["property"] = object.Property
	omap["value"] = object.Value

	return omap
}

func ResourcePluginPropertySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"property": &schema.Schema{
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

func ResourcePluginProperty() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePluginPropertySchema(),
	}
}

func SetPluginPropertiesFromMap(object *PluginProperties, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPluginPropertiesFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPluginPropertyObj := vmap["plugin_property"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mPluginPropertyObj) {
		log.Printf("Setting plugin_property  PluginProperty <<%T>> => %#v", mPluginPropertyObj, mPluginPropertyObj)
		for _, v := range mPluginPropertyObj.([]interface{}) {
			mPluginProperty := new(PluginProperty)
			SetPluginPropertyFromMap(mPluginProperty, d, m, v)
			object.AddPluginProperty(mPluginProperty)
		}
	}

	log.Printf("FINISHED PluginProperties object: %#v", object)
}

func TakePluginPropertiesAsMap(object *PluginProperties) map[string]interface{} {
	omap := make(map[string]interface{})

	var plugin_property_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.PluginProperty {
		plugin_property_map = append(plugin_property_map, TakePluginPropertyAsMap(&v))
	}
	omap["plugin_property"] = plugin_property_map

	return omap
}

func ResourcePluginPropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"plugin_property": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: PluginProperty
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePluginProperty(),
		},
	}
}

func ResourcePluginProperties() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePluginPropertiesSchema(),
	}
}

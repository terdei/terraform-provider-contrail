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

func SetPortMapFromMap(object *PortMap, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPortMapFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSrcPortObj := vmap["src_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSrcPortObj) {
		log.Printf("Setting src_port  SrcPort <<%T>> => %#v", mSrcPortObj, mSrcPortObj)
		mSrcPort := mSrcPortObj.(int)
		object.SrcPort = mSrcPort
	}

	mDstPortObj := vmap["dst_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDstPortObj) {
		log.Printf("Setting dst_port  DstPort <<%T>> => %#v", mDstPortObj, mDstPortObj)
		mDstPort := mDstPortObj.(int)
		object.DstPort = mDstPort
	}

	log.Printf("FINISHED PortMap object: %#v", object)
}

func TakePortMapAsMap(object *PortMap) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["src_port"] = object.SrcPort
	omap["dst_port"] = object.DstPort

	return omap
}

func ResourcePortMapSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"src_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"dst_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourcePortMap() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePortMapSchema(),
	}
}

func SetPortMappingsFromMap(object *PortMappings, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPortMappingsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPortMappingsObj := vmap["port_mappings"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mPortMappingsObj) {
		log.Printf("Setting port_mappings  PortMappings <<%T>> => %#v", mPortMappingsObj, mPortMappingsObj)
		for _, v := range mPortMappingsObj.([]interface{}) {
			mPortMappings := new(PortMap)
			SetPortMapFromMap(mPortMappings, d, m, v)
			object.AddPortMappings(mPortMappings)
		}
	}

	log.Printf("FINISHED PortMappings object: %#v", object)
}

func TakePortMappingsAsMap(object *PortMappings) map[string]interface{} {
	omap := make(map[string]interface{})

	var port_mappings_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.PortMappings {
		port_mappings_map = append(port_mappings_map, TakePortMapAsMap(&v))
	}
	omap["port_mappings"] = port_mappings_map

	return omap
}

func ResourcePortMappingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port_mappings": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: PortMap
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortMap(),
		},
	}
}

func ResourcePortMappings() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePortMappingsSchema(),
	}
}

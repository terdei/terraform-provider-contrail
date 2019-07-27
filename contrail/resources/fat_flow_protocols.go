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

func SetProtocolTypeFromMap(object *ProtocolType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetProtocolTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mProtocolObj := vmap["protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolObj) {
		log.Printf("Setting protocol  Protocol <<%T>> => %#v", mProtocolObj, mProtocolObj)
		mProtocol := mProtocolObj.(string)
		object.Protocol = mProtocol
	}

	mPortObj := vmap["port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPortObj) {
		log.Printf("Setting port  Port <<%T>> => %#v", mPortObj, mPortObj)
		mPort := mPortObj.(int)
		object.Port = mPort
	}

	log.Printf("FINISHED ProtocolType object: %#v", object)
}

func TakeProtocolTypeAsMap(object *ProtocolType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["protocol"] = object.Protocol
	omap["port"] = object.Port

	return omap
}

func ResourceProtocolTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceProtocolType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceProtocolTypeSchema(),
	}
}

func SetFatFlowProtocolsFromMap(object *FatFlowProtocols, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFatFlowProtocolsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mFatFlowProtocolObj := vmap["fat_flow_protocol"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mFatFlowProtocolObj) {
		log.Printf("Setting fat_flow_protocol  FatFlowProtocol <<%T>> => %#v", mFatFlowProtocolObj, mFatFlowProtocolObj)
		for _, v := range mFatFlowProtocolObj.([]interface{}) {
			mFatFlowProtocol := new(ProtocolType)
			SetProtocolTypeFromMap(mFatFlowProtocol, d, m, v)
			object.AddFatFlowProtocol(mFatFlowProtocol)
		}
	}

	log.Printf("FINISHED FatFlowProtocols object: %#v", object)
}

func TakeFatFlowProtocolsAsMap(object *FatFlowProtocols) map[string]interface{} {
	omap := make(map[string]interface{})

	var fat_flow_protocol_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.FatFlowProtocol {
		fat_flow_protocol_map = append(fat_flow_protocol_map, TakeProtocolTypeAsMap(&v))
	}
	omap["fat_flow_protocol"] = fat_flow_protocol_map

	return omap
}

func ResourceFatFlowProtocolsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"fat_flow_protocol": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: ProtocolType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceProtocolType(),
		},
	}
}

func ResourceFatFlowProtocols() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFatFlowProtocolsSchema(),
	}
}

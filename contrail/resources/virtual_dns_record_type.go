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

func SetVirtualDnsRecordTypeFromMap(object *VirtualDnsRecordType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualDnsRecordTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRecordNameObj := vmap["record_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordNameObj) {
		log.Printf("Setting record_name  RecordName <<%T>> => %#v", mRecordNameObj, mRecordNameObj)
		mRecordName := mRecordNameObj.(string)
		object.RecordName = mRecordName
	}

	mRecordTypeObj := vmap["record_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordTypeObj) {
		log.Printf("Setting record_type  RecordType <<%T>> => %#v", mRecordTypeObj, mRecordTypeObj)
		mRecordType := mRecordTypeObj.(string)
		object.RecordType = mRecordType
	}

	mRecordClassObj := vmap["record_class"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordClassObj) {
		log.Printf("Setting record_class  RecordClass <<%T>> => %#v", mRecordClassObj, mRecordClassObj)
		mRecordClass := mRecordClassObj.(string)
		object.RecordClass = mRecordClass
	}

	mRecordDataObj := vmap["record_data"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordDataObj) {
		log.Printf("Setting record_data  RecordData <<%T>> => %#v", mRecordDataObj, mRecordDataObj)
		mRecordData := mRecordDataObj.(string)
		object.RecordData = mRecordData
	}

	mRecordTtlSecondsObj := vmap["record_ttl_seconds"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordTtlSecondsObj) {
		log.Printf("Setting record_ttl_seconds  RecordTtlSeconds <<%T>> => %#v", mRecordTtlSecondsObj, mRecordTtlSecondsObj)
		mRecordTtlSeconds := mRecordTtlSecondsObj.(int)
		object.RecordTtlSeconds = mRecordTtlSeconds
	}

	mRecordMxPreferenceObj := vmap["record_mx_preference"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordMxPreferenceObj) {
		log.Printf("Setting record_mx_preference  RecordMxPreference <<%T>> => %#v", mRecordMxPreferenceObj, mRecordMxPreferenceObj)
		mRecordMxPreference := mRecordMxPreferenceObj.(int)
		object.RecordMxPreference = mRecordMxPreference
	}

	log.Printf("FINISHED VirtualDnsRecordType object: %#v", object)
}

func TakeVirtualDnsRecordTypeAsMap(object *VirtualDnsRecordType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["record_name"] = object.RecordName
	omap["record_type"] = object.RecordType
	omap["record_class"] = object.RecordClass
	omap["record_data"] = object.RecordData
	omap["record_ttl_seconds"] = object.RecordTtlSeconds
	omap["record_mx_preference"] = object.RecordMxPreference

	return omap
}

func ResourceVirtualDnsRecordTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"record_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"record_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"record_class": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"record_data": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"record_ttl_seconds": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"record_mx_preference": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceVirtualDnsRecordType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualDnsRecordTypeSchema(),
	}
}

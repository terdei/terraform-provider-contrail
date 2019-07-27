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

func SetEcmpHashingIncludeFieldsFromMap(object *EcmpHashingIncludeFields, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetEcmpHashingIncludeFieldsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mHashingConfiguredObj := vmap["hashing_configured"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mHashingConfiguredObj) {
		log.Printf("Setting hashing_configured  HashingConfigured <<%T>> => %#v", mHashingConfiguredObj, mHashingConfiguredObj)
		mHashingConfigured := mHashingConfiguredObj.(bool)
		object.HashingConfigured = mHashingConfigured
	}

	mSourceIpObj := vmap["source_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSourceIpObj) {
		log.Printf("Setting source_ip  SourceIp <<%T>> => %#v", mSourceIpObj, mSourceIpObj)
		mSourceIp := mSourceIpObj.(bool)
		object.SourceIp = mSourceIp
	}

	mDestinationIpObj := vmap["destination_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDestinationIpObj) {
		log.Printf("Setting destination_ip  DestinationIp <<%T>> => %#v", mDestinationIpObj, mDestinationIpObj)
		mDestinationIp := mDestinationIpObj.(bool)
		object.DestinationIp = mDestinationIp
	}

	mIpProtocolObj := vmap["ip_protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpProtocolObj) {
		log.Printf("Setting ip_protocol  IpProtocol <<%T>> => %#v", mIpProtocolObj, mIpProtocolObj)
		mIpProtocol := mIpProtocolObj.(bool)
		object.IpProtocol = mIpProtocol
	}

	mSourcePortObj := vmap["source_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSourcePortObj) {
		log.Printf("Setting source_port  SourcePort <<%T>> => %#v", mSourcePortObj, mSourcePortObj)
		mSourcePort := mSourcePortObj.(bool)
		object.SourcePort = mSourcePort
	}

	mDestinationPortObj := vmap["destination_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDestinationPortObj) {
		log.Printf("Setting destination_port  DestinationPort <<%T>> => %#v", mDestinationPortObj, mDestinationPortObj)
		mDestinationPort := mDestinationPortObj.(bool)
		object.DestinationPort = mDestinationPort
	}

	log.Printf("FINISHED EcmpHashingIncludeFields object: %#v", object)
}

func TakeEcmpHashingIncludeFieldsAsMap(object *EcmpHashingIncludeFields) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["hashing_configured"] = object.HashingConfigured
	omap["source_ip"] = object.SourceIp
	omap["destination_ip"] = object.DestinationIp
	omap["ip_protocol"] = object.IpProtocol
	omap["source_port"] = object.SourcePort
	omap["destination_port"] = object.DestinationPort

	return omap
}

func ResourceEcmpHashingIncludeFieldsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hashing_configured": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"source_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"destination_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"ip_protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"source_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"destination_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceEcmpHashingIncludeFields() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceEcmpHashingIncludeFieldsSchema(),
	}
}

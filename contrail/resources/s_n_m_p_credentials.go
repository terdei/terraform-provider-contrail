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

func SetSNMPCredentialsFromMap(object *SNMPCredentials, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSNMPCredentialsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mVersionObj := vmap["version"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVersionObj) {
		log.Printf("Setting version  Version <<%T>> => %#v", mVersionObj, mVersionObj)
		mVersion := mVersionObj.(int)
		object.Version = mVersion
	}

	mLocalPortObj := vmap["local_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLocalPortObj) {
		log.Printf("Setting local_port  LocalPort <<%T>> => %#v", mLocalPortObj, mLocalPortObj)
		mLocalPort := mLocalPortObj.(int)
		object.LocalPort = mLocalPort
	}

	mRetriesObj := vmap["retries"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRetriesObj) {
		log.Printf("Setting retries  Retries <<%T>> => %#v", mRetriesObj, mRetriesObj)
		mRetries := mRetriesObj.(int)
		object.Retries = mRetries
	}

	mTimeoutObj := vmap["timeout"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTimeoutObj) {
		log.Printf("Setting timeout  Timeout <<%T>> => %#v", mTimeoutObj, mTimeoutObj)
		mTimeout := mTimeoutObj.(int)
		object.Timeout = mTimeout
	}

	mV2CommunityObj := vmap["v2_community"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV2CommunityObj) {
		log.Printf("Setting v2_community  V2Community <<%T>> => %#v", mV2CommunityObj, mV2CommunityObj)
		mV2Community := mV2CommunityObj.(string)
		object.V2Community = mV2Community
	}

	mV3SecurityNameObj := vmap["v3_security_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3SecurityNameObj) {
		log.Printf("Setting v3_security_name  V3SecurityName <<%T>> => %#v", mV3SecurityNameObj, mV3SecurityNameObj)
		mV3SecurityName := mV3SecurityNameObj.(string)
		object.V3SecurityName = mV3SecurityName
	}

	mV3SecurityLevelObj := vmap["v3_security_level"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3SecurityLevelObj) {
		log.Printf("Setting v3_security_level  V3SecurityLevel <<%T>> => %#v", mV3SecurityLevelObj, mV3SecurityLevelObj)
		mV3SecurityLevel := mV3SecurityLevelObj.(string)
		object.V3SecurityLevel = mV3SecurityLevel
	}

	mV3SecurityEngineIdObj := vmap["v3_security_engine_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3SecurityEngineIdObj) {
		log.Printf("Setting v3_security_engine_id  V3SecurityEngineId <<%T>> => %#v", mV3SecurityEngineIdObj, mV3SecurityEngineIdObj)
		mV3SecurityEngineId := mV3SecurityEngineIdObj.(string)
		object.V3SecurityEngineId = mV3SecurityEngineId
	}

	mV3ContextObj := vmap["v3_context"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3ContextObj) {
		log.Printf("Setting v3_context  V3Context <<%T>> => %#v", mV3ContextObj, mV3ContextObj)
		mV3Context := mV3ContextObj.(string)
		object.V3Context = mV3Context
	}

	mV3ContextEngineIdObj := vmap["v3_context_engine_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3ContextEngineIdObj) {
		log.Printf("Setting v3_context_engine_id  V3ContextEngineId <<%T>> => %#v", mV3ContextEngineIdObj, mV3ContextEngineIdObj)
		mV3ContextEngineId := mV3ContextEngineIdObj.(string)
		object.V3ContextEngineId = mV3ContextEngineId
	}

	mV3AuthenticationProtocolObj := vmap["v3_authentication_protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3AuthenticationProtocolObj) {
		log.Printf("Setting v3_authentication_protocol  V3AuthenticationProtocol <<%T>> => %#v", mV3AuthenticationProtocolObj, mV3AuthenticationProtocolObj)
		mV3AuthenticationProtocol := mV3AuthenticationProtocolObj.(string)
		object.V3AuthenticationProtocol = mV3AuthenticationProtocol
	}

	mV3AuthenticationPasswordObj := vmap["v3_authentication_password"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3AuthenticationPasswordObj) {
		log.Printf("Setting v3_authentication_password  V3AuthenticationPassword <<%T>> => %#v", mV3AuthenticationPasswordObj, mV3AuthenticationPasswordObj)
		mV3AuthenticationPassword := mV3AuthenticationPasswordObj.(string)
		object.V3AuthenticationPassword = mV3AuthenticationPassword
	}

	mV3PrivacyProtocolObj := vmap["v3_privacy_protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3PrivacyProtocolObj) {
		log.Printf("Setting v3_privacy_protocol  V3PrivacyProtocol <<%T>> => %#v", mV3PrivacyProtocolObj, mV3PrivacyProtocolObj)
		mV3PrivacyProtocol := mV3PrivacyProtocolObj.(string)
		object.V3PrivacyProtocol = mV3PrivacyProtocol
	}

	mV3PrivacyPasswordObj := vmap["v3_privacy_password"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3PrivacyPasswordObj) {
		log.Printf("Setting v3_privacy_password  V3PrivacyPassword <<%T>> => %#v", mV3PrivacyPasswordObj, mV3PrivacyPasswordObj)
		mV3PrivacyPassword := mV3PrivacyPasswordObj.(string)
		object.V3PrivacyPassword = mV3PrivacyPassword
	}

	mV3EngineIdObj := vmap["v3_engine_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3EngineIdObj) {
		log.Printf("Setting v3_engine_id  V3EngineId <<%T>> => %#v", mV3EngineIdObj, mV3EngineIdObj)
		mV3EngineId := mV3EngineIdObj.(string)
		object.V3EngineId = mV3EngineId
	}

	mV3EngineBootsObj := vmap["v3_engine_boots"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3EngineBootsObj) {
		log.Printf("Setting v3_engine_boots  V3EngineBoots <<%T>> => %#v", mV3EngineBootsObj, mV3EngineBootsObj)
		mV3EngineBoots := mV3EngineBootsObj.(int)
		object.V3EngineBoots = mV3EngineBoots
	}

	mV3EngineTimeObj := vmap["v3_engine_time"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mV3EngineTimeObj) {
		log.Printf("Setting v3_engine_time  V3EngineTime <<%T>> => %#v", mV3EngineTimeObj, mV3EngineTimeObj)
		mV3EngineTime := mV3EngineTimeObj.(int)
		object.V3EngineTime = mV3EngineTime
	}

	log.Printf("FINISHED SNMPCredentials object: %#v", object)
}

func TakeSNMPCredentialsAsMap(object *SNMPCredentials) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["version"] = object.Version
	omap["local_port"] = object.LocalPort
	omap["retries"] = object.Retries
	omap["timeout"] = object.Timeout
	omap["v2_community"] = object.V2Community
	omap["v3_security_name"] = object.V3SecurityName
	omap["v3_security_level"] = object.V3SecurityLevel
	omap["v3_security_engine_id"] = object.V3SecurityEngineId
	omap["v3_context"] = object.V3Context
	omap["v3_context_engine_id"] = object.V3ContextEngineId
	omap["v3_authentication_protocol"] = object.V3AuthenticationProtocol
	omap["v3_authentication_password"] = object.V3AuthenticationPassword
	omap["v3_privacy_protocol"] = object.V3PrivacyProtocol
	omap["v3_privacy_password"] = object.V3PrivacyPassword
	omap["v3_engine_id"] = object.V3EngineId
	omap["v3_engine_boots"] = object.V3EngineBoots
	omap["v3_engine_time"] = object.V3EngineTime

	return omap
}

func ResourceSNMPCredentialsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"local_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"retries": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"timeout": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"v2_community": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_security_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_security_level": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_security_engine_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_context": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_context_engine_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_authentication_protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_authentication_password": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_privacy_protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_privacy_password": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_engine_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"v3_engine_boots": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"v3_engine_time": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceSNMPCredentials() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSNMPCredentialsSchema(),
	}
}

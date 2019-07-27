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

func SetVirtualIpTypeFromMap(object *VirtualIpType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualIpTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAddressObj := vmap["address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAddressObj) {
		log.Printf("Setting address  Address <<%T>> => %#v", mAddressObj, mAddressObj)
		mAddress := mAddressObj.(string)
		object.Address = mAddress
	}

	mStatusObj := vmap["status"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mStatusObj) {
		log.Printf("Setting status  Status <<%T>> => %#v", mStatusObj, mStatusObj)
		mStatus := mStatusObj.(string)
		object.Status = mStatus
	}

	mStatusDescriptionObj := vmap["status_description"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mStatusDescriptionObj) {
		log.Printf("Setting status_description  StatusDescription <<%T>> => %#v", mStatusDescriptionObj, mStatusDescriptionObj)
		mStatusDescription := mStatusDescriptionObj.(string)
		object.StatusDescription = mStatusDescription
	}

	mAdminStateObj := vmap["admin_state"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAdminStateObj) {
		log.Printf("Setting admin_state  AdminState <<%T>> => %#v", mAdminStateObj, mAdminStateObj)
		mAdminState := mAdminStateObj.(bool)
		object.AdminState = mAdminState
	}

	mProtocolObj := vmap["protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolObj) {
		log.Printf("Setting protocol  Protocol <<%T>> => %#v", mProtocolObj, mProtocolObj)
		mProtocol := mProtocolObj.(string)
		object.Protocol = mProtocol
	}

	mProtocolPortObj := vmap["protocol_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolPortObj) {
		log.Printf("Setting protocol_port  ProtocolPort <<%T>> => %#v", mProtocolPortObj, mProtocolPortObj)
		mProtocolPort := mProtocolPortObj.(int)
		object.ProtocolPort = mProtocolPort
	}

	mConnectionLimitObj := vmap["connection_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mConnectionLimitObj) {
		log.Printf("Setting connection_limit  ConnectionLimit <<%T>> => %#v", mConnectionLimitObj, mConnectionLimitObj)
		mConnectionLimit := mConnectionLimitObj.(int)
		object.ConnectionLimit = mConnectionLimit
	}

	mSubnetIdObj := vmap["subnet_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSubnetIdObj) {
		log.Printf("Setting subnet_id  SubnetId <<%T>> => %#v", mSubnetIdObj, mSubnetIdObj)
		mSubnetId := mSubnetIdObj.(string)
		object.SubnetId = mSubnetId
	}

	mPersistenceCookieNameObj := vmap["persistence_cookie_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPersistenceCookieNameObj) {
		log.Printf("Setting persistence_cookie_name  PersistenceCookieName <<%T>> => %#v", mPersistenceCookieNameObj, mPersistenceCookieNameObj)
		mPersistenceCookieName := mPersistenceCookieNameObj.(string)
		object.PersistenceCookieName = mPersistenceCookieName
	}

	mPersistenceTypeObj := vmap["persistence_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPersistenceTypeObj) {
		log.Printf("Setting persistence_type  PersistenceType <<%T>> => %#v", mPersistenceTypeObj, mPersistenceTypeObj)
		mPersistenceType := mPersistenceTypeObj.(string)
		object.PersistenceType = mPersistenceType
	}

	log.Printf("FINISHED VirtualIpType object: %#v", object)
}

func TakeVirtualIpTypeAsMap(object *VirtualIpType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["address"] = object.Address
	omap["status"] = object.Status
	omap["status_description"] = object.StatusDescription
	omap["admin_state"] = object.AdminState
	omap["protocol"] = object.Protocol
	omap["protocol_port"] = object.ProtocolPort
	omap["connection_limit"] = object.ConnectionLimit
	omap["subnet_id"] = object.SubnetId
	omap["persistence_cookie_name"] = object.PersistenceCookieName
	omap["persistence_type"] = object.PersistenceType

	return omap
}

func ResourceVirtualIpTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"status": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"status_description": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"admin_state": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"protocol_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"connection_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"subnet_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"persistence_cookie_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"persistence_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceVirtualIpType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualIpTypeSchema(),
	}
}

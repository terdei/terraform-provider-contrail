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

func SetLoadbalancerPoolTypeFromMap(object *LoadbalancerPoolType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLoadbalancerPoolTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

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

	mLoadbalancerMethodObj := vmap["loadbalancer_method"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLoadbalancerMethodObj) {
		log.Printf("Setting loadbalancer_method  LoadbalancerMethod <<%T>> => %#v", mLoadbalancerMethodObj, mLoadbalancerMethodObj)
		mLoadbalancerMethod := mLoadbalancerMethodObj.(string)
		object.LoadbalancerMethod = mLoadbalancerMethod
	}

	mSubnetIdObj := vmap["subnet_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSubnetIdObj) {
		log.Printf("Setting subnet_id  SubnetId <<%T>> => %#v", mSubnetIdObj, mSubnetIdObj)
		mSubnetId := mSubnetIdObj.(string)
		object.SubnetId = mSubnetId
	}

	mSessionPersistenceObj := vmap["session_persistence"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSessionPersistenceObj) {
		log.Printf("Setting session_persistence  SessionPersistence <<%T>> => %#v", mSessionPersistenceObj, mSessionPersistenceObj)
		mSessionPersistence := mSessionPersistenceObj.(string)
		object.SessionPersistence = mSessionPersistence
	}

	mPersistenceCookieNameObj := vmap["persistence_cookie_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPersistenceCookieNameObj) {
		log.Printf("Setting persistence_cookie_name  PersistenceCookieName <<%T>> => %#v", mPersistenceCookieNameObj, mPersistenceCookieNameObj)
		mPersistenceCookieName := mPersistenceCookieNameObj.(string)
		object.PersistenceCookieName = mPersistenceCookieName
	}

	log.Printf("FINISHED LoadbalancerPoolType object: %#v", object)
}

func TakeLoadbalancerPoolTypeAsMap(object *LoadbalancerPoolType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["status"] = object.Status
	omap["status_description"] = object.StatusDescription
	omap["admin_state"] = object.AdminState
	omap["protocol"] = object.Protocol
	omap["loadbalancer_method"] = object.LoadbalancerMethod
	omap["subnet_id"] = object.SubnetId
	omap["session_persistence"] = object.SessionPersistence
	omap["persistence_cookie_name"] = object.PersistenceCookieName

	return omap
}

func ResourceLoadbalancerPoolTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"loadbalancer_method": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"subnet_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"session_persistence": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"persistence_cookie_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceLoadbalancerPoolType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLoadbalancerPoolTypeSchema(),
	}
}

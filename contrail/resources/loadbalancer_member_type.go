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

func SetLoadbalancerMemberTypeFromMap(object *LoadbalancerMemberType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLoadbalancerMemberTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAdminStateObj := vmap["admin_state"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAdminStateObj) {
		log.Printf("Setting admin_state  AdminState <<%T>> => %#v", mAdminStateObj, mAdminStateObj)
		mAdminState := mAdminStateObj.(bool)
		object.AdminState = mAdminState
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

	mProtocolPortObj := vmap["protocol_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolPortObj) {
		log.Printf("Setting protocol_port  ProtocolPort <<%T>> => %#v", mProtocolPortObj, mProtocolPortObj)
		mProtocolPort := mProtocolPortObj.(int)
		object.ProtocolPort = mProtocolPort
	}

	mWeightObj := vmap["weight"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mWeightObj) {
		log.Printf("Setting weight  Weight <<%T>> => %#v", mWeightObj, mWeightObj)
		mWeight := mWeightObj.(int)
		object.Weight = mWeight
	}

	mAddressObj := vmap["address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAddressObj) {
		log.Printf("Setting address  Address <<%T>> => %#v", mAddressObj, mAddressObj)
		mAddress := mAddressObj.(string)
		object.Address = mAddress
	}

	log.Printf("FINISHED LoadbalancerMemberType object: %#v", object)
}

func TakeLoadbalancerMemberTypeAsMap(object *LoadbalancerMemberType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["admin_state"] = object.AdminState
	omap["status"] = object.Status
	omap["status_description"] = object.StatusDescription
	omap["protocol_port"] = object.ProtocolPort
	omap["weight"] = object.Weight
	omap["address"] = object.Address

	return omap
}

func ResourceLoadbalancerMemberTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
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
		"protocol_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"weight": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceLoadbalancerMemberType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLoadbalancerMemberTypeSchema(),
	}
}

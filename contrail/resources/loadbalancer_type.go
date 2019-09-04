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

func SetLoadbalancerTypeFromMap(object *LoadbalancerType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLoadbalancerTypeFromMAP] map = %#v", val)

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

	mProvisioningStatusObj := vmap["provisioning_status"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProvisioningStatusObj) {
		log.Printf("Setting provisioning_status  ProvisioningStatus <<%T>> => %#v", mProvisioningStatusObj, mProvisioningStatusObj)
		mProvisioningStatus := mProvisioningStatusObj.(string)
		object.ProvisioningStatus = mProvisioningStatus
	}

	mOperatingStatusObj := vmap["operating_status"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOperatingStatusObj) {
		log.Printf("Setting operating_status  OperatingStatus <<%T>> => %#v", mOperatingStatusObj, mOperatingStatusObj)
		mOperatingStatus := mOperatingStatusObj.(string)
		object.OperatingStatus = mOperatingStatus
	}

	mVipSubnetIdObj := vmap["vip_subnet_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVipSubnetIdObj) {
		log.Printf("Setting vip_subnet_id  VipSubnetId <<%T>> => %#v", mVipSubnetIdObj, mVipSubnetIdObj)
		mVipSubnetId := mVipSubnetIdObj.(string)
		object.VipSubnetId = mVipSubnetId
	}

	mVipAddressObj := vmap["vip_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVipAddressObj) {
		log.Printf("Setting vip_address  VipAddress <<%T>> => %#v", mVipAddressObj, mVipAddressObj)
		mVipAddress := mVipAddressObj.(string)
		object.VipAddress = mVipAddress
	}

	mAdminStateObj := vmap["admin_state"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAdminStateObj) {
		log.Printf("Setting admin_state  AdminState <<%T>> => %#v", mAdminStateObj, mAdminStateObj)
		mAdminState := mAdminStateObj.(bool)
		object.AdminState = mAdminState
	}

	log.Printf("FINISHED LoadbalancerType object: %#v", object)
}

func TakeLoadbalancerTypeAsMap(object *LoadbalancerType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["status"] = object.Status
	omap["provisioning_status"] = object.ProvisioningStatus
	omap["operating_status"] = object.OperatingStatus
	omap["vip_subnet_id"] = object.VipSubnetId
	omap["vip_address"] = object.VipAddress
	omap["admin_state"] = object.AdminState

	return omap
}

func ResourceLoadbalancerTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"status": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"provisioning_status": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"operating_status": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"vip_subnet_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"vip_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"admin_state": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceLoadbalancerType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLoadbalancerTypeSchema(),
	}
}

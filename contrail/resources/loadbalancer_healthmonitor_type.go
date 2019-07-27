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

func SetLoadbalancerHealthmonitorTypeFromMap(object *LoadbalancerHealthmonitorType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLoadbalancerHealthmonitorTypeFromMAP] map = %#v", val)

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

	mMonitorTypeObj := vmap["monitor_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMonitorTypeObj) {
		log.Printf("Setting monitor_type  MonitorType <<%T>> => %#v", mMonitorTypeObj, mMonitorTypeObj)
		mMonitorType := mMonitorTypeObj.(string)
		object.MonitorType = mMonitorType
	}

	mDelayObj := vmap["delay"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDelayObj) {
		log.Printf("Setting delay  Delay <<%T>> => %#v", mDelayObj, mDelayObj)
		mDelay := mDelayObj.(int)
		object.Delay = mDelay
	}

	mTimeoutObj := vmap["timeout"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTimeoutObj) {
		log.Printf("Setting timeout  Timeout <<%T>> => %#v", mTimeoutObj, mTimeoutObj)
		mTimeout := mTimeoutObj.(int)
		object.Timeout = mTimeout
	}

	mMaxRetriesObj := vmap["max_retries"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMaxRetriesObj) {
		log.Printf("Setting max_retries  MaxRetries <<%T>> => %#v", mMaxRetriesObj, mMaxRetriesObj)
		mMaxRetries := mMaxRetriesObj.(int)
		object.MaxRetries = mMaxRetries
	}

	mHttpMethodObj := vmap["http_method"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mHttpMethodObj) {
		log.Printf("Setting http_method  HttpMethod <<%T>> => %#v", mHttpMethodObj, mHttpMethodObj)
		mHttpMethod := mHttpMethodObj.(string)
		object.HttpMethod = mHttpMethod
	}

	mUrlPathObj := vmap["url_path"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUrlPathObj) {
		log.Printf("Setting url_path  UrlPath <<%T>> => %#v", mUrlPathObj, mUrlPathObj)
		mUrlPath := mUrlPathObj.(string)
		object.UrlPath = mUrlPath
	}

	mExpectedCodesObj := vmap["expected_codes"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mExpectedCodesObj) {
		log.Printf("Setting expected_codes  ExpectedCodes <<%T>> => %#v", mExpectedCodesObj, mExpectedCodesObj)
		mExpectedCodes := mExpectedCodesObj.(string)
		object.ExpectedCodes = mExpectedCodes
	}

	log.Printf("FINISHED LoadbalancerHealthmonitorType object: %#v", object)
}

func TakeLoadbalancerHealthmonitorTypeAsMap(object *LoadbalancerHealthmonitorType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["admin_state"] = object.AdminState
	omap["monitor_type"] = object.MonitorType
	omap["delay"] = object.Delay
	omap["timeout"] = object.Timeout
	omap["max_retries"] = object.MaxRetries
	omap["http_method"] = object.HttpMethod
	omap["url_path"] = object.UrlPath
	omap["expected_codes"] = object.ExpectedCodes

	return omap
}

func ResourceLoadbalancerHealthmonitorTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"monitor_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"delay": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"timeout": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"max_retries": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"http_method": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"url_path": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"expected_codes": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceLoadbalancerHealthmonitorType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLoadbalancerHealthmonitorTypeSchema(),
	}
}

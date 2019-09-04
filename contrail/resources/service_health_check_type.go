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

func SetServiceHealthCheckTypeFromMap(object *ServiceHealthCheckType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceHealthCheckTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mEnabledObj := vmap["enabled"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEnabledObj) {
		log.Printf("Setting enabled  Enabled <<%T>> => %#v", mEnabledObj, mEnabledObj)
		mEnabled := mEnabledObj.(bool)
		object.Enabled = mEnabled
	}

	mHealthCheckTypeObj := vmap["health_check_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mHealthCheckTypeObj) {
		log.Printf("Setting health_check_type  HealthCheckType <<%T>> => %#v", mHealthCheckTypeObj, mHealthCheckTypeObj)
		mHealthCheckType := mHealthCheckTypeObj.(string)
		object.HealthCheckType = mHealthCheckType
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

	mDelayusecsObj := vmap["delayUsecs"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDelayusecsObj) {
		log.Printf("Setting delayUsecs  Delayusecs <<%T>> => %#v", mDelayusecsObj, mDelayusecsObj)
		mDelayusecs := mDelayusecsObj.(int)
		object.Delayusecs = mDelayusecs
	}

	mTimeoutObj := vmap["timeout"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTimeoutObj) {
		log.Printf("Setting timeout  Timeout <<%T>> => %#v", mTimeoutObj, mTimeoutObj)
		mTimeout := mTimeoutObj.(int)
		object.Timeout = mTimeout
	}

	mTimeoutusecsObj := vmap["timeoutUsecs"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTimeoutusecsObj) {
		log.Printf("Setting timeoutUsecs  Timeoutusecs <<%T>> => %#v", mTimeoutusecsObj, mTimeoutusecsObj)
		mTimeoutusecs := mTimeoutusecsObj.(int)
		object.Timeoutusecs = mTimeoutusecs
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

	log.Printf("FINISHED ServiceHealthCheckType object: %#v", object)
}

func TakeServiceHealthCheckTypeAsMap(object *ServiceHealthCheckType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["enabled"] = object.Enabled
	omap["health_check_type"] = object.HealthCheckType
	omap["monitor_type"] = object.MonitorType
	omap["delay"] = object.Delay
	omap["delayUsecs"] = object.Delayusecs
	omap["timeout"] = object.Timeout
	omap["timeoutUsecs"] = object.Timeoutusecs
	omap["max_retries"] = object.MaxRetries
	omap["http_method"] = object.HttpMethod
	omap["url_path"] = object.UrlPath
	omap["expected_codes"] = object.ExpectedCodes

	return omap
}

func ResourceServiceHealthCheckTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"health_check_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
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
		"delayusecs": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"timeout": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"timeoutusecs": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
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

func ResourceServiceHealthCheckType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceHealthCheckTypeSchema(),
	}
}

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

func SetGracefulRestartParametersTypeFromMap(object *GracefulRestartParametersType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetGracefulRestartParametersTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mEnableObj := vmap["enable"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEnableObj) {
		log.Printf("Setting enable  Enable <<%T>> => %#v", mEnableObj, mEnableObj)
		mEnable := mEnableObj.(bool)
		object.Enable = mEnable
	}

	mRestartTimeObj := vmap["restart_time"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRestartTimeObj) {
		log.Printf("Setting restart_time  RestartTime <<%T>> => %#v", mRestartTimeObj, mRestartTimeObj)
		mRestartTime := mRestartTimeObj.(int)
		object.RestartTime = mRestartTime
	}

	mLongLivedRestartTimeObj := vmap["long_lived_restart_time"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLongLivedRestartTimeObj) {
		log.Printf("Setting long_lived_restart_time  LongLivedRestartTime <<%T>> => %#v", mLongLivedRestartTimeObj, mLongLivedRestartTimeObj)
		mLongLivedRestartTime := mLongLivedRestartTimeObj.(int)
		object.LongLivedRestartTime = mLongLivedRestartTime
	}

	mEndOfRibTimeoutObj := vmap["end_of_rib_timeout"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEndOfRibTimeoutObj) {
		log.Printf("Setting end_of_rib_timeout  EndOfRibTimeout <<%T>> => %#v", mEndOfRibTimeoutObj, mEndOfRibTimeoutObj)
		mEndOfRibTimeout := mEndOfRibTimeoutObj.(int)
		object.EndOfRibTimeout = mEndOfRibTimeout
	}

	mBgpHelperEnableObj := vmap["bgp_helper_enable"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mBgpHelperEnableObj) {
		log.Printf("Setting bgp_helper_enable  BgpHelperEnable <<%T>> => %#v", mBgpHelperEnableObj, mBgpHelperEnableObj)
		mBgpHelperEnable := mBgpHelperEnableObj.(bool)
		object.BgpHelperEnable = mBgpHelperEnable
	}

	mXmppHelperEnableObj := vmap["xmpp_helper_enable"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mXmppHelperEnableObj) {
		log.Printf("Setting xmpp_helper_enable  XmppHelperEnable <<%T>> => %#v", mXmppHelperEnableObj, mXmppHelperEnableObj)
		mXmppHelperEnable := mXmppHelperEnableObj.(bool)
		object.XmppHelperEnable = mXmppHelperEnable
	}

	log.Printf("FINISHED GracefulRestartParametersType object: %#v", object)
}

func TakeGracefulRestartParametersTypeAsMap(object *GracefulRestartParametersType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["enable"] = object.Enable
	omap["restart_time"] = object.RestartTime
	omap["long_lived_restart_time"] = object.LongLivedRestartTime
	omap["end_of_rib_timeout"] = object.EndOfRibTimeout
	omap["bgp_helper_enable"] = object.BgpHelperEnable
	omap["xmpp_helper_enable"] = object.XmppHelperEnable

	return omap
}

func ResourceGracefulRestartParametersTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"restart_time": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"long_lived_restart_time": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"end_of_rib_timeout": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"bgp_helper_enable": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"xmpp_helper_enable": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceGracefulRestartParametersType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceGracefulRestartParametersTypeSchema(),
	}
}

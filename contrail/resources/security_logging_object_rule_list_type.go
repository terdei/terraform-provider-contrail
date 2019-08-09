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

func SetSecurityLoggingObjectRuleEntryTypeFromMap(object *SecurityLoggingObjectRuleEntryType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSecurityLoggingObjectRuleEntryTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRuleUuidObj := vmap["rule_uuid"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRuleUuidObj) {
		log.Printf("Setting rule_uuid  RuleUuid <<%T>> => %#v", mRuleUuidObj, mRuleUuidObj)
		mRuleUuid := mRuleUuidObj.(string)
		object.RuleUuid = mRuleUuid
	}

	mRateObj := vmap["rate"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRateObj) {
		log.Printf("Setting rate  Rate <<%T>> => %#v", mRateObj, mRateObj)
		mRate := mRateObj.(int)
		object.Rate = mRate
	}

	log.Printf("FINISHED SecurityLoggingObjectRuleEntryType object: %#v", object)
}

func TakeSecurityLoggingObjectRuleEntryTypeAsMap(object *SecurityLoggingObjectRuleEntryType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["rule_uuid"] = object.RuleUuid
	omap["rate"] = object.Rate

	return omap
}

func ResourceSecurityLoggingObjectRuleEntryTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule_uuid": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"rate": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceSecurityLoggingObjectRuleEntryType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSecurityLoggingObjectRuleEntryTypeSchema(),
	}
}

func SetSecurityLoggingObjectRuleListTypeFromMap(object *SecurityLoggingObjectRuleListType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSecurityLoggingObjectRuleListTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRuleObj := vmap["rule"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mRuleObj) {
		log.Printf("Setting rule  Rule <<%T>> => %#v", mRuleObj, mRuleObj)
		for _, v := range mRuleObj.([]interface{}) {
			mRule := new(SecurityLoggingObjectRuleEntryType)
			SetSecurityLoggingObjectRuleEntryTypeFromMap(mRule, d, m, v)
			object.AddRule(mRule)
		}
	}

	log.Printf("FINISHED SecurityLoggingObjectRuleListType object: %#v", object)
}

func TakeSecurityLoggingObjectRuleListTypeAsMap(object *SecurityLoggingObjectRuleListType) map[string]interface{} {
	omap := make(map[string]interface{})

	var rule_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Rule {
		rule_map = append(rule_map, TakeSecurityLoggingObjectRuleEntryTypeAsMap(&v))
	}
	omap["rule"] = rule_map

	return omap
}

func ResourceSecurityLoggingObjectRuleListTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: SecurityLoggingObjectRuleEntryType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSecurityLoggingObjectRuleEntryType(),
		},
	}
}

func ResourceSecurityLoggingObjectRuleListType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSecurityLoggingObjectRuleListTypeSchema(),
	}
}

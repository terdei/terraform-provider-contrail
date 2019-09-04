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

func SetAclRuleTypeFromMap(object *AclRuleType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAclRuleTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mMatchConditionObj := vmap["match_condition"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mMatchConditionObj) {
		log.Printf("Setting match_condition  MatchCondition <<%T>> => %#v", mMatchConditionObj, mMatchConditionObj)
		mMatchCondition := new(MatchConditionType)
		SetMatchConditionTypeFromMap(mMatchCondition, d, m, mMatchConditionObj)
		object.MatchCondition = mMatchCondition
	}

	mActionListObj := vmap["action_list"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mActionListObj) {
		log.Printf("Setting action_list  ActionList <<%T>> => %#v", mActionListObj, mActionListObj)
		mActionList := new(ActionListType)
		SetActionListTypeFromMap(mActionList, d, m, mActionListObj)
		object.ActionList = mActionList
	}

	mRuleUuidObj := vmap["rule_uuid"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRuleUuidObj) {
		log.Printf("Setting rule_uuid  RuleUuid <<%T>> => %#v", mRuleUuidObj, mRuleUuidObj)
		mRuleUuid := mRuleUuidObj.(string)
		object.RuleUuid = mRuleUuid
	}

	mDirectionObj := vmap["direction"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDirectionObj) {
		log.Printf("Setting direction  Direction <<%T>> => %#v", mDirectionObj, mDirectionObj)
		mDirection := mDirectionObj.(string)
		object.Direction = mDirection
	}

	log.Printf("FINISHED AclRuleType object: %#v", object)
}

func TakeAclRuleTypeAsMap(object *AclRuleType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.MatchCondition != nil {
		omap["match_condition"] = TakeMatchConditionTypeAsMap(object.MatchCondition)
	}
	if object.ActionList != nil {
		omap["action_list"] = TakeActionListTypeAsMap(object.ActionList)
	}
	omap["rule_uuid"] = object.RuleUuid
	omap["direction"] = object.Direction

	return omap
}

func ResourceAclRuleTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"match_condition": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: MatchConditionType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceMatchConditionType(),
		},
		"action_list": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: ActionListType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceActionListType(),
		},
		"rule_uuid": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"direction": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceAclRuleType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAclRuleTypeSchema(),
	}
}

func SetAclEntriesTypeFromMap(object *AclEntriesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAclEntriesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mDynamicObj := vmap["dynamic"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDynamicObj) {
		log.Printf("Setting dynamic  Dynamic <<%T>> => %#v", mDynamicObj, mDynamicObj)
		mDynamic := mDynamicObj.(bool)
		object.Dynamic = mDynamic
	}

	mAclRuleObj := vmap["acl_rule"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mAclRuleObj) {
		log.Printf("Setting acl_rule  AclRule <<%T>> => %#v", mAclRuleObj, mAclRuleObj)
		for _, v := range mAclRuleObj.([]interface{}) {
			mAclRule := new(AclRuleType)
			SetAclRuleTypeFromMap(mAclRule, d, m, v)
			object.AddAclRule(mAclRule)
		}
	}

	log.Printf("FINISHED AclEntriesType object: %#v", object)
}

func TakeAclEntriesTypeAsMap(object *AclEntriesType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["dynamic"] = object.Dynamic
	var acl_rule_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.AclRule {
		acl_rule_map = append(acl_rule_map, TakeAclRuleTypeAsMap(&v))
	}
	omap["acl_rule"] = acl_rule_map

	return omap
}

func ResourceAclEntriesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dynamic": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"acl_rule": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AclRuleType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceAclRuleType(),
		},
	}
}

func ResourceAclEntriesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAclEntriesTypeSchema(),
	}
}

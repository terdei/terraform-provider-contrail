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

func SetRbacPermTypeFromMap(object *RbacPermType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRbacPermTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRoleNameObj := vmap["role_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRoleNameObj) {
		log.Printf("Setting role_name  RoleName <<%T>> => %#v", mRoleNameObj, mRoleNameObj)
		mRoleName := mRoleNameObj.(string)
		object.RoleName = mRoleName
	}

	mRoleCrudObj := vmap["role_crud"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRoleCrudObj) {
		log.Printf("Setting role_crud  RoleCrud <<%T>> => %#v", mRoleCrudObj, mRoleCrudObj)
		mRoleCrud := mRoleCrudObj.(string)
		object.RoleCrud = mRoleCrud
	}

	log.Printf("FINISHED RbacPermType object: %#v", object)
}

func TakeRbacPermTypeAsMap(object *RbacPermType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["role_name"] = object.RoleName
	omap["role_crud"] = object.RoleCrud

	return omap
}

func ResourceRbacPermTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"role_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"role_crud": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceRbacPermType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRbacPermTypeSchema(),
	}
}

func SetRbacRuleTypeFromMap(object *RbacRuleType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRbacRuleTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRuleObjectObj := vmap["rule_object"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRuleObjectObj) {
		log.Printf("Setting rule_object  RuleObject <<%T>> => %#v", mRuleObjectObj, mRuleObjectObj)
		mRuleObject := mRuleObjectObj.(string)
		object.RuleObject = mRuleObject
	}

	mRuleFieldObj := vmap["rule_field"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRuleFieldObj) {
		log.Printf("Setting rule_field  RuleField <<%T>> => %#v", mRuleFieldObj, mRuleFieldObj)
		mRuleField := mRuleFieldObj.(string)
		object.RuleField = mRuleField
	}

	mRulePermsObj := vmap["rule_perms"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mRulePermsObj) {
		log.Printf("Setting rule_perms  RulePerms <<%T>> => %#v", mRulePermsObj, mRulePermsObj)
		for _, v := range mRulePermsObj.([]interface{}) {
			mRulePerms := new(RbacPermType)
			SetRbacPermTypeFromMap(mRulePerms, d, m, v)
			object.AddRulePerms(mRulePerms)
		}
	}

	log.Printf("FINISHED RbacRuleType object: %#v", object)
}

func TakeRbacRuleTypeAsMap(object *RbacRuleType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["rule_object"] = object.RuleObject
	omap["rule_field"] = object.RuleField
	var rule_perms_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.RulePerms {
		rule_perms_map = append(rule_perms_map, TakeRbacPermTypeAsMap(&v))
	}
	omap["rule_perms"] = rule_perms_map

	return omap
}

func ResourceRbacRuleTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule_object": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"rule_field": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"rule_perms": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: RbacPermType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceRbacPermType(),
		},
	}
}

func ResourceRbacRuleType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRbacRuleTypeSchema(),
	}
}

func SetRbacRuleEntriesTypeFromMap(object *RbacRuleEntriesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRbacRuleEntriesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRbacRuleObj := vmap["rbac_rule"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mRbacRuleObj) {
		log.Printf("Setting rbac_rule  RbacRule <<%T>> => %#v", mRbacRuleObj, mRbacRuleObj)
		for _, v := range mRbacRuleObj.([]interface{}) {
			mRbacRule := new(RbacRuleType)
			SetRbacRuleTypeFromMap(mRbacRule, d, m, v)
			object.AddRbacRule(mRbacRule)
		}
	}

	log.Printf("FINISHED RbacRuleEntriesType object: %#v", object)
}

func TakeRbacRuleEntriesTypeAsMap(object *RbacRuleEntriesType) map[string]interface{} {
	omap := make(map[string]interface{})

	var rbac_rule_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.RbacRule {
		rbac_rule_map = append(rbac_rule_map, TakeRbacRuleTypeAsMap(&v))
	}
	omap["rbac_rule"] = rbac_rule_map

	return omap
}

func ResourceRbacRuleEntriesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rbac_rule": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: RbacRuleType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRbacRuleType(),
		},
	}
}

func ResourceRbacRuleEntriesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRbacRuleEntriesTypeSchema(),
	}
}

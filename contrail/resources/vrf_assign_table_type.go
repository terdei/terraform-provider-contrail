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

func SetVrfAssignRuleTypeFromMap(object *VrfAssignRuleType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVrfAssignRuleTypeFromMAP] map = %#v", val)

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

	mVlanTagObj := vmap["vlan_tag"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVlanTagObj) {
		log.Printf("Setting vlan_tag  VlanTag <<%T>> => %#v", mVlanTagObj, mVlanTagObj)
		mVlanTag := mVlanTagObj.(int)
		object.VlanTag = mVlanTag
	}

	mRoutingInstanceObj := vmap["routing_instance"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRoutingInstanceObj) {
		log.Printf("Setting routing_instance  RoutingInstance <<%T>> => %#v", mRoutingInstanceObj, mRoutingInstanceObj)
		mRoutingInstance := mRoutingInstanceObj.(string)
		object.RoutingInstance = mRoutingInstance
	}

	mIgnoreAclObj := vmap["ignore_acl"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIgnoreAclObj) {
		log.Printf("Setting ignore_acl  IgnoreAcl <<%T>> => %#v", mIgnoreAclObj, mIgnoreAclObj)
		mIgnoreAcl := mIgnoreAclObj.(bool)
		object.IgnoreAcl = mIgnoreAcl
	}

	log.Printf("FINISHED VrfAssignRuleType object: %#v", object)
}

func TakeVrfAssignRuleTypeAsMap(object *VrfAssignRuleType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.MatchCondition != nil {
		omap["match_condition"] = TakeMatchConditionTypeAsMap(object.MatchCondition)
	}
	omap["vlan_tag"] = object.VlanTag
	omap["routing_instance"] = object.RoutingInstance
	omap["ignore_acl"] = object.IgnoreAcl

	return omap
}

func ResourceVrfAssignRuleTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"match_condition": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: MatchConditionType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMatchConditionType(),
		},
		"vlan_tag": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"routing_instance": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ignore_acl": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceVrfAssignRuleType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVrfAssignRuleTypeSchema(),
	}
}

func SetVrfAssignTableTypeFromMap(object *VrfAssignTableType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVrfAssignTableTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mVrfAssignRuleObj := vmap["vrf_assign_rule"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mVrfAssignRuleObj) {
		log.Printf("Setting vrf_assign_rule  VrfAssignRule <<%T>> => %#v", mVrfAssignRuleObj, mVrfAssignRuleObj)
		for _, v := range mVrfAssignRuleObj.([]interface{}) {
			mVrfAssignRule := new(VrfAssignRuleType)
			SetVrfAssignRuleTypeFromMap(mVrfAssignRule, d, m, v)
			object.AddVrfAssignRule(mVrfAssignRule)
		}
	}

	log.Printf("FINISHED VrfAssignTableType object: %#v", object)
}

func TakeVrfAssignTableTypeAsMap(object *VrfAssignTableType) map[string]interface{} {
	omap := make(map[string]interface{})

	var vrf_assign_rule_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.VrfAssignRule {
		vrf_assign_rule_map = append(vrf_assign_rule_map, TakeVrfAssignRuleTypeAsMap(&v))
	}
	omap["vrf_assign_rule"] = vrf_assign_rule_map

	return omap
}

func ResourceVrfAssignTableTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vrf_assign_rule": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: VrfAssignRuleType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVrfAssignRuleType(),
		},
	}
}

func ResourceVrfAssignTableType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVrfAssignTableTypeSchema(),
	}
}

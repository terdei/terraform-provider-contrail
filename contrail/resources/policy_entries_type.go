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

func SetPolicyRuleTypeFromMap(object *PolicyRuleType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPolicyRuleTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRuleSequenceObj := vmap["rule_sequence"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mRuleSequenceObj) {
		log.Printf("Setting rule_sequence  RuleSequence <<%T>> => %#v", mRuleSequenceObj, mRuleSequenceObj)
		mRuleSequence := new(SequenceType)
		SetSequenceTypeFromMap(mRuleSequence, d, m, mRuleSequenceObj)
		object.RuleSequence = mRuleSequence
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

	mProtocolObj := vmap["protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolObj) {
		log.Printf("Setting protocol  Protocol <<%T>> => %#v", mProtocolObj, mProtocolObj)
		mProtocol := mProtocolObj.(string)
		object.Protocol = mProtocol
	}

	mSrcAddressesObj := vmap["src_addresses"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSrcAddressesObj) {
		log.Printf("Setting src_addresses  SrcAddresses <<%T>> => %#v", mSrcAddressesObj, mSrcAddressesObj)
		for _, v := range mSrcAddressesObj.([]interface{}) {
			mSrcAddresses := new(AddressType)
			SetAddressTypeFromMap(mSrcAddresses, d, m, v)
			object.AddSrcAddresses(mSrcAddresses)
		}
	}

	mSrcPortsObj := vmap["src_ports"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSrcPortsObj) {
		log.Printf("Setting src_ports  SrcPorts <<%T>> => %#v", mSrcPortsObj, mSrcPortsObj)
		for _, v := range mSrcPortsObj.([]interface{}) {
			mSrcPorts := new(PortType)
			SetPortTypeFromMap(mSrcPorts, d, m, v)
			object.AddSrcPorts(mSrcPorts)
		}
	}

	mApplicationObj := vmap["application"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mApplicationObj) {
		log.Printf("Setting application  Application <<%T>> => %#v", mApplicationObj, mApplicationObj)
		for _, v := range mApplicationObj.([]interface{}) {
			mApplication := v.(string)
			object.AddApplication(mApplication)
		}
	}

	mDstAddressesObj := vmap["dst_addresses"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mDstAddressesObj) {
		log.Printf("Setting dst_addresses  DstAddresses <<%T>> => %#v", mDstAddressesObj, mDstAddressesObj)
		for _, v := range mDstAddressesObj.([]interface{}) {
			mDstAddresses := new(AddressType)
			SetAddressTypeFromMap(mDstAddresses, d, m, v)
			object.AddDstAddresses(mDstAddresses)
		}
	}

	mDstPortsObj := vmap["dst_ports"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mDstPortsObj) {
		log.Printf("Setting dst_ports  DstPorts <<%T>> => %#v", mDstPortsObj, mDstPortsObj)
		for _, v := range mDstPortsObj.([]interface{}) {
			mDstPorts := new(PortType)
			SetPortTypeFromMap(mDstPorts, d, m, v)
			object.AddDstPorts(mDstPorts)
		}
	}

	mActionListObj := vmap["action_list"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mActionListObj) {
		log.Printf("Setting action_list  ActionList <<%T>> => %#v", mActionListObj, mActionListObj)
		mActionList := new(ActionListType)
		SetActionListTypeFromMap(mActionList, d, m, mActionListObj)
		object.ActionList = mActionList
	}

	mEthertypeObj := vmap["ethertype"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEthertypeObj) {
		log.Printf("Setting ethertype  Ethertype <<%T>> => %#v", mEthertypeObj, mEthertypeObj)
		mEthertype := mEthertypeObj.(string)
		object.Ethertype = mEthertype
	}

	mCreatedObj := vmap["created"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mCreatedObj) {
		log.Printf("Setting created  Created <<%T>> => %#v", mCreatedObj, mCreatedObj)
		mCreated := mCreatedObj.(string)
		object.Created = mCreated
	}

	mLastModifiedObj := vmap["last_modified"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLastModifiedObj) {
		log.Printf("Setting last_modified  LastModified <<%T>> => %#v", mLastModifiedObj, mLastModifiedObj)
		mLastModified := mLastModifiedObj.(string)
		object.LastModified = mLastModified
	}

	log.Printf("FINISHED PolicyRuleType object: %#v", object)
}

func TakePolicyRuleTypeAsMap(object *PolicyRuleType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.RuleSequence != nil {
		omap["rule_sequence"] = TakeSequenceTypeAsMap(object.RuleSequence)
	}
	omap["rule_uuid"] = object.RuleUuid
	omap["direction"] = object.Direction
	omap["protocol"] = object.Protocol
	var src_addresses_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.SrcAddresses {
		src_addresses_map = append(src_addresses_map, TakeAddressTypeAsMap(&v))
	}
	omap["src_addresses"] = src_addresses_map
	var src_ports_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.SrcPorts {
		src_ports_map = append(src_ports_map, TakePortTypeAsMap(&v))
	}
	omap["src_ports"] = src_ports_map
	application_arr := make([]string, len(object.Application))
	for _, v := range object.Application {
		application_arr = append(application_arr, v)
	}
	omap["application"] = application_arr
	var dst_addresses_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.DstAddresses {
		dst_addresses_map = append(dst_addresses_map, TakeAddressTypeAsMap(&v))
	}
	omap["dst_addresses"] = dst_addresses_map
	var dst_ports_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.DstPorts {
		dst_ports_map = append(dst_ports_map, TakePortTypeAsMap(&v))
	}
	omap["dst_ports"] = dst_ports_map
	if object.ActionList != nil {
		omap["action_list"] = TakeActionListTypeAsMap(object.ActionList)
	}
	omap["ethertype"] = object.Ethertype
	omap["created"] = object.Created
	omap["last_modified"] = object.LastModified

	return omap
}

func ResourcePolicyRuleTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule_sequence": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SequenceType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSequenceType(),
		},
		"rule_uuid": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"direction": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"src_addresses": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AddressType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceAddressType(),
		},
		"src_ports": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: PortType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortType(),
		},
		"application": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"dst_addresses": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AddressType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceAddressType(),
		},
		"dst_ports": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: PortType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortType(),
		},
		"action_list": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: ActionListType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceActionListType(),
		},
		"ethertype": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"created": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:dateTime
			Optional: true,
			Type:     schema.TypeString,
		},
		"last_modified": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:dateTime
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourcePolicyRuleType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePolicyRuleTypeSchema(),
	}
}

func SetPolicyEntriesTypeFromMap(object *PolicyEntriesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPolicyEntriesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPolicyRuleObj := vmap["policy_rule"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mPolicyRuleObj) {
		log.Printf("Setting policy_rule  PolicyRule <<%T>> => %#v", mPolicyRuleObj, mPolicyRuleObj)
		for _, v := range mPolicyRuleObj.([]interface{}) {
			mPolicyRule := new(PolicyRuleType)
			SetPolicyRuleTypeFromMap(mPolicyRule, d, m, v)
			object.AddPolicyRule(mPolicyRule)
		}
	}

	log.Printf("FINISHED PolicyEntriesType object: %#v", object)
}

func TakePolicyEntriesTypeAsMap(object *PolicyEntriesType) map[string]interface{} {
	omap := make(map[string]interface{})

	var policy_rule_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.PolicyRule {
		policy_rule_map = append(policy_rule_map, TakePolicyRuleTypeAsMap(&v))
	}
	omap["policy_rule"] = policy_rule_map

	return omap
}

func ResourcePolicyEntriesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"policy_rule": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: PolicyRuleType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourcePolicyRuleType(),
		},
	}
}

func ResourcePolicyEntriesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePolicyEntriesTypeSchema(),
	}
}

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

func SetActionListTypeFromMap(object *ActionListType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetActionListTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSimpleActionObj := vmap["simple_action"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSimpleActionObj) {
		log.Printf("Setting simple_action  SimpleAction <<%T>> => %#v", mSimpleActionObj, mSimpleActionObj)
		mSimpleAction := mSimpleActionObj.(string)
		object.SimpleAction = mSimpleAction
	}

	mGatewayNameObj := vmap["gateway_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mGatewayNameObj) {
		log.Printf("Setting gateway_name  GatewayName <<%T>> => %#v", mGatewayNameObj, mGatewayNameObj)
		mGatewayName := mGatewayNameObj.(string)
		object.GatewayName = mGatewayName
	}

	mApplyServiceObj := vmap["apply_service"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mApplyServiceObj) {
		log.Printf("Setting apply_service  ApplyService <<%T>> => %#v", mApplyServiceObj, mApplyServiceObj)
		for _, v := range mApplyServiceObj.([]interface{}) {
			mApplyService := v.(string)
			object.AddApplyService(mApplyService)
		}
	}

	mMirrorToObj := vmap["mirror_to"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mMirrorToObj) {
		log.Printf("Setting mirror_to  MirrorTo <<%T>> => %#v", mMirrorToObj, mMirrorToObj)
		mMirrorTo := new(MirrorActionType)
		SetMirrorActionTypeFromMap(mMirrorTo, d, m, mMirrorToObj)
		object.MirrorTo = mMirrorTo
	}

	mAssignRoutingInstanceObj := vmap["assign_routing_instance"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAssignRoutingInstanceObj) {
		log.Printf("Setting assign_routing_instance  AssignRoutingInstance <<%T>> => %#v", mAssignRoutingInstanceObj, mAssignRoutingInstanceObj)
		mAssignRoutingInstance := mAssignRoutingInstanceObj.(string)
		object.AssignRoutingInstance = mAssignRoutingInstance
	}

	mLogObj := vmap["log"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLogObj) {
		log.Printf("Setting log  Log <<%T>> => %#v", mLogObj, mLogObj)
		mLog := mLogObj.(bool)
		object.Log = mLog
	}

	mAlertObj := vmap["alert"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAlertObj) {
		log.Printf("Setting alert  Alert <<%T>> => %#v", mAlertObj, mAlertObj)
		mAlert := mAlertObj.(bool)
		object.Alert = mAlert
	}

	mQosActionObj := vmap["qos_action"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mQosActionObj) {
		log.Printf("Setting qos_action  QosAction <<%T>> => %#v", mQosActionObj, mQosActionObj)
		mQosAction := mQosActionObj.(string)
		object.QosAction = mQosAction
	}

	log.Printf("FINISHED ActionListType object: %#v", object)
}

func TakeActionListTypeAsMap(object *ActionListType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["simple_action"] = object.SimpleAction
	omap["gateway_name"] = object.GatewayName
	apply_service_arr := make([]string, len(object.ApplyService))
	for _, v := range object.ApplyService {
		apply_service_arr = append(apply_service_arr, v)
	}
	omap["apply_service"] = apply_service_arr
	if object.MirrorTo != nil {
		omap["mirror_to"] = TakeMirrorActionTypeAsMap(object.MirrorTo)
	}
	omap["assign_routing_instance"] = object.AssignRoutingInstance
	omap["log"] = object.Log
	omap["alert"] = object.Alert
	omap["qos_action"] = object.QosAction

	return omap
}

func ResourceActionListTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"simple_action": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"gateway_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"apply_service": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"mirror_to": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: MirrorActionType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMirrorActionType(),
		},
		"assign_routing_instance": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"log": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"alert": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"qos_action": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceActionListType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceActionListTypeSchema(),
	}
}

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

func SetAlarmOperand2FromMap(object *AlarmOperand2, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAlarmOperand2FromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mUveAttributeObj := vmap["uve_attribute"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUveAttributeObj) {
		log.Printf("Setting uve_attribute  UveAttribute <<%T>> => %#v", mUveAttributeObj, mUveAttributeObj)
		mUveAttribute := mUveAttributeObj.(string)
		object.UveAttribute = mUveAttribute
	}

	mJsonValueObj := vmap["json_value"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mJsonValueObj) {
		log.Printf("Setting json_value  JsonValue <<%T>> => %#v", mJsonValueObj, mJsonValueObj)
		mJsonValue := mJsonValueObj.(string)
		object.JsonValue = mJsonValue
	}

	log.Printf("FINISHED AlarmOperand2 object: %#v", object)
}

func TakeAlarmOperand2AsMap(object *AlarmOperand2) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["uve_attribute"] = object.UveAttribute
	omap["json_value"] = object.JsonValue

	return omap
}

func ResourceAlarmOperand2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uve_attribute": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"json_value": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceAlarmOperand2() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAlarmOperand2Schema(),
	}
}

func SetAlarmExpressionFromMap(object *AlarmExpression, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAlarmExpressionFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mOperationObj := vmap["operation"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOperationObj) {
		log.Printf("Setting operation  Operation <<%T>> => %#v", mOperationObj, mOperationObj)
		mOperation := mOperationObj.(string)
		object.Operation = mOperation
	}

	mOperand1Obj := vmap["operand1"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOperand1Obj) {
		log.Printf("Setting operand1  Operand1 <<%T>> => %#v", mOperand1Obj, mOperand1Obj)
		mOperand1 := mOperand1Obj.(string)
		object.Operand1 = mOperand1
	}

	mOperand2Obj := vmap["operand2"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mOperand2Obj) {
		log.Printf("Setting operand2  Operand2 <<%T>> => %#v", mOperand2Obj, mOperand2Obj)
		mOperand2 := new(AlarmOperand2)
		SetAlarmOperand2FromMap(mOperand2, d, m, mOperand2Obj)
		object.Operand2 = mOperand2
	}

	mVariablesObj := vmap["variables"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mVariablesObj) {
		log.Printf("Setting variables  Variables <<%T>> => %#v", mVariablesObj, mVariablesObj)
		for _, v := range mVariablesObj.([]interface{}) {
			mVariables := v.(string)
			object.AddVariables(mVariables)
		}
	}

	log.Printf("FINISHED AlarmExpression object: %#v", object)
}

func TakeAlarmExpressionAsMap(object *AlarmExpression) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["operation"] = object.Operation
	omap["operand1"] = object.Operand1
	if object.Operand2 != nil {
		omap["operand2"] = TakeAlarmOperand2AsMap(object.Operand2)
	}
	variables_arr := make([]string, len(object.Variables))
	for _, v := range object.Variables {
		variables_arr = append(variables_arr, v)
	}
	omap["variables"] = variables_arr

	return omap
}

func ResourceAlarmExpressionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"operation": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"operand1": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"operand2": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: AlarmOperand2
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceAlarmOperand2(),
		},
		"variables": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceAlarmExpression() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAlarmExpressionSchema(),
	}
}

func SetAlarmAndListFromMap(object *AlarmAndList, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAlarmAndListFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAndListObj := vmap["and_list"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mAndListObj) {
		log.Printf("Setting and_list  AndList <<%T>> => %#v", mAndListObj, mAndListObj)
		for _, v := range mAndListObj.([]interface{}) {
			mAndList := new(AlarmExpression)
			SetAlarmExpressionFromMap(mAndList, d, m, v)
			object.AddAndList(mAndList)
		}
	}

	log.Printf("FINISHED AlarmAndList object: %#v", object)
}

func TakeAlarmAndListAsMap(object *AlarmAndList) map[string]interface{} {
	omap := make(map[string]interface{})

	var and_list_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.AndList {
		and_list_map = append(and_list_map, TakeAlarmExpressionAsMap(&v))
	}
	omap["and_list"] = and_list_map

	return omap
}

func ResourceAlarmAndListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"and_list": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AlarmExpression
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAlarmExpression(),
		},
	}
}

func ResourceAlarmAndList() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAlarmAndListSchema(),
	}
}

func SetAlarmOrListFromMap(object *AlarmOrList, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetAlarmOrListFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mOrListObj := vmap["or_list"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mOrListObj) {
		log.Printf("Setting or_list  OrList <<%T>> => %#v", mOrListObj, mOrListObj)
		for _, v := range mOrListObj.([]interface{}) {
			mOrList := new(AlarmAndList)
			SetAlarmAndListFromMap(mOrList, d, m, v)
			object.AddOrList(mOrList)
		}
	}

	log.Printf("FINISHED AlarmOrList object: %#v", object)
}

func TakeAlarmOrListAsMap(object *AlarmOrList) map[string]interface{} {
	omap := make(map[string]interface{})

	var or_list_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.OrList {
		or_list_map = append(or_list_map, TakeAlarmAndListAsMap(&v))
	}
	omap["or_list"] = or_list_map

	return omap
}

func ResourceAlarmOrListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"or_list": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AlarmAndList
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAlarmAndList(),
		},
	}
}

func ResourceAlarmOrList() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceAlarmOrListSchema(),
	}
}

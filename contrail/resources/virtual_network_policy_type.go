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

func SetTimerTypeFromMap(object *TimerType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetTimerTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mStartTimeObj := vmap["start_time"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mStartTimeObj) {
		log.Printf("Setting start_time  StartTime <<%T>> => %#v", mStartTimeObj, mStartTimeObj)
		mStartTime := mStartTimeObj.(string)
		object.StartTime = mStartTime
	}

	mOnIntervalObj := vmap["on_interval"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOnIntervalObj) {
		log.Printf("Setting on_interval  OnInterval <<%T>> => %#v", mOnIntervalObj, mOnIntervalObj)
		mOnInterval := mOnIntervalObj.(uint64)
		object.OnInterval = mOnInterval
	}

	mOffIntervalObj := vmap["off_interval"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOffIntervalObj) {
		log.Printf("Setting off_interval  OffInterval <<%T>> => %#v", mOffIntervalObj, mOffIntervalObj)
		mOffInterval := mOffIntervalObj.(uint64)
		object.OffInterval = mOffInterval
	}

	mEndTimeObj := vmap["end_time"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEndTimeObj) {
		log.Printf("Setting end_time  EndTime <<%T>> => %#v", mEndTimeObj, mEndTimeObj)
		mEndTime := mEndTimeObj.(string)
		object.EndTime = mEndTime
	}

	log.Printf("FINISHED TimerType object: %#v", object)
}

func TakeTimerTypeAsMap(object *TimerType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["start_time"] = object.StartTime
	omap["on_interval"] = object.OnInterval
	omap["off_interval"] = object.OffInterval
	omap["end_time"] = object.EndTime

	return omap
}

func ResourceTimerTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start_time": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:dateTime
			Required: true,
			Type:     schema.TypeString,
		},
		"on_interval": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:time
			Required: true,
			Type:     schema.TypeInt,
		},
		"off_interval": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:time
			Required: true,
			Type:     schema.TypeInt,
		},
		"end_time": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:dateTime
			Required: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceTimerType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceTimerTypeSchema(),
	}
}

func SetVirtualNetworkPolicyTypeFromMap(object *VirtualNetworkPolicyType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualNetworkPolicyTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSequenceObj := vmap["sequence"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSequenceObj) {
		log.Printf("Setting sequence  Sequence <<%T>> => %#v", mSequenceObj, mSequenceObj)
		mSequence := new(SequenceType)
		SetSequenceTypeFromMap(mSequence, d, m, mSequenceObj)
		object.Sequence = mSequence
	}

	mTimerObj := vmap["timer"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mTimerObj) {
		log.Printf("Setting timer  Timer <<%T>> => %#v", mTimerObj, mTimerObj)
		mTimer := new(TimerType)
		SetTimerTypeFromMap(mTimer, d, m, mTimerObj)
		object.Timer = mTimer
	}

	log.Printf("FINISHED VirtualNetworkPolicyType object: %#v", object)
}

func TakeVirtualNetworkPolicyTypeAsMap(object *VirtualNetworkPolicyType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Sequence != nil {
		omap["sequence"] = TakeSequenceTypeAsMap(object.Sequence)
	}
	if object.Timer != nil {
		omap["timer"] = TakeTimerTypeAsMap(object.Timer)
	}

	return omap
}

func ResourceVirtualNetworkPolicyTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sequence": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SequenceType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSequenceType(),
		},
		"timer": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: TimerType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTimerType(),
		},
	}
}

func ResourceVirtualNetworkPolicyType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualNetworkPolicyTypeSchema(),
	}
}

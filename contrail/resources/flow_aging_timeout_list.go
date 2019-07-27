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

func SetFlowAgingTimeoutFromMap(object *FlowAgingTimeout, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFlowAgingTimeoutFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mProtocolObj := vmap["protocol"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolObj) {
		log.Printf("Setting protocol  Protocol <<%T>> => %#v", mProtocolObj, mProtocolObj)
		mProtocol := mProtocolObj.(string)
		object.Protocol = mProtocol
	}

	mPortObj := vmap["port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPortObj) {
		log.Printf("Setting port  Port <<%T>> => %#v", mPortObj, mPortObj)
		mPort := mPortObj.(int)
		object.Port = mPort
	}

	mTimeoutInSecondsObj := vmap["timeout_in_seconds"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTimeoutInSecondsObj) {
		log.Printf("Setting timeout_in_seconds  TimeoutInSeconds <<%T>> => %#v", mTimeoutInSecondsObj, mTimeoutInSecondsObj)
		mTimeoutInSeconds := mTimeoutInSecondsObj.(int)
		object.TimeoutInSeconds = mTimeoutInSeconds
	}

	log.Printf("FINISHED FlowAgingTimeout object: %#v", object)
}

func TakeFlowAgingTimeoutAsMap(object *FlowAgingTimeout) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["protocol"] = object.Protocol
	omap["port"] = object.Port
	omap["timeout_in_seconds"] = object.TimeoutInSeconds

	return omap
}

func ResourceFlowAgingTimeoutSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"timeout_in_seconds": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceFlowAgingTimeout() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFlowAgingTimeoutSchema(),
	}
}

func SetFlowAgingTimeoutListFromMap(object *FlowAgingTimeoutList, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFlowAgingTimeoutListFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mFlowAgingTimeoutObj := vmap["flow_aging_timeout"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mFlowAgingTimeoutObj) {
		log.Printf("Setting flow_aging_timeout  FlowAgingTimeout <<%T>> => %#v", mFlowAgingTimeoutObj, mFlowAgingTimeoutObj)
		for _, v := range mFlowAgingTimeoutObj.([]interface{}) {
			mFlowAgingTimeout := new(FlowAgingTimeout)
			SetFlowAgingTimeoutFromMap(mFlowAgingTimeout, d, m, v)
			object.AddFlowAgingTimeout(mFlowAgingTimeout)
		}
	}

	log.Printf("FINISHED FlowAgingTimeoutList object: %#v", object)
}

func TakeFlowAgingTimeoutListAsMap(object *FlowAgingTimeoutList) map[string]interface{} {
	omap := make(map[string]interface{})

	var flow_aging_timeout_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.FlowAgingTimeout {
		flow_aging_timeout_map = append(flow_aging_timeout_map, TakeFlowAgingTimeoutAsMap(&v))
	}
	omap["flow_aging_timeout"] = flow_aging_timeout_map

	return omap
}

func ResourceFlowAgingTimeoutListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"flow_aging_timeout": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: FlowAgingTimeout
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFlowAgingTimeout(),
		},
	}
}

func ResourceFlowAgingTimeoutList() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFlowAgingTimeoutListSchema(),
	}
}

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

func SetQosIdForwardingClassPairFromMap(object *QosIdForwardingClassPair, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetQosIdForwardingClassPairFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mKeyObj := vmap["key"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mKeyObj) {
		log.Printf("Setting key  Key <<%T>> => %#v", mKeyObj, mKeyObj)
		mKey := mKeyObj.(int)
		object.Key = mKey
	}

	mForwardingClassIdObj := vmap["forwarding_class_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mForwardingClassIdObj) {
		log.Printf("Setting forwarding_class_id  ForwardingClassId <<%T>> => %#v", mForwardingClassIdObj, mForwardingClassIdObj)
		mForwardingClassId := mForwardingClassIdObj.(int)
		object.ForwardingClassId = mForwardingClassId
	}

	log.Printf("FINISHED QosIdForwardingClassPair object: %#v", object)
}

func TakeQosIdForwardingClassPairAsMap(object *QosIdForwardingClassPair) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["key"] = object.Key
	omap["forwarding_class_id"] = object.ForwardingClassId

	return omap
}

func ResourceQosIdForwardingClassPairSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"forwarding_class_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceQosIdForwardingClassPair() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceQosIdForwardingClassPairSchema(),
	}
}

func SetQosIdForwardingClassPairsFromMap(object *QosIdForwardingClassPairs, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetQosIdForwardingClassPairsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mQosIdForwardingClassPairObj := vmap["qos_id_forwarding_class_pair"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mQosIdForwardingClassPairObj) {
		log.Printf("Setting qos_id_forwarding_class_pair  QosIdForwardingClassPair <<%T>> => %#v", mQosIdForwardingClassPairObj, mQosIdForwardingClassPairObj)
		for _, v := range mQosIdForwardingClassPairObj.([]interface{}) {
			mQosIdForwardingClassPair := new(QosIdForwardingClassPair)
			SetQosIdForwardingClassPairFromMap(mQosIdForwardingClassPair, d, m, v)
			object.AddQosIdForwardingClassPair(mQosIdForwardingClassPair)
		}
	}

	log.Printf("FINISHED QosIdForwardingClassPairs object: %#v", object)
}

func TakeQosIdForwardingClassPairsAsMap(object *QosIdForwardingClassPairs) map[string]interface{} {
	omap := make(map[string]interface{})

	var qos_id_forwarding_class_pair_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.QosIdForwardingClassPair {
		qos_id_forwarding_class_pair_map = append(qos_id_forwarding_class_pair_map, TakeQosIdForwardingClassPairAsMap(&v))
	}
	omap["qos_id_forwarding_class_pair"] = qos_id_forwarding_class_pair_map

	return omap
}

func ResourceQosIdForwardingClassPairsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"qos_id_forwarding_class_pair": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: QosIdForwardingClassPair
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceQosIdForwardingClassPair(),
		},
	}
}

func ResourceQosIdForwardingClassPairs() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceQosIdForwardingClassPairsSchema(),
	}
}

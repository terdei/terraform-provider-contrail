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

func SetDiscoveryPubSubEndPointTypeFromMap(object *DiscoveryPubSubEndPointType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetDiscoveryPubSubEndPointTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mEpTypeObj := vmap["ep_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEpTypeObj) {
		log.Printf("Setting ep_type  EpType <<%T>> => %#v", mEpTypeObj, mEpTypeObj)
		mEpType := mEpTypeObj.(string)
		object.EpType = mEpType
	}

	mEpIdObj := vmap["ep_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEpIdObj) {
		log.Printf("Setting ep_id  EpId <<%T>> => %#v", mEpIdObj, mEpIdObj)
		mEpId := mEpIdObj.(string)
		object.EpId = mEpId
	}

	mEpPrefixObj := vmap["ep_prefix"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mEpPrefixObj) {
		log.Printf("Setting ep_prefix  EpPrefix <<%T>> => %#v", mEpPrefixObj, mEpPrefixObj)
		mEpPrefix := new(SubnetType)
		SetSubnetTypeFromMap(mEpPrefix, d, m, mEpPrefixObj)
		object.EpPrefix = mEpPrefix
	}

	mEpVersionObj := vmap["ep_version"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEpVersionObj) {
		log.Printf("Setting ep_version  EpVersion <<%T>> => %#v", mEpVersionObj, mEpVersionObj)
		mEpVersion := mEpVersionObj.(string)
		object.EpVersion = mEpVersion
	}

	log.Printf("FINISHED DiscoveryPubSubEndPointType object: %#v", object)
}

func TakeDiscoveryPubSubEndPointTypeAsMap(object *DiscoveryPubSubEndPointType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["ep_type"] = object.EpType
	omap["ep_id"] = object.EpId
	if object.EpPrefix != nil {
		omap["ep_prefix"] = TakeSubnetTypeAsMap(object.EpPrefix)
	}
	omap["ep_version"] = object.EpVersion

	return omap
}

func ResourceDiscoveryPubSubEndPointTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ep_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"ep_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ep_prefix": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
		"ep_version": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceDiscoveryPubSubEndPointType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceDiscoveryPubSubEndPointTypeSchema(),
	}
}

func SetDiscoveryServiceAssignmentTypeFromMap(object *DiscoveryServiceAssignmentType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetDiscoveryServiceAssignmentTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPublisherObj := vmap["publisher"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mPublisherObj) {
		log.Printf("Setting publisher  Publisher <<%T>> => %#v", mPublisherObj, mPublisherObj)
		mPublisher := new(DiscoveryPubSubEndPointType)
		SetDiscoveryPubSubEndPointTypeFromMap(mPublisher, d, m, mPublisherObj)
		object.Publisher = mPublisher
	}

	mSubscriberObj := vmap["subscriber"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSubscriberObj) {
		log.Printf("Setting subscriber  Subscriber <<%T>> => %#v", mSubscriberObj, mSubscriberObj)
		for _, v := range mSubscriberObj.([]interface{}) {
			mSubscriber := new(DiscoveryPubSubEndPointType)
			SetDiscoveryPubSubEndPointTypeFromMap(mSubscriber, d, m, v)
			object.AddSubscriber(mSubscriber)
		}
	}

	log.Printf("FINISHED DiscoveryServiceAssignmentType object: %#v", object)
}

func TakeDiscoveryServiceAssignmentTypeAsMap(object *DiscoveryServiceAssignmentType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Publisher != nil {
		omap["publisher"] = TakeDiscoveryPubSubEndPointTypeAsMap(object.Publisher)
	}
	var subscriber_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Subscriber {
		subscriber_map = append(subscriber_map, TakeDiscoveryPubSubEndPointTypeAsMap(&v))
	}
	omap["subscriber"] = subscriber_map

	return omap
}

func ResourceDiscoveryServiceAssignmentTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"publisher": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: DiscoveryPubSubEndPointType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceDiscoveryPubSubEndPointType(),
		},
		"subscriber": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: DiscoveryPubSubEndPointType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceDiscoveryPubSubEndPointType(),
		},
	}
}

func ResourceDiscoveryServiceAssignmentType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceDiscoveryServiceAssignmentTypeSchema(),
	}
}

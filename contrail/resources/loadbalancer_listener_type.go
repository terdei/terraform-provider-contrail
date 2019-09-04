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

func SetLoadbalancerListenerTypeFromMap(object *LoadbalancerListenerType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLoadbalancerListenerTypeFromMAP] map = %#v", val)

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

	mProtocolPortObj := vmap["protocol_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolPortObj) {
		log.Printf("Setting protocol_port  ProtocolPort <<%T>> => %#v", mProtocolPortObj, mProtocolPortObj)
		mProtocolPort := mProtocolPortObj.(int)
		object.ProtocolPort = mProtocolPort
	}

	mAdminStateObj := vmap["admin_state"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAdminStateObj) {
		log.Printf("Setting admin_state  AdminState <<%T>> => %#v", mAdminStateObj, mAdminStateObj)
		mAdminState := mAdminStateObj.(bool)
		object.AdminState = mAdminState
	}

	mConnectionLimitObj := vmap["connection_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mConnectionLimitObj) {
		log.Printf("Setting connection_limit  ConnectionLimit <<%T>> => %#v", mConnectionLimitObj, mConnectionLimitObj)
		mConnectionLimit := mConnectionLimitObj.(int)
		object.ConnectionLimit = mConnectionLimit
	}

	mDefaultTlsContainerObj := vmap["default_tls_container"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDefaultTlsContainerObj) {
		log.Printf("Setting default_tls_container  DefaultTlsContainer <<%T>> => %#v", mDefaultTlsContainerObj, mDefaultTlsContainerObj)
		mDefaultTlsContainer := mDefaultTlsContainerObj.(string)
		object.DefaultTlsContainer = mDefaultTlsContainer
	}

	mSniContainersObj := vmap["sni_containers"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mSniContainersObj) {
		log.Printf("Setting sni_containers  SniContainers <<%T>> => %#v", mSniContainersObj, mSniContainersObj)
		for _, v := range mSniContainersObj.([]interface{}) {
			mSniContainers := v.(string)
			object.AddSniContainers(mSniContainers)
		}
	}

	log.Printf("FINISHED LoadbalancerListenerType object: %#v", object)
}

func TakeLoadbalancerListenerTypeAsMap(object *LoadbalancerListenerType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["protocol"] = object.Protocol
	omap["protocol_port"] = object.ProtocolPort
	omap["admin_state"] = object.AdminState
	omap["connection_limit"] = object.ConnectionLimit
	omap["default_tls_container"] = object.DefaultTlsContainer
	sni_containers_arr := make([]string, len(object.SniContainers))
	for _, v := range object.SniContainers {
		sni_containers_arr = append(sni_containers_arr, v)
	}
	omap["sni_containers"] = sni_containers_arr

	return omap
}

func ResourceLoadbalancerListenerTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"protocol_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"admin_state": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"connection_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"default_tls_container": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"sni_containers": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceLoadbalancerListenerType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLoadbalancerListenerTypeSchema(),
	}
}

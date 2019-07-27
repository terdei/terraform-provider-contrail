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

func SetMirrorActionTypeFromMap(object *MirrorActionType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetMirrorActionTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAnalyzerNameObj := vmap["analyzer_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAnalyzerNameObj) {
		log.Printf("Setting analyzer_name  AnalyzerName <<%T>> => %#v", mAnalyzerNameObj, mAnalyzerNameObj)
		mAnalyzerName := mAnalyzerNameObj.(string)
		object.AnalyzerName = mAnalyzerName
	}

	mEncapsulationObj := vmap["encapsulation"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEncapsulationObj) {
		log.Printf("Setting encapsulation  Encapsulation <<%T>> => %#v", mEncapsulationObj, mEncapsulationObj)
		mEncapsulation := mEncapsulationObj.(string)
		object.Encapsulation = mEncapsulation
	}

	mAnalyzerIpAddressObj := vmap["analyzer_ip_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAnalyzerIpAddressObj) {
		log.Printf("Setting analyzer_ip_address  AnalyzerIpAddress <<%T>> => %#v", mAnalyzerIpAddressObj, mAnalyzerIpAddressObj)
		mAnalyzerIpAddress := mAnalyzerIpAddressObj.(string)
		object.AnalyzerIpAddress = mAnalyzerIpAddress
	}

	mAnalyzerMacAddressObj := vmap["analyzer_mac_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAnalyzerMacAddressObj) {
		log.Printf("Setting analyzer_mac_address  AnalyzerMacAddress <<%T>> => %#v", mAnalyzerMacAddressObj, mAnalyzerMacAddressObj)
		mAnalyzerMacAddress := mAnalyzerMacAddressObj.(string)
		object.AnalyzerMacAddress = mAnalyzerMacAddress
	}

	mRoutingInstanceObj := vmap["routing_instance"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRoutingInstanceObj) {
		log.Printf("Setting routing_instance  RoutingInstance <<%T>> => %#v", mRoutingInstanceObj, mRoutingInstanceObj)
		mRoutingInstance := mRoutingInstanceObj.(string)
		object.RoutingInstance = mRoutingInstance
	}

	mUdpPortObj := vmap["udp_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUdpPortObj) {
		log.Printf("Setting udp_port  UdpPort <<%T>> => %#v", mUdpPortObj, mUdpPortObj)
		mUdpPort := mUdpPortObj.(int)
		object.UdpPort = mUdpPort
	}

	mJuniperHeaderObj := vmap["juniper_header"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mJuniperHeaderObj) {
		log.Printf("Setting juniper_header  JuniperHeader <<%T>> => %#v", mJuniperHeaderObj, mJuniperHeaderObj)
		mJuniperHeader := mJuniperHeaderObj.(bool)
		object.JuniperHeader = mJuniperHeader
	}

	mNhModeObj := vmap["nh_mode"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNhModeObj) {
		log.Printf("Setting nh_mode  NhMode <<%T>> => %#v", mNhModeObj, mNhModeObj)
		mNhMode := mNhModeObj.(string)
		object.NhMode = mNhMode
	}

	mStaticNhHeaderObj := vmap["static_nh_header"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mStaticNhHeaderObj) {
		log.Printf("Setting static_nh_header  StaticNhHeader <<%T>> => %#v", mStaticNhHeaderObj, mStaticNhHeaderObj)
		mStaticNhHeader := new(StaticMirrorNhType)
		SetStaticMirrorNhTypeFromMap(mStaticNhHeader, d, m, mStaticNhHeaderObj)
		object.StaticNhHeader = mStaticNhHeader
	}

	log.Printf("FINISHED MirrorActionType object: %#v", object)
}

func TakeMirrorActionTypeAsMap(object *MirrorActionType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["analyzer_name"] = object.AnalyzerName
	omap["encapsulation"] = object.Encapsulation
	omap["analyzer_ip_address"] = object.AnalyzerIpAddress
	omap["analyzer_mac_address"] = object.AnalyzerMacAddress
	omap["routing_instance"] = object.RoutingInstance
	omap["udp_port"] = object.UdpPort
	omap["juniper_header"] = object.JuniperHeader
	omap["nh_mode"] = object.NhMode
	if object.StaticNhHeader != nil {
		omap["static_nh_header"] = TakeStaticMirrorNhTypeAsMap(object.StaticNhHeader)
	}

	return omap
}

func ResourceMirrorActionTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"analyzer_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"encapsulation": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"analyzer_ip_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"analyzer_mac_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"routing_instance": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"udp_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"juniper_header": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"nh_mode": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"static_nh_header": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: StaticMirrorNhType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceStaticMirrorNhType(),
		},
	}
}

func ResourceMirrorActionType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceMirrorActionTypeSchema(),
	}
}

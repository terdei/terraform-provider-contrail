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

func SetFirewallServiceTypeFromMap(object *FirewallServiceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFirewallServiceTypeFromMAP] map = %#v", val)

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

	mProtocolIdObj := vmap["protocol_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProtocolIdObj) {
		log.Printf("Setting protocol_id  ProtocolId <<%T>> => %#v", mProtocolIdObj, mProtocolIdObj)
		mProtocolId := mProtocolIdObj.(int)
		object.ProtocolId = mProtocolId
	}

	mSrcPortsObj := vmap["src_ports"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSrcPortsObj) {
		log.Printf("Setting src_ports  SrcPorts <<%T>> => %#v", mSrcPortsObj, mSrcPortsObj)
		mSrcPorts := new(PortType)
		SetPortTypeFromMap(mSrcPorts, d, m, mSrcPortsObj)
		object.SrcPorts = mSrcPorts
	}

	mDstPortsObj := vmap["dst_ports"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mDstPortsObj) {
		log.Printf("Setting dst_ports  DstPorts <<%T>> => %#v", mDstPortsObj, mDstPortsObj)
		mDstPorts := new(PortType)
		SetPortTypeFromMap(mDstPorts, d, m, mDstPortsObj)
		object.DstPorts = mDstPorts
	}

	log.Printf("FINISHED FirewallServiceType object: %#v", object)
}

func TakeFirewallServiceTypeAsMap(object *FirewallServiceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["protocol"] = object.Protocol
	omap["protocol_id"] = object.ProtocolId
	if object.SrcPorts != nil {
		omap["src_ports"] = TakePortTypeAsMap(object.SrcPorts)
	}
	if object.DstPorts != nil {
		omap["dst_ports"] = TakePortTypeAsMap(object.DstPorts)
	}

	return omap
}

func ResourceFirewallServiceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"protocol_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"src_ports": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: PortType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortType(),
		},
		"dst_ports": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: PortType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortType(),
		},
	}
}

func ResourceFirewallServiceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFirewallServiceTypeSchema(),
	}
}

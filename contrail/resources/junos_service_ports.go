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

func SetJunosServicePortsFromMap(object *JunosServicePorts, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetJunosServicePortsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mServicePortObj := vmap["service_port"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mServicePortObj) {
		log.Printf("Setting service_port  ServicePort <<%T>> => %#v", mServicePortObj, mServicePortObj)
		for _, v := range mServicePortObj.([]interface{}) {
			mServicePort := v.(string)
			object.AddServicePort(mServicePort)
		}
	}

	log.Printf("FINISHED JunosServicePorts object: %#v", object)
}

func TakeJunosServicePortsAsMap(object *JunosServicePorts) map[string]interface{} {
	omap := make(map[string]interface{})

	service_port_arr := make([]string, len(object.ServicePort))
	for _, v := range object.ServicePort {
		service_port_arr = append(service_port_arr, v)
	}
	omap["service_port"] = service_port_arr

	return omap
}

func ResourceJunosServicePortsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_port": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceJunosServicePorts() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceJunosServicePortsSchema(),
	}
}

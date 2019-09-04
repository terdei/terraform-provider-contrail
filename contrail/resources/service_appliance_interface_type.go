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

func SetServiceApplianceInterfaceTypeFromMap(object *ServiceApplianceInterfaceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceApplianceInterfaceTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mInterfaceTypeObj := vmap["interface_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mInterfaceTypeObj) {
		log.Printf("Setting interface_type  InterfaceType <<%T>> => %#v", mInterfaceTypeObj, mInterfaceTypeObj)
		mInterfaceType := mInterfaceTypeObj.(string)
		object.InterfaceType = mInterfaceType
	}

	log.Printf("FINISHED ServiceApplianceInterfaceType object: %#v", object)
}

func TakeServiceApplianceInterfaceTypeAsMap(object *ServiceApplianceInterfaceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["interface_type"] = object.InterfaceType

	return omap
}

func ResourceServiceApplianceInterfaceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interface_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceServiceApplianceInterfaceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceApplianceInterfaceTypeSchema(),
	}
}

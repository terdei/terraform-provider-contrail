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

func SetEncapsulationPrioritiesTypeFromMap(object *EncapsulationPrioritiesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetEncapsulationPrioritiesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mEncapsulationObj := vmap["encapsulation"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mEncapsulationObj) {
		log.Printf("Setting encapsulation  Encapsulation <<%T>> => %#v", mEncapsulationObj, mEncapsulationObj)
		for _, v := range mEncapsulationObj.([]interface{}) {
			mEncapsulation := v.(string)
			object.AddEncapsulation(mEncapsulation)
		}
	}

	log.Printf("FINISHED EncapsulationPrioritiesType object: %#v", object)
}

func TakeEncapsulationPrioritiesTypeAsMap(object *EncapsulationPrioritiesType) map[string]interface{} {
	omap := make(map[string]interface{})

	encapsulation_arr := make([]string, len(object.Encapsulation))
	for _, v := range object.Encapsulation {
		encapsulation_arr = append(encapsulation_arr, v)
	}
	omap["encapsulation"] = encapsulation_arr

	return omap
}

func ResourceEncapsulationPrioritiesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"encapsulation": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceEncapsulationPrioritiesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceEncapsulationPrioritiesTypeSchema(),
	}
}

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

func SetSubnetListTypeFromMap(object *SubnetListType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSubnetListTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSubnetObj := vmap["subnet"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSubnetObj) {
		log.Printf("Setting subnet  Subnet <<%T>> => %#v", mSubnetObj, mSubnetObj)
		for _, v := range mSubnetObj.([]interface{}) {
			mSubnet := new(SubnetType)
			SetSubnetTypeFromMap(mSubnet, d, m, v)
			object.AddSubnet(mSubnet)
		}
	}

	log.Printf("FINISHED SubnetListType object: %#v", object)
}

func TakeSubnetListTypeAsMap(object *SubnetListType) map[string]interface{} {
	omap := make(map[string]interface{})

	var subnet_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Subnet {
		subnet_map = append(subnet_map, TakeSubnetTypeAsMap(&v))
	}
	omap["subnet"] = subnet_map

	return omap
}

func ResourceSubnetListTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
	}
}

func ResourceSubnetListType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSubnetListTypeSchema(),
	}
}

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

func SetFloatingIpPoolTypeFromMap(object *FloatingIpPoolType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFloatingIpPoolTypeFromMAP] map = %#v", val)

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

	log.Printf("FINISHED FloatingIpPoolType object: %#v", object)
}

func TakeFloatingIpPoolTypeAsMap(object *FloatingIpPoolType) map[string]interface{} {
	omap := make(map[string]interface{})

	var subnet_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Subnet {
		subnet_map = append(subnet_map, TakeSubnetTypeAsMap(&v))
	}
	omap["subnet"] = subnet_map

	return omap
}

func ResourceFloatingIpPoolTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
	}
}

func ResourceFloatingIpPoolType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFloatingIpPoolTypeSchema(),
	}
}

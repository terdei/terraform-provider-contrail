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

func SetIpamSubnetsFromMap(object *IpamSubnets, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetIpamSubnetsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSubnetsObj := vmap["subnets"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSubnetsObj) {
		log.Printf("Setting subnets  Subnets <<%T>> => %#v", mSubnetsObj, mSubnetsObj)
		for _, v := range mSubnetsObj.([]interface{}) {
			mSubnets := new(IpamSubnetType)
			SetIpamSubnetTypeFromMap(mSubnets, d, m, v)
			object.AddSubnets(mSubnets)
		}
	}

	log.Printf("FINISHED IpamSubnets object: %#v", object)
}

func TakeIpamSubnetsAsMap(object *IpamSubnets) map[string]interface{} {
	omap := make(map[string]interface{})

	var subnets_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Subnets {
		subnets_map = append(subnets_map, TakeIpamSubnetTypeAsMap(&v))
	}
	omap["subnets"] = subnets_map

	return omap
}

func ResourceIpamSubnetsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnets": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: IpamSubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIpamSubnetType(),
		},
	}
}

func ResourceIpamSubnets() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceIpamSubnetsSchema(),
	}
}

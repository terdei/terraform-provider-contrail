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

func SetFloatingIpPoolSubnetTypeFromMap(object *FloatingIpPoolSubnetType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFloatingIpPoolSubnetTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSubnetUuidObj := vmap["subnet_uuid"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mSubnetUuidObj) {
		log.Printf("Setting subnet_uuid  SubnetUuid <<%T>> => %#v", mSubnetUuidObj, mSubnetUuidObj)
		for _, v := range mSubnetUuidObj.([]interface{}) {
			mSubnetUuid := v.(string)
			object.AddSubnetUuid(mSubnetUuid)
		}
	}

	log.Printf("FINISHED FloatingIpPoolSubnetType object: %#v", object)
}

func TakeFloatingIpPoolSubnetTypeAsMap(object *FloatingIpPoolSubnetType) map[string]interface{} {
	omap := make(map[string]interface{})

	subnet_uuid_arr := make([]string, len(object.SubnetUuid))
	for _, v := range object.SubnetUuid {
		subnet_uuid_arr = append(subnet_uuid_arr, v)
	}
	omap["subnet_uuid"] = subnet_uuid_arr

	return omap
}

func ResourceFloatingIpPoolSubnetTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet_uuid": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceFloatingIpPoolSubnetType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFloatingIpPoolSubnetTypeSchema(),
	}
}

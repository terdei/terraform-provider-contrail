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

func SetVnSubnetsTypeFromMap(object *VnSubnetsType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVnSubnetsTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mIpamSubnetsObj := vmap["ipam_subnets"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mIpamSubnetsObj) {
		log.Printf("Setting ipam_subnets  IpamSubnets <<%T>> => %#v", mIpamSubnetsObj, mIpamSubnetsObj)
		for _, v := range mIpamSubnetsObj.([]interface{}) {
			mIpamSubnets := new(IpamSubnetType)
			SetIpamSubnetTypeFromMap(mIpamSubnets, d, m, v)
			object.AddIpamSubnets(mIpamSubnets)
		}
	}

	mHostRoutesObj := vmap["host_routes"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mHostRoutesObj) {
		log.Printf("Setting host_routes  HostRoutes <<%T>> => %#v", mHostRoutesObj, mHostRoutesObj)
		mHostRoutes := new(RouteTableType)
		SetRouteTableTypeFromMap(mHostRoutes, d, m, mHostRoutesObj)
		object.HostRoutes = mHostRoutes
	}

	log.Printf("FINISHED VnSubnetsType object: %#v", object)
}

func TakeVnSubnetsTypeAsMap(object *VnSubnetsType) map[string]interface{} {
	omap := make(map[string]interface{})

	var ipam_subnets_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.IpamSubnets {
		ipam_subnets_map = append(ipam_subnets_map, TakeIpamSubnetTypeAsMap(&v))
	}
	omap["ipam_subnets"] = ipam_subnets_map
	if object.HostRoutes != nil {
		omap["host_routes"] = TakeRouteTableTypeAsMap(object.HostRoutes)
	}

	return omap
}

func ResourceVnSubnetsTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipam_subnets": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: IpamSubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIpamSubnetType(),
		},
		"host_routes": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: RouteTableType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTableType(),
		},
	}
}

func ResourceVnSubnetsType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVnSubnetsTypeSchema(),
	}
}

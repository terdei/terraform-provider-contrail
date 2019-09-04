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

func SetVirtualRouterNetworkIpamTypeFromMap(object *VirtualRouterNetworkIpamType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualRouterNetworkIpamTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAllocationPoolsObj := vmap["allocation_pools"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mAllocationPoolsObj) {
		log.Printf("Setting allocation_pools  AllocationPools <<%T>> => %#v", mAllocationPoolsObj, mAllocationPoolsObj)
		for _, v := range mAllocationPoolsObj.([]interface{}) {
			mAllocationPools := new(AllocationPoolType)
			SetAllocationPoolTypeFromMap(mAllocationPools, d, m, v)
			object.AddAllocationPools(mAllocationPools)
		}
	}

	mSubnetObj := vmap["subnet"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mSubnetObj) {
		log.Printf("Setting subnet  Subnet <<%T>> => %#v", mSubnetObj, mSubnetObj)
		for _, v := range mSubnetObj.([]interface{}) {
			mSubnet := new(SubnetType)
			SetSubnetTypeFromMap(mSubnet, d, m, v)
			object.AddSubnet(mSubnet)
		}
	}

	log.Printf("FINISHED VirtualRouterNetworkIpamType object: %#v", object)
}

func TakeVirtualRouterNetworkIpamTypeAsMap(object *VirtualRouterNetworkIpamType) map[string]interface{} {
	omap := make(map[string]interface{})

	var allocation_pools_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.AllocationPools {
		allocation_pools_map = append(allocation_pools_map, TakeAllocationPoolTypeAsMap(&v))
	}
	omap["allocation_pools"] = allocation_pools_map
	var subnet_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Subnet {
		subnet_map = append(subnet_map, TakeSubnetTypeAsMap(&v))
	}
	omap["subnet"] = subnet_map

	return omap
}

func ResourceVirtualRouterNetworkIpamTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocation_pools": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AllocationPoolType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAllocationPoolType(),
		},
		"subnet": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
	}
}

func ResourceVirtualRouterNetworkIpamType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualRouterNetworkIpamTypeSchema(),
	}
}

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

func SetRouteTargetListFromMap(object *RouteTargetList, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRouteTargetListFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRouteTargetObj := vmap["route_target"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mRouteTargetObj) {
		log.Printf("Setting route_target  RouteTarget <<%T>> => %#v", mRouteTargetObj, mRouteTargetObj)
		for _, v := range mRouteTargetObj.([]interface{}) {
			mRouteTarget := v.(string)
			object.AddRouteTarget(mRouteTarget)
		}
	}

	log.Printf("FINISHED RouteTargetList object: %#v", object)
}

func TakeRouteTargetListAsMap(object *RouteTargetList) map[string]interface{} {
	omap := make(map[string]interface{})

	route_target_arr := make([]string, len(object.RouteTarget))
	for _, v := range object.RouteTarget {
		route_target_arr = append(route_target_arr, v)
	}
	omap["route_target"] = route_target_arr

	return omap
}

func ResourceRouteTargetListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"route_target": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceRouteTargetList() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRouteTargetListSchema(),
	}
}

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

func SetRouteTableTypeFromMap(object *RouteTableType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetRouteTableTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mRouteObj := vmap["route"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mRouteObj) {
		log.Printf("Setting route  Route <<%T>> => %#v", mRouteObj, mRouteObj)
		for _, v := range mRouteObj.([]interface{}) {
			mRoute := new(RouteType)
			SetRouteTypeFromMap(mRoute, d, m, v)
			object.AddRoute(mRoute)
		}
	}

	log.Printf("FINISHED RouteTableType object: %#v", object)
}

func TakeRouteTableTypeAsMap(object *RouteTableType) map[string]interface{} {
	omap := make(map[string]interface{})

	var route_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Route {
		route_map = append(route_map, TakeRouteTypeAsMap(&v))
	}
	omap["route"] = route_map

	return omap
}

func ResourceRouteTableTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"route": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: RouteType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteType(),
		},
	}
}

func ResourceRouteTableType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceRouteTableTypeSchema(),
	}
}

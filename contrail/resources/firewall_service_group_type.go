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

func SetFirewallServiceGroupTypeFromMap(object *FirewallServiceGroupType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetFirewallServiceGroupTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mFirewallServiceObj := vmap["firewall_service"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mFirewallServiceObj) {
		log.Printf("Setting firewall_service  FirewallService <<%T>> => %#v", mFirewallServiceObj, mFirewallServiceObj)
		for _, v := range mFirewallServiceObj.([]interface{}) {
			mFirewallService := new(FirewallServiceType)
			SetFirewallServiceTypeFromMap(mFirewallService, d, m, v)
			object.AddFirewallService(mFirewallService)
		}
	}

	log.Printf("FINISHED FirewallServiceGroupType object: %#v", object)
}

func TakeFirewallServiceGroupTypeAsMap(object *FirewallServiceGroupType) map[string]interface{} {
	omap := make(map[string]interface{})

	var firewall_service_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.FirewallService {
		firewall_service_map = append(firewall_service_map, TakeFirewallServiceTypeAsMap(&v))
	}
	omap["firewall_service"] = firewall_service_map

	return omap
}

func ResourceFirewallServiceGroupTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"firewall_service": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: FirewallServiceType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallServiceType(),
		},
	}
}

func ResourceFirewallServiceGroupType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceFirewallServiceGroupTypeSchema(),
	}
}

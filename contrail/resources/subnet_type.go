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

func SetSubnetTypeFromMap(object *SubnetType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetSubnetTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mIpPrefixObj := vmap["ip_prefix"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpPrefixObj) {
		log.Printf("Setting ip_prefix  IpPrefix <<%T>> => %#v", mIpPrefixObj, mIpPrefixObj)
		mIpPrefix := mIpPrefixObj.(string)
		object.IpPrefix = mIpPrefix
	}

	mIpPrefixLenObj := vmap["ip_prefix_len"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpPrefixLenObj) {
		log.Printf("Setting ip_prefix_len  IpPrefixLen <<%T>> => %#v", mIpPrefixLenObj, mIpPrefixLenObj)
		mIpPrefixLen := mIpPrefixLenObj.(int)
		object.IpPrefixLen = mIpPrefixLen
	}

	log.Printf("FINISHED SubnetType object: %#v", object)
}

func TakeSubnetTypeAsMap(object *SubnetType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["ip_prefix"] = object.IpPrefix
	omap["ip_prefix_len"] = object.IpPrefixLen

	return omap
}

func ResourceSubnetTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip_prefix": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ip_prefix_len": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceSubnetType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceSubnetTypeSchema(),
	}
}

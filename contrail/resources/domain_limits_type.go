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

func SetDomainLimitsTypeFromMap(object *DomainLimitsType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetDomainLimitsTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mProjectLimitObj := vmap["project_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mProjectLimitObj) {
		log.Printf("Setting project_limit  ProjectLimit <<%T>> => %#v", mProjectLimitObj, mProjectLimitObj)
		mProjectLimit := mProjectLimitObj.(int)
		object.ProjectLimit = mProjectLimit
	}

	mVirtualNetworkLimitObj := vmap["virtual_network_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualNetworkLimitObj) {
		log.Printf("Setting virtual_network_limit  VirtualNetworkLimit <<%T>> => %#v", mVirtualNetworkLimitObj, mVirtualNetworkLimitObj)
		mVirtualNetworkLimit := mVirtualNetworkLimitObj.(int)
		object.VirtualNetworkLimit = mVirtualNetworkLimit
	}

	mSecurityGroupLimitObj := vmap["security_group_limit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSecurityGroupLimitObj) {
		log.Printf("Setting security_group_limit  SecurityGroupLimit <<%T>> => %#v", mSecurityGroupLimitObj, mSecurityGroupLimitObj)
		mSecurityGroupLimit := mSecurityGroupLimitObj.(int)
		object.SecurityGroupLimit = mSecurityGroupLimit
	}

	log.Printf("FINISHED DomainLimitsType object: %#v", object)
}

func TakeDomainLimitsTypeAsMap(object *DomainLimitsType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["project_limit"] = object.ProjectLimit
	omap["virtual_network_limit"] = object.VirtualNetworkLimit
	omap["security_group_limit"] = object.SecurityGroupLimit

	return omap
}

func ResourceDomainLimitsTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"project_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_network_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"security_group_limit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceDomainLimitsType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceDomainLimitsTypeSchema(),
	}
}

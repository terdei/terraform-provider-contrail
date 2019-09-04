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

func SetShareTypeFromMap(object *ShareType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetShareTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mTenantObj := vmap["tenant"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTenantObj) {
		log.Printf("Setting tenant  Tenant <<%T>> => %#v", mTenantObj, mTenantObj)
		mTenant := mTenantObj.(string)
		object.Tenant = mTenant
	}

	mTenantAccessObj := vmap["tenant_access"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTenantAccessObj) {
		log.Printf("Setting tenant_access  TenantAccess <<%T>> => %#v", mTenantAccessObj, mTenantAccessObj)
		mTenantAccess := mTenantAccessObj.(int)
		object.TenantAccess = mTenantAccess
	}

	log.Printf("FINISHED ShareType object: %#v", object)
}

func TakeShareTypeAsMap(object *ShareType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["tenant"] = object.Tenant
	omap["tenant_access"] = object.TenantAccess

	return omap
}

func ResourceShareTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tenant": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"tenant_access": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceShareType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceShareTypeSchema(),
	}
}

func SetPermType2FromMap(object *PermType2, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPermType2FromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mOwnerObj := vmap["owner"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOwnerObj) {
		log.Printf("Setting owner  Owner <<%T>> => %#v", mOwnerObj, mOwnerObj)
		mOwner := mOwnerObj.(string)
		object.Owner = mOwner
	}

	mOwnerAccessObj := vmap["owner_access"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOwnerAccessObj) {
		log.Printf("Setting owner_access  OwnerAccess <<%T>> => %#v", mOwnerAccessObj, mOwnerAccessObj)
		mOwnerAccess := mOwnerAccessObj.(int)
		object.OwnerAccess = mOwnerAccess
	}

	mGlobalAccessObj := vmap["global_access"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mGlobalAccessObj) {
		log.Printf("Setting global_access  GlobalAccess <<%T>> => %#v", mGlobalAccessObj, mGlobalAccessObj)
		mGlobalAccess := mGlobalAccessObj.(int)
		object.GlobalAccess = mGlobalAccess
	}

	mShareObj := vmap["share"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mShareObj) {
		log.Printf("Setting share  Share <<%T>> => %#v", mShareObj, mShareObj)
		for _, v := range mShareObj.([]interface{}) {
			mShare := new(ShareType)
			SetShareTypeFromMap(mShare, d, m, v)
			object.AddShare(mShare)
		}
	}

	log.Printf("FINISHED PermType2 object: %#v", object)
}

func TakePermType2AsMap(object *PermType2) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["owner"] = object.Owner
	omap["owner_access"] = object.OwnerAccess
	omap["global_access"] = object.GlobalAccess
	var share_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Share {
		share_map = append(share_map, TakeShareTypeAsMap(&v))
	}
	omap["share"] = share_map

	return omap
}

func ResourcePermType2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"owner": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"owner_access": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"global_access": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"share": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: ShareType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceShareType(),
		},
	}
}

func ResourcePermType2() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePermType2Schema(),
	}
}

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

func SetPermTypeFromMap(object *PermType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetPermTypeFromMAP] map = %#v", val)

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

	mGroupObj := vmap["group"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mGroupObj) {
		log.Printf("Setting group  Group <<%T>> => %#v", mGroupObj, mGroupObj)
		mGroup := mGroupObj.(string)
		object.Group = mGroup
	}

	mGroupAccessObj := vmap["group_access"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mGroupAccessObj) {
		log.Printf("Setting group_access  GroupAccess <<%T>> => %#v", mGroupAccessObj, mGroupAccessObj)
		mGroupAccess := mGroupAccessObj.(int)
		object.GroupAccess = mGroupAccess
	}

	mOtherAccessObj := vmap["other_access"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOtherAccessObj) {
		log.Printf("Setting other_access  OtherAccess <<%T>> => %#v", mOtherAccessObj, mOtherAccessObj)
		mOtherAccess := mOtherAccessObj.(int)
		object.OtherAccess = mOtherAccess
	}

	log.Printf("FINISHED PermType object: %#v", object)
}

func TakePermTypeAsMap(object *PermType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["owner"] = object.Owner
	omap["owner_access"] = object.OwnerAccess
	omap["group"] = object.Group
	omap["group_access"] = object.GroupAccess
	omap["other_access"] = object.OtherAccess

	return omap
}

func ResourcePermTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"owner": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"owner_access": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"group": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"group_access": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"other_access": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourcePermType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourcePermTypeSchema(),
	}
}

func SetUuidTypeFromMap(object *UuidType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetUuidTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mUuidMslongObj := vmap["uuid_mslong"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUuidMslongObj) {
		log.Printf("Setting uuid_mslong  UuidMslong <<%T>> => %#v", mUuidMslongObj, mUuidMslongObj)
		mUuidMslong := mUuidMslongObj.(uint64)
		object.UuidMslong = mUuidMslong
	}

	mUuidLslongObj := vmap["uuid_lslong"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUuidLslongObj) {
		log.Printf("Setting uuid_lslong  UuidLslong <<%T>> => %#v", mUuidLslongObj, mUuidLslongObj)
		mUuidLslong := mUuidLslongObj.(uint64)
		object.UuidLslong = mUuidLslong
	}

	log.Printf("FINISHED UuidType object: %#v", object)
}

func TakeUuidTypeAsMap(object *UuidType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["uuid_mslong"] = object.UuidMslong
	omap["uuid_lslong"] = object.UuidLslong

	return omap
}

func ResourceUuidTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid_mslong": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:unsignedLong
			Optional: true,
			Type:     schema.TypeInt,
		},
		"uuid_lslong": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:unsignedLong
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceUuidType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceUuidTypeSchema(),
	}
}

func SetIdPermsTypeFromMap(object *IdPermsType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetIdPermsTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mPermissionsObj := vmap["permissions"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mPermissionsObj) {
		log.Printf("Setting permissions  Permissions <<%T>> => %#v", mPermissionsObj, mPermissionsObj)
		mPermissions := new(PermType)
		SetPermTypeFromMap(mPermissions, d, m, mPermissionsObj)
		object.Permissions = mPermissions
	}

	mUuidObj := vmap["uuid"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mUuidObj) {
		log.Printf("Setting uuid  Uuid <<%T>> => %#v", mUuidObj, mUuidObj)
		mUuid := new(UuidType)
		SetUuidTypeFromMap(mUuid, d, m, mUuidObj)
		object.Uuid = mUuid
	}

	mEnableObj := vmap["enable"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEnableObj) {
		log.Printf("Setting enable  Enable <<%T>> => %#v", mEnableObj, mEnableObj)
		mEnable := mEnableObj.(bool)
		object.Enable = mEnable
	}

	mCreatedObj := vmap["created"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mCreatedObj) {
		log.Printf("Setting created  Created <<%T>> => %#v", mCreatedObj, mCreatedObj)
		mCreated := mCreatedObj.(string)
		object.Created = mCreated
	}

	mLastModifiedObj := vmap["last_modified"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLastModifiedObj) {
		log.Printf("Setting last_modified  LastModified <<%T>> => %#v", mLastModifiedObj, mLastModifiedObj)
		mLastModified := mLastModifiedObj.(string)
		object.LastModified = mLastModified
	}

	mDescriptionObj := vmap["description"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDescriptionObj) {
		log.Printf("Setting description  Description <<%T>> => %#v", mDescriptionObj, mDescriptionObj)
		mDescription := mDescriptionObj.(string)
		object.Description = mDescription
	}

	mUserVisibleObj := vmap["user_visible"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUserVisibleObj) {
		log.Printf("Setting user_visible  UserVisible <<%T>> => %#v", mUserVisibleObj, mUserVisibleObj)
		mUserVisible := mUserVisibleObj.(bool)
		object.UserVisible = mUserVisible
	}

	mCreatorObj := vmap["creator"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mCreatorObj) {
		log.Printf("Setting creator  Creator <<%T>> => %#v", mCreatorObj, mCreatorObj)
		mCreator := mCreatorObj.(string)
		object.Creator = mCreator
	}

	log.Printf("FINISHED IdPermsType object: %#v", object)
}

func TakeIdPermsTypeAsMap(object *IdPermsType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Permissions != nil {
		omap["permissions"] = TakePermTypeAsMap(object.Permissions)
	}
	if object.Uuid != nil {
		omap["uuid"] = TakeUuidTypeAsMap(object.Uuid)
	}
	omap["enable"] = object.Enable
	omap["created"] = object.Created
	omap["last_modified"] = object.LastModified
	omap["description"] = object.Description
	omap["user_visible"] = object.UserVisible
	omap["creator"] = object.Creator

	return omap
}

func ResourceIdPermsTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"permissions": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: PermType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePermType(),
		},
		"uuid": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: UuidType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceUuidType(),
		},
		"enable": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Required: true,
			Type:     schema.TypeBool,
		},
		"created": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:dateTime
			Optional: true,
			Type:     schema.TypeString,
		},
		"last_modified": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:dateTime
			Optional: true,
			Type:     schema.TypeString,
		},
		"description": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"user_visible": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"creator": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceIdPermsType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceIdPermsTypeSchema(),
	}
}

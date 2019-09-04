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

func SetInterfaceMirrorTypeFromMap(object *InterfaceMirrorType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetInterfaceMirrorTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mTrafficDirectionObj := vmap["traffic_direction"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mTrafficDirectionObj) {
		log.Printf("Setting traffic_direction  TrafficDirection <<%T>> => %#v", mTrafficDirectionObj, mTrafficDirectionObj)
		mTrafficDirection := mTrafficDirectionObj.(string)
		object.TrafficDirection = mTrafficDirection
	}

	mMirrorToObj := vmap["mirror_to"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mMirrorToObj) {
		log.Printf("Setting mirror_to  MirrorTo <<%T>> => %#v", mMirrorToObj, mMirrorToObj)
		mMirrorTo := new(MirrorActionType)
		SetMirrorActionTypeFromMap(mMirrorTo, d, m, mMirrorToObj)
		object.MirrorTo = mMirrorTo
	}

	log.Printf("FINISHED InterfaceMirrorType object: %#v", object)
}

func TakeInterfaceMirrorTypeAsMap(object *InterfaceMirrorType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["traffic_direction"] = object.TrafficDirection
	if object.MirrorTo != nil {
		omap["mirror_to"] = TakeMirrorActionTypeAsMap(object.MirrorTo)
	}

	return omap
}

func ResourceInterfaceMirrorTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"traffic_direction": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"mirror_to": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: MirrorActionType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceMirrorActionType(),
		},
	}
}

func ResourceInterfaceMirrorType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceInterfaceMirrorTypeSchema(),
	}
}

func SetVirtualMachineInterfacePropertiesTypeFromMap(object *VirtualMachineInterfacePropertiesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualMachineInterfacePropertiesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mServiceInterfaceTypeObj := vmap["service_interface_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceInterfaceTypeObj) {
		log.Printf("Setting service_interface_type  ServiceInterfaceType <<%T>> => %#v", mServiceInterfaceTypeObj, mServiceInterfaceTypeObj)
		mServiceInterfaceType := mServiceInterfaceTypeObj.(string)
		object.ServiceInterfaceType = mServiceInterfaceType
	}

	mInterfaceMirrorObj := vmap["interface_mirror"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mInterfaceMirrorObj) {
		log.Printf("Setting interface_mirror  InterfaceMirror <<%T>> => %#v", mInterfaceMirrorObj, mInterfaceMirrorObj)
		mInterfaceMirror := new(InterfaceMirrorType)
		SetInterfaceMirrorTypeFromMap(mInterfaceMirror, d, m, mInterfaceMirrorObj)
		object.InterfaceMirror = mInterfaceMirror
	}

	mLocalPreferenceObj := vmap["local_preference"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLocalPreferenceObj) {
		log.Printf("Setting local_preference  LocalPreference <<%T>> => %#v", mLocalPreferenceObj, mLocalPreferenceObj)
		mLocalPreference := mLocalPreferenceObj.(int)
		object.LocalPreference = mLocalPreference
	}

	mSubInterfaceVlanTagObj := vmap["sub_interface_vlan_tag"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSubInterfaceVlanTagObj) {
		log.Printf("Setting sub_interface_vlan_tag  SubInterfaceVlanTag <<%T>> => %#v", mSubInterfaceVlanTagObj, mSubInterfaceVlanTagObj)
		mSubInterfaceVlanTag := mSubInterfaceVlanTagObj.(int)
		object.SubInterfaceVlanTag = mSubInterfaceVlanTag
	}

	log.Printf("FINISHED VirtualMachineInterfacePropertiesType object: %#v", object)
}

func TakeVirtualMachineInterfacePropertiesTypeAsMap(object *VirtualMachineInterfacePropertiesType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["service_interface_type"] = object.ServiceInterfaceType
	if object.InterfaceMirror != nil {
		omap["interface_mirror"] = TakeInterfaceMirrorTypeAsMap(object.InterfaceMirror)
	}
	omap["local_preference"] = object.LocalPreference
	omap["sub_interface_vlan_tag"] = object.SubInterfaceVlanTag

	return omap
}

func ResourceVirtualMachineInterfacePropertiesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_interface_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"interface_mirror": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: InterfaceMirrorType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceInterfaceMirrorType(),
		},
		"local_preference": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"sub_interface_vlan_tag": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceVirtualMachineInterfacePropertiesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualMachineInterfacePropertiesTypeSchema(),
	}
}

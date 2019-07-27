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

func SetServiceInstanceInterfaceTypeFromMap(object *ServiceInstanceInterfaceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceInstanceInterfaceTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mVirtualNetworkObj := vmap["virtual_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualNetworkObj) {
		log.Printf("Setting virtual_network  VirtualNetwork <<%T>> => %#v", mVirtualNetworkObj, mVirtualNetworkObj)
		mVirtualNetwork := mVirtualNetworkObj.(string)
		object.VirtualNetwork = mVirtualNetwork
	}

	mIpAddressObj := vmap["ip_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpAddressObj) {
		log.Printf("Setting ip_address  IpAddress <<%T>> => %#v", mIpAddressObj, mIpAddressObj)
		mIpAddress := mIpAddressObj.(string)
		object.IpAddress = mIpAddress
	}

	mStaticRoutesObj := vmap["static_routes"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mStaticRoutesObj) {
		log.Printf("Setting static_routes  StaticRoutes <<%T>> => %#v", mStaticRoutesObj, mStaticRoutesObj)
		mStaticRoutes := new(RouteTableType)
		SetRouteTableTypeFromMap(mStaticRoutes, d, m, mStaticRoutesObj)
		object.StaticRoutes = mStaticRoutes
	}

	mAllowedAddressPairsObj := vmap["allowed_address_pairs"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mAllowedAddressPairsObj) {
		log.Printf("Setting allowed_address_pairs  AllowedAddressPairs <<%T>> => %#v", mAllowedAddressPairsObj, mAllowedAddressPairsObj)
		mAllowedAddressPairs := new(AllowedAddressPairs)
		SetAllowedAddressPairsFromMap(mAllowedAddressPairs, d, m, mAllowedAddressPairsObj)
		object.AllowedAddressPairs = mAllowedAddressPairs
	}

	log.Printf("FINISHED ServiceInstanceInterfaceType object: %#v", object)
}

func TakeServiceInstanceInterfaceTypeAsMap(object *ServiceInstanceInterfaceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["virtual_network"] = object.VirtualNetwork
	omap["ip_address"] = object.IpAddress
	if object.StaticRoutes != nil {
		omap["static_routes"] = TakeRouteTableTypeAsMap(object.StaticRoutes)
	}
	if object.AllowedAddressPairs != nil {
		omap["allowed_address_pairs"] = TakeAllowedAddressPairsAsMap(object.AllowedAddressPairs)
	}

	return omap
}

func ResourceServiceInstanceInterfaceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"virtual_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"ip_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"static_routes": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: RouteTableType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTableType(),
		},
		"allowed_address_pairs": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: AllowedAddressPairs
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAllowedAddressPairs(),
		},
	}
}

func ResourceServiceInstanceInterfaceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceInstanceInterfaceTypeSchema(),
	}
}

func SetServiceScaleOutTypeFromMap(object *ServiceScaleOutType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceScaleOutTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mMaxInstancesObj := vmap["max_instances"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mMaxInstancesObj) {
		log.Printf("Setting max_instances  MaxInstances <<%T>> => %#v", mMaxInstancesObj, mMaxInstancesObj)
		mMaxInstances := mMaxInstancesObj.(int)
		object.MaxInstances = mMaxInstances
	}

	mAutoScaleObj := vmap["auto_scale"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAutoScaleObj) {
		log.Printf("Setting auto_scale  AutoScale <<%T>> => %#v", mAutoScaleObj, mAutoScaleObj)
		mAutoScale := mAutoScaleObj.(bool)
		object.AutoScale = mAutoScale
	}

	log.Printf("FINISHED ServiceScaleOutType object: %#v", object)
}

func TakeServiceScaleOutTypeAsMap(object *ServiceScaleOutType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["max_instances"] = object.MaxInstances
	omap["auto_scale"] = object.AutoScale

	return omap
}

func ResourceServiceScaleOutTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"max_instances": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"auto_scale": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Required: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceServiceScaleOutType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceScaleOutTypeSchema(),
	}
}

func SetServiceInstanceTypeFromMap(object *ServiceInstanceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceInstanceTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mAutoPolicyObj := vmap["auto_policy"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAutoPolicyObj) {
		log.Printf("Setting auto_policy  AutoPolicy <<%T>> => %#v", mAutoPolicyObj, mAutoPolicyObj)
		mAutoPolicy := mAutoPolicyObj.(bool)
		object.AutoPolicy = mAutoPolicy
	}

	mAvailabilityZoneObj := vmap["availability_zone"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAvailabilityZoneObj) {
		log.Printf("Setting availability_zone  AvailabilityZone <<%T>> => %#v", mAvailabilityZoneObj, mAvailabilityZoneObj)
		mAvailabilityZone := mAvailabilityZoneObj.(string)
		object.AvailabilityZone = mAvailabilityZone
	}

	mManagementVirtualNetworkObj := vmap["management_virtual_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mManagementVirtualNetworkObj) {
		log.Printf("Setting management_virtual_network  ManagementVirtualNetwork <<%T>> => %#v", mManagementVirtualNetworkObj, mManagementVirtualNetworkObj)
		mManagementVirtualNetwork := mManagementVirtualNetworkObj.(string)
		object.ManagementVirtualNetwork = mManagementVirtualNetwork
	}

	mLeftVirtualNetworkObj := vmap["left_virtual_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLeftVirtualNetworkObj) {
		log.Printf("Setting left_virtual_network  LeftVirtualNetwork <<%T>> => %#v", mLeftVirtualNetworkObj, mLeftVirtualNetworkObj)
		mLeftVirtualNetwork := mLeftVirtualNetworkObj.(string)
		object.LeftVirtualNetwork = mLeftVirtualNetwork
	}

	mLeftIpAddressObj := vmap["left_ip_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLeftIpAddressObj) {
		log.Printf("Setting left_ip_address  LeftIpAddress <<%T>> => %#v", mLeftIpAddressObj, mLeftIpAddressObj)
		mLeftIpAddress := mLeftIpAddressObj.(string)
		object.LeftIpAddress = mLeftIpAddress
	}

	mRightVirtualNetworkObj := vmap["right_virtual_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRightVirtualNetworkObj) {
		log.Printf("Setting right_virtual_network  RightVirtualNetwork <<%T>> => %#v", mRightVirtualNetworkObj, mRightVirtualNetworkObj)
		mRightVirtualNetwork := mRightVirtualNetworkObj.(string)
		object.RightVirtualNetwork = mRightVirtualNetwork
	}

	mRightIpAddressObj := vmap["right_ip_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRightIpAddressObj) {
		log.Printf("Setting right_ip_address  RightIpAddress <<%T>> => %#v", mRightIpAddressObj, mRightIpAddressObj)
		mRightIpAddress := mRightIpAddressObj.(string)
		object.RightIpAddress = mRightIpAddress
	}

	mInterfaceListObj := vmap["interface_list"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mInterfaceListObj) {
		log.Printf("Setting interface_list  InterfaceList <<%T>> => %#v", mInterfaceListObj, mInterfaceListObj)
		for _, v := range mInterfaceListObj.([]interface{}) {
			mInterfaceList := new(ServiceInstanceInterfaceType)
			SetServiceInstanceInterfaceTypeFromMap(mInterfaceList, d, m, v)
			object.AddInterfaceList(mInterfaceList)
		}
	}

	mScaleOutObj := vmap["scale_out"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mScaleOutObj) {
		log.Printf("Setting scale_out  ScaleOut <<%T>> => %#v", mScaleOutObj, mScaleOutObj)
		mScaleOut := new(ServiceScaleOutType)
		SetServiceScaleOutTypeFromMap(mScaleOut, d, m, mScaleOutObj)
		object.ScaleOut = mScaleOut
	}

	mHaModeObj := vmap["ha_mode"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mHaModeObj) {
		log.Printf("Setting ha_mode  HaMode <<%T>> => %#v", mHaModeObj, mHaModeObj)
		mHaMode := mHaModeObj.(string)
		object.HaMode = mHaMode
	}

	mVirtualRouterIdObj := vmap["virtual_router_id"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualRouterIdObj) {
		log.Printf("Setting virtual_router_id  VirtualRouterId <<%T>> => %#v", mVirtualRouterIdObj, mVirtualRouterIdObj)
		mVirtualRouterId := mVirtualRouterIdObj.(string)
		object.VirtualRouterId = mVirtualRouterId
	}

	log.Printf("FINISHED ServiceInstanceType object: %#v", object)
}

func TakeServiceInstanceTypeAsMap(object *ServiceInstanceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["auto_policy"] = object.AutoPolicy
	omap["availability_zone"] = object.AvailabilityZone
	omap["management_virtual_network"] = object.ManagementVirtualNetwork
	omap["left_virtual_network"] = object.LeftVirtualNetwork
	omap["left_ip_address"] = object.LeftIpAddress
	omap["right_virtual_network"] = object.RightVirtualNetwork
	omap["right_ip_address"] = object.RightIpAddress
	var interface_list_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.InterfaceList {
		interface_list_map = append(interface_list_map, TakeServiceInstanceInterfaceTypeAsMap(&v))
	}
	omap["interface_list"] = interface_list_map
	if object.ScaleOut != nil {
		omap["scale_out"] = TakeServiceScaleOutTypeAsMap(object.ScaleOut)
	}
	omap["ha_mode"] = object.HaMode
	omap["virtual_router_id"] = object.VirtualRouterId

	return omap
}

func ResourceServiceInstanceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_policy": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"availability_zone": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"management_virtual_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"left_virtual_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"left_ip_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"right_virtual_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"right_ip_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"interface_list": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: ServiceInstanceInterfaceType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceInstanceInterfaceType(),
		},
		"scale_out": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: ServiceScaleOutType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceScaleOutType(),
		},
		"ha_mode": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"virtual_router_id": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceServiceInstanceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceInstanceTypeSchema(),
	}
}

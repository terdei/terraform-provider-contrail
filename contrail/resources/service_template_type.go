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

func SetServiceTemplateInterfaceTypeFromMap(object *ServiceTemplateInterfaceType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceTemplateInterfaceTypeFromMAP] map = %#v", val)

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

	mSharedIpObj := vmap["shared_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSharedIpObj) {
		log.Printf("Setting shared_ip  SharedIp <<%T>> => %#v", mSharedIpObj, mSharedIpObj)
		mSharedIp := mSharedIpObj.(bool)
		object.SharedIp = mSharedIp
	}

	mStaticRouteEnableObj := vmap["static_route_enable"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mStaticRouteEnableObj) {
		log.Printf("Setting static_route_enable  StaticRouteEnable <<%T>> => %#v", mStaticRouteEnableObj, mStaticRouteEnableObj)
		mStaticRouteEnable := mStaticRouteEnableObj.(bool)
		object.StaticRouteEnable = mStaticRouteEnable
	}

	log.Printf("FINISHED ServiceTemplateInterfaceType object: %#v", object)
}

func TakeServiceTemplateInterfaceTypeAsMap(object *ServiceTemplateInterfaceType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["service_interface_type"] = object.ServiceInterfaceType
	omap["shared_ip"] = object.SharedIp
	omap["static_route_enable"] = object.StaticRouteEnable

	return omap
}

func ResourceServiceTemplateInterfaceTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_interface_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"shared_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"static_route_enable": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
	}
}

func ResourceServiceTemplateInterfaceType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceTemplateInterfaceTypeSchema(),
	}
}

func SetServiceTemplateTypeFromMap(object *ServiceTemplateType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetServiceTemplateTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mVersionObj := vmap["version"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVersionObj) {
		log.Printf("Setting version  Version <<%T>> => %#v", mVersionObj, mVersionObj)
		mVersion := mVersionObj.(int)
		object.Version = mVersion
	}

	mServiceModeObj := vmap["service_mode"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceModeObj) {
		log.Printf("Setting service_mode  ServiceMode <<%T>> => %#v", mServiceModeObj, mServiceModeObj)
		mServiceMode := mServiceModeObj.(string)
		object.ServiceMode = mServiceMode
	}

	mServiceTypeObj := vmap["service_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceTypeObj) {
		log.Printf("Setting service_type  ServiceType <<%T>> => %#v", mServiceTypeObj, mServiceTypeObj)
		mServiceType := mServiceTypeObj.(string)
		object.ServiceType = mServiceType
	}

	mImageNameObj := vmap["image_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mImageNameObj) {
		log.Printf("Setting image_name  ImageName <<%T>> => %#v", mImageNameObj, mImageNameObj)
		mImageName := mImageNameObj.(string)
		object.ImageName = mImageName
	}

	mServiceScalingObj := vmap["service_scaling"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceScalingObj) {
		log.Printf("Setting service_scaling  ServiceScaling <<%T>> => %#v", mServiceScalingObj, mServiceScalingObj)
		mServiceScaling := mServiceScalingObj.(bool)
		object.ServiceScaling = mServiceScaling
	}

	mInterfaceTypeObj := vmap["interface_type"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mInterfaceTypeObj) {
		log.Printf("Setting interface_type  InterfaceType <<%T>> => %#v", mInterfaceTypeObj, mInterfaceTypeObj)
		for _, v := range mInterfaceTypeObj.([]interface{}) {
			mInterfaceType := new(ServiceTemplateInterfaceType)
			SetServiceTemplateInterfaceTypeFromMap(mInterfaceType, d, m, v)
			object.AddInterfaceType(mInterfaceType)
		}
	}

	mFlavorObj := vmap["flavor"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mFlavorObj) {
		log.Printf("Setting flavor  Flavor <<%T>> => %#v", mFlavorObj, mFlavorObj)
		mFlavor := mFlavorObj.(string)
		object.Flavor = mFlavor
	}

	mOrderedInterfacesObj := vmap["ordered_interfaces"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mOrderedInterfacesObj) {
		log.Printf("Setting ordered_interfaces  OrderedInterfaces <<%T>> => %#v", mOrderedInterfacesObj, mOrderedInterfacesObj)
		mOrderedInterfaces := mOrderedInterfacesObj.(bool)
		object.OrderedInterfaces = mOrderedInterfaces
	}

	mServiceVirtualizationTypeObj := vmap["service_virtualization_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceVirtualizationTypeObj) {
		log.Printf("Setting service_virtualization_type  ServiceVirtualizationType <<%T>> => %#v", mServiceVirtualizationTypeObj, mServiceVirtualizationTypeObj)
		mServiceVirtualizationType := mServiceVirtualizationTypeObj.(string)
		object.ServiceVirtualizationType = mServiceVirtualizationType
	}

	mAvailabilityZoneEnableObj := vmap["availability_zone_enable"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAvailabilityZoneEnableObj) {
		log.Printf("Setting availability_zone_enable  AvailabilityZoneEnable <<%T>> => %#v", mAvailabilityZoneEnableObj, mAvailabilityZoneEnableObj)
		mAvailabilityZoneEnable := mAvailabilityZoneEnableObj.(bool)
		object.AvailabilityZoneEnable = mAvailabilityZoneEnable
	}

	mVrouterInstanceTypeObj := vmap["vrouter_instance_type"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVrouterInstanceTypeObj) {
		log.Printf("Setting vrouter_instance_type  VrouterInstanceType <<%T>> => %#v", mVrouterInstanceTypeObj, mVrouterInstanceTypeObj)
		mVrouterInstanceType := mVrouterInstanceTypeObj.(string)
		object.VrouterInstanceType = mVrouterInstanceType
	}

	mInstanceDataObj := vmap["instance_data"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mInstanceDataObj) {
		log.Printf("Setting instance_data  InstanceData <<%T>> => %#v", mInstanceDataObj, mInstanceDataObj)
		mInstanceData := mInstanceDataObj.(string)
		object.InstanceData = mInstanceData
	}

	log.Printf("FINISHED ServiceTemplateType object: %#v", object)
}

func TakeServiceTemplateTypeAsMap(object *ServiceTemplateType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["version"] = object.Version
	omap["service_mode"] = object.ServiceMode
	omap["service_type"] = object.ServiceType
	omap["image_name"] = object.ImageName
	omap["service_scaling"] = object.ServiceScaling
	var interface_type_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.InterfaceType {
		interface_type_map = append(interface_type_map, TakeServiceTemplateInterfaceTypeAsMap(&v))
	}
	omap["interface_type"] = interface_type_map
	omap["flavor"] = object.Flavor
	omap["ordered_interfaces"] = object.OrderedInterfaces
	omap["service_virtualization_type"] = object.ServiceVirtualizationType
	omap["availability_zone_enable"] = object.AvailabilityZoneEnable
	omap["vrouter_instance_type"] = object.VrouterInstanceType
	omap["instance_data"] = object.InstanceData

	return omap
}

func ResourceServiceTemplateTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"service_mode": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"service_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"image_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"service_scaling": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"interface_type": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: ServiceTemplateInterfaceType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceTemplateInterfaceType(),
		},
		"flavor": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ordered_interfaces": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"service_virtualization_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"availability_zone_enable": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"vrouter_instance_type": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"instance_data": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceServiceTemplateType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceServiceTemplateTypeSchema(),
	}
}

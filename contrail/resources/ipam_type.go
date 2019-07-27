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

func SetIpAddressesTypeFromMap(object *IpAddressesType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetIpAddressesTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mIpAddressObj := vmap["ip_address"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mIpAddressObj) {
		log.Printf("Setting ip_address  IpAddress <<%T>> => %#v", mIpAddressObj, mIpAddressObj)
		for _, v := range mIpAddressObj.([]interface{}) {
			mIpAddress := v.(string)
			object.AddIpAddress(mIpAddress)
		}
	}

	log.Printf("FINISHED IpAddressesType object: %#v", object)
}

func TakeIpAddressesTypeAsMap(object *IpAddressesType) map[string]interface{} {
	omap := make(map[string]interface{})

	ip_address_arr := make([]string, len(object.IpAddress))
	for _, v := range object.IpAddress {
		ip_address_arr = append(ip_address_arr, v)
	}
	omap["ip_address"] = ip_address_arr

	return omap
}

func ResourceIpAddressesTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip_address": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceIpAddressesType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceIpAddressesTypeSchema(),
	}
}

func SetIpamDnsAddressTypeFromMap(object *IpamDnsAddressType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetIpamDnsAddressTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mTenantDnsServerAddressObj := vmap["tenant_dns_server_address"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mTenantDnsServerAddressObj) {
		log.Printf("Setting tenant_dns_server_address  TenantDnsServerAddress <<%T>> => %#v", mTenantDnsServerAddressObj, mTenantDnsServerAddressObj)
		mTenantDnsServerAddress := new(IpAddressesType)
		SetIpAddressesTypeFromMap(mTenantDnsServerAddress, d, m, mTenantDnsServerAddressObj)
		object.TenantDnsServerAddress = mTenantDnsServerAddress
	}

	mVirtualDnsServerNameObj := vmap["virtual_dns_server_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualDnsServerNameObj) {
		log.Printf("Setting virtual_dns_server_name  VirtualDnsServerName <<%T>> => %#v", mVirtualDnsServerNameObj, mVirtualDnsServerNameObj)
		mVirtualDnsServerName := mVirtualDnsServerNameObj.(string)
		object.VirtualDnsServerName = mVirtualDnsServerName
	}

	log.Printf("FINISHED IpamDnsAddressType object: %#v", object)
}

func TakeIpamDnsAddressTypeAsMap(object *IpamDnsAddressType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.TenantDnsServerAddress != nil {
		omap["tenant_dns_server_address"] = TakeIpAddressesTypeAsMap(object.TenantDnsServerAddress)
	}
	omap["virtual_dns_server_name"] = object.VirtualDnsServerName

	return omap
}

func ResourceIpamDnsAddressTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tenant_dns_server_address": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: IpAddressesType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIpAddressesType(),
		},
		"virtual_dns_server_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceIpamDnsAddressType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceIpamDnsAddressTypeSchema(),
	}
}

func SetIpamTypeFromMap(object *IpamType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetIpamTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mIpamMethodObj := vmap["ipam_method"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpamMethodObj) {
		log.Printf("Setting ipam_method  IpamMethod <<%T>> => %#v", mIpamMethodObj, mIpamMethodObj)
		mIpamMethod := mIpamMethodObj.(string)
		object.IpamMethod = mIpamMethod
	}

	mIpamDnsMethodObj := vmap["ipam_dns_method"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpamDnsMethodObj) {
		log.Printf("Setting ipam_dns_method  IpamDnsMethod <<%T>> => %#v", mIpamDnsMethodObj, mIpamDnsMethodObj)
		mIpamDnsMethod := mIpamDnsMethodObj.(string)
		object.IpamDnsMethod = mIpamDnsMethod
	}

	mIpamDnsServerObj := vmap["ipam_dns_server"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mIpamDnsServerObj) {
		log.Printf("Setting ipam_dns_server  IpamDnsServer <<%T>> => %#v", mIpamDnsServerObj, mIpamDnsServerObj)
		mIpamDnsServer := new(IpamDnsAddressType)
		SetIpamDnsAddressTypeFromMap(mIpamDnsServer, d, m, mIpamDnsServerObj)
		object.IpamDnsServer = mIpamDnsServer
	}

	mDhcpOptionListObj := vmap["dhcp_option_list"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mDhcpOptionListObj) {
		log.Printf("Setting dhcp_option_list  DhcpOptionList <<%T>> => %#v", mDhcpOptionListObj, mDhcpOptionListObj)
		mDhcpOptionList := new(DhcpOptionsListType)
		SetDhcpOptionsListTypeFromMap(mDhcpOptionList, d, m, mDhcpOptionListObj)
		object.DhcpOptionList = mDhcpOptionList
	}

	mCidrBlockObj := vmap["cidr_block"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mCidrBlockObj) {
		log.Printf("Setting cidr_block  CidrBlock <<%T>> => %#v", mCidrBlockObj, mCidrBlockObj)
		mCidrBlock := new(SubnetType)
		SetSubnetTypeFromMap(mCidrBlock, d, m, mCidrBlockObj)
		object.CidrBlock = mCidrBlock
	}

	mHostRoutesObj := vmap["host_routes"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mHostRoutesObj) {
		log.Printf("Setting host_routes  HostRoutes <<%T>> => %#v", mHostRoutesObj, mHostRoutesObj)
		mHostRoutes := new(RouteTableType)
		SetRouteTableTypeFromMap(mHostRoutes, d, m, mHostRoutesObj)
		object.HostRoutes = mHostRoutes
	}

	log.Printf("FINISHED IpamType object: %#v", object)
}

func TakeIpamTypeAsMap(object *IpamType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["ipam_method"] = object.IpamMethod
	omap["ipam_dns_method"] = object.IpamDnsMethod
	if object.IpamDnsServer != nil {
		omap["ipam_dns_server"] = TakeIpamDnsAddressTypeAsMap(object.IpamDnsServer)
	}
	if object.DhcpOptionList != nil {
		omap["dhcp_option_list"] = TakeDhcpOptionsListTypeAsMap(object.DhcpOptionList)
	}
	if object.CidrBlock != nil {
		omap["cidr_block"] = TakeSubnetTypeAsMap(object.CidrBlock)
	}
	if object.HostRoutes != nil {
		omap["host_routes"] = TakeRouteTableTypeAsMap(object.HostRoutes)
	}

	return omap
}

func ResourceIpamTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipam_method": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ipam_dns_method": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ipam_dns_server": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: IpamDnsAddressType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIpamDnsAddressType(),
		},
		"dhcp_option_list": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: DhcpOptionsListType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDhcpOptionsListType(),
		},
		"cidr_block": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SubnetType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
		"host_routes": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: RouteTableType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTableType(),
		},
	}
}

func ResourceIpamType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceIpamTypeSchema(),
	}
}

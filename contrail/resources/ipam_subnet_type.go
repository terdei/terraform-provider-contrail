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

func SetIpamSubnetTypeFromMap(object *IpamSubnetType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetIpamSubnetTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mSubnetObj := vmap["subnet"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSubnetObj) {
		log.Printf("Setting subnet  Subnet <<%T>> => %#v", mSubnetObj, mSubnetObj)
		mSubnet := new(SubnetType)
		SetSubnetTypeFromMap(mSubnet, d, m, mSubnetObj)
		object.Subnet = mSubnet
	}

	mDefaultGatewayObj := vmap["default_gateway"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDefaultGatewayObj) {
		log.Printf("Setting default_gateway  DefaultGateway <<%T>> => %#v", mDefaultGatewayObj, mDefaultGatewayObj)
		mDefaultGateway := mDefaultGatewayObj.(string)
		object.DefaultGateway = mDefaultGateway
	}

	mDnsServerAddressObj := vmap["dns_server_address"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDnsServerAddressObj) {
		log.Printf("Setting dns_server_address  DnsServerAddress <<%T>> => %#v", mDnsServerAddressObj, mDnsServerAddressObj)
		mDnsServerAddress := mDnsServerAddressObj.(string)
		object.DnsServerAddress = mDnsServerAddress
	}

	mSubnetUuidObj := vmap["subnet_uuid"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSubnetUuidObj) {
		log.Printf("Setting subnet_uuid  SubnetUuid <<%T>> => %#v", mSubnetUuidObj, mSubnetUuidObj)
		mSubnetUuid := mSubnetUuidObj.(string)
		object.SubnetUuid = mSubnetUuid
	}

	mEnableDhcpObj := vmap["enable_dhcp"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mEnableDhcpObj) {
		log.Printf("Setting enable_dhcp  EnableDhcp <<%T>> => %#v", mEnableDhcpObj, mEnableDhcpObj)
		mEnableDhcp := mEnableDhcpObj.(bool)
		object.EnableDhcp = mEnableDhcp
	}

	mDnsNameserversObj := vmap["dns_nameservers"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mDnsNameserversObj) {
		log.Printf("Setting dns_nameservers  DnsNameservers <<%T>> => %#v", mDnsNameserversObj, mDnsNameserversObj)
		for _, v := range mDnsNameserversObj.([]interface{}) {
			mDnsNameservers := v.(string)
			object.AddDnsNameservers(mDnsNameservers)
		}
	}

	mAllocationPoolsObj := vmap["allocation_pools"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mAllocationPoolsObj) {
		log.Printf("Setting allocation_pools  AllocationPools <<%T>> => %#v", mAllocationPoolsObj, mAllocationPoolsObj)
		for _, v := range mAllocationPoolsObj.([]interface{}) {
			mAllocationPools := new(AllocationPoolType)
			SetAllocationPoolTypeFromMap(mAllocationPools, d, m, v)
			object.AddAllocationPools(mAllocationPools)
		}
	}

	mAddrFromStartObj := vmap["addr_from_start"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAddrFromStartObj) {
		log.Printf("Setting addr_from_start  AddrFromStart <<%T>> => %#v", mAddrFromStartObj, mAddrFromStartObj)
		mAddrFromStart := mAddrFromStartObj.(bool)
		object.AddrFromStart = mAddrFromStart
	}

	mDhcpOptionListObj := vmap["dhcp_option_list"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mDhcpOptionListObj) {
		log.Printf("Setting dhcp_option_list  DhcpOptionList <<%T>> => %#v", mDhcpOptionListObj, mDhcpOptionListObj)
		mDhcpOptionList := new(DhcpOptionsListType)
		SetDhcpOptionsListTypeFromMap(mDhcpOptionList, d, m, mDhcpOptionListObj)
		object.DhcpOptionList = mDhcpOptionList
	}

	mHostRoutesObj := vmap["host_routes"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mHostRoutesObj) {
		log.Printf("Setting host_routes  HostRoutes <<%T>> => %#v", mHostRoutesObj, mHostRoutesObj)
		mHostRoutes := new(RouteTableType)
		SetRouteTableTypeFromMap(mHostRoutes, d, m, mHostRoutesObj)
		object.HostRoutes = mHostRoutes
	}

	mSubnetNameObj := vmap["subnet_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSubnetNameObj) {
		log.Printf("Setting subnet_name  SubnetName <<%T>> => %#v", mSubnetNameObj, mSubnetNameObj)
		mSubnetName := mSubnetNameObj.(string)
		object.SubnetName = mSubnetName
	}

	mAllocUnitObj := vmap["alloc_unit"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAllocUnitObj) {
		log.Printf("Setting alloc_unit  AllocUnit <<%T>> => %#v", mAllocUnitObj, mAllocUnitObj)
		mAllocUnit := mAllocUnitObj.(int)
		object.AllocUnit = mAllocUnit
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

	log.Printf("FINISHED IpamSubnetType object: %#v", object)
}

func TakeIpamSubnetTypeAsMap(object *IpamSubnetType) map[string]interface{} {
	omap := make(map[string]interface{})

	if object.Subnet != nil {
		omap["subnet"] = TakeSubnetTypeAsMap(object.Subnet)
	}
	omap["default_gateway"] = object.DefaultGateway
	omap["dns_server_address"] = object.DnsServerAddress
	omap["subnet_uuid"] = object.SubnetUuid
	omap["enable_dhcp"] = object.EnableDhcp
	dns_nameservers_arr := make([]string, len(object.DnsNameservers))
	for _, v := range object.DnsNameservers {
		dns_nameservers_arr = append(dns_nameservers_arr, v)
	}
	omap["dns_nameservers"] = dns_nameservers_arr
	var allocation_pools_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.AllocationPools {
		allocation_pools_map = append(allocation_pools_map, TakeAllocationPoolTypeAsMap(&v))
	}
	omap["allocation_pools"] = allocation_pools_map
	omap["addr_from_start"] = object.AddrFromStart
	if object.DhcpOptionList != nil {
		omap["dhcp_option_list"] = TakeDhcpOptionsListTypeAsMap(object.DhcpOptionList)
	}
	if object.HostRoutes != nil {
		omap["host_routes"] = TakeRouteTableTypeAsMap(object.HostRoutes)
	}
	omap["subnet_name"] = object.SubnetName
	omap["alloc_unit"] = object.AllocUnit
	omap["created"] = object.Created
	omap["last_modified"] = object.LastModified

	return omap
}

func ResourceIpamSubnetTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: SubnetType
			Required: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
		"default_gateway": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"dns_server_address": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"subnet_uuid": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"enable_dhcp": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"dns_nameservers": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"allocation_pools": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: AllocationPoolType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAllocationPoolType(),
		},
		"addr_from_start": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"dhcp_option_list": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: DhcpOptionsListType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDhcpOptionsListType(),
		},
		"host_routes": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: RouteTableType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTableType(),
		},
		"subnet_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"alloc_unit": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
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
	}
}

func ResourceIpamSubnetType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceIpamSubnetTypeSchema(),
	}
}

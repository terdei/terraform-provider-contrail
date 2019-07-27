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

func SetQuotaTypeFromMap(object *QuotaType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetQuotaTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mDefaultsObj := vmap["defaults"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDefaultsObj) {
		log.Printf("Setting defaults  Defaults <<%T>> => %#v", mDefaultsObj, mDefaultsObj)
		mDefaults := mDefaultsObj.(int)
		object.Defaults = mDefaults
	}

	mFloatingIpObj := vmap["floating_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mFloatingIpObj) {
		log.Printf("Setting floating_ip  FloatingIp <<%T>> => %#v", mFloatingIpObj, mFloatingIpObj)
		mFloatingIp := mFloatingIpObj.(int)
		object.FloatingIp = mFloatingIp
	}

	mInstanceIpObj := vmap["instance_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mInstanceIpObj) {
		log.Printf("Setting instance_ip  InstanceIp <<%T>> => %#v", mInstanceIpObj, mInstanceIpObj)
		mInstanceIp := mInstanceIpObj.(int)
		object.InstanceIp = mInstanceIp
	}

	mVirtualMachineInterfaceObj := vmap["virtual_machine_interface"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualMachineInterfaceObj) {
		log.Printf("Setting virtual_machine_interface  VirtualMachineInterface <<%T>> => %#v", mVirtualMachineInterfaceObj, mVirtualMachineInterfaceObj)
		mVirtualMachineInterface := mVirtualMachineInterfaceObj.(int)
		object.VirtualMachineInterface = mVirtualMachineInterface
	}

	mVirtualNetworkObj := vmap["virtual_network"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualNetworkObj) {
		log.Printf("Setting virtual_network  VirtualNetwork <<%T>> => %#v", mVirtualNetworkObj, mVirtualNetworkObj)
		mVirtualNetwork := mVirtualNetworkObj.(int)
		object.VirtualNetwork = mVirtualNetwork
	}

	mVirtualRouterObj := vmap["virtual_router"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualRouterObj) {
		log.Printf("Setting virtual_router  VirtualRouter <<%T>> => %#v", mVirtualRouterObj, mVirtualRouterObj)
		mVirtualRouter := mVirtualRouterObj.(int)
		object.VirtualRouter = mVirtualRouter
	}

	mVirtualDnsObj := vmap["virtual_DNS"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualDnsObj) {
		log.Printf("Setting virtual_DNS  VirtualDns <<%T>> => %#v", mVirtualDnsObj, mVirtualDnsObj)
		mVirtualDns := mVirtualDnsObj.(int)
		object.VirtualDns = mVirtualDns
	}

	mVirtualDnsRecordObj := vmap["virtual_DNS_record"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualDnsRecordObj) {
		log.Printf("Setting virtual_DNS_record  VirtualDnsRecord <<%T>> => %#v", mVirtualDnsRecordObj, mVirtualDnsRecordObj)
		mVirtualDnsRecord := mVirtualDnsRecordObj.(int)
		object.VirtualDnsRecord = mVirtualDnsRecord
	}

	mBgpRouterObj := vmap["bgp_router"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mBgpRouterObj) {
		log.Printf("Setting bgp_router  BgpRouter <<%T>> => %#v", mBgpRouterObj, mBgpRouterObj)
		mBgpRouter := mBgpRouterObj.(int)
		object.BgpRouter = mBgpRouter
	}

	mNetworkIpamObj := vmap["network_ipam"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNetworkIpamObj) {
		log.Printf("Setting network_ipam  NetworkIpam <<%T>> => %#v", mNetworkIpamObj, mNetworkIpamObj)
		mNetworkIpam := mNetworkIpamObj.(int)
		object.NetworkIpam = mNetworkIpam
	}

	mAccessControlListObj := vmap["access_control_list"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mAccessControlListObj) {
		log.Printf("Setting access_control_list  AccessControlList <<%T>> => %#v", mAccessControlListObj, mAccessControlListObj)
		mAccessControlList := mAccessControlListObj.(int)
		object.AccessControlList = mAccessControlList
	}

	mNetworkPolicyObj := vmap["network_policy"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNetworkPolicyObj) {
		log.Printf("Setting network_policy  NetworkPolicy <<%T>> => %#v", mNetworkPolicyObj, mNetworkPolicyObj)
		mNetworkPolicy := mNetworkPolicyObj.(int)
		object.NetworkPolicy = mNetworkPolicy
	}

	mFloatingIpPoolObj := vmap["floating_ip_pool"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mFloatingIpPoolObj) {
		log.Printf("Setting floating_ip_pool  FloatingIpPool <<%T>> => %#v", mFloatingIpPoolObj, mFloatingIpPoolObj)
		mFloatingIpPool := mFloatingIpPoolObj.(int)
		object.FloatingIpPool = mFloatingIpPool
	}

	mServiceTemplateObj := vmap["service_template"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceTemplateObj) {
		log.Printf("Setting service_template  ServiceTemplate <<%T>> => %#v", mServiceTemplateObj, mServiceTemplateObj)
		mServiceTemplate := mServiceTemplateObj.(int)
		object.ServiceTemplate = mServiceTemplate
	}

	mServiceInstanceObj := vmap["service_instance"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServiceInstanceObj) {
		log.Printf("Setting service_instance  ServiceInstance <<%T>> => %#v", mServiceInstanceObj, mServiceInstanceObj)
		mServiceInstance := mServiceInstanceObj.(int)
		object.ServiceInstance = mServiceInstance
	}

	mLogicalRouterObj := vmap["logical_router"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLogicalRouterObj) {
		log.Printf("Setting logical_router  LogicalRouter <<%T>> => %#v", mLogicalRouterObj, mLogicalRouterObj)
		mLogicalRouter := mLogicalRouterObj.(int)
		object.LogicalRouter = mLogicalRouter
	}

	mSecurityGroupObj := vmap["security_group"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSecurityGroupObj) {
		log.Printf("Setting security_group  SecurityGroup <<%T>> => %#v", mSecurityGroupObj, mSecurityGroupObj)
		mSecurityGroup := mSecurityGroupObj.(int)
		object.SecurityGroup = mSecurityGroup
	}

	mSecurityGroupRuleObj := vmap["security_group_rule"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSecurityGroupRuleObj) {
		log.Printf("Setting security_group_rule  SecurityGroupRule <<%T>> => %#v", mSecurityGroupRuleObj, mSecurityGroupRuleObj)
		mSecurityGroupRule := mSecurityGroupRuleObj.(int)
		object.SecurityGroupRule = mSecurityGroupRule
	}

	mSubnetObj := vmap["subnet"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mSubnetObj) {
		log.Printf("Setting subnet  Subnet <<%T>> => %#v", mSubnetObj, mSubnetObj)
		mSubnet := mSubnetObj.(int)
		object.Subnet = mSubnet
	}

	mGlobalVrouterConfigObj := vmap["global_vrouter_config"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mGlobalVrouterConfigObj) {
		log.Printf("Setting global_vrouter_config  GlobalVrouterConfig <<%T>> => %#v", mGlobalVrouterConfigObj, mGlobalVrouterConfigObj)
		mGlobalVrouterConfig := mGlobalVrouterConfigObj.(int)
		object.GlobalVrouterConfig = mGlobalVrouterConfig
	}

	mLoadbalancerPoolObj := vmap["loadbalancer_pool"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLoadbalancerPoolObj) {
		log.Printf("Setting loadbalancer_pool  LoadbalancerPool <<%T>> => %#v", mLoadbalancerPoolObj, mLoadbalancerPoolObj)
		mLoadbalancerPool := mLoadbalancerPoolObj.(int)
		object.LoadbalancerPool = mLoadbalancerPool
	}

	mLoadbalancerMemberObj := vmap["loadbalancer_member"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLoadbalancerMemberObj) {
		log.Printf("Setting loadbalancer_member  LoadbalancerMember <<%T>> => %#v", mLoadbalancerMemberObj, mLoadbalancerMemberObj)
		mLoadbalancerMember := mLoadbalancerMemberObj.(int)
		object.LoadbalancerMember = mLoadbalancerMember
	}

	mLoadbalancerHealthmonitorObj := vmap["loadbalancer_healthmonitor"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLoadbalancerHealthmonitorObj) {
		log.Printf("Setting loadbalancer_healthmonitor  LoadbalancerHealthmonitor <<%T>> => %#v", mLoadbalancerHealthmonitorObj, mLoadbalancerHealthmonitorObj)
		mLoadbalancerHealthmonitor := mLoadbalancerHealthmonitorObj.(int)
		object.LoadbalancerHealthmonitor = mLoadbalancerHealthmonitor
	}

	mVirtualIpObj := vmap["virtual_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mVirtualIpObj) {
		log.Printf("Setting virtual_ip  VirtualIp <<%T>> => %#v", mVirtualIpObj, mVirtualIpObj)
		mVirtualIp := mVirtualIpObj.(int)
		object.VirtualIp = mVirtualIp
	}

	log.Printf("FINISHED QuotaType object: %#v", object)
}

func TakeQuotaTypeAsMap(object *QuotaType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["defaults"] = object.Defaults
	omap["floating_ip"] = object.FloatingIp
	omap["instance_ip"] = object.InstanceIp
	omap["virtual_machine_interface"] = object.VirtualMachineInterface
	omap["virtual_network"] = object.VirtualNetwork
	omap["virtual_router"] = object.VirtualRouter
	omap["virtual_DNS"] = object.VirtualDns
	omap["virtual_DNS_record"] = object.VirtualDnsRecord
	omap["bgp_router"] = object.BgpRouter
	omap["network_ipam"] = object.NetworkIpam
	omap["access_control_list"] = object.AccessControlList
	omap["network_policy"] = object.NetworkPolicy
	omap["floating_ip_pool"] = object.FloatingIpPool
	omap["service_template"] = object.ServiceTemplate
	omap["service_instance"] = object.ServiceInstance
	omap["logical_router"] = object.LogicalRouter
	omap["security_group"] = object.SecurityGroup
	omap["security_group_rule"] = object.SecurityGroupRule
	omap["subnet"] = object.Subnet
	omap["global_vrouter_config"] = object.GlobalVrouterConfig
	omap["loadbalancer_pool"] = object.LoadbalancerPool
	omap["loadbalancer_member"] = object.LoadbalancerMember
	omap["loadbalancer_healthmonitor"] = object.LoadbalancerHealthmonitor
	omap["virtual_ip"] = object.VirtualIp

	return omap
}

func ResourceQuotaTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"defaults": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"floating_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"instance_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_machine_interface": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_network": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_router": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_dns": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_dns_record": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"bgp_router": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"network_ipam": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"access_control_list": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"network_policy": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"floating_ip_pool": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"service_template": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"service_instance": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"logical_router": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"security_group": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"security_group_rule": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"subnet": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"global_vrouter_config": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"loadbalancer_pool": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"loadbalancer_member": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"loadbalancer_healthmonitor": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"virtual_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceQuotaType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceQuotaTypeSchema(),
	}
}

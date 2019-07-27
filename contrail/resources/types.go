// Automatically generated. DO NOT EDIT.
// (Type Map)
//

package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

var _ = log.Printf // Avoid import errors if not used

var ContrailResourcesMap = map[string]*schema.Resource{
	"contrail_access_control_list":            ResourceAccessControlList(),
	"contrail_alarm":                          ResourceAlarm(),
	"contrail_alias_ip":                       ResourceAliasIp(),
	"contrail_alias_ip_refs":                  ResourceAliasIpRefs(),
	"contrail_alias_ip_pool":                  ResourceAliasIpPool(),
	"contrail_analytics_node":                 ResourceAnalyticsNode(),
	"contrail_api_access_list":                ResourceApiAccessList(),
	"contrail_bgp_as_a_service":               ResourceBgpAsAService(),
	"contrail_bgp_as_a_service_refs":          ResourceBgpAsAServiceRefs(),
	"contrail_bgp_router":                     ResourceBgpRouter(),
	"contrail_config_node":                    ResourceConfigNode(),
	"contrail_config_root":                    ResourceConfigRoot(),
	"contrail_customer_attachment":            ResourceCustomerAttachment(),
	"contrail_customer_attachment_refs":       ResourceCustomerAttachmentRefs(),
	"contrail_database_node":                  ResourceDatabaseNode(),
	"contrail_discovery_service_assignment":   ResourceDiscoveryServiceAssignment(),
	"contrail_domain":                         ResourceDomain(),
	"contrail_dsa_rule":                       ResourceDsaRule(),
	"contrail_floating_ip":                    ResourceFloatingIp(),
	"contrail_floating_ip_refs":               ResourceFloatingIpRefs(),
	"contrail_floating_ip_pool":               ResourceFloatingIpPool(),
	"contrail_forwarding_class":               ResourceForwardingClass(),
	"contrail_forwarding_class_refs":          ResourceForwardingClassRefs(),
	"contrail_global_qos_config":              ResourceGlobalQosConfig(),
	"contrail_global_system_config":           ResourceGlobalSystemConfig(),
	"contrail_global_system_config_refs":      ResourceGlobalSystemConfigRefs(),
	"contrail_global_vrouter_config":          ResourceGlobalVrouterConfig(),
	"contrail_instance_ip":                    ResourceInstanceIp(),
	"contrail_instance_ip_refs":               ResourceInstanceIpRefs(),
	"contrail_interface_route_table":          ResourceInterfaceRouteTable(),
	"contrail_interface_route_table_refs":     ResourceInterfaceRouteTableRefs(),
	"contrail_loadbalancer":                   ResourceLoadbalancer(),
	"contrail_loadbalancer_refs":              ResourceLoadbalancerRefs(),
	"contrail_loadbalancer_healthmonitor":     ResourceLoadbalancerHealthmonitor(),
	"contrail_loadbalancer_listener":          ResourceLoadbalancerListener(),
	"contrail_loadbalancer_listener_refs":     ResourceLoadbalancerListenerRefs(),
	"contrail_loadbalancer_member":            ResourceLoadbalancerMember(),
	"contrail_loadbalancer_pool":              ResourceLoadbalancerPool(),
	"contrail_loadbalancer_pool_refs":         ResourceLoadbalancerPoolRefs(),
	"contrail_logical_interface":              ResourceLogicalInterface(),
	"contrail_logical_interface_refs":         ResourceLogicalInterfaceRefs(),
	"contrail_logical_router":                 ResourceLogicalRouter(),
	"contrail_logical_router_refs":            ResourceLogicalRouterRefs(),
	"contrail_namespace":                      ResourceNamespace(),
	"contrail_network_ipam":                   ResourceNetworkIpam(),
	"contrail_network_ipam_refs":              ResourceNetworkIpamRefs(),
	"contrail_network_policy":                 ResourceNetworkPolicy(),
	"contrail_physical_interface":             ResourcePhysicalInterface(),
	"contrail_physical_interface_refs":        ResourcePhysicalInterfaceRefs(),
	"contrail_physical_router":                ResourcePhysicalRouter(),
	"contrail_physical_router_refs":           ResourcePhysicalRouterRefs(),
	"contrail_port_tuple":                     ResourcePortTuple(),
	"contrail_project":                        ResourceProject(),
	"contrail_project_refs":                   ResourceProjectRefs(),
	"contrail_provider_attachment":            ResourceProviderAttachment(),
	"contrail_provider_attachment_refs":       ResourceProviderAttachmentRefs(),
	"contrail_qos_config":                     ResourceQosConfig(),
	"contrail_qos_config_refs":                ResourceQosConfigRefs(),
	"contrail_qos_queue":                      ResourceQosQueue(),
	"contrail_route_aggregate":                ResourceRouteAggregate(),
	"contrail_route_aggregate_refs":           ResourceRouteAggregateRefs(),
	"contrail_route_table":                    ResourceRouteTable(),
	"contrail_route_target":                   ResourceRouteTarget(),
	"contrail_routing_instance":               ResourceRoutingInstance(),
	"contrail_routing_policy":                 ResourceRoutingPolicy(),
	"contrail_routing_policy_refs":            ResourceRoutingPolicyRefs(),
	"contrail_security_group":                 ResourceSecurityGroup(),
	"contrail_service_appliance":              ResourceServiceAppliance(),
	"contrail_service_appliance_refs":         ResourceServiceApplianceRefs(),
	"contrail_service_appliance_set":          ResourceServiceApplianceSet(),
	"contrail_service_health_check":           ResourceServiceHealthCheck(),
	"contrail_service_health_check_refs":      ResourceServiceHealthCheckRefs(),
	"contrail_service_instance":               ResourceServiceInstance(),
	"contrail_service_instance_refs":          ResourceServiceInstanceRefs(),
	"contrail_service_template":               ResourceServiceTemplate(),
	"contrail_service_template_refs":          ResourceServiceTemplateRefs(),
	"contrail_subnet":                         ResourceSubnet(),
	"contrail_subnet_refs":                    ResourceSubnetRefs(),
	"contrail_virtual_dns":                    ResourceVirtualDns(),
	"contrail_virtual_dns_record":             ResourceVirtualDnsRecord(),
	"contrail_virtual_ip":                     ResourceVirtualIp(),
	"contrail_virtual_ip_refs":                ResourceVirtualIpRefs(),
	"contrail_virtual_machine":                ResourceVirtualMachine(),
	"contrail_virtual_machine_refs":           ResourceVirtualMachineRefs(),
	"contrail_virtual_machine_interface":      ResourceVirtualMachineInterface(),
	"contrail_virtual_machine_interface_refs": ResourceVirtualMachineInterfaceRefs(),
	"contrail_virtual_network":                ResourceVirtualNetwork(),
	"contrail_virtual_network_refs":           ResourceVirtualNetworkRefs(),
	"contrail_virtual_router":                 ResourceVirtualRouter(),
	"contrail_virtual_router_refs":            ResourceVirtualRouterRefs(),
}

func CheckTerraformMap(x interface{}) bool {
	//log.Printf("Checking %T => %#v", x, x)
	slice, ok := x.([]interface{})
	if ok && len(slice) > 0 {
		return true
	}
	switch x.(type) {
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case bool:
		return true
	case float32:
		return true
	case float64:
		return true
	case string:
		return true
	case int:
		return true
	}
	return false
}

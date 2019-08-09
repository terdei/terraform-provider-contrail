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
	"contrail_access_control_list":               ResourceAccessControlList(),
	"contrail_access_control_list_refs":          ResourceAccessControlListRefs(),
	"contrail_address_group":                     ResourceAddressGroup(),
	"contrail_address_group_refs":                ResourceAddressGroupRefs(),
	"contrail_alarm":                             ResourceAlarm(),
	"contrail_alarm_refs":                        ResourceAlarmRefs(),
	"contrail_alias_ip":                          ResourceAliasIp(),
	"contrail_alias_ip_refs":                     ResourceAliasIpRefs(),
	"contrail_alias_ip_pool":                     ResourceAliasIpPool(),
	"contrail_alias_ip_pool_refs":                ResourceAliasIpPoolRefs(),
	"contrail_analytics_node":                    ResourceAnalyticsNode(),
	"contrail_analytics_node_refs":               ResourceAnalyticsNodeRefs(),
	"contrail_api_access_list":                   ResourceApiAccessList(),
	"contrail_api_access_list_refs":              ResourceApiAccessListRefs(),
	"contrail_application_policy_set":            ResourceApplicationPolicySet(),
	"contrail_application_policy_set_refs":       ResourceApplicationPolicySetRefs(),
	"contrail_bgp_as_a_service":                  ResourceBgpAsAService(),
	"contrail_bgp_as_a_service_refs":             ResourceBgpAsAServiceRefs(),
	"contrail_bgp_router":                        ResourceBgpRouter(),
	"contrail_bgp_router_refs":                   ResourceBgpRouterRefs(),
	"contrail_bgpvpn":                            ResourceBgpvpn(),
	"contrail_bgpvpn_refs":                       ResourceBgpvpnRefs(),
	"contrail_bridge_domain":                     ResourceBridgeDomain(),
	"contrail_bridge_domain_refs":                ResourceBridgeDomainRefs(),
	"contrail_config_node":                       ResourceConfigNode(),
	"contrail_config_node_refs":                  ResourceConfigNodeRefs(),
	"contrail_config_root":                       ResourceConfigRoot(),
	"contrail_config_root_refs":                  ResourceConfigRootRefs(),
	"contrail_customer_attachment":               ResourceCustomerAttachment(),
	"contrail_customer_attachment_refs":          ResourceCustomerAttachmentRefs(),
	"contrail_database_node":                     ResourceDatabaseNode(),
	"contrail_database_node_refs":                ResourceDatabaseNodeRefs(),
	"contrail_discovery_service_assignment":      ResourceDiscoveryServiceAssignment(),
	"contrail_discovery_service_assignment_refs": ResourceDiscoveryServiceAssignmentRefs(),
	"contrail_domain":                            ResourceDomain(),
	"contrail_domain_refs":                       ResourceDomainRefs(),
	"contrail_dsa_rule":                          ResourceDsaRule(),
	"contrail_dsa_rule_refs":                     ResourceDsaRuleRefs(),
	"contrail_e2_service_provider":               ResourceE2ServiceProvider(),
	"contrail_e2_service_provider_refs":          ResourceE2ServiceProviderRefs(),
	"contrail_firewall_policy":                   ResourceFirewallPolicy(),
	"contrail_firewall_policy_refs":              ResourceFirewallPolicyRefs(),
	"contrail_firewall_rule":                     ResourceFirewallRule(),
	"contrail_firewall_rule_refs":                ResourceFirewallRuleRefs(),
	"contrail_floating_ip":                       ResourceFloatingIp(),
	"contrail_floating_ip_refs":                  ResourceFloatingIpRefs(),
	"contrail_floating_ip_pool":                  ResourceFloatingIpPool(),
	"contrail_floating_ip_pool_refs":             ResourceFloatingIpPoolRefs(),
	"contrail_forwarding_class":                  ResourceForwardingClass(),
	"contrail_forwarding_class_refs":             ResourceForwardingClassRefs(),
	"contrail_global_qos_config":                 ResourceGlobalQosConfig(),
	"contrail_global_qos_config_refs":            ResourceGlobalQosConfigRefs(),
	"contrail_global_system_config":              ResourceGlobalSystemConfig(),
	"contrail_global_system_config_refs":         ResourceGlobalSystemConfigRefs(),
	"contrail_global_vrouter_config":             ResourceGlobalVrouterConfig(),
	"contrail_global_vrouter_config_refs":        ResourceGlobalVrouterConfigRefs(),
	"contrail_instance_ip":                       ResourceInstanceIp(),
	"contrail_instance_ip_refs":                  ResourceInstanceIpRefs(),
	"contrail_interface_route_table":             ResourceInterfaceRouteTable(),
	"contrail_interface_route_table_refs":        ResourceInterfaceRouteTableRefs(),
	"contrail_loadbalancer":                      ResourceLoadbalancer(),
	"contrail_loadbalancer_refs":                 ResourceLoadbalancerRefs(),
	"contrail_loadbalancer_healthmonitor":        ResourceLoadbalancerHealthmonitor(),
	"contrail_loadbalancer_healthmonitor_refs":   ResourceLoadbalancerHealthmonitorRefs(),
	"contrail_loadbalancer_listener":             ResourceLoadbalancerListener(),
	"contrail_loadbalancer_listener_refs":        ResourceLoadbalancerListenerRefs(),
	"contrail_loadbalancer_member":               ResourceLoadbalancerMember(),
	"contrail_loadbalancer_member_refs":          ResourceLoadbalancerMemberRefs(),
	"contrail_loadbalancer_pool":                 ResourceLoadbalancerPool(),
	"contrail_loadbalancer_pool_refs":            ResourceLoadbalancerPoolRefs(),
	"contrail_logical_interface":                 ResourceLogicalInterface(),
	"contrail_logical_interface_refs":            ResourceLogicalInterfaceRefs(),
	"contrail_logical_router":                    ResourceLogicalRouter(),
	"contrail_logical_router_refs":               ResourceLogicalRouterRefs(),
	"contrail_namespace":                         ResourceNamespace(),
	"contrail_namespace_refs":                    ResourceNamespaceRefs(),
	"contrail_network_device_config":             ResourceNetworkDeviceConfig(),
	"contrail_network_device_config_refs":        ResourceNetworkDeviceConfigRefs(),
	"contrail_network_ipam":                      ResourceNetworkIpam(),
	"contrail_network_ipam_refs":                 ResourceNetworkIpamRefs(),
	"contrail_network_policy":                    ResourceNetworkPolicy(),
	"contrail_network_policy_refs":               ResourceNetworkPolicyRefs(),
	"contrail_peering_policy":                    ResourcePeeringPolicy(),
	"contrail_peering_policy_refs":               ResourcePeeringPolicyRefs(),
	"contrail_physical_interface":                ResourcePhysicalInterface(),
	"contrail_physical_interface_refs":           ResourcePhysicalInterfaceRefs(),
	"contrail_physical_router":                   ResourcePhysicalRouter(),
	"contrail_physical_router_refs":              ResourcePhysicalRouterRefs(),
	"contrail_policy_management":                 ResourcePolicyManagement(),
	"contrail_policy_management_refs":            ResourcePolicyManagementRefs(),
	"contrail_port_tuple":                        ResourcePortTuple(),
	"contrail_port_tuple_refs":                   ResourcePortTupleRefs(),
	"contrail_project":                           ResourceProject(),
	"contrail_project_refs":                      ResourceProjectRefs(),
	"contrail_provider_attachment":               ResourceProviderAttachment(),
	"contrail_provider_attachment_refs":          ResourceProviderAttachmentRefs(),
	"contrail_qos_config":                        ResourceQosConfig(),
	"contrail_qos_config_refs":                   ResourceQosConfigRefs(),
	"contrail_qos_queue":                         ResourceQosQueue(),
	"contrail_qos_queue_refs":                    ResourceQosQueueRefs(),
	"contrail_route_aggregate":                   ResourceRouteAggregate(),
	"contrail_route_aggregate_refs":              ResourceRouteAggregateRefs(),
	"contrail_route_table":                       ResourceRouteTable(),
	"contrail_route_table_refs":                  ResourceRouteTableRefs(),
	"contrail_route_target":                      ResourceRouteTarget(),
	"contrail_route_target_refs":                 ResourceRouteTargetRefs(),
	"contrail_routing_instance":                  ResourceRoutingInstance(),
	"contrail_routing_instance_refs":             ResourceRoutingInstanceRefs(),
	"contrail_routing_policy":                    ResourceRoutingPolicy(),
	"contrail_routing_policy_refs":               ResourceRoutingPolicyRefs(),
	"contrail_security_group":                    ResourceSecurityGroup(),
	"contrail_security_group_refs":               ResourceSecurityGroupRefs(),
	"contrail_security_logging_object":           ResourceSecurityLoggingObject(),
	"contrail_security_logging_object_refs":      ResourceSecurityLoggingObjectRefs(),
	"contrail_service_appliance":                 ResourceServiceAppliance(),
	"contrail_service_appliance_refs":            ResourceServiceApplianceRefs(),
	"contrail_service_appliance_set":             ResourceServiceApplianceSet(),
	"contrail_service_appliance_set_refs":        ResourceServiceApplianceSetRefs(),
	"contrail_service_connection_module":         ResourceServiceConnectionModule(),
	"contrail_service_connection_module_refs":    ResourceServiceConnectionModuleRefs(),
	"contrail_service_endpoint":                  ResourceServiceEndpoint(),
	"contrail_service_endpoint_refs":             ResourceServiceEndpointRefs(),
	"contrail_service_group":                     ResourceServiceGroup(),
	"contrail_service_group_refs":                ResourceServiceGroupRefs(),
	"contrail_service_health_check":              ResourceServiceHealthCheck(),
	"contrail_service_health_check_refs":         ResourceServiceHealthCheckRefs(),
	"contrail_service_instance":                  ResourceServiceInstance(),
	"contrail_service_instance_refs":             ResourceServiceInstanceRefs(),
	"contrail_service_object":                    ResourceServiceObject(),
	"contrail_service_object_refs":               ResourceServiceObjectRefs(),
	"contrail_service_template":                  ResourceServiceTemplate(),
	"contrail_service_template_refs":             ResourceServiceTemplateRefs(),
	"contrail_subnet":                            ResourceSubnet(),
	"contrail_subnet_refs":                       ResourceSubnetRefs(),
	"contrail_tag":                               ResourceTag(),
	"contrail_tag_refs":                          ResourceTagRefs(),
	"contrail_tag_type":                          ResourceTagType(),
	"contrail_tag_type_refs":                     ResourceTagTypeRefs(),
	"contrail_virtual_dns":                       ResourceVirtualDns(),
	"contrail_virtual_dns_refs":                  ResourceVirtualDnsRefs(),
	"contrail_virtual_dns_record":                ResourceVirtualDnsRecord(),
	"contrail_virtual_dns_record_refs":           ResourceVirtualDnsRecordRefs(),
	"contrail_virtual_ip":                        ResourceVirtualIp(),
	"contrail_virtual_ip_refs":                   ResourceVirtualIpRefs(),
	"contrail_virtual_machine":                   ResourceVirtualMachine(),
	"contrail_virtual_machine_refs":              ResourceVirtualMachineRefs(),
	"contrail_virtual_machine_interface":         ResourceVirtualMachineInterface(),
	"contrail_virtual_machine_interface_refs":    ResourceVirtualMachineInterfaceRefs(),
	"contrail_virtual_network":                   ResourceVirtualNetwork(),
	"contrail_virtual_network_refs":              ResourceVirtualNetworkRefs(),
	"contrail_virtual_router":                    ResourceVirtualRouter(),
	"contrail_virtual_router_refs":               ResourceVirtualRouterRefs(),
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

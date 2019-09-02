//
// Automatically generated. DO NOT EDIT.
// (Object)
//

package resources

import (
	"fmt"
	"log"
	"strings"

	"github.com/Juniper/contrail-go-api"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
)

var _ = spew.Dump // Avoid import errors if not used

func SetVirtualMachineInterfaceFromResource(object *VirtualMachineInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetVirtualMachineInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("ecmp_hashing_include_fields"); ok {
		member := new(EcmpHashingIncludeFields)
		SetEcmpHashingIncludeFieldsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetEcmpHashingIncludeFields(member)
	}
	if val, ok := d.GetOk("port_security_enabled"); ok {
		object.SetPortSecurityEnabled(val.(bool))
	}
	if val, ok := d.GetOk("virtual_machine_interface_mac_addresses"); ok {
		member := new(MacAddressesType)
		SetMacAddressesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceMacAddresses(member)
	}
	if val, ok := d.GetOk("virtual_machine_interface_dhcp_option_list"); ok {
		member := new(DhcpOptionsListType)
		SetDhcpOptionsListTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceDhcpOptionList(member)
	}
	if val, ok := d.GetOk("virtual_machine_interface_host_routes"); ok {
		member := new(RouteTableType)
		SetRouteTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceHostRoutes(member)
	}
	if val, ok := d.GetOk("virtual_machine_interface_allowed_address_pairs"); ok {
		member := new(AllowedAddressPairs)
		SetAllowedAddressPairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceAllowedAddressPairs(member)
	}
	if val, ok := d.GetOk("vrf_assign_table"); ok {
		member := new(VrfAssignTableType)
		SetVrfAssignTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVrfAssignTable(member)
	}
	if val, ok := d.GetOk("virtual_machine_interface_device_owner"); ok {
		object.SetVirtualMachineInterfaceDeviceOwner(val.(string))
	}
	if val, ok := d.GetOk("virtual_machine_interface_disable_policy"); ok {
		object.SetVirtualMachineInterfaceDisablePolicy(val.(bool))
	}
	if val, ok := d.GetOk("virtual_machine_interface_properties"); ok {
		member := new(VirtualMachineInterfacePropertiesType)
		SetVirtualMachineInterfacePropertiesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceProperties(member)
	}
	if val, ok := d.GetOk("virtual_machine_interface_bindings"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceBindings(member)
	}
	if val, ok := d.GetOk("virtual_machine_interface_fat_flow_protocols"); ok {
		member := new(FatFlowProtocols)
		SetFatFlowProtocolsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualMachineInterfaceFatFlowProtocols(member)
	}
	if val, ok := d.GetOk("vlan_tag_based_bridge_domain"); ok {
		object.SetVlanTagBasedBridgeDomain(val.(bool))
	}
	if val, ok := d.GetOk("id_perms"); ok {
		member := new(IdPermsType)
		SetIdPermsTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetIdPerms(member)
	}
	if val, ok := d.GetOk("perms2"); ok {
		member := new(PermType2)
		SetPermType2FromMap(member, d, m, (val.([]interface{}))[0])
		object.SetPerms2(member)
	}
	if val, ok := d.GetOk("annotations"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetAnnotations(member)
	}
	if val, ok := d.GetOk("display_name"); ok {
		object.SetDisplayName(val.(string))
	}

}

func SetRefsVirtualMachineInterfaceFromResource(object *VirtualMachineInterface, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualMachineInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("security_logging_object_refs"); ok {
		log.Printf("Got ref security_logging_object_refs -- will call: object.AddSecurityLoggingObject(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("security-logging-object", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving security-logging-object by Uuid = %v as ref for SecurityLoggingObject on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddSecurityLoggingObject(refObj.(*SecurityLoggingObject))
		}
	}
	if val, ok := d.GetOk("qos_config_refs"); ok {
		log.Printf("Got ref qos_config_refs -- will call: object.AddQosConfig(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("qos-config", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving qos-config by Uuid = %v as ref for QosConfig on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddQosConfig(refObj.(*QosConfig))
		}
	}
	if val, ok := d.GetOk("security_group_refs"); ok {
		log.Printf("Got ref security_group_refs -- will call: object.AddSecurityGroup(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("security-group", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving security-group by Uuid = %v as ref for SecurityGroup on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddSecurityGroup(refObj.(*SecurityGroup))
		}
	}
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine-interface by Uuid = %v as ref for VirtualMachineInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
		}
	}
	if val, ok := d.GetOk("virtual_machine_refs"); ok {
		log.Printf("Got ref virtual_machine_refs -- will call: object.AddVirtualMachine(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine by Uuid = %v as ref for VirtualMachine on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachine(refObj.(*VirtualMachine))
		}
	}
	if val, ok := d.GetOk("routing_instance_refs"); ok {
		log.Printf("Got ref routing_instance_refs -- will call: object.AddRoutingInstance(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("routing-instance", refId.(string))
			dataObj := new(PolicyBasedForwardingRuleType)
			SetPolicyBasedForwardingRuleTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving routing-instance by Uuid = %v as ref for RoutingInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddRoutingInstance(refObj.(*RoutingInstance), *dataObj)
		}
	}
	if val, ok := d.GetOk("port_tuple_refs"); ok {
		log.Printf("Got ref port_tuple_refs -- will call: object.AddPortTuple(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("port-tuple", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving port-tuple by Uuid = %v as ref for PortTuple on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPortTuple(refObj.(*PortTuple))
		}
	}
	if val, ok := d.GetOk("service_health_check_refs"); ok {
		log.Printf("Got ref service_health_check_refs -- will call: object.AddServiceHealthCheck(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-health-check", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-health-check by Uuid = %v as ref for ServiceHealthCheck on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceHealthCheck(refObj.(*ServiceHealthCheck))
		}
	}
	if val, ok := d.GetOk("interface_route_table_refs"); ok {
		log.Printf("Got ref interface_route_table_refs -- will call: object.AddInterfaceRouteTable(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("interface-route-table", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving interface-route-table by Uuid = %v as ref for InterfaceRouteTable on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddInterfaceRouteTable(refObj.(*InterfaceRouteTable))
		}
	}
	if val, ok := d.GetOk("physical_interface_refs"); ok {
		log.Printf("Got ref physical_interface_refs -- will call: object.AddPhysicalInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("physical-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving physical-interface by Uuid = %v as ref for PhysicalInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPhysicalInterface(refObj.(*PhysicalInterface))
		}
	}
	if val, ok := d.GetOk("bridge_domain_refs"); ok {
		log.Printf("Got ref bridge_domain_refs -- will call: object.AddBridgeDomain(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("bridge-domain", refId.(string))
			dataObj := new(BridgeDomainMembershipType)
			SetBridgeDomainMembershipTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving bridge-domain by Uuid = %v as ref for BridgeDomain on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddBridgeDomain(refObj.(*BridgeDomain), *dataObj)
		}
	}
	if val, ok := d.GetOk("service_endpoint_refs"); ok {
		log.Printf("Got ref service_endpoint_refs -- will call: object.AddServiceEndpoint(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-endpoint", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-endpoint by Uuid = %v as ref for ServiceEndpoint on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceEndpoint(refObj.(*ServiceEndpoint))
		}
	}
	if val, ok := d.GetOk("tag_refs"); ok {
		log.Printf("Got ref tag_refs -- will call: object.AddTag(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("tag", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving tag by Uuid = %v as ref for Tag on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddTag(refObj.(*Tag))
		}
	}

	return nil
}

func SetReqRefsVirtualMachineInterfaceFromResource(object *VirtualMachineInterface, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualMachineInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_network_refs"); ok {
		log.Printf("Got ref virtual_network_refs -- will call: object.AddVirtualNetwork(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-network", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-network by Uuid = %v as ref for VirtualNetwork on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualNetwork(refObj.(*VirtualNetwork))
		}
	}
	if val, ok := d.GetOk("bgp_router_refs"); ok {
		log.Printf("Got ref bgp_router_refs -- will call: object.AddBgpRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("bgp-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving bgp-router by Uuid = %v as ref for BgpRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddBgpRouter(refObj.(*BgpRouter))
		}
	}

	return nil
}

func DeleteRefsVirtualMachineInterfaceFromResource(object *VirtualMachineInterface, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsVirtualMachineInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("security_logging_object_refs"); ok {
		log.Printf("Got ref security_logging_object_refs -- will call: object.DeleteSecurityLoggingObject(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteSecurityLoggingObject(refId.(string))
		}
	}
	if val, ok := d.GetOk("qos_config_refs"); ok {
		log.Printf("Got ref qos_config_refs -- will call: object.DeleteQosConfig(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteQosConfig(refId.(string))
		}
	}
	if val, ok := d.GetOk("security_group_refs"); ok {
		log.Printf("Got ref security_group_refs -- will call: object.DeleteSecurityGroup(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteSecurityGroup(refId.(string))
		}
	}
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.DeleteVirtualMachineInterface(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualMachineInterface(refId.(string))
		}
	}
	if val, ok := d.GetOk("virtual_machine_refs"); ok {
		log.Printf("Got ref virtual_machine_refs -- will call: object.DeleteVirtualMachine(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualMachine(refId.(string))
		}
	}
	if val, ok := d.GetOk("routing_instance_refs"); ok {
		log.Printf("Got ref routing_instance_refs -- will call: object.DeleteRoutingInstance(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteRoutingInstance(refId.(string))
		}
	}
	if val, ok := d.GetOk("port_tuple_refs"); ok {
		log.Printf("Got ref port_tuple_refs -- will call: object.DeletePortTuple(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePortTuple(refId.(string))
		}
	}
	if val, ok := d.GetOk("service_health_check_refs"); ok {
		log.Printf("Got ref service_health_check_refs -- will call: object.DeleteServiceHealthCheck(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceHealthCheck(refId.(string))
		}
	}
	if val, ok := d.GetOk("interface_route_table_refs"); ok {
		log.Printf("Got ref interface_route_table_refs -- will call: object.DeleteInterfaceRouteTable(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteInterfaceRouteTable(refId.(string))
		}
	}
	if val, ok := d.GetOk("physical_interface_refs"); ok {
		log.Printf("Got ref physical_interface_refs -- will call: object.DeletePhysicalInterface(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePhysicalInterface(refId.(string))
		}
	}
	if val, ok := d.GetOk("bridge_domain_refs"); ok {
		log.Printf("Got ref bridge_domain_refs -- will call: object.DeleteBridgeDomain(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteBridgeDomain(refId.(string))
		}
	}
	if val, ok := d.GetOk("service_endpoint_refs"); ok {
		log.Printf("Got ref service_endpoint_refs -- will call: object.DeleteServiceEndpoint(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceEndpoint(refId.(string))
		}
	}
	if val, ok := d.GetOk("tag_refs"); ok {
		log.Printf("Got ref tag_refs -- will call: object.DeleteTag(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteTag(refId.(string))
		}
	}

	return nil
}

func WriteVirtualMachineInterfaceToResource(object VirtualMachineInterface, d *schema.ResourceData, m interface{}) {

	ecmp_hashing_include_fieldsObj := object.GetEcmpHashingIncludeFields()
	d.Set("ecmp_hashing_include_fields", TakeEcmpHashingIncludeFieldsAsMap(&ecmp_hashing_include_fieldsObj))
	d.Set("port_security_enabled", object.GetPortSecurityEnabled())
	virtual_machine_interface_mac_addressesObj := object.GetVirtualMachineInterfaceMacAddresses()
	d.Set("virtual_machine_interface_mac_addresses", TakeMacAddressesTypeAsMap(&virtual_machine_interface_mac_addressesObj))
	virtual_machine_interface_dhcp_option_listObj := object.GetVirtualMachineInterfaceDhcpOptionList()
	d.Set("virtual_machine_interface_dhcp_option_list", TakeDhcpOptionsListTypeAsMap(&virtual_machine_interface_dhcp_option_listObj))
	virtual_machine_interface_host_routesObj := object.GetVirtualMachineInterfaceHostRoutes()
	d.Set("virtual_machine_interface_host_routes", TakeRouteTableTypeAsMap(&virtual_machine_interface_host_routesObj))
	virtual_machine_interface_allowed_address_pairsObj := object.GetVirtualMachineInterfaceAllowedAddressPairs()
	d.Set("virtual_machine_interface_allowed_address_pairs", TakeAllowedAddressPairsAsMap(&virtual_machine_interface_allowed_address_pairsObj))
	vrf_assign_tableObj := object.GetVrfAssignTable()
	d.Set("vrf_assign_table", TakeVrfAssignTableTypeAsMap(&vrf_assign_tableObj))
	d.Set("virtual_machine_interface_device_owner", object.GetVirtualMachineInterfaceDeviceOwner())
	d.Set("virtual_machine_interface_disable_policy", object.GetVirtualMachineInterfaceDisablePolicy())
	virtual_machine_interface_propertiesObj := object.GetVirtualMachineInterfaceProperties()
	d.Set("virtual_machine_interface_properties", TakeVirtualMachineInterfacePropertiesTypeAsMap(&virtual_machine_interface_propertiesObj))
	virtual_machine_interface_bindingsObj := object.GetVirtualMachineInterfaceBindings()
	d.Set("virtual_machine_interface_bindings", TakeKeyValuePairsAsMap(&virtual_machine_interface_bindingsObj))
	virtual_machine_interface_fat_flow_protocolsObj := object.GetVirtualMachineInterfaceFatFlowProtocols()
	d.Set("virtual_machine_interface_fat_flow_protocols", TakeFatFlowProtocolsAsMap(&virtual_machine_interface_fat_flow_protocolsObj))
	d.Set("vlan_tag_based_bridge_domain", object.GetVlanTagBasedBridgeDomain())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

	if ref, err := object.GetVirtualNetworkRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_network_refs", refList)
	}
	if ref, err := object.GetBgpRouterRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("bgp_router_refs", refList)
	}
}

func WriteVirtualMachineInterfaceRefsToResource(object VirtualMachineInterface, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetSecurityLoggingObjectRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("security_logging_object_refs", refList)
	}
	if ref, err := object.GetQosConfigRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("qos_config_refs", refList)
	}
	if ref, err := object.GetSecurityGroupRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("security_group_refs", refList)
	}
	if ref, err := object.GetVirtualMachineInterfaceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_machine_interface_refs", refList)
	}
	if ref, err := object.GetVirtualMachineRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_machine_refs", refList)
	}
	if ref, err := object.GetRoutingInstanceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("routing_instance_refs", refList)
	}
	if ref, err := object.GetPortTupleRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("port_tuple_refs", refList)
	}
	if ref, err := object.GetServiceHealthCheckRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_health_check_refs", refList)
	}
	if ref, err := object.GetInterfaceRouteTableRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("interface_route_table_refs", refList)
	}
	if ref, err := object.GetPhysicalInterfaceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("physical_interface_refs", refList)
	}
	if ref, err := object.GetBridgeDomainRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("bridge_domain_refs", refList)
	}
	if ref, err := object.GetServiceEndpointRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_endpoint_refs", refList)
	}
	if ref, err := object.GetTagRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("tag_refs", refList)
	}
}

func TakeVirtualMachineInterfaceAsMap(object *VirtualMachineInterface) map[string]interface{} {
	omap := make(map[string]interface{})

	ecmp_hashing_include_fieldsObj := object.GetEcmpHashingIncludeFields()
	omap["ecmp_hashing_include_fields"] = TakeEcmpHashingIncludeFieldsAsMap(&ecmp_hashing_include_fieldsObj)
	omap["port_security_enabled"] = object.GetPortSecurityEnabled()
	virtual_machine_interface_mac_addressesObj := object.GetVirtualMachineInterfaceMacAddresses()
	omap["virtual_machine_interface_mac_addresses"] = TakeMacAddressesTypeAsMap(&virtual_machine_interface_mac_addressesObj)
	virtual_machine_interface_dhcp_option_listObj := object.GetVirtualMachineInterfaceDhcpOptionList()
	omap["virtual_machine_interface_dhcp_option_list"] = TakeDhcpOptionsListTypeAsMap(&virtual_machine_interface_dhcp_option_listObj)
	virtual_machine_interface_host_routesObj := object.GetVirtualMachineInterfaceHostRoutes()
	omap["virtual_machine_interface_host_routes"] = TakeRouteTableTypeAsMap(&virtual_machine_interface_host_routesObj)
	virtual_machine_interface_allowed_address_pairsObj := object.GetVirtualMachineInterfaceAllowedAddressPairs()
	omap["virtual_machine_interface_allowed_address_pairs"] = TakeAllowedAddressPairsAsMap(&virtual_machine_interface_allowed_address_pairsObj)
	vrf_assign_tableObj := object.GetVrfAssignTable()
	omap["vrf_assign_table"] = TakeVrfAssignTableTypeAsMap(&vrf_assign_tableObj)
	omap["virtual_machine_interface_device_owner"] = object.GetVirtualMachineInterfaceDeviceOwner()
	omap["virtual_machine_interface_disable_policy"] = object.GetVirtualMachineInterfaceDisablePolicy()
	virtual_machine_interface_propertiesObj := object.GetVirtualMachineInterfaceProperties()
	omap["virtual_machine_interface_properties"] = TakeVirtualMachineInterfacePropertiesTypeAsMap(&virtual_machine_interface_propertiesObj)
	virtual_machine_interface_bindingsObj := object.GetVirtualMachineInterfaceBindings()
	omap["virtual_machine_interface_bindings"] = TakeKeyValuePairsAsMap(&virtual_machine_interface_bindingsObj)
	virtual_machine_interface_fat_flow_protocolsObj := object.GetVirtualMachineInterfaceFatFlowProtocols()
	omap["virtual_machine_interface_fat_flow_protocols"] = TakeFatFlowProtocolsAsMap(&virtual_machine_interface_fat_flow_protocolsObj)
	omap["vlan_tag_based_bridge_domain"] = object.GetVlanTagBasedBridgeDomain()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateVirtualMachineInterfaceFromResource(object *VirtualMachineInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("ecmp_hashing_include_fields") {
		if val, ok := d.GetOk("ecmp_hashing_include_fields"); ok {
			member := new(EcmpHashingIncludeFields)
			SetEcmpHashingIncludeFieldsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetEcmpHashingIncludeFields(member)
		}
	}
	if d.HasChange("port_security_enabled") {
		if val, ok := d.GetOk("port_security_enabled"); ok {
			object.SetPortSecurityEnabled(val.(bool))
		}
	}
	if d.HasChange("virtual_machine_interface_mac_addresses") {
		if val, ok := d.GetOk("virtual_machine_interface_mac_addresses"); ok {
			member := new(MacAddressesType)
			SetMacAddressesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceMacAddresses(member)
		}
	}
	if d.HasChange("virtual_machine_interface_dhcp_option_list") {
		if val, ok := d.GetOk("virtual_machine_interface_dhcp_option_list"); ok {
			member := new(DhcpOptionsListType)
			SetDhcpOptionsListTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceDhcpOptionList(member)
		}
	}
	if d.HasChange("virtual_machine_interface_host_routes") {
		if val, ok := d.GetOk("virtual_machine_interface_host_routes"); ok {
			member := new(RouteTableType)
			SetRouteTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceHostRoutes(member)
		}
	}
	if d.HasChange("virtual_machine_interface_allowed_address_pairs") {
		if val, ok := d.GetOk("virtual_machine_interface_allowed_address_pairs"); ok {
			member := new(AllowedAddressPairs)
			SetAllowedAddressPairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceAllowedAddressPairs(member)
		}
	}
	if d.HasChange("vrf_assign_table") {
		if val, ok := d.GetOk("vrf_assign_table"); ok {
			member := new(VrfAssignTableType)
			SetVrfAssignTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVrfAssignTable(member)
		}
	}
	if d.HasChange("virtual_machine_interface_device_owner") {
		if val, ok := d.GetOk("virtual_machine_interface_device_owner"); ok {
			object.SetVirtualMachineInterfaceDeviceOwner(val.(string))
		}
	}
	if d.HasChange("virtual_machine_interface_disable_policy") {
		if val, ok := d.GetOk("virtual_machine_interface_disable_policy"); ok {
			object.SetVirtualMachineInterfaceDisablePolicy(val.(bool))
		}
	}
	if d.HasChange("virtual_machine_interface_properties") {
		if val, ok := d.GetOk("virtual_machine_interface_properties"); ok {
			member := new(VirtualMachineInterfacePropertiesType)
			SetVirtualMachineInterfacePropertiesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceProperties(member)
		}
	}
	if d.HasChange("virtual_machine_interface_bindings") {
		if val, ok := d.GetOk("virtual_machine_interface_bindings"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceBindings(member)
		}
	}
	if d.HasChange("virtual_machine_interface_fat_flow_protocols") {
		if val, ok := d.GetOk("virtual_machine_interface_fat_flow_protocols"); ok {
			member := new(FatFlowProtocols)
			SetFatFlowProtocolsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualMachineInterfaceFatFlowProtocols(member)
		}
	}
	if d.HasChange("vlan_tag_based_bridge_domain") {
		if val, ok := d.GetOk("vlan_tag_based_bridge_domain"); ok {
			object.SetVlanTagBasedBridgeDomain(val.(bool))
		}
	}
	if d.HasChange("id_perms") {
		if val, ok := d.GetOk("id_perms"); ok {
			member := new(IdPermsType)
			SetIdPermsTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetIdPerms(member)
		}
	}
	if d.HasChange("perms2") {
		if val, ok := d.GetOk("perms2"); ok {
			member := new(PermType2)
			SetPermType2FromMap(member, d, m, (val.([]interface{}))[0])
			object.SetPerms2(member)
		}
	}
	if d.HasChange("annotations") {
		if val, ok := d.GetOk("annotations"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetAnnotations(member)
		}
	}
	if d.HasChange("display_name") {
		if val, ok := d.GetOk("display_name"); ok {
			object.SetDisplayName(val.(string))
		}
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("virtual_network_refs") {
		object.ClearVirtualNetwork()
		if val, ok := d.GetOk("virtual_network_refs"); ok {
			log.Printf("Got ref virtual_network_refs -- will call: object.AddVirtualNetwork(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("virtual-network", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddVirtualNetwork(refObj.(*VirtualNetwork))
			}
		}
	}
	if d.HasChange("bgp_router_refs") {
		object.ClearBgpRouter()
		if val, ok := d.GetOk("bgp_router_refs"); ok {
			log.Printf("Got ref bgp_router_refs -- will call: object.AddBgpRouter(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("bgp-router", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddBgpRouter(refObj.(*BgpRouter))
			}
		}
	}

}

func UpdateVirtualMachineInterfaceRefsFromResource(object *VirtualMachineInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("security_logging_object_refs") {
		object.ClearSecurityLoggingObject()
		if val, ok := d.GetOk("security_logging_object_refs"); ok {
			log.Printf("Got ref security_logging_object_refs -- will call: object.AddSecurityLoggingObject(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("security-logging-object", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddSecurityLoggingObject(refObj.(*SecurityLoggingObject))
			}
		}
	}
	if d.HasChange("qos_config_refs") {
		object.ClearQosConfig()
		if val, ok := d.GetOk("qos_config_refs"); ok {
			log.Printf("Got ref qos_config_refs -- will call: object.AddQosConfig(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("qos-config", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddQosConfig(refObj.(*QosConfig))
			}
		}
	}
	if d.HasChange("security_group_refs") {
		object.ClearSecurityGroup()
		if val, ok := d.GetOk("security_group_refs"); ok {
			log.Printf("Got ref security_group_refs -- will call: object.AddSecurityGroup(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("security-group", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddSecurityGroup(refObj.(*SecurityGroup))
			}
		}
	}
	if d.HasChange("virtual_machine_interface_refs") {
		object.ClearVirtualMachineInterface()
		if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
			log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("virtual-machine-interface", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
			}
		}
	}
	if d.HasChange("virtual_machine_refs") {
		object.ClearVirtualMachine()
		if val, ok := d.GetOk("virtual_machine_refs"); ok {
			log.Printf("Got ref virtual_machine_refs -- will call: object.AddVirtualMachine(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("virtual-machine", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddVirtualMachine(refObj.(*VirtualMachine))
			}
		}
	}
	if d.HasChange("routing_instance_refs") {
		object.ClearRoutingInstance()
		if val, ok := d.GetOk("routing_instance_refs"); ok {
			log.Printf("Got ref routing_instance_refs -- will call: object.AddRoutingInstance(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("routing-instance", refId.(string))
				dataObj := new(PolicyBasedForwardingRuleType)
				SetPolicyBasedForwardingRuleTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddRoutingInstance(refObj.(*RoutingInstance), *dataObj)
			}
		}
	}
	if d.HasChange("port_tuple_refs") {
		object.ClearPortTuple()
		if val, ok := d.GetOk("port_tuple_refs"); ok {
			log.Printf("Got ref port_tuple_refs -- will call: object.AddPortTuple(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("port-tuple", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddPortTuple(refObj.(*PortTuple))
			}
		}
	}
	if d.HasChange("service_health_check_refs") {
		object.ClearServiceHealthCheck()
		if val, ok := d.GetOk("service_health_check_refs"); ok {
			log.Printf("Got ref service_health_check_refs -- will call: object.AddServiceHealthCheck(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-health-check", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceHealthCheck(refObj.(*ServiceHealthCheck))
			}
		}
	}
	if d.HasChange("interface_route_table_refs") {
		object.ClearInterfaceRouteTable()
		if val, ok := d.GetOk("interface_route_table_refs"); ok {
			log.Printf("Got ref interface_route_table_refs -- will call: object.AddInterfaceRouteTable(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("interface-route-table", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddInterfaceRouteTable(refObj.(*InterfaceRouteTable))
			}
		}
	}
	if d.HasChange("physical_interface_refs") {
		object.ClearPhysicalInterface()
		if val, ok := d.GetOk("physical_interface_refs"); ok {
			log.Printf("Got ref physical_interface_refs -- will call: object.AddPhysicalInterface(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("physical-interface", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddPhysicalInterface(refObj.(*PhysicalInterface))
			}
		}
	}
	if d.HasChange("bridge_domain_refs") {
		object.ClearBridgeDomain()
		if val, ok := d.GetOk("bridge_domain_refs"); ok {
			log.Printf("Got ref bridge_domain_refs -- will call: object.AddBridgeDomain(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("bridge-domain", refId.(string))
				dataObj := new(BridgeDomainMembershipType)
				SetBridgeDomainMembershipTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddBridgeDomain(refObj.(*BridgeDomain), *dataObj)
			}
		}
	}
	if d.HasChange("service_endpoint_refs") {
		object.ClearServiceEndpoint()
		if val, ok := d.GetOk("service_endpoint_refs"); ok {
			log.Printf("Got ref service_endpoint_refs -- will call: object.AddServiceEndpoint(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-endpoint", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceEndpoint(refObj.(*ServiceEndpoint))
			}
		}
	}
	if d.HasChange("tag_refs") {
		object.ClearTag()
		if val, ok := d.GetOk("tag_refs"); ok {
			log.Printf("Got ref tag_refs -- will call: object.AddTag(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("tag", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddTag(refObj.(*Tag))
			}
		}
	}

}

func ResourceVirtualMachineInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualMachineInterfaceCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(VirtualMachineInterface)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceVirtualMachineInterfaceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "VirtualMachineInterface", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetVirtualMachineInterfaceFromResource(object, d, m)

	if err := SetReqRefsVirtualMachineInterfaceFromResource(object, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceReqRefsCreate] Set required refs on object VirtualMachineInterface on %v (%v)", client.GetServer(), err)
	}

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceCreate] Creation of resource VirtualMachineInterface on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceVirtualMachineInterfaceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualMachineInterfaceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsCreate] Missing 'uuid' field for resource VirtualMachineInterface")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-machine-interface", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsCreate] Retrieving VirtualMachineInterface with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualMachineInterface := obj.(*VirtualMachineInterface) // Fully set by Contrail backend
	if err := SetRefsVirtualMachineInterfaceFromResource(objVirtualMachineInterface, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsCreate] Set refs on object VirtualMachineInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualMachineInterface.GetHref())
	if err := client.Update(objVirtualMachineInterface); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsCreate] Update refs for resource VirtualMachineInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualMachineInterface.GetUuid())
	return nil
}

func ResourceVirtualMachineInterfaceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualMachineInterfaceRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-machine-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRead] Read resource virtual-machine-interface on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualMachineInterface)
	WriteVirtualMachineInterfaceToResource(*object, d, m)
	return nil
}

func ResourceVirtualMachineInterfaceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualMachineInterfaceRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-machine-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsRead] Read resource virtual-machine-interface on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualMachineInterface)
	WriteVirtualMachineInterfaceRefsToResource(*object, d, m)
	return nil
}

func ResourceVirtualMachineInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualMachineInterfaceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-machine-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceUpdate] Retrieving VirtualMachineInterface with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualMachineInterface)
	UpdateVirtualMachineInterfaceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceUpdate] Update of resource VirtualMachineInterface on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualMachineInterfaceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualMachineInterfaceRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-machine-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsUpdate] Retrieving VirtualMachineInterface with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualMachineInterface)
	UpdateVirtualMachineInterfaceRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsUpdate] Update of resource VirtualMachineInterface on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualMachineInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualMachineInterfaceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("virtual-machine-interface", d.Id()); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceDelete] Deletion of resource virtual-machine-interface on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceVirtualMachineInterfaceRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualMachineInterfaceRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsDelete] Missing 'uuid' field for resource VirtualMachineInterface")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-machine-interface", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsDelete] Retrieving VirtualMachineInterface with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualMachineInterface := obj.(*VirtualMachineInterface) // Fully set by Contrail backend
	if err := DeleteRefsVirtualMachineInterfaceFromResource(objVirtualMachineInterface, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsDelete] Set refs on object VirtualMachineInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualMachineInterface.GetHref())
	if err := client.Update(objVirtualMachineInterface); err != nil {
		return fmt.Errorf("[ResourceVirtualMachineInterfaceRefsDelete] Delete refs for resource VirtualMachineInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualMachineInterface.GetUuid())
	return nil
}

func ResourceVirtualMachineInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"parent_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		// Properties
		"ecmp_hashing_include_fields": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceEcmpHashingIncludeFields(),
		},
		"port_security_enabled": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"virtual_machine_interface_mac_addresses": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMacAddressesType(),
		},
		"virtual_machine_interface_dhcp_option_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDhcpOptionsListType(),
		},
		"virtual_machine_interface_host_routes": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTableType(),
		},
		"virtual_machine_interface_allowed_address_pairs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAllowedAddressPairs(),
		},
		"vrf_assign_table": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVrfAssignTableType(),
		},
		"virtual_machine_interface_device_owner": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"virtual_machine_interface_disable_policy": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"virtual_machine_interface_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachineInterfacePropertiesType(),
		},
		"virtual_machine_interface_bindings": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceKeyValuePairs(),
		},
		"virtual_machine_interface_fat_flow_protocols": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFatFlowProtocols(),
		},
		"vlan_tag_based_bridge_domain": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"id_perms": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIdPermsType(),
		},
		"perms2": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePermType2(),
		},
		"annotations": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceKeyValuePairs(),
		},
		"display_name": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"virtual_network_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"bgp_router_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}

func ResourceVirtualMachineInterfaceRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"security_logging_object_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"qos_config_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"security_group_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"virtual_machine_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"virtual_machine_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"routing_instance_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"attr": &schema.Schema{
						Type:     schema.TypeList,
						Elem:     ResourcePolicyBasedForwardingRuleType(),
						Required: true,
					},
				},
			},
		},
		"port_tuple_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"service_health_check_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"interface_route_table_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"physical_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"bridge_domain_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"attr": &schema.Schema{
						Type:     schema.TypeList,
						Elem:     ResourceBridgeDomainMembershipType(),
						Required: true,
					},
				},
			},
		},
		"service_endpoint_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}

func ResourceVirtualMachineInterface() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualMachineInterfaceCreate,
		Read:   ResourceVirtualMachineInterfaceRead,
		Update: ResourceVirtualMachineInterfaceUpdate,
		Delete: ResourceVirtualMachineInterfaceDelete,
		Schema: ResourceVirtualMachineInterfaceSchema(),
	}
}

func ResourceVirtualMachineInterfaceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualMachineInterfaceRefsCreate,
		Read:   ResourceVirtualMachineInterfaceRefsRead,
		Update: ResourceVirtualMachineInterfaceRefsUpdate,
		Delete: ResourceVirtualMachineInterfaceRefsDelete,
		Schema: ResourceVirtualMachineInterfaceRefsSchema(),
	}
}

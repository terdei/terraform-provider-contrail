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

func SetVirtualNetworkFromResource(object *VirtualNetwork, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetVirtualNetworkFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("ecmp_hashing_include_fields"); ok {
		member := new(EcmpHashingIncludeFields)
		SetEcmpHashingIncludeFieldsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetEcmpHashingIncludeFields(member)
	}
	if val, ok := d.GetOk("virtual_network_properties"); ok {
		member := new(VirtualNetworkType)
		SetVirtualNetworkTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualNetworkProperties(member)
	}
	if val, ok := d.GetOk("provider_properties"); ok {
		member := new(ProviderDetails)
		SetProviderDetailsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetProviderProperties(member)
	}
	if val, ok := d.GetOk("is_provider_network"); ok {
		object.SetIsProviderNetwork(val.(bool))
	}
	if val, ok := d.GetOk("virtual_network_network_id"); ok {
		object.SetVirtualNetworkNetworkId(val.(int))
	}
	if val, ok := d.GetOk("port_security_enabled"); ok {
		object.SetPortSecurityEnabled(val.(bool))
	}
	if val, ok := d.GetOk("route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetRouteTargetList(member)
	}
	if val, ok := d.GetOk("import_route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetImportRouteTargetList(member)
	}
	if val, ok := d.GetOk("export_route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetExportRouteTargetList(member)
	}
	if val, ok := d.GetOk("router_external"); ok {
		object.SetRouterExternal(val.(bool))
	}
	if val, ok := d.GetOk("is_shared"); ok {
		object.SetIsShared(val.(bool))
	}
	if val, ok := d.GetOk("external_ipam"); ok {
		object.SetExternalIpam(val.(bool))
	}
	if val, ok := d.GetOk("flood_unknown_unicast"); ok {
		object.SetFloodUnknownUnicast(val.(bool))
	}
	if val, ok := d.GetOk("multi_policy_service_chains_enabled"); ok {
		object.SetMultiPolicyServiceChainsEnabled(val.(bool))
	}
	if val, ok := d.GetOk("address_allocation_mode"); ok {
		object.SetAddressAllocationMode(val.(string))
	}
	if val, ok := d.GetOk("mac_learning_enabled"); ok {
		object.SetMacLearningEnabled(val.(bool))
	}
	if val, ok := d.GetOk("mac_limit_control"); ok {
		member := new(MACLimitControlType)
		SetMACLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMacLimitControl(member)
	}
	if val, ok := d.GetOk("mac_move_control"); ok {
		member := new(MACMoveLimitControlType)
		SetMACMoveLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMacMoveControl(member)
	}
	if val, ok := d.GetOk("mac_aging_time"); ok {
		object.SetMacAgingTime(val.(int))
	}
	if val, ok := d.GetOk("pbb_evpn_enable"); ok {
		object.SetPbbEvpnEnable(val.(bool))
	}
	if val, ok := d.GetOk("pbb_etree_enable"); ok {
		object.SetPbbEtreeEnable(val.(bool))
	}
	if val, ok := d.GetOk("layer2_control_word"); ok {
		object.SetLayer2ControlWord(val.(bool))
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

func SetRefsVirtualNetworkFromResource(object *VirtualNetwork, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualNetworkFromResource] key = %v, prefix = %v", key, prefix)
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
	if val, ok := d.GetOk("network_ipam_refs"); ok {
		log.Printf("Got ref network_ipam_refs -- will call: object.AddNetworkIpam(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("network-ipam", refId.(string))
			dataObj := new(VnSubnetsType)
			SetVnSubnetsTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving network-ipam by Uuid = %v as ref for NetworkIpam on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddNetworkIpam(refObj.(*NetworkIpam), *dataObj)
		}
	}
	if val, ok := d.GetOk("network_policy_refs"); ok {
		log.Printf("Got ref network_policy_refs -- will call: object.AddNetworkPolicy(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("network-policy", refId.(string))
			dataObj := new(VirtualNetworkPolicyType)
			SetVirtualNetworkPolicyTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving network-policy by Uuid = %v as ref for NetworkPolicy on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddNetworkPolicy(refObj.(*NetworkPolicy), *dataObj)
		}
	}
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
	if val, ok := d.GetOk("route_table_refs"); ok {
		log.Printf("Got ref route_table_refs -- will call: object.AddRouteTable(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("route-table", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving route-table by Uuid = %v as ref for RouteTable on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddRouteTable(refObj.(*RouteTable))
		}
	}
	if val, ok := d.GetOk("bgpvpn_refs"); ok {
		log.Printf("Got ref bgpvpn_refs -- will call: object.AddBgpvpn(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("bgpvpn", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving bgpvpn by Uuid = %v as ref for Bgpvpn on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddBgpvpn(refObj.(*Bgpvpn))
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

func DeleteRefsVirtualNetworkFromResource(object *VirtualNetwork, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsVirtualNetworkFromResource] key = %v, prefix = %v", key, prefix)
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
	if val, ok := d.GetOk("network_ipam_refs"); ok {
		log.Printf("Got ref network_ipam_refs -- will call: object.DeleteNetworkIpam(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteNetworkIpam(refId.(string))
		}
	}
	if val, ok := d.GetOk("network_policy_refs"); ok {
		log.Printf("Got ref network_policy_refs -- will call: object.DeleteNetworkPolicy(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteNetworkPolicy(refId.(string))
		}
	}
	if val, ok := d.GetOk("virtual_network_refs"); ok {
		log.Printf("Got ref virtual_network_refs -- will call: object.DeleteVirtualNetwork(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualNetwork(refId.(string))
		}
	}
	if val, ok := d.GetOk("route_table_refs"); ok {
		log.Printf("Got ref route_table_refs -- will call: object.DeleteRouteTable(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteRouteTable(refId.(string))
		}
	}
	if val, ok := d.GetOk("bgpvpn_refs"); ok {
		log.Printf("Got ref bgpvpn_refs -- will call: object.DeleteBgpvpn(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteBgpvpn(refId.(string))
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

func WriteVirtualNetworkToResource(object VirtualNetwork, d *schema.ResourceData, m interface{}) {

	ecmp_hashing_include_fieldsObj := object.GetEcmpHashingIncludeFields()
	d.Set("ecmp_hashing_include_fields", TakeEcmpHashingIncludeFieldsAsMap(&ecmp_hashing_include_fieldsObj))
	virtual_network_propertiesObj := object.GetVirtualNetworkProperties()
	d.Set("virtual_network_properties", TakeVirtualNetworkTypeAsMap(&virtual_network_propertiesObj))
	provider_propertiesObj := object.GetProviderProperties()
	d.Set("provider_properties", TakeProviderDetailsAsMap(&provider_propertiesObj))
	d.Set("is_provider_network", object.GetIsProviderNetwork())
	d.Set("virtual_network_network_id", object.GetVirtualNetworkNetworkId())
	d.Set("port_security_enabled", object.GetPortSecurityEnabled())
	route_target_listObj := object.GetRouteTargetList()
	d.Set("route_target_list", TakeRouteTargetListAsMap(&route_target_listObj))
	import_route_target_listObj := object.GetImportRouteTargetList()
	d.Set("import_route_target_list", TakeRouteTargetListAsMap(&import_route_target_listObj))
	export_route_target_listObj := object.GetExportRouteTargetList()
	d.Set("export_route_target_list", TakeRouteTargetListAsMap(&export_route_target_listObj))
	d.Set("router_external", object.GetRouterExternal())
	d.Set("is_shared", object.GetIsShared())
	d.Set("external_ipam", object.GetExternalIpam())
	d.Set("flood_unknown_unicast", object.GetFloodUnknownUnicast())
	d.Set("multi_policy_service_chains_enabled", object.GetMultiPolicyServiceChainsEnabled())
	d.Set("address_allocation_mode", object.GetAddressAllocationMode())
	d.Set("mac_learning_enabled", object.GetMacLearningEnabled())
	mac_limit_controlObj := object.GetMacLimitControl()
	d.Set("mac_limit_control", TakeMACLimitControlTypeAsMap(&mac_limit_controlObj))
	mac_move_controlObj := object.GetMacMoveControl()
	d.Set("mac_move_control", TakeMACMoveLimitControlTypeAsMap(&mac_move_controlObj))
	d.Set("mac_aging_time", object.GetMacAgingTime())
	d.Set("pbb_evpn_enable", object.GetPbbEvpnEnable())
	d.Set("pbb_etree_enable", object.GetPbbEtreeEnable())
	d.Set("layer2_control_word", object.GetLayer2ControlWord())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteVirtualNetworkRefsToResource(object VirtualNetwork, d *schema.ResourceData, m interface{}) {

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
	if ref, err := object.GetNetworkIpamRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("network_ipam_refs", refList)
	}
	if ref, err := object.GetNetworkPolicyRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("network_policy_refs", refList)
	}
	if ref, err := object.GetVirtualNetworkRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_network_refs", refList)
	}
	if ref, err := object.GetRouteTableRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("route_table_refs", refList)
	}
	if ref, err := object.GetBgpvpnRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("bgpvpn_refs", refList)
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

func TakeVirtualNetworkAsMap(object *VirtualNetwork) map[string]interface{} {
	omap := make(map[string]interface{})

	ecmp_hashing_include_fieldsObj := object.GetEcmpHashingIncludeFields()
	omap["ecmp_hashing_include_fields"] = TakeEcmpHashingIncludeFieldsAsMap(&ecmp_hashing_include_fieldsObj)
	virtual_network_propertiesObj := object.GetVirtualNetworkProperties()
	omap["virtual_network_properties"] = TakeVirtualNetworkTypeAsMap(&virtual_network_propertiesObj)
	provider_propertiesObj := object.GetProviderProperties()
	omap["provider_properties"] = TakeProviderDetailsAsMap(&provider_propertiesObj)
	omap["is_provider_network"] = object.GetIsProviderNetwork()
	omap["virtual_network_network_id"] = object.GetVirtualNetworkNetworkId()
	omap["port_security_enabled"] = object.GetPortSecurityEnabled()
	route_target_listObj := object.GetRouteTargetList()
	omap["route_target_list"] = TakeRouteTargetListAsMap(&route_target_listObj)
	import_route_target_listObj := object.GetImportRouteTargetList()
	omap["import_route_target_list"] = TakeRouteTargetListAsMap(&import_route_target_listObj)
	export_route_target_listObj := object.GetExportRouteTargetList()
	omap["export_route_target_list"] = TakeRouteTargetListAsMap(&export_route_target_listObj)
	omap["router_external"] = object.GetRouterExternal()
	omap["is_shared"] = object.GetIsShared()
	omap["external_ipam"] = object.GetExternalIpam()
	omap["flood_unknown_unicast"] = object.GetFloodUnknownUnicast()
	omap["multi_policy_service_chains_enabled"] = object.GetMultiPolicyServiceChainsEnabled()
	omap["address_allocation_mode"] = object.GetAddressAllocationMode()
	omap["mac_learning_enabled"] = object.GetMacLearningEnabled()
	mac_limit_controlObj := object.GetMacLimitControl()
	omap["mac_limit_control"] = TakeMACLimitControlTypeAsMap(&mac_limit_controlObj)
	mac_move_controlObj := object.GetMacMoveControl()
	omap["mac_move_control"] = TakeMACMoveLimitControlTypeAsMap(&mac_move_controlObj)
	omap["mac_aging_time"] = object.GetMacAgingTime()
	omap["pbb_evpn_enable"] = object.GetPbbEvpnEnable()
	omap["pbb_etree_enable"] = object.GetPbbEtreeEnable()
	omap["layer2_control_word"] = object.GetLayer2ControlWord()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateVirtualNetworkFromResource(object *VirtualNetwork, d *schema.ResourceData, m interface{}, prefix ...string) {
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
	if d.HasChange("virtual_network_properties") {
		if val, ok := d.GetOk("virtual_network_properties"); ok {
			member := new(VirtualNetworkType)
			SetVirtualNetworkTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualNetworkProperties(member)
		}
	}
	if d.HasChange("provider_properties") {
		if val, ok := d.GetOk("provider_properties"); ok {
			member := new(ProviderDetails)
			SetProviderDetailsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetProviderProperties(member)
		}
	}
	if d.HasChange("is_provider_network") {
		if val, ok := d.GetOk("is_provider_network"); ok {
			object.SetIsProviderNetwork(val.(bool))
		}
	}
	if d.HasChange("virtual_network_network_id") {
		if val, ok := d.GetOk("virtual_network_network_id"); ok {
			object.SetVirtualNetworkNetworkId(val.(int))
		}
	}
	if d.HasChange("port_security_enabled") {
		if val, ok := d.GetOk("port_security_enabled"); ok {
			object.SetPortSecurityEnabled(val.(bool))
		}
	}
	if d.HasChange("route_target_list") {
		if val, ok := d.GetOk("route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetRouteTargetList(member)
		}
	}
	if d.HasChange("import_route_target_list") {
		if val, ok := d.GetOk("import_route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetImportRouteTargetList(member)
		}
	}
	if d.HasChange("export_route_target_list") {
		if val, ok := d.GetOk("export_route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetExportRouteTargetList(member)
		}
	}
	if d.HasChange("router_external") {
		if val, ok := d.GetOk("router_external"); ok {
			object.SetRouterExternal(val.(bool))
		}
	}
	if d.HasChange("is_shared") {
		if val, ok := d.GetOk("is_shared"); ok {
			object.SetIsShared(val.(bool))
		}
	}
	if d.HasChange("external_ipam") {
		if val, ok := d.GetOk("external_ipam"); ok {
			object.SetExternalIpam(val.(bool))
		}
	}
	if d.HasChange("flood_unknown_unicast") {
		if val, ok := d.GetOk("flood_unknown_unicast"); ok {
			object.SetFloodUnknownUnicast(val.(bool))
		}
	}
	if d.HasChange("multi_policy_service_chains_enabled") {
		if val, ok := d.GetOk("multi_policy_service_chains_enabled"); ok {
			object.SetMultiPolicyServiceChainsEnabled(val.(bool))
		}
	}
	if d.HasChange("address_allocation_mode") {
		if val, ok := d.GetOk("address_allocation_mode"); ok {
			object.SetAddressAllocationMode(val.(string))
		}
	}
	if d.HasChange("mac_learning_enabled") {
		if val, ok := d.GetOk("mac_learning_enabled"); ok {
			object.SetMacLearningEnabled(val.(bool))
		}
	}
	if d.HasChange("mac_limit_control") {
		if val, ok := d.GetOk("mac_limit_control"); ok {
			member := new(MACLimitControlType)
			SetMACLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMacLimitControl(member)
		}
	}
	if d.HasChange("mac_move_control") {
		if val, ok := d.GetOk("mac_move_control"); ok {
			member := new(MACMoveLimitControlType)
			SetMACMoveLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMacMoveControl(member)
		}
	}
	if d.HasChange("mac_aging_time") {
		if val, ok := d.GetOk("mac_aging_time"); ok {
			object.SetMacAgingTime(val.(int))
		}
	}
	if d.HasChange("pbb_evpn_enable") {
		if val, ok := d.GetOk("pbb_evpn_enable"); ok {
			object.SetPbbEvpnEnable(val.(bool))
		}
	}
	if d.HasChange("pbb_etree_enable") {
		if val, ok := d.GetOk("pbb_etree_enable"); ok {
			object.SetPbbEtreeEnable(val.(bool))
		}
	}
	if d.HasChange("layer2_control_word") {
		if val, ok := d.GetOk("layer2_control_word"); ok {
			object.SetLayer2ControlWord(val.(bool))
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

}

func UpdateVirtualNetworkRefsFromResource(object *VirtualNetwork, d *schema.ResourceData, m interface{}, prefix ...string) {
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
	if d.HasChange("network_ipam_refs") {
		object.ClearNetworkIpam()
		if val, ok := d.GetOk("network_ipam_refs"); ok {
			log.Printf("Got ref network_ipam_refs -- will call: object.AddNetworkIpam(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("network-ipam", refId.(string))
				dataObj := new(VnSubnetsType)
				SetVnSubnetsTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddNetworkIpam(refObj.(*NetworkIpam), *dataObj)
			}
		}
	}
	if d.HasChange("network_policy_refs") {
		object.ClearNetworkPolicy()
		if val, ok := d.GetOk("network_policy_refs"); ok {
			log.Printf("Got ref network_policy_refs -- will call: object.AddNetworkPolicy(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("network-policy", refId.(string))
				dataObj := new(VirtualNetworkPolicyType)
				SetVirtualNetworkPolicyTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddNetworkPolicy(refObj.(*NetworkPolicy), *dataObj)
			}
		}
	}
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
	if d.HasChange("route_table_refs") {
		object.ClearRouteTable()
		if val, ok := d.GetOk("route_table_refs"); ok {
			log.Printf("Got ref route_table_refs -- will call: object.AddRouteTable(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("route-table", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddRouteTable(refObj.(*RouteTable))
			}
		}
	}
	if d.HasChange("bgpvpn_refs") {
		object.ClearBgpvpn()
		if val, ok := d.GetOk("bgpvpn_refs"); ok {
			log.Printf("Got ref bgpvpn_refs -- will call: object.AddBgpvpn(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("bgpvpn", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddBgpvpn(refObj.(*Bgpvpn))
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

func ResourceVirtualNetworkCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualNetworkCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(VirtualNetwork)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceVirtualNetworkCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "VirtualNetwork", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceVirtualNetworkCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "VirtualNetwork", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetVirtualNetworkFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkCreate] Creation of resource VirtualNetwork on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceVirtualNetworkRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualNetworkRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualNetworkRefsCreate] Missing 'uuid' field for resource VirtualNetwork")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-network", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsCreate] Retrieving VirtualNetwork with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualNetwork := obj.(*VirtualNetwork) // Fully set by Contrail backend
	if err := SetRefsVirtualNetworkFromResource(objVirtualNetwork, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsCreate] Set refs on object VirtualNetwork (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualNetwork.GetHref())
	if err := client.Update(objVirtualNetwork); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsCreate] Update refs for resource VirtualNetwork (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualNetwork.GetUuid())
	return nil
}

func ResourceVirtualNetworkRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualNetworkRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-network", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRead] Read resource virtual-network on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualNetwork)
	WriteVirtualNetworkToResource(*object, d, m)
	return nil
}

func ResourceVirtualNetworkRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualNetworkRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-network", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsRead] Read resource virtual-network on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualNetwork)
	WriteVirtualNetworkRefsToResource(*object, d, m)
	return nil
}

func ResourceVirtualNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualNetworkUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-network", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkUpdate] Retrieving VirtualNetwork with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualNetwork)
	UpdateVirtualNetworkFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkUpdate] Update of resource VirtualNetwork on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualNetworkRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualNetworkRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-network", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsUpdate] Retrieving VirtualNetwork with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualNetwork)
	UpdateVirtualNetworkRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsUpdate] Update of resource VirtualNetwork on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualNetworkDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualNetworkDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("virtual-network", d.Id()); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkDelete] Deletion of resource virtual-network on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceVirtualNetworkRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualNetworkRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualNetworkRefsDelete] Missing 'uuid' field for resource VirtualNetwork")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-network", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsDelete] Retrieving VirtualNetwork with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualNetwork := obj.(*VirtualNetwork) // Fully set by Contrail backend
	if err := DeleteRefsVirtualNetworkFromResource(objVirtualNetwork, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsDelete] Set refs on object VirtualNetwork (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualNetwork.GetHref())
	if err := client.Update(objVirtualNetwork); err != nil {
		return fmt.Errorf("[ResourceVirtualNetworkRefsDelete] Delete refs for resource VirtualNetwork (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualNetwork.GetUuid())
	return nil
}

func ResourceVirtualNetworkSchema() map[string]*schema.Schema {
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
		"virtual_network_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualNetworkType(),
		},
		"provider_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceProviderDetails(),
		},
		"is_provider_network": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"virtual_network_network_id": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"port_security_enabled": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
		},
		"import_route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
		},
		"export_route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
		},
		"router_external": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"is_shared": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"external_ipam": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"flood_unknown_unicast": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"multi_policy_service_chains_enabled": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"address_allocation_mode": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"mac_learning_enabled": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"mac_limit_control": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMACLimitControlType(),
		},
		"mac_move_control": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMACMoveLimitControlType(),
		},
		"mac_aging_time": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"pbb_evpn_enable": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"pbb_etree_enable": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"layer2_control_word": &schema.Schema{
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
	}
}

func ResourceVirtualNetworkRefsSchema() map[string]*schema.Schema {
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
		"network_ipam_refs": &schema.Schema{
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
						Elem:     ResourceVnSubnetsType(),
						Required: true,
					},
				},
			},
		},
		"network_policy_refs": &schema.Schema{
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
						Elem:     ResourceVirtualNetworkPolicyType(),
						Required: true,
					},
				},
			},
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
		"route_table_refs": &schema.Schema{
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
		"bgpvpn_refs": &schema.Schema{
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

func ResourceVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualNetworkCreate,
		Read:   ResourceVirtualNetworkRead,
		Update: ResourceVirtualNetworkUpdate,
		Delete: ResourceVirtualNetworkDelete,
		Schema: ResourceVirtualNetworkSchema(),
	}
}

func ResourceVirtualNetworkRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualNetworkRefsCreate,
		Read:   ResourceVirtualNetworkRefsRead,
		Update: ResourceVirtualNetworkRefsUpdate,
		Delete: ResourceVirtualNetworkRefsDelete,
		Schema: ResourceVirtualNetworkRefsSchema(),
	}
}

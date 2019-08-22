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

func SetGlobalSystemConfigFromResource(object *GlobalSystemConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetGlobalSystemConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("autonomous_system"); ok {
		object.SetAutonomousSystem(val.(int))
	}
	if val, ok := d.GetOk("config_version"); ok {
		object.SetConfigVersion(val.(string))
	}
	if val, ok := d.GetOk("graceful_restart_parameters"); ok {
		member := new(GracefulRestartParametersType)
		SetGracefulRestartParametersTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetGracefulRestartParameters(member)
	}
	if val, ok := d.GetOk("plugin_tuning"); ok {
		member := new(PluginProperties)
		SetPluginPropertiesFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetPluginTuning(member)
	}
	if val, ok := d.GetOk("ibgp_auto_mesh"); ok {
		object.SetIbgpAutoMesh(val.(bool))
	}
	if val, ok := d.GetOk("bgp_always_compare_med"); ok {
		object.SetBgpAlwaysCompareMed(val.(bool))
	}
	if val, ok := d.GetOk("ip_fabric_subnets"); ok {
		member := new(SubnetListType)
		SetSubnetListTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetIpFabricSubnets(member)
	}
	if val, ok := d.GetOk("bgpaas_parameters"); ok {
		member := new(BGPaaServiceParametersType)
		SetBGPaaServiceParametersTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetBgpaasParameters(member)
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
	if val, ok := d.GetOk("alarm_enable"); ok {
		object.SetAlarmEnable(val.(bool))
	}
	if val, ok := d.GetOk("user_defined_log_statistics"); ok {
		member := new(UserDefinedLogStatList)
		SetUserDefinedLogStatListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetUserDefinedLogStatistics(member)
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

func SetRefsGlobalSystemConfigFromResource(object *GlobalSystemConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsGlobalSystemConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func SetReqRefsGlobalSystemConfigFromResource(object *GlobalSystemConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsGlobalSystemConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsGlobalSystemConfigFromResource(object *GlobalSystemConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsGlobalSystemConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteGlobalSystemConfigToResource(object GlobalSystemConfig, d *schema.ResourceData, m interface{}) {

	d.Set("autonomous_system", object.GetAutonomousSystem())
	d.Set("config_version", object.GetConfigVersion())
	graceful_restart_parametersObj := object.GetGracefulRestartParameters()
	d.Set("graceful_restart_parameters", TakeGracefulRestartParametersTypeAsMap(&graceful_restart_parametersObj))
	plugin_tuningObj := object.GetPluginTuning()
	d.Set("plugin_tuning", TakePluginPropertiesAsMap(&plugin_tuningObj))
	d.Set("ibgp_auto_mesh", object.GetIbgpAutoMesh())
	d.Set("bgp_always_compare_med", object.GetBgpAlwaysCompareMed())
	ip_fabric_subnetsObj := object.GetIpFabricSubnets()
	d.Set("ip_fabric_subnets", TakeSubnetListTypeAsMap(&ip_fabric_subnetsObj))
	bgpaas_parametersObj := object.GetBgpaasParameters()
	d.Set("bgpaas_parameters", TakeBGPaaServiceParametersTypeAsMap(&bgpaas_parametersObj))
	mac_limit_controlObj := object.GetMacLimitControl()
	d.Set("mac_limit_control", TakeMACLimitControlTypeAsMap(&mac_limit_controlObj))
	mac_move_controlObj := object.GetMacMoveControl()
	d.Set("mac_move_control", TakeMACMoveLimitControlTypeAsMap(&mac_move_controlObj))
	d.Set("mac_aging_time", object.GetMacAgingTime())
	d.Set("alarm_enable", object.GetAlarmEnable())
	user_defined_log_statisticsObj := object.GetUserDefinedLogStatistics()
	d.Set("user_defined_log_statistics", TakeUserDefinedLogStatListAsMap(&user_defined_log_statisticsObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

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

func TakeGlobalSystemConfigAsMap(object *GlobalSystemConfig) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["autonomous_system"] = object.GetAutonomousSystem()
	omap["config_version"] = object.GetConfigVersion()
	graceful_restart_parametersObj := object.GetGracefulRestartParameters()
	omap["graceful_restart_parameters"] = TakeGracefulRestartParametersTypeAsMap(&graceful_restart_parametersObj)
	plugin_tuningObj := object.GetPluginTuning()
	omap["plugin_tuning"] = TakePluginPropertiesAsMap(&plugin_tuningObj)
	omap["ibgp_auto_mesh"] = object.GetIbgpAutoMesh()
	omap["bgp_always_compare_med"] = object.GetBgpAlwaysCompareMed()
	ip_fabric_subnetsObj := object.GetIpFabricSubnets()
	omap["ip_fabric_subnets"] = TakeSubnetListTypeAsMap(&ip_fabric_subnetsObj)
	bgpaas_parametersObj := object.GetBgpaasParameters()
	omap["bgpaas_parameters"] = TakeBGPaaServiceParametersTypeAsMap(&bgpaas_parametersObj)
	mac_limit_controlObj := object.GetMacLimitControl()
	omap["mac_limit_control"] = TakeMACLimitControlTypeAsMap(&mac_limit_controlObj)
	mac_move_controlObj := object.GetMacMoveControl()
	omap["mac_move_control"] = TakeMACMoveLimitControlTypeAsMap(&mac_move_controlObj)
	omap["mac_aging_time"] = object.GetMacAgingTime()
	omap["alarm_enable"] = object.GetAlarmEnable()
	user_defined_log_statisticsObj := object.GetUserDefinedLogStatistics()
	omap["user_defined_log_statistics"] = TakeUserDefinedLogStatListAsMap(&user_defined_log_statisticsObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateGlobalSystemConfigFromResource(object *GlobalSystemConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("autonomous_system") {
		if val, ok := d.GetOk("autonomous_system"); ok {
			object.SetAutonomousSystem(val.(int))
		}
	}
	if d.HasChange("config_version") {
		if val, ok := d.GetOk("config_version"); ok {
			object.SetConfigVersion(val.(string))
		}
	}
	if d.HasChange("graceful_restart_parameters") {
		if val, ok := d.GetOk("graceful_restart_parameters"); ok {
			member := new(GracefulRestartParametersType)
			SetGracefulRestartParametersTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetGracefulRestartParameters(member)
		}
	}
	if d.HasChange("plugin_tuning") {
		if val, ok := d.GetOk("plugin_tuning"); ok {
			member := new(PluginProperties)
			SetPluginPropertiesFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetPluginTuning(member)
		}
	}
	if d.HasChange("ibgp_auto_mesh") {
		if val, ok := d.GetOk("ibgp_auto_mesh"); ok {
			object.SetIbgpAutoMesh(val.(bool))
		}
	}
	if d.HasChange("bgp_always_compare_med") {
		if val, ok := d.GetOk("bgp_always_compare_med"); ok {
			object.SetBgpAlwaysCompareMed(val.(bool))
		}
	}
	if d.HasChange("ip_fabric_subnets") {
		if val, ok := d.GetOk("ip_fabric_subnets"); ok {
			member := new(SubnetListType)
			SetSubnetListTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetIpFabricSubnets(member)
		}
	}
	if d.HasChange("bgpaas_parameters") {
		if val, ok := d.GetOk("bgpaas_parameters"); ok {
			member := new(BGPaaServiceParametersType)
			SetBGPaaServiceParametersTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetBgpaasParameters(member)
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
	if d.HasChange("alarm_enable") {
		if val, ok := d.GetOk("alarm_enable"); ok {
			object.SetAlarmEnable(val.(bool))
		}
	}
	if d.HasChange("user_defined_log_statistics") {
		if val, ok := d.GetOk("user_defined_log_statistics"); ok {
			member := new(UserDefinedLogStatList)
			SetUserDefinedLogStatListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetUserDefinedLogStatistics(member)
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
	if d.HasChange("bgp_router_refs") {
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

func ResourceGlobalSystemConfigCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalSystemConfigCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(GlobalSystemConfig)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceGlobalSystemConfigCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "GlobalSystemConfig", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetGlobalSystemConfigFromResource(object, d, m)

	if err := SetReqRefsGlobalSystemConfigFromResource(object, d, m); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigReqRefsCreate] Set required refs on object GlobalSystemConfig on %v (%v)", client.GetServer(), err)
	}

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigCreate] Creation of resource GlobalSystemConfig on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceGlobalSystemConfigRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalSystemConfigRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsCreate] Missing 'uuid' field for resource GlobalSystemConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("global-system-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsCreate] Retrieving GlobalSystemConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objGlobalSystemConfig := obj.(*GlobalSystemConfig) // Fully set by Contrail backend
	if err := SetRefsGlobalSystemConfigFromResource(objGlobalSystemConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsCreate] Set refs on object GlobalSystemConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objGlobalSystemConfig.GetHref())
	if err := client.Update(objGlobalSystemConfig); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsCreate] Update refs for resource GlobalSystemConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objGlobalSystemConfig.GetUuid())
	return nil
}

func ResourceGlobalSystemConfigRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalSystemConfigREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("global-system-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRead] Read resource global-system-config on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*GlobalSystemConfig)
	WriteGlobalSystemConfigToResource(*object, d, m)
	return nil
}

func ResourceGlobalSystemConfigRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalSystemConfigRefsREAD")
	return nil
}

func ResourceGlobalSystemConfigUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalSystemConfigUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("global-system-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigResourceUpdate] Retrieving GlobalSystemConfig with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*GlobalSystemConfig)
	UpdateGlobalSystemConfigFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigUpdate] Update of resource GlobalSystemConfig on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceGlobalSystemConfigRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalSystemConfigRefsUpdate")
	return nil
}

func ResourceGlobalSystemConfigDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalSystemConfigDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("global-system-config", d.Id()); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigDelete] Deletion of resource global-system-config on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceGlobalSystemConfigRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalSystemConfigRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsDelete] Missing 'uuid' field for resource GlobalSystemConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("global-system-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsDelete] Retrieving GlobalSystemConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objGlobalSystemConfig := obj.(*GlobalSystemConfig) // Fully set by Contrail backend
	if err := DeleteRefsGlobalSystemConfigFromResource(objGlobalSystemConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsDelete] Set refs on object GlobalSystemConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objGlobalSystemConfig.GetHref())
	if err := client.Update(objGlobalSystemConfig); err != nil {
		return fmt.Errorf("[ResourceGlobalSystemConfigRefsDelete] Delete refs for resource GlobalSystemConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objGlobalSystemConfig.GetUuid())
	return nil
}

func ResourceGlobalSystemConfigSchema() map[string]*schema.Schema {
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
		"autonomous_system": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"config_version": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"graceful_restart_parameters": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceGracefulRestartParametersType(),
		},
		"plugin_tuning": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePluginProperties(),
		},
		"ibgp_auto_mesh": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"bgp_always_compare_med": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"ip_fabric_subnets": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetListType(),
		},
		"bgpaas_parameters": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceBGPaaServiceParametersType(),
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
		"alarm_enable": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"user_defined_log_statistics": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceUserDefinedLogStatList(),
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

func ResourceGlobalSystemConfigRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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

func ResourceGlobalSystemConfig() *schema.Resource {
	return &schema.Resource{
		Create: ResourceGlobalSystemConfigCreate,
		Read:   ResourceGlobalSystemConfigRead,
		Update: ResourceGlobalSystemConfigUpdate,
		Delete: ResourceGlobalSystemConfigDelete,
		Schema: ResourceGlobalSystemConfigSchema(),
	}
}

func ResourceGlobalSystemConfigRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceGlobalSystemConfigRefsCreate,
		Read:   ResourceGlobalSystemConfigRefsRead,
		Update: ResourceGlobalSystemConfigRefsUpdate,
		Delete: ResourceGlobalSystemConfigRefsDelete,
		Schema: ResourceGlobalSystemConfigRefsSchema(),
	}
}

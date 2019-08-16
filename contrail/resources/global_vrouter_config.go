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

func SetGlobalVrouterConfigFromResource(object *GlobalVrouterConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetGlobalVrouterConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("ecmp_hashing_include_fields"); ok {
		member := new(EcmpHashingIncludeFields)
		SetEcmpHashingIncludeFieldsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetEcmpHashingIncludeFields(member)
	}
	if val, ok := d.GetOk("linklocal_services"); ok {
		member := new(LinklocalServicesTypes)
		SetLinklocalServicesTypesFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLinklocalServices(member)
	}
	if val, ok := d.GetOk("encapsulation_priorities"); ok {
		member := new(EncapsulationPrioritiesType)
		SetEncapsulationPrioritiesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetEncapsulationPriorities(member)
	}
	if val, ok := d.GetOk("vxlan_network_identifier_mode"); ok {
		object.SetVxlanNetworkIdentifierMode(val.(string))
	}
	if val, ok := d.GetOk("flow_export_rate"); ok {
		object.SetFlowExportRate(val.(int))
	}
	if val, ok := d.GetOk("flow_aging_timeout_list"); ok {
		member := new(FlowAgingTimeoutList)
		SetFlowAgingTimeoutListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetFlowAgingTimeoutList(member)
	}
	if val, ok := d.GetOk("enable_security_logging"); ok {
		object.SetEnableSecurityLogging(val.(bool))
	}
	if val, ok := d.GetOk("forwarding_mode"); ok {
		object.SetForwardingMode(val.(string))
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

func SetRefsGlobalVrouterConfigFromResource(object *GlobalVrouterConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsGlobalVrouterConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsGlobalVrouterConfigFromResource(object *GlobalVrouterConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsGlobalVrouterConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteGlobalVrouterConfigToResource(object GlobalVrouterConfig, d *schema.ResourceData, m interface{}) {

	ecmp_hashing_include_fieldsObj := object.GetEcmpHashingIncludeFields()
	d.Set("ecmp_hashing_include_fields", TakeEcmpHashingIncludeFieldsAsMap(&ecmp_hashing_include_fieldsObj))
	linklocal_servicesObj := object.GetLinklocalServices()
	d.Set("linklocal_services", TakeLinklocalServicesTypesAsMap(&linklocal_servicesObj))
	encapsulation_prioritiesObj := object.GetEncapsulationPriorities()
	d.Set("encapsulation_priorities", TakeEncapsulationPrioritiesTypeAsMap(&encapsulation_prioritiesObj))
	d.Set("vxlan_network_identifier_mode", object.GetVxlanNetworkIdentifierMode())
	d.Set("flow_export_rate", object.GetFlowExportRate())
	flow_aging_timeout_listObj := object.GetFlowAgingTimeoutList()
	d.Set("flow_aging_timeout_list", TakeFlowAgingTimeoutListAsMap(&flow_aging_timeout_listObj))
	d.Set("enable_security_logging", object.GetEnableSecurityLogging())
	d.Set("forwarding_mode", object.GetForwardingMode())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeGlobalVrouterConfigAsMap(object *GlobalVrouterConfig) map[string]interface{} {
	omap := make(map[string]interface{})

	ecmp_hashing_include_fieldsObj := object.GetEcmpHashingIncludeFields()
	omap["ecmp_hashing_include_fields"] = TakeEcmpHashingIncludeFieldsAsMap(&ecmp_hashing_include_fieldsObj)
	linklocal_servicesObj := object.GetLinklocalServices()
	omap["linklocal_services"] = TakeLinklocalServicesTypesAsMap(&linklocal_servicesObj)
	encapsulation_prioritiesObj := object.GetEncapsulationPriorities()
	omap["encapsulation_priorities"] = TakeEncapsulationPrioritiesTypeAsMap(&encapsulation_prioritiesObj)
	omap["vxlan_network_identifier_mode"] = object.GetVxlanNetworkIdentifierMode()
	omap["flow_export_rate"] = object.GetFlowExportRate()
	flow_aging_timeout_listObj := object.GetFlowAgingTimeoutList()
	omap["flow_aging_timeout_list"] = TakeFlowAgingTimeoutListAsMap(&flow_aging_timeout_listObj)
	omap["enable_security_logging"] = object.GetEnableSecurityLogging()
	omap["forwarding_mode"] = object.GetForwardingMode()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateGlobalVrouterConfigFromResource(object *GlobalVrouterConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
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
	if d.HasChange("linklocal_services") {
		if val, ok := d.GetOk("linklocal_services"); ok {
			member := new(LinklocalServicesTypes)
			SetLinklocalServicesTypesFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLinklocalServices(member)
		}
	}
	if d.HasChange("encapsulation_priorities") {
		if val, ok := d.GetOk("encapsulation_priorities"); ok {
			member := new(EncapsulationPrioritiesType)
			SetEncapsulationPrioritiesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetEncapsulationPriorities(member)
		}
	}
	if d.HasChange("vxlan_network_identifier_mode") {
		if val, ok := d.GetOk("vxlan_network_identifier_mode"); ok {
			object.SetVxlanNetworkIdentifierMode(val.(string))
		}
	}
	if d.HasChange("flow_export_rate") {
		if val, ok := d.GetOk("flow_export_rate"); ok {
			object.SetFlowExportRate(val.(int))
		}
	}
	if d.HasChange("flow_aging_timeout_list") {
		if val, ok := d.GetOk("flow_aging_timeout_list"); ok {
			member := new(FlowAgingTimeoutList)
			SetFlowAgingTimeoutListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetFlowAgingTimeoutList(member)
		}
	}
	if d.HasChange("enable_security_logging") {
		if val, ok := d.GetOk("enable_security_logging"); ok {
			object.SetEnableSecurityLogging(val.(bool))
		}
	}
	if d.HasChange("forwarding_mode") {
		if val, ok := d.GetOk("forwarding_mode"); ok {
			object.SetForwardingMode(val.(string))
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

func ResourceGlobalVrouterConfigCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalVrouterConfigCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(GlobalVrouterConfig)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceGlobalVrouterConfigCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "GlobalVrouterConfig", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetGlobalVrouterConfigFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigCreate] Creation of resource GlobalVrouterConfig on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceGlobalVrouterConfigRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalVrouterConfigRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsCreate] Missing 'uuid' field for resource GlobalVrouterConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("global-vrouter-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsCreate] Retrieving GlobalVrouterConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objGlobalVrouterConfig := obj.(*GlobalVrouterConfig) // Fully set by Contrail backend
	if err := SetRefsGlobalVrouterConfigFromResource(objGlobalVrouterConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsCreate] Set refs on object GlobalVrouterConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objGlobalVrouterConfig.GetHref())
	if err := client.Update(objGlobalVrouterConfig); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsCreate] Update refs for resource GlobalVrouterConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objGlobalVrouterConfig.GetUuid())
	return nil
}

func ResourceGlobalVrouterConfigRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalVrouterConfigREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("global-vrouter-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRead] Read resource global-vrouter-config on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*GlobalVrouterConfig)
	WriteGlobalVrouterConfigToResource(*object, d, m)
	return nil
}

func ResourceGlobalVrouterConfigRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalVrouterConfigRefsREAD")
	return nil
}

func ResourceGlobalVrouterConfigUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalVrouterConfigUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("global-vrouter-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigResourceUpdate] Retrieving GlobalVrouterConfig with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*GlobalVrouterConfig)
	UpdateGlobalVrouterConfigFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigUpdate] Update of resource GlobalVrouterConfig on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceGlobalVrouterConfigRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalVrouterConfigRefsUpdate")
	return nil
}

func ResourceGlobalVrouterConfigDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalVrouterConfigDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("global-vrouter-config", d.Id()); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigDelete] Deletion of resource global-vrouter-config on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceGlobalVrouterConfigRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalVrouterConfigRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsDelete] Missing 'uuid' field for resource GlobalVrouterConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("global-vrouter-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsDelete] Retrieving GlobalVrouterConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objGlobalVrouterConfig := obj.(*GlobalVrouterConfig) // Fully set by Contrail backend
	if err := DeleteRefsGlobalVrouterConfigFromResource(objGlobalVrouterConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsDelete] Set refs on object GlobalVrouterConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objGlobalVrouterConfig.GetHref())
	if err := client.Update(objGlobalVrouterConfig); err != nil {
		return fmt.Errorf("[ResourceGlobalVrouterConfigRefsDelete] Delete refs for resource GlobalVrouterConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objGlobalVrouterConfig.GetUuid())
	return nil
}

func ResourceGlobalVrouterConfigSchema() map[string]*schema.Schema {
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
		"linklocal_services": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLinklocalServicesTypes(),
		},
		"encapsulation_priorities": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceEncapsulationPrioritiesType(),
		},
		"vxlan_network_identifier_mode": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"flow_export_rate": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"flow_aging_timeout_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFlowAgingTimeoutList(),
		},
		"enable_security_logging": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"forwarding_mode": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
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

func ResourceGlobalVrouterConfigRefsSchema() map[string]*schema.Schema {
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

func ResourceGlobalVrouterConfig() *schema.Resource {
	return &schema.Resource{
		Create: ResourceGlobalVrouterConfigCreate,
		Read:   ResourceGlobalVrouterConfigRead,
		Update: ResourceGlobalVrouterConfigUpdate,
		Delete: ResourceGlobalVrouterConfigDelete,
		Schema: ResourceGlobalVrouterConfigSchema(),
	}
}

func ResourceGlobalVrouterConfigRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceGlobalVrouterConfigRefsCreate,
		Read:   ResourceGlobalVrouterConfigRefsRead,
		Update: ResourceGlobalVrouterConfigRefsUpdate,
		Delete: ResourceGlobalVrouterConfigRefsDelete,
		Schema: ResourceGlobalVrouterConfigRefsSchema(),
	}
}

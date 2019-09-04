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

func SetQosConfigFromResource(object *QosConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetQosConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("qos_config_type"); ok {
		object.SetQosConfigType(val.(string))
	}
	if val, ok := d.GetOk("dscp_entries"); ok {
		member := new(QosIdForwardingClassPairs)
		SetQosIdForwardingClassPairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetDscpEntries(member)
	}
	if val, ok := d.GetOk("vlan_priority_entries"); ok {
		member := new(QosIdForwardingClassPairs)
		SetQosIdForwardingClassPairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVlanPriorityEntries(member)
	}
	if val, ok := d.GetOk("mpls_exp_entries"); ok {
		member := new(QosIdForwardingClassPairs)
		SetQosIdForwardingClassPairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMplsExpEntries(member)
	}
	if val, ok := d.GetOk("default_forwarding_class_id"); ok {
		object.SetDefaultForwardingClassId(val.(int))
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

func SetRefsQosConfigFromResource(object *QosConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsQosConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("global_system_config_refs"); ok {
		log.Printf("Got ref global_system_config_refs -- will call: object.AddGlobalSystemConfig(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("global-system-config", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving global-system-config by Uuid = %v as ref for GlobalSystemConfig on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddGlobalSystemConfig(refObj.(*GlobalSystemConfig))
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

func DeleteRefsQosConfigFromResource(object *QosConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsQosConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("global_system_config_refs"); ok {
		log.Printf("Got ref global_system_config_refs -- will call: object.DeleteGlobalSystemConfig(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteGlobalSystemConfig(refId.(string))
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

func WriteQosConfigToResource(object QosConfig, d *schema.ResourceData, m interface{}) {

	d.Set("qos_config_type", object.GetQosConfigType())
	dscp_entriesObj := object.GetDscpEntries()
	d.Set("dscp_entries", TakeQosIdForwardingClassPairsAsMap(&dscp_entriesObj))
	vlan_priority_entriesObj := object.GetVlanPriorityEntries()
	d.Set("vlan_priority_entries", TakeQosIdForwardingClassPairsAsMap(&vlan_priority_entriesObj))
	mpls_exp_entriesObj := object.GetMplsExpEntries()
	d.Set("mpls_exp_entries", TakeQosIdForwardingClassPairsAsMap(&mpls_exp_entriesObj))
	d.Set("default_forwarding_class_id", object.GetDefaultForwardingClassId())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteQosConfigRefsToResource(object QosConfig, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetGlobalSystemConfigRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("global_system_config_refs", refList)
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

func TakeQosConfigAsMap(object *QosConfig) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["qos_config_type"] = object.GetQosConfigType()
	dscp_entriesObj := object.GetDscpEntries()
	omap["dscp_entries"] = TakeQosIdForwardingClassPairsAsMap(&dscp_entriesObj)
	vlan_priority_entriesObj := object.GetVlanPriorityEntries()
	omap["vlan_priority_entries"] = TakeQosIdForwardingClassPairsAsMap(&vlan_priority_entriesObj)
	mpls_exp_entriesObj := object.GetMplsExpEntries()
	omap["mpls_exp_entries"] = TakeQosIdForwardingClassPairsAsMap(&mpls_exp_entriesObj)
	omap["default_forwarding_class_id"] = object.GetDefaultForwardingClassId()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateQosConfigFromResource(object *QosConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("qos_config_type") {
		if val, ok := d.GetOk("qos_config_type"); ok {
			object.SetQosConfigType(val.(string))
		}
	}
	if d.HasChange("dscp_entries") {
		if val, ok := d.GetOk("dscp_entries"); ok {
			member := new(QosIdForwardingClassPairs)
			SetQosIdForwardingClassPairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetDscpEntries(member)
		}
	}
	if d.HasChange("vlan_priority_entries") {
		if val, ok := d.GetOk("vlan_priority_entries"); ok {
			member := new(QosIdForwardingClassPairs)
			SetQosIdForwardingClassPairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVlanPriorityEntries(member)
		}
	}
	if d.HasChange("mpls_exp_entries") {
		if val, ok := d.GetOk("mpls_exp_entries"); ok {
			member := new(QosIdForwardingClassPairs)
			SetQosIdForwardingClassPairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMplsExpEntries(member)
		}
	}
	if d.HasChange("default_forwarding_class_id") {
		if val, ok := d.GetOk("default_forwarding_class_id"); ok {
			object.SetDefaultForwardingClassId(val.(int))
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

func UpdateQosConfigRefsFromResource(object *QosConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("global_system_config_refs") {
		object.ClearGlobalSystemConfig()
		if val, ok := d.GetOk("global_system_config_refs"); ok {
			log.Printf("Got ref global_system_config_refs -- will call: object.AddGlobalSystemConfig(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("global-system-config", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddGlobalSystemConfig(refObj.(*GlobalSystemConfig))
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

func ResourceQosConfigCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceQosConfigCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(QosConfig)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceQosConfigCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "QosConfig", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetQosConfigFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceQosConfigCreate] Creation of resource QosConfig on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceQosConfigRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceQosConfigRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceQosConfigRefsCreate] Missing 'uuid' field for resource QosConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("qos-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsCreate] Retrieving QosConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objQosConfig := obj.(*QosConfig) // Fully set by Contrail backend
	if err := SetRefsQosConfigFromResource(objQosConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsCreate] Set refs on object QosConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objQosConfig.GetHref())
	if err := client.Update(objQosConfig); err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsCreate] Update refs for resource QosConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objQosConfig.GetUuid())
	return nil
}

func ResourceQosConfigRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosConfigRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("qos-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosConfigRead] Read resource qos-config on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*QosConfig)
	WriteQosConfigToResource(*object, d, m)
	return nil
}

func ResourceQosConfigRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosConfigRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("qos-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsRead] Read resource qos-config on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*QosConfig)
	WriteQosConfigRefsToResource(*object, d, m)
	return nil
}

func ResourceQosConfigUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosConfigUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("qos-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosConfigUpdate] Retrieving QosConfig with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*QosConfig)
	UpdateQosConfigFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceQosConfigUpdate] Update of resource QosConfig on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceQosConfigRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosConfigRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("qos-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsUpdate] Retrieving QosConfig with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*QosConfig)
	UpdateQosConfigRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsUpdate] Update of resource QosConfig on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceQosConfigDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosConfigDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("qos-config", d.Id()); err != nil {
		return fmt.Errorf("[ResourceQosConfigDelete] Deletion of resource qos-config on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceQosConfigRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceQosConfigRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceQosConfigRefsDelete] Missing 'uuid' field for resource QosConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("qos-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsDelete] Retrieving QosConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objQosConfig := obj.(*QosConfig) // Fully set by Contrail backend
	if err := DeleteRefsQosConfigFromResource(objQosConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsDelete] Set refs on object QosConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objQosConfig.GetHref())
	if err := client.Update(objQosConfig); err != nil {
		return fmt.Errorf("[ResourceQosConfigRefsDelete] Delete refs for resource QosConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objQosConfig.GetUuid())
	return nil
}

func ResourceQosConfigSchema() map[string]*schema.Schema {
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
		"qos_config_type": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"dscp_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceQosIdForwardingClassPairs(),
		},
		"vlan_priority_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceQosIdForwardingClassPairs(),
		},
		"mpls_exp_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceQosIdForwardingClassPairs(),
		},
		"default_forwarding_class_id": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
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

func ResourceQosConfigRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"global_system_config_refs": &schema.Schema{
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

func ResourceQosConfig() *schema.Resource {
	return &schema.Resource{
		Create: ResourceQosConfigCreate,
		Read:   ResourceQosConfigRead,
		Update: ResourceQosConfigUpdate,
		Delete: ResourceQosConfigDelete,
		Schema: ResourceQosConfigSchema(),
	}
}

func ResourceQosConfigRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceQosConfigRefsCreate,
		Read:   ResourceQosConfigRefsRead,
		Update: ResourceQosConfigRefsUpdate,
		Delete: ResourceQosConfigRefsDelete,
		Schema: ResourceQosConfigRefsSchema(),
	}
}

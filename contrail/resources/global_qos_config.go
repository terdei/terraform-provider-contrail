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

func SetGlobalQosConfigFromResource(object *GlobalQosConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetGlobalQosConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("control_traffic_dscp"); ok {
		member := new(ControlTrafficDscpType)
		SetControlTrafficDscpTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetControlTrafficDscp(member)
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

func SetRefsGlobalQosConfigFromResource(object *GlobalQosConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsGlobalQosConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteGlobalQosConfigToResource(object GlobalQosConfig, d *schema.ResourceData, m interface{}) {

	control_traffic_dscpObj := object.GetControlTrafficDscp()
	d.Set("control_traffic_dscp", TakeControlTrafficDscpTypeAsMap(&control_traffic_dscpObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeGlobalQosConfigAsMap(object *GlobalQosConfig) map[string]interface{} {
	omap := make(map[string]interface{})

	control_traffic_dscpObj := object.GetControlTrafficDscp()
	omap["control_traffic_dscp"] = TakeControlTrafficDscpTypeAsMap(&control_traffic_dscpObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateGlobalQosConfigFromResource(object *GlobalQosConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("control_traffic_dscp") {
		if val, ok := d.GetOk("control_traffic_dscp"); ok {
			member := new(ControlTrafficDscpType)
			SetControlTrafficDscpTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetControlTrafficDscp(member)
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

func ResourceGlobalQosConfigCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalQosConfigCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(GlobalQosConfig)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceGlobalQosConfigCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "GlobalQosConfig", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetGlobalQosConfigFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigCreate] Creation of resource GlobalQosConfig on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceGlobalQosConfigRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceGlobalQosConfigRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceGlobalQosConfigRefsCreate] Missing 'uuid' field for resource GlobalQosConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("global-qos-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigRefsCreate] Retrieving GlobalQosConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objGlobalQosConfig := obj.(*GlobalQosConfig) // Fully set by Contrail backend
	if err := SetRefsGlobalQosConfigFromResource(objGlobalQosConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigRefsCreate] Set refs on object GlobalQosConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objGlobalQosConfig.GetHref())
	if err := client.Update(objGlobalQosConfig); err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigRefsCreate] Update refs for resource GlobalQosConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objGlobalQosConfig.GetUuid())
	return nil
}

func ResourceGlobalQosConfigRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalQosConfigREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("global-qos-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigRead] Read resource global-qos-config on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*GlobalQosConfig)
	WriteGlobalQosConfigToResource(*object, d, m)
	return nil
}

func ResourceGlobalQosConfigRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalQosConfigRefsREAD")
	return nil
}

func ResourceGlobalQosConfigUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalQosConfigUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("global-qos-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigResourceUpdate] Retrieving GlobalQosConfig with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*GlobalQosConfig)
	UpdateGlobalQosConfigFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigUpdate] Update of resource GlobalQosConfig on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceGlobalQosConfigRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalQosConfigRefsUpdate")
	return nil
}

func ResourceGlobalQosConfigDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalQosConfigDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("global-qos-config", d.Id()); err != nil {
		return fmt.Errorf("[ResourceGlobalQosConfigDelete] Deletion of resource global-qos-config on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceGlobalQosConfigRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceGlobalQosConfigRefsDelete: %v", d.Id())
	return nil
}

func ResourceGlobalQosConfigSchema() map[string]*schema.Schema {
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
		"control_traffic_dscp": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceControlTrafficDscpType(),
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

func ResourceGlobalQosConfigRefsSchema() map[string]*schema.Schema {
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

func ResourceGlobalQosConfig() *schema.Resource {
	return &schema.Resource{
		Create: ResourceGlobalQosConfigCreate,
		Read:   ResourceGlobalQosConfigRead,
		Update: ResourceGlobalQosConfigUpdate,
		Delete: ResourceGlobalQosConfigDelete,
		Schema: ResourceGlobalQosConfigSchema(),
	}
}

func ResourceGlobalQosConfigRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceGlobalQosConfigRefsCreate,
		Read:   ResourceGlobalQosConfigRefsRead,
		Update: ResourceGlobalQosConfigRefsUpdate,
		Delete: ResourceGlobalQosConfigRefsDelete,
		Schema: ResourceGlobalQosConfigRefsSchema(),
	}
}

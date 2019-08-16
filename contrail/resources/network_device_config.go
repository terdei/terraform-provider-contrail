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

func SetNetworkDeviceConfigFromResource(object *NetworkDeviceConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetNetworkDeviceConfigFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsNetworkDeviceConfigFromResource(object *NetworkDeviceConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsNetworkDeviceConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("physical_router_refs"); ok {
		log.Printf("Got ref physical_router_refs -- will call: object.AddPhysicalRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("physical-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving physical-router by Uuid = %v as ref for PhysicalRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPhysicalRouter(refObj.(*PhysicalRouter))
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

func DeleteRefsNetworkDeviceConfigFromResource(object *NetworkDeviceConfig, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsNetworkDeviceConfigFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("physical_router_refs"); ok {
		log.Printf("Got ref physical_router_refs -- will call: object.DeletePhysicalRouter(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePhysicalRouter(refId.(string))
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

func WriteNetworkDeviceConfigToResource(object NetworkDeviceConfig, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeNetworkDeviceConfigAsMap(object *NetworkDeviceConfig) map[string]interface{} {
	omap := make(map[string]interface{})

	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateNetworkDeviceConfigFromResource(object *NetworkDeviceConfig, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
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

func ResourceNetworkDeviceConfigCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkDeviceConfigCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(NetworkDeviceConfig)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceNetworkDeviceConfigCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "NetworkDeviceConfig", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetNetworkDeviceConfigFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigCreate] Creation of resource NetworkDeviceConfig on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceNetworkDeviceConfigRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkDeviceConfigRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsCreate] Missing 'uuid' field for resource NetworkDeviceConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("network-device-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsCreate] Retrieving NetworkDeviceConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNetworkDeviceConfig := obj.(*NetworkDeviceConfig) // Fully set by Contrail backend
	if err := SetRefsNetworkDeviceConfigFromResource(objNetworkDeviceConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsCreate] Set refs on object NetworkDeviceConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNetworkDeviceConfig.GetHref())
	if err := client.Update(objNetworkDeviceConfig); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsCreate] Update refs for resource NetworkDeviceConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNetworkDeviceConfig.GetUuid())
	return nil
}

func ResourceNetworkDeviceConfigRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkDeviceConfigREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("network-device-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRead] Read resource network-device-config on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*NetworkDeviceConfig)
	WriteNetworkDeviceConfigToResource(*object, d, m)
	return nil
}

func ResourceNetworkDeviceConfigRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkDeviceConfigRefsREAD")
	return nil
}

func ResourceNetworkDeviceConfigUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkDeviceConfigUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("network-device-config", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigResourceUpdate] Retrieving NetworkDeviceConfig with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*NetworkDeviceConfig)
	UpdateNetworkDeviceConfigFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigUpdate] Update of resource NetworkDeviceConfig on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceNetworkDeviceConfigRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkDeviceConfigRefsUpdate")
	return nil
}

func ResourceNetworkDeviceConfigDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkDeviceConfigDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("network-device-config", d.Id()); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigDelete] Deletion of resource network-device-config on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceNetworkDeviceConfigRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkDeviceConfigRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsDelete] Missing 'uuid' field for resource NetworkDeviceConfig")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("network-device-config", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsDelete] Retrieving NetworkDeviceConfig with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNetworkDeviceConfig := obj.(*NetworkDeviceConfig) // Fully set by Contrail backend
	if err := DeleteRefsNetworkDeviceConfigFromResource(objNetworkDeviceConfig, d, m); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsDelete] Set refs on object NetworkDeviceConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNetworkDeviceConfig.GetHref())
	if err := client.Update(objNetworkDeviceConfig); err != nil {
		return fmt.Errorf("[ResourceNetworkDeviceConfigRefsDelete] Delete refs for resource NetworkDeviceConfig (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNetworkDeviceConfig.GetUuid())
	return nil
}

func ResourceNetworkDeviceConfigSchema() map[string]*schema.Schema {
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

func ResourceNetworkDeviceConfigRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"physical_router_refs": &schema.Schema{
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

func ResourceNetworkDeviceConfig() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNetworkDeviceConfigCreate,
		Read:   ResourceNetworkDeviceConfigRead,
		Update: ResourceNetworkDeviceConfigUpdate,
		Delete: ResourceNetworkDeviceConfigDelete,
		Schema: ResourceNetworkDeviceConfigSchema(),
	}
}

func ResourceNetworkDeviceConfigRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNetworkDeviceConfigRefsCreate,
		Read:   ResourceNetworkDeviceConfigRefsRead,
		Update: ResourceNetworkDeviceConfigRefsUpdate,
		Delete: ResourceNetworkDeviceConfigRefsDelete,
		Schema: ResourceNetworkDeviceConfigRefsSchema(),
	}
}

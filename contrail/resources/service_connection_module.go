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

func SetServiceConnectionModuleFromResource(object *ServiceConnectionModule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceConnectionModuleFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("e2_service"); ok {
		object.SetE2Service(val.(string))
	}
	if val, ok := d.GetOk("service_type"); ok {
		object.SetServiceType(val.(string))
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

func SetRefsServiceConnectionModuleFromResource(object *ServiceConnectionModule, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceConnectionModuleFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_object_refs"); ok {
		log.Printf("Got ref service_object_refs -- will call: object.AddServiceObject(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-object", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-object by Uuid = %v as ref for ServiceObject on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceObject(refObj.(*ServiceObject))
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

func DeleteRefsServiceConnectionModuleFromResource(object *ServiceConnectionModule, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceConnectionModuleFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_object_refs"); ok {
		log.Printf("Got ref service_object_refs -- will call: object.DeleteServiceObject(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceObject(refId.(string))
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

func WriteServiceConnectionModuleToResource(object ServiceConnectionModule, d *schema.ResourceData, m interface{}) {

	d.Set("e2_service", object.GetE2Service())
	d.Set("service_type", object.GetServiceType())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceConnectionModuleRefsToResource(object ServiceConnectionModule, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetServiceObjectRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_object_refs", refList)
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

func TakeServiceConnectionModuleAsMap(object *ServiceConnectionModule) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["e2_service"] = object.GetE2Service()
	omap["service_type"] = object.GetServiceType()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceConnectionModuleFromResource(object *ServiceConnectionModule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("e2_service") {
		if val, ok := d.GetOk("e2_service"); ok {
			object.SetE2Service(val.(string))
		}
	}
	if d.HasChange("service_type") {
		if val, ok := d.GetOk("service_type"); ok {
			object.SetServiceType(val.(string))
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

func UpdateServiceConnectionModuleRefsFromResource(object *ServiceConnectionModule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("service_object_refs") {
		object.ClearServiceObject()
		if val, ok := d.GetOk("service_object_refs"); ok {
			log.Printf("Got ref service_object_refs -- will call: object.AddServiceObject(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-object", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceObject(refObj.(*ServiceObject))
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

func ResourceServiceConnectionModuleCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceConnectionModuleCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceConnectionModule)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceServiceConnectionModuleCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "ServiceConnectionModule", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceServiceConnectionModuleCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceConnectionModule", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceConnectionModuleFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleCreate] Creation of resource ServiceConnectionModule on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceConnectionModuleRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceConnectionModuleRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsCreate] Missing 'uuid' field for resource ServiceConnectionModule")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-connection-module", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsCreate] Retrieving ServiceConnectionModule with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceConnectionModule := obj.(*ServiceConnectionModule) // Fully set by Contrail backend
	if err := SetRefsServiceConnectionModuleFromResource(objServiceConnectionModule, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsCreate] Set refs on object ServiceConnectionModule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceConnectionModule.GetHref())
	if err := client.Update(objServiceConnectionModule); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsCreate] Update refs for resource ServiceConnectionModule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceConnectionModule.GetUuid())
	return nil
}

func ResourceServiceConnectionModuleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceConnectionModuleRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-connection-module", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRead] Read resource service-connection-module on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceConnectionModule)
	WriteServiceConnectionModuleToResource(*object, d, m)
	return nil
}

func ResourceServiceConnectionModuleRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceConnectionModuleRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-connection-module", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsRead] Read resource service-connection-module on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceConnectionModule)
	WriteServiceConnectionModuleRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceConnectionModuleUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceConnectionModuleUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-connection-module", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleUpdate] Retrieving ServiceConnectionModule with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceConnectionModule)
	UpdateServiceConnectionModuleFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleUpdate] Update of resource ServiceConnectionModule on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceConnectionModuleRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceConnectionModuleRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-connection-module", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsUpdate] Retrieving ServiceConnectionModule with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceConnectionModule)
	UpdateServiceConnectionModuleRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsUpdate] Update of resource ServiceConnectionModule on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceConnectionModuleDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceConnectionModuleDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-connection-module", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleDelete] Deletion of resource service-connection-module on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceConnectionModuleRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceConnectionModuleRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsDelete] Missing 'uuid' field for resource ServiceConnectionModule")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-connection-module", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsDelete] Retrieving ServiceConnectionModule with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceConnectionModule := obj.(*ServiceConnectionModule) // Fully set by Contrail backend
	if err := DeleteRefsServiceConnectionModuleFromResource(objServiceConnectionModule, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsDelete] Set refs on object ServiceConnectionModule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceConnectionModule.GetHref())
	if err := client.Update(objServiceConnectionModule); err != nil {
		return fmt.Errorf("[ResourceServiceConnectionModuleRefsDelete] Delete refs for resource ServiceConnectionModule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceConnectionModule.GetUuid())
	return nil
}

func ResourceServiceConnectionModuleSchema() map[string]*schema.Schema {
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
		"e2_service": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"service_type": &schema.Schema{
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

func ResourceServiceConnectionModuleRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_object_refs": &schema.Schema{
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

func ResourceServiceConnectionModule() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceConnectionModuleCreate,
		Read:   ResourceServiceConnectionModuleRead,
		Update: ResourceServiceConnectionModuleUpdate,
		Delete: ResourceServiceConnectionModuleDelete,
		Schema: ResourceServiceConnectionModuleSchema(),
	}
}

func ResourceServiceConnectionModuleRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceConnectionModuleRefsCreate,
		Read:   ResourceServiceConnectionModuleRefsRead,
		Update: ResourceServiceConnectionModuleRefsUpdate,
		Delete: ResourceServiceConnectionModuleRefsDelete,
		Schema: ResourceServiceConnectionModuleRefsSchema(),
	}
}

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

func SetServiceEndpointFromResource(object *ServiceEndpoint, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceEndpointFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsServiceEndpointFromResource(object *ServiceEndpoint, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceEndpointFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_connection_module_refs"); ok {
		log.Printf("Got ref service_connection_module_refs -- will call: object.AddServiceConnectionModule(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-connection-module", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-connection-module by Uuid = %v as ref for ServiceConnectionModule on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceConnectionModule(refObj.(*ServiceConnectionModule))
		}
	}
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

func DeleteRefsServiceEndpointFromResource(object *ServiceEndpoint, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceEndpointFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_connection_module_refs"); ok {
		log.Printf("Got ref service_connection_module_refs -- will call: object.DeleteServiceConnectionModule(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceConnectionModule(refId.(string))
		}
	}
	if val, ok := d.GetOk("physical_router_refs"); ok {
		log.Printf("Got ref physical_router_refs -- will call: object.DeletePhysicalRouter(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePhysicalRouter(refId.(string))
		}
	}
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

func WriteServiceEndpointToResource(object ServiceEndpoint, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceEndpointRefsToResource(object ServiceEndpoint, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetServiceConnectionModuleRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_connection_module_refs", refList)
	}
	if ref, err := object.GetPhysicalRouterRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("physical_router_refs", refList)
	}
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

func TakeServiceEndpointAsMap(object *ServiceEndpoint) map[string]interface{} {
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

func UpdateServiceEndpointFromResource(object *ServiceEndpoint, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateServiceEndpointRefsFromResource(object *ServiceEndpoint, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("service_connection_module_refs") {
		object.ClearServiceConnectionModule()
		if val, ok := d.GetOk("service_connection_module_refs"); ok {
			log.Printf("Got ref service_connection_module_refs -- will call: object.AddServiceConnectionModule(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-connection-module", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceConnectionModule(refObj.(*ServiceConnectionModule))
			}
		}
	}
	if d.HasChange("physical_router_refs") {
		object.ClearPhysicalRouter()
		if val, ok := d.GetOk("physical_router_refs"); ok {
			log.Printf("Got ref physical_router_refs -- will call: object.AddPhysicalRouter(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("physical-router", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddPhysicalRouter(refObj.(*PhysicalRouter))
			}
		}
	}
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

func ResourceServiceEndpointCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceEndpointCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceEndpoint)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceServiceEndpointCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "ServiceEndpoint", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceServiceEndpointCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceEndpoint", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceEndpointFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointCreate] Creation of resource ServiceEndpoint on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceEndpointRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceEndpointRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceEndpointRefsCreate] Missing 'uuid' field for resource ServiceEndpoint")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-endpoint", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsCreate] Retrieving ServiceEndpoint with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceEndpoint := obj.(*ServiceEndpoint) // Fully set by Contrail backend
	if err := SetRefsServiceEndpointFromResource(objServiceEndpoint, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsCreate] Set refs on object ServiceEndpoint (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceEndpoint.GetHref())
	if err := client.Update(objServiceEndpoint); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsCreate] Update refs for resource ServiceEndpoint (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceEndpoint.GetUuid())
	return nil
}

func ResourceServiceEndpointRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceEndpointRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-endpoint", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRead] Read resource service-endpoint on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceEndpoint)
	WriteServiceEndpointToResource(*object, d, m)
	return nil
}

func ResourceServiceEndpointRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceEndpointRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-endpoint", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsRead] Read resource service-endpoint on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceEndpoint)
	WriteServiceEndpointRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceEndpointUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceEndpointUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-endpoint", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceEndpointUpdate] Retrieving ServiceEndpoint with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceEndpoint)
	UpdateServiceEndpointFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointUpdate] Update of resource ServiceEndpoint on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceEndpointRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceEndpointRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-endpoint", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsUpdate] Retrieving ServiceEndpoint with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceEndpoint)
	UpdateServiceEndpointRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsUpdate] Update of resource ServiceEndpoint on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceEndpointDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceEndpointDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-endpoint", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointDelete] Deletion of resource service-endpoint on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceEndpointRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceEndpointRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceEndpointRefsDelete] Missing 'uuid' field for resource ServiceEndpoint")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-endpoint", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsDelete] Retrieving ServiceEndpoint with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceEndpoint := obj.(*ServiceEndpoint) // Fully set by Contrail backend
	if err := DeleteRefsServiceEndpointFromResource(objServiceEndpoint, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsDelete] Set refs on object ServiceEndpoint (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceEndpoint.GetHref())
	if err := client.Update(objServiceEndpoint); err != nil {
		return fmt.Errorf("[ResourceServiceEndpointRefsDelete] Delete refs for resource ServiceEndpoint (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceEndpoint.GetUuid())
	return nil
}

func ResourceServiceEndpointSchema() map[string]*schema.Schema {
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

func ResourceServiceEndpointRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_connection_module_refs": &schema.Schema{
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

func ResourceServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceEndpointCreate,
		Read:   ResourceServiceEndpointRead,
		Update: ResourceServiceEndpointUpdate,
		Delete: ResourceServiceEndpointDelete,
		Schema: ResourceServiceEndpointSchema(),
	}
}

func ResourceServiceEndpointRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceEndpointRefsCreate,
		Read:   ResourceServiceEndpointRefsRead,
		Update: ResourceServiceEndpointRefsUpdate,
		Delete: ResourceServiceEndpointRefsDelete,
		Schema: ResourceServiceEndpointRefsSchema(),
	}
}

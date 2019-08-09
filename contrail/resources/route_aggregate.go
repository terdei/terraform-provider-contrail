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

func SetRouteAggregateFromResource(object *RouteAggregate, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetRouteAggregateFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsRouteAggregateFromResource(object *RouteAggregate, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsRouteAggregateFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-instance", refId.(string))
			dataObj := new(ServiceInterfaceTag)
			SetServiceInterfaceTagFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-instance by Uuid = %v as ref for ServiceInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceInstance(refObj.(*ServiceInstance), *dataObj)
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

func WriteRouteAggregateToResource(object RouteAggregate, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeRouteAggregateAsMap(object *RouteAggregate) map[string]interface{} {
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

func UpdateRouteAggregateFromResource(object *RouteAggregate, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceRouteAggregateCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteAggregateCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(RouteAggregate)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceRouteAggregateCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "RouteAggregate", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetRouteAggregateFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceRouteAggregateCreate] Creation of resource RouteAggregate on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceRouteAggregateRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteAggregateRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRouteAggregateRefsCreate] Missing 'uuid' field for resource RouteAggregate")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("route-aggregate", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRouteAggregateRefsCreate] Retrieving RouteAggregate with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRouteAggregate := obj.(*RouteAggregate) // Fully set by Contrail backend
	if err := SetRefsRouteAggregateFromResource(objRouteAggregate, d, m); err != nil {
		return fmt.Errorf("[ResourceRouteAggregateRefsCreate] Set refs on object RouteAggregate (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRouteAggregate.GetHref())
	if err := client.Update(objRouteAggregate); err != nil {
		return fmt.Errorf("[ResourceRouteAggregateRefsCreate] Update refs for resource RouteAggregate (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRouteAggregate.GetUuid())
	return nil
}

func ResourceRouteAggregateRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteAggregateREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("route-aggregate", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteAggregateRead] Read resource route-aggregate on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RouteAggregate)
	WriteRouteAggregateToResource(*object, d, m)
	return nil
}

func ResourceRouteAggregateRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteAggregateRefsREAD")
	return nil
}

func ResourceRouteAggregateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteAggregateUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("route-aggregate", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteAggregateResourceUpdate] Retrieving RouteAggregate with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RouteAggregate)
	UpdateRouteAggregateFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRouteAggregateUpdate] Update of resource RouteAggregate on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRouteAggregateRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteAggregateRefsUpdate")
	return nil
}

func ResourceRouteAggregateDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteAggregateDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("route-aggregate", d.Id()); err != nil {
		return fmt.Errorf("[ResourceRouteAggregateDelete] Deletion of resource route-aggregate on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceRouteAggregateRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteAggregateRefsDelete: %v", d.Id())
	return nil
}

func ResourceRouteAggregateSchema() map[string]*schema.Schema {
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

func ResourceRouteAggregateRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_instance_refs": &schema.Schema{
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
						Elem:     ResourceServiceInterfaceTag(),
						Required: true,
					},
				},
			},
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourceRouteAggregate() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteAggregateCreate,
		Read:   ResourceRouteAggregateRead,
		Update: ResourceRouteAggregateUpdate,
		Delete: ResourceRouteAggregateDelete,
		Schema: ResourceRouteAggregateSchema(),
	}
}

func ResourceRouteAggregateRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteAggregateRefsCreate,
		Read:   ResourceRouteAggregateRefsRead,
		Update: ResourceRouteAggregateRefsUpdate,
		Delete: ResourceRouteAggregateRefsDelete,
		Schema: ResourceRouteAggregateRefsSchema(),
	}
}

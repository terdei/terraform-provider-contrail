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

func SetRoutingInstanceFromResource(object *RoutingInstance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetRoutingInstanceFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsRoutingInstanceFromResource(object *RoutingInstance, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsRoutingInstanceFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsRoutingInstanceFromResource(object *RoutingInstance, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsRoutingInstanceFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteRoutingInstanceToResource(object RoutingInstance, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteRoutingInstanceRefsToResource(object RoutingInstance, d *schema.ResourceData, m interface{}) {

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

func TakeRoutingInstanceAsMap(object *RoutingInstance) map[string]interface{} {
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

func UpdateRoutingInstanceFromResource(object *RoutingInstance, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateRoutingInstanceRefsFromResource(object *RoutingInstance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
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

func ResourceRoutingInstanceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRoutingInstanceCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(RoutingInstance)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceRoutingInstanceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "RoutingInstance", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceRoutingInstanceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "RoutingInstance", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetRoutingInstanceFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceCreate] Creation of resource RoutingInstance on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceRoutingInstanceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRoutingInstanceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRoutingInstanceRefsCreate] Missing 'uuid' field for resource RoutingInstance")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("routing-instance", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsCreate] Retrieving RoutingInstance with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRoutingInstance := obj.(*RoutingInstance) // Fully set by Contrail backend
	if err := SetRefsRoutingInstanceFromResource(objRoutingInstance, d, m); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsCreate] Set refs on object RoutingInstance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRoutingInstance.GetHref())
	if err := client.Update(objRoutingInstance); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsCreate] Update refs for resource RoutingInstance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRoutingInstance.GetUuid())
	return nil
}

func ResourceRoutingInstanceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingInstanceRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("routing-instance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRead] Read resource routing-instance on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RoutingInstance)
	WriteRoutingInstanceToResource(*object, d, m)
	return nil
}

func ResourceRoutingInstanceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingInstanceRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("routing-instance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsRead] Read resource routing-instance on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RoutingInstance)
	WriteRoutingInstanceRefsToResource(*object, d, m)
	return nil
}

func ResourceRoutingInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingInstanceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("routing-instance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceUpdate] Retrieving RoutingInstance with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RoutingInstance)
	UpdateRoutingInstanceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceUpdate] Update of resource RoutingInstance on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRoutingInstanceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingInstanceRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("routing-instance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsUpdate] Retrieving RoutingInstance with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RoutingInstance)
	UpdateRoutingInstanceRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsUpdate] Update of resource RoutingInstance on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRoutingInstanceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingInstanceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("routing-instance", d.Id()); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceDelete] Deletion of resource routing-instance on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceRoutingInstanceRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRoutingInstanceRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRoutingInstanceRefsDelete] Missing 'uuid' field for resource RoutingInstance")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("routing-instance", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsDelete] Retrieving RoutingInstance with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRoutingInstance := obj.(*RoutingInstance) // Fully set by Contrail backend
	if err := DeleteRefsRoutingInstanceFromResource(objRoutingInstance, d, m); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsDelete] Set refs on object RoutingInstance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRoutingInstance.GetHref())
	if err := client.Update(objRoutingInstance); err != nil {
		return fmt.Errorf("[ResourceRoutingInstanceRefsDelete] Delete refs for resource RoutingInstance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRoutingInstance.GetUuid())
	return nil
}

func ResourceRoutingInstanceSchema() map[string]*schema.Schema {
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

func ResourceRoutingInstanceRefsSchema() map[string]*schema.Schema {
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

func ResourceRoutingInstance() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRoutingInstanceCreate,
		Read:   ResourceRoutingInstanceRead,
		Update: ResourceRoutingInstanceUpdate,
		Delete: ResourceRoutingInstanceDelete,
		Schema: ResourceRoutingInstanceSchema(),
	}
}

func ResourceRoutingInstanceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRoutingInstanceRefsCreate,
		Read:   ResourceRoutingInstanceRefsRead,
		Update: ResourceRoutingInstanceRefsUpdate,
		Delete: ResourceRoutingInstanceRefsDelete,
		Schema: ResourceRoutingInstanceRefsSchema(),
	}
}

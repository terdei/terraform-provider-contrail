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

func SetRouteTargetFromResource(object *RouteTarget, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetRouteTargetFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsRouteTargetFromResource(object *RouteTarget, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsRouteTargetFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsRouteTargetFromResource(object *RouteTarget, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsRouteTargetFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteRouteTargetToResource(object RouteTarget, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteRouteTargetRefsToResource(object RouteTarget, d *schema.ResourceData, m interface{}) {

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

func TakeRouteTargetAsMap(object *RouteTarget) map[string]interface{} {
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

func UpdateRouteTargetFromResource(object *RouteTarget, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateRouteTargetRefsFromResource(object *RouteTarget, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceRouteTargetCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteTargetCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(RouteTarget)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceRouteTargetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "RouteTarget", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceRouteTargetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "RouteTarget", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetRouteTargetFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceRouteTargetCreate] Creation of resource RouteTarget on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceRouteTargetRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteTargetRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRouteTargetRefsCreate] Missing 'uuid' field for resource RouteTarget")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("route-target", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsCreate] Retrieving RouteTarget with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRouteTarget := obj.(*RouteTarget) // Fully set by Contrail backend
	if err := SetRefsRouteTargetFromResource(objRouteTarget, d, m); err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsCreate] Set refs on object RouteTarget (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRouteTarget.GetHref())
	if err := client.Update(objRouteTarget); err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsCreate] Update refs for resource RouteTarget (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRouteTarget.GetUuid())
	return nil
}

func ResourceRouteTargetRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTargetRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("route-target", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteTargetRead] Read resource route-target on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RouteTarget)
	WriteRouteTargetToResource(*object, d, m)
	return nil
}

func ResourceRouteTargetRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTargetRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("route-target", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsRead] Read resource route-target on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RouteTarget)
	WriteRouteTargetRefsToResource(*object, d, m)
	return nil
}

func ResourceRouteTargetUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTargetUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("route-target", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteTargetUpdate] Retrieving RouteTarget with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RouteTarget)
	UpdateRouteTargetFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRouteTargetUpdate] Update of resource RouteTarget on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRouteTargetRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTargetRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("route-target", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsUpdate] Retrieving RouteTarget with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RouteTarget)
	UpdateRouteTargetRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsUpdate] Update of resource RouteTarget on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRouteTargetDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTargetDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("route-target", d.Id()); err != nil {
		return fmt.Errorf("[ResourceRouteTargetDelete] Deletion of resource route-target on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceRouteTargetRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteTargetRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRouteTargetRefsDelete] Missing 'uuid' field for resource RouteTarget")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("route-target", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsDelete] Retrieving RouteTarget with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRouteTarget := obj.(*RouteTarget) // Fully set by Contrail backend
	if err := DeleteRefsRouteTargetFromResource(objRouteTarget, d, m); err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsDelete] Set refs on object RouteTarget (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRouteTarget.GetHref())
	if err := client.Update(objRouteTarget); err != nil {
		return fmt.Errorf("[ResourceRouteTargetRefsDelete] Delete refs for resource RouteTarget (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRouteTarget.GetUuid())
	return nil
}

func ResourceRouteTargetSchema() map[string]*schema.Schema {
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

func ResourceRouteTargetRefsSchema() map[string]*schema.Schema {
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

func ResourceRouteTarget() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteTargetCreate,
		Read:   ResourceRouteTargetRead,
		Update: ResourceRouteTargetUpdate,
		Delete: ResourceRouteTargetDelete,
		Schema: ResourceRouteTargetSchema(),
	}
}

func ResourceRouteTargetRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteTargetRefsCreate,
		Read:   ResourceRouteTargetRefsRead,
		Update: ResourceRouteTargetRefsUpdate,
		Delete: ResourceRouteTargetRefsDelete,
		Schema: ResourceRouteTargetRefsSchema(),
	}
}

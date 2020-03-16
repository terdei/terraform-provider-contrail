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

func SetBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetBgpRouterFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsBgpRouterFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsBgpRouterFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteBgpRouterToResource(object BgpRouter, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteBgpRouterRefsToResource(object BgpRouter, d *schema.ResourceData, m interface{}) {

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

func TakeBgpRouterAsMap(object *BgpRouter) map[string]interface{} {
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

func UpdateBgpRouterFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateBgpRouterRefsFromResource(object *BgpRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceBgpRouterCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpRouterCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(BgpRouter)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceBgpRouterCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "BgpRouter", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceBgpRouterCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "BgpRouter", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetBgpRouterFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceBgpRouterCreate] Creation of resource BgpRouter on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceBgpRouterRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpRouterRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Missing 'uuid' field for resource BgpRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgp-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Retrieving BgpRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpRouter := obj.(*BgpRouter) // Fully set by Contrail backend
	if err := SetRefsBgpRouterFromResource(objBgpRouter, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Set refs on object BgpRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpRouter.GetHref())
	if err := client.Update(objBgpRouter); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsCreate] Update refs for resource BgpRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpRouter.GetUuid())
	return nil
}

func ResourceBgpRouterRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgp-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRead] Read resource bgp-router on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BgpRouter)
	WriteBgpRouterToResource(*object, d, m)
	return nil
}

func ResourceBgpRouterRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgp-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsRead] Read resource bgp-router on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BgpRouter)
	WriteBgpRouterRefsToResource(*object, d, m)
	return nil
}

func ResourceBgpRouterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgp-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterUpdate] Retrieving BgpRouter with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BgpRouter)
	UpdateBgpRouterFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpRouterUpdate] Update of resource BgpRouter on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpRouterRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgp-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsUpdate] Retrieving BgpRouter with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BgpRouter)
	UpdateBgpRouterRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsUpdate] Update of resource BgpRouter on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpRouterDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpRouterDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("bgp-router", d.Id()); err != nil {
		return fmt.Errorf("[ResourceBgpRouterDelete] Deletion of resource bgp-router on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceBgpRouterRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpRouterRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpRouterRefsDelete] Missing 'uuid' field for resource BgpRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgp-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsDelete] Retrieving BgpRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpRouter := obj.(*BgpRouter) // Fully set by Contrail backend
	if err := DeleteRefsBgpRouterFromResource(objBgpRouter, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsDelete] Set refs on object BgpRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpRouter.GetHref())
	if err := client.Update(objBgpRouter); err != nil {
		return fmt.Errorf("[ResourceBgpRouterRefsDelete] Delete refs for resource BgpRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpRouter.GetUuid())
	return nil
}

func ResourceBgpRouterSchema() map[string]*schema.Schema {
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

func ResourceBgpRouterRefsSchema() map[string]*schema.Schema {
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

func ResourceBgpRouter() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpRouterCreate,
		Read:   ResourceBgpRouterRead,
		Update: ResourceBgpRouterUpdate,
		Delete: ResourceBgpRouterDelete,
		Schema: ResourceBgpRouterSchema(),
	}
}

func ResourceBgpRouterRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpRouterRefsCreate,
		Read:   ResourceBgpRouterRefsRead,
		Update: ResourceBgpRouterRefsUpdate,
		Delete: ResourceBgpRouterRefsDelete,
		Schema: ResourceBgpRouterRefsSchema(),
	}
}

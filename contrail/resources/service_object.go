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

func SetServiceObjectFromResource(object *ServiceObject, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceObjectFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsServiceObjectFromResource(object *ServiceObject, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceObjectFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsServiceObjectFromResource(object *ServiceObject, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceObjectFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteServiceObjectToResource(object ServiceObject, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceObjectRefsToResource(object ServiceObject, d *schema.ResourceData, m interface{}) {

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

func TakeServiceObjectAsMap(object *ServiceObject) map[string]interface{} {
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

func UpdateServiceObjectFromResource(object *ServiceObject, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateServiceObjectRefsFromResource(object *ServiceObject, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceServiceObjectCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceObjectCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceObject)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceObjectCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceObject", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceObjectFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceObjectCreate] Creation of resource ServiceObject on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceObjectRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceObjectRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceObjectRefsCreate] Missing 'uuid' field for resource ServiceObject")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-object", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsCreate] Retrieving ServiceObject with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceObject := obj.(*ServiceObject) // Fully set by Contrail backend
	if err := SetRefsServiceObjectFromResource(objServiceObject, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsCreate] Set refs on object ServiceObject (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceObject.GetHref())
	if err := client.Update(objServiceObject); err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsCreate] Update refs for resource ServiceObject (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceObject.GetUuid())
	return nil
}

func ResourceServiceObjectRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceObjectRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-object", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceObjectRead] Read resource service-object on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceObject)
	WriteServiceObjectToResource(*object, d, m)
	return nil
}

func ResourceServiceObjectRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceObjectRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-object", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsRead] Read resource service-object on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceObject)
	WriteServiceObjectRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceObjectUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceObjectUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-object", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceObjectUpdate] Retrieving ServiceObject with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceObject)
	UpdateServiceObjectFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceObjectUpdate] Update of resource ServiceObject on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceObjectRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceObjectRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-object", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsUpdate] Retrieving ServiceObject with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceObject)
	UpdateServiceObjectRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsUpdate] Update of resource ServiceObject on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceObjectDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceObjectDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-object", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceObjectDelete] Deletion of resource service-object on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceObjectRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceObjectRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceObjectRefsDelete] Missing 'uuid' field for resource ServiceObject")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-object", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsDelete] Retrieving ServiceObject with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceObject := obj.(*ServiceObject) // Fully set by Contrail backend
	if err := DeleteRefsServiceObjectFromResource(objServiceObject, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsDelete] Set refs on object ServiceObject (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceObject.GetHref())
	if err := client.Update(objServiceObject); err != nil {
		return fmt.Errorf("[ResourceServiceObjectRefsDelete] Delete refs for resource ServiceObject (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceObject.GetUuid())
	return nil
}

func ResourceServiceObjectSchema() map[string]*schema.Schema {
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

func ResourceServiceObjectRefsSchema() map[string]*schema.Schema {
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

func ResourceServiceObject() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceObjectCreate,
		Read:   ResourceServiceObjectRead,
		Update: ResourceServiceObjectUpdate,
		Delete: ResourceServiceObjectDelete,
		Schema: ResourceServiceObjectSchema(),
	}
}

func ResourceServiceObjectRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceObjectRefsCreate,
		Read:   ResourceServiceObjectRefsRead,
		Update: ResourceServiceObjectRefsUpdate,
		Delete: ResourceServiceObjectRefsDelete,
		Schema: ResourceServiceObjectRefsSchema(),
	}
}

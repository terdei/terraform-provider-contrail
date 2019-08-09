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

func SetTagTypeFromResource(object *TagType, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetTagTypeFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("tag_type_id"); ok {
		object.SetTagTypeId(val.(string))
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

func SetRefsTagTypeFromResource(object *TagType, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsTagTypeFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteTagTypeToResource(object TagType, d *schema.ResourceData, m interface{}) {

	d.Set("tag_type_id", object.GetTagTypeId())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeTagTypeAsMap(object *TagType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["tag_type_id"] = object.GetTagTypeId()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateTagTypeFromResource(object *TagType, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("tag_type_id") {
		if val, ok := d.GetOk("tag_type_id"); ok {
			object.SetTagTypeId(val.(string))
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

func ResourceTagTypeCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceTagTypeCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(TagType)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceTagTypeCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "TagType", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetTagTypeFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceTagTypeCreate] Creation of resource TagType on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceTagTypeRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceTagTypeRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceTagTypeRefsCreate] Missing 'uuid' field for resource TagType")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("tag-type", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceTagTypeRefsCreate] Retrieving TagType with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objTagType := obj.(*TagType) // Fully set by Contrail backend
	if err := SetRefsTagTypeFromResource(objTagType, d, m); err != nil {
		return fmt.Errorf("[ResourceTagTypeRefsCreate] Set refs on object TagType (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objTagType.GetHref())
	if err := client.Update(objTagType); err != nil {
		return fmt.Errorf("[ResourceTagTypeRefsCreate] Update refs for resource TagType (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objTagType.GetUuid())
	return nil
}

func ResourceTagTypeRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagTypeREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("tag-type", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceTagTypeRead] Read resource tag-type on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*TagType)
	WriteTagTypeToResource(*object, d, m)
	return nil
}

func ResourceTagTypeRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagTypeRefsREAD")
	return nil
}

func ResourceTagTypeUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagTypeUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("tag-type", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceTagTypeResourceUpdate] Retrieving TagType with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*TagType)
	UpdateTagTypeFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceTagTypeUpdate] Update of resource TagType on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceTagTypeRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagTypeRefsUpdate")
	return nil
}

func ResourceTagTypeDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagTypeDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("tag-type", d.Id()); err != nil {
		return fmt.Errorf("[ResourceTagTypeDelete] Deletion of resource tag-type on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceTagTypeRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagTypeRefsDelete: %v", d.Id())
	return nil
}

func ResourceTagTypeSchema() map[string]*schema.Schema {
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
		"tag_type_id": &schema.Schema{
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

func ResourceTagTypeRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourceTagType() *schema.Resource {
	return &schema.Resource{
		Create: ResourceTagTypeCreate,
		Read:   ResourceTagTypeRead,
		Update: ResourceTagTypeUpdate,
		Delete: ResourceTagTypeDelete,
		Schema: ResourceTagTypeSchema(),
	}
}

func ResourceTagTypeRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceTagTypeRefsCreate,
		Read:   ResourceTagTypeRefsRead,
		Update: ResourceTagTypeRefsUpdate,
		Delete: ResourceTagTypeRefsDelete,
		Schema: ResourceTagTypeRefsSchema(),
	}
}

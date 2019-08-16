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

func SetTagFromResource(object *Tag, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetTagFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("tag_type_name"); ok {
		object.SetTagTypeName(val.(string))
	}
	if val, ok := d.GetOk("tag_value"); ok {
		object.SetTagValue(val.(string))
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

func SetRefsTagFromResource(object *Tag, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsTagFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("tag_type_refs"); ok {
		log.Printf("Got ref tag_type_refs -- will call: object.AddTagType(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("tag-type", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving tag-type by Uuid = %v as ref for TagType on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddTagType(refObj.(*TagType))
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

func WriteTagToResource(object Tag, d *schema.ResourceData, m interface{}) {

	d.Set("tag_type_name", object.GetTagTypeName())
	d.Set("tag_value", object.GetTagValue())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeTagAsMap(object *Tag) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["tag_type_name"] = object.GetTagTypeName()
	omap["tag_value"] = object.GetTagValue()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateTagFromResource(object *Tag, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("tag_type_name") {
		if val, ok := d.GetOk("tag_type_name"); ok {
			object.SetTagTypeName(val.(string))
		}
	}
	if d.HasChange("tag_value") {
		if val, ok := d.GetOk("tag_value"); ok {
			object.SetTagValue(val.(string))
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

func ResourceTagCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceTagCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Tag)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceTagCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Tag", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetTagFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceTagCreate] Creation of resource Tag on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceTagRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceTagRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceTagRefsCreate] Missing 'uuid' field for resource Tag")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("tag", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceTagRefsCreate] Retrieving Tag with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objTag := obj.(*Tag) // Fully set by Contrail backend
	if err := SetRefsTagFromResource(objTag, d, m); err != nil {
		return fmt.Errorf("[ResourceTagRefsCreate] Set refs on object Tag (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objTag.GetHref())
	if err := client.Update(objTag); err != nil {
		return fmt.Errorf("[ResourceTagRefsCreate] Update refs for resource Tag (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objTag.GetUuid())
	return nil
}

func ResourceTagRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("tag", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceTagRead] Read resource tag on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Tag)
	WriteTagToResource(*object, d, m)
	return nil
}

func ResourceTagRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagRefsREAD")
	return nil
}

func ResourceTagUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("tag", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceTagResourceUpdate] Retrieving Tag with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Tag)
	UpdateTagFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceTagUpdate] Update of resource Tag on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceTagRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagRefsUpdate")
	return nil
}

func ResourceTagDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("tag", d.Id()); err != nil {
		return fmt.Errorf("[ResourceTagDelete] Deletion of resource tag on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceTagRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceTagRefsDelete: %v", d.Id())
	return nil
}

func ResourceTagSchema() map[string]*schema.Schema {
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
		"tag_type_name": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"tag_value": &schema.Schema{
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

func ResourceTagRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tag_type_refs": &schema.Schema{
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

func ResourceTag() *schema.Resource {
	return &schema.Resource{
		Create: ResourceTagCreate,
		Read:   ResourceTagRead,
		Update: ResourceTagUpdate,
		Delete: ResourceTagDelete,
		Schema: ResourceTagSchema(),
	}
}

func ResourceTagRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceTagRefsCreate,
		Read:   ResourceTagRefsRead,
		Update: ResourceTagRefsUpdate,
		Delete: ResourceTagRefsDelete,
		Schema: ResourceTagRefsSchema(),
	}
}

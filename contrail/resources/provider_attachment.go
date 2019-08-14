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

func SetProviderAttachmentFromResource(object *ProviderAttachment, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetProviderAttachmentFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsProviderAttachmentFromResource(object *ProviderAttachment, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsProviderAttachmentFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_router_refs"); ok {
		log.Printf("Got ref virtual_router_refs -- will call: object.AddVirtualRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-router by Uuid = %v as ref for VirtualRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualRouter(refObj.(*VirtualRouter))
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

func WriteProviderAttachmentToResource(object ProviderAttachment, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeProviderAttachmentAsMap(object *ProviderAttachment) map[string]interface{} {
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

func UpdateProviderAttachmentFromResource(object *ProviderAttachment, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceProviderAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceProviderAttachmentCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ProviderAttachment)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceProviderAttachmentCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ProviderAttachment", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetProviderAttachmentFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentCreate] Creation of resource ProviderAttachment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceProviderAttachmentRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceProviderAttachmentRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceProviderAttachmentRefsCreate] Missing 'uuid' field for resource ProviderAttachment")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("provider-attachment", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentRefsCreate] Retrieving ProviderAttachment with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objProviderAttachment := obj.(*ProviderAttachment) // Fully set by Contrail backend
	if err := SetRefsProviderAttachmentFromResource(objProviderAttachment, d, m); err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentRefsCreate] Set refs on object ProviderAttachment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objProviderAttachment.GetHref())
	if err := client.Update(objProviderAttachment); err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentRefsCreate] Update refs for resource ProviderAttachment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objProviderAttachment.GetUuid())
	return nil
}

func ResourceProviderAttachmentRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProviderAttachmentREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("provider-attachment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentRead] Read resource provider-attachment on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ProviderAttachment)
	WriteProviderAttachmentToResource(*object, d, m)
	return nil
}

func ResourceProviderAttachmentRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProviderAttachmentRefsREAD")
	return nil
}

func ResourceProviderAttachmentUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProviderAttachmentUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("provider-attachment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentResourceUpdate] Retrieving ProviderAttachment with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ProviderAttachment)
	UpdateProviderAttachmentFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentUpdate] Update of resource ProviderAttachment on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceProviderAttachmentRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProviderAttachmentRefsUpdate")
	return nil
}

func ResourceProviderAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProviderAttachmentDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("provider-attachment", d.Id()); err != nil {
		return fmt.Errorf("[ResourceProviderAttachmentDelete] Deletion of resource provider-attachment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceProviderAttachmentRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProviderAttachmentRefsDelete: %v", d.Id())
	return nil
}

func ResourceProviderAttachmentSchema() map[string]*schema.Schema {
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

func ResourceProviderAttachmentRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_router_refs": &schema.Schema{
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

func ResourceProviderAttachment() *schema.Resource {
	return &schema.Resource{
		Create: ResourceProviderAttachmentCreate,
		Read:   ResourceProviderAttachmentRead,
		Update: ResourceProviderAttachmentUpdate,
		Delete: ResourceProviderAttachmentDelete,
		Schema: ResourceProviderAttachmentSchema(),
	}
}

func ResourceProviderAttachmentRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceProviderAttachmentRefsCreate,
		Read:   ResourceProviderAttachmentRefsRead,
		Update: ResourceProviderAttachmentRefsUpdate,
		Delete: ResourceProviderAttachmentRefsDelete,
		Schema: ResourceProviderAttachmentRefsSchema(),
	}
}

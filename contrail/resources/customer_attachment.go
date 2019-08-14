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

func SetCustomerAttachmentFromResource(object *CustomerAttachment, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetCustomerAttachmentFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsCustomerAttachmentFromResource(object *CustomerAttachment, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsCustomerAttachmentFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine-interface by Uuid = %v as ref for VirtualMachineInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
		}
	}
	if val, ok := d.GetOk("floating_ip_refs"); ok {
		log.Printf("Got ref floating_ip_refs -- will call: object.AddFloatingIp(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("floating-ip", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving floating-ip by Uuid = %v as ref for FloatingIp on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddFloatingIp(refObj.(*FloatingIp))
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

func WriteCustomerAttachmentToResource(object CustomerAttachment, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeCustomerAttachmentAsMap(object *CustomerAttachment) map[string]interface{} {
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

func UpdateCustomerAttachmentFromResource(object *CustomerAttachment, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceCustomerAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceCustomerAttachmentCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(CustomerAttachment)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceCustomerAttachmentCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "CustomerAttachment", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetCustomerAttachmentFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentCreate] Creation of resource CustomerAttachment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceCustomerAttachmentRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceCustomerAttachmentRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceCustomerAttachmentRefsCreate] Missing 'uuid' field for resource CustomerAttachment")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("customer-attachment", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentRefsCreate] Retrieving CustomerAttachment with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objCustomerAttachment := obj.(*CustomerAttachment) // Fully set by Contrail backend
	if err := SetRefsCustomerAttachmentFromResource(objCustomerAttachment, d, m); err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentRefsCreate] Set refs on object CustomerAttachment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objCustomerAttachment.GetHref())
	if err := client.Update(objCustomerAttachment); err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentRefsCreate] Update refs for resource CustomerAttachment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objCustomerAttachment.GetUuid())
	return nil
}

func ResourceCustomerAttachmentRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceCustomerAttachmentREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("customer-attachment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentRead] Read resource customer-attachment on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*CustomerAttachment)
	WriteCustomerAttachmentToResource(*object, d, m)
	return nil
}

func ResourceCustomerAttachmentRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceCustomerAttachmentRefsREAD")
	return nil
}

func ResourceCustomerAttachmentUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceCustomerAttachmentUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("customer-attachment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentResourceUpdate] Retrieving CustomerAttachment with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*CustomerAttachment)
	UpdateCustomerAttachmentFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentUpdate] Update of resource CustomerAttachment on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceCustomerAttachmentRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceCustomerAttachmentRefsUpdate")
	return nil
}

func ResourceCustomerAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceCustomerAttachmentDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("customer-attachment", d.Id()); err != nil {
		return fmt.Errorf("[ResourceCustomerAttachmentDelete] Deletion of resource customer-attachment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceCustomerAttachmentRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceCustomerAttachmentRefsDelete: %v", d.Id())
	return nil
}

func ResourceCustomerAttachmentSchema() map[string]*schema.Schema {
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

func ResourceCustomerAttachmentRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_machine_interface_refs": &schema.Schema{
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
		"floating_ip_refs": &schema.Schema{
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

func ResourceCustomerAttachment() *schema.Resource {
	return &schema.Resource{
		Create: ResourceCustomerAttachmentCreate,
		Read:   ResourceCustomerAttachmentRead,
		Update: ResourceCustomerAttachmentUpdate,
		Delete: ResourceCustomerAttachmentDelete,
		Schema: ResourceCustomerAttachmentSchema(),
	}
}

func ResourceCustomerAttachmentRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceCustomerAttachmentRefsCreate,
		Read:   ResourceCustomerAttachmentRefsRead,
		Update: ResourceCustomerAttachmentRefsUpdate,
		Delete: ResourceCustomerAttachmentRefsDelete,
		Schema: ResourceCustomerAttachmentRefsSchema(),
	}
}

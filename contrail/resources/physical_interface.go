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

func SetPhysicalInterfaceFromResource(object *PhysicalInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetPhysicalInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("ethernet_segment_identifier"); ok {
		object.SetEthernetSegmentIdentifier(val.(string))
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

func SetRefsPhysicalInterfaceFromResource(object *PhysicalInterface, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsPhysicalInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("physical_interface_refs"); ok {
		log.Printf("Got ref physical_interface_refs -- will call: object.AddPhysicalInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("physical-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving physical-interface by Uuid = %v as ref for PhysicalInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPhysicalInterface(refObj.(*PhysicalInterface))
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

func WritePhysicalInterfaceToResource(object PhysicalInterface, d *schema.ResourceData, m interface{}) {

	d.Set("ethernet_segment_identifier", object.GetEthernetSegmentIdentifier())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakePhysicalInterfaceAsMap(object *PhysicalInterface) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["ethernet_segment_identifier"] = object.GetEthernetSegmentIdentifier()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdatePhysicalInterfaceFromResource(object *PhysicalInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("ethernet_segment_identifier") {
		if val, ok := d.GetOk("ethernet_segment_identifier"); ok {
			object.SetEthernetSegmentIdentifier(val.(string))
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

func ResourcePhysicalInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePhysicalInterfaceCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(PhysicalInterface)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourcePhysicalInterfaceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "PhysicalInterface", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetPhysicalInterfaceFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceCreate] Creation of resource PhysicalInterface on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourcePhysicalInterfaceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePhysicalInterfaceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePhysicalInterfaceRefsCreate] Missing 'uuid' field for resource PhysicalInterface")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("physical-interface", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceRefsCreate] Retrieving PhysicalInterface with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPhysicalInterface := obj.(*PhysicalInterface) // Fully set by Contrail backend
	if err := SetRefsPhysicalInterfaceFromResource(objPhysicalInterface, d, m); err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceRefsCreate] Set refs on object PhysicalInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPhysicalInterface.GetHref())
	if err := client.Update(objPhysicalInterface); err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceRefsCreate] Update refs for resource PhysicalInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPhysicalInterface.GetUuid())
	return nil
}

func ResourcePhysicalInterfaceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalInterfaceREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("physical-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceRead] Read resource physical-interface on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*PhysicalInterface)
	WritePhysicalInterfaceToResource(*object, d, m)
	return nil
}

func ResourcePhysicalInterfaceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalInterfaceRefsREAD")
	return nil
}

func ResourcePhysicalInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalInterfaceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("physical-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceResourceUpdate] Retrieving PhysicalInterface with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*PhysicalInterface)
	UpdatePhysicalInterfaceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceUpdate] Update of resource PhysicalInterface on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourcePhysicalInterfaceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalInterfaceRefsUpdate")
	return nil
}

func ResourcePhysicalInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalInterfaceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("physical-interface", d.Id()); err != nil {
		return fmt.Errorf("[ResourcePhysicalInterfaceDelete] Deletion of resource physical-interface on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourcePhysicalInterfaceRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalInterfaceRefsDelete: %v", d.Id())
	return nil
}

func ResourcePhysicalInterfaceSchema() map[string]*schema.Schema {
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
		"ethernet_segment_identifier": &schema.Schema{
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

func ResourcePhysicalInterfaceRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"physical_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePhysicalInterface(),
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourcePhysicalInterface() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePhysicalInterfaceCreate,
		Read:   ResourcePhysicalInterfaceRead,
		Update: ResourcePhysicalInterfaceUpdate,
		Delete: ResourcePhysicalInterfaceDelete,
		Schema: ResourcePhysicalInterfaceSchema(),
	}
}

func ResourcePhysicalInterfaceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePhysicalInterfaceRefsCreate,
		Read:   ResourcePhysicalInterfaceRefsRead,
		Update: ResourcePhysicalInterfaceRefsUpdate,
		Delete: ResourcePhysicalInterfaceRefsDelete,
		Schema: ResourcePhysicalInterfaceRefsSchema(),
	}
}

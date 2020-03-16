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

func SetDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetDiscoveryServiceAssignmentFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsDiscoveryServiceAssignmentFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsDiscoveryServiceAssignmentFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteDiscoveryServiceAssignmentToResource(object DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteDiscoveryServiceAssignmentRefsToResource(object DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}) {

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

func TakeDiscoveryServiceAssignmentAsMap(object *DiscoveryServiceAssignment) map[string]interface{} {
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

func UpdateDiscoveryServiceAssignmentFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateDiscoveryServiceAssignmentRefsFromResource(object *DiscoveryServiceAssignment, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceDiscoveryServiceAssignmentCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDiscoveryServiceAssignmentCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(DiscoveryServiceAssignment)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceDiscoveryServiceAssignmentCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "DiscoveryServiceAssignment", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceDiscoveryServiceAssignmentCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "DiscoveryServiceAssignment", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetDiscoveryServiceAssignmentFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentCreate] Creation of resource DiscoveryServiceAssignment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDiscoveryServiceAssignmentRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Missing 'uuid' field for resource DiscoveryServiceAssignment")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("discovery-service-assignment", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Retrieving DiscoveryServiceAssignment with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDiscoveryServiceAssignment := obj.(*DiscoveryServiceAssignment) // Fully set by Contrail backend
	if err := SetRefsDiscoveryServiceAssignmentFromResource(objDiscoveryServiceAssignment, d, m); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Set refs on object DiscoveryServiceAssignment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDiscoveryServiceAssignment.GetHref())
	if err := client.Update(objDiscoveryServiceAssignment); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsCreate] Update refs for resource DiscoveryServiceAssignment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDiscoveryServiceAssignment.GetUuid())
	return nil
}

func ResourceDiscoveryServiceAssignmentRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("discovery-service-assignment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRead] Read resource discovery-service-assignment on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*DiscoveryServiceAssignment)
	WriteDiscoveryServiceAssignmentToResource(*object, d, m)
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("discovery-service-assignment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsRead] Read resource discovery-service-assignment on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*DiscoveryServiceAssignment)
	WriteDiscoveryServiceAssignmentRefsToResource(*object, d, m)
	return nil
}

func ResourceDiscoveryServiceAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("discovery-service-assignment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentUpdate] Retrieving DiscoveryServiceAssignment with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*DiscoveryServiceAssignment)
	UpdateDiscoveryServiceAssignmentFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentUpdate] Update of resource DiscoveryServiceAssignment on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("discovery-service-assignment", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsUpdate] Retrieving DiscoveryServiceAssignment with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*DiscoveryServiceAssignment)
	UpdateDiscoveryServiceAssignmentRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsUpdate] Update of resource DiscoveryServiceAssignment on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDiscoveryServiceAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDiscoveryServiceAssignmentDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("discovery-service-assignment", d.Id()); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentDelete] Deletion of resource discovery-service-assignment on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceDiscoveryServiceAssignmentRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDiscoveryServiceAssignmentRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsDelete] Missing 'uuid' field for resource DiscoveryServiceAssignment")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("discovery-service-assignment", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsDelete] Retrieving DiscoveryServiceAssignment with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDiscoveryServiceAssignment := obj.(*DiscoveryServiceAssignment) // Fully set by Contrail backend
	if err := DeleteRefsDiscoveryServiceAssignmentFromResource(objDiscoveryServiceAssignment, d, m); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsDelete] Set refs on object DiscoveryServiceAssignment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDiscoveryServiceAssignment.GetHref())
	if err := client.Update(objDiscoveryServiceAssignment); err != nil {
		return fmt.Errorf("[ResourceDiscoveryServiceAssignmentRefsDelete] Delete refs for resource DiscoveryServiceAssignment (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDiscoveryServiceAssignment.GetUuid())
	return nil
}

func ResourceDiscoveryServiceAssignmentSchema() map[string]*schema.Schema {
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

func ResourceDiscoveryServiceAssignmentRefsSchema() map[string]*schema.Schema {
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

func ResourceDiscoveryServiceAssignment() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDiscoveryServiceAssignmentCreate,
		Read:   ResourceDiscoveryServiceAssignmentRead,
		Update: ResourceDiscoveryServiceAssignmentUpdate,
		Delete: ResourceDiscoveryServiceAssignmentDelete,
		Schema: ResourceDiscoveryServiceAssignmentSchema(),
	}
}

func ResourceDiscoveryServiceAssignmentRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDiscoveryServiceAssignmentRefsCreate,
		Read:   ResourceDiscoveryServiceAssignmentRefsRead,
		Update: ResourceDiscoveryServiceAssignmentRefsUpdate,
		Delete: ResourceDiscoveryServiceAssignmentRefsDelete,
		Schema: ResourceDiscoveryServiceAssignmentRefsSchema(),
	}
}

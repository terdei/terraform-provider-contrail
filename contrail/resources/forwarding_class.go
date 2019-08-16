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

func SetForwardingClassFromResource(object *ForwardingClass, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetForwardingClassFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("forwarding_class_dscp"); ok {
		object.SetForwardingClassDscp(val.(int))
	}
	if val, ok := d.GetOk("forwarding_class_vlan_priority"); ok {
		object.SetForwardingClassVlanPriority(val.(int))
	}
	if val, ok := d.GetOk("forwarding_class_mpls_exp"); ok {
		object.SetForwardingClassMplsExp(val.(int))
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

func SetRefsForwardingClassFromResource(object *ForwardingClass, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsForwardingClassFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("qos_queue_refs"); ok {
		log.Printf("Got ref qos_queue_refs -- will call: object.AddQosQueue(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("qos-queue", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving qos-queue by Uuid = %v as ref for QosQueue on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddQosQueue(refObj.(*QosQueue))
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

func DeleteRefsForwardingClassFromResource(object *ForwardingClass, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsForwardingClassFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("qos_queue_refs"); ok {
		log.Printf("Got ref qos_queue_refs -- will call: object.DeleteQosQueue(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteQosQueue(refId.(string))
		}
	}
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

func WriteForwardingClassToResource(object ForwardingClass, d *schema.ResourceData, m interface{}) {

	d.Set("forwarding_class_dscp", object.GetForwardingClassDscp())
	d.Set("forwarding_class_vlan_priority", object.GetForwardingClassVlanPriority())
	d.Set("forwarding_class_mpls_exp", object.GetForwardingClassMplsExp())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeForwardingClassAsMap(object *ForwardingClass) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["forwarding_class_dscp"] = object.GetForwardingClassDscp()
	omap["forwarding_class_vlan_priority"] = object.GetForwardingClassVlanPriority()
	omap["forwarding_class_mpls_exp"] = object.GetForwardingClassMplsExp()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateForwardingClassFromResource(object *ForwardingClass, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("forwarding_class_dscp") {
		if val, ok := d.GetOk("forwarding_class_dscp"); ok {
			object.SetForwardingClassDscp(val.(int))
		}
	}
	if d.HasChange("forwarding_class_vlan_priority") {
		if val, ok := d.GetOk("forwarding_class_vlan_priority"); ok {
			object.SetForwardingClassVlanPriority(val.(int))
		}
	}
	if d.HasChange("forwarding_class_mpls_exp") {
		if val, ok := d.GetOk("forwarding_class_mpls_exp"); ok {
			object.SetForwardingClassMplsExp(val.(int))
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

func ResourceForwardingClassCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceForwardingClassCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ForwardingClass)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceForwardingClassCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ForwardingClass", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetForwardingClassFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceForwardingClassCreate] Creation of resource ForwardingClass on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceForwardingClassRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceForwardingClassRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceForwardingClassRefsCreate] Missing 'uuid' field for resource ForwardingClass")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("forwarding-class", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceForwardingClassRefsCreate] Retrieving ForwardingClass with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objForwardingClass := obj.(*ForwardingClass) // Fully set by Contrail backend
	if err := SetRefsForwardingClassFromResource(objForwardingClass, d, m); err != nil {
		return fmt.Errorf("[ResourceForwardingClassRefsCreate] Set refs on object ForwardingClass (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objForwardingClass.GetHref())
	if err := client.Update(objForwardingClass); err != nil {
		return fmt.Errorf("[ResourceForwardingClassRefsCreate] Update refs for resource ForwardingClass (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objForwardingClass.GetUuid())
	return nil
}

func ResourceForwardingClassRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceForwardingClassREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("forwarding-class", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceForwardingClassRead] Read resource forwarding-class on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ForwardingClass)
	WriteForwardingClassToResource(*object, d, m)
	return nil
}

func ResourceForwardingClassRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceForwardingClassRefsREAD")
	return nil
}

func ResourceForwardingClassUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceForwardingClassUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("forwarding-class", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceForwardingClassResourceUpdate] Retrieving ForwardingClass with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ForwardingClass)
	UpdateForwardingClassFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceForwardingClassUpdate] Update of resource ForwardingClass on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceForwardingClassRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceForwardingClassRefsUpdate")
	return nil
}

func ResourceForwardingClassDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceForwardingClassDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("forwarding-class", d.Id()); err != nil {
		return fmt.Errorf("[ResourceForwardingClassDelete] Deletion of resource forwarding-class on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceForwardingClassRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceForwardingClassRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceForwardingClassRefsDelete] Missing 'uuid' field for resource ForwardingClass")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("forwarding-class", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceForwardingClassRefsDelete] Retrieving ForwardingClass with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objForwardingClass := obj.(*ForwardingClass) // Fully set by Contrail backend
	if err := DeleteRefsForwardingClassFromResource(objForwardingClass, d, m); err != nil {
		return fmt.Errorf("[ResourceForwardingClassRefsDelete] Set refs on object ForwardingClass (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objForwardingClass.GetHref())
	if err := client.Update(objForwardingClass); err != nil {
		return fmt.Errorf("[ResourceForwardingClassRefsDelete] Delete refs for resource ForwardingClass (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objForwardingClass.GetUuid())
	return nil
}

func ResourceForwardingClassSchema() map[string]*schema.Schema {
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
		"forwarding_class_dscp": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"forwarding_class_vlan_priority": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"forwarding_class_mpls_exp": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
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

func ResourceForwardingClassRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"qos_queue_refs": &schema.Schema{
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

func ResourceForwardingClass() *schema.Resource {
	return &schema.Resource{
		Create: ResourceForwardingClassCreate,
		Read:   ResourceForwardingClassRead,
		Update: ResourceForwardingClassUpdate,
		Delete: ResourceForwardingClassDelete,
		Schema: ResourceForwardingClassSchema(),
	}
}

func ResourceForwardingClassRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceForwardingClassRefsCreate,
		Read:   ResourceForwardingClassRefsRead,
		Update: ResourceForwardingClassRefsUpdate,
		Delete: ResourceForwardingClassRefsDelete,
		Schema: ResourceForwardingClassRefsSchema(),
	}
}

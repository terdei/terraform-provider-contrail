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

func SetQosQueueFromResource(object *QosQueue, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetQosQueueFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("min_bandwidth"); ok {
		object.SetMinBandwidth(val.(int))
	}
	if val, ok := d.GetOk("max_bandwidth"); ok {
		object.SetMaxBandwidth(val.(int))
	}
	if val, ok := d.GetOk("qos_queue_identifier"); ok {
		object.SetQosQueueIdentifier(val.(int))
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

func SetRefsQosQueueFromResource(object *QosQueue, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsQosQueueFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsQosQueueFromResource(object *QosQueue, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsQosQueueFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteQosQueueToResource(object QosQueue, d *schema.ResourceData, m interface{}) {

	d.Set("min_bandwidth", object.GetMinBandwidth())
	d.Set("max_bandwidth", object.GetMaxBandwidth())
	d.Set("qos_queue_identifier", object.GetQosQueueIdentifier())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteQosQueueRefsToResource(object QosQueue, d *schema.ResourceData, m interface{}) {

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

func TakeQosQueueAsMap(object *QosQueue) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["min_bandwidth"] = object.GetMinBandwidth()
	omap["max_bandwidth"] = object.GetMaxBandwidth()
	omap["qos_queue_identifier"] = object.GetQosQueueIdentifier()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateQosQueueFromResource(object *QosQueue, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("min_bandwidth") {
		if val, ok := d.GetOk("min_bandwidth"); ok {
			object.SetMinBandwidth(val.(int))
		}
	}
	if d.HasChange("max_bandwidth") {
		if val, ok := d.GetOk("max_bandwidth"); ok {
			object.SetMaxBandwidth(val.(int))
		}
	}
	if d.HasChange("qos_queue_identifier") {
		if val, ok := d.GetOk("qos_queue_identifier"); ok {
			object.SetQosQueueIdentifier(val.(int))
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

func UpdateQosQueueRefsFromResource(object *QosQueue, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceQosQueueCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceQosQueueCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(QosQueue)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceQosQueueCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "QosQueue", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceQosQueueCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "QosQueue", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetQosQueueFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceQosQueueCreate] Creation of resource QosQueue on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceQosQueueRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceQosQueueRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceQosQueueRefsCreate] Missing 'uuid' field for resource QosQueue")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("qos-queue", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsCreate] Retrieving QosQueue with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objQosQueue := obj.(*QosQueue) // Fully set by Contrail backend
	if err := SetRefsQosQueueFromResource(objQosQueue, d, m); err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsCreate] Set refs on object QosQueue (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objQosQueue.GetHref())
	if err := client.Update(objQosQueue); err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsCreate] Update refs for resource QosQueue (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objQosQueue.GetUuid())
	return nil
}

func ResourceQosQueueRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosQueueRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("qos-queue", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosQueueRead] Read resource qos-queue on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*QosQueue)
	WriteQosQueueToResource(*object, d, m)
	return nil
}

func ResourceQosQueueRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosQueueRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("qos-queue", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsRead] Read resource qos-queue on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*QosQueue)
	WriteQosQueueRefsToResource(*object, d, m)
	return nil
}

func ResourceQosQueueUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosQueueUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("qos-queue", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosQueueUpdate] Retrieving QosQueue with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*QosQueue)
	UpdateQosQueueFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceQosQueueUpdate] Update of resource QosQueue on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceQosQueueRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosQueueRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("qos-queue", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsUpdate] Retrieving QosQueue with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*QosQueue)
	UpdateQosQueueRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsUpdate] Update of resource QosQueue on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceQosQueueDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceQosQueueDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("qos-queue", d.Id()); err != nil {
		return fmt.Errorf("[ResourceQosQueueDelete] Deletion of resource qos-queue on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceQosQueueRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceQosQueueRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceQosQueueRefsDelete] Missing 'uuid' field for resource QosQueue")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("qos-queue", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsDelete] Retrieving QosQueue with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objQosQueue := obj.(*QosQueue) // Fully set by Contrail backend
	if err := DeleteRefsQosQueueFromResource(objQosQueue, d, m); err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsDelete] Set refs on object QosQueue (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objQosQueue.GetHref())
	if err := client.Update(objQosQueue); err != nil {
		return fmt.Errorf("[ResourceQosQueueRefsDelete] Delete refs for resource QosQueue (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objQosQueue.GetUuid())
	return nil
}

func ResourceQosQueueSchema() map[string]*schema.Schema {
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
		"min_bandwidth": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"max_bandwidth": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"qos_queue_identifier": &schema.Schema{
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

func ResourceQosQueueRefsSchema() map[string]*schema.Schema {
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

func ResourceQosQueue() *schema.Resource {
	return &schema.Resource{
		Create: ResourceQosQueueCreate,
		Read:   ResourceQosQueueRead,
		Update: ResourceQosQueueUpdate,
		Delete: ResourceQosQueueDelete,
		Schema: ResourceQosQueueSchema(),
	}
}

func ResourceQosQueueRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceQosQueueRefsCreate,
		Read:   ResourceQosQueueRefsRead,
		Update: ResourceQosQueueRefsUpdate,
		Delete: ResourceQosQueueRefsDelete,
		Schema: ResourceQosQueueRefsSchema(),
	}
}

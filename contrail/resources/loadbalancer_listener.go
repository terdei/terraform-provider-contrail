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

func SetLoadbalancerListenerFromResource(object *LoadbalancerListener, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLoadbalancerListenerFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_listener_properties"); ok {
		member := new(LoadbalancerListenerType)
		SetLoadbalancerListenerTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLoadbalancerListenerProperties(member)
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

func SetRefsLoadbalancerListenerFromResource(object *LoadbalancerListener, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLoadbalancerListenerFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_refs"); ok {
		log.Printf("Got ref loadbalancer_refs -- will call: object.AddLoadbalancer(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("loadbalancer", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving loadbalancer by Uuid = %v as ref for Loadbalancer on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddLoadbalancer(refObj.(*Loadbalancer))
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

func DeleteRefsLoadbalancerListenerFromResource(object *LoadbalancerListener, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsLoadbalancerListenerFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_refs"); ok {
		log.Printf("Got ref loadbalancer_refs -- will call: object.DeleteLoadbalancer(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteLoadbalancer(refId.(string))
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

func WriteLoadbalancerListenerToResource(object LoadbalancerListener, d *schema.ResourceData, m interface{}) {

	loadbalancer_listener_propertiesObj := object.GetLoadbalancerListenerProperties()
	d.Set("loadbalancer_listener_properties", TakeLoadbalancerListenerTypeAsMap(&loadbalancer_listener_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteLoadbalancerListenerRefsToResource(object LoadbalancerListener, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetLoadbalancerRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("loadbalancer_refs", refList)
	}
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

func TakeLoadbalancerListenerAsMap(object *LoadbalancerListener) map[string]interface{} {
	omap := make(map[string]interface{})

	loadbalancer_listener_propertiesObj := object.GetLoadbalancerListenerProperties()
	omap["loadbalancer_listener_properties"] = TakeLoadbalancerListenerTypeAsMap(&loadbalancer_listener_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLoadbalancerListenerFromResource(object *LoadbalancerListener, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("loadbalancer_listener_properties") {
		if val, ok := d.GetOk("loadbalancer_listener_properties"); ok {
			member := new(LoadbalancerListenerType)
			SetLoadbalancerListenerTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLoadbalancerListenerProperties(member)
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

func UpdateLoadbalancerListenerRefsFromResource(object *LoadbalancerListener, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("loadbalancer_refs") {
		object.ClearLoadbalancer()
		if val, ok := d.GetOk("loadbalancer_refs"); ok {
			log.Printf("Got ref loadbalancer_refs -- will call: object.AddLoadbalancer(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("loadbalancer", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddLoadbalancer(refObj.(*Loadbalancer))
			}
		}
	}
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

func ResourceLoadbalancerListenerCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerListenerCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LoadbalancerListener)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceLoadbalancerListenerCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "LoadbalancerListener", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceLoadbalancerListenerCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LoadbalancerListener", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLoadbalancerListenerFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerCreate] Creation of resource LoadbalancerListener on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLoadbalancerListenerRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerListenerRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsCreate] Missing 'uuid' field for resource LoadbalancerListener")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-listener", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsCreate] Retrieving LoadbalancerListener with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerListener := obj.(*LoadbalancerListener) // Fully set by Contrail backend
	if err := SetRefsLoadbalancerListenerFromResource(objLoadbalancerListener, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsCreate] Set refs on object LoadbalancerListener (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerListener.GetHref())
	if err := client.Update(objLoadbalancerListener); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsCreate] Update refs for resource LoadbalancerListener (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerListener.GetUuid())
	return nil
}

func ResourceLoadbalancerListenerRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerListenerRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-listener", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRead] Read resource loadbalancer-listener on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerListener)
	WriteLoadbalancerListenerToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerListenerRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerListenerRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-listener", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsRead] Read resource loadbalancer-listener on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerListener)
	WriteLoadbalancerListenerRefsToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerListenerUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerListenerUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-listener", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerUpdate] Retrieving LoadbalancerListener with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerListener)
	UpdateLoadbalancerListenerFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerUpdate] Update of resource LoadbalancerListener on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerListenerRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerListenerRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-listener", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsUpdate] Retrieving LoadbalancerListener with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerListener)
	UpdateLoadbalancerListenerRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsUpdate] Update of resource LoadbalancerListener on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerListenerDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerListenerDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("loadbalancer-listener", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerDelete] Deletion of resource loadbalancer-listener on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLoadbalancerListenerRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerListenerRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsDelete] Missing 'uuid' field for resource LoadbalancerListener")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-listener", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsDelete] Retrieving LoadbalancerListener with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerListener := obj.(*LoadbalancerListener) // Fully set by Contrail backend
	if err := DeleteRefsLoadbalancerListenerFromResource(objLoadbalancerListener, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsDelete] Set refs on object LoadbalancerListener (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerListener.GetHref())
	if err := client.Update(objLoadbalancerListener); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerRefsDelete] Delete refs for resource LoadbalancerListener (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerListener.GetUuid())
	return nil
}

func ResourceLoadbalancerListenerSchema() map[string]*schema.Schema {
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
		"loadbalancer_listener_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLoadbalancerListenerType(),
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

func ResourceLoadbalancerListenerRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"loadbalancer_refs": &schema.Schema{
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

func ResourceLoadbalancerListener() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerListenerCreate,
		Read:   ResourceLoadbalancerListenerRead,
		Update: ResourceLoadbalancerListenerUpdate,
		Delete: ResourceLoadbalancerListenerDelete,
		Schema: ResourceLoadbalancerListenerSchema(),
	}
}

func ResourceLoadbalancerListenerRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerListenerRefsCreate,
		Read:   ResourceLoadbalancerListenerRefsRead,
		Update: ResourceLoadbalancerListenerRefsUpdate,
		Delete: ResourceLoadbalancerListenerRefsDelete,
		Schema: ResourceLoadbalancerListenerRefsSchema(),
	}
}

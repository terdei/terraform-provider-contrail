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

func ResourceLoadbalancerListenerCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerListenerCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LoadbalancerListener)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceLoadbalancerListenerCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LoadbalancerListener", client.GetServer(), err)
		}
		object.SetParent(parent)
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
	log.Printf("ResourceLoadbalancerListenerREAD")
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
	log.Printf("ResourceLoadbalancerListenerRefsREAD")
	return nil
}

func ResourceLoadbalancerListenerUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerListenerUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-listener", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerListenerResourceUpdate] Retrieving LoadbalancerListener with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
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
	log.Printf("ResourceLoadbalancerListenerRefsDelete: %v", d.Id())
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
			Elem:     ResourceLoadbalancer(),
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

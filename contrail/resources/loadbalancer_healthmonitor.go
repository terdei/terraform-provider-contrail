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

func SetLoadbalancerHealthmonitorFromResource(object *LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLoadbalancerHealthmonitorFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_healthmonitor_properties"); ok {
		member := new(LoadbalancerHealthmonitorType)
		SetLoadbalancerHealthmonitorTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLoadbalancerHealthmonitorProperties(member)
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

func SetRefsLoadbalancerHealthmonitorFromResource(object *LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLoadbalancerHealthmonitorFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsLoadbalancerHealthmonitorFromResource(object *LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsLoadbalancerHealthmonitorFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteLoadbalancerHealthmonitorToResource(object LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}) {

	loadbalancer_healthmonitor_propertiesObj := object.GetLoadbalancerHealthmonitorProperties()
	d.Set("loadbalancer_healthmonitor_properties", TakeLoadbalancerHealthmonitorTypeAsMap(&loadbalancer_healthmonitor_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteLoadbalancerHealthmonitorRefsToResource(object LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}) {

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

func TakeLoadbalancerHealthmonitorAsMap(object *LoadbalancerHealthmonitor) map[string]interface{} {
	omap := make(map[string]interface{})

	loadbalancer_healthmonitor_propertiesObj := object.GetLoadbalancerHealthmonitorProperties()
	omap["loadbalancer_healthmonitor_properties"] = TakeLoadbalancerHealthmonitorTypeAsMap(&loadbalancer_healthmonitor_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLoadbalancerHealthmonitorFromResource(object *LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("loadbalancer_healthmonitor_properties") {
		if val, ok := d.GetOk("loadbalancer_healthmonitor_properties"); ok {
			member := new(LoadbalancerHealthmonitorType)
			SetLoadbalancerHealthmonitorTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLoadbalancerHealthmonitorProperties(member)
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

func UpdateLoadbalancerHealthmonitorRefsFromResource(object *LoadbalancerHealthmonitor, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceLoadbalancerHealthmonitorCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerHealthmonitorCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LoadbalancerHealthmonitor)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceLoadbalancerHealthmonitorCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "LoadbalancerHealthmonitor", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceLoadbalancerHealthmonitorCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LoadbalancerHealthmonitor", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLoadbalancerHealthmonitorFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorCreate] Creation of resource LoadbalancerHealthmonitor on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLoadbalancerHealthmonitorRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerHealthmonitorRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsCreate] Missing 'uuid' field for resource LoadbalancerHealthmonitor")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-healthmonitor", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsCreate] Retrieving LoadbalancerHealthmonitor with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerHealthmonitor := obj.(*LoadbalancerHealthmonitor) // Fully set by Contrail backend
	if err := SetRefsLoadbalancerHealthmonitorFromResource(objLoadbalancerHealthmonitor, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsCreate] Set refs on object LoadbalancerHealthmonitor (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerHealthmonitor.GetHref())
	if err := client.Update(objLoadbalancerHealthmonitor); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsCreate] Update refs for resource LoadbalancerHealthmonitor (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerHealthmonitor.GetUuid())
	return nil
}

func ResourceLoadbalancerHealthmonitorRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerHealthmonitorRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-healthmonitor", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRead] Read resource loadbalancer-healthmonitor on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerHealthmonitor)
	WriteLoadbalancerHealthmonitorToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerHealthmonitorRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerHealthmonitorRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-healthmonitor", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsRead] Read resource loadbalancer-healthmonitor on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerHealthmonitor)
	WriteLoadbalancerHealthmonitorRefsToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerHealthmonitorUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerHealthmonitorUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-healthmonitor", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorUpdate] Retrieving LoadbalancerHealthmonitor with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerHealthmonitor)
	UpdateLoadbalancerHealthmonitorFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorUpdate] Update of resource LoadbalancerHealthmonitor on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerHealthmonitorRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerHealthmonitorRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-healthmonitor", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsUpdate] Retrieving LoadbalancerHealthmonitor with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerHealthmonitor)
	UpdateLoadbalancerHealthmonitorRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsUpdate] Update of resource LoadbalancerHealthmonitor on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerHealthmonitorDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerHealthmonitorDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("loadbalancer-healthmonitor", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorDelete] Deletion of resource loadbalancer-healthmonitor on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLoadbalancerHealthmonitorRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerHealthmonitorRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsDelete] Missing 'uuid' field for resource LoadbalancerHealthmonitor")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-healthmonitor", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsDelete] Retrieving LoadbalancerHealthmonitor with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerHealthmonitor := obj.(*LoadbalancerHealthmonitor) // Fully set by Contrail backend
	if err := DeleteRefsLoadbalancerHealthmonitorFromResource(objLoadbalancerHealthmonitor, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsDelete] Set refs on object LoadbalancerHealthmonitor (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerHealthmonitor.GetHref())
	if err := client.Update(objLoadbalancerHealthmonitor); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerHealthmonitorRefsDelete] Delete refs for resource LoadbalancerHealthmonitor (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerHealthmonitor.GetUuid())
	return nil
}

func ResourceLoadbalancerHealthmonitorSchema() map[string]*schema.Schema {
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
		"loadbalancer_healthmonitor_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLoadbalancerHealthmonitorType(),
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

func ResourceLoadbalancerHealthmonitorRefsSchema() map[string]*schema.Schema {
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

func ResourceLoadbalancerHealthmonitor() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerHealthmonitorCreate,
		Read:   ResourceLoadbalancerHealthmonitorRead,
		Update: ResourceLoadbalancerHealthmonitorUpdate,
		Delete: ResourceLoadbalancerHealthmonitorDelete,
		Schema: ResourceLoadbalancerHealthmonitorSchema(),
	}
}

func ResourceLoadbalancerHealthmonitorRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerHealthmonitorRefsCreate,
		Read:   ResourceLoadbalancerHealthmonitorRefsRead,
		Update: ResourceLoadbalancerHealthmonitorRefsUpdate,
		Delete: ResourceLoadbalancerHealthmonitorRefsDelete,
		Schema: ResourceLoadbalancerHealthmonitorRefsSchema(),
	}
}

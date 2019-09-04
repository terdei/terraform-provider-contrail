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

func SetLoadbalancerMemberFromResource(object *LoadbalancerMember, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLoadbalancerMemberFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_member_properties"); ok {
		member := new(LoadbalancerMemberType)
		SetLoadbalancerMemberTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLoadbalancerMemberProperties(member)
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

func SetRefsLoadbalancerMemberFromResource(object *LoadbalancerMember, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLoadbalancerMemberFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsLoadbalancerMemberFromResource(object *LoadbalancerMember, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsLoadbalancerMemberFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteLoadbalancerMemberToResource(object LoadbalancerMember, d *schema.ResourceData, m interface{}) {

	loadbalancer_member_propertiesObj := object.GetLoadbalancerMemberProperties()
	d.Set("loadbalancer_member_properties", TakeLoadbalancerMemberTypeAsMap(&loadbalancer_member_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteLoadbalancerMemberRefsToResource(object LoadbalancerMember, d *schema.ResourceData, m interface{}) {

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

func TakeLoadbalancerMemberAsMap(object *LoadbalancerMember) map[string]interface{} {
	omap := make(map[string]interface{})

	loadbalancer_member_propertiesObj := object.GetLoadbalancerMemberProperties()
	omap["loadbalancer_member_properties"] = TakeLoadbalancerMemberTypeAsMap(&loadbalancer_member_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLoadbalancerMemberFromResource(object *LoadbalancerMember, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("loadbalancer_member_properties") {
		if val, ok := d.GetOk("loadbalancer_member_properties"); ok {
			member := new(LoadbalancerMemberType)
			SetLoadbalancerMemberTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLoadbalancerMemberProperties(member)
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

func UpdateLoadbalancerMemberRefsFromResource(object *LoadbalancerMember, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceLoadbalancerMemberCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerMemberCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LoadbalancerMember)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceLoadbalancerMemberCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LoadbalancerMember", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLoadbalancerMemberFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberCreate] Creation of resource LoadbalancerMember on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLoadbalancerMemberRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerMemberRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsCreate] Missing 'uuid' field for resource LoadbalancerMember")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-member", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsCreate] Retrieving LoadbalancerMember with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerMember := obj.(*LoadbalancerMember) // Fully set by Contrail backend
	if err := SetRefsLoadbalancerMemberFromResource(objLoadbalancerMember, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsCreate] Set refs on object LoadbalancerMember (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerMember.GetHref())
	if err := client.Update(objLoadbalancerMember); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsCreate] Update refs for resource LoadbalancerMember (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerMember.GetUuid())
	return nil
}

func ResourceLoadbalancerMemberRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerMemberRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-member", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRead] Read resource loadbalancer-member on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerMember)
	WriteLoadbalancerMemberToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerMemberRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerMemberRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-member", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsRead] Read resource loadbalancer-member on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerMember)
	WriteLoadbalancerMemberRefsToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerMemberUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerMemberUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-member", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberUpdate] Retrieving LoadbalancerMember with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerMember)
	UpdateLoadbalancerMemberFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberUpdate] Update of resource LoadbalancerMember on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerMemberRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerMemberRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-member", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsUpdate] Retrieving LoadbalancerMember with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerMember)
	UpdateLoadbalancerMemberRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsUpdate] Update of resource LoadbalancerMember on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerMemberDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerMemberDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("loadbalancer-member", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberDelete] Deletion of resource loadbalancer-member on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLoadbalancerMemberRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerMemberRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsDelete] Missing 'uuid' field for resource LoadbalancerMember")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-member", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsDelete] Retrieving LoadbalancerMember with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerMember := obj.(*LoadbalancerMember) // Fully set by Contrail backend
	if err := DeleteRefsLoadbalancerMemberFromResource(objLoadbalancerMember, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsDelete] Set refs on object LoadbalancerMember (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerMember.GetHref())
	if err := client.Update(objLoadbalancerMember); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerMemberRefsDelete] Delete refs for resource LoadbalancerMember (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerMember.GetUuid())
	return nil
}

func ResourceLoadbalancerMemberSchema() map[string]*schema.Schema {
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
		"loadbalancer_member_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLoadbalancerMemberType(),
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

func ResourceLoadbalancerMemberRefsSchema() map[string]*schema.Schema {
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

func ResourceLoadbalancerMember() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerMemberCreate,
		Read:   ResourceLoadbalancerMemberRead,
		Update: ResourceLoadbalancerMemberUpdate,
		Delete: ResourceLoadbalancerMemberDelete,
		Schema: ResourceLoadbalancerMemberSchema(),
	}
}

func ResourceLoadbalancerMemberRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerMemberRefsCreate,
		Read:   ResourceLoadbalancerMemberRefsRead,
		Update: ResourceLoadbalancerMemberRefsUpdate,
		Delete: ResourceLoadbalancerMemberRefsDelete,
		Schema: ResourceLoadbalancerMemberRefsSchema(),
	}
}

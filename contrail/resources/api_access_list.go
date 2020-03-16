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

func SetApiAccessListFromResource(object *ApiAccessList, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetApiAccessListFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("api_access_list_entries"); ok {
		member := new(RbacRuleEntriesType)
		SetRbacRuleEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetApiAccessListEntries(member)
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

func SetRefsApiAccessListFromResource(object *ApiAccessList, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsApiAccessListFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsApiAccessListFromResource(object *ApiAccessList, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsApiAccessListFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteApiAccessListToResource(object ApiAccessList, d *schema.ResourceData, m interface{}) {

	api_access_list_entriesObj := object.GetApiAccessListEntries()
	d.Set("api_access_list_entries", TakeRbacRuleEntriesTypeAsMap(&api_access_list_entriesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteApiAccessListRefsToResource(object ApiAccessList, d *schema.ResourceData, m interface{}) {

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

func TakeApiAccessListAsMap(object *ApiAccessList) map[string]interface{} {
	omap := make(map[string]interface{})

	api_access_list_entriesObj := object.GetApiAccessListEntries()
	omap["api_access_list_entries"] = TakeRbacRuleEntriesTypeAsMap(&api_access_list_entriesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateApiAccessListFromResource(object *ApiAccessList, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("api_access_list_entries") {
		if val, ok := d.GetOk("api_access_list_entries"); ok {
			member := new(RbacRuleEntriesType)
			SetRbacRuleEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetApiAccessListEntries(member)
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

func UpdateApiAccessListRefsFromResource(object *ApiAccessList, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceApiAccessListCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceApiAccessListCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ApiAccessList)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceApiAccessListCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "ApiAccessList", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceApiAccessListCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ApiAccessList", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetApiAccessListFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceApiAccessListCreate] Creation of resource ApiAccessList on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceApiAccessListRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceApiAccessListRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceApiAccessListRefsCreate] Missing 'uuid' field for resource ApiAccessList")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("api-access-list", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsCreate] Retrieving ApiAccessList with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objApiAccessList := obj.(*ApiAccessList) // Fully set by Contrail backend
	if err := SetRefsApiAccessListFromResource(objApiAccessList, d, m); err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsCreate] Set refs on object ApiAccessList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objApiAccessList.GetHref())
	if err := client.Update(objApiAccessList); err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsCreate] Update refs for resource ApiAccessList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objApiAccessList.GetUuid())
	return nil
}

func ResourceApiAccessListRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApiAccessListRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("api-access-list", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApiAccessListRead] Read resource api-access-list on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ApiAccessList)
	WriteApiAccessListToResource(*object, d, m)
	return nil
}

func ResourceApiAccessListRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApiAccessListRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("api-access-list", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsRead] Read resource api-access-list on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ApiAccessList)
	WriteApiAccessListRefsToResource(*object, d, m)
	return nil
}

func ResourceApiAccessListUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApiAccessListUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("api-access-list", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApiAccessListUpdate] Retrieving ApiAccessList with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ApiAccessList)
	UpdateApiAccessListFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceApiAccessListUpdate] Update of resource ApiAccessList on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceApiAccessListRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApiAccessListRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("api-access-list", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsUpdate] Retrieving ApiAccessList with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ApiAccessList)
	UpdateApiAccessListRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsUpdate] Update of resource ApiAccessList on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceApiAccessListDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApiAccessListDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("api-access-list", d.Id()); err != nil {
		return fmt.Errorf("[ResourceApiAccessListDelete] Deletion of resource api-access-list on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceApiAccessListRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceApiAccessListRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceApiAccessListRefsDelete] Missing 'uuid' field for resource ApiAccessList")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("api-access-list", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsDelete] Retrieving ApiAccessList with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objApiAccessList := obj.(*ApiAccessList) // Fully set by Contrail backend
	if err := DeleteRefsApiAccessListFromResource(objApiAccessList, d, m); err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsDelete] Set refs on object ApiAccessList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objApiAccessList.GetHref())
	if err := client.Update(objApiAccessList); err != nil {
		return fmt.Errorf("[ResourceApiAccessListRefsDelete] Delete refs for resource ApiAccessList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objApiAccessList.GetUuid())
	return nil
}

func ResourceApiAccessListSchema() map[string]*schema.Schema {
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
		"api_access_list_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRbacRuleEntriesType(),
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

func ResourceApiAccessListRefsSchema() map[string]*schema.Schema {
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

func ResourceApiAccessList() *schema.Resource {
	return &schema.Resource{
		Create: ResourceApiAccessListCreate,
		Read:   ResourceApiAccessListRead,
		Update: ResourceApiAccessListUpdate,
		Delete: ResourceApiAccessListDelete,
		Schema: ResourceApiAccessListSchema(),
	}
}

func ResourceApiAccessListRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceApiAccessListRefsCreate,
		Read:   ResourceApiAccessListRefsRead,
		Update: ResourceApiAccessListRefsUpdate,
		Delete: ResourceApiAccessListRefsDelete,
		Schema: ResourceApiAccessListRefsSchema(),
	}
}

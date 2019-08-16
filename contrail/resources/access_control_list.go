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

func SetAccessControlListFromResource(object *AccessControlList, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetAccessControlListFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("access_control_list_entries"); ok {
		member := new(AclEntriesType)
		SetAclEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetAccessControlListEntries(member)
	}
	if val, ok := d.GetOk("access_control_list_hash"); ok {
		object.SetAccessControlListHash(val.(uint64))
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

func SetRefsAccessControlListFromResource(object *AccessControlList, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsAccessControlListFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsAccessControlListFromResource(object *AccessControlList, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsAccessControlListFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteAccessControlListToResource(object AccessControlList, d *schema.ResourceData, m interface{}) {

	access_control_list_entriesObj := object.GetAccessControlListEntries()
	d.Set("access_control_list_entries", TakeAclEntriesTypeAsMap(&access_control_list_entriesObj))
	d.Set("access_control_list_hash", object.GetAccessControlListHash())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeAccessControlListAsMap(object *AccessControlList) map[string]interface{} {
	omap := make(map[string]interface{})

	access_control_list_entriesObj := object.GetAccessControlListEntries()
	omap["access_control_list_entries"] = TakeAclEntriesTypeAsMap(&access_control_list_entriesObj)
	omap["access_control_list_hash"] = object.GetAccessControlListHash()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateAccessControlListFromResource(object *AccessControlList, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("access_control_list_entries") {
		if val, ok := d.GetOk("access_control_list_entries"); ok {
			member := new(AclEntriesType)
			SetAclEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetAccessControlListEntries(member)
		}
	}
	if d.HasChange("access_control_list_hash") {
		if val, ok := d.GetOk("access_control_list_hash"); ok {
			object.SetAccessControlListHash(val.(uint64))
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

func ResourceAccessControlListCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAccessControlListCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(AccessControlList)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceAccessControlListCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "AccessControlList", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetAccessControlListFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceAccessControlListCreate] Creation of resource AccessControlList on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceAccessControlListRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAccessControlListRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAccessControlListRefsCreate] Missing 'uuid' field for resource AccessControlList")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("access-control-list", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAccessControlListRefsCreate] Retrieving AccessControlList with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAccessControlList := obj.(*AccessControlList) // Fully set by Contrail backend
	if err := SetRefsAccessControlListFromResource(objAccessControlList, d, m); err != nil {
		return fmt.Errorf("[ResourceAccessControlListRefsCreate] Set refs on object AccessControlList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAccessControlList.GetHref())
	if err := client.Update(objAccessControlList); err != nil {
		return fmt.Errorf("[ResourceAccessControlListRefsCreate] Update refs for resource AccessControlList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAccessControlList.GetUuid())
	return nil
}

func ResourceAccessControlListRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAccessControlListREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("access-control-list", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAccessControlListRead] Read resource access-control-list on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AccessControlList)
	WriteAccessControlListToResource(*object, d, m)
	return nil
}

func ResourceAccessControlListRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAccessControlListRefsREAD")
	return nil
}

func ResourceAccessControlListUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAccessControlListUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("access-control-list", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAccessControlListResourceUpdate] Retrieving AccessControlList with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AccessControlList)
	UpdateAccessControlListFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAccessControlListUpdate] Update of resource AccessControlList on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAccessControlListRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAccessControlListRefsUpdate")
	return nil
}

func ResourceAccessControlListDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAccessControlListDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("access-control-list", d.Id()); err != nil {
		return fmt.Errorf("[ResourceAccessControlListDelete] Deletion of resource access-control-list on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceAccessControlListRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAccessControlListRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAccessControlListRefsDelete] Missing 'uuid' field for resource AccessControlList")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("access-control-list", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAccessControlListRefsDelete] Retrieving AccessControlList with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAccessControlList := obj.(*AccessControlList) // Fully set by Contrail backend
	if err := DeleteRefsAccessControlListFromResource(objAccessControlList, d, m); err != nil {
		return fmt.Errorf("[ResourceAccessControlListRefsDelete] Set refs on object AccessControlList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAccessControlList.GetHref())
	if err := client.Update(objAccessControlList); err != nil {
		return fmt.Errorf("[ResourceAccessControlListRefsDelete] Delete refs for resource AccessControlList (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAccessControlList.GetUuid())
	return nil
}

func ResourceAccessControlListSchema() map[string]*schema.Schema {
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
		"access_control_list_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAclEntriesType(),
		},
		"access_control_list_hash": &schema.Schema{
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

func ResourceAccessControlListRefsSchema() map[string]*schema.Schema {
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

func ResourceAccessControlList() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAccessControlListCreate,
		Read:   ResourceAccessControlListRead,
		Update: ResourceAccessControlListUpdate,
		Delete: ResourceAccessControlListDelete,
		Schema: ResourceAccessControlListSchema(),
	}
}

func ResourceAccessControlListRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAccessControlListRefsCreate,
		Read:   ResourceAccessControlListRefsRead,
		Update: ResourceAccessControlListRefsUpdate,
		Delete: ResourceAccessControlListRefsDelete,
		Schema: ResourceAccessControlListRefsSchema(),
	}
}

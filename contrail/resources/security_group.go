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

func SetSecurityGroupFromResource(object *SecurityGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetSecurityGroupFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("security_group_id"); ok {
		object.SetSecurityGroupId(val.(string))
	}
	if val, ok := d.GetOk("configured_security_group_id"); ok {
		object.SetConfiguredSecurityGroupId(val.(int))
	}
	if val, ok := d.GetOk("security_group_entries"); ok {
		member := new(PolicyEntriesType)
		SetPolicyEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetSecurityGroupEntries(member)
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

func SetRefsSecurityGroupFromResource(object *SecurityGroup, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsSecurityGroupFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteSecurityGroupToResource(object SecurityGroup, d *schema.ResourceData, m interface{}) {

	d.Set("security_group_id", object.GetSecurityGroupId())
	d.Set("configured_security_group_id", object.GetConfiguredSecurityGroupId())
	security_group_entriesObj := object.GetSecurityGroupEntries()
	d.Set("security_group_entries", TakePolicyEntriesTypeAsMap(&security_group_entriesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeSecurityGroupAsMap(object *SecurityGroup) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["security_group_id"] = object.GetSecurityGroupId()
	omap["configured_security_group_id"] = object.GetConfiguredSecurityGroupId()
	security_group_entriesObj := object.GetSecurityGroupEntries()
	omap["security_group_entries"] = TakePolicyEntriesTypeAsMap(&security_group_entriesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateSecurityGroupFromResource(object *SecurityGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("security_group_id") {
		if val, ok := d.GetOk("security_group_id"); ok {
			object.SetSecurityGroupId(val.(string))
		}
	}
	if d.HasChange("configured_security_group_id") {
		if val, ok := d.GetOk("configured_security_group_id"); ok {
			object.SetConfiguredSecurityGroupId(val.(int))
		}
	}
	if d.HasChange("security_group_entries") {
		if val, ok := d.GetOk("security_group_entries"); ok {
			member := new(PolicyEntriesType)
			SetPolicyEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetSecurityGroupEntries(member)
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

func ResourceSecurityGroupCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceSecurityGroupCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(SecurityGroup)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceSecurityGroupCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "SecurityGroup", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetSecurityGroupFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceSecurityGroupCreate] Creation of resource SecurityGroup on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceSecurityGroupRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceSecurityGroupRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceSecurityGroupRefsCreate] Missing 'uuid' field for resource SecurityGroup")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("security-group", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceSecurityGroupRefsCreate] Retrieving SecurityGroup with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objSecurityGroup := obj.(*SecurityGroup) // Fully set by Contrail backend
	if err := SetRefsSecurityGroupFromResource(objSecurityGroup, d, m); err != nil {
		return fmt.Errorf("[ResourceSecurityGroupRefsCreate] Set refs on object SecurityGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objSecurityGroup.GetHref())
	if err := client.Update(objSecurityGroup); err != nil {
		return fmt.Errorf("[ResourceSecurityGroupRefsCreate] Update refs for resource SecurityGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objSecurityGroup.GetUuid())
	return nil
}

func ResourceSecurityGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityGroupREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("security-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceSecurityGroupRead] Read resource security-group on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*SecurityGroup)
	WriteSecurityGroupToResource(*object, d, m)
	return nil
}

func ResourceSecurityGroupRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityGroupRefsREAD")
	return nil
}

func ResourceSecurityGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityGroupUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("security-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceSecurityGroupResourceUpdate] Retrieving SecurityGroup with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*SecurityGroup)
	UpdateSecurityGroupFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceSecurityGroupUpdate] Update of resource SecurityGroup on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceSecurityGroupRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityGroupRefsUpdate")
	return nil
}

func ResourceSecurityGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityGroupDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("security-group", d.Id()); err != nil {
		return fmt.Errorf("[ResourceSecurityGroupDelete] Deletion of resource security-group on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceSecurityGroupRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityGroupRefsDelete: %v", d.Id())
	return nil
}

func ResourceSecurityGroupSchema() map[string]*schema.Schema {
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
		"security_group_id": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"configured_security_group_id": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"security_group_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePolicyEntriesType(),
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

func ResourceSecurityGroupRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSecurityGroupCreate,
		Read:   ResourceSecurityGroupRead,
		Update: ResourceSecurityGroupUpdate,
		Delete: ResourceSecurityGroupDelete,
		Schema: ResourceSecurityGroupSchema(),
	}
}

func ResourceSecurityGroupRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSecurityGroupRefsCreate,
		Read:   ResourceSecurityGroupRefsRead,
		Update: ResourceSecurityGroupRefsUpdate,
		Delete: ResourceSecurityGroupRefsDelete,
		Schema: ResourceSecurityGroupRefsSchema(),
	}
}

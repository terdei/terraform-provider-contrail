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

func SetPolicyManagementFromResource(object *PolicyManagement, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetPolicyManagementFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsPolicyManagementFromResource(object *PolicyManagement, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsPolicyManagementFromResource] key = %v, prefix = %v", key, prefix)
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

func WritePolicyManagementToResource(object PolicyManagement, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakePolicyManagementAsMap(object *PolicyManagement) map[string]interface{} {
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

func UpdatePolicyManagementFromResource(object *PolicyManagement, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourcePolicyManagementCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePolicyManagementCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(PolicyManagement)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourcePolicyManagementCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "PolicyManagement", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetPolicyManagementFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourcePolicyManagementCreate] Creation of resource PolicyManagement on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourcePolicyManagementRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePolicyManagementRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePolicyManagementRefsCreate] Missing 'uuid' field for resource PolicyManagement")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("policy-management", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePolicyManagementRefsCreate] Retrieving PolicyManagement with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPolicyManagement := obj.(*PolicyManagement) // Fully set by Contrail backend
	if err := SetRefsPolicyManagementFromResource(objPolicyManagement, d, m); err != nil {
		return fmt.Errorf("[ResourcePolicyManagementRefsCreate] Set refs on object PolicyManagement (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPolicyManagement.GetHref())
	if err := client.Update(objPolicyManagement); err != nil {
		return fmt.Errorf("[ResourcePolicyManagementRefsCreate] Update refs for resource PolicyManagement (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPolicyManagement.GetUuid())
	return nil
}

func ResourcePolicyManagementRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePolicyManagementREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("policy-management", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePolicyManagementRead] Read resource policy-management on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*PolicyManagement)
	WritePolicyManagementToResource(*object, d, m)
	return nil
}

func ResourcePolicyManagementRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePolicyManagementRefsREAD")
	return nil
}

func ResourcePolicyManagementUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePolicyManagementUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("policy-management", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePolicyManagementResourceUpdate] Retrieving PolicyManagement with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*PolicyManagement)
	UpdatePolicyManagementFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourcePolicyManagementUpdate] Update of resource PolicyManagement on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourcePolicyManagementRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePolicyManagementRefsUpdate")
	return nil
}

func ResourcePolicyManagementDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePolicyManagementDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("policy-management", d.Id()); err != nil {
		return fmt.Errorf("[ResourcePolicyManagementDelete] Deletion of resource policy-management on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourcePolicyManagementRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePolicyManagementRefsDelete: %v", d.Id())
	return nil
}

func ResourcePolicyManagementSchema() map[string]*schema.Schema {
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

func ResourcePolicyManagementRefsSchema() map[string]*schema.Schema {
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

func ResourcePolicyManagement() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePolicyManagementCreate,
		Read:   ResourcePolicyManagementRead,
		Update: ResourcePolicyManagementUpdate,
		Delete: ResourcePolicyManagementDelete,
		Schema: ResourcePolicyManagementSchema(),
	}
}

func ResourcePolicyManagementRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePolicyManagementRefsCreate,
		Read:   ResourcePolicyManagementRefsRead,
		Update: ResourcePolicyManagementRefsUpdate,
		Delete: ResourcePolicyManagementRefsDelete,
		Schema: ResourcePolicyManagementRefsSchema(),
	}
}

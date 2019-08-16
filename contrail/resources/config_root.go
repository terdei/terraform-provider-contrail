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

func SetConfigRootFromResource(object *ConfigRoot, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetConfigRootFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsConfigRootFromResource(object *ConfigRoot, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsConfigRootFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsConfigRootFromResource(object *ConfigRoot, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsConfigRootFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteConfigRootToResource(object ConfigRoot, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeConfigRootAsMap(object *ConfigRoot) map[string]interface{} {
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

func UpdateConfigRootFromResource(object *ConfigRoot, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceConfigRootCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceConfigRootCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ConfigRoot)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceConfigRootCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ConfigRoot", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetConfigRootFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceConfigRootCreate] Creation of resource ConfigRoot on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceConfigRootRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceConfigRootRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceConfigRootRefsCreate] Missing 'uuid' field for resource ConfigRoot")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("config-root", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceConfigRootRefsCreate] Retrieving ConfigRoot with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objConfigRoot := obj.(*ConfigRoot) // Fully set by Contrail backend
	if err := SetRefsConfigRootFromResource(objConfigRoot, d, m); err != nil {
		return fmt.Errorf("[ResourceConfigRootRefsCreate] Set refs on object ConfigRoot (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objConfigRoot.GetHref())
	if err := client.Update(objConfigRoot); err != nil {
		return fmt.Errorf("[ResourceConfigRootRefsCreate] Update refs for resource ConfigRoot (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objConfigRoot.GetUuid())
	return nil
}

func ResourceConfigRootRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigRootREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("config-root", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceConfigRootRead] Read resource config-root on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ConfigRoot)
	WriteConfigRootToResource(*object, d, m)
	return nil
}

func ResourceConfigRootRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigRootRefsREAD")
	return nil
}

func ResourceConfigRootUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigRootUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("config-root", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceConfigRootResourceUpdate] Retrieving ConfigRoot with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ConfigRoot)
	UpdateConfigRootFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceConfigRootUpdate] Update of resource ConfigRoot on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceConfigRootRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigRootRefsUpdate")
	return nil
}

func ResourceConfigRootDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigRootDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("config-root", d.Id()); err != nil {
		return fmt.Errorf("[ResourceConfigRootDelete] Deletion of resource config-root on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceConfigRootRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceConfigRootRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceConfigRootRefsDelete] Missing 'uuid' field for resource ConfigRoot")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("config-root", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceConfigRootRefsDelete] Retrieving ConfigRoot with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objConfigRoot := obj.(*ConfigRoot) // Fully set by Contrail backend
	if err := DeleteRefsConfigRootFromResource(objConfigRoot, d, m); err != nil {
		return fmt.Errorf("[ResourceConfigRootRefsDelete] Set refs on object ConfigRoot (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objConfigRoot.GetHref())
	if err := client.Update(objConfigRoot); err != nil {
		return fmt.Errorf("[ResourceConfigRootRefsDelete] Delete refs for resource ConfigRoot (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objConfigRoot.GetUuid())
	return nil
}

func ResourceConfigRootSchema() map[string]*schema.Schema {
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

func ResourceConfigRootRefsSchema() map[string]*schema.Schema {
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

func ResourceConfigRoot() *schema.Resource {
	return &schema.Resource{
		Create: ResourceConfigRootCreate,
		Read:   ResourceConfigRootRead,
		Update: ResourceConfigRootUpdate,
		Delete: ResourceConfigRootDelete,
		Schema: ResourceConfigRootSchema(),
	}
}

func ResourceConfigRootRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceConfigRootRefsCreate,
		Read:   ResourceConfigRootRefsRead,
		Update: ResourceConfigRootRefsUpdate,
		Delete: ResourceConfigRootRefsDelete,
		Schema: ResourceConfigRootRefsSchema(),
	}
}

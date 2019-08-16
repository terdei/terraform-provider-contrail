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

func SetAliasIpPoolFromResource(object *AliasIpPool, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetAliasIpPoolFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsAliasIpPoolFromResource(object *AliasIpPool, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsAliasIpPoolFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsAliasIpPoolFromResource(object *AliasIpPool, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsAliasIpPoolFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteAliasIpPoolToResource(object AliasIpPool, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeAliasIpPoolAsMap(object *AliasIpPool) map[string]interface{} {
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

func UpdateAliasIpPoolFromResource(object *AliasIpPool, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceAliasIpPoolCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAliasIpPoolCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(AliasIpPool)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceAliasIpPoolCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "AliasIpPool", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetAliasIpPoolFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolCreate] Creation of resource AliasIpPool on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceAliasIpPoolRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAliasIpPoolRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAliasIpPoolRefsCreate] Missing 'uuid' field for resource AliasIpPool")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("alias-ip-pool", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRefsCreate] Retrieving AliasIpPool with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAliasIpPool := obj.(*AliasIpPool) // Fully set by Contrail backend
	if err := SetRefsAliasIpPoolFromResource(objAliasIpPool, d, m); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRefsCreate] Set refs on object AliasIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAliasIpPool.GetHref())
	if err := client.Update(objAliasIpPool); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRefsCreate] Update refs for resource AliasIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAliasIpPool.GetUuid())
	return nil
}

func ResourceAliasIpPoolRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpPoolREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("alias-ip-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRead] Read resource alias-ip-pool on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AliasIpPool)
	WriteAliasIpPoolToResource(*object, d, m)
	return nil
}

func ResourceAliasIpPoolRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpPoolRefsREAD")
	return nil
}

func ResourceAliasIpPoolUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpPoolUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("alias-ip-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolResourceUpdate] Retrieving AliasIpPool with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AliasIpPool)
	UpdateAliasIpPoolFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolUpdate] Update of resource AliasIpPool on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAliasIpPoolRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpPoolRefsUpdate")
	return nil
}

func ResourceAliasIpPoolDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpPoolDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("alias-ip-pool", d.Id()); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolDelete] Deletion of resource alias-ip-pool on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceAliasIpPoolRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAliasIpPoolRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAliasIpPoolRefsDelete] Missing 'uuid' field for resource AliasIpPool")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("alias-ip-pool", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRefsDelete] Retrieving AliasIpPool with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAliasIpPool := obj.(*AliasIpPool) // Fully set by Contrail backend
	if err := DeleteRefsAliasIpPoolFromResource(objAliasIpPool, d, m); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRefsDelete] Set refs on object AliasIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAliasIpPool.GetHref())
	if err := client.Update(objAliasIpPool); err != nil {
		return fmt.Errorf("[ResourceAliasIpPoolRefsDelete] Delete refs for resource AliasIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAliasIpPool.GetUuid())
	return nil
}

func ResourceAliasIpPoolSchema() map[string]*schema.Schema {
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

func ResourceAliasIpPoolRefsSchema() map[string]*schema.Schema {
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

func ResourceAliasIpPool() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAliasIpPoolCreate,
		Read:   ResourceAliasIpPoolRead,
		Update: ResourceAliasIpPoolUpdate,
		Delete: ResourceAliasIpPoolDelete,
		Schema: ResourceAliasIpPoolSchema(),
	}
}

func ResourceAliasIpPoolRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAliasIpPoolRefsCreate,
		Read:   ResourceAliasIpPoolRefsRead,
		Update: ResourceAliasIpPoolRefsUpdate,
		Delete: ResourceAliasIpPoolRefsDelete,
		Schema: ResourceAliasIpPoolRefsSchema(),
	}
}

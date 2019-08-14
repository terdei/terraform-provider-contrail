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

func SetPortTupleFromResource(object *PortTuple, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetPortTupleFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsPortTupleFromResource(object *PortTuple, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsPortTupleFromResource] key = %v, prefix = %v", key, prefix)
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

func WritePortTupleToResource(object PortTuple, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakePortTupleAsMap(object *PortTuple) map[string]interface{} {
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

func UpdatePortTupleFromResource(object *PortTuple, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourcePortTupleCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePortTupleCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(PortTuple)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourcePortTupleCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "PortTuple", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetPortTupleFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourcePortTupleCreate] Creation of resource PortTuple on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourcePortTupleRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePortTupleRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePortTupleRefsCreate] Missing 'uuid' field for resource PortTuple")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("port-tuple", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePortTupleRefsCreate] Retrieving PortTuple with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPortTuple := obj.(*PortTuple) // Fully set by Contrail backend
	if err := SetRefsPortTupleFromResource(objPortTuple, d, m); err != nil {
		return fmt.Errorf("[ResourcePortTupleRefsCreate] Set refs on object PortTuple (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPortTuple.GetHref())
	if err := client.Update(objPortTuple); err != nil {
		return fmt.Errorf("[ResourcePortTupleRefsCreate] Update refs for resource PortTuple (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPortTuple.GetUuid())
	return nil
}

func ResourcePortTupleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePortTupleREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("port-tuple", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePortTupleRead] Read resource port-tuple on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*PortTuple)
	WritePortTupleToResource(*object, d, m)
	return nil
}

func ResourcePortTupleRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePortTupleRefsREAD")
	return nil
}

func ResourcePortTupleUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePortTupleUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("port-tuple", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePortTupleResourceUpdate] Retrieving PortTuple with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*PortTuple)
	UpdatePortTupleFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourcePortTupleUpdate] Update of resource PortTuple on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourcePortTupleRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePortTupleRefsUpdate")
	return nil
}

func ResourcePortTupleDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePortTupleDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("port-tuple", d.Id()); err != nil {
		return fmt.Errorf("[ResourcePortTupleDelete] Deletion of resource port-tuple on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourcePortTupleRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePortTupleRefsDelete: %v", d.Id())
	return nil
}

func ResourcePortTupleSchema() map[string]*schema.Schema {
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

func ResourcePortTupleRefsSchema() map[string]*schema.Schema {
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

func ResourcePortTuple() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePortTupleCreate,
		Read:   ResourcePortTupleRead,
		Update: ResourcePortTupleUpdate,
		Delete: ResourcePortTupleDelete,
		Schema: ResourcePortTupleSchema(),
	}
}

func ResourcePortTupleRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePortTupleRefsCreate,
		Read:   ResourcePortTupleRefsRead,
		Update: ResourcePortTupleRefsUpdate,
		Delete: ResourcePortTupleRefsDelete,
		Schema: ResourcePortTupleRefsSchema(),
	}
}

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

func SetNamespaceFromResource(object *Namespace, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetNamespaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("namespace_cidr"); ok {
		member := new(SubnetType)
		SetSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetNamespaceCidr(member)
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

func SetRefsNamespaceFromResource(object *Namespace, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsNamespaceFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteNamespaceToResource(object Namespace, d *schema.ResourceData, m interface{}) {

	namespace_cidrObj := object.GetNamespaceCidr()
	d.Set("namespace_cidr", TakeSubnetTypeAsMap(&namespace_cidrObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeNamespaceAsMap(object *Namespace) map[string]interface{} {
	omap := make(map[string]interface{})

	namespace_cidrObj := object.GetNamespaceCidr()
	omap["namespace_cidr"] = TakeSubnetTypeAsMap(&namespace_cidrObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateNamespaceFromResource(object *Namespace, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("namespace_cidr") {
		if val, ok := d.GetOk("namespace_cidr"); ok {
			member := new(SubnetType)
			SetSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetNamespaceCidr(member)
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

func ResourceNamespaceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNamespaceCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Namespace)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceNamespaceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Namespace", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetNamespaceFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceNamespaceCreate] Creation of resource Namespace on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceNamespaceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNamespaceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNamespaceRefsCreate] Missing 'uuid' field for resource Namespace")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("namespace", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNamespaceRefsCreate] Retrieving Namespace with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNamespace := obj.(*Namespace) // Fully set by Contrail backend
	if err := SetRefsNamespaceFromResource(objNamespace, d, m); err != nil {
		return fmt.Errorf("[ResourceNamespaceRefsCreate] Set refs on object Namespace (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNamespace.GetHref())
	if err := client.Update(objNamespace); err != nil {
		return fmt.Errorf("[ResourceNamespaceRefsCreate] Update refs for resource Namespace (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNamespace.GetUuid())
	return nil
}

func ResourceNamespaceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNamespaceREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("namespace", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNamespaceRead] Read resource namespace on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Namespace)
	WriteNamespaceToResource(*object, d, m)
	return nil
}

func ResourceNamespaceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNamespaceRefsREAD")
	return nil
}

func ResourceNamespaceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNamespaceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("namespace", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNamespaceResourceUpdate] Retrieving Namespace with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Namespace)
	UpdateNamespaceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceNamespaceUpdate] Update of resource Namespace on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceNamespaceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNamespaceRefsUpdate")
	return nil
}

func ResourceNamespaceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNamespaceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("namespace", d.Id()); err != nil {
		return fmt.Errorf("[ResourceNamespaceDelete] Deletion of resource namespace on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceNamespaceRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNamespaceRefsDelete: %v", d.Id())
	return nil
}

func ResourceNamespaceSchema() map[string]*schema.Schema {
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
		"namespace_cidr": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
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

func ResourceNamespaceRefsSchema() map[string]*schema.Schema {
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

func ResourceNamespace() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNamespaceCreate,
		Read:   ResourceNamespaceRead,
		Update: ResourceNamespaceUpdate,
		Delete: ResourceNamespaceDelete,
		Schema: ResourceNamespaceSchema(),
	}
}

func ResourceNamespaceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNamespaceRefsCreate,
		Read:   ResourceNamespaceRefsRead,
		Update: ResourceNamespaceRefsUpdate,
		Delete: ResourceNamespaceRefsDelete,
		Schema: ResourceNamespaceRefsSchema(),
	}
}

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

func SetAliasIpFromResource(object *AliasIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetAliasIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("alias_ip_address"); ok {
		object.SetAliasIpAddress(val.(string))
	}
	if val, ok := d.GetOk("alias_ip_address_family"); ok {
		object.SetAliasIpAddressFamily(val.(string))
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

func SetRefsAliasIpFromResource(object *AliasIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsAliasIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("project_refs"); ok {
		log.Printf("Got ref project_refs -- will call: object.AddProject(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("project", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving project by Uuid = %v as ref for Project on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddProject(refObj.(*Project))
		}
	}
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine-interface by Uuid = %v as ref for VirtualMachineInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
		}
	}
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

func WriteAliasIpToResource(object AliasIp, d *schema.ResourceData, m interface{}) {

	d.Set("alias_ip_address", object.GetAliasIpAddress())
	d.Set("alias_ip_address_family", object.GetAliasIpAddressFamily())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeAliasIpAsMap(object *AliasIp) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["alias_ip_address"] = object.GetAliasIpAddress()
	omap["alias_ip_address_family"] = object.GetAliasIpAddressFamily()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateAliasIpFromResource(object *AliasIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("alias_ip_address") {
		if val, ok := d.GetOk("alias_ip_address"); ok {
			object.SetAliasIpAddress(val.(string))
		}
	}
	if d.HasChange("alias_ip_address_family") {
		if val, ok := d.GetOk("alias_ip_address_family"); ok {
			object.SetAliasIpAddressFamily(val.(string))
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

func ResourceAliasIpCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAliasIpCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(AliasIp)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceAliasIpCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "AliasIp", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetAliasIpFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceAliasIpCreate] Creation of resource AliasIp on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceAliasIpRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAliasIpRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAliasIpRefsCreate] Missing 'uuid' field for resource AliasIp")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("alias-ip", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpRefsCreate] Retrieving AliasIp with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAliasIp := obj.(*AliasIp) // Fully set by Contrail backend
	if err := SetRefsAliasIpFromResource(objAliasIp, d, m); err != nil {
		return fmt.Errorf("[ResourceAliasIpRefsCreate] Set refs on object AliasIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAliasIp.GetHref())
	if err := client.Update(objAliasIp); err != nil {
		return fmt.Errorf("[ResourceAliasIpRefsCreate] Update refs for resource AliasIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAliasIp.GetUuid())
	return nil
}

func ResourceAliasIpRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("alias-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpRead] Read resource alias-ip on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AliasIp)
	WriteAliasIpToResource(*object, d, m)
	return nil
}

func ResourceAliasIpRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpRefsREAD")
	return nil
}

func ResourceAliasIpUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("alias-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAliasIpResourceUpdate] Retrieving AliasIp with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AliasIp)
	UpdateAliasIpFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAliasIpUpdate] Update of resource AliasIp on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAliasIpRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpRefsUpdate")
	return nil
}

func ResourceAliasIpDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("alias-ip", d.Id()); err != nil {
		return fmt.Errorf("[ResourceAliasIpDelete] Deletion of resource alias-ip on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceAliasIpRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAliasIpRefsDelete: %v", d.Id())
	return nil
}

func ResourceAliasIpSchema() map[string]*schema.Schema {
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
		"alias_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"alias_ip_address_family": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
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

func ResourceAliasIpRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"project_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceProject(),
		},
		"virtual_machine_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachineInterface(),
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourceAliasIp() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAliasIpCreate,
		Read:   ResourceAliasIpRead,
		Update: ResourceAliasIpUpdate,
		Delete: ResourceAliasIpDelete,
		Schema: ResourceAliasIpSchema(),
	}
}

func ResourceAliasIpRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAliasIpRefsCreate,
		Read:   ResourceAliasIpRefsRead,
		Update: ResourceAliasIpRefsUpdate,
		Delete: ResourceAliasIpRefsDelete,
		Schema: ResourceAliasIpRefsSchema(),
	}
}

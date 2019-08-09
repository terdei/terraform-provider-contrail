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

func SetVirtualDnsFromResource(object *VirtualDns, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetVirtualDnsFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_dns_data"); ok {
		member := new(VirtualDnsType)
		SetVirtualDnsTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualDnsData(member)
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

func SetRefsVirtualDnsFromResource(object *VirtualDns, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualDnsFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteVirtualDnsToResource(object VirtualDns, d *schema.ResourceData, m interface{}) {

	virtual_dns_dataObj := object.GetVirtualDnsData()
	d.Set("virtual_dns_data", TakeVirtualDnsTypeAsMap(&virtual_dns_dataObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeVirtualDnsAsMap(object *VirtualDns) map[string]interface{} {
	omap := make(map[string]interface{})

	virtual_dns_dataObj := object.GetVirtualDnsData()
	omap["virtual_dns_data"] = TakeVirtualDnsTypeAsMap(&virtual_dns_dataObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateVirtualDnsFromResource(object *VirtualDns, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("virtual_dns_data") {
		if val, ok := d.GetOk("virtual_dns_data"); ok {
			member := new(VirtualDnsType)
			SetVirtualDnsTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualDnsData(member)
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

func ResourceVirtualDnsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualDnsCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(VirtualDns)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceVirtualDnsCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "VirtualDns", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetVirtualDnsFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsCreate] Creation of resource VirtualDns on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceVirtualDnsRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualDnsRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualDnsRefsCreate] Missing 'uuid' field for resource VirtualDns")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-DNS", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRefsCreate] Retrieving VirtualDns with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualDns := obj.(*VirtualDns) // Fully set by Contrail backend
	if err := SetRefsVirtualDnsFromResource(objVirtualDns, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRefsCreate] Set refs on object VirtualDns (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualDns.GetHref())
	if err := client.Update(objVirtualDns); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRefsCreate] Update refs for resource VirtualDns (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualDns.GetUuid())
	return nil
}

func ResourceVirtualDnsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-DNS", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRead] Read resource virtual-DNS on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualDns)
	WriteVirtualDnsToResource(*object, d, m)
	return nil
}

func ResourceVirtualDnsRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRefsREAD")
	return nil
}

func ResourceVirtualDnsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-DNS", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsResourceUpdate] Retrieving VirtualDns with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualDns)
	UpdateVirtualDnsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsUpdate] Update of resource VirtualDns on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualDnsRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRefsUpdate")
	return nil
}

func ResourceVirtualDnsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("virtual-DNS", d.Id()); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsDelete] Deletion of resource virtual-DNS on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceVirtualDnsRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRefsDelete: %v", d.Id())
	return nil
}

func ResourceVirtualDnsSchema() map[string]*schema.Schema {
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
		"virtual_dns_data": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualDnsType(),
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

func ResourceVirtualDnsRefsSchema() map[string]*schema.Schema {
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

func ResourceVirtualDns() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualDnsCreate,
		Read:   ResourceVirtualDnsRead,
		Update: ResourceVirtualDnsUpdate,
		Delete: ResourceVirtualDnsDelete,
		Schema: ResourceVirtualDnsSchema(),
	}
}

func ResourceVirtualDnsRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualDnsRefsCreate,
		Read:   ResourceVirtualDnsRefsRead,
		Update: ResourceVirtualDnsRefsUpdate,
		Delete: ResourceVirtualDnsRefsDelete,
		Schema: ResourceVirtualDnsRefsSchema(),
	}
}

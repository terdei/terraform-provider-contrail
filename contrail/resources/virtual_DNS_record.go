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

func SetVirtualDnsRecordFromResource(object *VirtualDnsRecord, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetVirtualDnsRecordFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_dns_record_data"); ok {
		member := new(VirtualDnsRecordType)
		SetVirtualDnsRecordTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualDnsRecordData(member)
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

func SetRefsVirtualDnsRecordFromResource(object *VirtualDnsRecord, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualDnsRecordFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsVirtualDnsRecordFromResource(object *VirtualDnsRecord, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsVirtualDnsRecordFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteVirtualDnsRecordToResource(object VirtualDnsRecord, d *schema.ResourceData, m interface{}) {

	virtual_dns_record_dataObj := object.GetVirtualDnsRecordData()
	d.Set("virtual_dns_record_data", TakeVirtualDnsRecordTypeAsMap(&virtual_dns_record_dataObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeVirtualDnsRecordAsMap(object *VirtualDnsRecord) map[string]interface{} {
	omap := make(map[string]interface{})

	virtual_dns_record_dataObj := object.GetVirtualDnsRecordData()
	omap["virtual_dns_record_data"] = TakeVirtualDnsRecordTypeAsMap(&virtual_dns_record_dataObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateVirtualDnsRecordFromResource(object *VirtualDnsRecord, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("virtual_dns_record_data") {
		if val, ok := d.GetOk("virtual_dns_record_data"); ok {
			member := new(VirtualDnsRecordType)
			SetVirtualDnsRecordTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualDnsRecordData(member)
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

func ResourceVirtualDnsRecordCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualDnsRecordCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(VirtualDnsRecord)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceVirtualDnsRecordCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "VirtualDnsRecord", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetVirtualDnsRecordFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordCreate] Creation of resource VirtualDnsRecord on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceVirtualDnsRecordRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualDnsRecordRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsCreate] Missing 'uuid' field for resource VirtualDnsRecord")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-DNS-record", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsCreate] Retrieving VirtualDnsRecord with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualDnsRecord := obj.(*VirtualDnsRecord) // Fully set by Contrail backend
	if err := SetRefsVirtualDnsRecordFromResource(objVirtualDnsRecord, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsCreate] Set refs on object VirtualDnsRecord (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualDnsRecord.GetHref())
	if err := client.Update(objVirtualDnsRecord); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsCreate] Update refs for resource VirtualDnsRecord (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualDnsRecord.GetUuid())
	return nil
}

func ResourceVirtualDnsRecordRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRecordREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-DNS-record", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRead] Read resource virtual-DNS-record on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualDnsRecord)
	WriteVirtualDnsRecordToResource(*object, d, m)
	return nil
}

func ResourceVirtualDnsRecordRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRecordRefsREAD")
	return nil
}

func ResourceVirtualDnsRecordUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRecordUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-DNS-record", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordResourceUpdate] Retrieving VirtualDnsRecord with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualDnsRecord)
	UpdateVirtualDnsRecordFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordUpdate] Update of resource VirtualDnsRecord on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualDnsRecordRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRecordRefsUpdate")
	return nil
}

func ResourceVirtualDnsRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualDnsRecordDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("virtual-DNS-record", d.Id()); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordDelete] Deletion of resource virtual-DNS-record on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceVirtualDnsRecordRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualDnsRecordRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsDelete] Missing 'uuid' field for resource VirtualDnsRecord")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-DNS-record", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsDelete] Retrieving VirtualDnsRecord with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualDnsRecord := obj.(*VirtualDnsRecord) // Fully set by Contrail backend
	if err := DeleteRefsVirtualDnsRecordFromResource(objVirtualDnsRecord, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsDelete] Set refs on object VirtualDnsRecord (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualDnsRecord.GetHref())
	if err := client.Update(objVirtualDnsRecord); err != nil {
		return fmt.Errorf("[ResourceVirtualDnsRecordRefsDelete] Delete refs for resource VirtualDnsRecord (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualDnsRecord.GetUuid())
	return nil
}

func ResourceVirtualDnsRecordSchema() map[string]*schema.Schema {
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
		"virtual_dns_record_data": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualDnsRecordType(),
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

func ResourceVirtualDnsRecordRefsSchema() map[string]*schema.Schema {
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

func ResourceVirtualDnsRecord() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualDnsRecordCreate,
		Read:   ResourceVirtualDnsRecordRead,
		Update: ResourceVirtualDnsRecordUpdate,
		Delete: ResourceVirtualDnsRecordDelete,
		Schema: ResourceVirtualDnsRecordSchema(),
	}
}

func ResourceVirtualDnsRecordRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualDnsRecordRefsCreate,
		Read:   ResourceVirtualDnsRecordRefsRead,
		Update: ResourceVirtualDnsRecordRefsUpdate,
		Delete: ResourceVirtualDnsRecordRefsDelete,
		Schema: ResourceVirtualDnsRecordRefsSchema(),
	}
}

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

func SetAddressGroupFromResource(object *AddressGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetAddressGroupFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("address_group_prefix"); ok {
		member := new(SubnetListType)
		SetSubnetListTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetAddressGroupPrefix(member)
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

func SetRefsAddressGroupFromResource(object *AddressGroup, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsAddressGroupFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsAddressGroupFromResource(object *AddressGroup, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsAddressGroupFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteAddressGroupToResource(object AddressGroup, d *schema.ResourceData, m interface{}) {

	address_group_prefixObj := object.GetAddressGroupPrefix()
	d.Set("address_group_prefix", TakeSubnetListTypeAsMap(&address_group_prefixObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteAddressGroupRefsToResource(object AddressGroup, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetTagRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("tag_refs", refList)
	}
}

func TakeAddressGroupAsMap(object *AddressGroup) map[string]interface{} {
	omap := make(map[string]interface{})

	address_group_prefixObj := object.GetAddressGroupPrefix()
	omap["address_group_prefix"] = TakeSubnetListTypeAsMap(&address_group_prefixObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateAddressGroupFromResource(object *AddressGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("address_group_prefix") {
		if val, ok := d.GetOk("address_group_prefix"); ok {
			member := new(SubnetListType)
			SetSubnetListTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetAddressGroupPrefix(member)
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

func UpdateAddressGroupRefsFromResource(object *AddressGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("tag_refs") {
		object.ClearTag()
		if val, ok := d.GetOk("tag_refs"); ok {
			log.Printf("Got ref tag_refs -- will call: object.AddTag(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("tag", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddTag(refObj.(*Tag))
			}
		}
	}

}

func ResourceAddressGroupCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAddressGroupCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(AddressGroup)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceAddressGroupCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "AddressGroup", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetAddressGroupFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceAddressGroupCreate] Creation of resource AddressGroup on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceAddressGroupRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAddressGroupRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAddressGroupRefsCreate] Missing 'uuid' field for resource AddressGroup")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("address-group", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsCreate] Retrieving AddressGroup with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAddressGroup := obj.(*AddressGroup) // Fully set by Contrail backend
	if err := SetRefsAddressGroupFromResource(objAddressGroup, d, m); err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsCreate] Set refs on object AddressGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAddressGroup.GetHref())
	if err := client.Update(objAddressGroup); err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsCreate] Update refs for resource AddressGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAddressGroup.GetUuid())
	return nil
}

func ResourceAddressGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAddressGroupRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("address-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAddressGroupRead] Read resource address-group on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AddressGroup)
	WriteAddressGroupToResource(*object, d, m)
	return nil
}

func ResourceAddressGroupRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAddressGroupRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("address-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsRead] Read resource address-group on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AddressGroup)
	WriteAddressGroupRefsToResource(*object, d, m)
	return nil
}

func ResourceAddressGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAddressGroupUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("address-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAddressGroupUpdate] Retrieving AddressGroup with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AddressGroup)
	UpdateAddressGroupFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAddressGroupUpdate] Update of resource AddressGroup on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAddressGroupRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAddressGroupRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("address-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsUpdate] Retrieving AddressGroup with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AddressGroup)
	UpdateAddressGroupRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsUpdate] Update of resource AddressGroup on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAddressGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAddressGroupDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("address-group", d.Id()); err != nil {
		return fmt.Errorf("[ResourceAddressGroupDelete] Deletion of resource address-group on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceAddressGroupRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAddressGroupRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAddressGroupRefsDelete] Missing 'uuid' field for resource AddressGroup")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("address-group", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsDelete] Retrieving AddressGroup with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAddressGroup := obj.(*AddressGroup) // Fully set by Contrail backend
	if err := DeleteRefsAddressGroupFromResource(objAddressGroup, d, m); err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsDelete] Set refs on object AddressGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAddressGroup.GetHref())
	if err := client.Update(objAddressGroup); err != nil {
		return fmt.Errorf("[ResourceAddressGroupRefsDelete] Delete refs for resource AddressGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAddressGroup.GetUuid())
	return nil
}

func ResourceAddressGroupSchema() map[string]*schema.Schema {
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
		"address_group_prefix": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetListType(),
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

func ResourceAddressGroupRefsSchema() map[string]*schema.Schema {
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

func ResourceAddressGroup() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAddressGroupCreate,
		Read:   ResourceAddressGroupRead,
		Update: ResourceAddressGroupUpdate,
		Delete: ResourceAddressGroupDelete,
		Schema: ResourceAddressGroupSchema(),
	}
}

func ResourceAddressGroupRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAddressGroupRefsCreate,
		Read:   ResourceAddressGroupRefsRead,
		Update: ResourceAddressGroupRefsUpdate,
		Delete: ResourceAddressGroupRefsDelete,
		Schema: ResourceAddressGroupRefsSchema(),
	}
}

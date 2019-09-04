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

func SetServiceGroupFromResource(object *ServiceGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceGroupFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_group_firewall_service_list"); ok {
		member := new(FirewallServiceGroupType)
		SetFirewallServiceGroupTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceGroupFirewallServiceList(member)
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

func SetRefsServiceGroupFromResource(object *ServiceGroup, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceGroupFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsServiceGroupFromResource(object *ServiceGroup, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceGroupFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteServiceGroupToResource(object ServiceGroup, d *schema.ResourceData, m interface{}) {

	service_group_firewall_service_listObj := object.GetServiceGroupFirewallServiceList()
	d.Set("service_group_firewall_service_list", TakeFirewallServiceGroupTypeAsMap(&service_group_firewall_service_listObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceGroupRefsToResource(object ServiceGroup, d *schema.ResourceData, m interface{}) {

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

func TakeServiceGroupAsMap(object *ServiceGroup) map[string]interface{} {
	omap := make(map[string]interface{})

	service_group_firewall_service_listObj := object.GetServiceGroupFirewallServiceList()
	omap["service_group_firewall_service_list"] = TakeFirewallServiceGroupTypeAsMap(&service_group_firewall_service_listObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceGroupFromResource(object *ServiceGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("service_group_firewall_service_list") {
		if val, ok := d.GetOk("service_group_firewall_service_list"); ok {
			member := new(FirewallServiceGroupType)
			SetFirewallServiceGroupTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceGroupFirewallServiceList(member)
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

func UpdateServiceGroupRefsFromResource(object *ServiceGroup, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceServiceGroupCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceGroupCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceGroup)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceGroupCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceGroup", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceGroupFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceGroupCreate] Creation of resource ServiceGroup on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceGroupRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceGroupRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceGroupRefsCreate] Missing 'uuid' field for resource ServiceGroup")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-group", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsCreate] Retrieving ServiceGroup with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceGroup := obj.(*ServiceGroup) // Fully set by Contrail backend
	if err := SetRefsServiceGroupFromResource(objServiceGroup, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsCreate] Set refs on object ServiceGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceGroup.GetHref())
	if err := client.Update(objServiceGroup); err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsCreate] Update refs for resource ServiceGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceGroup.GetUuid())
	return nil
}

func ResourceServiceGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceGroupRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceGroupRead] Read resource service-group on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceGroup)
	WriteServiceGroupToResource(*object, d, m)
	return nil
}

func ResourceServiceGroupRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceGroupRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsRead] Read resource service-group on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceGroup)
	WriteServiceGroupRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceGroupUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceGroupUpdate] Retrieving ServiceGroup with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceGroup)
	UpdateServiceGroupFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceGroupUpdate] Update of resource ServiceGroup on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceGroupRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceGroupRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-group", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsUpdate] Retrieving ServiceGroup with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceGroup)
	UpdateServiceGroupRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsUpdate] Update of resource ServiceGroup on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceGroupDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-group", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceGroupDelete] Deletion of resource service-group on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceGroupRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceGroupRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceGroupRefsDelete] Missing 'uuid' field for resource ServiceGroup")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-group", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsDelete] Retrieving ServiceGroup with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceGroup := obj.(*ServiceGroup) // Fully set by Contrail backend
	if err := DeleteRefsServiceGroupFromResource(objServiceGroup, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsDelete] Set refs on object ServiceGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceGroup.GetHref())
	if err := client.Update(objServiceGroup); err != nil {
		return fmt.Errorf("[ResourceServiceGroupRefsDelete] Delete refs for resource ServiceGroup (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceGroup.GetUuid())
	return nil
}

func ResourceServiceGroupSchema() map[string]*schema.Schema {
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
		"service_group_firewall_service_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallServiceGroupType(),
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

func ResourceServiceGroupRefsSchema() map[string]*schema.Schema {
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

func ResourceServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceGroupCreate,
		Read:   ResourceServiceGroupRead,
		Update: ResourceServiceGroupUpdate,
		Delete: ResourceServiceGroupDelete,
		Schema: ResourceServiceGroupSchema(),
	}
}

func ResourceServiceGroupRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceGroupRefsCreate,
		Read:   ResourceServiceGroupRefsRead,
		Update: ResourceServiceGroupRefsUpdate,
		Delete: ResourceServiceGroupRefsDelete,
		Schema: ResourceServiceGroupRefsSchema(),
	}
}

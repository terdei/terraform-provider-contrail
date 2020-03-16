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

func SetServiceApplianceSetFromResource(object *ServiceApplianceSet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceApplianceSetFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_appliance_set_properties"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceApplianceSetProperties(member)
	}
	if val, ok := d.GetOk("service_appliance_driver"); ok {
		object.SetServiceApplianceDriver(val.(string))
	}
	if val, ok := d.GetOk("service_appliance_ha_mode"); ok {
		object.SetServiceApplianceHaMode(val.(string))
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

func SetRefsServiceApplianceSetFromResource(object *ServiceApplianceSet, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceApplianceSetFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsServiceApplianceSetFromResource(object *ServiceApplianceSet, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceApplianceSetFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteServiceApplianceSetToResource(object ServiceApplianceSet, d *schema.ResourceData, m interface{}) {

	service_appliance_set_propertiesObj := object.GetServiceApplianceSetProperties()
	d.Set("service_appliance_set_properties", TakeKeyValuePairsAsMap(&service_appliance_set_propertiesObj))
	d.Set("service_appliance_driver", object.GetServiceApplianceDriver())
	d.Set("service_appliance_ha_mode", object.GetServiceApplianceHaMode())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceApplianceSetRefsToResource(object ServiceApplianceSet, d *schema.ResourceData, m interface{}) {

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

func TakeServiceApplianceSetAsMap(object *ServiceApplianceSet) map[string]interface{} {
	omap := make(map[string]interface{})

	service_appliance_set_propertiesObj := object.GetServiceApplianceSetProperties()
	omap["service_appliance_set_properties"] = TakeKeyValuePairsAsMap(&service_appliance_set_propertiesObj)
	omap["service_appliance_driver"] = object.GetServiceApplianceDriver()
	omap["service_appliance_ha_mode"] = object.GetServiceApplianceHaMode()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceApplianceSetFromResource(object *ServiceApplianceSet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("service_appliance_set_properties") {
		if val, ok := d.GetOk("service_appliance_set_properties"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceApplianceSetProperties(member)
		}
	}
	if d.HasChange("service_appliance_driver") {
		if val, ok := d.GetOk("service_appliance_driver"); ok {
			object.SetServiceApplianceDriver(val.(string))
		}
	}
	if d.HasChange("service_appliance_ha_mode") {
		if val, ok := d.GetOk("service_appliance_ha_mode"); ok {
			object.SetServiceApplianceHaMode(val.(string))
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

func UpdateServiceApplianceSetRefsFromResource(object *ServiceApplianceSet, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceServiceApplianceSetCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceSetCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceApplianceSet)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceServiceApplianceSetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "ServiceApplianceSet", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceServiceApplianceSetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceApplianceSet", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceApplianceSetFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetCreate] Creation of resource ServiceApplianceSet on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceApplianceSetRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceSetRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsCreate] Missing 'uuid' field for resource ServiceApplianceSet")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-appliance-set", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsCreate] Retrieving ServiceApplianceSet with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceApplianceSet := obj.(*ServiceApplianceSet) // Fully set by Contrail backend
	if err := SetRefsServiceApplianceSetFromResource(objServiceApplianceSet, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsCreate] Set refs on object ServiceApplianceSet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceApplianceSet.GetHref())
	if err := client.Update(objServiceApplianceSet); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsCreate] Update refs for resource ServiceApplianceSet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceApplianceSet.GetUuid())
	return nil
}

func ResourceServiceApplianceSetRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceSetRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-appliance-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRead] Read resource service-appliance-set on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceApplianceSet)
	WriteServiceApplianceSetToResource(*object, d, m)
	return nil
}

func ResourceServiceApplianceSetRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceSetRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-appliance-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsRead] Read resource service-appliance-set on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceApplianceSet)
	WriteServiceApplianceSetRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceApplianceSetUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceSetUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-appliance-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetUpdate] Retrieving ServiceApplianceSet with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceApplianceSet)
	UpdateServiceApplianceSetFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetUpdate] Update of resource ServiceApplianceSet on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceApplianceSetRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceSetRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-appliance-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsUpdate] Retrieving ServiceApplianceSet with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceApplianceSet)
	UpdateServiceApplianceSetRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsUpdate] Update of resource ServiceApplianceSet on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceApplianceSetDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceSetDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-appliance-set", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetDelete] Deletion of resource service-appliance-set on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceApplianceSetRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceSetRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsDelete] Missing 'uuid' field for resource ServiceApplianceSet")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-appliance-set", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsDelete] Retrieving ServiceApplianceSet with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceApplianceSet := obj.(*ServiceApplianceSet) // Fully set by Contrail backend
	if err := DeleteRefsServiceApplianceSetFromResource(objServiceApplianceSet, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsDelete] Set refs on object ServiceApplianceSet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceApplianceSet.GetHref())
	if err := client.Update(objServiceApplianceSet); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetRefsDelete] Delete refs for resource ServiceApplianceSet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceApplianceSet.GetUuid())
	return nil
}

func ResourceServiceApplianceSetSchema() map[string]*schema.Schema {
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
		"service_appliance_set_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceKeyValuePairs(),
		},
		"service_appliance_driver": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"service_appliance_ha_mode": &schema.Schema{
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

func ResourceServiceApplianceSetRefsSchema() map[string]*schema.Schema {
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

func ResourceServiceApplianceSet() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceApplianceSetCreate,
		Read:   ResourceServiceApplianceSetRead,
		Update: ResourceServiceApplianceSetUpdate,
		Delete: ResourceServiceApplianceSetDelete,
		Schema: ResourceServiceApplianceSetSchema(),
	}
}

func ResourceServiceApplianceSetRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceApplianceSetRefsCreate,
		Read:   ResourceServiceApplianceSetRefsRead,
		Update: ResourceServiceApplianceSetRefsUpdate,
		Delete: ResourceServiceApplianceSetRefsDelete,
		Schema: ResourceServiceApplianceSetRefsSchema(),
	}
}

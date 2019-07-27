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

func ResourceServiceApplianceSetCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceSetCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceApplianceSet)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceApplianceSetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceApplianceSet", client.GetServer(), err)
		}
		object.SetParent(parent)
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
	log.Printf("ResourceServiceApplianceSetREAD")
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
	log.Printf("ResourceServiceApplianceSetRefsREAD")
	return nil
}

func ResourceServiceApplianceSetUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceSetUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-appliance-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceSetResourceUpdate] Retrieving ServiceApplianceSet with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
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
	log.Printf("ResourceServiceApplianceSetRefsDelete: %v", d.Id())
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

func ResourceServiceApplianceSet() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceApplianceSetCreate,
		Read:   ResourceServiceApplianceSetRead,
		Update: ResourceServiceApplianceSetUpdate,
		Delete: ResourceServiceApplianceSetDelete,
		Schema: ResourceServiceApplianceSetSchema(),
	}
}

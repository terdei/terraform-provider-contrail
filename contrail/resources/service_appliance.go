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

func SetServiceApplianceFromResource(object *ServiceAppliance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceApplianceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_appliance_user_credentials"); ok {
		member := new(UserCredentials)
		SetUserCredentialsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceApplianceUserCredentials(member)
	}
	if val, ok := d.GetOk("service_appliance_ip_address"); ok {
		object.SetServiceApplianceIpAddress(val.(string))
	}
	if val, ok := d.GetOk("service_appliance_properties"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceApplianceProperties(member)
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

func SetRefsServiceApplianceFromResource(object *ServiceAppliance, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceApplianceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("physical_interface_refs"); ok {
		log.Printf("Got ref physical_interface_refs -- will call: object.AddPhysicalInterface(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("physical-interface", refId.(string))
			dataObj := new(ServiceApplianceInterfaceType)
			SetServiceApplianceInterfaceTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving physical-interface by Uuid = %v as ref for PhysicalInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPhysicalInterface(refObj.(*PhysicalInterface), *dataObj)
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

func DeleteRefsServiceApplianceFromResource(object *ServiceAppliance, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceApplianceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("physical_interface_refs"); ok {
		log.Printf("Got ref physical_interface_refs -- will call: object.DeletePhysicalInterface(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePhysicalInterface(refId.(string))
		}
	}
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

func WriteServiceApplianceToResource(object ServiceAppliance, d *schema.ResourceData, m interface{}) {

	service_appliance_user_credentialsObj := object.GetServiceApplianceUserCredentials()
	d.Set("service_appliance_user_credentials", TakeUserCredentialsAsMap(&service_appliance_user_credentialsObj))
	d.Set("service_appliance_ip_address", object.GetServiceApplianceIpAddress())
	service_appliance_propertiesObj := object.GetServiceApplianceProperties()
	d.Set("service_appliance_properties", TakeKeyValuePairsAsMap(&service_appliance_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceApplianceRefsToResource(object ServiceAppliance, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetPhysicalInterfaceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("physical_interface_refs", refList)
	}
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

func TakeServiceApplianceAsMap(object *ServiceAppliance) map[string]interface{} {
	omap := make(map[string]interface{})

	service_appliance_user_credentialsObj := object.GetServiceApplianceUserCredentials()
	omap["service_appliance_user_credentials"] = TakeUserCredentialsAsMap(&service_appliance_user_credentialsObj)
	omap["service_appliance_ip_address"] = object.GetServiceApplianceIpAddress()
	service_appliance_propertiesObj := object.GetServiceApplianceProperties()
	omap["service_appliance_properties"] = TakeKeyValuePairsAsMap(&service_appliance_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceApplianceFromResource(object *ServiceAppliance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("service_appliance_user_credentials") {
		if val, ok := d.GetOk("service_appliance_user_credentials"); ok {
			member := new(UserCredentials)
			SetUserCredentialsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceApplianceUserCredentials(member)
		}
	}
	if d.HasChange("service_appliance_ip_address") {
		if val, ok := d.GetOk("service_appliance_ip_address"); ok {
			object.SetServiceApplianceIpAddress(val.(string))
		}
	}
	if d.HasChange("service_appliance_properties") {
		if val, ok := d.GetOk("service_appliance_properties"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceApplianceProperties(member)
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

func UpdateServiceApplianceRefsFromResource(object *ServiceAppliance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("physical_interface_refs") {
		object.ClearPhysicalInterface()
		if val, ok := d.GetOk("physical_interface_refs"); ok {
			log.Printf("Got ref physical_interface_refs -- will call: object.AddPhysicalInterface(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("physical-interface", refId.(string))
				dataObj := new(ServiceApplianceInterfaceType)
				SetServiceApplianceInterfaceTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddPhysicalInterface(refObj.(*PhysicalInterface), *dataObj)
			}
		}
	}
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

func ResourceServiceApplianceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceAppliance)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceApplianceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceAppliance", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceApplianceFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceCreate] Creation of resource ServiceAppliance on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceApplianceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceApplianceRefsCreate] Missing 'uuid' field for resource ServiceAppliance")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-appliance", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsCreate] Retrieving ServiceAppliance with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceAppliance := obj.(*ServiceAppliance) // Fully set by Contrail backend
	if err := SetRefsServiceApplianceFromResource(objServiceAppliance, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsCreate] Set refs on object ServiceAppliance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceAppliance.GetHref())
	if err := client.Update(objServiceAppliance); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsCreate] Update refs for resource ServiceAppliance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceAppliance.GetUuid())
	return nil
}

func ResourceServiceApplianceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-appliance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRead] Read resource service-appliance on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceAppliance)
	WriteServiceApplianceToResource(*object, d, m)
	return nil
}

func ResourceServiceApplianceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-appliance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsRead] Read resource service-appliance on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceAppliance)
	WriteServiceApplianceRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceApplianceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-appliance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceUpdate] Retrieving ServiceAppliance with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceAppliance)
	UpdateServiceApplianceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceUpdate] Update of resource ServiceAppliance on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceApplianceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-appliance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsUpdate] Retrieving ServiceAppliance with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceAppliance)
	UpdateServiceApplianceRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsUpdate] Update of resource ServiceAppliance on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceApplianceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceApplianceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-appliance", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceDelete] Deletion of resource service-appliance on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceApplianceRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceApplianceRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceApplianceRefsDelete] Missing 'uuid' field for resource ServiceAppliance")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-appliance", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsDelete] Retrieving ServiceAppliance with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceAppliance := obj.(*ServiceAppliance) // Fully set by Contrail backend
	if err := DeleteRefsServiceApplianceFromResource(objServiceAppliance, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsDelete] Set refs on object ServiceAppliance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceAppliance.GetHref())
	if err := client.Update(objServiceAppliance); err != nil {
		return fmt.Errorf("[ResourceServiceApplianceRefsDelete] Delete refs for resource ServiceAppliance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceAppliance.GetUuid())
	return nil
}

func ResourceServiceApplianceSchema() map[string]*schema.Schema {
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
		"service_appliance_user_credentials": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceUserCredentials(),
		},
		"service_appliance_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"service_appliance_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceKeyValuePairs(),
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

func ResourceServiceApplianceRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"physical_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"to": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"attr": &schema.Schema{
						Type:     schema.TypeList,
						Elem:     ResourceServiceApplianceInterfaceType(),
						Required: true,
					},
				},
			},
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

func ResourceServiceAppliance() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceApplianceCreate,
		Read:   ResourceServiceApplianceRead,
		Update: ResourceServiceApplianceUpdate,
		Delete: ResourceServiceApplianceDelete,
		Schema: ResourceServiceApplianceSchema(),
	}
}

func ResourceServiceApplianceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceApplianceRefsCreate,
		Read:   ResourceServiceApplianceRefsRead,
		Update: ResourceServiceApplianceRefsUpdate,
		Delete: ResourceServiceApplianceRefsDelete,
		Schema: ResourceServiceApplianceRefsSchema(),
	}
}

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

func SetServiceInstanceFromResource(object *ServiceInstance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceInstanceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_properties"); ok {
		member := new(ServiceInstanceType)
		SetServiceInstanceTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceInstanceProperties(member)
	}
	if val, ok := d.GetOk("service_instance_bindings"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceInstanceBindings(member)
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

func SetRefsServiceInstanceFromResource(object *ServiceInstance, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceInstanceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_template_refs"); ok {
		log.Printf("Got ref service_template_refs -- will call: object.AddServiceTemplate(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-template", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-template by Uuid = %v as ref for ServiceTemplate on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceTemplate(refObj.(*ServiceTemplate))
		}
	}
	if val, ok := d.GetOk("instance_ip_refs"); ok {
		log.Printf("Got ref instance_ip_refs -- will call: object.AddInstanceIp(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("instance-ip", refId.(string))
			dataObj := new(ServiceInterfaceTag)
			SetServiceInterfaceTagFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving instance-ip by Uuid = %v as ref for InstanceIp on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddInstanceIp(refObj.(*InstanceIp), *dataObj)
		}
	}

	return nil
}

func WriteServiceInstanceToResource(object ServiceInstance, d *schema.ResourceData, m interface{}) {

	service_instance_propertiesObj := object.GetServiceInstanceProperties()
	d.Set("service_instance_properties", TakeServiceInstanceTypeAsMap(&service_instance_propertiesObj))
	service_instance_bindingsObj := object.GetServiceInstanceBindings()
	d.Set("service_instance_bindings", TakeKeyValuePairsAsMap(&service_instance_bindingsObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeServiceInstanceAsMap(object *ServiceInstance) map[string]interface{} {
	omap := make(map[string]interface{})

	service_instance_propertiesObj := object.GetServiceInstanceProperties()
	omap["service_instance_properties"] = TakeServiceInstanceTypeAsMap(&service_instance_propertiesObj)
	service_instance_bindingsObj := object.GetServiceInstanceBindings()
	omap["service_instance_bindings"] = TakeKeyValuePairsAsMap(&service_instance_bindingsObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceInstanceFromResource(object *ServiceInstance, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("service_instance_properties") {
		if val, ok := d.GetOk("service_instance_properties"); ok {
			member := new(ServiceInstanceType)
			SetServiceInstanceTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceInstanceProperties(member)
		}
	}
	if d.HasChange("service_instance_bindings") {
		if val, ok := d.GetOk("service_instance_bindings"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceInstanceBindings(member)
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

func ResourceServiceInstanceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceInstanceCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceInstance)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceInstanceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceInstance", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceInstanceFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceInstanceCreate] Creation of resource ServiceInstance on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceInstanceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceInstanceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceInstanceRefsCreate] Missing 'uuid' field for resource ServiceInstance")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-instance", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceInstanceRefsCreate] Retrieving ServiceInstance with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceInstance := obj.(*ServiceInstance) // Fully set by Contrail backend
	if err := SetRefsServiceInstanceFromResource(objServiceInstance, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceInstanceRefsCreate] Set refs on object ServiceInstance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceInstance.GetHref())
	if err := client.Update(objServiceInstance); err != nil {
		return fmt.Errorf("[ResourceServiceInstanceRefsCreate] Update refs for resource ServiceInstance (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceInstance.GetUuid())
	return nil
}

func ResourceServiceInstanceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceInstanceREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-instance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceInstanceRead] Read resource service-instance on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceInstance)
	WriteServiceInstanceToResource(*object, d, m)
	return nil
}

func ResourceServiceInstanceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceInstanceRefsREAD")
	return nil
}

func ResourceServiceInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceInstanceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-instance", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceInstanceResourceUpdate] Retrieving ServiceInstance with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceInstance)
	UpdateServiceInstanceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceInstanceUpdate] Update of resource ServiceInstance on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceInstanceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceInstanceRefsUpdate")
	return nil
}

func ResourceServiceInstanceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceInstanceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-instance", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceInstanceDelete] Deletion of resource service-instance on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceInstanceRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceInstanceRefsDelete: %v", d.Id())
	return nil
}

func ResourceServiceInstanceSchema() map[string]*schema.Schema {
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
		"service_instance_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceInstanceType(),
		},
		"service_instance_bindings": &schema.Schema{
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

func ResourceServiceInstanceRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_template_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceTemplate(),
		},
		"instance_ip_refs": &schema.Schema{
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
						Elem:     ResourceServiceInterfaceTag(),
						Required: true,
					},
				},
			},
		},
	}
}

func ResourceServiceInstance() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceInstanceCreate,
		Read:   ResourceServiceInstanceRead,
		Update: ResourceServiceInstanceUpdate,
		Delete: ResourceServiceInstanceDelete,
		Schema: ResourceServiceInstanceSchema(),
	}
}

func ResourceServiceInstanceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceInstanceRefsCreate,
		Read:   ResourceServiceInstanceRefsRead,
		Update: ResourceServiceInstanceRefsUpdate,
		Delete: ResourceServiceInstanceRefsDelete,
		Schema: ResourceServiceInstanceRefsSchema(),
	}
}

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

func SetServiceHealthCheckFromResource(object *ServiceHealthCheck, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceHealthCheckFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_health_check_properties"); ok {
		member := new(ServiceHealthCheckType)
		SetServiceHealthCheckTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceHealthCheckProperties(member)
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

func SetRefsServiceHealthCheckFromResource(object *ServiceHealthCheck, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceHealthCheckFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-instance", refId.(string))
			dataObj := new(ServiceInterfaceTag)
			SetServiceInterfaceTagFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-instance by Uuid = %v as ref for ServiceInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceInstance(refObj.(*ServiceInstance), *dataObj)
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

func DeleteRefsServiceHealthCheckFromResource(object *ServiceHealthCheck, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsServiceHealthCheckFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.DeleteServiceInstance(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceInstance(refId.(string))
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

func WriteServiceHealthCheckToResource(object ServiceHealthCheck, d *schema.ResourceData, m interface{}) {

	service_health_check_propertiesObj := object.GetServiceHealthCheckProperties()
	d.Set("service_health_check_properties", TakeServiceHealthCheckTypeAsMap(&service_health_check_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteServiceHealthCheckRefsToResource(object ServiceHealthCheck, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetServiceInstanceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_instance_refs", refList)
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

func TakeServiceHealthCheckAsMap(object *ServiceHealthCheck) map[string]interface{} {
	omap := make(map[string]interface{})

	service_health_check_propertiesObj := object.GetServiceHealthCheckProperties()
	omap["service_health_check_properties"] = TakeServiceHealthCheckTypeAsMap(&service_health_check_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceHealthCheckFromResource(object *ServiceHealthCheck, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("service_health_check_properties") {
		if val, ok := d.GetOk("service_health_check_properties"); ok {
			member := new(ServiceHealthCheckType)
			SetServiceHealthCheckTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceHealthCheckProperties(member)
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

func UpdateServiceHealthCheckRefsFromResource(object *ServiceHealthCheck, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("service_instance_refs") {
		object.ClearServiceInstance()
		if val, ok := d.GetOk("service_instance_refs"); ok {
			log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-instance", refId.(string))
				dataObj := new(ServiceInterfaceTag)
				SetServiceInterfaceTagFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceInstance(refObj.(*ServiceInstance), *dataObj)
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

func ResourceServiceHealthCheckCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceHealthCheckCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceHealthCheck)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceHealthCheckCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceHealthCheck", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceHealthCheckFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckCreate] Creation of resource ServiceHealthCheck on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceHealthCheckRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceHealthCheckRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsCreate] Missing 'uuid' field for resource ServiceHealthCheck")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-health-check", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsCreate] Retrieving ServiceHealthCheck with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceHealthCheck := obj.(*ServiceHealthCheck) // Fully set by Contrail backend
	if err := SetRefsServiceHealthCheckFromResource(objServiceHealthCheck, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsCreate] Set refs on object ServiceHealthCheck (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceHealthCheck.GetHref())
	if err := client.Update(objServiceHealthCheck); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsCreate] Update refs for resource ServiceHealthCheck (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceHealthCheck.GetUuid())
	return nil
}

func ResourceServiceHealthCheckRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceHealthCheckRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-health-check", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRead] Read resource service-health-check on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceHealthCheck)
	WriteServiceHealthCheckToResource(*object, d, m)
	return nil
}

func ResourceServiceHealthCheckRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceHealthCheckRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-health-check", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsRead] Read resource service-health-check on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceHealthCheck)
	WriteServiceHealthCheckRefsToResource(*object, d, m)
	return nil
}

func ResourceServiceHealthCheckUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceHealthCheckUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-health-check", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckUpdate] Retrieving ServiceHealthCheck with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceHealthCheck)
	UpdateServiceHealthCheckFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckUpdate] Update of resource ServiceHealthCheck on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceHealthCheckRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceHealthCheckRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-health-check", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsUpdate] Retrieving ServiceHealthCheck with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceHealthCheck)
	UpdateServiceHealthCheckRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsUpdate] Update of resource ServiceHealthCheck on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceHealthCheckDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceHealthCheckDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-health-check", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckDelete] Deletion of resource service-health-check on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceHealthCheckRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceHealthCheckRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsDelete] Missing 'uuid' field for resource ServiceHealthCheck")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-health-check", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsDelete] Retrieving ServiceHealthCheck with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceHealthCheck := obj.(*ServiceHealthCheck) // Fully set by Contrail backend
	if err := DeleteRefsServiceHealthCheckFromResource(objServiceHealthCheck, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsDelete] Set refs on object ServiceHealthCheck (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceHealthCheck.GetHref())
	if err := client.Update(objServiceHealthCheck); err != nil {
		return fmt.Errorf("[ResourceServiceHealthCheckRefsDelete] Delete refs for resource ServiceHealthCheck (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceHealthCheck.GetUuid())
	return nil
}

func ResourceServiceHealthCheckSchema() map[string]*schema.Schema {
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
		"service_health_check_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceHealthCheckType(),
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

func ResourceServiceHealthCheckRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_instance_refs": &schema.Schema{
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

func ResourceServiceHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceHealthCheckCreate,
		Read:   ResourceServiceHealthCheckRead,
		Update: ResourceServiceHealthCheckUpdate,
		Delete: ResourceServiceHealthCheckDelete,
		Schema: ResourceServiceHealthCheckSchema(),
	}
}

func ResourceServiceHealthCheckRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceHealthCheckRefsCreate,
		Read:   ResourceServiceHealthCheckRefsRead,
		Update: ResourceServiceHealthCheckRefsUpdate,
		Delete: ResourceServiceHealthCheckRefsDelete,
		Schema: ResourceServiceHealthCheckRefsSchema(),
	}
}

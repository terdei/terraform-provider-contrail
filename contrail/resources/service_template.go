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

func SetServiceTemplateFromResource(object *ServiceTemplate, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetServiceTemplateFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_template_properties"); ok {
		member := new(ServiceTemplateType)
		SetServiceTemplateTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetServiceTemplateProperties(member)
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

func SetRefsServiceTemplateFromResource(object *ServiceTemplate, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsServiceTemplateFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_appliance_set_refs"); ok {
		log.Printf("Got ref service_appliance_set_refs -- will call: object.AddServiceApplianceSet(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-appliance-set", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-appliance-set by Uuid = %v as ref for ServiceApplianceSet on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceApplianceSet(refObj.(*ServiceApplianceSet))
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

func WriteServiceTemplateToResource(object ServiceTemplate, d *schema.ResourceData, m interface{}) {

	service_template_propertiesObj := object.GetServiceTemplateProperties()
	d.Set("service_template_properties", TakeServiceTemplateTypeAsMap(&service_template_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeServiceTemplateAsMap(object *ServiceTemplate) map[string]interface{} {
	omap := make(map[string]interface{})

	service_template_propertiesObj := object.GetServiceTemplateProperties()
	omap["service_template_properties"] = TakeServiceTemplateTypeAsMap(&service_template_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateServiceTemplateFromResource(object *ServiceTemplate, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("service_template_properties") {
		if val, ok := d.GetOk("service_template_properties"); ok {
			member := new(ServiceTemplateType)
			SetServiceTemplateTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetServiceTemplateProperties(member)
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

func ResourceServiceTemplateCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceTemplateCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ServiceTemplate)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceServiceTemplateCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ServiceTemplate", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetServiceTemplateFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceServiceTemplateCreate] Creation of resource ServiceTemplate on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceServiceTemplateRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceServiceTemplateRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceServiceTemplateRefsCreate] Missing 'uuid' field for resource ServiceTemplate")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("service-template", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceServiceTemplateRefsCreate] Retrieving ServiceTemplate with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objServiceTemplate := obj.(*ServiceTemplate) // Fully set by Contrail backend
	if err := SetRefsServiceTemplateFromResource(objServiceTemplate, d, m); err != nil {
		return fmt.Errorf("[ResourceServiceTemplateRefsCreate] Set refs on object ServiceTemplate (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objServiceTemplate.GetHref())
	if err := client.Update(objServiceTemplate); err != nil {
		return fmt.Errorf("[ResourceServiceTemplateRefsCreate] Update refs for resource ServiceTemplate (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objServiceTemplate.GetUuid())
	return nil
}

func ResourceServiceTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceTemplateREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("service-template", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceTemplateRead] Read resource service-template on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ServiceTemplate)
	WriteServiceTemplateToResource(*object, d, m)
	return nil
}

func ResourceServiceTemplateRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceTemplateRefsREAD")
	return nil
}

func ResourceServiceTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceTemplateUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("service-template", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceServiceTemplateResourceUpdate] Retrieving ServiceTemplate with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ServiceTemplate)
	UpdateServiceTemplateFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceServiceTemplateUpdate] Update of resource ServiceTemplate on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceServiceTemplateRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceTemplateRefsUpdate")
	return nil
}

func ResourceServiceTemplateDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceTemplateDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("service-template", d.Id()); err != nil {
		return fmt.Errorf("[ResourceServiceTemplateDelete] Deletion of resource service-template on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceServiceTemplateRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceServiceTemplateRefsDelete: %v", d.Id())
	return nil
}

func ResourceServiceTemplateSchema() map[string]*schema.Schema {
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
		"service_template_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceTemplateType(),
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

func ResourceServiceTemplateRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_appliance_set_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceApplianceSet(),
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourceServiceTemplate() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceTemplateCreate,
		Read:   ResourceServiceTemplateRead,
		Update: ResourceServiceTemplateUpdate,
		Delete: ResourceServiceTemplateDelete,
		Schema: ResourceServiceTemplateSchema(),
	}
}

func ResourceServiceTemplateRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceServiceTemplateRefsCreate,
		Read:   ResourceServiceTemplateRefsRead,
		Update: ResourceServiceTemplateRefsUpdate,
		Delete: ResourceServiceTemplateRefsDelete,
		Schema: ResourceServiceTemplateRefsSchema(),
	}
}

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

func SetLoadbalancerFromResource(object *Loadbalancer, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLoadbalancerFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_properties"); ok {
		member := new(LoadbalancerType)
		SetLoadbalancerTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLoadbalancerProperties(member)
	}
	if val, ok := d.GetOk("loadbalancer_provider"); ok {
		object.SetLoadbalancerProvider(val.(string))
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

func SetRefsLoadbalancerFromResource(object *Loadbalancer, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLoadbalancerFromResource] key = %v, prefix = %v", key, prefix)
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
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-instance", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-instance by Uuid = %v as ref for ServiceInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceInstance(refObj.(*ServiceInstance))
		}
	}
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine-interface by Uuid = %v as ref for VirtualMachineInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
		}
	}

	return nil
}

func WriteLoadbalancerToResource(object Loadbalancer, d *schema.ResourceData, m interface{}) {

	loadbalancer_propertiesObj := object.GetLoadbalancerProperties()
	d.Set("loadbalancer_properties", TakeLoadbalancerTypeAsMap(&loadbalancer_propertiesObj))
	d.Set("loadbalancer_provider", object.GetLoadbalancerProvider())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeLoadbalancerAsMap(object *Loadbalancer) map[string]interface{} {
	omap := make(map[string]interface{})

	loadbalancer_propertiesObj := object.GetLoadbalancerProperties()
	omap["loadbalancer_properties"] = TakeLoadbalancerTypeAsMap(&loadbalancer_propertiesObj)
	omap["loadbalancer_provider"] = object.GetLoadbalancerProvider()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLoadbalancerFromResource(object *Loadbalancer, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("loadbalancer_properties") {
		if val, ok := d.GetOk("loadbalancer_properties"); ok {
			member := new(LoadbalancerType)
			SetLoadbalancerTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLoadbalancerProperties(member)
		}
	}
	if d.HasChange("loadbalancer_provider") {
		if val, ok := d.GetOk("loadbalancer_provider"); ok {
			object.SetLoadbalancerProvider(val.(string))
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

func ResourceLoadbalancerCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Loadbalancer)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceLoadbalancerCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Loadbalancer", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLoadbalancerFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerCreate] Creation of resource Loadbalancer on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLoadbalancerRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerRefsCreate] Missing 'uuid' field for resource Loadbalancer")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerRefsCreate] Retrieving Loadbalancer with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancer := obj.(*Loadbalancer) // Fully set by Contrail backend
	if err := SetRefsLoadbalancerFromResource(objLoadbalancer, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerRefsCreate] Set refs on object Loadbalancer (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancer.GetHref())
	if err := client.Update(objLoadbalancer); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerRefsCreate] Update refs for resource Loadbalancer (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancer.GetUuid())
	return nil
}

func ResourceLoadbalancerRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerRead] Read resource loadbalancer on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Loadbalancer)
	WriteLoadbalancerToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerRefsREAD")
	return nil
}

func ResourceLoadbalancerUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerResourceUpdate] Retrieving Loadbalancer with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Loadbalancer)
	UpdateLoadbalancerFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerUpdate] Update of resource Loadbalancer on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerRefsUpdate")
	return nil
}

func ResourceLoadbalancerDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("loadbalancer", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerDelete] Deletion of resource loadbalancer on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLoadbalancerRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerRefsDelete: %v", d.Id())
	return nil
}

func ResourceLoadbalancerSchema() map[string]*schema.Schema {
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
		"loadbalancer_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLoadbalancerType(),
		},
		"loadbalancer_provider": &schema.Schema{
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

func ResourceLoadbalancerRefsSchema() map[string]*schema.Schema {
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
		"service_instance_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceInstance(),
		},
		"virtual_machine_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachineInterface(),
		},
	}
}

func ResourceLoadbalancer() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerCreate,
		Read:   ResourceLoadbalancerRead,
		Update: ResourceLoadbalancerUpdate,
		Delete: ResourceLoadbalancerDelete,
		Schema: ResourceLoadbalancerSchema(),
	}
}

func ResourceLoadbalancerRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerRefsCreate,
		Read:   ResourceLoadbalancerRefsRead,
		Update: ResourceLoadbalancerRefsUpdate,
		Delete: ResourceLoadbalancerRefsDelete,
		Schema: ResourceLoadbalancerRefsSchema(),
	}
}

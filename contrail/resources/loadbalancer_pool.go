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

func SetLoadbalancerPoolFromResource(object *LoadbalancerPool, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLoadbalancerPoolFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_pool_properties"); ok {
		member := new(LoadbalancerPoolType)
		SetLoadbalancerPoolTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLoadbalancerPoolProperties(member)
	}
	if val, ok := d.GetOk("loadbalancer_pool_provider"); ok {
		object.SetLoadbalancerPoolProvider(val.(string))
	}
	if val, ok := d.GetOk("loadbalancer_pool_custom_attributes"); ok {
		member := new(KeyValuePairs)
		SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetLoadbalancerPoolCustomAttributes(member)
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

func SetRefsLoadbalancerPoolFromResource(object *LoadbalancerPool, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLoadbalancerPoolFromResource] key = %v, prefix = %v", key, prefix)
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
	if val, ok := d.GetOk("loadbalancer_listener_refs"); ok {
		log.Printf("Got ref loadbalancer_listener_refs -- will call: object.AddLoadbalancerListener(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("loadbalancer-listener", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving loadbalancer-listener by Uuid = %v as ref for LoadbalancerListener on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddLoadbalancerListener(refObj.(*LoadbalancerListener))
		}
	}
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
	if val, ok := d.GetOk("loadbalancer_healthmonitor_refs"); ok {
		log.Printf("Got ref loadbalancer_healthmonitor_refs -- will call: object.AddLoadbalancerHealthmonitor(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("loadbalancer-healthmonitor", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving loadbalancer-healthmonitor by Uuid = %v as ref for LoadbalancerHealthmonitor on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddLoadbalancerHealthmonitor(refObj.(*LoadbalancerHealthmonitor))
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

func DeleteRefsLoadbalancerPoolFromResource(object *LoadbalancerPool, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsLoadbalancerPoolFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.DeleteServiceInstance(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceInstance(refId.(string))
		}
	}
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.DeleteVirtualMachineInterface(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualMachineInterface(refId.(string))
		}
	}
	if val, ok := d.GetOk("loadbalancer_listener_refs"); ok {
		log.Printf("Got ref loadbalancer_listener_refs -- will call: object.DeleteLoadbalancerListener(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteLoadbalancerListener(refId.(string))
		}
	}
	if val, ok := d.GetOk("service_appliance_set_refs"); ok {
		log.Printf("Got ref service_appliance_set_refs -- will call: object.DeleteServiceApplianceSet(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceApplianceSet(refId.(string))
		}
	}
	if val, ok := d.GetOk("loadbalancer_healthmonitor_refs"); ok {
		log.Printf("Got ref loadbalancer_healthmonitor_refs -- will call: object.DeleteLoadbalancerHealthmonitor(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteLoadbalancerHealthmonitor(refId.(string))
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

func WriteLoadbalancerPoolToResource(object LoadbalancerPool, d *schema.ResourceData, m interface{}) {

	loadbalancer_pool_propertiesObj := object.GetLoadbalancerPoolProperties()
	d.Set("loadbalancer_pool_properties", TakeLoadbalancerPoolTypeAsMap(&loadbalancer_pool_propertiesObj))
	d.Set("loadbalancer_pool_provider", object.GetLoadbalancerPoolProvider())
	loadbalancer_pool_custom_attributesObj := object.GetLoadbalancerPoolCustomAttributes()
	d.Set("loadbalancer_pool_custom_attributes", TakeKeyValuePairsAsMap(&loadbalancer_pool_custom_attributesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteLoadbalancerPoolRefsToResource(object LoadbalancerPool, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetServiceInstanceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_instance_refs", refList)
	}
	if ref, err := object.GetVirtualMachineInterfaceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_machine_interface_refs", refList)
	}
	if ref, err := object.GetLoadbalancerListenerRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("loadbalancer_listener_refs", refList)
	}
	if ref, err := object.GetServiceApplianceSetRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_appliance_set_refs", refList)
	}
	if ref, err := object.GetLoadbalancerHealthmonitorRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("loadbalancer_healthmonitor_refs", refList)
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

func TakeLoadbalancerPoolAsMap(object *LoadbalancerPool) map[string]interface{} {
	omap := make(map[string]interface{})

	loadbalancer_pool_propertiesObj := object.GetLoadbalancerPoolProperties()
	omap["loadbalancer_pool_properties"] = TakeLoadbalancerPoolTypeAsMap(&loadbalancer_pool_propertiesObj)
	omap["loadbalancer_pool_provider"] = object.GetLoadbalancerPoolProvider()
	loadbalancer_pool_custom_attributesObj := object.GetLoadbalancerPoolCustomAttributes()
	omap["loadbalancer_pool_custom_attributes"] = TakeKeyValuePairsAsMap(&loadbalancer_pool_custom_attributesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLoadbalancerPoolFromResource(object *LoadbalancerPool, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("loadbalancer_pool_properties") {
		if val, ok := d.GetOk("loadbalancer_pool_properties"); ok {
			member := new(LoadbalancerPoolType)
			SetLoadbalancerPoolTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLoadbalancerPoolProperties(member)
		}
	}
	if d.HasChange("loadbalancer_pool_provider") {
		if val, ok := d.GetOk("loadbalancer_pool_provider"); ok {
			object.SetLoadbalancerPoolProvider(val.(string))
		}
	}
	if d.HasChange("loadbalancer_pool_custom_attributes") {
		if val, ok := d.GetOk("loadbalancer_pool_custom_attributes"); ok {
			member := new(KeyValuePairs)
			SetKeyValuePairsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetLoadbalancerPoolCustomAttributes(member)
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

func UpdateLoadbalancerPoolRefsFromResource(object *LoadbalancerPool, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("service_instance_refs") {
		object.ClearServiceInstance()
		if val, ok := d.GetOk("service_instance_refs"); ok {
			log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-instance", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceInstance(refObj.(*ServiceInstance))
			}
		}
	}
	if d.HasChange("virtual_machine_interface_refs") {
		object.ClearVirtualMachineInterface()
		if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
			log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("virtual-machine-interface", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
			}
		}
	}
	if d.HasChange("loadbalancer_listener_refs") {
		object.ClearLoadbalancerListener()
		if val, ok := d.GetOk("loadbalancer_listener_refs"); ok {
			log.Printf("Got ref loadbalancer_listener_refs -- will call: object.AddLoadbalancerListener(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("loadbalancer-listener", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddLoadbalancerListener(refObj.(*LoadbalancerListener))
			}
		}
	}
	if d.HasChange("service_appliance_set_refs") {
		object.ClearServiceApplianceSet()
		if val, ok := d.GetOk("service_appliance_set_refs"); ok {
			log.Printf("Got ref service_appliance_set_refs -- will call: object.AddServiceApplianceSet(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-appliance-set", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceApplianceSet(refObj.(*ServiceApplianceSet))
			}
		}
	}
	if d.HasChange("loadbalancer_healthmonitor_refs") {
		object.ClearLoadbalancerHealthmonitor()
		if val, ok := d.GetOk("loadbalancer_healthmonitor_refs"); ok {
			log.Printf("Got ref loadbalancer_healthmonitor_refs -- will call: object.AddLoadbalancerHealthmonitor(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("loadbalancer-healthmonitor", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddLoadbalancerHealthmonitor(refObj.(*LoadbalancerHealthmonitor))
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

func ResourceLoadbalancerPoolCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerPoolCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LoadbalancerPool)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceLoadbalancerPoolCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "LoadbalancerPool", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceLoadbalancerPoolCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LoadbalancerPool", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLoadbalancerPoolFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolCreate] Creation of resource LoadbalancerPool on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLoadbalancerPoolRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerPoolRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsCreate] Missing 'uuid' field for resource LoadbalancerPool")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-pool", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsCreate] Retrieving LoadbalancerPool with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerPool := obj.(*LoadbalancerPool) // Fully set by Contrail backend
	if err := SetRefsLoadbalancerPoolFromResource(objLoadbalancerPool, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsCreate] Set refs on object LoadbalancerPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerPool.GetHref())
	if err := client.Update(objLoadbalancerPool); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsCreate] Update refs for resource LoadbalancerPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerPool.GetUuid())
	return nil
}

func ResourceLoadbalancerPoolRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerPoolRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRead] Read resource loadbalancer-pool on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerPool)
	WriteLoadbalancerPoolToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerPoolRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerPoolRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("loadbalancer-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsRead] Read resource loadbalancer-pool on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LoadbalancerPool)
	WriteLoadbalancerPoolRefsToResource(*object, d, m)
	return nil
}

func ResourceLoadbalancerPoolUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerPoolUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolUpdate] Retrieving LoadbalancerPool with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerPool)
	UpdateLoadbalancerPoolFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolUpdate] Update of resource LoadbalancerPool on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerPoolRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerPoolRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("loadbalancer-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsUpdate] Retrieving LoadbalancerPool with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LoadbalancerPool)
	UpdateLoadbalancerPoolRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsUpdate] Update of resource LoadbalancerPool on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLoadbalancerPoolDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLoadbalancerPoolDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("loadbalancer-pool", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolDelete] Deletion of resource loadbalancer-pool on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLoadbalancerPoolRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLoadbalancerPoolRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsDelete] Missing 'uuid' field for resource LoadbalancerPool")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("loadbalancer-pool", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsDelete] Retrieving LoadbalancerPool with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLoadbalancerPool := obj.(*LoadbalancerPool) // Fully set by Contrail backend
	if err := DeleteRefsLoadbalancerPoolFromResource(objLoadbalancerPool, d, m); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsDelete] Set refs on object LoadbalancerPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLoadbalancerPool.GetHref())
	if err := client.Update(objLoadbalancerPool); err != nil {
		return fmt.Errorf("[ResourceLoadbalancerPoolRefsDelete] Delete refs for resource LoadbalancerPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLoadbalancerPool.GetUuid())
	return nil
}

func ResourceLoadbalancerPoolSchema() map[string]*schema.Schema {
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
		"loadbalancer_pool_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLoadbalancerPoolType(),
		},
		"loadbalancer_pool_provider": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"loadbalancer_pool_custom_attributes": &schema.Schema{
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

func ResourceLoadbalancerPoolRefsSchema() map[string]*schema.Schema {
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
				},
			},
		},
		"virtual_machine_interface_refs": &schema.Schema{
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
		"loadbalancer_listener_refs": &schema.Schema{
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
		"service_appliance_set_refs": &schema.Schema{
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
		"loadbalancer_healthmonitor_refs": &schema.Schema{
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

func ResourceLoadbalancerPool() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerPoolCreate,
		Read:   ResourceLoadbalancerPoolRead,
		Update: ResourceLoadbalancerPoolUpdate,
		Delete: ResourceLoadbalancerPoolDelete,
		Schema: ResourceLoadbalancerPoolSchema(),
	}
}

func ResourceLoadbalancerPoolRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLoadbalancerPoolRefsCreate,
		Read:   ResourceLoadbalancerPoolRefsRead,
		Update: ResourceLoadbalancerPoolRefsUpdate,
		Delete: ResourceLoadbalancerPoolRefsDelete,
		Schema: ResourceLoadbalancerPoolRefsSchema(),
	}
}

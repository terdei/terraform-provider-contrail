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

func SetBgpAsAServiceFromResource(object *BgpAsAService, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetBgpAsAServiceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("autonomous_system"); ok {
		object.SetAutonomousSystem(val.(int))
	}
	if val, ok := d.GetOk("bgpaas_shared"); ok {
		object.SetBgpaasShared(val.(bool))
	}
	if val, ok := d.GetOk("bgpaas_ip_address"); ok {
		object.SetBgpaasIpAddress(val.(string))
	}
	if val, ok := d.GetOk("bgpaas_session_attributes"); ok {
		object.SetBgpaasSessionAttributes(val.(string))
	}
	if val, ok := d.GetOk("bgpaas_ipv4_mapped_ipv6_nexthop"); ok {
		object.SetBgpaasIpv4MappedIpv6Nexthop(val.(bool))
	}
	if val, ok := d.GetOk("bgpaas_suppress_route_advertisement"); ok {
		object.SetBgpaasSuppressRouteAdvertisement(val.(bool))
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

func SetRefsBgpAsAServiceFromResource(object *BgpAsAService, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsBgpAsAServiceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_health_check_refs"); ok {
		log.Printf("Got ref service_health_check_refs -- will call: object.AddServiceHealthCheck(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-health-check", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-health-check by Uuid = %v as ref for ServiceHealthCheck on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceHealthCheck(refObj.(*ServiceHealthCheck))
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

func SetReqRefsBgpAsAServiceFromResource(object *BgpAsAService, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsBgpAsAServiceFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsBgpAsAServiceFromResource(object *BgpAsAService, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsBgpAsAServiceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_health_check_refs"); ok {
		log.Printf("Got ref service_health_check_refs -- will call: object.DeleteServiceHealthCheck(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceHealthCheck(refId.(string))
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

func WriteBgpAsAServiceToResource(object BgpAsAService, d *schema.ResourceData, m interface{}) {

	d.Set("autonomous_system", object.GetAutonomousSystem())
	d.Set("bgpaas_shared", object.GetBgpaasShared())
	d.Set("bgpaas_ip_address", object.GetBgpaasIpAddress())
	d.Set("bgpaas_session_attributes", object.GetBgpaasSessionAttributes())
	d.Set("bgpaas_ipv4_mapped_ipv6_nexthop", object.GetBgpaasIpv4MappedIpv6Nexthop())
	d.Set("bgpaas_suppress_route_advertisement", object.GetBgpaasSuppressRouteAdvertisement())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

	if ref, err := object.GetVirtualMachineInterfaceRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_machine_interface_refs", refList)
	}
}

func WriteBgpAsAServiceRefsToResource(object BgpAsAService, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetServiceHealthCheckRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("service_health_check_refs", refList)
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

func TakeBgpAsAServiceAsMap(object *BgpAsAService) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["autonomous_system"] = object.GetAutonomousSystem()
	omap["bgpaas_shared"] = object.GetBgpaasShared()
	omap["bgpaas_ip_address"] = object.GetBgpaasIpAddress()
	omap["bgpaas_session_attributes"] = object.GetBgpaasSessionAttributes()
	omap["bgpaas_ipv4_mapped_ipv6_nexthop"] = object.GetBgpaasIpv4MappedIpv6Nexthop()
	omap["bgpaas_suppress_route_advertisement"] = object.GetBgpaasSuppressRouteAdvertisement()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateBgpAsAServiceFromResource(object *BgpAsAService, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("autonomous_system") {
		if val, ok := d.GetOk("autonomous_system"); ok {
			object.SetAutonomousSystem(val.(int))
		}
	}
	if d.HasChange("bgpaas_shared") {
		if val, ok := d.GetOk("bgpaas_shared"); ok {
			object.SetBgpaasShared(val.(bool))
		}
	}
	if d.HasChange("bgpaas_ip_address") {
		if val, ok := d.GetOk("bgpaas_ip_address"); ok {
			object.SetBgpaasIpAddress(val.(string))
		}
	}
	if d.HasChange("bgpaas_session_attributes") {
		if val, ok := d.GetOk("bgpaas_session_attributes"); ok {
			object.SetBgpaasSessionAttributes(val.(string))
		}
	}
	if d.HasChange("bgpaas_ipv4_mapped_ipv6_nexthop") {
		if val, ok := d.GetOk("bgpaas_ipv4_mapped_ipv6_nexthop"); ok {
			object.SetBgpaasIpv4MappedIpv6Nexthop(val.(bool))
		}
	}
	if d.HasChange("bgpaas_suppress_route_advertisement") {
		if val, ok := d.GetOk("bgpaas_suppress_route_advertisement"); ok {
			object.SetBgpaasSuppressRouteAdvertisement(val.(bool))
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

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
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

}

func UpdateBgpAsAServiceRefsFromResource(object *BgpAsAService, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("service_health_check_refs") {
		object.ClearServiceHealthCheck()
		if val, ok := d.GetOk("service_health_check_refs"); ok {
			log.Printf("Got ref service_health_check_refs -- will call: object.AddServiceHealthCheck(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("service-health-check", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddServiceHealthCheck(refObj.(*ServiceHealthCheck))
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

func ResourceBgpAsAServiceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpAsAServiceCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(BgpAsAService)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceBgpAsAServiceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "BgpAsAService", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceBgpAsAServiceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "BgpAsAService", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetBgpAsAServiceFromResource(object, d, m)

	if err := SetReqRefsBgpAsAServiceFromResource(object, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceReqRefsCreate] Set required refs on object BgpAsAService on %v (%v)", client.GetServer(), err)
	}

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceCreate] Creation of resource BgpAsAService on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceBgpAsAServiceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpAsAServiceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsCreate] Missing 'uuid' field for resource BgpAsAService")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgp-as-a-service", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsCreate] Retrieving BgpAsAService with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpAsAService := obj.(*BgpAsAService) // Fully set by Contrail backend
	if err := SetRefsBgpAsAServiceFromResource(objBgpAsAService, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsCreate] Set refs on object BgpAsAService (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpAsAService.GetHref())
	if err := client.Update(objBgpAsAService); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsCreate] Update refs for resource BgpAsAService (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpAsAService.GetUuid())
	return nil
}

func ResourceBgpAsAServiceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpAsAServiceRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgp-as-a-service", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRead] Read resource bgp-as-a-service on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BgpAsAService)
	WriteBgpAsAServiceToResource(*object, d, m)
	return nil
}

func ResourceBgpAsAServiceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpAsAServiceRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgp-as-a-service", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsRead] Read resource bgp-as-a-service on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BgpAsAService)
	WriteBgpAsAServiceRefsToResource(*object, d, m)
	return nil
}

func ResourceBgpAsAServiceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpAsAServiceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgp-as-a-service", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceUpdate] Retrieving BgpAsAService with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BgpAsAService)
	UpdateBgpAsAServiceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceUpdate] Update of resource BgpAsAService on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpAsAServiceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpAsAServiceRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgp-as-a-service", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsUpdate] Retrieving BgpAsAService with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BgpAsAService)
	UpdateBgpAsAServiceRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsUpdate] Update of resource BgpAsAService on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpAsAServiceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpAsAServiceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("bgp-as-a-service", d.Id()); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceDelete] Deletion of resource bgp-as-a-service on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceBgpAsAServiceRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpAsAServiceRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsDelete] Missing 'uuid' field for resource BgpAsAService")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgp-as-a-service", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsDelete] Retrieving BgpAsAService with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpAsAService := obj.(*BgpAsAService) // Fully set by Contrail backend
	if err := DeleteRefsBgpAsAServiceFromResource(objBgpAsAService, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsDelete] Set refs on object BgpAsAService (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpAsAService.GetHref())
	if err := client.Update(objBgpAsAService); err != nil {
		return fmt.Errorf("[ResourceBgpAsAServiceRefsDelete] Delete refs for resource BgpAsAService (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpAsAService.GetUuid())
	return nil
}

func ResourceBgpAsAServiceSchema() map[string]*schema.Schema {
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
		"autonomous_system": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"bgpaas_shared": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"bgpaas_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"bgpaas_session_attributes": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"bgpaas_ipv4_mapped_ipv6_nexthop": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"bgpaas_suppress_route_advertisement": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
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
	}
}

func ResourceBgpAsAServiceRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_health_check_refs": &schema.Schema{
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

func ResourceBgpAsAService() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpAsAServiceCreate,
		Read:   ResourceBgpAsAServiceRead,
		Update: ResourceBgpAsAServiceUpdate,
		Delete: ResourceBgpAsAServiceDelete,
		Schema: ResourceBgpAsAServiceSchema(),
	}
}

func ResourceBgpAsAServiceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpAsAServiceRefsCreate,
		Read:   ResourceBgpAsAServiceRefsRead,
		Update: ResourceBgpAsAServiceRefsUpdate,
		Delete: ResourceBgpAsAServiceRefsDelete,
		Schema: ResourceBgpAsAServiceRefsSchema(),
	}
}

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

func SetInstanceIpFromResource(object *InstanceIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetInstanceIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("instance_ip_address"); ok {
		object.SetInstanceIpAddress(val.(string))
	}
	if val, ok := d.GetOk("instance_ip_family"); ok {
		object.SetInstanceIpFamily(val.(string))
	}
	if val, ok := d.GetOk("instance_ip_mode"); ok {
		object.SetInstanceIpMode(val.(string))
	}
	if val, ok := d.GetOk("secondary_ip_tracking_ip"); ok {
		member := new(SubnetType)
		SetSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetSecondaryIpTrackingIp(member)
	}
	if val, ok := d.GetOk("subnet_uuid"); ok {
		object.SetSubnetUuid(val.(string))
	}
	if val, ok := d.GetOk("instance_ip_secondary"); ok {
		object.SetInstanceIpSecondary(val.(bool))
	}
	if val, ok := d.GetOk("instance_ip_local_ip"); ok {
		object.SetInstanceIpLocalIp(val.(bool))
	}
	if val, ok := d.GetOk("service_instance_ip"); ok {
		object.SetServiceInstanceIp(val.(bool))
	}
	if val, ok := d.GetOk("service_health_check_ip"); ok {
		object.SetServiceHealthCheckIp(val.(bool))
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

func SetRefsInstanceIpFromResource(object *InstanceIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsInstanceIpFromResource] key = %v, prefix = %v", key, prefix)
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
	if val, ok := d.GetOk("physical_router_refs"); ok {
		log.Printf("Got ref physical_router_refs -- will call: object.AddPhysicalRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("physical-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving physical-router by Uuid = %v as ref for PhysicalRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPhysicalRouter(refObj.(*PhysicalRouter))
		}
	}
	if val, ok := d.GetOk("virtual_router_refs"); ok {
		log.Printf("Got ref virtual_router_refs -- will call: object.AddVirtualRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-router by Uuid = %v as ref for VirtualRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualRouter(refObj.(*VirtualRouter))
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

func SetReqRefsInstanceIpFromResource(object *InstanceIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsInstanceIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_network_refs"); ok {
		log.Printf("Got ref virtual_network_refs -- will call: object.AddVirtualNetwork(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-network", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-network by Uuid = %v as ref for VirtualNetwork on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualNetwork(refObj.(*VirtualNetwork))
		}
	}
	if val, ok := d.GetOk("network_ipam_refs"); ok {
		log.Printf("Got ref network_ipam_refs -- will call: object.AddNetworkIpam(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("network-ipam", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving network-ipam by Uuid = %v as ref for NetworkIpam on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddNetworkIpam(refObj.(*NetworkIpam))
		}
	}

	return nil
}

func DeleteRefsInstanceIpFromResource(object *InstanceIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsInstanceIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.DeleteVirtualMachineInterface(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualMachineInterface(refId.(string))
		}
	}
	if val, ok := d.GetOk("physical_router_refs"); ok {
		log.Printf("Got ref physical_router_refs -- will call: object.DeletePhysicalRouter(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePhysicalRouter(refId.(string))
		}
	}
	if val, ok := d.GetOk("virtual_router_refs"); ok {
		log.Printf("Got ref virtual_router_refs -- will call: object.DeleteVirtualRouter(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualRouter(refId.(string))
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

func WriteInstanceIpToResource(object InstanceIp, d *schema.ResourceData, m interface{}) {

	d.Set("instance_ip_address", object.GetInstanceIpAddress())
	d.Set("instance_ip_family", object.GetInstanceIpFamily())
	d.Set("instance_ip_mode", object.GetInstanceIpMode())
	secondary_ip_tracking_ipObj := object.GetSecondaryIpTrackingIp()
	d.Set("secondary_ip_tracking_ip", TakeSubnetTypeAsMap(&secondary_ip_tracking_ipObj))
	d.Set("subnet_uuid", object.GetSubnetUuid())
	d.Set("instance_ip_secondary", object.GetInstanceIpSecondary())
	d.Set("instance_ip_local_ip", object.GetInstanceIpLocalIp())
	d.Set("service_instance_ip", object.GetServiceInstanceIp())
	d.Set("service_health_check_ip", object.GetServiceHealthCheckIp())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

	if ref, err := object.GetVirtualNetworkRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("virtual_network_refs", refList)
	}
	if ref, err := object.GetNetworkIpamRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("network_ipam_refs", refList)
	}
}

func TakeInstanceIpAsMap(object *InstanceIp) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["instance_ip_address"] = object.GetInstanceIpAddress()
	omap["instance_ip_family"] = object.GetInstanceIpFamily()
	omap["instance_ip_mode"] = object.GetInstanceIpMode()
	secondary_ip_tracking_ipObj := object.GetSecondaryIpTrackingIp()
	omap["secondary_ip_tracking_ip"] = TakeSubnetTypeAsMap(&secondary_ip_tracking_ipObj)
	omap["subnet_uuid"] = object.GetSubnetUuid()
	omap["instance_ip_secondary"] = object.GetInstanceIpSecondary()
	omap["instance_ip_local_ip"] = object.GetInstanceIpLocalIp()
	omap["service_instance_ip"] = object.GetServiceInstanceIp()
	omap["service_health_check_ip"] = object.GetServiceHealthCheckIp()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateInstanceIpFromResource(object *InstanceIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("instance_ip_address") {
		if val, ok := d.GetOk("instance_ip_address"); ok {
			object.SetInstanceIpAddress(val.(string))
		}
	}
	if d.HasChange("instance_ip_family") {
		if val, ok := d.GetOk("instance_ip_family"); ok {
			object.SetInstanceIpFamily(val.(string))
		}
	}
	if d.HasChange("instance_ip_mode") {
		if val, ok := d.GetOk("instance_ip_mode"); ok {
			object.SetInstanceIpMode(val.(string))
		}
	}
	if d.HasChange("secondary_ip_tracking_ip") {
		if val, ok := d.GetOk("secondary_ip_tracking_ip"); ok {
			member := new(SubnetType)
			SetSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetSecondaryIpTrackingIp(member)
		}
	}
	if d.HasChange("subnet_uuid") {
		if val, ok := d.GetOk("subnet_uuid"); ok {
			object.SetSubnetUuid(val.(string))
		}
	}
	if d.HasChange("instance_ip_secondary") {
		if val, ok := d.GetOk("instance_ip_secondary"); ok {
			object.SetInstanceIpSecondary(val.(bool))
		}
	}
	if d.HasChange("instance_ip_local_ip") {
		if val, ok := d.GetOk("instance_ip_local_ip"); ok {
			object.SetInstanceIpLocalIp(val.(bool))
		}
	}
	if d.HasChange("service_instance_ip") {
		if val, ok := d.GetOk("service_instance_ip"); ok {
			object.SetServiceInstanceIp(val.(bool))
		}
	}
	if d.HasChange("service_health_check_ip") {
		if val, ok := d.GetOk("service_health_check_ip"); ok {
			object.SetServiceHealthCheckIp(val.(bool))
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
	if d.HasChange("virtual_network_refs") {
		if val, ok := d.GetOk("virtual_network_refs"); ok {
			log.Printf("Got ref virtual_network_refs -- will call: object.AddVirtualNetwork(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("virtual-network", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddVirtualNetwork(refObj.(*VirtualNetwork))
			}
		}
	}
	if d.HasChange("network_ipam_refs") {
		if val, ok := d.GetOk("network_ipam_refs"); ok {
			log.Printf("Got ref network_ipam_refs -- will call: object.AddNetworkIpam(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("network-ipam", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddNetworkIpam(refObj.(*NetworkIpam))
			}
		}
	}

}

func ResourceInstanceIpCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceInstanceIpCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(InstanceIp)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceInstanceIpCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "InstanceIp", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetInstanceIpFromResource(object, d, m)

	if err := SetReqRefsInstanceIpFromResource(object, d, m); err != nil {
		return fmt.Errorf("[ResourceInstanceIpReqRefsCreate] Set required refs on object InstanceIp on %v (%v)", client.GetServer(), err)
	}

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceInstanceIpCreate] Creation of resource InstanceIp on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceInstanceIpRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceInstanceIpRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceInstanceIpRefsCreate] Missing 'uuid' field for resource InstanceIp")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("instance-ip", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceInstanceIpRefsCreate] Retrieving InstanceIp with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objInstanceIp := obj.(*InstanceIp) // Fully set by Contrail backend
	if err := SetRefsInstanceIpFromResource(objInstanceIp, d, m); err != nil {
		return fmt.Errorf("[ResourceInstanceIpRefsCreate] Set refs on object InstanceIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objInstanceIp.GetHref())
	if err := client.Update(objInstanceIp); err != nil {
		return fmt.Errorf("[ResourceInstanceIpRefsCreate] Update refs for resource InstanceIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objInstanceIp.GetUuid())
	return nil
}

func ResourceInstanceIpRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInstanceIpREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("instance-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceInstanceIpRead] Read resource instance-ip on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*InstanceIp)
	WriteInstanceIpToResource(*object, d, m)
	return nil
}

func ResourceInstanceIpRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInstanceIpRefsREAD")
	return nil
}

func ResourceInstanceIpUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInstanceIpUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("instance-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceInstanceIpResourceUpdate] Retrieving InstanceIp with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*InstanceIp)
	UpdateInstanceIpFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceInstanceIpUpdate] Update of resource InstanceIp on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceInstanceIpRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInstanceIpRefsUpdate")
	return nil
}

func ResourceInstanceIpDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInstanceIpDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("instance-ip", d.Id()); err != nil {
		return fmt.Errorf("[ResourceInstanceIpDelete] Deletion of resource instance-ip on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceInstanceIpRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceInstanceIpRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceInstanceIpRefsDelete] Missing 'uuid' field for resource InstanceIp")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("instance-ip", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceInstanceIpRefsDelete] Retrieving InstanceIp with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objInstanceIp := obj.(*InstanceIp) // Fully set by Contrail backend
	if err := DeleteRefsInstanceIpFromResource(objInstanceIp, d, m); err != nil {
		return fmt.Errorf("[ResourceInstanceIpRefsDelete] Set refs on object InstanceIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objInstanceIp.GetHref())
	if err := client.Update(objInstanceIp); err != nil {
		return fmt.Errorf("[ResourceInstanceIpRefsDelete] Delete refs for resource InstanceIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objInstanceIp.GetUuid())
	return nil
}

func ResourceInstanceIpSchema() map[string]*schema.Schema {
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
		"instance_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"instance_ip_family": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"instance_ip_mode": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"secondary_ip_tracking_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
		},
		"subnet_uuid": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"instance_ip_secondary": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"instance_ip_local_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"service_instance_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"service_health_check_ip": &schema.Schema{
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
		"virtual_network_refs": &schema.Schema{
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
		"network_ipam_refs": &schema.Schema{
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

func ResourceInstanceIpRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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
		"physical_router_refs": &schema.Schema{
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
		"virtual_router_refs": &schema.Schema{
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

func ResourceInstanceIp() *schema.Resource {
	return &schema.Resource{
		Create: ResourceInstanceIpCreate,
		Read:   ResourceInstanceIpRead,
		Update: ResourceInstanceIpUpdate,
		Delete: ResourceInstanceIpDelete,
		Schema: ResourceInstanceIpSchema(),
	}
}

func ResourceInstanceIpRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceInstanceIpRefsCreate,
		Read:   ResourceInstanceIpRefsRead,
		Update: ResourceInstanceIpRefsUpdate,
		Delete: ResourceInstanceIpRefsDelete,
		Schema: ResourceInstanceIpRefsSchema(),
	}
}

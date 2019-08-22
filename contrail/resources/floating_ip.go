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

func SetFloatingIpFromResource(object *FloatingIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetFloatingIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("floating_ip_address"); ok {
		object.SetFloatingIpAddress(val.(string))
	}
	if val, ok := d.GetOk("floating_ip_is_virtual_ip"); ok {
		object.SetFloatingIpIsVirtualIp(val.(bool))
	}
	if val, ok := d.GetOk("floating_ip_fixed_ip_address"); ok {
		object.SetFloatingIpFixedIpAddress(val.(string))
	}
	if val, ok := d.GetOk("floating_ip_address_family"); ok {
		object.SetFloatingIpAddressFamily(val.(string))
	}
	if val, ok := d.GetOk("floating_ip_port_mappings_enable"); ok {
		object.SetFloatingIpPortMappingsEnable(val.(bool))
	}
	if val, ok := d.GetOk("floating_ip_port_mappings"); ok {
		member := new(PortMappings)
		SetPortMappingsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetFloatingIpPortMappings(member)
	}
	if val, ok := d.GetOk("floating_ip_traffic_direction"); ok {
		object.SetFloatingIpTrafficDirection(val.(string))
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

func SetRefsFloatingIpFromResource(object *FloatingIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsFloatingIpFromResource] key = %v, prefix = %v", key, prefix)
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

func SetReqRefsFloatingIpFromResource(object *FloatingIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsFloatingIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("project_refs"); ok {
		log.Printf("Got ref project_refs -- will call: object.AddProject(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("project", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving project by Uuid = %v as ref for Project on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddProject(refObj.(*Project))
		}
	}

	return nil
}

func DeleteRefsFloatingIpFromResource(object *FloatingIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsFloatingIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.DeleteVirtualMachineInterface(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualMachineInterface(refId.(string))
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

func WriteFloatingIpToResource(object FloatingIp, d *schema.ResourceData, m interface{}) {

	d.Set("floating_ip_address", object.GetFloatingIpAddress())
	d.Set("floating_ip_is_virtual_ip", object.GetFloatingIpIsVirtualIp())
	d.Set("floating_ip_fixed_ip_address", object.GetFloatingIpFixedIpAddress())
	d.Set("floating_ip_address_family", object.GetFloatingIpAddressFamily())
	d.Set("floating_ip_port_mappings_enable", object.GetFloatingIpPortMappingsEnable())
	floating_ip_port_mappingsObj := object.GetFloatingIpPortMappings()
	d.Set("floating_ip_port_mappings", TakePortMappingsAsMap(&floating_ip_port_mappingsObj))
	d.Set("floating_ip_traffic_direction", object.GetFloatingIpTrafficDirection())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

	if ref, err := object.GetProjectRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("project_refs", refList)
	}
}

func TakeFloatingIpAsMap(object *FloatingIp) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["floating_ip_address"] = object.GetFloatingIpAddress()
	omap["floating_ip_is_virtual_ip"] = object.GetFloatingIpIsVirtualIp()
	omap["floating_ip_fixed_ip_address"] = object.GetFloatingIpFixedIpAddress()
	omap["floating_ip_address_family"] = object.GetFloatingIpAddressFamily()
	omap["floating_ip_port_mappings_enable"] = object.GetFloatingIpPortMappingsEnable()
	floating_ip_port_mappingsObj := object.GetFloatingIpPortMappings()
	omap["floating_ip_port_mappings"] = TakePortMappingsAsMap(&floating_ip_port_mappingsObj)
	omap["floating_ip_traffic_direction"] = object.GetFloatingIpTrafficDirection()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateFloatingIpFromResource(object *FloatingIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("floating_ip_address") {
		if val, ok := d.GetOk("floating_ip_address"); ok {
			object.SetFloatingIpAddress(val.(string))
		}
	}
	if d.HasChange("floating_ip_is_virtual_ip") {
		if val, ok := d.GetOk("floating_ip_is_virtual_ip"); ok {
			object.SetFloatingIpIsVirtualIp(val.(bool))
		}
	}
	if d.HasChange("floating_ip_fixed_ip_address") {
		if val, ok := d.GetOk("floating_ip_fixed_ip_address"); ok {
			object.SetFloatingIpFixedIpAddress(val.(string))
		}
	}
	if d.HasChange("floating_ip_address_family") {
		if val, ok := d.GetOk("floating_ip_address_family"); ok {
			object.SetFloatingIpAddressFamily(val.(string))
		}
	}
	if d.HasChange("floating_ip_port_mappings_enable") {
		if val, ok := d.GetOk("floating_ip_port_mappings_enable"); ok {
			object.SetFloatingIpPortMappingsEnable(val.(bool))
		}
	}
	if d.HasChange("floating_ip_port_mappings") {
		if val, ok := d.GetOk("floating_ip_port_mappings"); ok {
			member := new(PortMappings)
			SetPortMappingsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetFloatingIpPortMappings(member)
		}
	}
	if d.HasChange("floating_ip_traffic_direction") {
		if val, ok := d.GetOk("floating_ip_traffic_direction"); ok {
			object.SetFloatingIpTrafficDirection(val.(string))
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
	if d.HasChange("project_refs") {
		if val, ok := d.GetOk("project_refs"); ok {
			log.Printf("Got ref project_refs -- will call: object.AddProject(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("project", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddProject(refObj.(*Project))
			}
		}
	}

}

func ResourceFloatingIpCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFloatingIpCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(FloatingIp)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceFloatingIpCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "FloatingIp", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetFloatingIpFromResource(object, d, m)

	if err := SetReqRefsFloatingIpFromResource(object, d, m); err != nil {
		return fmt.Errorf("[ResourceFloatingIpReqRefsCreate] Set required refs on object FloatingIp on %v (%v)", client.GetServer(), err)
	}

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceFloatingIpCreate] Creation of resource FloatingIp on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceFloatingIpRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFloatingIpRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFloatingIpRefsCreate] Missing 'uuid' field for resource FloatingIp")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("floating-ip", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpRefsCreate] Retrieving FloatingIp with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFloatingIp := obj.(*FloatingIp) // Fully set by Contrail backend
	if err := SetRefsFloatingIpFromResource(objFloatingIp, d, m); err != nil {
		return fmt.Errorf("[ResourceFloatingIpRefsCreate] Set refs on object FloatingIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFloatingIp.GetHref())
	if err := client.Update(objFloatingIp); err != nil {
		return fmt.Errorf("[ResourceFloatingIpRefsCreate] Update refs for resource FloatingIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFloatingIp.GetUuid())
	return nil
}

func ResourceFloatingIpRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("floating-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpRead] Read resource floating-ip on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*FloatingIp)
	WriteFloatingIpToResource(*object, d, m)
	return nil
}

func ResourceFloatingIpRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpRefsREAD")
	return nil
}

func ResourceFloatingIpUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("floating-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpResourceUpdate] Retrieving FloatingIp with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*FloatingIp)
	UpdateFloatingIpFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceFloatingIpUpdate] Update of resource FloatingIp on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceFloatingIpRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpRefsUpdate")
	return nil
}

func ResourceFloatingIpDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("floating-ip", d.Id()); err != nil {
		return fmt.Errorf("[ResourceFloatingIpDelete] Deletion of resource floating-ip on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceFloatingIpRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFloatingIpRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFloatingIpRefsDelete] Missing 'uuid' field for resource FloatingIp")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("floating-ip", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpRefsDelete] Retrieving FloatingIp with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFloatingIp := obj.(*FloatingIp) // Fully set by Contrail backend
	if err := DeleteRefsFloatingIpFromResource(objFloatingIp, d, m); err != nil {
		return fmt.Errorf("[ResourceFloatingIpRefsDelete] Set refs on object FloatingIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFloatingIp.GetHref())
	if err := client.Update(objFloatingIp); err != nil {
		return fmt.Errorf("[ResourceFloatingIpRefsDelete] Delete refs for resource FloatingIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFloatingIp.GetUuid())
	return nil
}

func ResourceFloatingIpSchema() map[string]*schema.Schema {
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
		"floating_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"floating_ip_is_virtual_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"floating_ip_fixed_ip_address": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"floating_ip_address_family": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"floating_ip_port_mappings_enable": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"floating_ip_port_mappings": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePortMappings(),
		},
		"floating_ip_traffic_direction": &schema.Schema{
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
		"project_refs": &schema.Schema{
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

func ResourceFloatingIpRefsSchema() map[string]*schema.Schema {
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

func ResourceFloatingIp() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFloatingIpCreate,
		Read:   ResourceFloatingIpRead,
		Update: ResourceFloatingIpUpdate,
		Delete: ResourceFloatingIpDelete,
		Schema: ResourceFloatingIpSchema(),
	}
}

func ResourceFloatingIpRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFloatingIpRefsCreate,
		Read:   ResourceFloatingIpRefsRead,
		Update: ResourceFloatingIpRefsUpdate,
		Delete: ResourceFloatingIpRefsDelete,
		Schema: ResourceFloatingIpRefsSchema(),
	}
}

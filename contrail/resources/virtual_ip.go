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

func SetVirtualIpFromResource(object *VirtualIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetVirtualIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_ip_properties"); ok {
		member := new(VirtualIpType)
		SetVirtualIpTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetVirtualIpProperties(member)
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

func SetRefsVirtualIpFromResource(object *VirtualIp, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsVirtualIpFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("loadbalancer_pool_refs"); ok {
		log.Printf("Got ref loadbalancer_pool_refs -- will call: object.AddLoadbalancerPool(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("loadbalancer-pool", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving loadbalancer-pool by Uuid = %v as ref for LoadbalancerPool on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddLoadbalancerPool(refObj.(*LoadbalancerPool))
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

func WriteVirtualIpToResource(object VirtualIp, d *schema.ResourceData, m interface{}) {

	virtual_ip_propertiesObj := object.GetVirtualIpProperties()
	d.Set("virtual_ip_properties", TakeVirtualIpTypeAsMap(&virtual_ip_propertiesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeVirtualIpAsMap(object *VirtualIp) map[string]interface{} {
	omap := make(map[string]interface{})

	virtual_ip_propertiesObj := object.GetVirtualIpProperties()
	omap["virtual_ip_properties"] = TakeVirtualIpTypeAsMap(&virtual_ip_propertiesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateVirtualIpFromResource(object *VirtualIp, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("virtual_ip_properties") {
		if val, ok := d.GetOk("virtual_ip_properties"); ok {
			member := new(VirtualIpType)
			SetVirtualIpTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetVirtualIpProperties(member)
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

func ResourceVirtualIpCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualIpCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(VirtualIp)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceVirtualIpCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "VirtualIp", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetVirtualIpFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceVirtualIpCreate] Creation of resource VirtualIp on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceVirtualIpRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceVirtualIpRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceVirtualIpRefsCreate] Missing 'uuid' field for resource VirtualIp")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("virtual-ip", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceVirtualIpRefsCreate] Retrieving VirtualIp with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objVirtualIp := obj.(*VirtualIp) // Fully set by Contrail backend
	if err := SetRefsVirtualIpFromResource(objVirtualIp, d, m); err != nil {
		return fmt.Errorf("[ResourceVirtualIpRefsCreate] Set refs on object VirtualIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objVirtualIp.GetHref())
	if err := client.Update(objVirtualIp); err != nil {
		return fmt.Errorf("[ResourceVirtualIpRefsCreate] Update refs for resource VirtualIp (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objVirtualIp.GetUuid())
	return nil
}

func ResourceVirtualIpRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualIpREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("virtual-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualIpRead] Read resource virtual-ip on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*VirtualIp)
	WriteVirtualIpToResource(*object, d, m)
	return nil
}

func ResourceVirtualIpRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualIpRefsREAD")
	return nil
}

func ResourceVirtualIpUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualIpUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("virtual-ip", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceVirtualIpResourceUpdate] Retrieving VirtualIp with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*VirtualIp)
	UpdateVirtualIpFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceVirtualIpUpdate] Update of resource VirtualIp on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceVirtualIpRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualIpRefsUpdate")
	return nil
}

func ResourceVirtualIpDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualIpDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("virtual-ip", d.Id()); err != nil {
		return fmt.Errorf("[ResourceVirtualIpDelete] Deletion of resource virtual-ip on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceVirtualIpRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceVirtualIpRefsDelete: %v", d.Id())
	return nil
}

func ResourceVirtualIpSchema() map[string]*schema.Schema {
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
		"virtual_ip_properties": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualIpType(),
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

func ResourceVirtualIpRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"loadbalancer_pool_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLoadbalancerPool(),
		},
		"virtual_machine_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachineInterface(),
		},
	}
}

func ResourceVirtualIp() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualIpCreate,
		Read:   ResourceVirtualIpRead,
		Update: ResourceVirtualIpUpdate,
		Delete: ResourceVirtualIpDelete,
		Schema: ResourceVirtualIpSchema(),
	}
}

func ResourceVirtualIpRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceVirtualIpRefsCreate,
		Read:   ResourceVirtualIpRefsRead,
		Update: ResourceVirtualIpRefsUpdate,
		Delete: ResourceVirtualIpRefsDelete,
		Schema: ResourceVirtualIpRefsSchema(),
	}
}

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

func SetLogicalInterfaceFromResource(object *LogicalInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLogicalInterfaceFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("logical_interface_vlan_tag"); ok {
		object.SetLogicalInterfaceVlanTag(val.(int))
	}
	if val, ok := d.GetOk("logical_interface_type"); ok {
		object.SetLogicalInterfaceType(val.(string))
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

func SetRefsLogicalInterfaceFromResource(object *LogicalInterface, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLogicalInterfaceFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteLogicalInterfaceToResource(object LogicalInterface, d *schema.ResourceData, m interface{}) {

	d.Set("logical_interface_vlan_tag", object.GetLogicalInterfaceVlanTag())
	d.Set("logical_interface_type", object.GetLogicalInterfaceType())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeLogicalInterfaceAsMap(object *LogicalInterface) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["logical_interface_vlan_tag"] = object.GetLogicalInterfaceVlanTag()
	omap["logical_interface_type"] = object.GetLogicalInterfaceType()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLogicalInterfaceFromResource(object *LogicalInterface, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("logical_interface_vlan_tag") {
		if val, ok := d.GetOk("logical_interface_vlan_tag"); ok {
			object.SetLogicalInterfaceVlanTag(val.(int))
		}
	}
	if d.HasChange("logical_interface_type") {
		if val, ok := d.GetOk("logical_interface_type"); ok {
			object.SetLogicalInterfaceType(val.(string))
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

func ResourceLogicalInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLogicalInterfaceCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LogicalInterface)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceLogicalInterfaceCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LogicalInterface", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLogicalInterfaceFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceCreate] Creation of resource LogicalInterface on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLogicalInterfaceRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLogicalInterfaceRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLogicalInterfaceRefsCreate] Missing 'uuid' field for resource LogicalInterface")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("logical-interface", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceRefsCreate] Retrieving LogicalInterface with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLogicalInterface := obj.(*LogicalInterface) // Fully set by Contrail backend
	if err := SetRefsLogicalInterfaceFromResource(objLogicalInterface, d, m); err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceRefsCreate] Set refs on object LogicalInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLogicalInterface.GetHref())
	if err := client.Update(objLogicalInterface); err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceRefsCreate] Update refs for resource LogicalInterface (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLogicalInterface.GetUuid())
	return nil
}

func ResourceLogicalInterfaceRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalInterfaceREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("logical-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceRead] Read resource logical-interface on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LogicalInterface)
	WriteLogicalInterfaceToResource(*object, d, m)
	return nil
}

func ResourceLogicalInterfaceRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalInterfaceRefsREAD")
	return nil
}

func ResourceLogicalInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalInterfaceUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("logical-interface", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceResourceUpdate] Retrieving LogicalInterface with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LogicalInterface)
	UpdateLogicalInterfaceFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceUpdate] Update of resource LogicalInterface on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLogicalInterfaceRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalInterfaceRefsUpdate")
	return nil
}

func ResourceLogicalInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalInterfaceDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("logical-interface", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLogicalInterfaceDelete] Deletion of resource logical-interface on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLogicalInterfaceRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalInterfaceRefsDelete: %v", d.Id())
	return nil
}

func ResourceLogicalInterfaceSchema() map[string]*schema.Schema {
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
		"logical_interface_vlan_tag": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"logical_interface_type": &schema.Schema{
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

func ResourceLogicalInterfaceRefsSchema() map[string]*schema.Schema {
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

func ResourceLogicalInterface() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLogicalInterfaceCreate,
		Read:   ResourceLogicalInterfaceRead,
		Update: ResourceLogicalInterfaceUpdate,
		Delete: ResourceLogicalInterfaceDelete,
		Schema: ResourceLogicalInterfaceSchema(),
	}
}

func ResourceLogicalInterfaceRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLogicalInterfaceRefsCreate,
		Read:   ResourceLogicalInterfaceRefsRead,
		Update: ResourceLogicalInterfaceRefsUpdate,
		Delete: ResourceLogicalInterfaceRefsDelete,
		Schema: ResourceLogicalInterfaceRefsSchema(),
	}
}

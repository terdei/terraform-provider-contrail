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

func SetSubnetFromResource(object *Subnet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetSubnetFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("subnet_ip_prefix"); ok {
		member := new(SubnetType)
		SetSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetSubnetIpPrefix(member)
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

func SetRefsSubnetFromResource(object *Subnet, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsSubnetFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteSubnetToResource(object Subnet, d *schema.ResourceData, m interface{}) {

	subnet_ip_prefixObj := object.GetSubnetIpPrefix()
	d.Set("subnet_ip_prefix", TakeSubnetTypeAsMap(&subnet_ip_prefixObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeSubnetAsMap(object *Subnet) map[string]interface{} {
	omap := make(map[string]interface{})

	subnet_ip_prefixObj := object.GetSubnetIpPrefix()
	omap["subnet_ip_prefix"] = TakeSubnetTypeAsMap(&subnet_ip_prefixObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateSubnetFromResource(object *Subnet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("subnet_ip_prefix") {
		if val, ok := d.GetOk("subnet_ip_prefix"); ok {
			member := new(SubnetType)
			SetSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetSubnetIpPrefix(member)
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

func ResourceSubnetCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceSubnetCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Subnet)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceSubnetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Subnet", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetSubnetFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceSubnetCreate] Creation of resource Subnet on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceSubnetRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceSubnetRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceSubnetRefsCreate] Missing 'uuid' field for resource Subnet")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("subnet", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceSubnetRefsCreate] Retrieving Subnet with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objSubnet := obj.(*Subnet) // Fully set by Contrail backend
	if err := SetRefsSubnetFromResource(objSubnet, d, m); err != nil {
		return fmt.Errorf("[ResourceSubnetRefsCreate] Set refs on object Subnet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objSubnet.GetHref())
	if err := client.Update(objSubnet); err != nil {
		return fmt.Errorf("[ResourceSubnetRefsCreate] Update refs for resource Subnet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objSubnet.GetUuid())
	return nil
}

func ResourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSubnetREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("subnet", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceSubnetRead] Read resource subnet on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Subnet)
	WriteSubnetToResource(*object, d, m)
	return nil
}

func ResourceSubnetRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSubnetRefsREAD")
	return nil
}

func ResourceSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSubnetUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("subnet", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceSubnetResourceUpdate] Retrieving Subnet with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Subnet)
	UpdateSubnetFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceSubnetUpdate] Update of resource Subnet on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceSubnetRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSubnetRefsUpdate")
	return nil
}

func ResourceSubnetDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSubnetDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("subnet", d.Id()); err != nil {
		return fmt.Errorf("[ResourceSubnetDelete] Deletion of resource subnet on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceSubnetRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSubnetRefsDelete: %v", d.Id())
	return nil
}

func ResourceSubnetSchema() map[string]*schema.Schema {
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
		"subnet_ip_prefix": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSubnetType(),
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

func ResourceSubnetRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_machine_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachineInterface(),
		},
	}
}

func ResourceSubnet() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSubnetCreate,
		Read:   ResourceSubnetRead,
		Update: ResourceSubnetUpdate,
		Delete: ResourceSubnetDelete,
		Schema: ResourceSubnetSchema(),
	}
}

func ResourceSubnetRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSubnetRefsCreate,
		Read:   ResourceSubnetRefsRead,
		Update: ResourceSubnetRefsUpdate,
		Delete: ResourceSubnetRefsDelete,
		Schema: ResourceSubnetRefsSchema(),
	}
}

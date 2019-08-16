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

func SetNetworkIpamFromResource(object *NetworkIpam, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetNetworkIpamFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("network_ipam_mgmt"); ok {
		member := new(IpamType)
		SetIpamTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetNetworkIpamMgmt(member)
	}
	if val, ok := d.GetOk("ipam_subnets"); ok {
		member := new(IpamSubnets)
		SetIpamSubnetsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetIpamSubnets(member)
	}
	if val, ok := d.GetOk("ipam_subnet_method"); ok {
		object.SetIpamSubnetMethod(val.(string))
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

func SetRefsNetworkIpamFromResource(object *NetworkIpam, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsNetworkIpamFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_dns_refs"); ok {
		log.Printf("Got ref virtual_dns_refs -- will call: object.AddVirtualDns(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-DNS", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-DNS by Uuid = %v as ref for VirtualDns on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualDns(refObj.(*VirtualDns))
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

func DeleteRefsNetworkIpamFromResource(object *NetworkIpam, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsNetworkIpamFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_dns_refs"); ok {
		log.Printf("Got ref virtual_dns_refs -- will call: object.DeleteVirtualDns(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualDns(refId.(string))
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

func WriteNetworkIpamToResource(object NetworkIpam, d *schema.ResourceData, m interface{}) {

	network_ipam_mgmtObj := object.GetNetworkIpamMgmt()
	d.Set("network_ipam_mgmt", TakeIpamTypeAsMap(&network_ipam_mgmtObj))
	ipam_subnetsObj := object.GetIpamSubnets()
	d.Set("ipam_subnets", TakeIpamSubnetsAsMap(&ipam_subnetsObj))
	d.Set("ipam_subnet_method", object.GetIpamSubnetMethod())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeNetworkIpamAsMap(object *NetworkIpam) map[string]interface{} {
	omap := make(map[string]interface{})

	network_ipam_mgmtObj := object.GetNetworkIpamMgmt()
	omap["network_ipam_mgmt"] = TakeIpamTypeAsMap(&network_ipam_mgmtObj)
	ipam_subnetsObj := object.GetIpamSubnets()
	omap["ipam_subnets"] = TakeIpamSubnetsAsMap(&ipam_subnetsObj)
	omap["ipam_subnet_method"] = object.GetIpamSubnetMethod()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateNetworkIpamFromResource(object *NetworkIpam, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("network_ipam_mgmt") {
		if val, ok := d.GetOk("network_ipam_mgmt"); ok {
			member := new(IpamType)
			SetIpamTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetNetworkIpamMgmt(member)
		}
	}
	if d.HasChange("ipam_subnets") {
		if val, ok := d.GetOk("ipam_subnets"); ok {
			member := new(IpamSubnets)
			SetIpamSubnetsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetIpamSubnets(member)
		}
	}
	if d.HasChange("ipam_subnet_method") {
		if val, ok := d.GetOk("ipam_subnet_method"); ok {
			object.SetIpamSubnetMethod(val.(string))
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

func ResourceNetworkIpamCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkIpamCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(NetworkIpam)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceNetworkIpamCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "NetworkIpam", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetNetworkIpamFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamCreate] Creation of resource NetworkIpam on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceNetworkIpamRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkIpamRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNetworkIpamRefsCreate] Missing 'uuid' field for resource NetworkIpam")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("network-ipam", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRefsCreate] Retrieving NetworkIpam with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNetworkIpam := obj.(*NetworkIpam) // Fully set by Contrail backend
	if err := SetRefsNetworkIpamFromResource(objNetworkIpam, d, m); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRefsCreate] Set refs on object NetworkIpam (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNetworkIpam.GetHref())
	if err := client.Update(objNetworkIpam); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRefsCreate] Update refs for resource NetworkIpam (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNetworkIpam.GetUuid())
	return nil
}

func ResourceNetworkIpamRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkIpamREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("network-ipam", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRead] Read resource network-ipam on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*NetworkIpam)
	WriteNetworkIpamToResource(*object, d, m)
	return nil
}

func ResourceNetworkIpamRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkIpamRefsREAD")
	return nil
}

func ResourceNetworkIpamUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkIpamUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("network-ipam", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkIpamResourceUpdate] Retrieving NetworkIpam with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*NetworkIpam)
	UpdateNetworkIpamFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamUpdate] Update of resource NetworkIpam on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceNetworkIpamRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkIpamRefsUpdate")
	return nil
}

func ResourceNetworkIpamDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkIpamDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("network-ipam", d.Id()); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamDelete] Deletion of resource network-ipam on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceNetworkIpamRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkIpamRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNetworkIpamRefsDelete] Missing 'uuid' field for resource NetworkIpam")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("network-ipam", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRefsDelete] Retrieving NetworkIpam with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNetworkIpam := obj.(*NetworkIpam) // Fully set by Contrail backend
	if err := DeleteRefsNetworkIpamFromResource(objNetworkIpam, d, m); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRefsDelete] Set refs on object NetworkIpam (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNetworkIpam.GetHref())
	if err := client.Update(objNetworkIpam); err != nil {
		return fmt.Errorf("[ResourceNetworkIpamRefsDelete] Delete refs for resource NetworkIpam (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNetworkIpam.GetUuid())
	return nil
}

func ResourceNetworkIpamSchema() map[string]*schema.Schema {
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
		"network_ipam_mgmt": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIpamType(),
		},
		"ipam_subnets": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceIpamSubnets(),
		},
		"ipam_subnet_method": &schema.Schema{
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

func ResourceNetworkIpamRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_dns_refs": &schema.Schema{
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

func ResourceNetworkIpam() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNetworkIpamCreate,
		Read:   ResourceNetworkIpamRead,
		Update: ResourceNetworkIpamUpdate,
		Delete: ResourceNetworkIpamDelete,
		Schema: ResourceNetworkIpamSchema(),
	}
}

func ResourceNetworkIpamRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNetworkIpamRefsCreate,
		Read:   ResourceNetworkIpamRefsRead,
		Update: ResourceNetworkIpamRefsUpdate,
		Delete: ResourceNetworkIpamRefsDelete,
		Schema: ResourceNetworkIpamRefsSchema(),
	}
}

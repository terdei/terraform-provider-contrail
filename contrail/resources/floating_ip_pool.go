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

func SetFloatingIpPoolFromResource(object *FloatingIpPool, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetFloatingIpPoolFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("floating_ip_pool_subnets"); ok {
		member := new(FloatingIpPoolSubnetType)
		SetFloatingIpPoolSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetFloatingIpPoolSubnets(member)
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

func SetRefsFloatingIpPoolFromResource(object *FloatingIpPool, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsFloatingIpPoolFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsFloatingIpPoolFromResource(object *FloatingIpPool, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsFloatingIpPoolFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteFloatingIpPoolToResource(object FloatingIpPool, d *schema.ResourceData, m interface{}) {

	floating_ip_pool_subnetsObj := object.GetFloatingIpPoolSubnets()
	d.Set("floating_ip_pool_subnets", TakeFloatingIpPoolSubnetTypeAsMap(&floating_ip_pool_subnetsObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeFloatingIpPoolAsMap(object *FloatingIpPool) map[string]interface{} {
	omap := make(map[string]interface{})

	floating_ip_pool_subnetsObj := object.GetFloatingIpPoolSubnets()
	omap["floating_ip_pool_subnets"] = TakeFloatingIpPoolSubnetTypeAsMap(&floating_ip_pool_subnetsObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateFloatingIpPoolFromResource(object *FloatingIpPool, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("floating_ip_pool_subnets") {
		if val, ok := d.GetOk("floating_ip_pool_subnets"); ok {
			member := new(FloatingIpPoolSubnetType)
			SetFloatingIpPoolSubnetTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetFloatingIpPoolSubnets(member)
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

func ResourceFloatingIpPoolCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFloatingIpPoolCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(FloatingIpPool)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceFloatingIpPoolCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "FloatingIpPool", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetFloatingIpPoolFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolCreate] Creation of resource FloatingIpPool on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceFloatingIpPoolRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFloatingIpPoolRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsCreate] Missing 'uuid' field for resource FloatingIpPool")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("floating-ip-pool", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsCreate] Retrieving FloatingIpPool with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFloatingIpPool := obj.(*FloatingIpPool) // Fully set by Contrail backend
	if err := SetRefsFloatingIpPoolFromResource(objFloatingIpPool, d, m); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsCreate] Set refs on object FloatingIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFloatingIpPool.GetHref())
	if err := client.Update(objFloatingIpPool); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsCreate] Update refs for resource FloatingIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFloatingIpPool.GetUuid())
	return nil
}

func ResourceFloatingIpPoolRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpPoolREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("floating-ip-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRead] Read resource floating-ip-pool on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*FloatingIpPool)
	WriteFloatingIpPoolToResource(*object, d, m)
	return nil
}

func ResourceFloatingIpPoolRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpPoolRefsREAD")
	return nil
}

func ResourceFloatingIpPoolUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpPoolUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("floating-ip-pool", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolResourceUpdate] Retrieving FloatingIpPool with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*FloatingIpPool)
	UpdateFloatingIpPoolFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolUpdate] Update of resource FloatingIpPool on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceFloatingIpPoolRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpPoolRefsUpdate")
	return nil
}

func ResourceFloatingIpPoolDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFloatingIpPoolDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("floating-ip-pool", d.Id()); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolDelete] Deletion of resource floating-ip-pool on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceFloatingIpPoolRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFloatingIpPoolRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsDelete] Missing 'uuid' field for resource FloatingIpPool")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("floating-ip-pool", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsDelete] Retrieving FloatingIpPool with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFloatingIpPool := obj.(*FloatingIpPool) // Fully set by Contrail backend
	if err := DeleteRefsFloatingIpPoolFromResource(objFloatingIpPool, d, m); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsDelete] Set refs on object FloatingIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFloatingIpPool.GetHref())
	if err := client.Update(objFloatingIpPool); err != nil {
		return fmt.Errorf("[ResourceFloatingIpPoolRefsDelete] Delete refs for resource FloatingIpPool (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFloatingIpPool.GetUuid())
	return nil
}

func ResourceFloatingIpPoolSchema() map[string]*schema.Schema {
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
		"floating_ip_pool_subnets": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFloatingIpPoolSubnetType(),
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

func ResourceFloatingIpPoolRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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

func ResourceFloatingIpPool() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFloatingIpPoolCreate,
		Read:   ResourceFloatingIpPoolRead,
		Update: ResourceFloatingIpPoolUpdate,
		Delete: ResourceFloatingIpPoolDelete,
		Schema: ResourceFloatingIpPoolSchema(),
	}
}

func ResourceFloatingIpPoolRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFloatingIpPoolRefsCreate,
		Read:   ResourceFloatingIpPoolRefsRead,
		Update: ResourceFloatingIpPoolRefsUpdate,
		Delete: ResourceFloatingIpPoolRefsDelete,
		Schema: ResourceFloatingIpPoolRefsSchema(),
	}
}

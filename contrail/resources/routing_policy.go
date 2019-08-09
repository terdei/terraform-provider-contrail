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

func SetRoutingPolicyFromResource(object *RoutingPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetRoutingPolicyFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsRoutingPolicyFromResource(object *RoutingPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsRoutingPolicyFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-instance", refId.(string))
			dataObj := new(RoutingPolicyServiceInstanceType)
			SetRoutingPolicyServiceInstanceTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-instance by Uuid = %v as ref for ServiceInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceInstance(refObj.(*ServiceInstance), *dataObj)
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

func WriteRoutingPolicyToResource(object RoutingPolicy, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeRoutingPolicyAsMap(object *RoutingPolicy) map[string]interface{} {
	omap := make(map[string]interface{})

	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateRoutingPolicyFromResource(object *RoutingPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
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

func ResourceRoutingPolicyCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRoutingPolicyCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(RoutingPolicy)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceRoutingPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "RoutingPolicy", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetRoutingPolicyFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyCreate] Creation of resource RoutingPolicy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceRoutingPolicyRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRoutingPolicyRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRoutingPolicyRefsCreate] Missing 'uuid' field for resource RoutingPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("routing-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyRefsCreate] Retrieving RoutingPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRoutingPolicy := obj.(*RoutingPolicy) // Fully set by Contrail backend
	if err := SetRefsRoutingPolicyFromResource(objRoutingPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyRefsCreate] Set refs on object RoutingPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRoutingPolicy.GetHref())
	if err := client.Update(objRoutingPolicy); err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyRefsCreate] Update refs for resource RoutingPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRoutingPolicy.GetUuid())
	return nil
}

func ResourceRoutingPolicyRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingPolicyREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("routing-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyRead] Read resource routing-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RoutingPolicy)
	WriteRoutingPolicyToResource(*object, d, m)
	return nil
}

func ResourceRoutingPolicyRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingPolicyRefsREAD")
	return nil
}

func ResourceRoutingPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingPolicyUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("routing-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyResourceUpdate] Retrieving RoutingPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RoutingPolicy)
	UpdateRoutingPolicyFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyUpdate] Update of resource RoutingPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRoutingPolicyRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingPolicyRefsUpdate")
	return nil
}

func ResourceRoutingPolicyDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingPolicyDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("routing-policy", d.Id()); err != nil {
		return fmt.Errorf("[ResourceRoutingPolicyDelete] Deletion of resource routing-policy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceRoutingPolicyRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRoutingPolicyRefsDelete: %v", d.Id())
	return nil
}

func ResourceRoutingPolicySchema() map[string]*schema.Schema {
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

func ResourceRoutingPolicyRefsSchema() map[string]*schema.Schema {
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
					"attr": &schema.Schema{
						Type:     schema.TypeList,
						Elem:     ResourceRoutingPolicyServiceInstanceType(),
						Required: true,
					},
				},
			},
		},
		"tag_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTag(),
		},
	}
}

func ResourceRoutingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRoutingPolicyCreate,
		Read:   ResourceRoutingPolicyRead,
		Update: ResourceRoutingPolicyUpdate,
		Delete: ResourceRoutingPolicyDelete,
		Schema: ResourceRoutingPolicySchema(),
	}
}

func ResourceRoutingPolicyRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRoutingPolicyRefsCreate,
		Read:   ResourceRoutingPolicyRefsRead,
		Update: ResourceRoutingPolicyRefsUpdate,
		Delete: ResourceRoutingPolicyRefsDelete,
		Schema: ResourceRoutingPolicyRefsSchema(),
	}
}

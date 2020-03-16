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

func SetE2ServiceProviderFromResource(object *E2ServiceProvider, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetE2ServiceProviderFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("e2_service_provider_promiscuous"); ok {
		object.SetE2ServiceProviderPromiscuous(val.(bool))
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

func SetRefsE2ServiceProviderFromResource(object *E2ServiceProvider, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsE2ServiceProviderFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("peering_policy_refs"); ok {
		log.Printf("Got ref peering_policy_refs -- will call: object.AddPeeringPolicy(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("peering-policy", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving peering-policy by Uuid = %v as ref for PeeringPolicy on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddPeeringPolicy(refObj.(*PeeringPolicy))
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

func DeleteRefsE2ServiceProviderFromResource(object *E2ServiceProvider, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsE2ServiceProviderFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("peering_policy_refs"); ok {
		log.Printf("Got ref peering_policy_refs -- will call: object.DeletePeeringPolicy(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeletePeeringPolicy(refId.(string))
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

func WriteE2ServiceProviderToResource(object E2ServiceProvider, d *schema.ResourceData, m interface{}) {

	d.Set("e2_service_provider_promiscuous", object.GetE2ServiceProviderPromiscuous())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteE2ServiceProviderRefsToResource(object E2ServiceProvider, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetPeeringPolicyRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("peering_policy_refs", refList)
	}
	if ref, err := object.GetPhysicalRouterRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("physical_router_refs", refList)
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

func TakeE2ServiceProviderAsMap(object *E2ServiceProvider) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["e2_service_provider_promiscuous"] = object.GetE2ServiceProviderPromiscuous()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateE2ServiceProviderFromResource(object *E2ServiceProvider, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("e2_service_provider_promiscuous") {
		if val, ok := d.GetOk("e2_service_provider_promiscuous"); ok {
			object.SetE2ServiceProviderPromiscuous(val.(bool))
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

func UpdateE2ServiceProviderRefsFromResource(object *E2ServiceProvider, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("peering_policy_refs") {
		object.ClearPeeringPolicy()
		if val, ok := d.GetOk("peering_policy_refs"); ok {
			log.Printf("Got ref peering_policy_refs -- will call: object.AddPeeringPolicy(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("peering-policy", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddPeeringPolicy(refObj.(*PeeringPolicy))
			}
		}
	}
	if d.HasChange("physical_router_refs") {
		object.ClearPhysicalRouter()
		if val, ok := d.GetOk("physical_router_refs"); ok {
			log.Printf("Got ref physical_router_refs -- will call: object.AddPhysicalRouter(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("physical-router", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddPhysicalRouter(refObj.(*PhysicalRouter))
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

func ResourceE2ServiceProviderCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceE2ServiceProviderCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(E2ServiceProvider)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceE2ServiceProviderCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "E2ServiceProvider", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceE2ServiceProviderCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "E2ServiceProvider", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetE2ServiceProviderFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderCreate] Creation of resource E2ServiceProvider on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceE2ServiceProviderRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceE2ServiceProviderRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsCreate] Missing 'uuid' field for resource E2ServiceProvider")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("e2-service-provider", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsCreate] Retrieving E2ServiceProvider with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objE2ServiceProvider := obj.(*E2ServiceProvider) // Fully set by Contrail backend
	if err := SetRefsE2ServiceProviderFromResource(objE2ServiceProvider, d, m); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsCreate] Set refs on object E2ServiceProvider (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objE2ServiceProvider.GetHref())
	if err := client.Update(objE2ServiceProvider); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsCreate] Update refs for resource E2ServiceProvider (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objE2ServiceProvider.GetUuid())
	return nil
}

func ResourceE2ServiceProviderRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceE2ServiceProviderRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("e2-service-provider", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRead] Read resource e2-service-provider on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*E2ServiceProvider)
	WriteE2ServiceProviderToResource(*object, d, m)
	return nil
}

func ResourceE2ServiceProviderRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceE2ServiceProviderRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("e2-service-provider", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsRead] Read resource e2-service-provider on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*E2ServiceProvider)
	WriteE2ServiceProviderRefsToResource(*object, d, m)
	return nil
}

func ResourceE2ServiceProviderUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceE2ServiceProviderUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("e2-service-provider", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderUpdate] Retrieving E2ServiceProvider with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*E2ServiceProvider)
	UpdateE2ServiceProviderFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderUpdate] Update of resource E2ServiceProvider on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceE2ServiceProviderRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceE2ServiceProviderRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("e2-service-provider", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsUpdate] Retrieving E2ServiceProvider with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*E2ServiceProvider)
	UpdateE2ServiceProviderRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsUpdate] Update of resource E2ServiceProvider on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceE2ServiceProviderDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceE2ServiceProviderDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("e2-service-provider", d.Id()); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderDelete] Deletion of resource e2-service-provider on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceE2ServiceProviderRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceE2ServiceProviderRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsDelete] Missing 'uuid' field for resource E2ServiceProvider")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("e2-service-provider", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsDelete] Retrieving E2ServiceProvider with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objE2ServiceProvider := obj.(*E2ServiceProvider) // Fully set by Contrail backend
	if err := DeleteRefsE2ServiceProviderFromResource(objE2ServiceProvider, d, m); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsDelete] Set refs on object E2ServiceProvider (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objE2ServiceProvider.GetHref())
	if err := client.Update(objE2ServiceProvider); err != nil {
		return fmt.Errorf("[ResourceE2ServiceProviderRefsDelete] Delete refs for resource E2ServiceProvider (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objE2ServiceProvider.GetUuid())
	return nil
}

func ResourceE2ServiceProviderSchema() map[string]*schema.Schema {
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
		"e2_service_provider_promiscuous": &schema.Schema{
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
	}
}

func ResourceE2ServiceProviderRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"peering_policy_refs": &schema.Schema{
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

func ResourceE2ServiceProvider() *schema.Resource {
	return &schema.Resource{
		Create: ResourceE2ServiceProviderCreate,
		Read:   ResourceE2ServiceProviderRead,
		Update: ResourceE2ServiceProviderUpdate,
		Delete: ResourceE2ServiceProviderDelete,
		Schema: ResourceE2ServiceProviderSchema(),
	}
}

func ResourceE2ServiceProviderRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceE2ServiceProviderRefsCreate,
		Read:   ResourceE2ServiceProviderRefsRead,
		Update: ResourceE2ServiceProviderRefsUpdate,
		Delete: ResourceE2ServiceProviderRefsDelete,
		Schema: ResourceE2ServiceProviderRefsSchema(),
	}
}

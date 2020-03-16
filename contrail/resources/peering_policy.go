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

func SetPeeringPolicyFromResource(object *PeeringPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetPeeringPolicyFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("peering_service"); ok {
		object.SetPeeringService(val.(string))
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

func SetRefsPeeringPolicyFromResource(object *PeeringPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsPeeringPolicyFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsPeeringPolicyFromResource(object *PeeringPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsPeeringPolicyFromResource] key = %v, prefix = %v", key, prefix)
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

func WritePeeringPolicyToResource(object PeeringPolicy, d *schema.ResourceData, m interface{}) {

	d.Set("peering_service", object.GetPeeringService())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WritePeeringPolicyRefsToResource(object PeeringPolicy, d *schema.ResourceData, m interface{}) {

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

func TakePeeringPolicyAsMap(object *PeeringPolicy) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["peering_service"] = object.GetPeeringService()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdatePeeringPolicyFromResource(object *PeeringPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("peering_service") {
		if val, ok := d.GetOk("peering_service"); ok {
			object.SetPeeringService(val.(string))
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

func UpdatePeeringPolicyRefsFromResource(object *PeeringPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
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

func ResourcePeeringPolicyCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePeeringPolicyCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(PeeringPolicy)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourcePeeringPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "PeeringPolicy", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourcePeeringPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "PeeringPolicy", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetPeeringPolicyFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyCreate] Creation of resource PeeringPolicy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourcePeeringPolicyRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePeeringPolicyRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePeeringPolicyRefsCreate] Missing 'uuid' field for resource PeeringPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("peering-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsCreate] Retrieving PeeringPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPeeringPolicy := obj.(*PeeringPolicy) // Fully set by Contrail backend
	if err := SetRefsPeeringPolicyFromResource(objPeeringPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsCreate] Set refs on object PeeringPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPeeringPolicy.GetHref())
	if err := client.Update(objPeeringPolicy); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsCreate] Update refs for resource PeeringPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPeeringPolicy.GetUuid())
	return nil
}

func ResourcePeeringPolicyRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePeeringPolicyRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("peering-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRead] Read resource peering-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*PeeringPolicy)
	WritePeeringPolicyToResource(*object, d, m)
	return nil
}

func ResourcePeeringPolicyRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePeeringPolicyRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("peering-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsRead] Read resource peering-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*PeeringPolicy)
	WritePeeringPolicyRefsToResource(*object, d, m)
	return nil
}

func ResourcePeeringPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePeeringPolicyUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("peering-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyUpdate] Retrieving PeeringPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*PeeringPolicy)
	UpdatePeeringPolicyFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyUpdate] Update of resource PeeringPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourcePeeringPolicyRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePeeringPolicyRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("peering-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsUpdate] Retrieving PeeringPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*PeeringPolicy)
	UpdatePeeringPolicyRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsUpdate] Update of resource PeeringPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourcePeeringPolicyDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePeeringPolicyDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("peering-policy", d.Id()); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyDelete] Deletion of resource peering-policy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourcePeeringPolicyRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePeeringPolicyRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePeeringPolicyRefsDelete] Missing 'uuid' field for resource PeeringPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("peering-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsDelete] Retrieving PeeringPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPeeringPolicy := obj.(*PeeringPolicy) // Fully set by Contrail backend
	if err := DeleteRefsPeeringPolicyFromResource(objPeeringPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsDelete] Set refs on object PeeringPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPeeringPolicy.GetHref())
	if err := client.Update(objPeeringPolicy); err != nil {
		return fmt.Errorf("[ResourcePeeringPolicyRefsDelete] Delete refs for resource PeeringPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPeeringPolicy.GetUuid())
	return nil
}

func ResourcePeeringPolicySchema() map[string]*schema.Schema {
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
		"peering_service": &schema.Schema{
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

func ResourcePeeringPolicyRefsSchema() map[string]*schema.Schema {
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

func ResourcePeeringPolicy() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePeeringPolicyCreate,
		Read:   ResourcePeeringPolicyRead,
		Update: ResourcePeeringPolicyUpdate,
		Delete: ResourcePeeringPolicyDelete,
		Schema: ResourcePeeringPolicySchema(),
	}
}

func ResourcePeeringPolicyRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePeeringPolicyRefsCreate,
		Read:   ResourcePeeringPolicyRefsRead,
		Update: ResourcePeeringPolicyRefsUpdate,
		Delete: ResourcePeeringPolicyRefsDelete,
		Schema: ResourcePeeringPolicyRefsSchema(),
	}
}

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

func SetNetworkPolicyFromResource(object *NetworkPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetNetworkPolicyFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("network_policy_entries"); ok {
		member := new(PolicyEntriesType)
		SetPolicyEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetNetworkPolicyEntries(member)
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

func SetRefsNetworkPolicyFromResource(object *NetworkPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsNetworkPolicyFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsNetworkPolicyFromResource(object *NetworkPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsNetworkPolicyFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteNetworkPolicyToResource(object NetworkPolicy, d *schema.ResourceData, m interface{}) {

	network_policy_entriesObj := object.GetNetworkPolicyEntries()
	d.Set("network_policy_entries", TakePolicyEntriesTypeAsMap(&network_policy_entriesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteNetworkPolicyRefsToResource(object NetworkPolicy, d *schema.ResourceData, m interface{}) {

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

func TakeNetworkPolicyAsMap(object *NetworkPolicy) map[string]interface{} {
	omap := make(map[string]interface{})

	network_policy_entriesObj := object.GetNetworkPolicyEntries()
	omap["network_policy_entries"] = TakePolicyEntriesTypeAsMap(&network_policy_entriesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateNetworkPolicyFromResource(object *NetworkPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("network_policy_entries") {
		if val, ok := d.GetOk("network_policy_entries"); ok {
			member := new(PolicyEntriesType)
			SetPolicyEntriesTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetNetworkPolicyEntries(member)
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

func UpdateNetworkPolicyRefsFromResource(object *NetworkPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceNetworkPolicyCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkPolicyCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(NetworkPolicy)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceNetworkPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "NetworkPolicy", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceNetworkPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "NetworkPolicy", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetNetworkPolicyFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyCreate] Creation of resource NetworkPolicy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceNetworkPolicyRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkPolicyRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNetworkPolicyRefsCreate] Missing 'uuid' field for resource NetworkPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("network-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsCreate] Retrieving NetworkPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNetworkPolicy := obj.(*NetworkPolicy) // Fully set by Contrail backend
	if err := SetRefsNetworkPolicyFromResource(objNetworkPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsCreate] Set refs on object NetworkPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNetworkPolicy.GetHref())
	if err := client.Update(objNetworkPolicy); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsCreate] Update refs for resource NetworkPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNetworkPolicy.GetUuid())
	return nil
}

func ResourceNetworkPolicyRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkPolicyRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("network-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRead] Read resource network-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*NetworkPolicy)
	WriteNetworkPolicyToResource(*object, d, m)
	return nil
}

func ResourceNetworkPolicyRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkPolicyRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("network-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsRead] Read resource network-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*NetworkPolicy)
	WriteNetworkPolicyRefsToResource(*object, d, m)
	return nil
}

func ResourceNetworkPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkPolicyUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("network-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyUpdate] Retrieving NetworkPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*NetworkPolicy)
	UpdateNetworkPolicyFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyUpdate] Update of resource NetworkPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceNetworkPolicyRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkPolicyRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("network-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsUpdate] Retrieving NetworkPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*NetworkPolicy)
	UpdateNetworkPolicyRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsUpdate] Update of resource NetworkPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceNetworkPolicyDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceNetworkPolicyDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("network-policy", d.Id()); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyDelete] Deletion of resource network-policy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceNetworkPolicyRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceNetworkPolicyRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceNetworkPolicyRefsDelete] Missing 'uuid' field for resource NetworkPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("network-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsDelete] Retrieving NetworkPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objNetworkPolicy := obj.(*NetworkPolicy) // Fully set by Contrail backend
	if err := DeleteRefsNetworkPolicyFromResource(objNetworkPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsDelete] Set refs on object NetworkPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objNetworkPolicy.GetHref())
	if err := client.Update(objNetworkPolicy); err != nil {
		return fmt.Errorf("[ResourceNetworkPolicyRefsDelete] Delete refs for resource NetworkPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objNetworkPolicy.GetUuid())
	return nil
}

func ResourceNetworkPolicySchema() map[string]*schema.Schema {
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
		"network_policy_entries": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourcePolicyEntriesType(),
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

func ResourceNetworkPolicyRefsSchema() map[string]*schema.Schema {
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

func ResourceNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNetworkPolicyCreate,
		Read:   ResourceNetworkPolicyRead,
		Update: ResourceNetworkPolicyUpdate,
		Delete: ResourceNetworkPolicyDelete,
		Schema: ResourceNetworkPolicySchema(),
	}
}

func ResourceNetworkPolicyRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceNetworkPolicyRefsCreate,
		Read:   ResourceNetworkPolicyRefsRead,
		Update: ResourceNetworkPolicyRefsUpdate,
		Delete: ResourceNetworkPolicyRefsDelete,
		Schema: ResourceNetworkPolicyRefsSchema(),
	}
}

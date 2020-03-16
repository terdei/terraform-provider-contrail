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

func SetFirewallPolicyFromResource(object *FirewallPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetFirewallPolicyFromResource] key = %v, prefix = %v", key, prefix)
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

func SetRefsFirewallPolicyFromResource(object *FirewallPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsFirewallPolicyFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("firewall_rule_refs"); ok {
		log.Printf("Got ref firewall_rule_refs -- will call: object.AddFirewallRule(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("firewall-rule", refId.(string))
			dataObj := new(FirewallSequence)
			SetFirewallSequenceFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving firewall-rule by Uuid = %v as ref for FirewallRule on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddFirewallRule(refObj.(*FirewallRule), *dataObj)
		}
	}
	if val, ok := d.GetOk("security_logging_object_refs"); ok {
		log.Printf("Got ref security_logging_object_refs -- will call: object.AddSecurityLoggingObject(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("security-logging-object", refId.(string))
			dataObj := new(SloRateType)
			SetSloRateTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving security-logging-object by Uuid = %v as ref for SecurityLoggingObject on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddSecurityLoggingObject(refObj.(*SecurityLoggingObject), *dataObj)
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

func DeleteRefsFirewallPolicyFromResource(object *FirewallPolicy, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsFirewallPolicyFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("firewall_rule_refs"); ok {
		log.Printf("Got ref firewall_rule_refs -- will call: object.DeleteFirewallRule(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteFirewallRule(refId.(string))
		}
	}
	if val, ok := d.GetOk("security_logging_object_refs"); ok {
		log.Printf("Got ref security_logging_object_refs -- will call: object.DeleteSecurityLoggingObject(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteSecurityLoggingObject(refId.(string))
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

func WriteFirewallPolicyToResource(object FirewallPolicy, d *schema.ResourceData, m interface{}) {

	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteFirewallPolicyRefsToResource(object FirewallPolicy, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetFirewallRuleRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("firewall_rule_refs", refList)
	}
	if ref, err := object.GetSecurityLoggingObjectRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("security_logging_object_refs", refList)
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

func TakeFirewallPolicyAsMap(object *FirewallPolicy) map[string]interface{} {
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

func UpdateFirewallPolicyFromResource(object *FirewallPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func UpdateFirewallPolicyRefsFromResource(object *FirewallPolicy, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("firewall_rule_refs") {
		object.ClearFirewallRule()
		if val, ok := d.GetOk("firewall_rule_refs"); ok {
			log.Printf("Got ref firewall_rule_refs -- will call: object.AddFirewallRule(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("firewall-rule", refId.(string))
				dataObj := new(FirewallSequence)
				SetFirewallSequenceFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddFirewallRule(refObj.(*FirewallRule), *dataObj)
			}
		}
	}
	if d.HasChange("security_logging_object_refs") {
		object.ClearSecurityLoggingObject()
		if val, ok := d.GetOk("security_logging_object_refs"); ok {
			log.Printf("Got ref security_logging_object_refs -- will call: object.AddSecurityLoggingObject(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("security-logging-object", refId.(string))
				dataObj := new(SloRateType)
				SetSloRateTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddSecurityLoggingObject(refObj.(*SecurityLoggingObject), *dataObj)
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

func ResourceFirewallPolicyCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFirewallPolicyCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(FirewallPolicy)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceFirewallPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "FirewallPolicy", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceFirewallPolicyCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "FirewallPolicy", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetFirewallPolicyFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyCreate] Creation of resource FirewallPolicy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceFirewallPolicyRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFirewallPolicyRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFirewallPolicyRefsCreate] Missing 'uuid' field for resource FirewallPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("firewall-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsCreate] Retrieving FirewallPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFirewallPolicy := obj.(*FirewallPolicy) // Fully set by Contrail backend
	if err := SetRefsFirewallPolicyFromResource(objFirewallPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsCreate] Set refs on object FirewallPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFirewallPolicy.GetHref())
	if err := client.Update(objFirewallPolicy); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsCreate] Update refs for resource FirewallPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFirewallPolicy.GetUuid())
	return nil
}

func ResourceFirewallPolicyRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallPolicyRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("firewall-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRead] Read resource firewall-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*FirewallPolicy)
	WriteFirewallPolicyToResource(*object, d, m)
	return nil
}

func ResourceFirewallPolicyRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallPolicyRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("firewall-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsRead] Read resource firewall-policy on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*FirewallPolicy)
	WriteFirewallPolicyRefsToResource(*object, d, m)
	return nil
}

func ResourceFirewallPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallPolicyUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("firewall-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyUpdate] Retrieving FirewallPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*FirewallPolicy)
	UpdateFirewallPolicyFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyUpdate] Update of resource FirewallPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceFirewallPolicyRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallPolicyRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("firewall-policy", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsUpdate] Retrieving FirewallPolicy with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*FirewallPolicy)
	UpdateFirewallPolicyRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsUpdate] Update of resource FirewallPolicy on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceFirewallPolicyDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallPolicyDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("firewall-policy", d.Id()); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyDelete] Deletion of resource firewall-policy on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceFirewallPolicyRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFirewallPolicyRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFirewallPolicyRefsDelete] Missing 'uuid' field for resource FirewallPolicy")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("firewall-policy", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsDelete] Retrieving FirewallPolicy with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFirewallPolicy := obj.(*FirewallPolicy) // Fully set by Contrail backend
	if err := DeleteRefsFirewallPolicyFromResource(objFirewallPolicy, d, m); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsDelete] Set refs on object FirewallPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFirewallPolicy.GetHref())
	if err := client.Update(objFirewallPolicy); err != nil {
		return fmt.Errorf("[ResourceFirewallPolicyRefsDelete] Delete refs for resource FirewallPolicy (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFirewallPolicy.GetUuid())
	return nil
}

func ResourceFirewallPolicySchema() map[string]*schema.Schema {
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

func ResourceFirewallPolicyRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"firewall_rule_refs": &schema.Schema{
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
						Elem:     ResourceFirewallSequence(),
						Required: true,
					},
				},
			},
		},
		"security_logging_object_refs": &schema.Schema{
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
						Elem:     ResourceSloRateType(),
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

func ResourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFirewallPolicyCreate,
		Read:   ResourceFirewallPolicyRead,
		Update: ResourceFirewallPolicyUpdate,
		Delete: ResourceFirewallPolicyDelete,
		Schema: ResourceFirewallPolicySchema(),
	}
}

func ResourceFirewallPolicyRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFirewallPolicyRefsCreate,
		Read:   ResourceFirewallPolicyRefsRead,
		Update: ResourceFirewallPolicyRefsUpdate,
		Delete: ResourceFirewallPolicyRefsDelete,
		Schema: ResourceFirewallPolicyRefsSchema(),
	}
}

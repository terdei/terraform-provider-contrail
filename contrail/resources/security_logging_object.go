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

func SetSecurityLoggingObjectFromResource(object *SecurityLoggingObject, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetSecurityLoggingObjectFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("security_logging_object_rules"); ok {
		member := new(SecurityLoggingObjectRuleListType)
		SetSecurityLoggingObjectRuleListTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetSecurityLoggingObjectRules(member)
	}
	if val, ok := d.GetOk("security_logging_object_rate"); ok {
		object.SetSecurityLoggingObjectRate(val.(int))
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

func SetRefsSecurityLoggingObjectFromResource(object *SecurityLoggingObject, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsSecurityLoggingObjectFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("network_policy_refs"); ok {
		log.Printf("Got ref network_policy_refs -- will call: object.AddNetworkPolicy(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("network-policy", refId.(string))
			dataObj := new(SecurityLoggingObjectRuleListType)
			SetSecurityLoggingObjectRuleListTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving network-policy by Uuid = %v as ref for NetworkPolicy on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddNetworkPolicy(refObj.(*NetworkPolicy), *dataObj)
		}
	}
	if val, ok := d.GetOk("security_group_refs"); ok {
		log.Printf("Got ref security_group_refs -- will call: object.AddSecurityGroup(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("security-group", refId.(string))
			dataObj := new(SecurityLoggingObjectRuleListType)
			SetSecurityLoggingObjectRuleListTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving security-group by Uuid = %v as ref for SecurityGroup on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddSecurityGroup(refObj.(*SecurityGroup), *dataObj)
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

func WriteSecurityLoggingObjectToResource(object SecurityLoggingObject, d *schema.ResourceData, m interface{}) {

	security_logging_object_rulesObj := object.GetSecurityLoggingObjectRules()
	d.Set("security_logging_object_rules", TakeSecurityLoggingObjectRuleListTypeAsMap(&security_logging_object_rulesObj))
	d.Set("security_logging_object_rate", object.GetSecurityLoggingObjectRate())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeSecurityLoggingObjectAsMap(object *SecurityLoggingObject) map[string]interface{} {
	omap := make(map[string]interface{})

	security_logging_object_rulesObj := object.GetSecurityLoggingObjectRules()
	omap["security_logging_object_rules"] = TakeSecurityLoggingObjectRuleListTypeAsMap(&security_logging_object_rulesObj)
	omap["security_logging_object_rate"] = object.GetSecurityLoggingObjectRate()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateSecurityLoggingObjectFromResource(object *SecurityLoggingObject, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("security_logging_object_rules") {
		if val, ok := d.GetOk("security_logging_object_rules"); ok {
			member := new(SecurityLoggingObjectRuleListType)
			SetSecurityLoggingObjectRuleListTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetSecurityLoggingObjectRules(member)
		}
	}
	if d.HasChange("security_logging_object_rate") {
		if val, ok := d.GetOk("security_logging_object_rate"); ok {
			object.SetSecurityLoggingObjectRate(val.(int))
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

func ResourceSecurityLoggingObjectCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceSecurityLoggingObjectCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(SecurityLoggingObject)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceSecurityLoggingObjectCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "SecurityLoggingObject", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetSecurityLoggingObjectFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectCreate] Creation of resource SecurityLoggingObject on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceSecurityLoggingObjectRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceSecurityLoggingObjectRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceSecurityLoggingObjectRefsCreate] Missing 'uuid' field for resource SecurityLoggingObject")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("security-logging-object", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectRefsCreate] Retrieving SecurityLoggingObject with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objSecurityLoggingObject := obj.(*SecurityLoggingObject) // Fully set by Contrail backend
	if err := SetRefsSecurityLoggingObjectFromResource(objSecurityLoggingObject, d, m); err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectRefsCreate] Set refs on object SecurityLoggingObject (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objSecurityLoggingObject.GetHref())
	if err := client.Update(objSecurityLoggingObject); err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectRefsCreate] Update refs for resource SecurityLoggingObject (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objSecurityLoggingObject.GetUuid())
	return nil
}

func ResourceSecurityLoggingObjectRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityLoggingObjectREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("security-logging-object", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectRead] Read resource security-logging-object on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*SecurityLoggingObject)
	WriteSecurityLoggingObjectToResource(*object, d, m)
	return nil
}

func ResourceSecurityLoggingObjectRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityLoggingObjectRefsREAD")
	return nil
}

func ResourceSecurityLoggingObjectUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityLoggingObjectUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("security-logging-object", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectResourceUpdate] Retrieving SecurityLoggingObject with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*SecurityLoggingObject)
	UpdateSecurityLoggingObjectFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectUpdate] Update of resource SecurityLoggingObject on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceSecurityLoggingObjectRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityLoggingObjectRefsUpdate")
	return nil
}

func ResourceSecurityLoggingObjectDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityLoggingObjectDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("security-logging-object", d.Id()); err != nil {
		return fmt.Errorf("[ResourceSecurityLoggingObjectDelete] Deletion of resource security-logging-object on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceSecurityLoggingObjectRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceSecurityLoggingObjectRefsDelete: %v", d.Id())
	return nil
}

func ResourceSecurityLoggingObjectSchema() map[string]*schema.Schema {
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
		"security_logging_object_rules": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSecurityLoggingObjectRuleListType(),
		},
		"security_logging_object_rate": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
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

func ResourceSecurityLoggingObjectRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"network_policy_refs": &schema.Schema{
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
						Elem:     ResourceSecurityLoggingObjectRuleListType(),
						Required: true,
					},
				},
			},
		},
		"security_group_refs": &schema.Schema{
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
						Elem:     ResourceSecurityLoggingObjectRuleListType(),
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

func ResourceSecurityLoggingObject() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSecurityLoggingObjectCreate,
		Read:   ResourceSecurityLoggingObjectRead,
		Update: ResourceSecurityLoggingObjectUpdate,
		Delete: ResourceSecurityLoggingObjectDelete,
		Schema: ResourceSecurityLoggingObjectSchema(),
	}
}

func ResourceSecurityLoggingObjectRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceSecurityLoggingObjectRefsCreate,
		Read:   ResourceSecurityLoggingObjectRefsRead,
		Update: ResourceSecurityLoggingObjectRefsUpdate,
		Delete: ResourceSecurityLoggingObjectRefsDelete,
		Schema: ResourceSecurityLoggingObjectRefsSchema(),
	}
}

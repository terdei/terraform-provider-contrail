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

func SetFirewallRuleFromResource(object *FirewallRule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetFirewallRuleFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("action_list"); ok {
		member := new(ActionListType)
		SetActionListTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetActionList(member)
	}
	if val, ok := d.GetOk("service"); ok {
		member := new(FirewallServiceType)
		SetFirewallServiceTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetService(member)
	}
	if val, ok := d.GetOk("endpoint_1"); ok {
		member := new(FirewallRuleEndpointType)
		SetFirewallRuleEndpointTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetEndpoint1(member)
	}
	if val, ok := d.GetOk("endpoint_2"); ok {
		member := new(FirewallRuleEndpointType)
		SetFirewallRuleEndpointTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetEndpoint2(member)
	}
	if val, ok := d.GetOk("match_tags"); ok {
		member := new(FirewallRuleMatchTagsType)
		SetFirewallRuleMatchTagsTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMatchTags(member)
	}
	if val, ok := d.GetOk("match_tag_types"); ok {
		member := new(FirewallRuleMatchTagsTypeIdList)
		SetFirewallRuleMatchTagsTypeIdListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMatchTagTypes(member)
	}
	if val, ok := d.GetOk("direction"); ok {
		object.SetDirection(val.(string))
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

func SetRefsFirewallRuleFromResource(object *FirewallRule, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsFirewallRuleFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_group_refs"); ok {
		log.Printf("Got ref service_group_refs -- will call: object.AddServiceGroup(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-group", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-group by Uuid = %v as ref for ServiceGroup on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceGroup(refObj.(*ServiceGroup))
		}
	}
	if val, ok := d.GetOk("address_group_refs"); ok {
		log.Printf("Got ref address_group_refs -- will call: object.AddAddressGroup(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("address-group", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving address-group by Uuid = %v as ref for AddressGroup on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddAddressGroup(refObj.(*AddressGroup))
		}
	}
	if val, ok := d.GetOk("virtual_network_refs"); ok {
		log.Printf("Got ref virtual_network_refs -- will call: object.AddVirtualNetwork(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-network", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-network by Uuid = %v as ref for VirtualNetwork on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualNetwork(refObj.(*VirtualNetwork))
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

func WriteFirewallRuleToResource(object FirewallRule, d *schema.ResourceData, m interface{}) {

	action_listObj := object.GetActionList()
	d.Set("action_list", TakeActionListTypeAsMap(&action_listObj))
	serviceObj := object.GetService()
	d.Set("service", TakeFirewallServiceTypeAsMap(&serviceObj))
	endpoint_1Obj := object.GetEndpoint1()
	d.Set("endpoint_1", TakeFirewallRuleEndpointTypeAsMap(&endpoint_1Obj))
	endpoint_2Obj := object.GetEndpoint2()
	d.Set("endpoint_2", TakeFirewallRuleEndpointTypeAsMap(&endpoint_2Obj))
	match_tagsObj := object.GetMatchTags()
	d.Set("match_tags", TakeFirewallRuleMatchTagsTypeAsMap(&match_tagsObj))
	match_tag_typesObj := object.GetMatchTagTypes()
	d.Set("match_tag_types", TakeFirewallRuleMatchTagsTypeIdListAsMap(&match_tag_typesObj))
	d.Set("direction", object.GetDirection())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeFirewallRuleAsMap(object *FirewallRule) map[string]interface{} {
	omap := make(map[string]interface{})

	action_listObj := object.GetActionList()
	omap["action_list"] = TakeActionListTypeAsMap(&action_listObj)
	serviceObj := object.GetService()
	omap["service"] = TakeFirewallServiceTypeAsMap(&serviceObj)
	endpoint_1Obj := object.GetEndpoint1()
	omap["endpoint_1"] = TakeFirewallRuleEndpointTypeAsMap(&endpoint_1Obj)
	endpoint_2Obj := object.GetEndpoint2()
	omap["endpoint_2"] = TakeFirewallRuleEndpointTypeAsMap(&endpoint_2Obj)
	match_tagsObj := object.GetMatchTags()
	omap["match_tags"] = TakeFirewallRuleMatchTagsTypeAsMap(&match_tagsObj)
	match_tag_typesObj := object.GetMatchTagTypes()
	omap["match_tag_types"] = TakeFirewallRuleMatchTagsTypeIdListAsMap(&match_tag_typesObj)
	omap["direction"] = object.GetDirection()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateFirewallRuleFromResource(object *FirewallRule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("action_list") {
		if val, ok := d.GetOk("action_list"); ok {
			member := new(ActionListType)
			SetActionListTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetActionList(member)
		}
	}
	if d.HasChange("service") {
		if val, ok := d.GetOk("service"); ok {
			member := new(FirewallServiceType)
			SetFirewallServiceTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetService(member)
		}
	}
	if d.HasChange("endpoint_1") {
		if val, ok := d.GetOk("endpoint_1"); ok {
			member := new(FirewallRuleEndpointType)
			SetFirewallRuleEndpointTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetEndpoint1(member)
		}
	}
	if d.HasChange("endpoint_2") {
		if val, ok := d.GetOk("endpoint_2"); ok {
			member := new(FirewallRuleEndpointType)
			SetFirewallRuleEndpointTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetEndpoint2(member)
		}
	}
	if d.HasChange("match_tags") {
		if val, ok := d.GetOk("match_tags"); ok {
			member := new(FirewallRuleMatchTagsType)
			SetFirewallRuleMatchTagsTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMatchTags(member)
		}
	}
	if d.HasChange("match_tag_types") {
		if val, ok := d.GetOk("match_tag_types"); ok {
			member := new(FirewallRuleMatchTagsTypeIdList)
			SetFirewallRuleMatchTagsTypeIdListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMatchTagTypes(member)
		}
	}
	if d.HasChange("direction") {
		if val, ok := d.GetOk("direction"); ok {
			object.SetDirection(val.(string))
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

func ResourceFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFirewallRuleCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(FirewallRule)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceFirewallRuleCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "FirewallRule", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetFirewallRuleFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceFirewallRuleCreate] Creation of resource FirewallRule on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceFirewallRuleRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceFirewallRuleRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceFirewallRuleRefsCreate] Missing 'uuid' field for resource FirewallRule")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("firewall-rule", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceFirewallRuleRefsCreate] Retrieving FirewallRule with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objFirewallRule := obj.(*FirewallRule) // Fully set by Contrail backend
	if err := SetRefsFirewallRuleFromResource(objFirewallRule, d, m); err != nil {
		return fmt.Errorf("[ResourceFirewallRuleRefsCreate] Set refs on object FirewallRule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objFirewallRule.GetHref())
	if err := client.Update(objFirewallRule); err != nil {
		return fmt.Errorf("[ResourceFirewallRuleRefsCreate] Update refs for resource FirewallRule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objFirewallRule.GetUuid())
	return nil
}

func ResourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallRuleREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("firewall-rule", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFirewallRuleRead] Read resource firewall-rule on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*FirewallRule)
	WriteFirewallRuleToResource(*object, d, m)
	return nil
}

func ResourceFirewallRuleRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallRuleRefsREAD")
	return nil
}

func ResourceFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallRuleUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("firewall-rule", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceFirewallRuleResourceUpdate] Retrieving FirewallRule with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*FirewallRule)
	UpdateFirewallRuleFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceFirewallRuleUpdate] Update of resource FirewallRule on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceFirewallRuleRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallRuleRefsUpdate")
	return nil
}

func ResourceFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallRuleDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("firewall-rule", d.Id()); err != nil {
		return fmt.Errorf("[ResourceFirewallRuleDelete] Deletion of resource firewall-rule on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceFirewallRuleRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceFirewallRuleRefsDelete: %v", d.Id())
	return nil
}

func ResourceFirewallRuleSchema() map[string]*schema.Schema {
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
		"action_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceActionListType(),
		},
		"service": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallServiceType(),
		},
		"endpoint_1": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallRuleEndpointType(),
		},
		"endpoint_2": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallRuleEndpointType(),
		},
		"match_tags": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallRuleMatchTagsType(),
		},
		"match_tag_types": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceFirewallRuleMatchTagsTypeIdList(),
		},
		"direction": &schema.Schema{
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

func ResourceFirewallRuleRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_group_refs": &schema.Schema{
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
		"address_group_refs": &schema.Schema{
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
		"virtual_network_refs": &schema.Schema{
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

func ResourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFirewallRuleCreate,
		Read:   ResourceFirewallRuleRead,
		Update: ResourceFirewallRuleUpdate,
		Delete: ResourceFirewallRuleDelete,
		Schema: ResourceFirewallRuleSchema(),
	}
}

func ResourceFirewallRuleRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceFirewallRuleRefsCreate,
		Read:   ResourceFirewallRuleRefsRead,
		Update: ResourceFirewallRuleRefsUpdate,
		Delete: ResourceFirewallRuleRefsDelete,
		Schema: ResourceFirewallRuleRefsSchema(),
	}
}

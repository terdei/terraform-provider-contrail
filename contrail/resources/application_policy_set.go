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

func SetApplicationPolicySetFromResource(object *ApplicationPolicySet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetApplicationPolicySetFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("all_applications"); ok {
		object.SetAllApplications(val.(bool))
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

func SetRefsApplicationPolicySetFromResource(object *ApplicationPolicySet, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsApplicationPolicySetFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("firewall_policy_refs"); ok {
		log.Printf("Got ref firewall_policy_refs -- will call: object.AddFirewallPolicy(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("firewall-policy", refId.(string))
			dataObj := new(FirewallSequence)
			SetFirewallSequenceFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving firewall-policy by Uuid = %v as ref for FirewallPolicy on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddFirewallPolicy(refObj.(*FirewallPolicy), *dataObj)
		}
	}
	if val, ok := d.GetOk("global_vrouter_config_refs"); ok {
		log.Printf("Got ref global_vrouter_config_refs -- will call: object.AddGlobalVrouterConfig(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("global-vrouter-config", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving global-vrouter-config by Uuid = %v as ref for GlobalVrouterConfig on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddGlobalVrouterConfig(refObj.(*GlobalVrouterConfig))
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

func DeleteRefsApplicationPolicySetFromResource(object *ApplicationPolicySet, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsApplicationPolicySetFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("firewall_policy_refs"); ok {
		log.Printf("Got ref firewall_policy_refs -- will call: object.DeleteFirewallPolicy(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteFirewallPolicy(refId.(string))
		}
	}
	if val, ok := d.GetOk("global_vrouter_config_refs"); ok {
		log.Printf("Got ref global_vrouter_config_refs -- will call: object.DeleteGlobalVrouterConfig(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteGlobalVrouterConfig(refId.(string))
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

func WriteApplicationPolicySetToResource(object ApplicationPolicySet, d *schema.ResourceData, m interface{}) {

	d.Set("all_applications", object.GetAllApplications())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteApplicationPolicySetRefsToResource(object ApplicationPolicySet, d *schema.ResourceData, m interface{}) {

	if ref, err := object.GetFirewallPolicyRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("firewall_policy_refs", refList)
	}
	if ref, err := object.GetGlobalVrouterConfigRefs(); err != nil {
		var refList []interface{}
		for _, v := range ref {
			omap := make(map[string]interface{})
			omap["to"] = v.Uuid
			refList = append(refList, omap)
		}
		d.Set("global_vrouter_config_refs", refList)
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

func TakeApplicationPolicySetAsMap(object *ApplicationPolicySet) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["all_applications"] = object.GetAllApplications()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateApplicationPolicySetFromResource(object *ApplicationPolicySet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("all_applications") {
		if val, ok := d.GetOk("all_applications"); ok {
			object.SetAllApplications(val.(bool))
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

func UpdateApplicationPolicySetRefsFromResource(object *ApplicationPolicySet, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if d.HasChange("firewall_policy_refs") {
		object.ClearFirewallPolicy()
		if val, ok := d.GetOk("firewall_policy_refs"); ok {
			log.Printf("Got ref firewall_policy_refs -- will call: object.AddFirewallPolicy(refObj, *dataObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("firewall-policy", refId.(string))
				dataObj := new(FirewallSequence)
				SetFirewallSequenceFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
				log.Printf("Data obj: %+v", dataObj)
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddFirewallPolicy(refObj.(*FirewallPolicy), *dataObj)
			}
		}
	}
	if d.HasChange("global_vrouter_config_refs") {
		object.ClearGlobalVrouterConfig()
		if val, ok := d.GetOk("global_vrouter_config_refs"); ok {
			log.Printf("Got ref global_vrouter_config_refs -- will call: object.AddGlobalVrouterConfig(refObj)")
			for k, v := range val.([]interface{}) {
				log.Printf("Item: %+v => <%T> %+v", k, v, v)
				refId := (v.(map[string]interface{}))["to"]
				log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
				refObj, _ := client.FindByUuid("global-vrouter-config", refId.(string))
				log.Printf("Ref 'to' (OBJECT): %+v", refObj)
				object.AddGlobalVrouterConfig(refObj.(*GlobalVrouterConfig))
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

func ResourceApplicationPolicySetCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceApplicationPolicySetCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ApplicationPolicySet)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceApplicationPolicySetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "ApplicationPolicySet", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceApplicationPolicySetCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ApplicationPolicySet", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetApplicationPolicySetFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetCreate] Creation of resource ApplicationPolicySet on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceApplicationPolicySetRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceApplicationPolicySetRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsCreate] Missing 'uuid' field for resource ApplicationPolicySet")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("application-policy-set", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsCreate] Retrieving ApplicationPolicySet with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objApplicationPolicySet := obj.(*ApplicationPolicySet) // Fully set by Contrail backend
	if err := SetRefsApplicationPolicySetFromResource(objApplicationPolicySet, d, m); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsCreate] Set refs on object ApplicationPolicySet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objApplicationPolicySet.GetHref())
	if err := client.Update(objApplicationPolicySet); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsCreate] Update refs for resource ApplicationPolicySet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objApplicationPolicySet.GetUuid())
	return nil
}

func ResourceApplicationPolicySetRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApplicationPolicySetRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("application-policy-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRead] Read resource application-policy-set on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ApplicationPolicySet)
	WriteApplicationPolicySetToResource(*object, d, m)
	return nil
}

func ResourceApplicationPolicySetRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApplicationPolicySetRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("application-policy-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsRead] Read resource application-policy-set on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ApplicationPolicySet)
	WriteApplicationPolicySetRefsToResource(*object, d, m)
	return nil
}

func ResourceApplicationPolicySetUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApplicationPolicySetUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("application-policy-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetUpdate] Retrieving ApplicationPolicySet with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ApplicationPolicySet)
	UpdateApplicationPolicySetFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetUpdate] Update of resource ApplicationPolicySet on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceApplicationPolicySetRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApplicationPolicySetRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("application-policy-set", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsUpdate] Retrieving ApplicationPolicySet with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ApplicationPolicySet)
	UpdateApplicationPolicySetRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsUpdate] Update of resource ApplicationPolicySet on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceApplicationPolicySetDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceApplicationPolicySetDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("application-policy-set", d.Id()); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetDelete] Deletion of resource application-policy-set on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceApplicationPolicySetRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceApplicationPolicySetRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsDelete] Missing 'uuid' field for resource ApplicationPolicySet")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("application-policy-set", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsDelete] Retrieving ApplicationPolicySet with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objApplicationPolicySet := obj.(*ApplicationPolicySet) // Fully set by Contrail backend
	if err := DeleteRefsApplicationPolicySetFromResource(objApplicationPolicySet, d, m); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsDelete] Set refs on object ApplicationPolicySet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objApplicationPolicySet.GetHref())
	if err := client.Update(objApplicationPolicySet); err != nil {
		return fmt.Errorf("[ResourceApplicationPolicySetRefsDelete] Delete refs for resource ApplicationPolicySet (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objApplicationPolicySet.GetUuid())
	return nil
}

func ResourceApplicationPolicySetSchema() map[string]*schema.Schema {
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
		"all_applications": &schema.Schema{
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

func ResourceApplicationPolicySetRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"firewall_policy_refs": &schema.Schema{
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
		"global_vrouter_config_refs": &schema.Schema{
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

func ResourceApplicationPolicySet() *schema.Resource {
	return &schema.Resource{
		Create: ResourceApplicationPolicySetCreate,
		Read:   ResourceApplicationPolicySetRead,
		Update: ResourceApplicationPolicySetUpdate,
		Delete: ResourceApplicationPolicySetDelete,
		Schema: ResourceApplicationPolicySetSchema(),
	}
}

func ResourceApplicationPolicySetRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceApplicationPolicySetRefsCreate,
		Read:   ResourceApplicationPolicySetRefsRead,
		Update: ResourceApplicationPolicySetRefsUpdate,
		Delete: ResourceApplicationPolicySetRefsDelete,
		Schema: ResourceApplicationPolicySetRefsSchema(),
	}
}

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

func SetDsaRuleFromResource(object *DsaRule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetDsaRuleFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("dsa_rule_entry"); ok {
		member := new(DiscoveryServiceAssignmentType)
		SetDiscoveryServiceAssignmentTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetDsaRuleEntry(member)
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

func SetRefsDsaRuleFromResource(object *DsaRule, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsDsaRuleFromResource] key = %v, prefix = %v", key, prefix)

	return nil
}

func WriteDsaRuleToResource(object DsaRule, d *schema.ResourceData, m interface{}) {

	dsa_rule_entryObj := object.GetDsaRuleEntry()
	d.Set("dsa_rule_entry", TakeDiscoveryServiceAssignmentTypeAsMap(&dsa_rule_entryObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeDsaRuleAsMap(object *DsaRule) map[string]interface{} {
	omap := make(map[string]interface{})

	dsa_rule_entryObj := object.GetDsaRuleEntry()
	omap["dsa_rule_entry"] = TakeDiscoveryServiceAssignmentTypeAsMap(&dsa_rule_entryObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateDsaRuleFromResource(object *DsaRule, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("dsa_rule_entry") {
		if val, ok := d.GetOk("dsa_rule_entry"); ok {
			member := new(DiscoveryServiceAssignmentType)
			SetDiscoveryServiceAssignmentTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetDsaRuleEntry(member)
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

func ResourceDsaRuleCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDsaRuleCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(DsaRule)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceDsaRuleCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "DsaRule", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetDsaRuleFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceDsaRuleCreate] Creation of resource DsaRule on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceDsaRuleRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDsaRuleRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDsaRuleRefsCreate] Missing 'uuid' field for resource DsaRule")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("dsa-rule", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDsaRuleRefsCreate] Retrieving DsaRule with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDsaRule := obj.(*DsaRule) // Fully set by Contrail backend
	if err := SetRefsDsaRuleFromResource(objDsaRule, d, m); err != nil {
		return fmt.Errorf("[ResourceDsaRuleRefsCreate] Set refs on object DsaRule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDsaRule.GetHref())
	if err := client.Update(objDsaRule); err != nil {
		return fmt.Errorf("[ResourceDsaRuleRefsCreate] Update refs for resource DsaRule (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDsaRule.GetUuid())
	return nil
}

func ResourceDsaRuleRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDsaRuleREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("dsa-rule", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDsaRuleRead] Read resource dsa-rule on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*DsaRule)
	WriteDsaRuleToResource(*object, d, m)
	return nil
}

func ResourceDsaRuleRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDsaRuleRefsREAD")
	return nil
}

func ResourceDsaRuleUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDsaRuleUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("dsa-rule", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDsaRuleResourceUpdate] Retrieving DsaRule with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*DsaRule)
	UpdateDsaRuleFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDsaRuleUpdate] Update of resource DsaRule on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDsaRuleRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDsaRuleRefsUpdate")
	return nil
}

func ResourceDsaRuleDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDsaRuleDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("dsa-rule", d.Id()); err != nil {
		return fmt.Errorf("[ResourceDsaRuleDelete] Deletion of resource dsa-rule on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceDsaRuleRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDsaRuleRefsDelete: %v", d.Id())
	return nil
}

func ResourceDsaRuleSchema() map[string]*schema.Schema {
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
		"dsa_rule_entry": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDiscoveryServiceAssignmentType(),
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

func ResourceDsaRule() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDsaRuleCreate,
		Read:   ResourceDsaRuleRead,
		Update: ResourceDsaRuleUpdate,
		Delete: ResourceDsaRuleDelete,
		Schema: ResourceDsaRuleSchema(),
	}
}

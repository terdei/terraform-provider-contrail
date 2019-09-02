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

func SetBridgeDomainFromResource(object *BridgeDomain, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetBridgeDomainFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("mac_learning_enabled"); ok {
		object.SetMacLearningEnabled(val.(bool))
	}
	if val, ok := d.GetOk("mac_limit_control"); ok {
		member := new(MACLimitControlType)
		SetMACLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMacLimitControl(member)
	}
	if val, ok := d.GetOk("mac_move_control"); ok {
		member := new(MACMoveLimitControlType)
		SetMACMoveLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetMacMoveControl(member)
	}
	if val, ok := d.GetOk("mac_aging_time"); ok {
		object.SetMacAgingTime(val.(int))
	}
	if val, ok := d.GetOk("isid"); ok {
		object.SetIsid(val.(int))
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

func SetRefsBridgeDomainFromResource(object *BridgeDomain, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsBridgeDomainFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsBridgeDomainFromResource(object *BridgeDomain, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsBridgeDomainFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteBridgeDomainToResource(object BridgeDomain, d *schema.ResourceData, m interface{}) {

	d.Set("mac_learning_enabled", object.GetMacLearningEnabled())
	mac_limit_controlObj := object.GetMacLimitControl()
	d.Set("mac_limit_control", TakeMACLimitControlTypeAsMap(&mac_limit_controlObj))
	mac_move_controlObj := object.GetMacMoveControl()
	d.Set("mac_move_control", TakeMACMoveLimitControlTypeAsMap(&mac_move_controlObj))
	d.Set("mac_aging_time", object.GetMacAgingTime())
	d.Set("isid", object.GetIsid())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteBridgeDomainRefsToResource(object BridgeDomain, d *schema.ResourceData, m interface{}) {

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

func TakeBridgeDomainAsMap(object *BridgeDomain) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["mac_learning_enabled"] = object.GetMacLearningEnabled()
	mac_limit_controlObj := object.GetMacLimitControl()
	omap["mac_limit_control"] = TakeMACLimitControlTypeAsMap(&mac_limit_controlObj)
	mac_move_controlObj := object.GetMacMoveControl()
	omap["mac_move_control"] = TakeMACMoveLimitControlTypeAsMap(&mac_move_controlObj)
	omap["mac_aging_time"] = object.GetMacAgingTime()
	omap["isid"] = object.GetIsid()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateBridgeDomainFromResource(object *BridgeDomain, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("mac_learning_enabled") {
		if val, ok := d.GetOk("mac_learning_enabled"); ok {
			object.SetMacLearningEnabled(val.(bool))
		}
	}
	if d.HasChange("mac_limit_control") {
		if val, ok := d.GetOk("mac_limit_control"); ok {
			member := new(MACLimitControlType)
			SetMACLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMacLimitControl(member)
		}
	}
	if d.HasChange("mac_move_control") {
		if val, ok := d.GetOk("mac_move_control"); ok {
			member := new(MACMoveLimitControlType)
			SetMACMoveLimitControlTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetMacMoveControl(member)
		}
	}
	if d.HasChange("mac_aging_time") {
		if val, ok := d.GetOk("mac_aging_time"); ok {
			object.SetMacAgingTime(val.(int))
		}
	}
	if d.HasChange("isid") {
		if val, ok := d.GetOk("isid"); ok {
			object.SetIsid(val.(int))
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

func UpdateBridgeDomainRefsFromResource(object *BridgeDomain, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceBridgeDomainCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBridgeDomainCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(BridgeDomain)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceBridgeDomainCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "BridgeDomain", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetBridgeDomainFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainCreate] Creation of resource BridgeDomain on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceBridgeDomainRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBridgeDomainRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBridgeDomainRefsCreate] Missing 'uuid' field for resource BridgeDomain")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bridge-domain", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsCreate] Retrieving BridgeDomain with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBridgeDomain := obj.(*BridgeDomain) // Fully set by Contrail backend
	if err := SetRefsBridgeDomainFromResource(objBridgeDomain, d, m); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsCreate] Set refs on object BridgeDomain (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBridgeDomain.GetHref())
	if err := client.Update(objBridgeDomain); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsCreate] Update refs for resource BridgeDomain (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBridgeDomain.GetUuid())
	return nil
}

func ResourceBridgeDomainRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBridgeDomainRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bridge-domain", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRead] Read resource bridge-domain on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BridgeDomain)
	WriteBridgeDomainToResource(*object, d, m)
	return nil
}

func ResourceBridgeDomainRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBridgeDomainRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bridge-domain", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsRead] Read resource bridge-domain on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*BridgeDomain)
	WriteBridgeDomainRefsToResource(*object, d, m)
	return nil
}

func ResourceBridgeDomainUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBridgeDomainUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bridge-domain", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBridgeDomainUpdate] Retrieving BridgeDomain with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BridgeDomain)
	UpdateBridgeDomainFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainUpdate] Update of resource BridgeDomain on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBridgeDomainRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBridgeDomainRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bridge-domain", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsUpdate] Retrieving BridgeDomain with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*BridgeDomain)
	UpdateBridgeDomainRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsUpdate] Update of resource BridgeDomain on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBridgeDomainDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBridgeDomainDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("bridge-domain", d.Id()); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainDelete] Deletion of resource bridge-domain on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceBridgeDomainRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBridgeDomainRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBridgeDomainRefsDelete] Missing 'uuid' field for resource BridgeDomain")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bridge-domain", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsDelete] Retrieving BridgeDomain with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBridgeDomain := obj.(*BridgeDomain) // Fully set by Contrail backend
	if err := DeleteRefsBridgeDomainFromResource(objBridgeDomain, d, m); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsDelete] Set refs on object BridgeDomain (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBridgeDomain.GetHref())
	if err := client.Update(objBridgeDomain); err != nil {
		return fmt.Errorf("[ResourceBridgeDomainRefsDelete] Delete refs for resource BridgeDomain (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBridgeDomain.GetUuid())
	return nil
}

func ResourceBridgeDomainSchema() map[string]*schema.Schema {
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
		"mac_learning_enabled": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"mac_limit_control": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMACLimitControlType(),
		},
		"mac_move_control": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceMACMoveLimitControlType(),
		},
		"mac_aging_time": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"isid": &schema.Schema{
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

func ResourceBridgeDomainRefsSchema() map[string]*schema.Schema {
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

func ResourceBridgeDomain() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBridgeDomainCreate,
		Read:   ResourceBridgeDomainRead,
		Update: ResourceBridgeDomainUpdate,
		Delete: ResourceBridgeDomainDelete,
		Schema: ResourceBridgeDomainSchema(),
	}
}

func ResourceBridgeDomainRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBridgeDomainRefsCreate,
		Read:   ResourceBridgeDomainRefsRead,
		Update: ResourceBridgeDomainRefsUpdate,
		Delete: ResourceBridgeDomainRefsDelete,
		Schema: ResourceBridgeDomainRefsSchema(),
	}
}

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

func SetAlarmFromResource(object *Alarm, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetAlarmFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("uve_keys"); ok {
		member := new(UveKeysType)
		SetUveKeysTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetUveKeys(member)
	}
	if val, ok := d.GetOk("alarm_severity"); ok {
		object.SetAlarmSeverity(val.(int))
	}
	if val, ok := d.GetOk("alarm_rules"); ok {
		member := new(AlarmOrList)
		SetAlarmOrListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetAlarmRules(member)
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

func SetRefsAlarmFromResource(object *Alarm, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsAlarmFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsAlarmFromResource(object *Alarm, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsAlarmFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteAlarmToResource(object Alarm, d *schema.ResourceData, m interface{}) {

	uve_keysObj := object.GetUveKeys()
	d.Set("uve_keys", TakeUveKeysTypeAsMap(&uve_keysObj))
	d.Set("alarm_severity", object.GetAlarmSeverity())
	alarm_rulesObj := object.GetAlarmRules()
	d.Set("alarm_rules", TakeAlarmOrListAsMap(&alarm_rulesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeAlarmAsMap(object *Alarm) map[string]interface{} {
	omap := make(map[string]interface{})

	uve_keysObj := object.GetUveKeys()
	omap["uve_keys"] = TakeUveKeysTypeAsMap(&uve_keysObj)
	omap["alarm_severity"] = object.GetAlarmSeverity()
	alarm_rulesObj := object.GetAlarmRules()
	omap["alarm_rules"] = TakeAlarmOrListAsMap(&alarm_rulesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateAlarmFromResource(object *Alarm, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("uve_keys") {
		if val, ok := d.GetOk("uve_keys"); ok {
			member := new(UveKeysType)
			SetUveKeysTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetUveKeys(member)
		}
	}
	if d.HasChange("alarm_severity") {
		if val, ok := d.GetOk("alarm_severity"); ok {
			object.SetAlarmSeverity(val.(int))
		}
	}
	if d.HasChange("alarm_rules") {
		if val, ok := d.GetOk("alarm_rules"); ok {
			member := new(AlarmOrList)
			SetAlarmOrListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetAlarmRules(member)
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

func ResourceAlarmCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAlarmCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Alarm)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceAlarmCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Alarm", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetAlarmFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceAlarmCreate] Creation of resource Alarm on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceAlarmRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAlarmRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAlarmRefsCreate] Missing 'uuid' field for resource Alarm")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("alarm", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAlarmRefsCreate] Retrieving Alarm with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAlarm := obj.(*Alarm) // Fully set by Contrail backend
	if err := SetRefsAlarmFromResource(objAlarm, d, m); err != nil {
		return fmt.Errorf("[ResourceAlarmRefsCreate] Set refs on object Alarm (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAlarm.GetHref())
	if err := client.Update(objAlarm); err != nil {
		return fmt.Errorf("[ResourceAlarmRefsCreate] Update refs for resource Alarm (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAlarm.GetUuid())
	return nil
}

func ResourceAlarmRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAlarmREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("alarm", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAlarmRead] Read resource alarm on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Alarm)
	WriteAlarmToResource(*object, d, m)
	return nil
}

func ResourceAlarmRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAlarmRefsREAD")
	return nil
}

func ResourceAlarmUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAlarmUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("alarm", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAlarmResourceUpdate] Retrieving Alarm with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Alarm)
	UpdateAlarmFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAlarmUpdate] Update of resource Alarm on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAlarmRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAlarmRefsUpdate")
	return nil
}

func ResourceAlarmDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAlarmDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("alarm", d.Id()); err != nil {
		return fmt.Errorf("[ResourceAlarmDelete] Deletion of resource alarm on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceAlarmRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAlarmRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAlarmRefsDelete] Missing 'uuid' field for resource Alarm")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("alarm", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAlarmRefsDelete] Retrieving Alarm with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAlarm := obj.(*Alarm) // Fully set by Contrail backend
	if err := DeleteRefsAlarmFromResource(objAlarm, d, m); err != nil {
		return fmt.Errorf("[ResourceAlarmRefsDelete] Set refs on object Alarm (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAlarm.GetHref())
	if err := client.Update(objAlarm); err != nil {
		return fmt.Errorf("[ResourceAlarmRefsDelete] Delete refs for resource Alarm (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAlarm.GetUuid())
	return nil
}

func ResourceAlarmSchema() map[string]*schema.Schema {
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
		"uve_keys": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceUveKeysType(),
		},
		"alarm_severity": &schema.Schema{
			Optional: true,
			Type:     schema.TypeInt,
		},
		"alarm_rules": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceAlarmOrList(),
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

func ResourceAlarmRefsSchema() map[string]*schema.Schema {
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

func ResourceAlarm() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAlarmCreate,
		Read:   ResourceAlarmRead,
		Update: ResourceAlarmUpdate,
		Delete: ResourceAlarmDelete,
		Schema: ResourceAlarmSchema(),
	}
}

func ResourceAlarmRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAlarmRefsCreate,
		Read:   ResourceAlarmRefsRead,
		Update: ResourceAlarmRefsUpdate,
		Delete: ResourceAlarmRefsDelete,
		Schema: ResourceAlarmRefsSchema(),
	}
}

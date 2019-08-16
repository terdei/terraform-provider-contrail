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

func SetProjectFromResource(object *Project, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetProjectFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("quota"); ok {
		member := new(QuotaType)
		SetQuotaTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetQuota(member)
	}
	if val, ok := d.GetOk("vxlan_routing"); ok {
		object.SetVxlanRouting(val.(bool))
	}
	if val, ok := d.GetOk("alarm_enable"); ok {
		object.SetAlarmEnable(val.(bool))
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

func SetRefsProjectFromResource(object *Project, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsProjectFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("namespace_refs"); ok {
		log.Printf("Got ref namespace_refs -- will call: object.AddNamespace(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("namespace", refId.(string))
			dataObj := new(SubnetType)
			SetSubnetTypeFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving namespace by Uuid = %v as ref for Namespace on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddNamespace(refObj.(*Namespace), *dataObj)
		}
	}
	if val, ok := d.GetOk("floating_ip_pool_refs"); ok {
		log.Printf("Got ref floating_ip_pool_refs -- will call: object.AddFloatingIpPool(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("floating-ip-pool", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving floating-ip-pool by Uuid = %v as ref for FloatingIpPool on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddFloatingIpPool(refObj.(*FloatingIpPool))
		}
	}
	if val, ok := d.GetOk("alias_ip_pool_refs"); ok {
		log.Printf("Got ref alias_ip_pool_refs -- will call: object.AddAliasIpPool(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("alias-ip-pool", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving alias-ip-pool by Uuid = %v as ref for AliasIpPool on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddAliasIpPool(refObj.(*AliasIpPool))
		}
	}
	if val, ok := d.GetOk("application_policy_set_refs"); ok {
		log.Printf("Got ref application_policy_set_refs -- will call: object.AddApplicationPolicySet(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("application-policy-set", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving application-policy-set by Uuid = %v as ref for ApplicationPolicySet on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddApplicationPolicySet(refObj.(*ApplicationPolicySet))
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

func DeleteRefsProjectFromResource(object *Project, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsProjectFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("namespace_refs"); ok {
		log.Printf("Got ref namespace_refs -- will call: object.DeleteNamespace(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteNamespace(refId.(string))
		}
	}
	if val, ok := d.GetOk("floating_ip_pool_refs"); ok {
		log.Printf("Got ref floating_ip_pool_refs -- will call: object.DeleteFloatingIpPool(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteFloatingIpPool(refId.(string))
		}
	}
	if val, ok := d.GetOk("alias_ip_pool_refs"); ok {
		log.Printf("Got ref alias_ip_pool_refs -- will call: object.DeleteAliasIpPool(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteAliasIpPool(refId.(string))
		}
	}
	if val, ok := d.GetOk("application_policy_set_refs"); ok {
		log.Printf("Got ref application_policy_set_refs -- will call: object.DeleteApplicationPolicySet(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteApplicationPolicySet(refId.(string))
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

func WriteProjectToResource(object Project, d *schema.ResourceData, m interface{}) {

	quotaObj := object.GetQuota()
	d.Set("quota", TakeQuotaTypeAsMap(&quotaObj))
	d.Set("vxlan_routing", object.GetVxlanRouting())
	d.Set("alarm_enable", object.GetAlarmEnable())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeProjectAsMap(object *Project) map[string]interface{} {
	omap := make(map[string]interface{})

	quotaObj := object.GetQuota()
	omap["quota"] = TakeQuotaTypeAsMap(&quotaObj)
	omap["vxlan_routing"] = object.GetVxlanRouting()
	omap["alarm_enable"] = object.GetAlarmEnable()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateProjectFromResource(object *Project, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("quota") {
		if val, ok := d.GetOk("quota"); ok {
			member := new(QuotaType)
			SetQuotaTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetQuota(member)
		}
	}
	if d.HasChange("vxlan_routing") {
		if val, ok := d.GetOk("vxlan_routing"); ok {
			object.SetVxlanRouting(val.(bool))
		}
	}
	if d.HasChange("alarm_enable") {
		if val, ok := d.GetOk("alarm_enable"); ok {
			object.SetAlarmEnable(val.(bool))
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

func ResourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceProjectCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Project)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceProjectCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Project", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetProjectFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceProjectCreate] Creation of resource Project on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceProjectRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceProjectRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceProjectRefsCreate] Missing 'uuid' field for resource Project")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("project", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceProjectRefsCreate] Retrieving Project with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objProject := obj.(*Project) // Fully set by Contrail backend
	if err := SetRefsProjectFromResource(objProject, d, m); err != nil {
		return fmt.Errorf("[ResourceProjectRefsCreate] Set refs on object Project (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objProject.GetHref())
	if err := client.Update(objProject); err != nil {
		return fmt.Errorf("[ResourceProjectRefsCreate] Update refs for resource Project (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objProject.GetUuid())
	return nil
}

func ResourceProjectRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProjectREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("project", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceProjectRead] Read resource project on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Project)
	WriteProjectToResource(*object, d, m)
	return nil
}

func ResourceProjectRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProjectRefsREAD")
	return nil
}

func ResourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProjectUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("project", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceProjectResourceUpdate] Retrieving Project with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Project)
	UpdateProjectFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceProjectUpdate] Update of resource Project on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceProjectRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProjectRefsUpdate")
	return nil
}

func ResourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceProjectDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("project", d.Id()); err != nil {
		return fmt.Errorf("[ResourceProjectDelete] Deletion of resource project on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceProjectRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceProjectRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceProjectRefsDelete] Missing 'uuid' field for resource Project")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("project", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceProjectRefsDelete] Retrieving Project with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objProject := obj.(*Project) // Fully set by Contrail backend
	if err := DeleteRefsProjectFromResource(objProject, d, m); err != nil {
		return fmt.Errorf("[ResourceProjectRefsDelete] Set refs on object Project (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objProject.GetHref())
	if err := client.Update(objProject); err != nil {
		return fmt.Errorf("[ResourceProjectRefsDelete] Delete refs for resource Project (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objProject.GetUuid())
	return nil
}

func ResourceProjectSchema() map[string]*schema.Schema {
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
		"quota": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceQuotaType(),
		},
		"vxlan_routing": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"alarm_enable": &schema.Schema{
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

func ResourceProjectRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"namespace_refs": &schema.Schema{
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
						Elem:     ResourceSubnetType(),
						Required: true,
					},
				},
			},
		},
		"floating_ip_pool_refs": &schema.Schema{
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
		"alias_ip_pool_refs": &schema.Schema{
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
		"application_policy_set_refs": &schema.Schema{
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

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		Create: ResourceProjectCreate,
		Read:   ResourceProjectRead,
		Update: ResourceProjectUpdate,
		Delete: ResourceProjectDelete,
		Schema: ResourceProjectSchema(),
	}
}

func ResourceProjectRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceProjectRefsCreate,
		Read:   ResourceProjectRefsRead,
		Update: ResourceProjectRefsUpdate,
		Delete: ResourceProjectRefsDelete,
		Schema: ResourceProjectRefsSchema(),
	}
}

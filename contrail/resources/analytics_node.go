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

func SetAnalyticsNodeFromResource(object *AnalyticsNode, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetAnalyticsNodeFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("analytics_node_ip_address"); ok {
		object.SetAnalyticsNodeIpAddress(val.(string))
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

func SetRefsAnalyticsNodeFromResource(object *AnalyticsNode, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsAnalyticsNodeFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsAnalyticsNodeFromResource(object *AnalyticsNode, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsAnalyticsNodeFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteAnalyticsNodeToResource(object AnalyticsNode, d *schema.ResourceData, m interface{}) {

	d.Set("analytics_node_ip_address", object.GetAnalyticsNodeIpAddress())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteAnalyticsNodeRefsToResource(object AnalyticsNode, d *schema.ResourceData, m interface{}) {

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

func TakeAnalyticsNodeAsMap(object *AnalyticsNode) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["analytics_node_ip_address"] = object.GetAnalyticsNodeIpAddress()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateAnalyticsNodeFromResource(object *AnalyticsNode, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("analytics_node_ip_address") {
		if val, ok := d.GetOk("analytics_node_ip_address"); ok {
			object.SetAnalyticsNodeIpAddress(val.(string))
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

func UpdateAnalyticsNodeRefsFromResource(object *AnalyticsNode, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceAnalyticsNodeCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAnalyticsNodeCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(AnalyticsNode)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceAnalyticsNodeCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "AnalyticsNode", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetAnalyticsNodeFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeCreate] Creation of resource AnalyticsNode on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceAnalyticsNodeRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAnalyticsNodeRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsCreate] Missing 'uuid' field for resource AnalyticsNode")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("analytics-node", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsCreate] Retrieving AnalyticsNode with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAnalyticsNode := obj.(*AnalyticsNode) // Fully set by Contrail backend
	if err := SetRefsAnalyticsNodeFromResource(objAnalyticsNode, d, m); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsCreate] Set refs on object AnalyticsNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAnalyticsNode.GetHref())
	if err := client.Update(objAnalyticsNode); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsCreate] Update refs for resource AnalyticsNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAnalyticsNode.GetUuid())
	return nil
}

func ResourceAnalyticsNodeRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAnalyticsNodeRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("analytics-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRead] Read resource analytics-node on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AnalyticsNode)
	WriteAnalyticsNodeToResource(*object, d, m)
	return nil
}

func ResourceAnalyticsNodeRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAnalyticsNodeRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("analytics-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsRead] Read resource analytics-node on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*AnalyticsNode)
	WriteAnalyticsNodeRefsToResource(*object, d, m)
	return nil
}

func ResourceAnalyticsNodeUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAnalyticsNodeUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("analytics-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeUpdate] Retrieving AnalyticsNode with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AnalyticsNode)
	UpdateAnalyticsNodeFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeUpdate] Update of resource AnalyticsNode on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAnalyticsNodeRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAnalyticsNodeRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("analytics-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsUpdate] Retrieving AnalyticsNode with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*AnalyticsNode)
	UpdateAnalyticsNodeRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsUpdate] Update of resource AnalyticsNode on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceAnalyticsNodeDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceAnalyticsNodeDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("analytics-node", d.Id()); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeDelete] Deletion of resource analytics-node on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceAnalyticsNodeRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceAnalyticsNodeRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsDelete] Missing 'uuid' field for resource AnalyticsNode")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("analytics-node", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsDelete] Retrieving AnalyticsNode with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objAnalyticsNode := obj.(*AnalyticsNode) // Fully set by Contrail backend
	if err := DeleteRefsAnalyticsNodeFromResource(objAnalyticsNode, d, m); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsDelete] Set refs on object AnalyticsNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objAnalyticsNode.GetHref())
	if err := client.Update(objAnalyticsNode); err != nil {
		return fmt.Errorf("[ResourceAnalyticsNodeRefsDelete] Delete refs for resource AnalyticsNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objAnalyticsNode.GetUuid())
	return nil
}

func ResourceAnalyticsNodeSchema() map[string]*schema.Schema {
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
		"analytics_node_ip_address": &schema.Schema{
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

func ResourceAnalyticsNodeRefsSchema() map[string]*schema.Schema {
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

func ResourceAnalyticsNode() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAnalyticsNodeCreate,
		Read:   ResourceAnalyticsNodeRead,
		Update: ResourceAnalyticsNodeUpdate,
		Delete: ResourceAnalyticsNodeDelete,
		Schema: ResourceAnalyticsNodeSchema(),
	}
}

func ResourceAnalyticsNodeRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceAnalyticsNodeRefsCreate,
		Read:   ResourceAnalyticsNodeRefsRead,
		Update: ResourceAnalyticsNodeRefsUpdate,
		Delete: ResourceAnalyticsNodeRefsDelete,
		Schema: ResourceAnalyticsNodeRefsSchema(),
	}
}

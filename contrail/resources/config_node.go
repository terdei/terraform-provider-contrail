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

func SetConfigNodeFromResource(object *ConfigNode, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetConfigNodeFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("config_node_ip_address"); ok {
		object.SetConfigNodeIpAddress(val.(string))
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

func SetRefsConfigNodeFromResource(object *ConfigNode, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsConfigNodeFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsConfigNodeFromResource(object *ConfigNode, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsConfigNodeFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteConfigNodeToResource(object ConfigNode, d *schema.ResourceData, m interface{}) {

	d.Set("config_node_ip_address", object.GetConfigNodeIpAddress())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteConfigNodeRefsToResource(object ConfigNode, d *schema.ResourceData, m interface{}) {

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

func TakeConfigNodeAsMap(object *ConfigNode) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["config_node_ip_address"] = object.GetConfigNodeIpAddress()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateConfigNodeFromResource(object *ConfigNode, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("config_node_ip_address") {
		if val, ok := d.GetOk("config_node_ip_address"); ok {
			object.SetConfigNodeIpAddress(val.(string))
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

func UpdateConfigNodeRefsFromResource(object *ConfigNode, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceConfigNodeCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceConfigNodeCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(ConfigNode)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceConfigNodeCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "ConfigNode", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceConfigNodeCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "ConfigNode", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetConfigNodeFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceConfigNodeCreate] Creation of resource ConfigNode on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceConfigNodeRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceConfigNodeRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceConfigNodeRefsCreate] Missing 'uuid' field for resource ConfigNode")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("config-node", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsCreate] Retrieving ConfigNode with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objConfigNode := obj.(*ConfigNode) // Fully set by Contrail backend
	if err := SetRefsConfigNodeFromResource(objConfigNode, d, m); err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsCreate] Set refs on object ConfigNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objConfigNode.GetHref())
	if err := client.Update(objConfigNode); err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsCreate] Update refs for resource ConfigNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objConfigNode.GetUuid())
	return nil
}

func ResourceConfigNodeRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigNodeRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("config-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceConfigNodeRead] Read resource config-node on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ConfigNode)
	WriteConfigNodeToResource(*object, d, m)
	return nil
}

func ResourceConfigNodeRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigNodeRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("config-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsRead] Read resource config-node on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*ConfigNode)
	WriteConfigNodeRefsToResource(*object, d, m)
	return nil
}

func ResourceConfigNodeUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigNodeUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("config-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceConfigNodeUpdate] Retrieving ConfigNode with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ConfigNode)
	UpdateConfigNodeFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceConfigNodeUpdate] Update of resource ConfigNode on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceConfigNodeRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigNodeRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("config-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsUpdate] Retrieving ConfigNode with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*ConfigNode)
	UpdateConfigNodeRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsUpdate] Update of resource ConfigNode on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceConfigNodeDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceConfigNodeDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("config-node", d.Id()); err != nil {
		return fmt.Errorf("[ResourceConfigNodeDelete] Deletion of resource config-node on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceConfigNodeRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceConfigNodeRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceConfigNodeRefsDelete] Missing 'uuid' field for resource ConfigNode")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("config-node", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsDelete] Retrieving ConfigNode with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objConfigNode := obj.(*ConfigNode) // Fully set by Contrail backend
	if err := DeleteRefsConfigNodeFromResource(objConfigNode, d, m); err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsDelete] Set refs on object ConfigNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objConfigNode.GetHref())
	if err := client.Update(objConfigNode); err != nil {
		return fmt.Errorf("[ResourceConfigNodeRefsDelete] Delete refs for resource ConfigNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objConfigNode.GetUuid())
	return nil
}

func ResourceConfigNodeSchema() map[string]*schema.Schema {
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
		"config_node_ip_address": &schema.Schema{
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

func ResourceConfigNodeRefsSchema() map[string]*schema.Schema {
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

func ResourceConfigNode() *schema.Resource {
	return &schema.Resource{
		Create: ResourceConfigNodeCreate,
		Read:   ResourceConfigNodeRead,
		Update: ResourceConfigNodeUpdate,
		Delete: ResourceConfigNodeDelete,
		Schema: ResourceConfigNodeSchema(),
	}
}

func ResourceConfigNodeRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceConfigNodeRefsCreate,
		Read:   ResourceConfigNodeRefsRead,
		Update: ResourceConfigNodeRefsUpdate,
		Delete: ResourceConfigNodeRefsDelete,
		Schema: ResourceConfigNodeRefsSchema(),
	}
}

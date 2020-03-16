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

func SetDatabaseNodeFromResource(object *DatabaseNode, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetDatabaseNodeFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("database_node_ip_address"); ok {
		object.SetDatabaseNodeIpAddress(val.(string))
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

func SetRefsDatabaseNodeFromResource(object *DatabaseNode, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsDatabaseNodeFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsDatabaseNodeFromResource(object *DatabaseNode, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsDatabaseNodeFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteDatabaseNodeToResource(object DatabaseNode, d *schema.ResourceData, m interface{}) {

	d.Set("database_node_ip_address", object.GetDatabaseNodeIpAddress())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteDatabaseNodeRefsToResource(object DatabaseNode, d *schema.ResourceData, m interface{}) {

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

func TakeDatabaseNodeAsMap(object *DatabaseNode) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["database_node_ip_address"] = object.GetDatabaseNodeIpAddress()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateDatabaseNodeFromResource(object *DatabaseNode, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("database_node_ip_address") {
		if val, ok := d.GetOk("database_node_ip_address"); ok {
			object.SetDatabaseNodeIpAddress(val.(string))
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

func UpdateDatabaseNodeRefsFromResource(object *DatabaseNode, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceDatabaseNodeCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDatabaseNodeCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(DatabaseNode)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		puuid_parts := strings.Split(puuid, "/")
		if len(puuid_parts) > 1 {
			parent, err := client.FindByUuid(puuid_parts[0], puuid_parts[1])
			if err != nil {
				return fmt.Errorf("[ResourceDatabaseNodeCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid_parts[1], puuid_parts[0], d.Get("name"), "DatabaseNode", client.GetServer(), err)
			}
			object.SetParent(parent)
		} else {
			parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
			if err != nil {
				return fmt.Errorf("[ResourceDatabaseNodeCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "DatabaseNode", client.GetServer(), err)
			}
			object.SetParent(parent)
		}
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetDatabaseNodeFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeCreate] Creation of resource DatabaseNode on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceDatabaseNodeRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDatabaseNodeRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDatabaseNodeRefsCreate] Missing 'uuid' field for resource DatabaseNode")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("database-node", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsCreate] Retrieving DatabaseNode with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDatabaseNode := obj.(*DatabaseNode) // Fully set by Contrail backend
	if err := SetRefsDatabaseNodeFromResource(objDatabaseNode, d, m); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsCreate] Set refs on object DatabaseNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDatabaseNode.GetHref())
	if err := client.Update(objDatabaseNode); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsCreate] Update refs for resource DatabaseNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDatabaseNode.GetUuid())
	return nil
}

func ResourceDatabaseNodeRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDatabaseNodeRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("database-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRead] Read resource database-node on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*DatabaseNode)
	WriteDatabaseNodeToResource(*object, d, m)
	return nil
}

func ResourceDatabaseNodeRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDatabaseNodeRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("database-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsRead] Read resource database-node on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*DatabaseNode)
	WriteDatabaseNodeRefsToResource(*object, d, m)
	return nil
}

func ResourceDatabaseNodeUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDatabaseNodeUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("database-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeUpdate] Retrieving DatabaseNode with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*DatabaseNode)
	UpdateDatabaseNodeFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeUpdate] Update of resource DatabaseNode on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDatabaseNodeRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDatabaseNodeRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("database-node", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsUpdate] Retrieving DatabaseNode with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*DatabaseNode)
	UpdateDatabaseNodeRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsUpdate] Update of resource DatabaseNode on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDatabaseNodeDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDatabaseNodeDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("database-node", d.Id()); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeDelete] Deletion of resource database-node on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceDatabaseNodeRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDatabaseNodeRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDatabaseNodeRefsDelete] Missing 'uuid' field for resource DatabaseNode")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("database-node", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsDelete] Retrieving DatabaseNode with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDatabaseNode := obj.(*DatabaseNode) // Fully set by Contrail backend
	if err := DeleteRefsDatabaseNodeFromResource(objDatabaseNode, d, m); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsDelete] Set refs on object DatabaseNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDatabaseNode.GetHref())
	if err := client.Update(objDatabaseNode); err != nil {
		return fmt.Errorf("[ResourceDatabaseNodeRefsDelete] Delete refs for resource DatabaseNode (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDatabaseNode.GetUuid())
	return nil
}

func ResourceDatabaseNodeSchema() map[string]*schema.Schema {
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
		"database_node_ip_address": &schema.Schema{
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

func ResourceDatabaseNodeRefsSchema() map[string]*schema.Schema {
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

func ResourceDatabaseNode() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDatabaseNodeCreate,
		Read:   ResourceDatabaseNodeRead,
		Update: ResourceDatabaseNodeUpdate,
		Delete: ResourceDatabaseNodeDelete,
		Schema: ResourceDatabaseNodeSchema(),
	}
}

func ResourceDatabaseNodeRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDatabaseNodeRefsCreate,
		Read:   ResourceDatabaseNodeRefsRead,
		Update: ResourceDatabaseNodeRefsUpdate,
		Delete: ResourceDatabaseNodeRefsDelete,
		Schema: ResourceDatabaseNodeRefsSchema(),
	}
}

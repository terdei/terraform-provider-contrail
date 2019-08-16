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

func SetRouteTableFromResource(object *RouteTable, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetRouteTableFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("routes"); ok {
		member := new(RouteTableType)
		SetRouteTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetRoutes(member)
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

func SetRefsRouteTableFromResource(object *RouteTable, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsRouteTableFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsRouteTableFromResource(object *RouteTable, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsRouteTableFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteRouteTableToResource(object RouteTable, d *schema.ResourceData, m interface{}) {

	routesObj := object.GetRoutes()
	d.Set("routes", TakeRouteTableTypeAsMap(&routesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeRouteTableAsMap(object *RouteTable) map[string]interface{} {
	omap := make(map[string]interface{})

	routesObj := object.GetRoutes()
	omap["routes"] = TakeRouteTableTypeAsMap(&routesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateRouteTableFromResource(object *RouteTable, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("routes") {
		if val, ok := d.GetOk("routes"); ok {
			member := new(RouteTableType)
			SetRouteTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetRoutes(member)
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

func ResourceRouteTableCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteTableCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(RouteTable)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceRouteTableCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "RouteTable", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetRouteTableFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceRouteTableCreate] Creation of resource RouteTable on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceRouteTableRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteTableRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRouteTableRefsCreate] Missing 'uuid' field for resource RouteTable")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("route-table", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRouteTableRefsCreate] Retrieving RouteTable with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRouteTable := obj.(*RouteTable) // Fully set by Contrail backend
	if err := SetRefsRouteTableFromResource(objRouteTable, d, m); err != nil {
		return fmt.Errorf("[ResourceRouteTableRefsCreate] Set refs on object RouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRouteTable.GetHref())
	if err := client.Update(objRouteTable); err != nil {
		return fmt.Errorf("[ResourceRouteTableRefsCreate] Update refs for resource RouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRouteTable.GetUuid())
	return nil
}

func ResourceRouteTableRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTableREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("route-table", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteTableRead] Read resource route-table on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*RouteTable)
	WriteRouteTableToResource(*object, d, m)
	return nil
}

func ResourceRouteTableRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTableRefsREAD")
	return nil
}

func ResourceRouteTableUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTableUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("route-table", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceRouteTableResourceUpdate] Retrieving RouteTable with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*RouteTable)
	UpdateRouteTableFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceRouteTableUpdate] Update of resource RouteTable on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceRouteTableRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTableRefsUpdate")
	return nil
}

func ResourceRouteTableDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceRouteTableDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("route-table", d.Id()); err != nil {
		return fmt.Errorf("[ResourceRouteTableDelete] Deletion of resource route-table on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceRouteTableRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceRouteTableRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceRouteTableRefsDelete] Missing 'uuid' field for resource RouteTable")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("route-table", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceRouteTableRefsDelete] Retrieving RouteTable with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objRouteTable := obj.(*RouteTable) // Fully set by Contrail backend
	if err := DeleteRefsRouteTableFromResource(objRouteTable, d, m); err != nil {
		return fmt.Errorf("[ResourceRouteTableRefsDelete] Set refs on object RouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objRouteTable.GetHref())
	if err := client.Update(objRouteTable); err != nil {
		return fmt.Errorf("[ResourceRouteTableRefsDelete] Delete refs for resource RouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objRouteTable.GetUuid())
	return nil
}

func ResourceRouteTableSchema() map[string]*schema.Schema {
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
		"routes": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTableType(),
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

func ResourceRouteTableRefsSchema() map[string]*schema.Schema {
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

func ResourceRouteTable() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteTableCreate,
		Read:   ResourceRouteTableRead,
		Update: ResourceRouteTableUpdate,
		Delete: ResourceRouteTableDelete,
		Schema: ResourceRouteTableSchema(),
	}
}

func ResourceRouteTableRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceRouteTableRefsCreate,
		Read:   ResourceRouteTableRefsRead,
		Update: ResourceRouteTableRefsUpdate,
		Delete: ResourceRouteTableRefsDelete,
		Schema: ResourceRouteTableRefsSchema(),
	}
}

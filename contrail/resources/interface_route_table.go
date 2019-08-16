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

func SetInterfaceRouteTableFromResource(object *InterfaceRouteTable, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetInterfaceRouteTableFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("interface_route_table_routes"); ok {
		member := new(RouteTableType)
		SetRouteTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetInterfaceRouteTableRoutes(member)
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

func SetRefsInterfaceRouteTableFromResource(object *InterfaceRouteTable, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsInterfaceRouteTableFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj, *dataObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-instance", refId.(string))
			dataObj := new(ServiceInterfaceTag)
			SetServiceInterfaceTagFromMap(dataObj, d, m, (v.(map[string]interface{}))["attr"])
			log.Printf("Data obj: %+v", dataObj)
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-instance by Uuid = %v as ref for ServiceInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceInstance(refObj.(*ServiceInstance), *dataObj)
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

func DeleteRefsInterfaceRouteTableFromResource(object *InterfaceRouteTable, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsInterfaceRouteTableFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.DeleteServiceInstance(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteServiceInstance(refId.(string))
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

func WriteInterfaceRouteTableToResource(object InterfaceRouteTable, d *schema.ResourceData, m interface{}) {

	interface_route_table_routesObj := object.GetInterfaceRouteTableRoutes()
	d.Set("interface_route_table_routes", TakeRouteTableTypeAsMap(&interface_route_table_routesObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeInterfaceRouteTableAsMap(object *InterfaceRouteTable) map[string]interface{} {
	omap := make(map[string]interface{})

	interface_route_table_routesObj := object.GetInterfaceRouteTableRoutes()
	omap["interface_route_table_routes"] = TakeRouteTableTypeAsMap(&interface_route_table_routesObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateInterfaceRouteTableFromResource(object *InterfaceRouteTable, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("interface_route_table_routes") {
		if val, ok := d.GetOk("interface_route_table_routes"); ok {
			member := new(RouteTableType)
			SetRouteTableTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetInterfaceRouteTableRoutes(member)
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

func ResourceInterfaceRouteTableCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceInterfaceRouteTableCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(InterfaceRouteTable)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceInterfaceRouteTableCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "InterfaceRouteTable", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetInterfaceRouteTableFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableCreate] Creation of resource InterfaceRouteTable on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceInterfaceRouteTableRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceInterfaceRouteTableRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsCreate] Missing 'uuid' field for resource InterfaceRouteTable")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("interface-route-table", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsCreate] Retrieving InterfaceRouteTable with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objInterfaceRouteTable := obj.(*InterfaceRouteTable) // Fully set by Contrail backend
	if err := SetRefsInterfaceRouteTableFromResource(objInterfaceRouteTable, d, m); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsCreate] Set refs on object InterfaceRouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objInterfaceRouteTable.GetHref())
	if err := client.Update(objInterfaceRouteTable); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsCreate] Update refs for resource InterfaceRouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objInterfaceRouteTable.GetUuid())
	return nil
}

func ResourceInterfaceRouteTableRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInterfaceRouteTableREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("interface-route-table", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRead] Read resource interface-route-table on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*InterfaceRouteTable)
	WriteInterfaceRouteTableToResource(*object, d, m)
	return nil
}

func ResourceInterfaceRouteTableRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInterfaceRouteTableRefsREAD")
	return nil
}

func ResourceInterfaceRouteTableUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInterfaceRouteTableUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("interface-route-table", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableResourceUpdate] Retrieving InterfaceRouteTable with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*InterfaceRouteTable)
	UpdateInterfaceRouteTableFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableUpdate] Update of resource InterfaceRouteTable on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceInterfaceRouteTableRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInterfaceRouteTableRefsUpdate")
	return nil
}

func ResourceInterfaceRouteTableDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceInterfaceRouteTableDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("interface-route-table", d.Id()); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableDelete] Deletion of resource interface-route-table on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceInterfaceRouteTableRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceInterfaceRouteTableRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsDelete] Missing 'uuid' field for resource InterfaceRouteTable")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("interface-route-table", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsDelete] Retrieving InterfaceRouteTable with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objInterfaceRouteTable := obj.(*InterfaceRouteTable) // Fully set by Contrail backend
	if err := DeleteRefsInterfaceRouteTableFromResource(objInterfaceRouteTable, d, m); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsDelete] Set refs on object InterfaceRouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objInterfaceRouteTable.GetHref())
	if err := client.Update(objInterfaceRouteTable); err != nil {
		return fmt.Errorf("[ResourceInterfaceRouteTableRefsDelete] Delete refs for resource InterfaceRouteTable (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objInterfaceRouteTable.GetUuid())
	return nil
}

func ResourceInterfaceRouteTableSchema() map[string]*schema.Schema {
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
		"interface_route_table_routes": &schema.Schema{
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

func ResourceInterfaceRouteTableRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_instance_refs": &schema.Schema{
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
						Elem:     ResourceServiceInterfaceTag(),
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

func ResourceInterfaceRouteTable() *schema.Resource {
	return &schema.Resource{
		Create: ResourceInterfaceRouteTableCreate,
		Read:   ResourceInterfaceRouteTableRead,
		Update: ResourceInterfaceRouteTableUpdate,
		Delete: ResourceInterfaceRouteTableDelete,
		Schema: ResourceInterfaceRouteTableSchema(),
	}
}

func ResourceInterfaceRouteTableRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceInterfaceRouteTableRefsCreate,
		Read:   ResourceInterfaceRouteTableRefsRead,
		Update: ResourceInterfaceRouteTableRefsUpdate,
		Delete: ResourceInterfaceRouteTableRefsDelete,
		Schema: ResourceInterfaceRouteTableRefsSchema(),
	}
}

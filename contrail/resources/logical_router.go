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

func SetLogicalRouterFromResource(object *LogicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetLogicalRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("configured_route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetConfiguredRouteTargetList(member)
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

func SetRefsLogicalRouterFromResource(object *LogicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsLogicalRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_machine_interface_refs"); ok {
		log.Printf("Got ref virtual_machine_interface_refs -- will call: object.AddVirtualMachineInterface(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-machine-interface", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-machine-interface by Uuid = %v as ref for VirtualMachineInterface on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualMachineInterface(refObj.(*VirtualMachineInterface))
		}
	}
	if val, ok := d.GetOk("route_target_refs"); ok {
		log.Printf("Got ref route_target_refs -- will call: object.AddRouteTarget(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("route-target", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving route-target by Uuid = %v as ref for RouteTarget on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddRouteTarget(refObj.(*RouteTarget))
		}
	}
	if val, ok := d.GetOk("route_table_refs"); ok {
		log.Printf("Got ref route_table_refs -- will call: object.AddRouteTable(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("route-table", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving route-table by Uuid = %v as ref for RouteTable on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddRouteTable(refObj.(*RouteTable))
		}
	}
	if val, ok := d.GetOk("virtual_network_refs"); ok {
		log.Printf("Got ref virtual_network_refs -- will call: object.AddVirtualNetwork(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-network", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-network by Uuid = %v as ref for VirtualNetwork on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualNetwork(refObj.(*VirtualNetwork))
		}
	}
	if val, ok := d.GetOk("service_instance_refs"); ok {
		log.Printf("Got ref service_instance_refs -- will call: object.AddServiceInstance(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("service-instance", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving service-instance by Uuid = %v as ref for ServiceInstance on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddServiceInstance(refObj.(*ServiceInstance))
		}
	}

	return nil
}

func WriteLogicalRouterToResource(object LogicalRouter, d *schema.ResourceData, m interface{}) {

	configured_route_target_listObj := object.GetConfiguredRouteTargetList()
	d.Set("configured_route_target_list", TakeRouteTargetListAsMap(&configured_route_target_listObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeLogicalRouterAsMap(object *LogicalRouter) map[string]interface{} {
	omap := make(map[string]interface{})

	configured_route_target_listObj := object.GetConfiguredRouteTargetList()
	omap["configured_route_target_list"] = TakeRouteTargetListAsMap(&configured_route_target_listObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateLogicalRouterFromResource(object *LogicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("configured_route_target_list") {
		if val, ok := d.GetOk("configured_route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetConfiguredRouteTargetList(member)
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

func ResourceLogicalRouterCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLogicalRouterCreate")
	//log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(LogicalRouter)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceLogicalRouterCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "LogicalRouter", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetLogicalRouterFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceLogicalRouterCreate] Creation of resource LogicalRouter on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceLogicalRouterRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceLogicalRouterRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceLogicalRouterRefsCreate] Missing 'uuid' field for resource LogicalRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("logical-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceLogicalRouterRefsCreate] Retrieving LogicalRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objLogicalRouter := obj.(*LogicalRouter) // Fully set by Contrail backend
	if err := SetRefsLogicalRouterFromResource(objLogicalRouter, d, m); err != nil {
		return fmt.Errorf("[ResourceLogicalRouterRefsCreate] Set refs on object LogicalRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objLogicalRouter.GetHref())
	if err := client.Update(objLogicalRouter); err != nil {
		return fmt.Errorf("[ResourceLogicalRouterRefsCreate] Update refs for resource LogicalRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objLogicalRouter.GetUuid())
	return nil
}

func ResourceLogicalRouterRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalRouterREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("logical-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLogicalRouterRead] Read resource logical-router on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*LogicalRouter)
	WriteLogicalRouterToResource(*object, d, m)
	return nil
}

func ResourceLogicalRouterRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalRouterRefsREAD")
	return nil
}

func ResourceLogicalRouterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalRouterUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("logical-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceLogicalRouterResourceUpdate] Retrieving LogicalRouter with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*LogicalRouter)
	UpdateLogicalRouterFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceLogicalRouterUpdate] Update of resource LogicalRouter on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceLogicalRouterRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalRouterRefsUpdate")
	return nil
}

func ResourceLogicalRouterDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalRouterDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("logical-router", d.Id()); err != nil {
		return fmt.Errorf("[ResourceLogicalRouterDelete] Deletion of resource logical-router on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceLogicalRouterRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceLogicalRouterRefsDelete: %v", d.Id())
	return nil
}

func ResourceLogicalRouterSchema() map[string]*schema.Schema {
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
		"configured_route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
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

func ResourceLogicalRouterRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_machine_interface_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualMachineInterface(),
		},
		"route_target_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTarget(),
		},
		"route_table_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTable(),
		},
		"virtual_network_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceVirtualNetwork(),
		},
		"service_instance_refs": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceServiceInstance(),
		},
	}
}

func ResourceLogicalRouter() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLogicalRouterCreate,
		Read:   ResourceLogicalRouterRead,
		Update: ResourceLogicalRouterUpdate,
		Delete: ResourceLogicalRouterDelete,
		Schema: ResourceLogicalRouterSchema(),
	}
}

func ResourceLogicalRouterRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceLogicalRouterRefsCreate,
		Read:   ResourceLogicalRouterRefsRead,
		Update: ResourceLogicalRouterRefsUpdate,
		Delete: ResourceLogicalRouterRefsDelete,
		Schema: ResourceLogicalRouterRefsSchema(),
	}
}

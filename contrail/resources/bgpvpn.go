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

func SetBgpvpnFromResource(object *Bgpvpn, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetBgpvpnFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetRouteTargetList(member)
	}
	if val, ok := d.GetOk("import_route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetImportRouteTargetList(member)
	}
	if val, ok := d.GetOk("export_route_target_list"); ok {
		member := new(RouteTargetList)
		SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetExportRouteTargetList(member)
	}
	if val, ok := d.GetOk("bgpvpn_type"); ok {
		object.SetBgpvpnType(val.(string))
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

func SetRefsBgpvpnFromResource(object *Bgpvpn, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsBgpvpnFromResource] key = %v, prefix = %v", key, prefix)
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

func DeleteRefsBgpvpnFromResource(object *Bgpvpn, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsBgpvpnFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteBgpvpnToResource(object Bgpvpn, d *schema.ResourceData, m interface{}) {

	route_target_listObj := object.GetRouteTargetList()
	d.Set("route_target_list", TakeRouteTargetListAsMap(&route_target_listObj))
	import_route_target_listObj := object.GetImportRouteTargetList()
	d.Set("import_route_target_list", TakeRouteTargetListAsMap(&import_route_target_listObj))
	export_route_target_listObj := object.GetExportRouteTargetList()
	d.Set("export_route_target_list", TakeRouteTargetListAsMap(&export_route_target_listObj))
	d.Set("bgpvpn_type", object.GetBgpvpnType())
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func WriteBgpvpnRefsToResource(object Bgpvpn, d *schema.ResourceData, m interface{}) {

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

func TakeBgpvpnAsMap(object *Bgpvpn) map[string]interface{} {
	omap := make(map[string]interface{})

	route_target_listObj := object.GetRouteTargetList()
	omap["route_target_list"] = TakeRouteTargetListAsMap(&route_target_listObj)
	import_route_target_listObj := object.GetImportRouteTargetList()
	omap["import_route_target_list"] = TakeRouteTargetListAsMap(&import_route_target_listObj)
	export_route_target_listObj := object.GetExportRouteTargetList()
	omap["export_route_target_list"] = TakeRouteTargetListAsMap(&export_route_target_listObj)
	omap["bgpvpn_type"] = object.GetBgpvpnType()
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateBgpvpnFromResource(object *Bgpvpn, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("route_target_list") {
		if val, ok := d.GetOk("route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetRouteTargetList(member)
		}
	}
	if d.HasChange("import_route_target_list") {
		if val, ok := d.GetOk("import_route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetImportRouteTargetList(member)
		}
	}
	if d.HasChange("export_route_target_list") {
		if val, ok := d.GetOk("export_route_target_list"); ok {
			member := new(RouteTargetList)
			SetRouteTargetListFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetExportRouteTargetList(member)
		}
	}
	if d.HasChange("bgpvpn_type") {
		if val, ok := d.GetOk("bgpvpn_type"); ok {
			object.SetBgpvpnType(val.(string))
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

func UpdateBgpvpnRefsFromResource(object *Bgpvpn, d *schema.ResourceData, m interface{}, prefix ...string) {
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

func ResourceBgpvpnCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpvpnCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Bgpvpn)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceBgpvpnCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Bgpvpn", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetBgpvpnFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceBgpvpnCreate] Creation of resource Bgpvpn on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceBgpvpnRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpvpnRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpvpnRefsCreate] Missing 'uuid' field for resource Bgpvpn")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgpvpn", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsCreate] Retrieving Bgpvpn with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpvpn := obj.(*Bgpvpn) // Fully set by Contrail backend
	if err := SetRefsBgpvpnFromResource(objBgpvpn, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsCreate] Set refs on object Bgpvpn (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpvpn.GetHref())
	if err := client.Update(objBgpvpn); err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsCreate] Update refs for resource Bgpvpn (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpvpn.GetUuid())
	return nil
}

func ResourceBgpvpnRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpvpnRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgpvpn", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpvpnRead] Read resource bgpvpn on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Bgpvpn)
	WriteBgpvpnToResource(*object, d, m)
	return nil
}

func ResourceBgpvpnRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpvpnRefsRead")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("bgpvpn", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsRead] Read resource bgpvpn on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Bgpvpn)
	WriteBgpvpnRefsToResource(*object, d, m)
	return nil
}

func ResourceBgpvpnUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpvpnUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgpvpn", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpvpnUpdate] Retrieving Bgpvpn with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Bgpvpn)
	UpdateBgpvpnFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpvpnUpdate] Update of resource Bgpvpn on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpvpnRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpvpnRefsUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("bgpvpn", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsUpdate] Retrieving Bgpvpn with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Bgpvpn)
	UpdateBgpvpnRefsFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsUpdate] Update of resource Bgpvpn on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceBgpvpnDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceBgpvpnDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("bgpvpn", d.Id()); err != nil {
		return fmt.Errorf("[ResourceBgpvpnDelete] Deletion of resource bgpvpn on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceBgpvpnRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceBgpvpnRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceBgpvpnRefsDelete] Missing 'uuid' field for resource Bgpvpn")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("bgpvpn", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsDelete] Retrieving Bgpvpn with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objBgpvpn := obj.(*Bgpvpn) // Fully set by Contrail backend
	if err := DeleteRefsBgpvpnFromResource(objBgpvpn, d, m); err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsDelete] Set refs on object Bgpvpn (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objBgpvpn.GetHref())
	if err := client.Update(objBgpvpn); err != nil {
		return fmt.Errorf("[ResourceBgpvpnRefsDelete] Delete refs for resource Bgpvpn (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objBgpvpn.GetUuid())
	return nil
}

func ResourceBgpvpnSchema() map[string]*schema.Schema {
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
		"route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
		},
		"import_route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
		},
		"export_route_target_list": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceRouteTargetList(),
		},
		"bgpvpn_type": &schema.Schema{
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

func ResourceBgpvpnRefsSchema() map[string]*schema.Schema {
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

func ResourceBgpvpn() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpvpnCreate,
		Read:   ResourceBgpvpnRead,
		Update: ResourceBgpvpnUpdate,
		Delete: ResourceBgpvpnDelete,
		Schema: ResourceBgpvpnSchema(),
	}
}

func ResourceBgpvpnRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceBgpvpnRefsCreate,
		Read:   ResourceBgpvpnRefsRead,
		Update: ResourceBgpvpnRefsUpdate,
		Delete: ResourceBgpvpnRefsDelete,
		Schema: ResourceBgpvpnRefsSchema(),
	}
}

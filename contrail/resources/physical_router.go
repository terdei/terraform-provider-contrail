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

func SetPhysicalRouterFromResource(object *PhysicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetPhysicalRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("physical_router_management_ip"); ok {
		object.SetPhysicalRouterManagementIp(val.(string))
	}
	if val, ok := d.GetOk("physical_router_dataplane_ip"); ok {
		object.SetPhysicalRouterDataplaneIp(val.(string))
	}
	if val, ok := d.GetOk("physical_router_loopback_ip"); ok {
		object.SetPhysicalRouterLoopbackIp(val.(string))
	}
	if val, ok := d.GetOk("physical_router_vendor_name"); ok {
		object.SetPhysicalRouterVendorName(val.(string))
	}
	if val, ok := d.GetOk("physical_router_product_name"); ok {
		object.SetPhysicalRouterProductName(val.(string))
	}
	if val, ok := d.GetOk("physical_router_vnc_managed"); ok {
		object.SetPhysicalRouterVncManaged(val.(bool))
	}
	if val, ok := d.GetOk("physical_router_role"); ok {
		object.SetPhysicalRouterRole(val.(string))
	}
	if val, ok := d.GetOk("physical_router_snmp"); ok {
		object.SetPhysicalRouterSnmp(val.(bool))
	}
	if val, ok := d.GetOk("physical_router_lldp"); ok {
		object.SetPhysicalRouterLldp(val.(bool))
	}
	if val, ok := d.GetOk("physical_router_user_credentials"); ok {
		member := new(UserCredentials)
		SetUserCredentialsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetPhysicalRouterUserCredentials(member)
	}
	if val, ok := d.GetOk("physical_router_snmp_credentials"); ok {
		member := new(SNMPCredentials)
		SetSNMPCredentialsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetPhysicalRouterSnmpCredentials(member)
	}
	if val, ok := d.GetOk("physical_router_junos_service_ports"); ok {
		member := new(JunosServicePorts)
		SetJunosServicePortsFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetPhysicalRouterJunosServicePorts(member)
	}
	if val, ok := d.GetOk("telemetry_info"); ok {
		member := new(TelemetryStateInfo)
		SetTelemetryStateInfoFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetTelemetryInfo(member)
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

func SetRefsPhysicalRouterFromResource(object *PhysicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsPhysicalRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_router_refs"); ok {
		log.Printf("Got ref virtual_router_refs -- will call: object.AddVirtualRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("virtual-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving virtual-router by Uuid = %v as ref for VirtualRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddVirtualRouter(refObj.(*VirtualRouter))
		}
	}
	if val, ok := d.GetOk("bgp_router_refs"); ok {
		log.Printf("Got ref bgp_router_refs -- will call: object.AddBgpRouter(refObj)")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			log.Printf("Ref 'to': %#v (str->%v)", refId, refId.(string))
			refObj, err := client.FindByUuid("bgp-router", refId.(string))
			if err != nil {
				return fmt.Errorf("[SnippetSetObjRef] Retrieving bgp-router by Uuid = %v as ref for BgpRouter on %v (%v)", refId, client.GetServer(), err)
			}
			log.Printf("Ref 'to' (OBJECT): %+v", refObj)
			object.AddBgpRouter(refObj.(*BgpRouter))
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

func DeleteRefsPhysicalRouterFromResource(object *PhysicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[DeleteRefsPhysicalRouterFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("virtual_router_refs"); ok {
		log.Printf("Got ref virtual_router_refs -- will call: object.DeleteVirtualRouter(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualRouter(refId.(string))
		}
	}
	if val, ok := d.GetOk("bgp_router_refs"); ok {
		log.Printf("Got ref bgp_router_refs -- will call: object.DeleteBgpRouter(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteBgpRouter(refId.(string))
		}
	}
	if val, ok := d.GetOk("virtual_network_refs"); ok {
		log.Printf("Got ref virtual_network_refs -- will call: object.DeleteVirtualNetwork(refObj.(string))")
		for k, v := range val.([]interface{}) {
			log.Printf("Item: %+v => <%T> %+v", k, v, v)
			refId := (v.(map[string]interface{}))["to"]
			object.DeleteVirtualNetwork(refId.(string))
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

func WritePhysicalRouterToResource(object PhysicalRouter, d *schema.ResourceData, m interface{}) {

	d.Set("physical_router_management_ip", object.GetPhysicalRouterManagementIp())
	d.Set("physical_router_dataplane_ip", object.GetPhysicalRouterDataplaneIp())
	d.Set("physical_router_loopback_ip", object.GetPhysicalRouterLoopbackIp())
	d.Set("physical_router_vendor_name", object.GetPhysicalRouterVendorName())
	d.Set("physical_router_product_name", object.GetPhysicalRouterProductName())
	d.Set("physical_router_vnc_managed", object.GetPhysicalRouterVncManaged())
	d.Set("physical_router_role", object.GetPhysicalRouterRole())
	d.Set("physical_router_snmp", object.GetPhysicalRouterSnmp())
	d.Set("physical_router_lldp", object.GetPhysicalRouterLldp())
	physical_router_user_credentialsObj := object.GetPhysicalRouterUserCredentials()
	d.Set("physical_router_user_credentials", TakeUserCredentialsAsMap(&physical_router_user_credentialsObj))
	physical_router_snmp_credentialsObj := object.GetPhysicalRouterSnmpCredentials()
	d.Set("physical_router_snmp_credentials", TakeSNMPCredentialsAsMap(&physical_router_snmp_credentialsObj))
	physical_router_junos_service_portsObj := object.GetPhysicalRouterJunosServicePorts()
	d.Set("physical_router_junos_service_ports", TakeJunosServicePortsAsMap(&physical_router_junos_service_portsObj))
	telemetry_infoObj := object.GetTelemetryInfo()
	d.Set("telemetry_info", TakeTelemetryStateInfoAsMap(&telemetry_infoObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakePhysicalRouterAsMap(object *PhysicalRouter) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["physical_router_management_ip"] = object.GetPhysicalRouterManagementIp()
	omap["physical_router_dataplane_ip"] = object.GetPhysicalRouterDataplaneIp()
	omap["physical_router_loopback_ip"] = object.GetPhysicalRouterLoopbackIp()
	omap["physical_router_vendor_name"] = object.GetPhysicalRouterVendorName()
	omap["physical_router_product_name"] = object.GetPhysicalRouterProductName()
	omap["physical_router_vnc_managed"] = object.GetPhysicalRouterVncManaged()
	omap["physical_router_role"] = object.GetPhysicalRouterRole()
	omap["physical_router_snmp"] = object.GetPhysicalRouterSnmp()
	omap["physical_router_lldp"] = object.GetPhysicalRouterLldp()
	physical_router_user_credentialsObj := object.GetPhysicalRouterUserCredentials()
	omap["physical_router_user_credentials"] = TakeUserCredentialsAsMap(&physical_router_user_credentialsObj)
	physical_router_snmp_credentialsObj := object.GetPhysicalRouterSnmpCredentials()
	omap["physical_router_snmp_credentials"] = TakeSNMPCredentialsAsMap(&physical_router_snmp_credentialsObj)
	physical_router_junos_service_portsObj := object.GetPhysicalRouterJunosServicePorts()
	omap["physical_router_junos_service_ports"] = TakeJunosServicePortsAsMap(&physical_router_junos_service_portsObj)
	telemetry_infoObj := object.GetTelemetryInfo()
	omap["telemetry_info"] = TakeTelemetryStateInfoAsMap(&telemetry_infoObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdatePhysicalRouterFromResource(object *PhysicalRouter, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("physical_router_management_ip") {
		if val, ok := d.GetOk("physical_router_management_ip"); ok {
			object.SetPhysicalRouterManagementIp(val.(string))
		}
	}
	if d.HasChange("physical_router_dataplane_ip") {
		if val, ok := d.GetOk("physical_router_dataplane_ip"); ok {
			object.SetPhysicalRouterDataplaneIp(val.(string))
		}
	}
	if d.HasChange("physical_router_loopback_ip") {
		if val, ok := d.GetOk("physical_router_loopback_ip"); ok {
			object.SetPhysicalRouterLoopbackIp(val.(string))
		}
	}
	if d.HasChange("physical_router_vendor_name") {
		if val, ok := d.GetOk("physical_router_vendor_name"); ok {
			object.SetPhysicalRouterVendorName(val.(string))
		}
	}
	if d.HasChange("physical_router_product_name") {
		if val, ok := d.GetOk("physical_router_product_name"); ok {
			object.SetPhysicalRouterProductName(val.(string))
		}
	}
	if d.HasChange("physical_router_vnc_managed") {
		if val, ok := d.GetOk("physical_router_vnc_managed"); ok {
			object.SetPhysicalRouterVncManaged(val.(bool))
		}
	}
	if d.HasChange("physical_router_role") {
		if val, ok := d.GetOk("physical_router_role"); ok {
			object.SetPhysicalRouterRole(val.(string))
		}
	}
	if d.HasChange("physical_router_snmp") {
		if val, ok := d.GetOk("physical_router_snmp"); ok {
			object.SetPhysicalRouterSnmp(val.(bool))
		}
	}
	if d.HasChange("physical_router_lldp") {
		if val, ok := d.GetOk("physical_router_lldp"); ok {
			object.SetPhysicalRouterLldp(val.(bool))
		}
	}
	if d.HasChange("physical_router_user_credentials") {
		if val, ok := d.GetOk("physical_router_user_credentials"); ok {
			member := new(UserCredentials)
			SetUserCredentialsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetPhysicalRouterUserCredentials(member)
		}
	}
	if d.HasChange("physical_router_snmp_credentials") {
		if val, ok := d.GetOk("physical_router_snmp_credentials"); ok {
			member := new(SNMPCredentials)
			SetSNMPCredentialsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetPhysicalRouterSnmpCredentials(member)
		}
	}
	if d.HasChange("physical_router_junos_service_ports") {
		if val, ok := d.GetOk("physical_router_junos_service_ports"); ok {
			member := new(JunosServicePorts)
			SetJunosServicePortsFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetPhysicalRouterJunosServicePorts(member)
		}
	}
	if d.HasChange("telemetry_info") {
		if val, ok := d.GetOk("telemetry_info"); ok {
			member := new(TelemetryStateInfo)
			SetTelemetryStateInfoFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetTelemetryInfo(member)
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

func ResourcePhysicalRouterCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePhysicalRouterCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(PhysicalRouter)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourcePhysicalRouterCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "PhysicalRouter", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetPhysicalRouterFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterCreate] Creation of resource PhysicalRouter on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourcePhysicalRouterRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePhysicalRouterRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePhysicalRouterRefsCreate] Missing 'uuid' field for resource PhysicalRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("physical-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRefsCreate] Retrieving PhysicalRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPhysicalRouter := obj.(*PhysicalRouter) // Fully set by Contrail backend
	if err := SetRefsPhysicalRouterFromResource(objPhysicalRouter, d, m); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRefsCreate] Set refs on object PhysicalRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPhysicalRouter.GetHref())
	if err := client.Update(objPhysicalRouter); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRefsCreate] Update refs for resource PhysicalRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPhysicalRouter.GetUuid())
	return nil
}

func ResourcePhysicalRouterRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalRouterREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("physical-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRead] Read resource physical-router on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*PhysicalRouter)
	WritePhysicalRouterToResource(*object, d, m)
	return nil
}

func ResourcePhysicalRouterRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalRouterRefsREAD")
	return nil
}

func ResourcePhysicalRouterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalRouterUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("physical-router", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterResourceUpdate] Retrieving PhysicalRouter with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*PhysicalRouter)
	UpdatePhysicalRouterFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterUpdate] Update of resource PhysicalRouter on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourcePhysicalRouterRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalRouterRefsUpdate")
	return nil
}

func ResourcePhysicalRouterDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourcePhysicalRouterDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("physical-router", d.Id()); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterDelete] Deletion of resource physical-router on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourcePhysicalRouterRefsDelete(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourcePhysicalRouterRefsDelete")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourcePhysicalRouterRefsDelete] Missing 'uuid' field for resource PhysicalRouter")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("physical-router", uuid)
	if err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRefsDelete] Retrieving PhysicalRouter with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objPhysicalRouter := obj.(*PhysicalRouter) // Fully set by Contrail backend
	if err := DeleteRefsPhysicalRouterFromResource(objPhysicalRouter, d, m); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRefsDelete] Set refs on object PhysicalRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objPhysicalRouter.GetHref())
	if err := client.Update(objPhysicalRouter); err != nil {
		return fmt.Errorf("[ResourcePhysicalRouterRefsDelete] Delete refs for resource PhysicalRouter (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objPhysicalRouter.GetUuid())
	return nil
}

func ResourcePhysicalRouterSchema() map[string]*schema.Schema {
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
		"physical_router_management_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"physical_router_dataplane_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"physical_router_loopback_ip": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"physical_router_vendor_name": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"physical_router_product_name": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"physical_router_vnc_managed": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"physical_router_role": &schema.Schema{
			Optional: true,
			Type:     schema.TypeString,
		},
		"physical_router_snmp": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"physical_router_lldp": &schema.Schema{
			Optional: true,
			Type:     schema.TypeBool,
		},
		"physical_router_user_credentials": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceUserCredentials(),
		},
		"physical_router_snmp_credentials": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceSNMPCredentials(),
		},
		"physical_router_junos_service_ports": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceJunosServicePorts(),
		},
		"telemetry_info": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTelemetryStateInfo(),
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

func ResourcePhysicalRouterRefsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"virtual_router_refs": &schema.Schema{
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
		"bgp_router_refs": &schema.Schema{
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
		"virtual_network_refs": &schema.Schema{
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

func ResourcePhysicalRouter() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePhysicalRouterCreate,
		Read:   ResourcePhysicalRouterRead,
		Update: ResourcePhysicalRouterUpdate,
		Delete: ResourcePhysicalRouterDelete,
		Schema: ResourcePhysicalRouterSchema(),
	}
}

func ResourcePhysicalRouterRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourcePhysicalRouterRefsCreate,
		Read:   ResourcePhysicalRouterRefsRead,
		Update: ResourcePhysicalRouterRefsUpdate,
		Delete: ResourcePhysicalRouterRefsDelete,
		Schema: ResourcePhysicalRouterRefsSchema(),
	}
}

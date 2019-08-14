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

func SetDomainFromResource(object *Domain, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	log.Printf("[SetDomainFromResource] key = %v, prefix = %v", key, prefix)
	if val, ok := d.GetOk("domain_limits"); ok {
		member := new(DomainLimitsType)
		SetDomainLimitsTypeFromMap(member, d, m, (val.([]interface{}))[0])
		object.SetDomainLimits(member)
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

func SetRefsDomainFromResource(object *Domain, d *schema.ResourceData, m interface{}, prefix ...string) error {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	log.Printf("[SetRefsDomainFromResource] key = %v, prefix = %v", key, prefix)
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

func WriteDomainToResource(object Domain, d *schema.ResourceData, m interface{}) {

	domain_limitsObj := object.GetDomainLimits()
	d.Set("domain_limits", TakeDomainLimitsTypeAsMap(&domain_limitsObj))
	id_permsObj := object.GetIdPerms()
	d.Set("id_perms", TakeIdPermsTypeAsMap(&id_permsObj))
	perms2Obj := object.GetPerms2()
	d.Set("perms2", TakePermType2AsMap(&perms2Obj))
	annotationsObj := object.GetAnnotations()
	d.Set("annotations", TakeKeyValuePairsAsMap(&annotationsObj))
	d.Set("display_name", object.GetDisplayName())

}

func TakeDomainAsMap(object *Domain) map[string]interface{} {
	omap := make(map[string]interface{})

	domain_limitsObj := object.GetDomainLimits()
	omap["domain_limits"] = TakeDomainLimitsTypeAsMap(&domain_limitsObj)
	id_permsObj := object.GetIdPerms()
	omap["id_perms"] = TakeIdPermsTypeAsMap(&id_permsObj)
	perms2Obj := object.GetPerms2()
	omap["perms2"] = TakePermType2AsMap(&perms2Obj)
	annotationsObj := object.GetAnnotations()
	omap["annotations"] = TakeKeyValuePairsAsMap(&annotationsObj)
	omap["display_name"] = object.GetDisplayName()

	return omap
}

func UpdateDomainFromResource(object *Domain, d *schema.ResourceData, m interface{}, prefix ...string) {
	key := strings.Join(prefix, ".")
	if len(key) != 0 {
		key = key + "."
	}

	if d.HasChange("domain_limits") {
		if val, ok := d.GetOk("domain_limits"); ok {
			member := new(DomainLimitsType)
			SetDomainLimitsTypeFromMap(member, d, m, (val.([]interface{}))[0])
			object.SetDomainLimits(member)
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

func ResourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDomainCreate")
	log.Print(spew.Sdump(d))
	// SPEW
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	object := new(Domain)
	object.SetName(d.Get("name").(string))
	if puuid_obj, ok := d.GetOk("parent_uuid"); ok {
		puuid := puuid_obj.(string)
		parent, err := client.FindByUuid(object.GetDefaultParentType(), puuid)
		if err != nil {
			return fmt.Errorf("[ResourceDomainCreate] retrieving Parent with uuid %s of type %s for resource %s (%s) - on %v (%v)", puuid, object.GetDefaultParentType(), d.Get("name"), "Domain", client.GetServer(), err)
		}
		object.SetParent(parent)
	}
	//object.SetFQName(object.GetDefaultParentType(), strings.Split(d.Get("parent_fq_name").(string) + ":" + d.Get("name").(string), ":"))
	SetDomainFromResource(object, d, m)

	if err := client.Create(object); err != nil {
		return fmt.Errorf("[ResourceDomainCreate] Creation of resource Domain on %v: (%v)", client.GetServer(), err)
	}
	d.SetId(object.GetUuid())
	return nil
}

func ResourceDomainRefsCreate(d *schema.ResourceData, m interface{}) error {
	// SPEW
	log.Printf("ResourceDomainRefsCreate")
	//log.Printf("SPEW: %v", spew.Sdump(d))
	// SPEW

	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	uuid_obj, ok := d.GetOk("uuid")
	if ok == false {
		return fmt.Errorf("[ResourceDomainRefsCreate] Missing 'uuid' field for resource Domain")
	}
	uuid := uuid_obj.(string)
	obj, err := client.FindByUuid("domain", uuid)
	if err != nil {
		return fmt.Errorf("[ResourceDomainRefsCreate] Retrieving Domain with uuid %s on %v (%v)", uuid, client.GetServer(), err)
	}
	objDomain := obj.(*Domain) // Fully set by Contrail backend
	if err := SetRefsDomainFromResource(objDomain, d, m); err != nil {
		return fmt.Errorf("[ResourceDomainRefsCreate] Set refs on object Domain (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	log.Printf("Object href: %v", objDomain.GetHref())
	if err := client.Update(objDomain); err != nil {
		return fmt.Errorf("[ResourceDomainRefsCreate] Update refs for resource Domain (uuid: %v) on %v (%v)", uuid, client.GetServer(), err)
	}
	d.SetId(objDomain.GetUuid())
	return nil
}

func ResourceDomainRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDomainREAD")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	base, err := client.FindByUuid("domain", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDomainRead] Read resource domain on %v: (%v)", client.GetServer(), err)
	}
	object := base.(*Domain)
	WriteDomainToResource(*object, d, m)
	return nil
}

func ResourceDomainRefsRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDomainRefsREAD")
	return nil
}

func ResourceDomainUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDomainUpdate")
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	obj, err := client.FindByUuid("domain", d.Id())
	if err != nil {
		return fmt.Errorf("[ResourceDomainResourceUpdate] Retrieving Domain with uuid %s on %v (%v)", d.Id(), client.GetServer(), err)
	}
	uobject := obj.(*Domain)
	UpdateDomainFromResource(uobject, d, m)

	log.Printf("Object href: %v", uobject.GetHref())
	if err := client.Update(uobject); err != nil {
		return fmt.Errorf("[ResourceDomainUpdate] Update of resource Domain on %v: (%v)", client.GetServer(), err)
	}
	return nil
}

func ResourceDomainRefsUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDomainRefsUpdate")
	return nil
}

func ResourceDomainDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDomainDelete: %v", d.Id())
	client := m.(*contrail.Client)
	client.GetServer() // dummy call
	if err := client.DeleteByUuid("domain", d.Id()); err != nil {
		return fmt.Errorf("[ResourceDomainDelete] Deletion of resource domain on %v: (%v)", client.GetServer(), err)
	}
	d.SetId("")
	return nil
}

func ResourceDomainRefsDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("ResourceDomainRefsDelete: %v", d.Id())
	return nil
}

func ResourceDomainSchema() map[string]*schema.Schema {
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
		"domain_limits": &schema.Schema{
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDomainLimitsType(),
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

func ResourceDomainRefsSchema() map[string]*schema.Schema {
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

func ResourceDomain() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDomainCreate,
		Read:   ResourceDomainRead,
		Update: ResourceDomainUpdate,
		Delete: ResourceDomainDelete,
		Schema: ResourceDomainSchema(),
	}
}

func ResourceDomainRefs() *schema.Resource {
	return &schema.Resource{
		Create: ResourceDomainRefsCreate,
		Read:   ResourceDomainRefsRead,
		Update: ResourceDomainRefsUpdate,
		Delete: ResourceDomainRefsDelete,
		Schema: ResourceDomainRefsSchema(),
	}
}

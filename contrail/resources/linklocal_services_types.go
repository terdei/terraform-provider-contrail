//
// Automatically generated. DO NOT EDIT.
// (Struct Type [aka CType])
//

package resources

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

var _ = spew.Dump // Avoid import errors if not used

func SetLinklocalServiceEntryTypeFromMap(object *LinklocalServiceEntryType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLinklocalServiceEntryTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mLinklocalServiceNameObj := vmap["linklocal_service_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLinklocalServiceNameObj) {
		log.Printf("Setting linklocal_service_name  LinklocalServiceName <<%T>> => %#v", mLinklocalServiceNameObj, mLinklocalServiceNameObj)
		mLinklocalServiceName := mLinklocalServiceNameObj.(string)
		object.LinklocalServiceName = mLinklocalServiceName
	}

	mLinklocalServiceIpObj := vmap["linklocal_service_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLinklocalServiceIpObj) {
		log.Printf("Setting linklocal_service_ip  LinklocalServiceIp <<%T>> => %#v", mLinklocalServiceIpObj, mLinklocalServiceIpObj)
		mLinklocalServiceIp := mLinklocalServiceIpObj.(string)
		object.LinklocalServiceIp = mLinklocalServiceIp
	}

	mLinklocalServicePortObj := vmap["linklocal_service_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mLinklocalServicePortObj) {
		log.Printf("Setting linklocal_service_port  LinklocalServicePort <<%T>> => %#v", mLinklocalServicePortObj, mLinklocalServicePortObj)
		mLinklocalServicePort := mLinklocalServicePortObj.(int)
		object.LinklocalServicePort = mLinklocalServicePort
	}

	mIpFabricDnsServiceNameObj := vmap["ip_fabric_DNS_service_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpFabricDnsServiceNameObj) {
		log.Printf("Setting ip_fabric_DNS_service_name  IpFabricDnsServiceName <<%T>> => %#v", mIpFabricDnsServiceNameObj, mIpFabricDnsServiceNameObj)
		mIpFabricDnsServiceName := mIpFabricDnsServiceNameObj.(string)
		object.IpFabricDnsServiceName = mIpFabricDnsServiceName
	}

	mIpFabricServicePortObj := vmap["ip_fabric_service_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mIpFabricServicePortObj) {
		log.Printf("Setting ip_fabric_service_port  IpFabricServicePort <<%T>> => %#v", mIpFabricServicePortObj, mIpFabricServicePortObj)
		mIpFabricServicePort := mIpFabricServicePortObj.(int)
		object.IpFabricServicePort = mIpFabricServicePort
	}

	mIpFabricServiceIpObj := vmap["ip_fabric_service_ip"] // [CPLX; Seq -> 0; 1]
	if CheckTerraformMap(mIpFabricServiceIpObj) {
		log.Printf("Setting ip_fabric_service_ip  IpFabricServiceIp <<%T>> => %#v", mIpFabricServiceIpObj, mIpFabricServiceIpObj)
		for _, v := range mIpFabricServiceIpObj.([]interface{}) {
			mIpFabricServiceIp := v.(string)
			object.AddIpFabricServiceIp(mIpFabricServiceIp)
		}
	}

	log.Printf("FINISHED LinklocalServiceEntryType object: %#v", object)
}

func TakeLinklocalServiceEntryTypeAsMap(object *LinklocalServiceEntryType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["linklocal_service_name"] = object.LinklocalServiceName
	omap["linklocal_service_ip"] = object.LinklocalServiceIp
	omap["linklocal_service_port"] = object.LinklocalServicePort
	omap["ip_fabric_DNS_service_name"] = object.IpFabricDnsServiceName
	omap["ip_fabric_service_port"] = object.IpFabricServicePort
	ip_fabric_service_ip_arr := make([]string, len(object.IpFabricServiceIp))
	for _, v := range object.IpFabricServiceIp {
		ip_fabric_service_ip_arr = append(ip_fabric_service_ip_arr, v)
	}
	omap["ip_fabric_service_ip"] = ip_fabric_service_ip_arr

	return omap
}

func ResourceLinklocalServiceEntryTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"linklocal_service_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"linklocal_service_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"linklocal_service_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"ip_fabric_dns_service_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"ip_fabric_service_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Required: true,
			Type:     schema.TypeInt,
		},
		"ip_fabric_service_ip": &schema.Schema{
			// Cmplx: 0; Seq: True; Type: xsd:string
			Optional: true,
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func ResourceLinklocalServiceEntryType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLinklocalServiceEntryTypeSchema(),
	}
}

func SetLinklocalServicesTypesFromMap(object *LinklocalServicesTypes, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetLinklocalServicesTypesFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mLinklocalServiceEntryObj := vmap["linklocal_service_entry"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mLinklocalServiceEntryObj) {
		log.Printf("Setting linklocal_service_entry  LinklocalServiceEntry <<%T>> => %#v", mLinklocalServiceEntryObj, mLinklocalServiceEntryObj)
		for _, v := range mLinklocalServiceEntryObj.([]interface{}) {
			mLinklocalServiceEntry := new(LinklocalServiceEntryType)
			SetLinklocalServiceEntryTypeFromMap(mLinklocalServiceEntry, d, m, v)
			object.AddLinklocalServiceEntry(mLinklocalServiceEntry)
		}
	}

	log.Printf("FINISHED LinklocalServicesTypes object: %#v", object)
}

func TakeLinklocalServicesTypesAsMap(object *LinklocalServicesTypes) map[string]interface{} {
	omap := make(map[string]interface{})

	var linklocal_service_entry_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.LinklocalServiceEntry {
		linklocal_service_entry_map = append(linklocal_service_entry_map, TakeLinklocalServiceEntryTypeAsMap(&v))
	}
	omap["linklocal_service_entry"] = linklocal_service_entry_map

	return omap
}

func ResourceLinklocalServicesTypesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"linklocal_service_entry": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: LinklocalServiceEntryType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceLinklocalServiceEntryType(),
		},
	}
}

func ResourceLinklocalServicesTypes() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceLinklocalServicesTypesSchema(),
	}
}

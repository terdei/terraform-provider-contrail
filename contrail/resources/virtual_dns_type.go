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

func SetDnsSoaRecordTypeFromMap(object *DnsSoaRecordType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetDnsSoaRecordTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mNegativeCacheTtlSecondsObj := vmap["negative_cache_ttl_seconds"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNegativeCacheTtlSecondsObj) {
		log.Printf("Setting negative_cache_ttl_seconds  NegativeCacheTtlSeconds <<%T>> => %#v", mNegativeCacheTtlSecondsObj, mNegativeCacheTtlSecondsObj)
		mNegativeCacheTtlSeconds := mNegativeCacheTtlSecondsObj.(int)
		object.NegativeCacheTtlSeconds = mNegativeCacheTtlSeconds
	}

	log.Printf("FINISHED DnsSoaRecordType object: %#v", object)
}

func TakeDnsSoaRecordTypeAsMap(object *DnsSoaRecordType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["negative_cache_ttl_seconds"] = object.NegativeCacheTtlSeconds

	return omap
}

func ResourceDnsSoaRecordTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"negative_cache_ttl_seconds": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceDnsSoaRecordType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceDnsSoaRecordTypeSchema(),
	}
}

func SetVirtualDnsTypeFromMap(object *VirtualDnsType, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetVirtualDnsTypeFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mDomainNameObj := vmap["domain_name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDomainNameObj) {
		log.Printf("Setting domain_name  DomainName <<%T>> => %#v", mDomainNameObj, mDomainNameObj)
		mDomainName := mDomainNameObj.(string)
		object.DomainName = mDomainName
	}

	mDynamicRecordsFromClientObj := vmap["dynamic_records_from_client"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDynamicRecordsFromClientObj) {
		log.Printf("Setting dynamic_records_from_client  DynamicRecordsFromClient <<%T>> => %#v", mDynamicRecordsFromClientObj, mDynamicRecordsFromClientObj)
		mDynamicRecordsFromClient := mDynamicRecordsFromClientObj.(bool)
		object.DynamicRecordsFromClient = mDynamicRecordsFromClient
	}

	mRecordOrderObj := vmap["record_order"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRecordOrderObj) {
		log.Printf("Setting record_order  RecordOrder <<%T>> => %#v", mRecordOrderObj, mRecordOrderObj)
		mRecordOrder := mRecordOrderObj.(string)
		object.RecordOrder = mRecordOrder
	}

	mDefaultTtlSecondsObj := vmap["default_ttl_seconds"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mDefaultTtlSecondsObj) {
		log.Printf("Setting default_ttl_seconds  DefaultTtlSeconds <<%T>> => %#v", mDefaultTtlSecondsObj, mDefaultTtlSecondsObj)
		mDefaultTtlSeconds := mDefaultTtlSecondsObj.(int)
		object.DefaultTtlSeconds = mDefaultTtlSeconds
	}

	mNextVirtualDnsObj := vmap["next_virtual_DNS"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNextVirtualDnsObj) {
		log.Printf("Setting next_virtual_DNS  NextVirtualDns <<%T>> => %#v", mNextVirtualDnsObj, mNextVirtualDnsObj)
		mNextVirtualDns := mNextVirtualDnsObj.(string)
		object.NextVirtualDns = mNextVirtualDns
	}

	mFloatingIpRecordObj := vmap["floating_ip_record"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mFloatingIpRecordObj) {
		log.Printf("Setting floating_ip_record  FloatingIpRecord <<%T>> => %#v", mFloatingIpRecordObj, mFloatingIpRecordObj)
		mFloatingIpRecord := mFloatingIpRecordObj.(string)
		object.FloatingIpRecord = mFloatingIpRecord
	}

	mExternalVisibleObj := vmap["external_visible"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mExternalVisibleObj) {
		log.Printf("Setting external_visible  ExternalVisible <<%T>> => %#v", mExternalVisibleObj, mExternalVisibleObj)
		mExternalVisible := mExternalVisibleObj.(bool)
		object.ExternalVisible = mExternalVisible
	}

	mReverseResolutionObj := vmap["reverse_resolution"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mReverseResolutionObj) {
		log.Printf("Setting reverse_resolution  ReverseResolution <<%T>> => %#v", mReverseResolutionObj, mReverseResolutionObj)
		mReverseResolution := mReverseResolutionObj.(bool)
		object.ReverseResolution = mReverseResolution
	}

	mSoaRecordObj := vmap["soa_record"] // [CPLX; Seq -> 1; 0]
	if CheckTerraformMap(mSoaRecordObj) {
		log.Printf("Setting soa_record  SoaRecord <<%T>> => %#v", mSoaRecordObj, mSoaRecordObj)
		mSoaRecord := new(DnsSoaRecordType)
		SetDnsSoaRecordTypeFromMap(mSoaRecord, d, m, mSoaRecordObj)
		object.SoaRecord = mSoaRecord
	}

	log.Printf("FINISHED VirtualDnsType object: %#v", object)
}

func TakeVirtualDnsTypeAsMap(object *VirtualDnsType) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["domain_name"] = object.DomainName
	omap["dynamic_records_from_client"] = object.DynamicRecordsFromClient
	omap["record_order"] = object.RecordOrder
	omap["default_ttl_seconds"] = object.DefaultTtlSeconds
	omap["next_virtual_DNS"] = object.NextVirtualDns
	omap["floating_ip_record"] = object.FloatingIpRecord
	omap["external_visible"] = object.ExternalVisible
	omap["reverse_resolution"] = object.ReverseResolution
	if object.SoaRecord != nil {
		omap["soa_record"] = TakeDnsSoaRecordTypeAsMap(object.SoaRecord)
	}

	return omap
}

func ResourceVirtualDnsTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"domain_name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Required: true,
			Type:     schema.TypeString,
		},
		"dynamic_records_from_client": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"record_order": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"default_ttl_seconds": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
		"next_virtual_dns": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"floating_ip_record": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"external_visible": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"reverse_resolution": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:boolean
			Optional: true,
			Type:     schema.TypeBool,
		},
		"soa_record": &schema.Schema{
			// Cmplx: 1; Seq: False; Type: DnsSoaRecordType
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceDnsSoaRecordType(),
		},
	}
}

func ResourceVirtualDnsType() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceVirtualDnsTypeSchema(),
	}
}

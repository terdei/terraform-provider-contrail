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

func SetTelemetryResourceInfoFromMap(object *TelemetryResourceInfo, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetTelemetryResourceInfoFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mNameObj := vmap["name"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mNameObj) {
		log.Printf("Setting name  Name <<%T>> => %#v", mNameObj, mNameObj)
		mName := mNameObj.(string)
		object.Name = mName
	}

	mPathObj := vmap["path"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPathObj) {
		log.Printf("Setting path  Path <<%T>> => %#v", mPathObj, mPathObj)
		mPath := mPathObj.(string)
		object.Path = mPath
	}

	mRateObj := vmap["rate"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mRateObj) {
		log.Printf("Setting rate  Rate <<%T>> => %#v", mRateObj, mRateObj)
		mRate := mRateObj.(string)
		object.Rate = mRate
	}

	log.Printf("FINISHED TelemetryResourceInfo object: %#v", object)
}

func TakeTelemetryResourceInfoAsMap(object *TelemetryResourceInfo) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["name"] = object.Name
	omap["path"] = object.Path
	omap["rate"] = object.Rate

	return omap
}

func ResourceTelemetryResourceInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"path": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"rate": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceTelemetryResourceInfo() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceTelemetryResourceInfoSchema(),
	}
}

func SetTelemetryStateInfoFromMap(object *TelemetryStateInfo, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetTelemetryStateInfoFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mResourceObj := vmap["resource"] // [CPLX; Seq -> 1; 1]
	if CheckTerraformMap(mResourceObj) {
		log.Printf("Setting resource  Resource <<%T>> => %#v", mResourceObj, mResourceObj)
		for _, v := range mResourceObj.([]interface{}) {
			mResource := new(TelemetryResourceInfo)
			SetTelemetryResourceInfoFromMap(mResource, d, m, v)
			object.AddResource(mResource)
		}
	}

	mServerIpObj := vmap["server_ip"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServerIpObj) {
		log.Printf("Setting server_ip  ServerIp <<%T>> => %#v", mServerIpObj, mServerIpObj)
		mServerIp := mServerIpObj.(string)
		object.ServerIp = mServerIp
	}

	mServerPortObj := vmap["server_port"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mServerPortObj) {
		log.Printf("Setting server_port  ServerPort <<%T>> => %#v", mServerPortObj, mServerPortObj)
		mServerPort := mServerPortObj.(int)
		object.ServerPort = mServerPort
	}

	log.Printf("FINISHED TelemetryStateInfo object: %#v", object)
}

func TakeTelemetryStateInfoAsMap(object *TelemetryStateInfo) map[string]interface{} {
	omap := make(map[string]interface{})

	var resource_map []interface{}
	// COMPLEX SEQUENCE
	for _, v := range object.Resource {
		resource_map = append(resource_map, TakeTelemetryResourceInfoAsMap(&v))
	}
	omap["resource"] = resource_map
	omap["server_ip"] = object.ServerIp
	omap["server_port"] = object.ServerPort

	return omap
}

func ResourceTelemetryStateInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource": &schema.Schema{
			// Cmplx: 1; Seq: True; Type: TelemetryResourceInfo
			Optional: true,
			Type:     schema.TypeList,
			Elem:     ResourceTelemetryResourceInfo(),
		},
		"server_ip": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"server_port": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:integer
			Optional: true,
			Type:     schema.TypeInt,
		},
	}
}

func ResourceTelemetryStateInfo() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceTelemetryStateInfoSchema(),
	}
}

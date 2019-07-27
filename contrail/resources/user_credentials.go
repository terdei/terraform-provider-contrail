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

func SetUserCredentialsFromMap(object *UserCredentials, d *schema.ResourceData, m interface{}, val interface{}) {
	log.Printf("[SetUserCredentialsFromMAP] map = %#v", val)

	vmap, ok := val.(map[string]interface{})
	if ok == false {
		vmap = (val.([]interface{})[0]).(map[string]interface{})
	}

	// SPEW
	log.Printf("SPEW: %v", spew.Sdump(vmap))
	// SPEW

	mUsernameObj := vmap["username"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mUsernameObj) {
		log.Printf("Setting username  Username <<%T>> => %#v", mUsernameObj, mUsernameObj)
		mUsername := mUsernameObj.(string)
		object.Username = mUsername
	}

	mPasswordObj := vmap["password"] // [CPLX; Seq -> 0; 0]
	if CheckTerraformMap(mPasswordObj) {
		log.Printf("Setting password  Password <<%T>> => %#v", mPasswordObj, mPasswordObj)
		mPassword := mPasswordObj.(string)
		object.Password = mPassword
	}

	log.Printf("FINISHED UserCredentials object: %#v", object)
}

func TakeUserCredentialsAsMap(object *UserCredentials) map[string]interface{} {
	omap := make(map[string]interface{})

	omap["username"] = object.Username
	omap["password"] = object.Password

	return omap
}

func ResourceUserCredentialsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
		"password": &schema.Schema{
			// Cmplx: 0; Seq: False; Type: xsd:string
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func ResourceUserCredentials() *schema.Resource {
	return &schema.Resource{
		Schema: ResourceUserCredentialsSchema(),
	}
}

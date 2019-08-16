package resources

import (
	"fmt"
	"github.com/Juniper/contrail-go-api"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var (
	CONTRAIL_API_SERVER = os.Getenv("CONTRAIL_API_SERVER")
	OS_USERNAME         = os.Getenv("OS_USERNAME")
	OS_TENANT_NAME      = os.Getenv("OS_TENANT_NAME")
	OS_PROJECT_NAME     = os.Getenv("OS_PROJECT_NAME")
	OS_PASSWORD         = os.Getenv("OS_PASSWORD")
	OS_AUTH_TOKEN       = os.Getenv("OS_AUTH_TOKEN")
	OS_TOKEN            = os.Getenv("OS_TOKEN")
	OS_AUTH_URL         = os.Getenv("OS_AUTH_URL")
	TF_ACC              = os.Getenv("TF_ACC")
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"contrail": testAccProvider,
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	v := os.Getenv("OS_AUTH_URL")

	if TF_ACC == "" {
		t.Fatal("TF_ACC must be set for acceptance tests")
	}

	if v == "" {
		t.Fatal("OS_AUTH_URL must be set for acceptance tests")
	}

	if OS_PASSWORD == "" && OS_AUTH_TOKEN == "" && OS_TOKEN == "" {
		t.Fatal("OS_PASSWORD or OS_AUTH_TOKEN or OS_TOKEN must be set for acceptance tests")
	}

	if OS_AUTH_URL == "" {
		t.Fatal("OS_AUTH_URL must be set for acceptance tests")
	}

	if OS_PROJECT_NAME == "" && OS_TENANT_NAME == "" {
		t.Fatal("OS_PROJECT_NAME or OS_TENANT_NAME must be set for acceptance tests")
	}

	if OS_USERNAME == "" {
		t.Fatal("OS_USERNAME must be set for acceptance tests")
	}

	if CONTRAIL_API_SERVER == "" {
		t.Fatal("CONTRAIL_API_SERVER must be set for acceptance tests")
	}
}

func testAccCheckDestroy(s *terraform.State, typeName string, clientKey string) error {

	client := testAccProvider.Meta().(*contrail.Client)
	client.GetServer()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != typeName {
			continue
		}

		_, err := client.FindByUuid(clientKey, rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("%s still exists", typeName)
		}
	}

	return nil
}

func testAccCheckExists(key string, clientKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[key]
		if !ok {
			return fmt.Errorf("Not found: %s", key)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*contrail.Client)
		_, err := client.FindByUuid(clientKey, rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error creating %s : %s", clientKey, err)
		}

		return nil
	}
}

// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccContactResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Skip: contact resources are tenant-level billing/mailing objects that may require
	// specific tenant permissions or have undocumented API requirements
	t.Skip("Skipping: contact resource requires special tenant permissions (BAD_REQUEST)")

	rName := acctest.RandomName("tf-acc-test-contact")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_contact.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_contact"),
		Steps: []resource.TestStep{
			{
				Config: testAccContactConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "address1", "123 Test Street"),
					resource.TestCheckResourceAttr(resourceName, "city", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "state", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "zip_code", "98101"),
					resource.TestCheckResourceAttr(resourceName, "country", "US"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccContactImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccContactImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

func testAccContactConfig_basic(nsName, name string) string {
	// Contact resources are tenant-level objects for billing/mailing purposes
	// They need to be created in system namespace
	_ = nsName // unused when using system namespace
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_contact" "test" {
  name       = %[1]q
  namespace  = "system"

  address1     = "123 Test Street"
  city         = "Seattle"
  state        = "Washington"
  state_code   = "WA"
  zip_code     = "98101"
  country      = "US"
  phone_number = "+12065551234"
  contact_type = "MAILING"
}
`, name))
}

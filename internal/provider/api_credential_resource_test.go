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

func TestAccAPICredentialResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: api_credential API endpoint not available in staging environment (NOT_FOUND 404)")

	resourceName := "f5xc_api_credential.test"
	rName := acctest.RandomName("tf-test-cred")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_api_credential"),
		Steps: []resource.TestStep{
			{
				Config: testAccAPICredentialConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "password"},
				ImportStateIdFunc:       testAccAPICredentialImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccAPICredentialImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		name := rs.Primary.Attributes["name"]
		// API credential uses name-only import format (no namespace)
		return name, nil
	}
}

func testAccAPICredentialConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "f5xc_api_credential" "test" {
  name = %[1]q
  type = "API_TOKEN"
}
`, name)
}

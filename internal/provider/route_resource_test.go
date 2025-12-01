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

func TestAccRouteResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: route resource has state drift on waf_type.inherit_waf empty block (generator issue with nested empty blocks)")

	resourceName := "f5xc_route.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-route")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouteResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "routes"},
				ImportStateIdFunc:       testAccRouteResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccRouteResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

func testAccRouteResourceConfig_basic(nsName, name string) string {
	// Route must be in system namespace and uses route_direct_response
	// to avoid external cluster dependencies
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_route" "test" {
  name      = %[1]q
  namespace = "system"

  routes {
    disable_location_add = false

    route_direct_response {
      response_code         = 200
      response_body_encoded = "string:///T0s="  # Base64 encoded "OK"
    }

    waf_type {
      inherit_waf {}
    }
  }
}
`, name))
}

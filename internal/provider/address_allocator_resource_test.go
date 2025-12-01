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

func TestAccAddressAllocatorResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_address_allocator.test"
	rName := acctest.RandomName("tf-test-addr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_address_allocator"),
		Steps: []resource.TestStep{
			{
				Config: testAccAddressAllocatorConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "address_allocation_scheme"},
				ImportStateIdFunc:       testAccAddressAllocatorImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccAddressAllocatorImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccAddressAllocatorConfig_basic(name string) string {
	// Address allocator requires address_allocation_scheme block and address_pool
	// Valid mode values are "LOCAL" or "GLOBAL_PER_SITE_NODE"
	// Must explicitly set all nested block attributes to avoid state drift from API defaults
	return fmt.Sprintf(`
resource "f5xc_address_allocator" "test" {
  name      = %[1]q
  namespace = "system"

  mode = "LOCAL"

  address_pool = ["10.0.0.0/24"]

  address_allocation_scheme {
    allocation_unit                = 0
    local_interface_address_type   = "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN"
    local_interface_address_offset = 0
  }
}
`, name)
}

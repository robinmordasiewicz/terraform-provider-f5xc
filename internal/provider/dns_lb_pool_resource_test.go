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

func TestAccDNSLbPoolResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-pool")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_dns_lb_pool.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_lb_pool"),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSLbPoolConfig_basic(nsName, rName),
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
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccDNSLbPoolImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccDNSLbPoolImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccDNSLbPoolConfig_basic(nsName, name string) string {
	// DNS LB Pool can only be created in system namespace
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_lb_pool" "test" {
  name      = %[1]q
  namespace = "system"

  # TTL for DNS responses (required)
  ttl = 60

  # Load balancing mode (required)
  load_balancing_mode = "ROUND_ROBIN"

  a_pool {
    max_answers = 1

    members {
      name        = "member1"
      ip_endpoint = "192.168.1.10"
      priority    = 0
      ratio       = 0
      disable     = false
    }

    # Health check disabled by default
    disable_health_check {}
  }
}
`, name))
}

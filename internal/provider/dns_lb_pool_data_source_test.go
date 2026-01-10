// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccDnsLbPoolDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-pool")
	nsName := "" // unused, DNS LB Pool must be in system namespace
	resourceName := "f5xc_dns_lb_pool.test"
	dataSourceName := "data.f5xc_dns_lb_pool.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsLbPoolDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccDnsLbPoolDataSourceConfig_basic(nsName, name string) string {
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

data "f5xc_dns_lb_pool" "test" {
  depends_on = [f5xc_dns_lb_pool.test]
  name       = f5xc_dns_lb_pool.test.name
  namespace  = f5xc_dns_lb_pool.test.namespace
}
`, name))
}

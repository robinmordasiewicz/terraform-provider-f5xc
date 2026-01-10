// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccDNSLoadBalancerResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_load_balancer requires tenant info for pool references - add f5xc_tenant data source to enable")

	rName := acctest.RandomName("tf-acc-test-lb")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_dns_load_balancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_load_balancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSLoadBalancerConfig_basic(nsName, rName),
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
				ImportStateIdFunc:       testAccDNSLoadBalancerImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccDNSLoadBalancerImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccDNSLoadBalancerConfig_basic(nsName, name string) string {
	// DNS Load Balancer must be in system namespace
	_ = nsName // unused but kept for test signature consistency
	poolName := name + "-pool"
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
# First create a DNS LB Pool to reference
resource "f5xc_dns_lb_pool" "test" {
  name      = %[1]q
  namespace = "system"
  ttl       = 60
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

    disable_health_check {}
  }
}

resource "f5xc_dns_load_balancer" "test" {
  name      = %[2]q
  namespace = "system"

  record_type = "A"

  rule_list {
    rules {
      score = 100

      ip_prefix_list {
        ip_prefixes = ["0.0.0.0/0"]
      }

      pool {
        name      = f5xc_dns_lb_pool.test.name
        namespace = "system"
      }
    }
  }
}
`, poolName, name))
}

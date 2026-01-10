// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccBgpRoutingPolicyDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-bgprp")
	resourceName := "f5xc_bgp_routing_policy.test"
	dataSourceName := "data.f5xc_bgp_routing_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBgpRoutingPolicyDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccBgpRoutingPolicyDataSourceConfig_basic(name string) string {
	// BGP routing policies should be created in system namespace for networking configuration
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_bgp_routing_policy" "test" {
  name      = %[1]q
  namespace = "system"

  rules {
    match {
      as_path = ".*"
    }
    action {
      allow {}
    }
  }
}

data "f5xc_bgp_routing_policy" "test" {
  depends_on = [f5xc_bgp_routing_policy.test]
  name       = f5xc_bgp_routing_policy.test.name
  namespace  = f5xc_bgp_routing_policy.test.namespace
}
`, name))
}

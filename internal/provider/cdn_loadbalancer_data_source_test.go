// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccCdnLoadbalancerDataSource_basic(t *testing.T) {
	t.Skip("Skipping: CDN loadbalancer requires additional CDN infrastructure and origin pool configuration not available in standard test environments")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-cdn-lb")
	resourceName := "f5xc_cdn_loadbalancer.test"
	dataSourceName := "data.f5xc_cdn_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCdnLoadbalancerDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccCdnLoadbalancerDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_cdn_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  domains = ["%[1]s.example.com"]
}

data "f5xc_cdn_loadbalancer" "test" {
  depends_on = [f5xc_cdn_loadbalancer.test]
  name       = f5xc_cdn_loadbalancer.test.name
  namespace  = f5xc_cdn_loadbalancer.test.namespace
}
`, name))
}

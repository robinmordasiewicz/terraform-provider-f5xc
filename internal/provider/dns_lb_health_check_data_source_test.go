// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccDnsLbHealthCheckDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := "" // unused, DNS LB Health Check must be in system namespace
	resourceName := "f5xc_dns_lb_health_check.test"
	dataSourceName := "data.f5xc_dns_lb_health_check.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsLbHealthCheckDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccDnsLbHealthCheckDataSourceConfig_basic(nsName, name string) string {
	// DNS LB Health Check must be in system namespace
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_lb_health_check" "test" {
  name      = %[1]q
  namespace = "system"

  icmp_health_check {}
}

data "f5xc_dns_lb_health_check" "test" {
  depends_on = [f5xc_dns_lb_health_check.test]
  name       = f5xc_dns_lb_health_check.test.name
  namespace  = f5xc_dns_lb_health_check.test.namespace
}
`, name))
}

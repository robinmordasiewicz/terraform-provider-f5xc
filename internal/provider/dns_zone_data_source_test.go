// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccDnsZoneDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_zone requires domain ownership verification - use a real domain configured in your F5XC tenant")

	// DNS zones require domain name format and the account must own the domain
	rName := acctest.RandomName("tf-acc") + ".example.com"
	nsName := "" // unused, kept for signature consistency
	resourceName := "f5xc_dns_zone.test"
	dataSourceName := "data.f5xc_dns_zone.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsZoneDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccDnsZoneDataSourceConfig_basic(nsName, name string) string {
	// Note: DNS Zone requires domain ownership verification and must use system namespace
	// The nsName parameter is unused but kept for test signature consistency
	_ = nsName
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_zone" "test" {
  name      = %[1]q
  namespace = "system"

  # SOA parameters are required - use default
  default_soa_parameters {}

  # DNSSEC mode with disable
  dnssec_mode {
    disable {}
  }
}

data "f5xc_dns_zone" "test" {
  depends_on = [f5xc_dns_zone.test]
  name       = f5xc_dns_zone.test.name
  namespace  = f5xc_dns_zone.test.namespace
}
`, name))
}

// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccDnsDomainDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	// Generate a unique domain name using random string
	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	nsName := "" // unused, DNS domain must be in system namespace
	resourceName := "f5xc_dns_domain.test"
	dataSourceName := "data.f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsDomainDataSourceConfig_basic(nsName, domainName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccDnsDomainDataSourceConfig_basic(nsName, name string) string {
	// DNS Domain must be in system namespace
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
  name      = %[1]q
  namespace = "system"
}

data "f5xc_dns_domain" "test" {
  depends_on = [f5xc_dns_domain.test]
  name       = f5xc_dns_domain.test.name
  namespace  = f5xc_dns_domain.test.namespace
}
`, name))
}

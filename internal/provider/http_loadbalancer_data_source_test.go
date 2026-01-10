// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccHttpLoadbalancerDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-http-lb")
	resourceName := "f5xc_http_loadbalancer.test"
	dataSourceName := "data.f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHttpLoadbalancerDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccHttpLoadbalancerDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}
}

data "f5xc_http_loadbalancer" "test" {
  depends_on = [f5xc_http_loadbalancer.test]
  name       = f5xc_http_loadbalancer.test.name
  namespace  = f5xc_http_loadbalancer.test.namespace
}
`, name))
}

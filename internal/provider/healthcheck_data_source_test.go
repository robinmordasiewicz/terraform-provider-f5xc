// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test


import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccHealthcheckDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	resourceName := "f5xc_healthcheck.test"
	dataSourceName := "data.f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccHealthcheckDataSourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"
  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5
  tcp_health_check {}
}

data "f5xc_healthcheck" "test" {
  depends_on = [f5xc_healthcheck.test]
  name       = f5xc_healthcheck.test.name
  namespace  = f5xc_healthcheck.test.namespace
}
`, name))
}

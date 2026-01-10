// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccForwardProxyPolicyDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-fpp")
	resourceName := "f5xc_forward_proxy_policy.test"
	dataSourceName := "data.f5xc_forward_proxy_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardProxyPolicyDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}

func testAccForwardProxyPolicyDataSourceConfig_basic(name string) string {
	// Forward proxy policy must be in system namespace
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_forward_proxy_policy" "test" {
  name      = %[1]q
  namespace = "system"
}

data "f5xc_forward_proxy_policy" "test" {
  depends_on = [f5xc_forward_proxy_policy.test]
  name       = f5xc_forward_proxy_policy.test.name
  namespace  = f5xc_forward_proxy_policy.test.namespace
}
`, name))
}

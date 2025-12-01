// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

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

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_http_loadbalancer.test"
	dataSourceName := "data.f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccHttpLoadbalancerDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccHttpLoadbalancerDataSourceConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_origin_pool" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-pool"
  namespace  = f5xc_namespace.test.name
  origin_servers {
    public_ip {
      ip = "1.2.3.4"
    }
  }
  port               = 80
  endpoint_selection = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"
}

resource "f5xc_http_loadbalancer" "test" {
  depends_on = [time_sleep.wait_for_namespace, f5xc_origin_pool.test]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  domains = ["test.example.com"]
  http {
    dns_volterra_managed = false
  }
  default_route_pools {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = f5xc_namespace.test.name
    }
  }
}

data "f5xc_http_loadbalancer" "test" {
  depends_on = [f5xc_http_loadbalancer.test]
  name       = f5xc_http_loadbalancer.test.name
  namespace  = f5xc_http_loadbalancer.test.namespace
}
`, nsName, name))
}

// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test


import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccNetworkFirewallDataSource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_network_firewall.test"
	dataSourceName := "data.f5xc_network_firewall.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkFirewallDataSourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
				),
			},
		},
	})
}


func testAccNetworkFirewallDataSourceConfig_basic(nsName, name string) string {
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

resource "f5xc_forward_proxy_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "${%[2]q}-fpp"
  namespace  = f5xc_namespace.test.name
  proxy_label = "test-proxy"
  rule_list {
    rules {
      metadata {
        name = "rule1"
      }
      spec {
        action      = "ALLOW"
        any_client  = true
        any_dst     = true
        rule_description = "Allow all"
      }
    }
  }
}

resource "f5xc_network_firewall" "test" {
  depends_on = [time_sleep.wait_for_namespace, f5xc_forward_proxy_policy.test]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  active_forward_proxy_policies {
    forward_proxy_policies {
      name      = f5xc_forward_proxy_policy.test.name
      namespace = f5xc_namespace.test.name
    }
  }
}

data "f5xc_network_firewall" "test" {
  depends_on = [f5xc_network_firewall.test]
  name       = f5xc_network_firewall.test.name
  namespace  = f5xc_network_firewall.test.namespace
}
`, nsName, name))
}

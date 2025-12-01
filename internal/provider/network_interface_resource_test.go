// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

func TestAccNetworkInterfaceResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ni")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_network_interface.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_network_interface"),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkInterfaceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "ethernet_interface"},
				ImportStateIdFunc:       testAccNetworkInterfaceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccNetworkInterfaceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

func testAccNetworkInterfaceConfig_basic(nsName, name string) string {
	// Network interface must be in system namespace
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_network_interface" "test" {
  name      = %[1]q
  namespace = "system"

  ethernet_interface {
    device   = "eth0"
    mtu      = 0
    priority = 0
    cluster {}
    untagged {}
    dhcp_client {}
    site_local_inside_network {}
    monitor_disabled {}
    not_primary {}
    no_ipv6_address {}
  }
}
`, name))
}

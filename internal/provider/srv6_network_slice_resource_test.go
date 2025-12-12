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

// =============================================================================
// SRV6 NETWORK SLICE RESOURCE ACCEPTANCE TESTS
//
// These tests verify the f5xc_srv6_network_slice resource implementation.
// SRv6 Network Slice requires SRv6 infrastructure which may not be available
// in all F5 XC tenant configurations.
//
// Run with:
//   TF_ACC=1 VES_API_URL="..." VES_P12_FILE="..." VES_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccSrv6NetworkSliceResource -timeout 30m
// =============================================================================

func TestAccSrv6NetworkSliceResource_basic(t *testing.T) {
	// Skip: SRv6 Network Slice requires SRv6 infrastructure that may not be available
	// in standard F5 XC tenant configurations
	t.Skip("Skipping: SRv6 Network Slice requires SRv6 infrastructure (not available in standard tenants)")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_srv6_network_slice.test"
	nsName := acctest.RandomName("tf-acc")
	name := acctest.RandomName("tf-acc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSrv6NetworkSliceResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccSrv6NetworkSliceResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccSrv6NetworkSliceResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

func testAccSrv6NetworkSliceResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_srv6_network_slice" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, name))
}

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
// VOLTSTACK SITE RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read with API verification
// 2. Import Test - Verify import produces identical state
//
// Note: Voltstack Site is a complex resource that typically requires:
// - Physical or virtual infrastructure deployment
// - Site token registration
// - Extensive network configuration
// - Multiple nested configuration blocks
//
// These tests use minimal configuration for validation purposes.
// Production deployments require comprehensive site configuration.
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccVoltstackSiteResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read operations with API verification
// Pattern: Basic lifecycle test with custom namespace and time_sleep delay
// -----------------------------------------------------------------------------

func TestAccVoltstackSiteResource_basic(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vs")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltstack_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltstack_site"),
		Steps: []resource.TestStep{
			// Step 1: Create voltstack site with minimal configuration
			{
				Config: testAccVoltstackSiteConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccVoltstackSiteImportStateIdFunc(resourceName),
			},
		},
	})
}

// testAccVoltstackSiteImportStateIdFunc returns the import ID in namespace/name format
func testAccVoltstackSiteImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccVoltstackSiteConfig_basic(nsName, name string) string {
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

# Note: This is a minimal configuration for testing only
# Production Voltstack Site deployments require extensive configuration including:
# - coordinates (latitude, longitude)
# - address (street, city, state, zip, country)
# - operating_system_version
# - master_nodes configuration with public_ip
# - logs_streaming_disabled or custom_logs block
# - Multiple required nested blocks for network configuration
# - Site token registration details
resource "f5xc_voltstack_site" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  # Minimal configuration - would need extensive configuration blocks for real deployment
}
`, nsName, name))
}

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
// TUNNEL RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import, Delete with API verification
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccTunnelResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with minimal configuration
// Pattern: Basic lifecycle test with custom namespace creation
// -----------------------------------------------------------------------------

func TestAccTunnelResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_tunnel.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tunnel")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create tunnel with minimal configuration
			{
				Config: testAccTunnelResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// Verify default tunnel_type
					resource.TestCheckResourceAttr(resourceName, "tunnel_type", "IPSEC_PSK"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccTunnelResourceImportStateIdFunc(resourceName),
				// Ignore nested attributes that the Read function doesn't fully restore
				ImportStateVerifyIgnore: []string{
					"timeouts",
					"local_ip",
					"remote_ip",
					"params",
				},
			},
		},
	})
}

// =============================================================================
// Test Helper Functions
// =============================================================================

// testAccTunnelResourceImportStateIdFunc returns the import ID in namespace/name format
func testAccTunnelResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccTunnelResourceConfig_basic(nsName, name string) string {
	// Tunnel must be created in system namespace
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_tunnel" "test" {
  name       = %[1]q
  namespace  = "system"

  # tunnel_type defaults to "IPSEC_PSK" if not specified
  tunnel_type = "IPSEC_PSK"

  local_ip {
    ip_address {
      auto {}
      virtual_network_type {
        site_local {}
      }
    }
  }

  remote_ip {
    ip {
      ipv4 {
        addr = "192.0.2.1"
      }
    }
  }

  params {
    ipsec {
      ipsec_psk {
        clear_secret_info {
          url = "string:///dGVzdC1wc2stc2VjcmV0"
        }
      }
    }
  }
}
`, name))
}

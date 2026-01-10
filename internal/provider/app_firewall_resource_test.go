// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// TEST: Basic app_firewall creation with default detection settings
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccAppFirewallResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckAppFirewallDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAppFirewallExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Import test
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccAppFirewallImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: App firewall with labels and description
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccAppFirewallResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckAppFirewallDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_withLabelsSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAppFirewallExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test application firewall"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "security"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: App firewall with blocking mode
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccAppFirewallResource_blocking(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckAppFirewallDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_blockingSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAppFirewallExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: App firewall with monitoring mode
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccAppFirewallResource_monitoring(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckAppFirewallDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_monitoringSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAppFirewallExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccAppFirewallResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckAppFirewallDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAppFirewallExists(resourceName),
					acctest.CheckAppFirewallDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccAppFirewallResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckAppFirewallDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAppFirewallExists(resourceName),
				),
			},
			{
				Config:             testAccAppFirewallConfig_basicSystem(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccAppFirewallImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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
// CONFIG HELPERS
// =============================================================================

// testAccAppFirewallConfig_basicSystem uses the "system" namespace
// to avoid creating test namespaces (namespace DELETE returns 501)
func testAccAppFirewallConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  name       = %[1]q
  namespace  = "system"

  # Use default detection settings for simplicity
  default_detection_settings {}

  # Allow all response codes
  allow_all_response_codes {}

  # Blocking mode
  blocking {}

  # Use default blocking page
  use_default_blocking_page {}

  # Use default bot settings
  default_bot_setting {}

  # Use default anonymization
  default_anonymization {}
}
`, name)
}

func testAccAppFirewallConfig_withLabelsSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test application firewall"

  labels = {
    environment = "test"
    team        = "security"
  }

  # Use default detection settings
  default_detection_settings {}

  # Allow all response codes
  allow_all_response_codes {}

  # Blocking mode
  blocking {}

  # Use default blocking page
  use_default_blocking_page {}

  # Use default bot settings
  default_bot_setting {}

  # Use default anonymization
  default_anonymization {}
}
`, name)
}

func testAccAppFirewallConfig_blockingSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  name       = %[1]q
  namespace  = "system"

  # Use default detection settings
  default_detection_settings {}

  # Allow all response codes
  allow_all_response_codes {}

  # Blocking mode - actively block malicious requests
  blocking {}

  # Use default blocking page
  use_default_blocking_page {}

  # Use default bot settings
  default_bot_setting {}

  # Use default anonymization
  default_anonymization {}
}
`, name)
}

func testAccAppFirewallConfig_monitoringSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  name       = %[1]q
  namespace  = "system"

  # Use default detection settings
  default_detection_settings {}

  # Allow all response codes
  allow_all_response_codes {}

  # Monitoring mode - log but don't block
  monitoring {}

  # Use default blocking page
  use_default_blocking_page {}

  # Use default bot settings
  default_bot_setting {}

  # Use default anonymization
  default_anonymization {}
}
`, name)
}

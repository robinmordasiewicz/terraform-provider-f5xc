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
// TEST: Basic app_firewall creation with default detection settings
// =============================================================================
func TestAccAppFirewallResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_firewall"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
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
// =============================================================================
func TestAccAppFirewallResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_firewall"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_withLabels(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
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
// =============================================================================
func TestAccAppFirewallResource_blocking(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_firewall"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_blocking(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: App firewall with monitoring mode
// =============================================================================
func TestAccAppFirewallResource_monitoring(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_firewall"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_monitoring(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// =============================================================================
func TestAccAppFirewallResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_firewall"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					acctest.CheckAppFirewallDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// =============================================================================
func TestAccAppFirewallResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_app_firewall.test"
	rName := acctest.RandomName("tf-test-waf")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_app_firewall"),
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			{
				Config:             testAccAppFirewallConfig_basic(nsName, rName),
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

// testAccAppFirewallConfig_namespaceBase returns the namespace configuration
func testAccAppFirewallConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

# Wait for namespace to be ready before creating app_firewall
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccAppFirewallConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccAppFirewallConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

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
`, name))
}

func testAccAppFirewallConfig_withLabels(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccAppFirewallConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
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
`, name))
}

func testAccAppFirewallConfig_blocking(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccAppFirewallConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

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
`, name))
}

func testAccAppFirewallConfig_monitoring(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccAppFirewallConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_app_firewall" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

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
`, name))
}

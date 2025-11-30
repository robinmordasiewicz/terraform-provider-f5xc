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
// TEST: Basic service_policy creation with allow_all_requests
// =============================================================================
func TestAccServicePolicyResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
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
				ImportStateIdFunc:       testAccServicePolicyImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: Service policy with labels and description
// =============================================================================
func TestAccServicePolicyResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_withLabels(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test service policy"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "security"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Service policy with deny_all_requests
// =============================================================================
func TestAccServicePolicyResource_denyAll(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_denyAll(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Service policy with allow_list
// =============================================================================
func TestAccServicePolicyResource_allowList(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_allowList(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// =============================================================================
func TestAccServicePolicyResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
					acctest.CheckServicePolicyDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// =============================================================================
func TestAccServicePolicyResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
				),
			},
			{
				Config:             testAccServicePolicyConfig_basic(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccServicePolicyImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// testAccServicePolicyConfig_namespaceBase returns the namespace configuration
func testAccServicePolicyConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

# Wait for namespace to be ready before creating service_policy
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccServicePolicyConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccServicePolicyConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  # Allow all requests - simplest policy
  allow_all_requests {}

  # Apply to any server
  any_server {}
}
`, name))
}

func testAccServicePolicyConfig_withLabels(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccServicePolicyConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = "Test service policy"

  labels = {
    environment = "test"
    team        = "security"
  }

  # Allow all requests
  allow_all_requests {}

  # Apply to any server
  any_server {}
}
`, name))
}

func testAccServicePolicyConfig_denyAll(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccServicePolicyConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  # Deny all requests
  deny_all_requests {}

  # Apply to any server
  any_server {}
}
`, name))
}

func testAccServicePolicyConfig_allowList(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccServicePolicyConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  # Allow list with IP prefix
  allow_list {
    prefix_list {
      prefixes = ["10.0.0.0/8", "192.168.0.0/16"]
    }
    default_action_deny {}
  }

  # Apply to any server
  any_server {}
}
`, name))
}

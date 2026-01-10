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
// TEST: Basic service_policy creation with allow_all_requests
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccServicePolicyResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
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
				ImportStateIdFunc:       testAccServicePolicyImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: Service policy with labels and description
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccServicePolicyResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_withLabelsSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccServicePolicyResource_denyAll(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_denyAllSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccServicePolicyResource_allowList(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_allowListSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccServicePolicyResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_basicSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccServicePolicyResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy.test"
	rName := acctest.RandomName("tf-test-sp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckServicePolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckServicePolicyExists(resourceName),
				),
			},
			{
				Config:             testAccServicePolicyConfig_basicSystem(rName),
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
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccServicePolicyConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  name       = %[1]q
  namespace  = "system"

  # Allow all requests - simplest policy
  allow_all_requests {}

  # Apply to any server
  any_server {}
}
`, name)
}

func testAccServicePolicyConfig_withLabelsSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  name        = %[1]q
  namespace   = "system"
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
`, name)
}

func testAccServicePolicyConfig_denyAllSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  name       = %[1]q
  namespace  = "system"

  # Deny all requests
  deny_all_requests {}

  # Apply to any server
  any_server {}
}
`, name)
}

func testAccServicePolicyConfig_allowListSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_service_policy" "test" {
  name       = %[1]q
  namespace  = "system"

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
`, name)
}

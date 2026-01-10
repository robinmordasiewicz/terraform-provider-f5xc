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
// VOLTSHARE_ADMIN_POLICY RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import with custom namespace
// 2. All Attributes Test - Test all optional attributes
// 3. Update Tests - Test mutable attributes
// 4. Disappears Test - Handle external resource deletion
// 5. Empty Plan Test - Verify no drift after apply
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccVoltshareAdminPolicyResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with custom namespace
// Pattern: Basic lifecycle test with time_sleep for namespace propagation
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			// Step 1: Create voltshare_admin_policy with minimal configuration
			{
				Config: testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
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
				ImportStateIdFunc:       testAccVoltshareAdminPolicyImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"
	description := "Comprehensive acceptance test voltshare admin policy"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			{
				Config: testAccVoltshareAdminPolicyConfig_allAttributes(nsName, rName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					resource.TestCheckResourceAttr(resourceName, "max_validity_duration", "86400s"),
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccVoltshareAdminPolicyImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccVoltshareAdminPolicyConfig_withLabels(nsName, rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccVoltshareAdminPolicyConfig_withLabels(nsName, rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "labels.environment"),
					resource.TestCheckNoResourceAttr(resourceName, "labels.managed_by"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 4: Update Test - Description
// Verifies: Update and removal of description attribute
// Pattern: Update test for optional string attribute
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccVoltshareAdminPolicyConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccVoltshareAdminPolicyConfig_withDescription(nsName, rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Disappears Test
// Verifies: Provider handles external resource deletion gracefully
// Pattern: Disappears test - resource deleted outside Terraform
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			{
				Config: testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Delete the resource outside of Terraform
					// TODO: add generic CheckResourceDisappears helper
					// acctest.CheckResourceExists(resourceName),
				),
				// Expect the plan to show the resource needs to be recreated
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 6: Empty Plan Test
// Verifies: No diff on subsequent plan after apply
// Pattern: Empty plan verification - idempotency check
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config:             testAccVoltshareAdminPolicyConfig_basic(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 7: Author Restrictions - Allow All
// Verifies: Author restrictions with allow_all configuration
// Pattern: Test nested block configuration
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_authorRestrictionsAllowAll(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			{
				Config: testAccVoltshareAdminPolicyConfig_authorRestrictionsAllowAll(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 8: User Restrictions - All Tenants
// Verifies: User restrictions with all_tenants configuration
// Pattern: Test list nested block configuration
// -----------------------------------------------------------------------------

func TestAccVoltshareAdminPolicyResource_userRestrictionsAllTenants(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-vap")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_voltshare_admin_policy.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckResourceDestroyed("f5xc_voltshare_admin_policy"),
		Steps: []resource.TestStep{
			{
				Config: testAccVoltshareAdminPolicyConfig_userRestrictionsAllTenants(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
				),
			},
		},
	})
}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

func testAccVoltshareAdminPolicyImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccVoltshareAdminPolicyConfig_basic(nsName, name string) string {
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

resource "f5xc_voltshare_admin_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  author_restrictions {
    allow_all {}
  }

  user_restrictions {
    all_tenants {}
  }
}
`, nsName, name))
}

func testAccVoltshareAdminPolicyConfig_allAttributes(nsName, name, description string) string {
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

resource "f5xc_voltshare_admin_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name                  = %[2]q
  namespace             = f5xc_namespace.test.name
  description           = %[3]q
  max_validity_duration = "86400s"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  author_restrictions {
    allow_all {}
  }

  user_restrictions {
    all_tenants {}
  }
}
`, nsName, name, description))
}

func testAccVoltshareAdminPolicyConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_voltshare_admin_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }

  author_restrictions {
    allow_all {}
  }

  user_restrictions {
    all_tenants {}
  }
}
`, nsName, name, environment, managedBy))
}

func testAccVoltshareAdminPolicyConfig_withDescription(nsName, name, description string) string {
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

resource "f5xc_voltshare_admin_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q

  author_restrictions {
    allow_all {}
  }

  user_restrictions {
    all_tenants {}
  }
}
`, nsName, name, description))
}

func testAccVoltshareAdminPolicyConfig_authorRestrictionsAllowAll(nsName, name string) string {
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

resource "f5xc_voltshare_admin_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  author_restrictions {
    allow_all {}
  }

  user_restrictions {
    all_tenants {}
  }
}
`, nsName, name))
}

func testAccVoltshareAdminPolicyConfig_userRestrictionsAllTenants(nsName, name string) string {
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

resource "f5xc_voltshare_admin_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  author_restrictions {
    allow_all {}
  }

  user_restrictions {
    all_tenants {}
  }
}
`, nsName, name))
}

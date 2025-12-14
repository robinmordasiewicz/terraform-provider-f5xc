// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// SECUREMESH SITE RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Update, Delete with API verification
// 2. Import Test - Verify import produces identical state
// 3. Update Tests - Test all mutable attributes
// 4. All Attributes Test - Test all optional attributes
// 5. Empty Plan Test - Verify no diff after apply
// 6. Plan Checks - Verify planned actions
// 7. Known Values - Verify planned values match expected
// 8. Error/Validation Tests - Invalid configuration handling
// 9. Requires Replace Test - Name/namespace change forces replacement
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccSecuremeshSiteResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test with CheckDestroy
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_basic(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create securemesh_site with minimal configuration
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName, rName),
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
				ImportStateIdFunc:       testAccSecuremeshSiteResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, etc.)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_allAttributes(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"
	description := "Comprehensive acceptance test securemesh site"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSecuremeshSiteResourceConfig_allAttributes(nsName, rName, description),
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
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccSecuremeshSiteResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_updateLabels(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccSecuremeshSiteResourceConfig_withLabels(nsName, rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccSecuremeshSiteResourceConfig_withLabels(nsName, rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName, rName),
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

func TestAccSecuremeshSiteResource_updateDescription(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccSecuremeshSiteResourceConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccSecuremeshSiteResourceConfig_withDescription(nsName, rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Update Test - Annotations
// Verifies: Update and removal of annotations attribute
// Pattern: Update test for map attribute
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_updateAnnotations(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with annotations
			{
				Config: testAccSecuremeshSiteResourceConfig_withAnnotations(nsName, rName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
			// Step 2: Update annotations
			{
				Config: testAccSecuremeshSiteResourceConfig_withAnnotations(nsName, rName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			// Step 3: Remove annotations
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 6: Empty Plan Test
// Verifies: No diff on subsequent plan after apply
// Pattern: Empty plan verification - idempotency check
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_emptyPlan(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccSecuremeshSiteResourceConfig_allAttributes(nsName, rName, "Empty plan test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config: testAccSecuremeshSiteResourceConfig_allAttributes(nsName, rName, "Empty plan test"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 7: Plan Checks Test
// Verifies: Planned actions are correct (create, update, no-op)
// Pattern: Plan check validation
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_planChecks(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create - verify create action planned
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName, rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			// Step 2: Update - verify update action planned
			{
				Config: testAccSecuremeshSiteResourceConfig_withDescription(nsName, rName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			// Step 3: No change - verify no-op planned
			{
				Config: testAccSecuremeshSiteResourceConfig_withDescription(nsName, rName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 8: Known Values Plan Check
// Verifies: Planned values match expected values
// Pattern: ExpectKnownValue plan check
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_knownValues(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	rName := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSecuremeshSiteResourceConfig_withDescription(nsName, rName, "Known value test"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(rName),
						),
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("namespace"),
							knownvalue.StringExact(nsName),
						),
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("description"),
							knownvalue.StringExact("Known value test"),
						),
					},
				},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 9: Error Test - Invalid Name
// Verifies: Invalid configurations are rejected with appropriate errors
// Pattern: ExpectError test
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				// Name with invalid characters (uppercase)
				Config:      testAccSecuremeshSiteResourceConfig_basic(nsName, "Invalid-NAME-Test"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|must)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 10: Error Test - Name Too Long
// Verifies: Name length validation
// Pattern: ExpectError test for validation
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	// Create a name that exceeds the maximum length (typically 63 characters for K8s-style names)
	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccSecuremeshSiteResourceConfig_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|length|long|exceed|character)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 11: Error Test - Empty Name
// Verifies: Empty name is rejected
// Pattern: ExpectError test
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccSecuremeshSiteResourceConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|empty|required|blank)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 12: Requires Replace Test
// Verifies: Changing name or namespace forces resource replacement
// Pattern: ForceNew/RequiresReplace verification
// -----------------------------------------------------------------------------

func TestAccSecuremeshSiteResource_requiresReplace(t *testing.T) {
	t.Skip("Skipping: requires physical site infrastructure and registration token")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName1 := acctest.RandomName("tf-acc-test-ns")
	nsName2 := acctest.RandomName("tf-acc-test-ns")
	rName1 := acctest.RandomName("tf-acc-test-smsite")
	rName2 := acctest.RandomName("tf-acc-test-smsite")
	resourceName := "f5xc_securemesh_site.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_securemesh_site"),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName1, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName1, rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName2),
				),
			},
			// Step 3: Change namespace - should force replacement
			{
				Config: testAccSecuremeshSiteResourceConfig_basic(nsName2, rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName2),
				),
			},
		},
	})
}

// =============================================================================
// Test Helper Functions
// =============================================================================

func testAccSecuremeshSiteResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccSecuremeshSiteResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_securemesh_site" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  # Minimal required configuration - only name and namespace are required
  volterra_certified_hw = "generic-single-nic-volstack-combo"

  # Use master_node_configuration for node details (correct schema block)
  master_node_configuration {
    name = "master-0"
  }
}
`, nsName, name))
}

func testAccSecuremeshSiteResourceConfig_allAttributes(nsName, name, description string) string {
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

resource "f5xc_securemesh_site" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  volterra_certified_hw = "generic-single-nic-volstack-combo"

  master_node_configuration {
    name = "master-0"
  }
}
`, nsName, name, description))
}

func testAccSecuremeshSiteResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_securemesh_site" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }

  volterra_certified_hw = "generic-single-nic-volstack-combo"

  master_node_configuration {
    name = "master-0"
  }
}
`, nsName, name, environment, managedBy))
}

func testAccSecuremeshSiteResourceConfig_withDescription(nsName, name, description string) string {
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

resource "f5xc_securemesh_site" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q

  volterra_certified_hw = "generic-single-nic-volstack-combo"

  master_node_configuration {
    name = "master-0"
  }
}
`, nsName, name, description))
}

func testAccSecuremeshSiteResourceConfig_withAnnotations(nsName, name, value1, value2 string) string {
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

resource "f5xc_securemesh_site" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  annotations = {
    key1 = %[3]q
    key2 = %[4]q
  }

  volterra_certified_hw = "generic-single-nic-volstack-combo"

  master_node_configuration {
    name = "master-0"
  }
}
`, nsName, name, value1, value2))
}

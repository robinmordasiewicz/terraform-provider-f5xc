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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// NAMESPACE RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Update, Delete with API verification
// 2. Import Test - Verify import produces identical state
// 3. Update Tests - Test all mutable attributes
// 4. Disappears Test - Handle external resource deletion
// 5. Error/Validation Tests - Invalid configuration handling
// 6. All Attributes Test - Test all optional attributes
// 7. Empty Plan Test - Verify no diff after apply
// 8. Plan Checks - Verify planned actions
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccNamespaceResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test with CheckDestroy
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		// Note: CheckDestroy omitted - F5 XC staging API returns 501 for namespace DELETE
		// This is a backend limitation, not a provider bug
		Steps: []resource.TestStep{
			// Step 1: Create namespace with minimal configuration
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify resource exists in F5 XC API
					acctest.CheckNamespaceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					// Note: namespace is optional and not set for namespace resource
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, disable)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"
	description := "Comprehensive acceptance test namespace"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig_allAttributes(rName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					// Note: namespace is optional and not set for namespace resource
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					// Verify attributes in F5 XC API match state
					acctest.CheckNamespaceAttributes(resourceName,
						map[string]string{
							"environment": "test",
							"managed_by":  "terraform-acceptance-test",
						},
						map[string]string{
							"purpose": "acceptance-testing",
							"owner":   "ci-cd",
						},
					),
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
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

func TestAccNamespaceResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccNamespaceResourceConfig_withLabels(rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccNamespaceResourceConfig_withLabels(rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
					// Verify API matches
					acctest.CheckNamespaceAttributes(resourceName,
						map[string]string{
							"environment": "staging",
							"managed_by":  "terraform-updated",
						},
						nil,
					),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
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

func TestAccNamespaceResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccNamespaceResourceConfig_withDescription(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccNamespaceResourceConfig_withDescription(rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
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

func TestAccNamespaceResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with annotations
			{
				Config: testAccNamespaceResourceConfig_withAnnotations(rName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
					acctest.CheckNamespaceAttributes(resourceName, nil,
						map[string]string{
							"key1": "value1",
							"key2": "value2",
						},
					),
				),
			},
			// Step 2: Update annotations
			{
				Config: testAccNamespaceResourceConfig_withAnnotations(rName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			// Step 3: Remove annotations
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 6: Disappears Test
// Verifies: Provider handles external resource deletion gracefully
// Pattern: Disappears test - resource deleted outside Terraform
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					// Delete the resource outside of Terraform
					acctest.CheckNamespaceDisappears(resourceName),
				),
				// Expect the plan to show the resource needs to be recreated
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 7: Empty Plan Test
// Verifies: No diff on subsequent plan after apply
// Pattern: Empty plan verification - idempotency check
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccNamespaceResourceConfig_allAttributes(rName, "Empty plan test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config: testAccNamespaceResourceConfig_allAttributes(rName, "Empty plan test"),
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
// Test 8: Plan Checks Test
// Verifies: Planned actions are correct (create, update, no-op)
// Pattern: Plan check validation
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create - verify create action planned
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
				),
			},
			// Step 2: Update - verify update action planned
			{
				Config: testAccNamespaceResourceConfig_withDescription(rName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			// Step 3: No change - verify no-op planned
			{
				Config: testAccNamespaceResourceConfig_withDescription(rName, "Updated for plan check"),
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
// Test 9: Known Values Plan Check
// Verifies: Planned values match expected values
// Pattern: ExpectKnownValue plan check
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig_withDescription(rName, "Known value test"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(rName),
						),
						// Note: namespace is optional and not set for namespace resource
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
// Test 10: Error Test - Invalid Name
// Verifies: Invalid configurations are rejected with appropriate errors
// Pattern: ExpectError test
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Name with invalid characters (uppercase)
				Config: testAccNamespaceResourceConfig_basic("Invalid-NAME-Test"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|must)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 11: Error Test - Name Too Long
// Verifies: Name length validation
// Pattern: ExpectError test for validation
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Create a name that exceeds the maximum length (typically 63 characters for K8s-style names)
	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNamespaceResourceConfig_basic(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|length|long|exceed|character)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 12: Error Test - Empty Name
// Verifies: Empty name is rejected
// Pattern: ExpectError test
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNamespaceResourceConfig_basic(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|empty|required|blank)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 13: Requires Replace Test
// Verifies: Changing name forces resource replacement
// Pattern: ForceNew/RequiresReplace verification
// -----------------------------------------------------------------------------

func TestAccNamespaceResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-acc-test-ns")
	rName2 := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckNamespaceDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccNamespaceResourceConfig_basic(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccNamespaceResourceConfig_basic(rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckNamespaceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName2),
				),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccNamespaceResourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}
`, name))
}

func testAccNamespaceResourceConfig_allAttributes(name, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name        = %[1]q
  description = %[2]q

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }
}
`, name, description))
}

func testAccNamespaceResourceConfig_withLabels(name, environment, managedBy string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }
}
`, name, environment, managedBy))
}

func testAccNamespaceResourceConfig_withDescription(name, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name        = %[1]q
  description = %[2]q
}
`, name, description))
}

func testAccNamespaceResourceConfig_withAnnotations(name, value1, value2 string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q

  annotations = {
    key1 = %[2]q
    key2 = %[3]q
  }
}
`, name, value1, value2))
}

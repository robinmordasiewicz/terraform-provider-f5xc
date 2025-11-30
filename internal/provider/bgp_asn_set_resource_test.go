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
// BGP_ASN_SET RESOURCE ACCEPTANCE TESTS
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
// 9. AS Numbers Test - Test as_numbers list attribute
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccBGPAsnSetResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test with CheckDestroy
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create BGP ASN set with minimal configuration
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify resource exists in F5 XC API
					acctest.CheckBGPAsnSetExists(resourceName),
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
				ImportStateIdFunc:       testAccBGPAsnSetImportStateIdFunc(resourceName),
			},
		},
	})
}

// testAccBGPAsnSetImportStateIdFunc returns a function that generates the import ID
// for a BGP ASN set resource in the format "namespace/name"
func testAccBGPAsnSetImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]

		if namespace == "" || name == "" {
			return "", fmt.Errorf("namespace or name not set in state")
		}

		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, as_numbers)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.#", "2"),
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable", "description"},
				ImportStateIdFunc:       testAccBGPAsnSetImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccBGPAsnSetResourceConfig_withLabels(nsName, rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccBGPAsnSetResourceConfig_withLabels(nsName, rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
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

func TestAccBGPAsnSetResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccBGPAsnSetResourceConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccBGPAsnSetResourceConfig_withDescription(nsName, rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
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

func TestAccBGPAsnSetResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with annotations
			{
				Config: testAccBGPAsnSetResourceConfig_withAnnotations(nsName, rName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
			// Step 2: Update annotations
			{
				Config: testAccBGPAsnSetResourceConfig_withAnnotations(nsName, rName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			// Step 3: Remove annotations
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
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

func TestAccBGPAsnSetResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					// Delete the resource outside of Terraform
					acctest.CheckBGPAsnSetDisappears(resourceName),
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

func TestAccBGPAsnSetResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccBGPAsnSetResourceConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config: testAccBGPAsnSetResourceConfig_allAttributes(nsName, rName),
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

func TestAccBGPAsnSetResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create - verify create action planned
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
				),
			},
			// Step 2: Update - verify update action planned
			{
				Config: testAccBGPAsnSetResourceConfig_withLabels(nsName, rName, "test", "terraform"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			// Step 3: No change - verify no-op planned
			{
				Config: testAccBGPAsnSetResourceConfig_withLabels(nsName, rName, "test", "terraform"),
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

func TestAccBGPAsnSetResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName),
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

func TestAccBGPAsnSetResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			{
				// Name with invalid characters (uppercase)
				Config:      testAccBGPAsnSetResourceConfig_basic(nsName, "Invalid-NAME-Test"),
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

func TestAccBGPAsnSetResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	// Create a name that exceeds the maximum length (typically 63 characters for K8s-style names)
	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccBGPAsnSetResourceConfig_basic(nsName, longName),
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

func TestAccBGPAsnSetResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccBGPAsnSetResourceConfig_basic(nsName, ""),
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

func TestAccBGPAsnSetResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-acc-test-asn")
	rName2 := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccBGPAsnSetResourceConfig_basic(nsName, rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName2),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 14: AS Numbers Test
// Verifies: AS numbers list attribute configuration
// Pattern: List attribute test
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_asNumbers(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_withASNumbers(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.0", "64512"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.1", "64513"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.2", "64514"),
				),
			},
			// Import verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccBGPAsnSetImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccBGPAsnSetResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_bgp_asn_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  as_numbers = ["64512"]
}
`, nsName, name))
}

func testAccBGPAsnSetResourceConfig_allAttributes(nsName, name string) string {
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

resource "f5xc_bgp_asn_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  as_numbers = ["64512", "64513"]
}
`, nsName, name))
}

func testAccBGPAsnSetResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_bgp_asn_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }
}
`, nsName, name, environment, managedBy))
}

func testAccBGPAsnSetResourceConfig_withDescription(nsName, name, description string) string {
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

resource "f5xc_bgp_asn_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, name, description))
}

func testAccBGPAsnSetResourceConfig_withAnnotations(nsName, name, value1, value2 string) string {
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

resource "f5xc_bgp_asn_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  annotations = {
    key1 = %[3]q
    key2 = %[4]q
  }
}
`, nsName, name, value1, value2))
}

func testAccBGPAsnSetResourceConfig_withASNumbers(nsName, name string) string {
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

resource "f5xc_bgp_asn_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  as_numbers = ["64512", "64513", "64514"]
}
`, nsName, name))
}

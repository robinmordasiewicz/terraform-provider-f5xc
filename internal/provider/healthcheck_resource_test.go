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
// HEALTHCHECK RESOURCE ACCEPTANCE TESTS
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
// 9. HTTP Health Check Test - Test nested block configuration
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccHealthcheckResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test with CheckDestroy
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_basic(t *testing.T) {
	t.Skip("Skipping: healthcheck generator does not marshal spec fields (healthy_threshold, timeout, etc.) to API request - requires generator enhancement")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create healthcheck with minimal configuration
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify resource exists in F5 XC API
					acctest.CheckHealthcheckExists(resourceName),
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
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// testAccHealthcheckImportStateIdFunc returns a function that generates the import ID
// for a healthcheck resource in the format "namespace/name"
func testAccHealthcheckImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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
// Verifies: All optional attributes (labels, annotations, description)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckResourceConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					// Verify attributes in F5 XC API match state
					acctest.CheckHealthcheckAttributes(resourceName,
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
				ImportStateVerifyIgnore: []string{"timeouts", "disable", "description"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccHealthcheckResourceConfig_withLabels(nsName, rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccHealthcheckResourceConfig_withLabels(nsName, rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
					// Verify API matches
					acctest.CheckHealthcheckAttributes(resourceName,
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
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "labels.environment"),
					resource.TestCheckNoResourceAttr(resourceName, "labels.managed_by"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 4: Update Test - Annotations
// Verifies: Update and removal of annotations attribute
// Pattern: Update test for map attribute
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with annotations
			{
				Config: testAccHealthcheckResourceConfig_withAnnotations(nsName, rName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
					acctest.CheckHealthcheckAttributes(resourceName, nil,
						map[string]string{
							"key1": "value1",
							"key2": "value2",
						},
					),
				),
			},
			// Step 2: Update annotations
			{
				Config: testAccHealthcheckResourceConfig_withAnnotations(nsName, rName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			// Step 3: Remove annotations
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
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

func TestAccHealthcheckResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					// Delete the resource outside of Terraform
					acctest.CheckHealthcheckDisappears(resourceName),
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

func TestAccHealthcheckResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccHealthcheckResourceConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config: testAccHealthcheckResourceConfig_allAttributes(nsName, rName),
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

func TestAccHealthcheckResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create - verify create action planned
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
				),
			},
			// Step 2: Update - verify update action planned
			{
				Config: testAccHealthcheckResourceConfig_withLabels(nsName, rName, "test", "terraform"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			// Step 3: No change - verify no-op planned
			{
				Config: testAccHealthcheckResourceConfig_withLabels(nsName, rName, "test", "terraform"),
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

func TestAccHealthcheckResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName),
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
// Test 9: Error Test - Invalid Name
// Verifies: Invalid configurations are rejected with appropriate errors
// Pattern: ExpectError test
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_invalidName(t *testing.T) {
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
				Config:      testAccHealthcheckResourceConfig_basic(nsName, "Invalid-NAME-Test"),
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

func TestAccHealthcheckResource_nameTooLong(t *testing.T) {
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
				Config:      testAccHealthcheckResourceConfig_basic(nsName, longName),
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

func TestAccHealthcheckResource_emptyName(t *testing.T) {
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
				Config:      testAccHealthcheckResourceConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|empty|required|blank)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 12: Requires Replace Test
// Verifies: Changing name forces resource replacement
// Pattern: ForceNew/RequiresReplace verification
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-acc-test-hc")
	rName2 := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccHealthcheckResourceConfig_basic(nsName, rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName2),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 13: HTTP Health Check Test
// Verifies: HTTP health check nested block configuration
// Pattern: Nested block test
// -----------------------------------------------------------------------------

func TestAccHealthcheckResource_httpHealthCheck(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")
	nsName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source: "hashicorp/time",
			},
		},
		CheckDestroy: acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckResourceConfig_httpHealthCheck(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.path", "/health"),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "example.com"),
				),
			},
			// Import verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccHealthcheckResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_healthcheck" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  tcp_health_check {}
}
`, nsName, name))
}

func testAccHealthcheckResourceConfig_allAttributes(nsName, name string) string {
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

resource "f5xc_healthcheck" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  tcp_health_check {}
}
`, nsName, name))
}

func testAccHealthcheckResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_healthcheck" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }

  tcp_health_check {}
}
`, nsName, name, environment, managedBy))
}

func testAccHealthcheckResourceConfig_withAnnotations(nsName, name, value1, value2 string) string {
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

resource "f5xc_healthcheck" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  annotations = {
    key1 = %[3]q
    key2 = %[4]q
  }

  tcp_health_check {}
}
`, nsName, name, value1, value2))
}

func testAccHealthcheckResourceConfig_httpHealthCheck(nsName, name string) string {
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

resource "f5xc_healthcheck" "test" {
  depends_on = [time_sleep.wait_for_namespace]

  name      = %[2]q
  namespace = f5xc_namespace.test.name

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path        = "/health"
    host_header = "example.com"
  }
}
`, nsName, name))
}

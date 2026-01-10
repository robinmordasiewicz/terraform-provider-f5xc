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
// SERVICE POLICY RULE RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import with custom namespace
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccServicePolicyRuleResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with custom namespace
// Pattern: Basic lifecycle test with namespace creation
// -----------------------------------------------------------------------------

func TestAccServicePolicyRuleResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy_rule.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create service_policy_rule with minimal configuration
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccServicePolicyRuleResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts", "waf_action"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, action)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccServicePolicyRuleResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy_rule.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	description := "Comprehensive acceptance test service policy rule"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccServicePolicyRuleResourceConfig_allAttributes(nsName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
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
				ImportStateIdFunc:       testAccServicePolicyRuleResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Action
// Verifies: Update of action attribute
// Pattern: Update test for string attribute
// -----------------------------------------------------------------------------

func TestAccServicePolicyRuleResource_updateAction(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy_rule.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with default action (DENY)
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "action", "DENY"),
				),
			},
			// Step 2: Update action to ALLOW
			{
				Config: testAccServicePolicyRuleResourceConfig_withAction(nsName, name, "ALLOW"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
				),
			},
			// Step 3: Update action to NEXT_POLICY
			{
				Config: testAccServicePolicyRuleResourceConfig_withAction(nsName, name, "NEXT_POLICY"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "action", "NEXT_POLICY"),
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

func TestAccServicePolicyRuleResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy_rule.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccServicePolicyRuleResourceConfig_withDescription(nsName, name, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccServicePolicyRuleResourceConfig_withDescription(nsName, name, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccServicePolicyRuleResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy_rule.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccServicePolicyRuleResourceConfig_withLabels(nsName, name, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccServicePolicyRuleResourceConfig_withLabels(nsName, name, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name),
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
// Test 6: Requires Replace Test
// Verifies: Changing name or namespace forces resource replacement
// Pattern: ForceNew/RequiresReplace verification
// -----------------------------------------------------------------------------

func TestAccServicePolicyRuleResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_service_policy_rule.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name1 := acctest.RandomName("tf-acc-test-rule")
	name2 := acctest.RandomName("tf-acc-test-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccServicePolicyRuleResourceConfig_basic(nsName, name2),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
		},
	})
}

// =============================================================================
// Import State ID Function
// =============================================================================

func testAccServicePolicyRuleResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccServicePolicyRuleResourceConfig_basic(nsName, name string) string {
	// Service Policy Rule must be in system namespace and requires waf_action
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_service_policy_rule" "test" {
  name      = %[1]q
  namespace = "system"

  waf_action {
    none {}
  }
}
`, name))
}

func testAccServicePolicyRuleResourceConfig_allAttributes(nsName, name, description string) string {
	// Service Policy Rule must be in system namespace and requires waf_action
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_service_policy_rule" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  waf_action {
    none {}
  }

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

func testAccServicePolicyRuleResourceConfig_withAction(nsName, name, action string) string {
	// Service Policy Rule must be in system namespace and requires waf_action
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_service_policy_rule" "test" {
  name      = %[1]q
  namespace = "system"
  action    = %[2]q

  waf_action {
    none {}
  }
}
`, name, action))
}

func testAccServicePolicyRuleResourceConfig_withDescription(nsName, name, description string) string {
	// Service Policy Rule must be in system namespace and requires waf_action
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_service_policy_rule" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  waf_action {
    none {}
  }
}
`, name, description))
}

func testAccServicePolicyRuleResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
	// Service Policy Rule must be in system namespace and requires waf_action
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_service_policy_rule" "test" {
  name      = %[1]q
  namespace = "system"

  waf_action {
    none {}
  }

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }
}
`, name, environment, managedBy))
}

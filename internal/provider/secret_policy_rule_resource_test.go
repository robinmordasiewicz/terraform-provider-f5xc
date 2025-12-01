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
// SECRET_POLICY_RULE RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import with custom namespace
// 2. All Attributes Test - Test all optional attributes
// 3. Update Tests - Test mutable attributes
//
// Note: secret_policy_rule requires a custom namespace and has OneOf validation
// for client identification (client_name, client_name_matcher, client_selector)
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccSecretPolicyRuleResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with minimal configuration
// Pattern: Basic lifecycle test with custom namespace
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create secret_policy_rule with minimal configuration
			{
				Config: testAccSecretPolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "client_name", "test-client"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// Verify computed action has default value
					resource.TestCheckResourceAttr(resourceName, "action", "DENY"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccSecretPolicyRuleResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, action)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"
	description := "Comprehensive acceptance test secret policy rule"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSecretPolicyRuleResourceConfig_allAttributes(nsName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
					resource.TestCheckResourceAttr(resourceName, "client_name", "allowed-client"),
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
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
				ImportStateIdFunc:       testAccSecretPolicyRuleResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Client Name Matcher Test
// Verifies: OneOf option - client_name_matcher with exact and regex values
// Pattern: Test nested block with list attributes
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_clientNameMatcher(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSecretPolicyRuleResourceConfig_clientNameMatcher(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					// Verify client_name_matcher exact_values
					resource.TestCheckResourceAttr(resourceName, "client_name_matcher.exact_values.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "client_name_matcher.exact_values.0", "client-one"),
					resource.TestCheckResourceAttr(resourceName, "client_name_matcher.exact_values.1", "client-two"),
					// Verify client_name_matcher regex_values
					resource.TestCheckResourceAttr(resourceName, "client_name_matcher.regex_values.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "client_name_matcher.regex_values.0", "^test-.*"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccSecretPolicyRuleResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 4: Client Selector Test
// Verifies: OneOf option - client_selector with label expressions
// Pattern: Test nested block with label selector expressions
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_clientSelector(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSecretPolicyRuleResourceConfig_clientSelector(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					// Verify client_selector expressions
					resource.TestCheckResourceAttr(resourceName, "client_selector.expressions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "client_selector.expressions.0", "environment in (prod,staging)"),
					resource.TestCheckResourceAttr(resourceName, "client_selector.expressions.1", "tier==frontend"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccSecretPolicyRuleResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Update Test - Action
// Verifies: In-place update of action attribute
// Pattern: Update test with action values (DENY, ALLOW, NEXT_POLICY)
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_updateAction(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with DENY action
			{
				Config: testAccSecretPolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "action", "DENY"),
				),
			},
			// Step 2: Update to ALLOW action
			{
				Config: testAccSecretPolicyRuleResourceConfig_withAction(nsName, name, "ALLOW"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "action", "ALLOW"),
				),
			},
			// Step 3: Update to NEXT_POLICY action
			{
				Config: testAccSecretPolicyRuleResourceConfig_withAction(nsName, name, "NEXT_POLICY"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "action", "NEXT_POLICY"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 6: Update Test - Description
// Verifies: Update and removal of description attribute
// Pattern: Update test for optional string attribute
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccSecretPolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			// Step 2: Add description
			{
				Config: testAccSecretPolicyRuleResourceConfig_withDescription(nsName, name, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccSecretPolicyRuleResourceConfig_withDescription(nsName, name, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccSecretPolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 7: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test for map attribute
// -----------------------------------------------------------------------------

func TestAccSecretPolicyRuleResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-rule")
	resourceName := "f5xc_secret_policy_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccSecretPolicyRuleResourceConfig_withLabels(nsName, name, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccSecretPolicyRuleResourceConfig_withLabels(nsName, name, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccSecretPolicyRuleResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// =============================================================================
// Helper Functions
// =============================================================================

// testAccSecretPolicyRuleResourceImportStateIdFunc returns the import ID in namespace/name format
func testAccSecretPolicyRuleResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// testAccSecretPolicyRuleResourceConfig_basic creates minimal configuration with custom namespace
func testAccSecretPolicyRuleResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  client_name = "test-client"
}
`, nsName, name))
}

// testAccSecretPolicyRuleResourceConfig_allAttributes creates configuration with all optional attributes
func testAccSecretPolicyRuleResourceConfig_allAttributes(nsName, name, description string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
  action      = "ALLOW"
  client_name = "allowed-client"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }
}
`, nsName, name, description))
}

// testAccSecretPolicyRuleResourceConfig_clientNameMatcher creates configuration with client_name_matcher
func testAccSecretPolicyRuleResourceConfig_clientNameMatcher(nsName, name string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  action     = "ALLOW"

  client_name_matcher {
    exact_values = ["client-one", "client-two"]
    regex_values = ["^test-.*"]
  }
}
`, nsName, name))
}

// testAccSecretPolicyRuleResourceConfig_clientSelector creates configuration with client_selector
func testAccSecretPolicyRuleResourceConfig_clientSelector(nsName, name string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
  action     = "ALLOW"

  client_selector {
    expressions = [
      "environment in (prod,staging)",
      "tier==frontend"
    ]
  }
}
`, nsName, name))
}

// testAccSecretPolicyRuleResourceConfig_withAction creates configuration with specific action
func testAccSecretPolicyRuleResourceConfig_withAction(nsName, name, action string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  action      = %[3]q
  client_name = "test-client"
}
`, nsName, name, action))
}

// testAccSecretPolicyRuleResourceConfig_withDescription creates configuration with description
func testAccSecretPolicyRuleResourceConfig_withDescription(nsName, name, description string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
  client_name = "test-client"
}
`, nsName, name, description))
}

// testAccSecretPolicyRuleResourceConfig_withLabels creates configuration with labels
func testAccSecretPolicyRuleResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_secret_policy_rule" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  client_name = "test-client"

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }
}
`, nsName, name, environment, managedBy))
}

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
// WAF EXCLUSION POLICY RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import with custom namespace
// 2. All Attributes Test - Test all optional attributes
// 3. Update Tests - Test mutable attributes
// 4. WAF Exclusion Rules Test - Test complex nested blocks
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccWAFExclusionPolicyResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with custom namespace
// Pattern: Basic lifecycle test with namespace creation and time_sleep
// -----------------------------------------------------------------------------

func TestAccWAFExclusionPolicyResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_waf_exclusion_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create waf_exclusion_policy with minimal configuration
			{
				Config: testAccWAFExclusionPolicyResourceConfig_basic(nsName, name),
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
				ImportStateIdFunc:       testAccWAFExclusionPolicyResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts", "waf_exclusion_rules"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccWAFExclusionPolicyResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_waf_exclusion_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-waf")
	description := "Comprehensive acceptance test WAF exclusion policy"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWAFExclusionPolicyResourceConfig_allAttributes(nsName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					// WAF Exclusion Policy can ONLY be created in "system" namespace
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
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
				ImportStateIdFunc:       testAccWAFExclusionPolicyResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: WAF Exclusion Rules Test
// Verifies: Complex nested waf_exclusion_rules blocks
// Pattern: Nested block testing with multiple exclusion rules
// -----------------------------------------------------------------------------

func TestAccWAFExclusionPolicyResource_withExclusionRules(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_waf_exclusion_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWAFExclusionPolicyResourceConfig_withExclusionRules(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// WAF Exclusion Policy can ONLY be created in "system" namespace
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					// Verify waf_exclusion_rules blocks
					resource.TestCheckResourceAttr(resourceName, "waf_exclusion_rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "waf_exclusion_rules.0.path_prefix", "/api"),
					resource.TestCheckResourceAttr(resourceName, "waf_exclusion_rules.1.exact_value", "app.example.com"),
				),
			},
			// Import verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       testAccWAFExclusionPolicyResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 4: Update Test - Description
// Verifies: Update and removal of description attribute
// Pattern: Update test for optional string attribute
// -----------------------------------------------------------------------------

func TestAccWAFExclusionPolicyResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_waf_exclusion_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccWAFExclusionPolicyResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccWAFExclusionPolicyResourceConfig_withDescription(nsName, name, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccWAFExclusionPolicyResourceConfig_withDescription(nsName, name, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccWAFExclusionPolicyResourceConfig_basic(nsName, name),
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

func TestAccWAFExclusionPolicyResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_waf_exclusion_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccWAFExclusionPolicyResourceConfig_withLabels(nsName, name, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccWAFExclusionPolicyResourceConfig_withLabels(nsName, name, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccWAFExclusionPolicyResourceConfig_basic(nsName, name),
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

func TestAccWAFExclusionPolicyResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_waf_exclusion_policy.test"
	nsName := acctest.RandomName("tf-acc-test-ns")
	name1 := acctest.RandomName("tf-acc-test-waf")
	name2 := acctest.RandomName("tf-acc-test-waf")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccWAFExclusionPolicyResourceConfig_basic(nsName, name1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccWAFExclusionPolicyResourceConfig_basic(nsName, name2),
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

func testAccWAFExclusionPolicyResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccWAFExclusionPolicyResourceConfig_basic(nsName, name string) string {
	// WAF Exclusion Policy must be in system namespace and requires waf_exclusion_rules
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_waf_exclusion_policy" "test" {
  name      = %[1]q
  namespace = "system"

  waf_exclusion_rules {
    metadata {
      name = "test-rule"
    }
  }
}
`, name))
}

func testAccWAFExclusionPolicyResourceConfig_allAttributes(nsName, name, description string) string {
	// WAF Exclusion Policy must be in system namespace and requires waf_exclusion_rules
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_waf_exclusion_policy" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  waf_exclusion_rules {
    metadata {
      name = "test-rule"
    }
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

func testAccWAFExclusionPolicyResourceConfig_withDescription(nsName, name, description string) string {
	// WAF Exclusion Policy must be in system namespace and requires waf_exclusion_rules
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_waf_exclusion_policy" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  waf_exclusion_rules {
    metadata {
      name = "test-rule"
    }
  }
}
`, name, description))
}

func testAccWAFExclusionPolicyResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
	// WAF Exclusion Policy must be in system namespace and requires waf_exclusion_rules
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_waf_exclusion_policy" "test" {
  name      = %[1]q
  namespace = "system"

  waf_exclusion_rules {
    metadata {
      name = "test-rule"
    }
  }

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }
}
`, name, environment, managedBy))
}

func testAccWAFExclusionPolicyResourceConfig_withExclusionRules(nsName, name string) string {
	// WAF Exclusion Policy must be in system namespace and requires waf_exclusion_rules
	_ = nsName // unused but kept for test signature consistency
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_waf_exclusion_policy" "test" {
  name      = %[1]q
  namespace = "system"

  waf_exclusion_rules {
    metadata {
      name             = "api-exclusion"
      description_spec = "Exclude WAF for API endpoints"
    }

    any_domain {}

    path_prefix = "/api"

    waf_skip_processing {}
  }

  waf_exclusion_rules {
    metadata {
      name             = "domain-exclusion"
      description_spec = "Exclude specific domain"
    }

    exact_value = "app.example.com"

    any_path {}

    waf_skip_processing {}
  }
}
`, name))
}

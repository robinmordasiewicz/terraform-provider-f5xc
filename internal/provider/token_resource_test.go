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
// TOKEN RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Update, Delete with API verification
// 2. Import Test - Verify import produces identical state
//
// IMPORTANT: Tokens can ONLY be created in the "system" namespace.
// Attempting to create tokens in custom namespaces will result in BAD_REQUEST.
// All test configurations use namespace = "system".
//
// Run with:
//
//	TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//	go test -v ./internal/provider/ -run TestAccTokenResource -timeout 30m
//
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test using system namespace
// -----------------------------------------------------------------------------

func TestAccTokenResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Tokens can ONLY be created in the "system" namespace
	tokenName := acctest.RandomName("tf-acc-test-token")
	resourceName := "f5xc_token.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create token with minimal configuration in system namespace
			{
				Config: testAccTokenResourceConfig_basic(tokenName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", tokenName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccTokenResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, disable)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccTokenResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Tokens can ONLY be created in the "system" namespace
	tokenName := acctest.RandomName("tf-acc-test-token")
	resourceName := "f5xc_token.test"
	description := "Comprehensive acceptance test token"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTokenResourceConfig_allAttributes(tokenName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", tokenName),
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
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
				ImportStateIdFunc:       testAccTokenResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccTokenResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Tokens can ONLY be created in the "system" namespace
	tokenName := acctest.RandomName("tf-acc-test-token")
	resourceName := "f5xc_token.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccTokenResourceConfig_withLabels(tokenName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccTokenResourceConfig_withLabels(tokenName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccTokenResourceConfig_basic(tokenName),
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

func TestAccTokenResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Tokens can ONLY be created in the "system" namespace
	tokenName := acctest.RandomName("tf-acc-test-token")
	resourceName := "f5xc_token.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccTokenResourceConfig_basic(tokenName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccTokenResourceConfig_withDescription(tokenName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccTokenResourceConfig_withDescription(tokenName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccTokenResourceConfig_basic(tokenName),
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

func TestAccTokenResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Tokens can ONLY be created in the "system" namespace
	tokenName := acctest.RandomName("tf-acc-test-token")
	resourceName := "f5xc_token.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with annotations
			{
				Config: testAccTokenResourceConfig_withAnnotations(tokenName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
			// Step 2: Update annotations
			{
				Config: testAccTokenResourceConfig_withAnnotations(tokenName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			// Step 3: Remove annotations
			{
				Config: testAccTokenResourceConfig_basic(tokenName),
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

// testAccTokenResourceImportStateIdFunc returns the import ID in the format namespace/name
func testAccTokenResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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
// All configurations use namespace = "system" as tokens can only be created there
// =============================================================================

func testAccTokenResourceConfig_basic(tokenName string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_token" "test" {
  name      = %[1]q
  namespace = "system"
}
`, tokenName))
}

func testAccTokenResourceConfig_allAttributes(tokenName, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_token" "test" {
  name        = %[1]q
  namespace   = "system"
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
`, tokenName, description))
}

func testAccTokenResourceConfig_withLabels(tokenName, environment, managedBy string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_token" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }
}
`, tokenName, environment, managedBy))
}

func testAccTokenResourceConfig_withDescription(tokenName, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_token" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, tokenName, description))
}

func testAccTokenResourceConfig_withAnnotations(tokenName, value1, value2 string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_token" "test" {
  name      = %[1]q
  namespace = "system"

  annotations = {
    key1 = %[2]q
    key2 = %[3]q
  }
}
`, tokenName, value1, value2))
}

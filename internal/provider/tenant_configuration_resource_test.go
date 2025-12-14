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
// TENANT CONFIGURATION RESOURCE ACCEPTANCE TESTS
//
// These tests follow HashiCorp's acceptance testing best practices:
// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
//
// Test categories implemented:
// 1. Basic Lifecycle Test - Create, Read, Import with minimal config
// 2. All Attributes Test - Test all optional attributes and nested blocks
// 3. Update Tests - Test updating basic_configuration, brute_force_detection_settings, password_policy
// 4. Disappears Test - Handle external resource deletion
// 5. Error/Validation Tests - Invalid configuration handling
// 6. Empty Plan Test - Verify no diff after apply
// 7. Plan Checks - Verify planned actions
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccTenantConfigurationResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Import operations with minimal configuration
// Pattern: Basic lifecycle test
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_basic(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin
	// permissions and may conflict with existing tenant configuration
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create tenant configuration with minimal configuration
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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
				ImportStateIdFunc:       testAccTenantConfigurationResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes and nested blocks
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_allAttributes(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"
	description := "Comprehensive tenant configuration test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTenantConfigurationResourceConfig_allAttributes(nsName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					// basic_configuration block
					resource.TestCheckResourceAttr(resourceName, "basic_configuration.display_name", "Test Tenant"),
					// brute_force_detection_settings block
					resource.TestCheckResourceAttr(resourceName, "brute_force_detection_settings.max_login_failures", "5"),
					// password_policy block
					resource.TestCheckResourceAttr(resourceName, "password_policy.minimum_length", "12"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.digits", "2"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.lowercase_characters", "2"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.uppercase_characters", "2"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.special_characters", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.not_username", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.not_recently_used", "3"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.expire_password", "90"),
				),
			},
			// Import verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
				ImportStateIdFunc:       testAccTenantConfigurationResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Basic Configuration
// Verifies: In-place update of basic_configuration block
// Pattern: Update test for nested block
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_updateBasicConfiguration(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial display name
			{
				Config: testAccTenantConfigurationResourceConfig_withBasicConfig(nsName, name, "Initial Display"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "basic_configuration.display_name", "Initial Display"),
				),
			},
			// Step 2: Update display name
			{
				Config: testAccTenantConfigurationResourceConfig_withBasicConfig(nsName, name, "Updated Display"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "basic_configuration.display_name", "Updated Display"),
				),
			},
			// Step 3: Remove basic_configuration block
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 4: Update Test - Brute Force Detection Settings
// Verifies: Update of brute_force_detection_settings block
// Pattern: Update test for nested block with integer attribute
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_updateBruteForceDetection(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial max_login_failures
			{
				Config: testAccTenantConfigurationResourceConfig_withBruteForce(nsName, name, 5),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "brute_force_detection_settings.max_login_failures", "5"),
				),
			},
			// Step 2: Update max_login_failures
			{
				Config: testAccTenantConfigurationResourceConfig_withBruteForce(nsName, name, 10),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "brute_force_detection_settings.max_login_failures", "10"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Update Test - Password Policy
// Verifies: Update of password_policy block
// Pattern: Update test for nested block with multiple attributes
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_updatePasswordPolicy(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial password policy
			{
				Config: testAccTenantConfigurationResourceConfig_withPasswordPolicy(nsName, name, 8, 1, 1, 1, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "password_policy.minimum_length", "8"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.digits", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.lowercase_characters", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.uppercase_characters", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.special_characters", "1"),
				),
			},
			// Step 2: Update password policy requirements
			{
				Config: testAccTenantConfigurationResourceConfig_withPasswordPolicy(nsName, name, 12, 2, 2, 2, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "password_policy.minimum_length", "12"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.digits", "2"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.lowercase_characters", "2"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.uppercase_characters", "2"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.special_characters", "2"),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 6: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test for map attribute
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_updateLabels(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccTenantConfigurationResourceConfig_withLabels(nsName, name, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccTenantConfigurationResourceConfig_withLabels(nsName, name, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
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
// Test 7: Update Test - Description
// Verifies: Update and removal of description attribute
// Pattern: Update test for optional string attribute
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_updateDescription(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccTenantConfigurationResourceConfig_withDescription(nsName, name, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccTenantConfigurationResourceConfig_withDescription(nsName, name, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 8: Disappears Test
// Verifies: Provider handles external resource deletion gracefully
// Pattern: Disappears test - resource deleted outside Terraform
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_disappears(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
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
// Test 9: Empty Plan Test
// Verifies: No diff on subsequent plan after apply
// Pattern: Empty plan verification - idempotency check
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_emptyPlan(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"
	description := "Empty plan test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccTenantConfigurationResourceConfig_allAttributes(nsName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config: testAccTenantConfigurationResourceConfig_allAttributes(nsName, name, description),
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
// Test 10: Plan Checks Test
// Verifies: Planned actions are correct (create, update, no-op)
// Pattern: Plan check validation
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_planChecks(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create - verify create action planned
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName, name),
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
				Config: testAccTenantConfigurationResourceConfig_withDescription(nsName, name, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			// Step 3: No change - verify no-op planned
			{
				Config: testAccTenantConfigurationResourceConfig_withDescription(nsName, name, "Updated for plan check"),
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
// Test 11: Known Values Plan Check
// Verifies: Planned values match expected values
// Pattern: ExpectKnownValue plan check
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_knownValues(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	name := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTenantConfigurationResourceConfig_withDescription(nsName, name, "Known value test"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(name),
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
// Test 12: Error Test - Invalid Name
// Verifies: Invalid configurations are rejected with appropriate errors
// Pattern: ExpectError test
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_invalidName(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

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
				Config:      testAccTenantConfigurationResourceConfig_basic(nsName, "Invalid-NAME-Test"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|must)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 13: Error Test - Name Too Long
// Verifies: Name length validation
// Pattern: ExpectError test for validation
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_nameTooLong(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName := acctest.RandomName("tf-acc-test-ns")
	// Create a name that exceeds the maximum length (typically 63 characters for K8s-style names)
	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check-tenant-config"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccTenantConfigurationResourceConfig_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|length|long|exceed|character)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 14: Requires Replace Test
// Verifies: Changing name or namespace forces resource replacement
// Pattern: ForceNew/RequiresReplace verification
// -----------------------------------------------------------------------------

func TestAccTenantConfigurationResource_requiresReplace(t *testing.T) {
	// Skip: Tenant Configuration is a singleton resource that requires tenant admin permissions
	t.Skip("Skipping: Tenant Configuration is a singleton resource requiring tenant admin permissions")

	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	nsName1 := acctest.RandomName("tf-acc-test-ns")
	nsName2 := acctest.RandomName("tf-acc-test-ns")
	name1 := acctest.RandomName("tf-acc-test-tenant-config")
	name2 := acctest.RandomName("tf-acc-test-tenant-config")
	resourceName := "f5xc_tenant_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {Source: "hashicorp/time"},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with first name and namespace
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName1, name1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName1, name2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Step 3: Change namespace - should force replacement
			{
				Config: testAccTenantConfigurationResourceConfig_basic(nsName2, name2),
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
// Test Configuration Functions
// =============================================================================

func testAccTenantConfigurationResourceConfig_basic(nsName, name string) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, name))
}

func testAccTenantConfigurationResourceConfig_allAttributes(nsName, name, description string) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q

  labels = {
    environment = "test"
    managed_by  = "terraform"
  }

  annotations = {
    purpose = "acceptance-testing"
  }

  basic_configuration {
    display_name = "Test Tenant"
  }

  brute_force_detection_settings {
    max_login_failures = 5
  }

  password_policy {
    minimum_length         = 12
    digits                 = 2
    lowercase_characters   = 2
    uppercase_characters   = 2
    special_characters     = 1
    not_username           = true
    not_recently_used      = 3
    expire_password        = 90
  }
}
`, nsName, name, description))
}

func testAccTenantConfigurationResourceConfig_withBasicConfig(nsName, name, displayName string) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  basic_configuration {
    display_name = %[3]q
  }
}
`, nsName, name, displayName))
}

func testAccTenantConfigurationResourceConfig_withBruteForce(nsName, name string, maxFailures int) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  brute_force_detection_settings {
    max_login_failures = %[3]d
  }
}
`, nsName, name, maxFailures))
}

func testAccTenantConfigurationResourceConfig_withPasswordPolicy(nsName, name string, minLength, digits, lowercase, uppercase, special int) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  password_policy {
    minimum_length       = %[3]d
    digits               = %[4]d
    lowercase_characters = %[5]d
    uppercase_characters = %[6]d
    special_characters   = %[7]d
  }
}
`, nsName, name, minLength, digits, lowercase, uppercase, special))
}

func testAccTenantConfigurationResourceConfig_withLabels(nsName, name, environment, managedBy string) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
    environment = %[3]q
    managed_by  = %[4]q
  }
}
`, nsName, name, environment, managedBy))
}

func testAccTenantConfigurationResourceConfig_withDescription(nsName, name, description string) string {
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

resource "f5xc_tenant_configuration" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, name, description))
}

// testAccTenantConfigurationResourceImportStateIdFunc returns the import ID for the resource
func testAccTenantConfigurationResourceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
	}
}

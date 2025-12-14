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
// DNS_DOMAIN RESOURCE ACCEPTANCE TESTS
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
// 7. Error/Validation Tests - Invalid configuration handling
// 8. Known Values Test - Verify planned values match expected
// 9. Requires Replace Test - Verify attributes that force replacement
// 10. DNSSEC Mode Test - Test dnssec_mode attribute functionality
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccDNSDomainResource -timeout 30m
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// Verifies: Create, Read, Update, Delete operations with API verification
// Pattern: Basic lifecycle test with CheckDestroy
// -----------------------------------------------------------------------------

func TestAccDNSDomainResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	// Generate a unique domain name using random string
	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create dns_domain with minimal configuration
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify resource exists in Terraform state
					acctest.CheckResourceExists(resourceName),
					// Verify state attributes
					resource.TestCheckResourceAttr(resourceName, "name", domainName),
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
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 2: All Attributes Test
// Verifies: All optional attributes (labels, annotations, description, dnssec_mode, volterra_managed)
// Pattern: Comprehensive attribute coverage
// -----------------------------------------------------------------------------

func TestAccDNSDomainResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"
	description := "Comprehensive acceptance test dns domain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSDomainResourceConfig_allAttributes(domainName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// Verify all attributes in Terraform state
					resource.TestCheckResourceAttr(resourceName, "name", domainName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_mode", "DNSSEC_DISABLE"),
				),
			},
			// Import verification for all attributes
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable", "volterra_managed"},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 3: Update Test - Labels
// Verifies: In-place update of labels attribute
// Pattern: Update test with multiple steps
// -----------------------------------------------------------------------------

func TestAccDNSDomainResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccDNSDomainResourceConfig_withLabels(domainName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccDNSDomainResourceConfig_withLabels(domainName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			// Step 3: Remove all labels
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName),
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

func TestAccDNSDomainResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create without description
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			// Step 2: Add description
			{
				Config: testAccDNSDomainResourceConfig_withDescription(domainName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			// Step 3: Update description
			{
				Config: testAccDNSDomainResourceConfig_withDescription(domainName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			// Step 4: Remove description
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName),
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

func TestAccDNSDomainResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create with annotations
			{
				Config: testAccDNSDomainResourceConfig_withAnnotations(domainName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
			// Step 2: Update annotations
			{
				Config: testAccDNSDomainResourceConfig_withAnnotations(domainName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			// Step 3: Remove annotations
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName),
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

func TestAccDNSDomainResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create resource
			{
				Config: testAccDNSDomainResourceConfig_allAttributes(domainName, "Empty plan test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			// Step 2: Apply same config again - should produce empty plan
			{
				Config: testAccDNSDomainResourceConfig_allAttributes(domainName, "Empty plan test"),
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

func TestAccDNSDomainResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create - verify create action planned
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName),
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
				Config: testAccDNSDomainResourceConfig_withDescription(domainName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			// Step 3: No change - verify no-op planned
			{
				Config: testAccDNSDomainResourceConfig_withDescription(domainName, "Updated for plan check"),
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

func TestAccDNSDomainResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSDomainResourceConfig_withDescription(domainName, "Known value test"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(domainName),
						),
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("namespace"),
							knownvalue.StringExact("system"),
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

func TestAccDNSDomainResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Name with invalid characters (uppercase)
				Config:      testAccDNSDomainResourceConfig_basic("Invalid-NAME-Test.example.com"),
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

func TestAccDNSDomainResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	// Create a name that exceeds the maximum length (typically 63 characters for DNS labels)
	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check-for-dns-domain-resource-because-it-exceeds-maximum-length.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDNSDomainResourceConfig_basic(longName),
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

func TestAccDNSDomainResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDNSDomainResourceConfig_basic(""),
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

func TestAccDNSDomainResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName1 := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	domainName2 := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create with first name
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", domainName1),
				),
			},
			// Step 2: Change name - should force replacement
			{
				Config: testAccDNSDomainResourceConfig_basic(domainName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", domainName2),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 13: DNSSEC Mode Test
// Verifies: dnssec_mode attribute can be set and updated
// Pattern: Attribute-specific test
// -----------------------------------------------------------------------------

func TestAccDNSDomainResource_dnssecMode(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	t.Skip("Skipping: dns_domain creation is disabled in staging environment - requires delegated domain configuration")

	domainName := fmt.Sprintf("%s.example.com", acctest.RandomName("tf-acc-test-dns"))
	resourceName := "f5xc_dns_domain.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_dns_domain"),
		Steps: []resource.TestStep{
			// Step 1: Create with DNSSEC disabled (default)
			{
				Config: testAccDNSDomainResourceConfig_withDNSSEC(domainName, "DNSSEC_DISABLE"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dnssec_mode", "DNSSEC_DISABLE"),
				),
			},
			// Step 2: Update to enable DNSSEC
			{
				Config: testAccDNSDomainResourceConfig_withDNSSEC(domainName, "DNSSEC_ENABLE"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dnssec_mode", "DNSSEC_ENABLE"),
				),
			},
			// Step 3: Update back to disabled
			{
				Config: testAccDNSDomainResourceConfig_withDNSSEC(domainName, "DNSSEC_DISABLE"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dnssec_mode", "DNSSEC_DISABLE"),
				),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccDNSDomainResourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
  name      = %[1]q
  namespace = "system"
}
`, name))
}

func testAccDNSDomainResourceConfig_allAttributes(name, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
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

  dnssec_mode = "DNSSEC_DISABLE"

  volterra_managed {}
}
`, name, description))
}

func testAccDNSDomainResourceConfig_withLabels(name, environment, managedBy string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }
}
`, name, environment, managedBy))
}

func testAccDNSDomainResourceConfig_withDescription(name, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, name, description))
}

func testAccDNSDomainResourceConfig_withAnnotations(name, value1, value2 string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
  name      = %[1]q
  namespace = "system"

  annotations = {
    key1 = %[2]q
    key2 = %[3]q
  }
}
`, name, value1, value2))
}

func testAccDNSDomainResourceConfig_withDNSSEC(name, dnssecMode string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_dns_domain" "test" {
  name        = %[1]q
  namespace   = "system"
  dnssec_mode = %[2]q
}
`, name, dnssecMode))
}

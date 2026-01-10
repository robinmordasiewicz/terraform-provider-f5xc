// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================

// -----------------------------------------------------------------------------
// Test 1: Basic Lifecycle Test
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
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
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.#", "2"),
				),
			},
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
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_withLabelsSystem(rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_withLabelsSystem(rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
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
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "description"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_withDescriptionSystem(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_withDescriptionSystem(rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 5: Update Test - Annotations
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_withAnnotationsSystem(rName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_withAnnotationsSystem(rName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
				),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 6: Disappears Test
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					acctest.CheckBGPAsnSetDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 7: Empty Plan Test
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_allAttributesSystem(rName),
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
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_withLabelsSystem(rName, "test", "terraform"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				Config: testAccBGPAsnSetResourceConfig_withLabelsSystem(rName, "test", "terraform"),
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
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(rName),
						),
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("namespace"),
							knownvalue.StringExact("system"),
						),
					},
				},
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 10: Error Test - Invalid Name
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccBGPAsnSetResourceConfig_basicSystem("Invalid-NAME-Test"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|must)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 11: Error Test - Name Too Long
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccBGPAsnSetResourceConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|length|long|exceed|character)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 12: Error Test - Empty Name
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccBGPAsnSetResourceConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|empty|required|blank)`),
			},
		},
	})
}

// -----------------------------------------------------------------------------
// Test 13: Requires Replace Test
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-acc-test-asn")
	rName2 := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
				),
			},
			{
				Config: testAccBGPAsnSetResourceConfig_basicSystem(rName2),
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
// -----------------------------------------------------------------------------

func TestAccBGPAsnSetResource_asNumbers(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-asn")
	resourceName := "f5xc_bgp_asn_set.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckBGPAsnSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccBGPAsnSetResourceConfig_withASNumbersSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckBGPAsnSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.0", "64512"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.1", "64513"),
					resource.TestCheckResourceAttr(resourceName, "as_numbers.2", "64514"),
				),
			},
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
// Test Configuration Functions - Use "system" namespace
// =============================================================================

func testAccBGPAsnSetResourceConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_bgp_asn_set" "test" {
  name       = %[1]q
  namespace  = "system"
  as_numbers = ["64512"]
}
`, name)
}

func testAccBGPAsnSetResourceConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_bgp_asn_set" "test" {
  name      = %[1]q
  namespace = "system"

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
`, name)
}

func testAccBGPAsnSetResourceConfig_withLabelsSystem(name, environment, managedBy string) string {
	return fmt.Sprintf(`
resource "f5xc_bgp_asn_set" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }

  as_numbers = ["64512"]
}
`, name, environment, managedBy)
}

func testAccBGPAsnSetResourceConfig_withDescriptionSystem(name, description string) string {
	return fmt.Sprintf(`
resource "f5xc_bgp_asn_set" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
  as_numbers  = ["64512"]
}
`, name, description)
}

func testAccBGPAsnSetResourceConfig_withAnnotationsSystem(name, value1, value2 string) string {
	return fmt.Sprintf(`
resource "f5xc_bgp_asn_set" "test" {
  name      = %[1]q
  namespace = "system"

  annotations = {
    key1 = %[2]q
    key2 = %[3]q
  }

  as_numbers = ["64512"]
}
`, name, value1, value2)
}

func testAccBGPAsnSetResourceConfig_withASNumbersSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_bgp_asn_set" "test" {
  name       = %[1]q
  namespace  = "system"
  as_numbers = ["64512", "64513", "64514"]
}
`, name)
}

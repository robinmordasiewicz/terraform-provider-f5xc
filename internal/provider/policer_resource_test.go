// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// TEST: Basic policer creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Import test
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "policer_mode", "policer_type"},
				ImportStateIdFunc:       testAccPolicerImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set including policer-specific fields
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test policer description"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "terraform"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update labels
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withLabelsSystem(rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccPolicerConfig_withLabelsSystem(rName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update description
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withDescriptionSystem(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccPolicerConfig_withDescriptionSystem(rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update annotations
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withAnnotationsSystem(rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccPolicerConfig_withAnnotationsSystem(rName, map[string]string{
					"owner":   "team-b",
					"project": "alpha",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-b"),
					resource.TestCheckResourceAttr(resourceName, "annotations.project", "alpha"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					acctest.CheckPolicerDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
				),
			},
			{
				Config:             testAccPolicerConfig_allAttributesSystem(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for create
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_policer.test", plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: Known values using statecheck
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basicSystem(rName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("name"),
						knownvalue.StringExact(rName),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("id"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

// =============================================================================
// TEST: Invalid name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccPolicerConfig_basicSystem("Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccPolicerConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccPolicerConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-test-policer")
	rName2 := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists("f5xc_policer.test"),
					resource.TestCheckResourceAttr("f5xc_policer.test", "name", rName1),
				),
			},
			{
				Config: testAccPolicerConfig_basicSystem(rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_policer.test", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: Policer-specific attributes (burst_size, committed_information_rate)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccPolicerResource_rateLimits(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withRateLimitsSystem(rName, 10000, 5000),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "committed_information_rate", "10000"),
					resource.TestCheckResourceAttr(resourceName, "burst_size", "5000"),
				),
			},
			{
				Config: testAccPolicerConfig_withRateLimitsSystem(rName, 20000, 10000),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "committed_information_rate", "20000"),
					resource.TestCheckResourceAttr(resourceName, "burst_size", "10000"),
				),
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccPolicerImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccPolicerConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_policer" "test" {
  name                       = %[1]q
  namespace                  = "system"
  committed_information_rate = 10000
  burst_size                 = 5000
}
`, name)
}

func testAccPolicerConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_policer" "test" {
  name                       = %[1]q
  namespace                  = "system"
  description                = "Test policer description"
  committed_information_rate = 10000
  burst_size                 = 5000

  labels = {
    environment = "test"
    team        = "platform"
  }

  annotations = {
    owner = "terraform"
  }
}
`, name)
}

func testAccPolicerConfig_withLabelsSystem(name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_policer" "test" {
  name                       = %[1]q
  namespace                  = "system"
  committed_information_rate = 10000
  burst_size                 = 5000

  labels = {
%[2]s  }
}
`, name, labelsStr)
}

func testAccPolicerConfig_withDescriptionSystem(name, description string) string {
	return fmt.Sprintf(`
resource "f5xc_policer" "test" {
  name                       = %[1]q
  namespace                  = "system"
  description                = %[2]q
  committed_information_rate = 10000
  burst_size                 = 5000
}
`, name, description)
}

func testAccPolicerConfig_withAnnotationsSystem(name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_policer" "test" {
  name                       = %[1]q
  namespace                  = "system"
  committed_information_rate = 10000
  burst_size                 = 5000

  annotations = {
%[2]s  }
}
`, name, annotationsStr)
}

func testAccPolicerConfig_withRateLimitsSystem(name string, cir, burstSize int) string {
	return fmt.Sprintf(`
resource "f5xc_policer" "test" {
  name                       = %[1]q
  namespace                  = "system"
  committed_information_rate = %[2]d
  burst_size                 = %[3]d
}
`, name, cir, burstSize)
}

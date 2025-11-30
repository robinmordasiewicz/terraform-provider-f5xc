// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

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
// TEST: Basic geo_location_set creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccGeoLocationSetResource_basic(t *testing.T) {
	t.Skip("Skipping: geo_location_set API endpoint not available in staging environment (returns 404)")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
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
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccGeoLocationSetImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccGeoLocationSetResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test geo location set description"),
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
func TestAccGeoLocationSetResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withLabelsSystem(rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_withLabelsSystem(rName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
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
func TestAccGeoLocationSetResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withDescriptionSystem(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_withDescriptionSystem(rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
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
func TestAccGeoLocationSetResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withAnnotationsSystem(rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_withAnnotationsSystem(rName, map[string]string{
					"owner":   "team-b",
					"project": "alpha",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
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
func TestAccGeoLocationSetResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					acctest.CheckGeoLocationSetDisappears(resourceName),
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
func TestAccGeoLocationSetResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
				),
			},
			{
				Config:             testAccGeoLocationSetConfig_allAttributesSystem(rName),
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
func TestAccGeoLocationSetResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_geo_location_set.test", plancheck.ResourceActionCreate),
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
func TestAccGeoLocationSetResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basicSystem(rName),
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
func TestAccGeoLocationSetResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccGeoLocationSetConfig_basicSystem("Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccGeoLocationSetResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccGeoLocationSetConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccGeoLocationSetResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccGeoLocationSetConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccGeoLocationSetResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-test-geoloc")
	rName2 := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists("f5xc_geo_location_set.test"),
					resource.TestCheckResourceAttr("f5xc_geo_location_set.test", "name", rName1),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_basicSystem(rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_geo_location_set.test", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: With custom_geo_location_selector block
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccGeoLocationSetResource_customSelector(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withCustomSelectorSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccGeoLocationSetImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccGeoLocationSetConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  name      = %[1]q
  namespace = "system"
}
`, name)
}

func testAccGeoLocationSetConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test geo location set description"

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

func testAccGeoLocationSetConfig_withLabelsSystem(name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
%[2]s  }
}
`, name, labelsStr)
}

func testAccGeoLocationSetConfig_withDescriptionSystem(name, description string) string {
	return fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, name, description)
}

func testAccGeoLocationSetConfig_withAnnotationsSystem(name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  name      = %[1]q
  namespace = "system"

  annotations = {
%[2]s  }
}
`, name, annotationsStr)
}

func testAccGeoLocationSetConfig_withCustomSelectorSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  name      = %[1]q
  namespace = "system"

  custom_geo_location_selector {
    expressions = ["ves.io/country in (US, CA)"]
  }
}
`, name)
}

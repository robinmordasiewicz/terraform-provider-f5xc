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
// TEST: Basic data_group creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataGroupResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
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
				ImportStateIdFunc:       testAccDataGroupImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataGroupResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test data group description"),
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
func TestAccDataGroupResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withLabelsSystem(rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccDataGroupConfig_withLabelsSystem(rName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
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
func TestAccDataGroupResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withDescriptionSystem(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccDataGroupConfig_withDescriptionSystem(rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
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
func TestAccDataGroupResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withAnnotationsSystem(rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccDataGroupConfig_withAnnotationsSystem(rName, map[string]string{
					"owner":   "team-b",
					"project": "alpha",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
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
func TestAccDataGroupResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					acctest.CheckDataGroupDisappears(resourceName),
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
func TestAccDataGroupResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
				),
			},
			{
				Config:             testAccDataGroupConfig_allAttributesSystem(rName),
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
func TestAccDataGroupResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_data_group.test", plancheck.ResourceActionCreate),
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
func TestAccDataGroupResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basicSystem(rName),
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
func TestAccDataGroupResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataGroupConfig_basicSystem("Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataGroupResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataGroupConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataGroupResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataGroupConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataGroupResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-test-datagrp")
	rName2 := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists("f5xc_data_group.test"),
					resource.TestCheckResourceAttr("f5xc_data_group.test", "name", rName1),
				),
			},
			{
				Config: testAccDataGroupConfig_basicSystem(rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_data_group.test", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: With string_records block
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataGroupResource_stringRecords(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withStringRecordsSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccDataGroupImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccDataGroupConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  name      = %[1]q
  namespace = "system"

  string_records {}
}
`, name)
}

func testAccDataGroupConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test data group description"

  labels = {
    environment = "test"
    team        = "platform"
  }

  annotations = {
    owner = "terraform"
  }

  string_records {}
}
`, name)
}

func testAccDataGroupConfig_withLabelsSystem(name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
%[2]s  }

  string_records {}
}
`, name, labelsStr)
}

func testAccDataGroupConfig_withDescriptionSystem(name, description string) string {
	return fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  string_records {}
}
`, name, description)
}

func testAccDataGroupConfig_withAnnotationsSystem(name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  name      = %[1]q
  namespace = "system"

  annotations = {
%[2]s  }

  string_records {}
}
`, name, annotationsStr)
}

func testAccDataGroupConfig_withStringRecordsSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  name      = %[1]q
  namespace = "system"

  string_records {}
}
`, name)
}

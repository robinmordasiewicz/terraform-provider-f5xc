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
// =============================================================================
func TestAccDataGroupResource_basic(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
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
// =============================================================================
func TestAccDataGroupResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_allAttributes(nsName, rName),
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
// =============================================================================
func TestAccDataGroupResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withLabels(nsName, rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccDataGroupConfig_withLabels(nsName, rName, map[string]string{
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
// =============================================================================
func TestAccDataGroupResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccDataGroupConfig_withDescription(nsName, rName, "Updated description"),
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
// =============================================================================
func TestAccDataGroupResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withAnnotations(nsName, rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccDataGroupConfig_withAnnotations(nsName, rName, map[string]string{
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
// =============================================================================
func TestAccDataGroupResource_disappears(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccDataGroupResource_emptyPlan(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists(resourceName),
				),
			},
			{
				Config:             testAccDataGroupConfig_allAttributes(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for create
// =============================================================================
func TestAccDataGroupResource_planChecks(t *testing.T) {
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccDataGroupResource_knownValues(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccDataGroupResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataGroupConfig_basic(nsName, "Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// =============================================================================
func TestAccDataGroupResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")
	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataGroupConfig_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// =============================================================================
func TestAccDataGroupResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataGroupConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// =============================================================================
func TestAccDataGroupResource_requiresReplace(t *testing.T) {
	rName1 := acctest.RandomName("tf-test-datagrp")
	rName2 := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_basic(nsName, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataGroupExists("f5xc_data_group.test"),
					resource.TestCheckResourceAttr("f5xc_data_group.test", "name", rName1),
				),
			},
			{
				Config: testAccDataGroupConfig_basic(nsName, rName2),
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
// =============================================================================
func TestAccDataGroupResource_stringRecords(t *testing.T) {
	resourceName := "f5xc_data_group.test"
	rName := acctest.RandomName("tf-test-datagrp")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataGroupConfig_withStringRecords(nsName, rName),
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
// CONFIG HELPERS
// =============================================================================

// testAccDataGroupConfig_namespaceBase returns the namespace configuration
func testAccDataGroupConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

# Wait for namespace to be ready before creating data_group
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccDataGroupConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccDataGroupConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  string_records {}
}
`, name),
	)
}

func testAccDataGroupConfig_allAttributes(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccDataGroupConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
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
`, name),
	)
}

func testAccDataGroupConfig_withLabels(nsName, name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccDataGroupConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[2]s  }

  string_records {}
}
`, name, labelsStr),
	)
}

func testAccDataGroupConfig_withDescription(nsName, name, description string) string {
	return acctest.ConfigCompose(
		testAccDataGroupConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = %[2]q

  string_records {}
}
`, name, description),
	)
}

func testAccDataGroupConfig_withAnnotations(nsName, name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccDataGroupConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[2]s  }

  string_records {}
}
`, name, annotationsStr),
	)
}

func testAccDataGroupConfig_withStringRecords(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccDataGroupConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_group" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  string_records {}
}
`, name),
	)
}

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
// TEST: Basic policer creation
// =============================================================================
func TestAccPolicerResource_basic(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
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
				ImportStateVerifyIgnore: []string{"timeouts", "policer_mode", "policer_type"},
				ImportStateIdFunc:       testAccPolicerImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set including policer-specific fields
// =============================================================================
func TestAccPolicerResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_allAttributes(nsName, rName),
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
// =============================================================================
func TestAccPolicerResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withLabels(nsName, rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccPolicerConfig_withLabels(nsName, rName, map[string]string{
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
// =============================================================================
func TestAccPolicerResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccPolicerConfig_withDescription(nsName, rName, "Updated description"),
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
// =============================================================================
func TestAccPolicerResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withAnnotations(nsName, rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccPolicerConfig_withAnnotations(nsName, rName, map[string]string{
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
// =============================================================================
func TestAccPolicerResource_disappears(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccPolicerResource_emptyPlan(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
				),
			},
			{
				Config:             testAccPolicerConfig_allAttributes(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for create
// =============================================================================
func TestAccPolicerResource_planChecks(t *testing.T) {
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccPolicerResource_knownValues(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basic(nsName, rName),
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
func TestAccPolicerResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccPolicerConfig_basic(nsName, "Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// =============================================================================
func TestAccPolicerResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")
	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccPolicerConfig_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// =============================================================================
func TestAccPolicerResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccPolicerConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// =============================================================================
func TestAccPolicerResource_requiresReplace(t *testing.T) {
	rName1 := acctest.RandomName("tf-test-policer")
	rName2 := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_basic(nsName, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists("f5xc_policer.test"),
					resource.TestCheckResourceAttr("f5xc_policer.test", "name", rName1),
				),
			},
			{
				Config: testAccPolicerConfig_basic(nsName, rName2),
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
// =============================================================================
func TestAccPolicerResource_rateLimits(t *testing.T) {
	resourceName := "f5xc_policer.test"
	rName := acctest.RandomName("tf-test-policer")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckPolicerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicerConfig_withRateLimits(nsName, rName, 10000, 5000),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckPolicerExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "committed_information_rate", "10000"),
					resource.TestCheckResourceAttr(resourceName, "burst_size", "5000"),
				),
			},
			{
				Config: testAccPolicerConfig_withRateLimits(nsName, rName, 20000, 10000),
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
// CONFIG HELPERS
// =============================================================================

// testAccPolicerConfig_namespaceBase returns the namespace configuration
func testAccPolicerConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

# Wait for namespace to be ready before creating policer
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccPolicerConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccPolicerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_policer" "test" {
  depends_on                 = [time_sleep.wait_for_namespace]
  name                       = %[1]q
  namespace                  = f5xc_namespace.test.name
  committed_information_rate = 10000
  burst_size                 = 5000
}
`, name),
	)
}

func testAccPolicerConfig_allAttributes(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccPolicerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_policer" "test" {
  depends_on                 = [time_sleep.wait_for_namespace]
  name                       = %[1]q
  namespace                  = f5xc_namespace.test.name
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
`, name),
	)
}

func testAccPolicerConfig_withLabels(nsName, name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccPolicerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_policer" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[2]s  }
}
`, name, labelsStr),
	)
}

func testAccPolicerConfig_withDescription(nsName, name, description string) string {
	return acctest.ConfigCompose(
		testAccPolicerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_policer" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = %[2]q
}
`, name, description),
	)
}

func testAccPolicerConfig_withAnnotations(nsName, name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccPolicerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_policer" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[2]s  }
}
`, name, annotationsStr),
	)
}

func testAccPolicerConfig_withRateLimits(nsName, name string, cir, burstSize int) string {
	return acctest.ConfigCompose(
		testAccPolicerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_policer" "test" {
  depends_on                 = [time_sleep.wait_for_namespace]
  name                       = %[1]q
  namespace                  = f5xc_namespace.test.name
  committed_information_rate = %[2]d
  burst_size                 = %[3]d
}
`, name, cir, burstSize),
	)
}

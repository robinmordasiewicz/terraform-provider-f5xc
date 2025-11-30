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
// =============================================================================
func TestAccGeoLocationSetResource_basic(t *testing.T) {
	t.Skip("Skipping: geo_location_set API endpoint not available in staging environment (returns 404)")
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
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
				ImportStateIdFunc:       testAccGeoLocationSetImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set
// =============================================================================
func TestAccGeoLocationSetResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_allAttributes(nsName, rName),
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
// =============================================================================
func TestAccGeoLocationSetResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withLabels(nsName, rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_withLabels(nsName, rName, map[string]string{
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
// =============================================================================
func TestAccGeoLocationSetResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_withDescription(nsName, rName, "Updated description"),
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
// =============================================================================
func TestAccGeoLocationSetResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withAnnotations(nsName, rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_withAnnotations(nsName, rName, map[string]string{
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
// =============================================================================
func TestAccGeoLocationSetResource_disappears(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccGeoLocationSetResource_emptyPlan(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists(resourceName),
				),
			},
			{
				Config:             testAccGeoLocationSetConfig_allAttributes(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for create
// =============================================================================
func TestAccGeoLocationSetResource_planChecks(t *testing.T) {
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basic(nsName, rName),
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
// =============================================================================
func TestAccGeoLocationSetResource_knownValues(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basic(nsName, rName),
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
func TestAccGeoLocationSetResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccGeoLocationSetConfig_basic(nsName, "Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// =============================================================================
func TestAccGeoLocationSetResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")
	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccGeoLocationSetConfig_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// =============================================================================
func TestAccGeoLocationSetResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccGeoLocationSetConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// =============================================================================
func TestAccGeoLocationSetResource_requiresReplace(t *testing.T) {
	rName1 := acctest.RandomName("tf-test-geoloc")
	rName2 := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_basic(nsName, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckGeoLocationSetExists("f5xc_geo_location_set.test"),
					resource.TestCheckResourceAttr("f5xc_geo_location_set.test", "name", rName1),
				),
			},
			{
				Config: testAccGeoLocationSetConfig_basic(nsName, rName2),
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
// =============================================================================
func TestAccGeoLocationSetResource_customSelector(t *testing.T) {
	resourceName := "f5xc_geo_location_set.test"
	rName := acctest.RandomName("tf-test-geoloc")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckGeoLocationSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccGeoLocationSetConfig_withCustomSelector(nsName, rName),
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
// CONFIG HELPERS
// =============================================================================

// testAccGeoLocationSetConfig_namespaceBase returns the namespace configuration
func testAccGeoLocationSetConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

# Wait for namespace to be ready before creating geo_location_set
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccGeoLocationSetConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccGeoLocationSetConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name
}
`, name),
	)
}

func testAccGeoLocationSetConfig_allAttributes(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccGeoLocationSetConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = "Test geo location set description"

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

func testAccGeoLocationSetConfig_withLabels(nsName, name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccGeoLocationSetConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[2]s  }
}
`, name, labelsStr),
	)
}

func testAccGeoLocationSetConfig_withDescription(nsName, name, description string) string {
	return acctest.ConfigCompose(
		testAccGeoLocationSetConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = %[2]q
}
`, name, description),
	)
}

func testAccGeoLocationSetConfig_withAnnotations(nsName, name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccGeoLocationSetConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[2]s  }
}
`, name, annotationsStr),
	)
}

func testAccGeoLocationSetConfig_withCustomSelector(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccGeoLocationSetConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_geo_location_set" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  custom_geo_location_selector {
    expressions = ["ves.io/country in (US, CA)"]
  }
}
`, name),
	)
}

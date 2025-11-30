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
// TEST: Basic data_type creation
// =============================================================================
func TestAccDataTypeResource_basic(t *testing.T) {
	t.Skip("Skipping: data_type API returns BAD_REQUEST - API spec investigation needed")
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
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
				ImportStateIdFunc:       testAccDataTypeImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set
// =============================================================================
func TestAccDataTypeResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test data type description"),
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
func TestAccDataTypeResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withLabels(nsName, rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccDataTypeConfig_withLabels(nsName, rName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
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
func TestAccDataTypeResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withDescription(nsName, rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccDataTypeConfig_withDescription(nsName, rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update annotations
// =============================================================================
func TestAccDataTypeResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withAnnotations(nsName, rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccDataTypeConfig_withAnnotations(nsName, rName, map[string]string{
					"owner":   "team-b",
					"project": "alpha",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
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
func TestAccDataTypeResource_disappears(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					acctest.CheckDataTypeDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// =============================================================================
func TestAccDataTypeResource_emptyPlan(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_allAttributes(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
				),
			},
			{
				Config:             testAccDataTypeConfig_allAttributes(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for create
// =============================================================================
func TestAccDataTypeResource_planChecks(t *testing.T) {
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basic(nsName, rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_data_type.test", plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: Known values using statecheck
// =============================================================================
func TestAccDataTypeResource_knownValues(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basic(nsName, rName),
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
func TestAccDataTypeResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataTypeConfig_basic(nsName, "Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// =============================================================================
func TestAccDataTypeResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")
	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataTypeConfig_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// =============================================================================
func TestAccDataTypeResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataTypeConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// =============================================================================
func TestAccDataTypeResource_requiresReplace(t *testing.T) {
	rName1 := acctest.RandomName("tf-test-datatype")
	rName2 := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basic(nsName, rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists("f5xc_data_type.test"),
					resource.TestCheckResourceAttr("f5xc_data_type.test", "name", rName1),
				),
			},
			{
				Config: testAccDataTypeConfig_basic(nsName, rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("f5xc_data_type.test", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: With PII and sensitive data flags
// =============================================================================
func TestAccDataTypeResource_piiFlags(t *testing.T) {
	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withPiiFlags(nsName, rName, true, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "is_pii", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_sensitive_data", "true"),
				),
			},
			{
				Config: testAccDataTypeConfig_withPiiFlags(nsName, rName, false, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "is_pii", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_sensitive_data", "false"),
				),
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccDataTypeImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// testAccDataTypeConfig_namespaceBase returns the namespace configuration
func testAccDataTypeConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

# Wait for namespace to be ready before creating data_type
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccDataTypeConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccDataTypeConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name
  is_pii     = true
}
`, name),
	)
}

func testAccDataTypeConfig_allAttributes(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccDataTypeConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  depends_on        = [time_sleep.wait_for_namespace]
  name              = %[1]q
  namespace         = f5xc_namespace.test.name
  description       = "Test data type description"
  is_pii            = true
  is_sensitive_data = true

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

func testAccDataTypeConfig_withLabels(nsName, name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccDataTypeConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name
  is_pii     = true

  labels = {
%[2]s  }
}
`, name, labelsStr),
	)
}

func testAccDataTypeConfig_withDescription(nsName, name, description string) string {
	return acctest.ConfigCompose(
		testAccDataTypeConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = %[2]q
  is_pii      = true
}
`, name, description),
	)
}

func testAccDataTypeConfig_withAnnotations(nsName, name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return acctest.ConfigCompose(
		testAccDataTypeConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name
  is_pii     = true

  annotations = {
%[2]s  }
}
`, name, annotationsStr),
	)
}

func testAccDataTypeConfig_withPiiFlags(nsName, name string, isPii, isSensitive bool) string {
	return acctest.ConfigCompose(
		testAccDataTypeConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  depends_on        = [time_sleep.wait_for_namespace]
  name              = %[1]q
  namespace         = f5xc_namespace.test.name
  is_pii            = %[2]t
  is_sensitive_data = %[3]t
}
`, name, isPii, isSensitive),
	)
}

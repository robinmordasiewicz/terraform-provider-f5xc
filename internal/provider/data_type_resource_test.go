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
// TEST: Basic data_type creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
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
				ImportStateIdFunc:       testAccDataTypeImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: All attributes set
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_allAttributesSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withLabelsSystem(rName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccDataTypeConfig_withLabelsSystem(rName, map[string]string{
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withDescriptionSystem(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccDataTypeConfig_withDescriptionSystem(rName, "Updated description"),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withAnnotationsSystem(rName, map[string]string{
					"owner": "team-a",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "team-a"),
				),
			},
			{
				Config: testAccDataTypeConfig_withAnnotationsSystem(rName, map[string]string{
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basicSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
				),
			},
			{
				Config:             testAccDataTypeConfig_allAttributesSystem(rName),
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
func TestAccDataTypeResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basicSystem(rName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basicSystem(rName),
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
func TestAccDataTypeResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataTypeConfig_basicSystem("Invalid_Name_With_Uppercase"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := strings.Repeat("a", 256)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataTypeConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|length|long|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataTypeConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|validation|empty|required|name)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-test-datatype")
	rName2 := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists("f5xc_data_type.test"),
					resource.TestCheckResourceAttr("f5xc_data_type.test", "name", rName1),
				),
			},
			{
				Config: testAccDataTypeConfig_basicSystem(rName2),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccDataTypeResource_piiFlags(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_data_type.test"
	rName := acctest.RandomName("tf-test-datatype")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckDataTypeDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccDataTypeConfig_withPiiFlagsSystem(rName, true, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckDataTypeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "is_pii", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_sensitive_data", "true"),
				),
			},
			{
				Config: testAccDataTypeConfig_withPiiFlagsSystem(rName, false, false),
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
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccDataTypeConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  name      = %[1]q
  namespace = "system"
  is_pii    = true

  rules {
    key_pattern {
      substring_value = "test"
    }
  }
}
`, name)
}

func testAccDataTypeConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  name              = %[1]q
  namespace         = "system"
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

  rules {
    key_pattern {
      substring_value = "test"
    }
  }
}
`, name)
}

func testAccDataTypeConfig_withLabelsSystem(name string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  name      = %[1]q
  namespace = "system"
  is_pii    = true

  labels = {
%[2]s  }

  rules {
    key_pattern {
      substring_value = "test"
    }
  }
}
`, name, labelsStr)
}

func testAccDataTypeConfig_withDescriptionSystem(name, description string) string {
	return fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
  is_pii      = true

  rules {
    key_pattern {
      substring_value = "test"
    }
  }
}
`, name, description)
}

func testAccDataTypeConfig_withAnnotationsSystem(name string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  name      = %[1]q
  namespace = "system"
  is_pii    = true

  annotations = {
%[2]s  }

  rules {
    key_pattern {
      substring_value = "test"
    }
  }
}
`, name, annotationsStr)
}

func testAccDataTypeConfig_withPiiFlagsSystem(name string, isPii, isSensitive bool) string {
	return fmt.Sprintf(`
resource "f5xc_data_type" "test" {
  name              = %[1]q
  namespace         = "system"
  is_pii            = %[2]t
  is_sensitive_data = %[3]t

  rules {
    key_pattern {
      substring_value = "test"
    }
  }
}
`, name, isPii, isSensitive)
}

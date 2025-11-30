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
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// USER_IDENTIFICATION RESOURCE ACCEPTANCE TESTS
//
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// (namespace DELETE API returns 501 Not Implemented)
// =============================================================================

// testAccUserIdentificationImportStateIdFunc returns a function that generates the import ID
func testAccUserIdentificationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

// TestAccUserIdentificationResource_basic tests basic user_identification creation
func TestAccUserIdentificationResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_basicSystem(uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccUserIdentificationImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// TestAccUserIdentificationResource_allAttributes tests user_identification with all attributes
func TestAccUserIdentificationResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_allAttributesSystem(uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test user identification with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "security"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "testing"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccUserIdentificationImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// TestAccUserIdentificationResource_updateLabels tests updating labels
func TestAccUserIdentificationResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_withLabelsSystem(uiName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccUserIdentificationConfig_withLabelsSystem(uiName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// TestAccUserIdentificationResource_updateDescription tests updating description
func TestAccUserIdentificationResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_withDescriptionSystem(uiName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccUserIdentificationConfig_withDescriptionSystem(uiName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccUserIdentificationResource_updateAnnotations tests updating annotations
func TestAccUserIdentificationResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_withAnnotationsSystem(uiName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccUserIdentificationConfig_withAnnotationsSystem(uiName, map[string]string{
					"version": "v2",
					"owner":   "security-team",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "security-team"),
				),
			},
		},
	})
}

// TestAccUserIdentificationResource_disappears tests resource deletion outside of Terraform
func TestAccUserIdentificationResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_basicSystem(uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					acctest.CheckUserIdentificationDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccUserIdentificationResource_emptyPlan tests that no changes are detected after apply
func TestAccUserIdentificationResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_allAttributesSystem(uiName),
			},
			{
				Config:             testAccUserIdentificationConfig_allAttributesSystem(uiName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccUserIdentificationResource_planChecks tests plan-time validation
func TestAccUserIdentificationResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_basicSystem(uiName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

// TestAccUserIdentificationResource_knownValues tests that computed values are known after apply
func TestAccUserIdentificationResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_basicSystem(uiName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("id"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("name"),
						knownvalue.StringExact(uiName),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("namespace"),
						knownvalue.StringExact("system"),
					),
				},
			},
		},
	})
}

// TestAccUserIdentificationResource_invalidName tests validation for invalid name
func TestAccUserIdentificationResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	uiName := "INVALID_NAME_WITH_UPPERCASE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserIdentificationConfig_basicSystem(uiName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccUserIdentificationResource_nameTooLong tests validation for name exceeding max length
func TestAccUserIdentificationResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	uiName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserIdentificationConfig_basicSystem(uiName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccUserIdentificationResource_emptyName tests validation for empty name
func TestAccUserIdentificationResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserIdentificationConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccUserIdentificationResource_requiresReplace tests that name/namespace changes require replace
func TestAccUserIdentificationResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName1 := acctest.RandomName("ui1")
	uiName2 := acctest.RandomName("ui2")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_basicSystem(uiName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName1),
				),
			},
			{
				Config: testAccUserIdentificationConfig_basicSystem(uiName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// TestAccUserIdentificationResource_identificationRules tests user identification with rules
func TestAccUserIdentificationResource_identificationRules(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_user_identification.test"
	uiName := acctest.RandomName("ui")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationConfig_withRulesSystem(uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccUserIdentificationImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccUserIdentificationConfig_basicSystem(uiName string) string {
	return fmt.Sprintf(`
resource "f5xc_user_identification" "test" {
  name       = %[1]q
  namespace  = "system"

  rules {
    client_ip {}
  }
}
`, uiName)
}

func testAccUserIdentificationConfig_allAttributesSystem(uiName string) string {
	return fmt.Sprintf(`
resource "f5xc_user_identification" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test user identification with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "security"
  }

  annotations = {
    purpose = "testing"
  }

  rules {
    client_ip {}
  }
}
`, uiName)
}

func testAccUserIdentificationConfig_withLabelsSystem(uiName string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_user_identification" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
%[2]s  }

  rules {
    client_ip {}
  }
}
`, uiName, labelsStr)
}

func testAccUserIdentificationConfig_withDescriptionSystem(uiName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_user_identification" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  rules {
    client_ip {}
  }
}
`, uiName, description)
}

func testAccUserIdentificationConfig_withAnnotationsSystem(uiName string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_user_identification" "test" {
  name       = %[1]q
  namespace  = "system"

  annotations = {
%[2]s  }

  rules {
    client_ip {}
  }
}
`, uiName, annotationsStr)
}

func testAccUserIdentificationConfig_withRulesSystem(uiName string) string {
	return fmt.Sprintf(`
resource "f5xc_user_identification" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "User identification with identification rules"

  rules {
    client_ip {}
  }
}
`, uiName)
}

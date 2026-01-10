// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

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
// ALERT POLICY RESOURCE ACCEPTANCE TESTS
//
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// (namespace DELETE API returns 501 Not Implemented)
// =============================================================================

// testAccAlertPolicyImportStateIdFunc returns a function that generates the import ID
func testAccAlertPolicyImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// TestAccAlertPolicyResource_basic tests basic alert_policy creation
func TestAccAlertPolicyResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_basicSystem(apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccAlertPolicyImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// TestAccAlertPolicyResource_allAttributes tests alert_policy with all attributes
func TestAccAlertPolicyResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_allAttributesSystem(apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test alert policy with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "monitoring"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "testing"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccAlertPolicyImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// TestAccAlertPolicyResource_updateLabels tests updating labels
func TestAccAlertPolicyResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_withLabelsSystem(apName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccAlertPolicyConfig_withLabelsSystem(apName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// TestAccAlertPolicyResource_updateDescription tests updating description
func TestAccAlertPolicyResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_withDescriptionSystem(apName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccAlertPolicyConfig_withDescriptionSystem(apName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccAlertPolicyResource_updateAnnotations tests updating annotations
func TestAccAlertPolicyResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_withAnnotationsSystem(apName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccAlertPolicyConfig_withAnnotationsSystem(apName, map[string]string{
					"version": "v2",
					"owner":   "ops-team",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ops-team"),
				),
			},
		},
	})
}

// TestAccAlertPolicyResource_disappears tests resource deletion outside of Terraform
func TestAccAlertPolicyResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_basicSystem(apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					acctest.CheckAlertPolicyDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccAlertPolicyResource_emptyPlan tests that no changes are detected after apply
func TestAccAlertPolicyResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_allAttributesSystem(apName),
			},
			{
				Config:             testAccAlertPolicyConfig_allAttributesSystem(apName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccAlertPolicyResource_planChecks tests plan-time validation
func TestAccAlertPolicyResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_basicSystem(apName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

// TestAccAlertPolicyResource_knownValues tests that computed values are known after apply
func TestAccAlertPolicyResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_basicSystem(apName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("id"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("name"),
						knownvalue.StringExact(apName),
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

// TestAccAlertPolicyResource_invalidName tests validation for invalid name
func TestAccAlertPolicyResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	apName := "INVALID_NAME_WITH_UPPERCASE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertPolicyConfig_basicSystem(apName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccAlertPolicyResource_nameTooLong tests validation for name exceeding max length
func TestAccAlertPolicyResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	apName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertPolicyConfig_basicSystem(apName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccAlertPolicyResource_emptyName tests validation for empty name
func TestAccAlertPolicyResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAlertPolicyConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccAlertPolicyResource_requiresReplace tests that name changes require replace
func TestAccAlertPolicyResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName1 := acctest.RandomName("ap1")
	apName2 := acctest.RandomName("ap2")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_basicSystem(apName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName1),
				),
			},
			{
				Config: testAccAlertPolicyConfig_basicSystem(apName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// TestAccAlertPolicyResource_alertSettings tests alert policy with specific settings
func TestAccAlertPolicyResource_alertSettings(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_alert_policy.test"
	apName := acctest.RandomName("ap")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyConfig_withSettingsSystem(apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Alert policy with specific settings"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccAlertPolicyImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccAlertPolicyConfig_basicSystem(apName string) string {
	return fmt.Sprintf(`
resource "f5xc_alert_policy" "test" {
  name       = %[1]q
  namespace  = "system"
}
`, apName)
}

func testAccAlertPolicyConfig_allAttributesSystem(apName string) string {
	return fmt.Sprintf(`
resource "f5xc_alert_policy" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test alert policy with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "monitoring"
  }

  annotations = {
    purpose = "testing"
  }
}
`, apName)
}

func testAccAlertPolicyConfig_withLabelsSystem(apName string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_alert_policy" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
%[2]s  }
}
`, apName, labelsStr)
}

func testAccAlertPolicyConfig_withDescriptionSystem(apName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_alert_policy" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, apName, description)
}

func testAccAlertPolicyConfig_withAnnotationsSystem(apName string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_alert_policy" "test" {
  name       = %[1]q
  namespace  = "system"

  annotations = {
%[2]s  }
}
`, apName, annotationsStr)
}

func testAccAlertPolicyConfig_withSettingsSystem(apName string) string {
	return fmt.Sprintf(`
resource "f5xc_alert_policy" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Alert policy with specific settings"

  labels = {
    alert-type = "monitoring"
  }
}
`, apName)
}

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
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// RATE_LIMITER RESOURCE ACCEPTANCE TESTS
//
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// (namespace DELETE API returns 501 Not Implemented)
// =============================================================================

// testAccRateLimiterImportStateIdFunc returns a function that generates the import ID
func testAccRateLimiterImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// TestAccRateLimiterResource_basic tests basic rate_limiter creation
func TestAccRateLimiterResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_basicSystem(rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccRateLimiterImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// TestAccRateLimiterResource_allAttributes tests rate_limiter with all attributes
func TestAccRateLimiterResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_allAttributesSystem(rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test rate limiter with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "engineering"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "testing"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccRateLimiterImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// TestAccRateLimiterResource_updateLabels tests updating labels
func TestAccRateLimiterResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_withLabelsSystem(rlName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccRateLimiterConfig_withLabelsSystem(rlName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// TestAccRateLimiterResource_updateDescription tests updating description
func TestAccRateLimiterResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_withDescriptionSystem(rlName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccRateLimiterConfig_withDescriptionSystem(rlName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccRateLimiterResource_updateAnnotations tests updating annotations
func TestAccRateLimiterResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_withAnnotationsSystem(rlName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccRateLimiterConfig_withAnnotationsSystem(rlName, map[string]string{
					"version": "v2",
					"owner":   "security-team",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "security-team"),
				),
			},
		},
	})
}

// TestAccRateLimiterResource_disappears tests resource deletion outside of Terraform
func TestAccRateLimiterResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_basicSystem(rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					acctest.CheckRateLimiterDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccRateLimiterResource_emptyPlan tests that no changes are detected after apply
func TestAccRateLimiterResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_allAttributesSystem(rlName),
			},
			{
				Config:             testAccRateLimiterConfig_allAttributesSystem(rlName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccRateLimiterResource_planChecks tests plan-time validation
func TestAccRateLimiterResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_basicSystem(rlName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

// TestAccRateLimiterResource_knownValues tests that computed values are known after apply
func TestAccRateLimiterResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_basicSystem(rlName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("id"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("name"),
						knownvalue.StringExact(rlName),
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

// TestAccRateLimiterResource_invalidName tests validation for invalid name
func TestAccRateLimiterResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rlName := "INVALID_NAME_WITH_UPPERCASE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccRateLimiterConfig_basicSystem(rlName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccRateLimiterResource_nameTooLong tests validation for name exceeding max length
func TestAccRateLimiterResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rlName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccRateLimiterConfig_basicSystem(rlName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccRateLimiterResource_emptyName tests validation for empty name
func TestAccRateLimiterResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccRateLimiterConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccRateLimiterResource_requiresReplace tests that name/namespace changes require replace
func TestAccRateLimiterResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName1 := acctest.RandomName("rl1")
	rlName2 := acctest.RandomName("rl2")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_basicSystem(rlName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName1),
				),
			},
			{
				Config: testAccRateLimiterConfig_basicSystem(rlName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// TestAccRateLimiterResource_rateLimitValues tests rate limit configuration with limits block
func TestAccRateLimiterResource_rateLimitValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_rate_limiter.test"
	rlName := acctest.RandomName("rl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterConfig_withLimitsSystem(rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "limits.0.total_number", "100"),
					resource.TestCheckResourceAttr(resourceName, "limits.0.unit", "MINUTE"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccRateLimiterImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccRateLimiterConfig_basicSystem(rlName string) string {
	return fmt.Sprintf(`
resource "f5xc_rate_limiter" "test" {
  name       = %[1]q
  namespace  = "system"
}
`, rlName)
}

func testAccRateLimiterConfig_allAttributesSystem(rlName string) string {
	return fmt.Sprintf(`
resource "f5xc_rate_limiter" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test rate limiter with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "testing"
  }
}
`, rlName)
}

func testAccRateLimiterConfig_withLabelsSystem(rlName string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_rate_limiter" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
%[2]s  }
}
`, rlName, labelsStr)
}

func testAccRateLimiterConfig_withDescriptionSystem(rlName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_rate_limiter" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, rlName, description)
}

func testAccRateLimiterConfig_withAnnotationsSystem(rlName string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_rate_limiter" "test" {
  name       = %[1]q
  namespace  = "system"

  annotations = {
%[2]s  }
}
`, rlName, annotationsStr)
}

func testAccRateLimiterConfig_withLimitsSystem(rlName string) string {
	return fmt.Sprintf(`
resource "f5xc_rate_limiter" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Rate limiter with limits configuration"

  limits {
    total_number      = 100
    unit              = "MINUTE"
    burst_multiplier  = 2
    period_multiplier = 1

    leaky_bucket {}
  }
}
`, rlName)
}

// Copyright (c) F5, Inc.
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-basic")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_basic(nsName, rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-all")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_allAttributes(nsName, rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test rate limiter with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-lbl")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_withLabels(nsName, rlName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccRateLimiterResourceConfig_withLabels(nsName, rlName, map[string]string{
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-desc")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_withDescription(nsName, rlName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccRateLimiterResourceConfig_withDescription(nsName, rlName, "Updated description"),
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-ann")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_withAnnotations(nsName, rlName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccRateLimiterResourceConfig_withAnnotations(nsName, rlName, map[string]string{
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-disap")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_basic(nsName, rlName),
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
	nsName := acctest.RandomName("rate-limiter-emp")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_allAttributes(nsName, rlName),
			},
			{
				Config:             testAccRateLimiterResourceConfig_allAttributes(nsName, rlName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccRateLimiterResource_planChecks tests plan-time validation
func TestAccRateLimiterResource_planChecks(t *testing.T) {
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-plnchk")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_basic(nsName, rlName),
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-known")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_basic(nsName, rlName),
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
						knownvalue.StringExact(nsName),
					),
				},
			},
		},
	})
}

// TestAccRateLimiterResource_invalidName tests validation for invalid name
func TestAccRateLimiterResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("rate-limiter-inv")
	rlName := "INVALID_NAME_WITH_UPPERCASE"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccRateLimiterResourceConfig_basic(nsName, rlName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccRateLimiterResource_nameTooLong tests validation for name exceeding max length
func TestAccRateLimiterResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("rate-limiter-long")
	rlName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccRateLimiterResourceConfig_basic(nsName, rlName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccRateLimiterResource_emptyName tests validation for empty name
func TestAccRateLimiterResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("rate-limiter-empty")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccRateLimiterResourceConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccRateLimiterResource_requiresReplace tests that name/namespace changes require replace
func TestAccRateLimiterResource_requiresReplace(t *testing.T) {
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-rep")
	rlName1 := acctest.RandomName("rl1")
	rlName2 := acctest.RandomName("rl2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_basic(nsName, rlName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName1),
				),
			},
			{
				Config: testAccRateLimiterResourceConfig_basic(nsName, rlName2),
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
	resourceName := "f5xc_rate_limiter.test"
	nsName := acctest.RandomName("rate-limiter-limits")
	rlName := acctest.RandomName("rl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckRateLimiterDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRateLimiterResourceConfig_withLimits(nsName, rlName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckRateLimiterExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rlName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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

// Test configurations

func testAccRateLimiterResourceConfig_basic(nsName, rlName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_rate_limiter" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, rlName)
}

func testAccRateLimiterResourceConfig_allAttributes(nsName, rlName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_rate_limiter" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
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
`, nsName, rlName)
}

func testAccRateLimiterResourceConfig_withLabels(nsName, rlName string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_rate_limiter" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[3]s  }
}
`, nsName, rlName, labelsStr)
}

func testAccRateLimiterResourceConfig_withDescription(nsName, rlName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_rate_limiter" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, rlName, description)
}

func testAccRateLimiterResourceConfig_withAnnotations(nsName, rlName string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_rate_limiter" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[3]s  }
}
`, nsName, rlName, annotationsStr)
}

func testAccRateLimiterResourceConfig_withLimits(nsName, rlName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_rate_limiter" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Rate limiter with limits configuration"

  limits {
    total_number      = 100
    unit              = "MINUTE"
    burst_multiplier  = 2
    period_multiplier = 1

    leaky_bucket {}
  }
}
`, nsName, rlName)
}

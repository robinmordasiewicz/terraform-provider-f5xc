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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-basic")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_basic(nsName, apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-all")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_allAttributes(nsName, apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-lbl")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_withLabels(nsName, apName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccAlertPolicyResourceConfig_withLabels(nsName, apName, map[string]string{
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-desc")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_withDescription(nsName, apName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccAlertPolicyResourceConfig_withDescription(nsName, apName, "Updated description"),
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-ann")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_withAnnotations(nsName, apName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccAlertPolicyResourceConfig_withAnnotations(nsName, apName, map[string]string{
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-disap")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_basic(nsName, apName),
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
	nsName := acctest.RandomName("alert-pol-emp")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_allAttributes(nsName, apName),
			},
			{
				Config:             testAccAlertPolicyResourceConfig_allAttributes(nsName, apName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccAlertPolicyResource_planChecks tests plan-time validation
func TestAccAlertPolicyResource_planChecks(t *testing.T) {
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-plnchk")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_basic(nsName, apName),
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-known")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_basic(nsName, apName),
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
						knownvalue.StringExact(nsName),
					),
				},
			},
		},
	})
}

// TestAccAlertPolicyResource_invalidName tests validation for invalid name
func TestAccAlertPolicyResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("alert-pol-inv")
	apName := "INVALID_NAME_WITH_UPPERCASE"

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
				Config:      testAccAlertPolicyResourceConfig_basic(nsName, apName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccAlertPolicyResource_nameTooLong tests validation for name exceeding max length
func TestAccAlertPolicyResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("alert-pol-long")
	apName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

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
				Config:      testAccAlertPolicyResourceConfig_basic(nsName, apName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccAlertPolicyResource_emptyName tests validation for empty name
func TestAccAlertPolicyResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("alert-pol-empty")

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
				Config:      testAccAlertPolicyResourceConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccAlertPolicyResource_requiresReplace tests that name/namespace changes require replace
func TestAccAlertPolicyResource_requiresReplace(t *testing.T) {
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-rep")
	apName1 := acctest.RandomName("ap1")
	apName2 := acctest.RandomName("ap2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_basic(nsName, apName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName1),
				),
			},
			{
				Config: testAccAlertPolicyResourceConfig_basic(nsName, apName2),
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
	resourceName := "f5xc_alert_policy.test"
	nsName := acctest.RandomName("alert-pol-settings")
	apName := acctest.RandomName("ap")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckAlertPolicyDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAlertPolicyResourceConfig_withSettings(nsName, apName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckAlertPolicyExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", apName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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

// Test configurations

func testAccAlertPolicyResourceConfig_basic(nsName, apName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_alert_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, apName)
}

func testAccAlertPolicyResourceConfig_allAttributes(nsName, apName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_alert_policy" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
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
`, nsName, apName)
}

func testAccAlertPolicyResourceConfig_withLabels(nsName, apName string, labels map[string]string) string {
	labelsStr := ""
	for k, v := range labels {
		labelsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_alert_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[3]s  }
}
`, nsName, apName, labelsStr)
}

func testAccAlertPolicyResourceConfig_withDescription(nsName, apName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_alert_policy" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, apName, description)
}

func testAccAlertPolicyResourceConfig_withAnnotations(nsName, apName string, annotations map[string]string) string {
	annotationsStr := ""
	for k, v := range annotations {
		annotationsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_alert_policy" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[3]s  }
}
`, nsName, apName, annotationsStr)
}

func testAccAlertPolicyResourceConfig_withSettings(nsName, apName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_alert_policy" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Alert policy with specific settings"

  labels = {
    alert-type = "monitoring"
  }
}
`, nsName, apName)
}

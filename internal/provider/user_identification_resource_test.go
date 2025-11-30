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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-basic")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_basic(nsName, uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-all")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_allAttributes(nsName, uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test user identification with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-lbl")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_withLabels(nsName, uiName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccUserIdentificationResourceConfig_withLabels(nsName, uiName, map[string]string{
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-desc")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_withDescription(nsName, uiName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccUserIdentificationResourceConfig_withDescription(nsName, uiName, "Updated description"),
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-ann")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_withAnnotations(nsName, uiName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccUserIdentificationResourceConfig_withAnnotations(nsName, uiName, map[string]string{
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-disap")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_basic(nsName, uiName),
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
	nsName := acctest.RandomName("user-id-emp")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_allAttributes(nsName, uiName),
			},
			{
				Config:             testAccUserIdentificationResourceConfig_allAttributes(nsName, uiName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccUserIdentificationResource_planChecks tests plan-time validation
func TestAccUserIdentificationResource_planChecks(t *testing.T) {
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-plnchk")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_basic(nsName, uiName),
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-known")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_basic(nsName, uiName),
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
						knownvalue.StringExact(nsName),
					),
				},
			},
		},
	})
}

// TestAccUserIdentificationResource_invalidName tests validation for invalid name
func TestAccUserIdentificationResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("user-id-inv")
	uiName := "INVALID_NAME_WITH_UPPERCASE"

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
				Config:      testAccUserIdentificationResourceConfig_basic(nsName, uiName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccUserIdentificationResource_nameTooLong tests validation for name exceeding max length
func TestAccUserIdentificationResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("user-id-long")
	uiName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

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
				Config:      testAccUserIdentificationResourceConfig_basic(nsName, uiName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccUserIdentificationResource_emptyName tests validation for empty name
func TestAccUserIdentificationResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("user-id-empty")

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
				Config:      testAccUserIdentificationResourceConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccUserIdentificationResource_requiresReplace tests that name/namespace changes require replace
func TestAccUserIdentificationResource_requiresReplace(t *testing.T) {
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-rep")
	uiName1 := acctest.RandomName("ui1")
	uiName2 := acctest.RandomName("ui2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_basic(nsName, uiName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName1),
				),
			},
			{
				Config: testAccUserIdentificationResourceConfig_basic(nsName, uiName2),
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
	resourceName := "f5xc_user_identification.test"
	nsName := acctest.RandomName("user-id-rules")
	uiName := acctest.RandomName("ui")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckUserIdentificationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentificationResourceConfig_withRules(nsName, uiName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckUserIdentificationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", uiName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
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

// Test configurations

func testAccUserIdentificationResourceConfig_basic(nsName, uiName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_user_identification" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  rules {
    client_ip {}
  }
}
`, nsName, uiName)
}

func testAccUserIdentificationResourceConfig_allAttributes(nsName, uiName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_user_identification" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
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
`, nsName, uiName)
}

func testAccUserIdentificationResourceConfig_withLabels(nsName, uiName string, labels map[string]string) string {
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

resource "f5xc_user_identification" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[3]s  }

  rules {
    client_ip {}
  }
}
`, nsName, uiName, labelsStr)
}

func testAccUserIdentificationResourceConfig_withDescription(nsName, uiName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_user_identification" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q

  rules {
    client_ip {}
  }
}
`, nsName, uiName, description)
}

func testAccUserIdentificationResourceConfig_withAnnotations(nsName, uiName string, annotations map[string]string) string {
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

resource "f5xc_user_identification" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[3]s  }

  rules {
    client_ip {}
  }
}
`, nsName, uiName, annotationsStr)
}

func testAccUserIdentificationResourceConfig_withRules(nsName, uiName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_user_identification" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "User identification with identification rules"

  rules {
    client_ip {}
  }
}
`, nsName, uiName)
}

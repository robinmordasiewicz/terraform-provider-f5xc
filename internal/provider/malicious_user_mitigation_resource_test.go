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

// testAccMaliciousUserMitigationImportStateIdFunc returns a function that generates the import ID
func testAccMaliciousUserMitigationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

// TestAccMaliciousUserMitigationResource_basic tests basic malicious_user_mitigation creation
func TestAccMaliciousUserMitigationResource_basic(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-basic")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", mumName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccMaliciousUserMitigationImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_allAttributes tests malicious_user_mitigation with all attributes
func TestAccMaliciousUserMitigationResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-all")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_allAttributes(nsName, mumName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", mumName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test malicious user mitigation with all attributes"),
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
				ImportStateIdFunc:       testAccMaliciousUserMitigationImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable"},
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_updateLabels tests updating labels
func TestAccMaliciousUserMitigationResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-lbl")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withLabels(nsName, mumName, map[string]string{
					"environment": "dev",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withLabels(nsName, mumName, map[string]string{
					"environment": "prod",
					"team":        "platform",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_updateDescription tests updating description
func TestAccMaliciousUserMitigationResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-desc")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withDescription(nsName, mumName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withDescription(nsName, mumName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_updateAnnotations tests updating annotations
func TestAccMaliciousUserMitigationResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-ann")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withAnnotations(nsName, mumName, map[string]string{
					"version": "v1",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v1"),
				),
			},
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withAnnotations(nsName, mumName, map[string]string{
					"version": "v2",
					"owner":   "security-team",
				}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "security-team"),
				),
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_disappears tests resource deletion outside of Terraform
func TestAccMaliciousUserMitigationResource_disappears(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-disap")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					acctest.CheckMaliciousUserMitigationDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_emptyPlan tests that no changes are detected after apply
func TestAccMaliciousUserMitigationResource_emptyPlan(t *testing.T) {
	nsName := acctest.RandomName("mal-user-emp")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_allAttributes(nsName, mumName),
			},
			{
				Config:             testAccMaliciousUserMitigationResourceConfig_allAttributes(nsName, mumName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_planChecks tests plan-time validation
func TestAccMaliciousUserMitigationResource_planChecks(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-plnchk")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_knownValues tests that computed values are known after apply
func TestAccMaliciousUserMitigationResource_knownValues(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-known")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("id"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						resourceName,
						tfjsonpath.New("name"),
						knownvalue.StringExact(mumName),
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

// TestAccMaliciousUserMitigationResource_invalidName tests validation for invalid name
func TestAccMaliciousUserMitigationResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("mal-user-inv")
	mumName := "INVALID_NAME_WITH_UPPERCASE"

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
				Config:      testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|must|name)`),
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_nameTooLong tests validation for name exceeding max length
func TestAccMaliciousUserMitigationResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("mal-user-long")
	mumName := acctest.RandomName("this-name-is-way-too-long-and-exceeds-maximum-length-allowed-for-names-in-f5xc")

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
				Config:      testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName),
				ExpectError: regexp.MustCompile(`(?i)(too long|length|maximum|characters)`),
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_emptyName tests validation for empty name
func TestAccMaliciousUserMitigationResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("mal-user-empty")

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
				Config:      testAccMaliciousUserMitigationResourceConfig_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(empty|required|blank|must)`),
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_requiresReplace tests that name/namespace changes require replace
func TestAccMaliciousUserMitigationResource_requiresReplace(t *testing.T) {
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-rep")
	mumName1 := acctest.RandomName("mum1")
	mumName2 := acctest.RandomName("mum2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", mumName1),
				),
			},
			{
				Config: testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// TestAccMaliciousUserMitigationResource_mitigationRules tests malicious user mitigation with rules
func TestAccMaliciousUserMitigationResource_mitigationRules(t *testing.T) {
	// t.Skip("Skipping: mitigation_type.rules attributes have state drift - ImportStateVerify fails")
	resourceName := "f5xc_malicious_user_mitigation.test"
	nsName := acctest.RandomName("mal-user-rules")
	mumName := acctest.RandomName("mum")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckMaliciousUserMitigationDestroyed,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: "~> 0.9",
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccMaliciousUserMitigationResourceConfig_withMitigationType(nsName, mumName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckMaliciousUserMitigationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", mumName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccMaliciousUserMitigationImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// Test configurations

func testAccMaliciousUserMitigationResourceConfig_basic(nsName, mumName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, mumName)
}

func testAccMaliciousUserMitigationResourceConfig_allAttributes(nsName, mumName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_malicious_user_mitigation" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Test malicious user mitigation with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "security"
  }

  annotations = {
    purpose = "testing"
  }
}
`, nsName, mumName)
}

func testAccMaliciousUserMitigationResourceConfig_withLabels(nsName, mumName string, labels map[string]string) string {
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

resource "f5xc_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[3]s  }
}
`, nsName, mumName, labelsStr)
}

func testAccMaliciousUserMitigationResourceConfig_withDescription(nsName, mumName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_malicious_user_mitigation" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, mumName, description)
}

func testAccMaliciousUserMitigationResourceConfig_withAnnotations(nsName, mumName string, annotations map[string]string) string {
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

resource "f5xc_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[3]s  }
}
`, nsName, mumName, annotationsStr)
}

func testAccMaliciousUserMitigationResourceConfig_withMitigationType(nsName, mumName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_malicious_user_mitigation" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Malicious user mitigation with mitigation type configuration"

  mitigation_type {
    rules {
      threat_level {
        high {}
      }
      mitigation_action {
        block_temporarily {}
      }
    }
  }
}
`, nsName, mumName)
}

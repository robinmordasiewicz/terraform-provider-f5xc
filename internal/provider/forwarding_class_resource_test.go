// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// TestAccForwardingClassResource_basic tests basic forwarding_class creation
func TestAccForwardingClassResource_basic(t *testing.T) {
	t.Skip("Skipping: forwarding_class API endpoint not available in staging environment (returns 404)")
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basic(nsName, fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccForwardingClassImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"timeouts",
				},
			},
		},
	})
}

// TestAccForwardingClassResource_allAttributes tests forwarding_class with all optional attributes
func TestAccForwardingClassResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_allAttributes(nsName, fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test forwarding_class with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "engineering"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
				),
			},
		},
	})
}

// TestAccForwardingClassResource_updateLabels tests updating labels
func TestAccForwardingClassResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_withLabels(nsName, fcName, map[string]string{"env": "dev"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.env", "dev"),
				),
			},
			{
				Config: testAccForwardingClassResource_withLabels(nsName, fcName, map[string]string{"env": "prod", "tier": "frontend"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.env", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.tier", "frontend"),
				),
			},
		},
	})
}

// TestAccForwardingClassResource_updateDescription tests updating description
func TestAccForwardingClassResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_withDescription(nsName, fcName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccForwardingClassResource_withDescription(nsName, fcName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccForwardingClassResource_updateAnnotations tests updating annotations
func TestAccForwardingClassResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_withAnnotations(nsName, fcName, map[string]string{"key1": "value1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
				),
			},
			{
				Config: testAccForwardingClassResource_withAnnotations(nsName, fcName, map[string]string{"key1": "updated", "key2": "value2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
		},
	})
}

// TestAccForwardingClassResource_disappears tests that Terraform handles external deletion
func TestAccForwardingClassResource_disappears(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basic(nsName, fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					acctest.CheckForwardingClassDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccForwardingClassResource_emptyPlan tests that re-applying the same config produces no changes
func TestAccForwardingClassResource_emptyPlan(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basic(nsName, fcName),
			},
			{
				Config:             testAccForwardingClassResource_basic(nsName, fcName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccForwardingClassResource_planChecks tests plan-time checks
func TestAccForwardingClassResource_planChecks(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basic(nsName, fcName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("id")),
					},
				},
			},
			{
				Config: testAccForwardingClassResource_withDescription(nsName, fcName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
		},
	})
}

// TestAccForwardingClassResource_knownValues tests state check values
func TestAccForwardingClassResource_knownValues(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basic(nsName, fcName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fcName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(nsName)),
				},
			},
		},
	})
}

// TestAccForwardingClassResource_invalidName tests validation of invalid names
func TestAccForwardingClassResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccForwardingClassResource_basic(nsName, "Invalid_Name_With_Underscore"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|must)`),
			},
		},
	})
}

// TestAccForwardingClassResource_nameTooLong tests validation of names exceeding max length
func TestAccForwardingClassResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fc")
	longName := strings.Repeat("a", 65) // Exceeds typical 63 character limit

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccForwardingClassResource_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|too long|length|must)`),
			},
		},
	})
}

// TestAccForwardingClassResource_emptyName tests validation of empty names
func TestAccForwardingClassResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccForwardingClassResource_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|empty|required|must)`),
			},
		},
	})
}

// TestAccForwardingClassResource_requiresReplace tests that name change requires replacement
func TestAccForwardingClassResource_requiresReplace(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")
	newFcName := acctest.RandomName("tf-fc-new")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basic(nsName, fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
				),
			},
			{
				Config: testAccForwardingClassResource_basic(nsName, newFcName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// TestAccForwardingClassResource_qosSettings tests the QoS settings (tos_value)
func TestAccForwardingClassResource_qosSettings(t *testing.T) {
	resourceName := "f5xc_forwarding_class.test"
	nsName := acctest.RandomName("tf-ns-fc")
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_qosSettings(nsName, fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
					resource.TestCheckResourceAttr(resourceName, "tos_value", "40"),
					resource.TestCheckResourceAttr(resourceName, "description", "Forwarding class with QoS settings"),
				),
			},
		},
	})
}

// =============================================================================
// TEST CONFIGURATION HELPERS
// =============================================================================

func testAccForwardingClassImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccForwardingClassResource_basic(nsName, fcName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_forwarding_class" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name
}
`, nsName, fcName)
}

func testAccForwardingClassResource_allAttributes(nsName, fcName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_forwarding_class" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Test forwarding_class with all attributes"

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "acceptance-testing"
  }
}
`, nsName, fcName)
}

func testAccForwardingClassResource_withLabels(nsName, fcName string, labels map[string]string) string {
	labelsHCL := ""
	for k, v := range labels {
		labelsHCL += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_forwarding_class" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  labels = {
%[3]s  }
}
`, nsName, fcName, labelsHCL)
}

func testAccForwardingClassResource_withDescription(nsName, fcName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_forwarding_class" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
}
`, nsName, fcName, description)
}

func testAccForwardingClassResource_withAnnotations(nsName, fcName string, annotations map[string]string) string {
	annotationsHCL := ""
	for k, v := range annotations {
		annotationsHCL += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_forwarding_class" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[2]q
  namespace  = f5xc_namespace.test.name

  annotations = {
%[3]s  }
}
`, nsName, fcName, annotationsHCL)
}

func testAccForwardingClassResource_qosSettings(nsName, fcName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_forwarding_class" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Forwarding class with QoS settings"
  tos_value   = 40
}
`, nsName, fcName)
}

// Ensure config.TestStepConfigFunc is used
var _ config.TestStepConfigFunc = nil

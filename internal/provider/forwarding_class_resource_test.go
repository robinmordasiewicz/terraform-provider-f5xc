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
// TEST: Basic forwarding_class creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_basic(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basicSystem(fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
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

// =============================================================================
// TEST: All attributes set
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_allAttributes(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_allAttributesSystem(fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test forwarding_class with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "engineering"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update labels
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_updateLabels(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_withLabelsSystem(fcName, map[string]string{"env": "dev"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.env", "dev"),
				),
			},
			{
				Config: testAccForwardingClassResource_withLabelsSystem(fcName, map[string]string{"env": "prod", "tier": "frontend"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.env", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.tier", "frontend"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update description
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_updateDescription(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_withDescriptionSystem(fcName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccForwardingClassResource_withDescriptionSystem(fcName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
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
func TestAccForwardingClassResource_updateAnnotations(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_withAnnotationsSystem(fcName, map[string]string{"key1": "value1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
				),
			},
			{
				Config: testAccForwardingClassResource_withAnnotationsSystem(fcName, map[string]string{"key1": "updated", "key2": "value2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_disappears(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basicSystem(fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					acctest.CheckForwardingClassDisappears(resourceName),
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
func TestAccForwardingClassResource_emptyPlan(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basicSystem(fcName),
			},
			{
				Config:             testAccForwardingClassResource_basicSystem(fcName),
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
func TestAccForwardingClassResource_planChecks(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basicSystem(fcName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("id")),
					},
				},
			},
			{
				Config: testAccForwardingClassResource_withDescriptionSystem(fcName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
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
func TestAccForwardingClassResource_knownValues(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basicSystem(fcName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fcName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact("system")),
				},
			},
		},
	})
}

// =============================================================================
// TEST: Invalid name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_invalidName(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccForwardingClassResource_basicSystem("Invalid_Name_With_Underscore"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_nameTooLong(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := strings.Repeat("a", 65) // Exceeds typical 63 character limit

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccForwardingClassResource_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|too long|length|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_emptyName(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccForwardingClassResource_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|empty|required|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_requiresReplace(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")
	newFcName := acctest.RandomName("tf-fc-new")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_basicSystem(fcName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckForwardingClassExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fcName),
				),
			},
			{
				Config: testAccForwardingClassResource_basicSystem(newFcName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: QoS settings (tos_value)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccForwardingClassResource_qosSettings(t *testing.T) {
	t.Skip("Skipping: forwarding_class requires tenant quota - API returns 'Quota not configured for kind forwarding_class'")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_forwarding_class.test"
	fcName := acctest.RandomName("tf-fc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckForwardingClassDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccForwardingClassResource_qosSettingsSystem(fcName),
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
// HELPER: Import state ID function
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

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccForwardingClassResource_basicSystem(fcName string) string {
	return fmt.Sprintf(`
resource "f5xc_forwarding_class" "test" {
  name      = %[1]q
  namespace = "system"
}
`, fcName)
}

func testAccForwardingClassResource_allAttributesSystem(fcName string) string {
	return fmt.Sprintf(`
resource "f5xc_forwarding_class" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test forwarding_class with all attributes"

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "acceptance-testing"
  }
}
`, fcName)
}

func testAccForwardingClassResource_withLabelsSystem(fcName string, labels map[string]string) string {
	labelsHCL := ""
	for k, v := range labels {
		labelsHCL += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_forwarding_class" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
%[2]s  }
}
`, fcName, labelsHCL)
}

func testAccForwardingClassResource_withDescriptionSystem(fcName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_forwarding_class" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, fcName, description)
}

func testAccForwardingClassResource_withAnnotationsSystem(fcName string, annotations map[string]string) string {
	annotationsHCL := ""
	for k, v := range annotations {
		annotationsHCL += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_forwarding_class" "test" {
  name      = %[1]q
  namespace = "system"

  annotations = {
%[2]s  }
}
`, fcName, annotationsHCL)
}

func testAccForwardingClassResource_qosSettingsSystem(fcName string) string {
	return fmt.Sprintf(`
resource "f5xc_forwarding_class" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Forwarding class with QoS settings"
  tos_value   = 40
}
`, fcName)
}

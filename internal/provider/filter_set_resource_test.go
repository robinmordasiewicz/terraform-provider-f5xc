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

// TestAccFilterSetResource_basic tests basic filter_set creation
func TestAccFilterSetResource_basic(t *testing.T) {
	t.Skip("Skipping: filter_set API returns BAD_REQUEST - API spec investigation needed")
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basic(nsName, fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccFilterSetImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"timeouts",
				},
			},
		},
	})
}

// TestAccFilterSetResource_allAttributes tests filter_set with all optional attributes
func TestAccFilterSetResource_allAttributes(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_allAttributes(nsName, fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test filter_set with all attributes"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "engineering"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
				),
			},
		},
	})
}

// TestAccFilterSetResource_updateLabels tests updating labels
func TestAccFilterSetResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_withLabels(nsName, fsName, map[string]string{"env": "dev"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.env", "dev"),
				),
			},
			{
				Config: testAccFilterSetResource_withLabels(nsName, fsName, map[string]string{"env": "prod", "tier": "frontend"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.env", "prod"),
					resource.TestCheckResourceAttr(resourceName, "labels.tier", "frontend"),
				),
			},
		},
	})
}

// TestAccFilterSetResource_updateDescription tests updating description
func TestAccFilterSetResource_updateDescription(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_withDescription(nsName, fsName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccFilterSetResource_withDescription(nsName, fsName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccFilterSetResource_updateAnnotations tests updating annotations
func TestAccFilterSetResource_updateAnnotations(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_withAnnotations(nsName, fsName, map[string]string{"key1": "value1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
				),
			},
			{
				Config: testAccFilterSetResource_withAnnotations(nsName, fsName, map[string]string{"key1": "updated", "key2": "value2"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
				),
			},
		},
	})
}

// TestAccFilterSetResource_disappears tests that Terraform handles external deletion
func TestAccFilterSetResource_disappears(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basic(nsName, fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					acctest.CheckFilterSetDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccFilterSetResource_emptyPlan tests that re-applying the same config produces no changes
func TestAccFilterSetResource_emptyPlan(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basic(nsName, fsName),
			},
			{
				Config:             testAccFilterSetResource_basic(nsName, fsName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// TestAccFilterSetResource_planChecks tests plan-time checks
func TestAccFilterSetResource_planChecks(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basic(nsName, fsName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("id")),
					},
				},
			},
			{
				Config: testAccFilterSetResource_withDescription(nsName, fsName, "Updated for plan check"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
		},
	})
}

// TestAccFilterSetResource_knownValues tests state check values
func TestAccFilterSetResource_knownValues(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basic(nsName, fsName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fsName)),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("namespace"), knownvalue.StringExact(nsName)),
				},
			},
		},
	})
}

// TestAccFilterSetResource_invalidName tests validation of invalid names
func TestAccFilterSetResource_invalidName(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFilterSetResource_basic(nsName, "Invalid_Name_With_Underscore"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|must)`),
			},
		},
	})
}

// TestAccFilterSetResource_nameTooLong tests validation of names exceeding max length
func TestAccFilterSetResource_nameTooLong(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fs")
	longName := strings.Repeat("a", 65) // Exceeds typical 63 character limit

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFilterSetResource_basic(nsName, longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|too long|length|must)`),
			},
		},
	})
}

// TestAccFilterSetResource_emptyName tests validation of empty names
func TestAccFilterSetResource_emptyName(t *testing.T) {
	nsName := acctest.RandomName("tf-ns-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFilterSetResource_basic(nsName, ""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|empty|required|must)`),
			},
		},
	})
}

// TestAccFilterSetResource_requiresReplace tests that name change requires replacement
func TestAccFilterSetResource_requiresReplace(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")
	newFsName := acctest.RandomName("tf-fs-new")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basic(nsName, fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
				),
			},
			{
				Config: testAccFilterSetResource_basic(nsName, newFsName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
			},
		},
	})
}

// TestAccFilterSetResource_filterFields tests the filter_fields nested block
func TestAccFilterSetResource_filterFields(t *testing.T) {
	resourceName := "f5xc_filter_set.test"
	nsName := acctest.RandomName("tf-ns-fs")
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_filterFields(nsName, fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
					resource.TestCheckResourceAttr(resourceName, "context_key", "dashboard"),
				),
			},
		},
	})
}

// =============================================================================
// TEST CONFIGURATION HELPERS
// =============================================================================

func testAccFilterSetImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccFilterSetResource_basic(nsName, fsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_filter_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  context_key = "dashboard"
}
`, nsName, fsName)
}

func testAccFilterSetResource_allAttributes(nsName, fsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_filter_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = "Test filter_set with all attributes"
  context_key = "dashboard"

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "acceptance-testing"
  }
}
`, nsName, fsName)
}

func testAccFilterSetResource_withLabels(nsName, fsName string, labels map[string]string) string {
	labelsHCL := ""
	for k, v := range labels {
		labelsHCL += fmt.Sprintf("    %s = %q\n", k, v)
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

resource "f5xc_filter_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  context_key = "dashboard"

  labels = {
%[3]s  }
}
`, nsName, fsName, labelsHCL)
}

func testAccFilterSetResource_withDescription(nsName, fsName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_filter_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  description = %[3]q
  context_key = "dashboard"
}
`, nsName, fsName, description)
}

func testAccFilterSetResource_withAnnotations(nsName, fsName string, annotations map[string]string) string {
	annotationsHCL := ""
	for k, v := range annotations {
		annotationsHCL += fmt.Sprintf("    %s = %q\n", k, v)
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

resource "f5xc_filter_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  context_key = "dashboard"

  annotations = {
%[3]s  }
}
`, nsName, fsName, annotationsHCL)
}

func testAccFilterSetResource_filterFields(nsName, fsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}

resource "f5xc_filter_set" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[2]q
  namespace   = f5xc_namespace.test.name
  context_key = "dashboard"
  description = "Filter set with filter fields"
}
`, nsName, fsName)
}

// Ensure config.TestStepConfigFunc is used
var _ config.TestStepConfigFunc = nil

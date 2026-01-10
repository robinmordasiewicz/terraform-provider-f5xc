// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

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
// TEST: Basic filter_set creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_basic(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basicSystem(fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
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

// =============================================================================
// TEST: All attributes set
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_allAttributes(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_allAttributesSystem(fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test filter_set with all attributes"),
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
func TestAccFilterSetResource_updateLabels(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_withLabelsSystem(fsName, map[string]string{"env": "dev"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.env", "dev"),
				),
			},
			{
				Config: testAccFilterSetResource_withLabelsSystem(fsName, map[string]string{"env": "prod", "tier": "frontend"}),
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
func TestAccFilterSetResource_updateDescription(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_withDescriptionSystem(fsName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccFilterSetResource_withDescriptionSystem(fsName, "Updated description"),
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
func TestAccFilterSetResource_updateAnnotations(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_withAnnotationsSystem(fsName, map[string]string{"key1": "value1"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
				),
			},
			{
				Config: testAccFilterSetResource_withAnnotationsSystem(fsName, map[string]string{"key1": "updated", "key2": "value2"}),
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
func TestAccFilterSetResource_disappears(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basicSystem(fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					acctest.CheckFilterSetDisappears(resourceName),
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
func TestAccFilterSetResource_emptyPlan(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basicSystem(fsName),
			},
			{
				Config:             testAccFilterSetResource_basicSystem(fsName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for create and update
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_planChecks(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basicSystem(fsName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
						plancheck.ExpectUnknownValue(resourceName, tfjsonpath.New("id")),
					},
				},
			},
			{
				Config: testAccFilterSetResource_withDescriptionSystem(fsName, "Updated for plan check"),
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
func TestAccFilterSetResource_knownValues(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basicSystem(fsName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(fsName)),
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
func TestAccFilterSetResource_invalidName(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFilterSetResource_basicSystem("Invalid_Name_With_Underscore"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_nameTooLong(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := strings.Repeat("a", 65) // Exceeds typical 63 character limit

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFilterSetResource_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|too long|length|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation error)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_emptyName(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccFilterSetResource_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|validation|empty|required|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: RequiresReplace on name change
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_requiresReplace(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")
	newFsName := acctest.RandomName("tf-fs-new")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_basicSystem(fsName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckFilterSetExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fsName),
				),
			},
			{
				Config: testAccFilterSetResource_basicSystem(newFsName),
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
// TEST: With filter_fields nested block
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccFilterSetResource_filterFields(t *testing.T) {
	// Removed skip - filter_set now includes required filter_fields
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_filter_set.test"
	fsName := acctest.RandomName("tf-fs")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckFilterSetDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterSetResource_filterFieldsSystem(fsName),
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
// HELPER: Import state ID function
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

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccFilterSetResource_basicSystem(fsName string) string {
	return fmt.Sprintf(`
resource "f5xc_filter_set" "test" {
  name        = %[1]q
  namespace   = "system"
  context_key = "dashboard"

  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }
}
`, fsName)
}

func testAccFilterSetResource_allAttributesSystem(fsName string) string {
	return fmt.Sprintf(`
resource "f5xc_filter_set" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Test filter_set with all attributes"
  context_key = "dashboard"

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "acceptance-testing"
  }

  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }
}
`, fsName)
}

func testAccFilterSetResource_withLabelsSystem(fsName string, labels map[string]string) string {
	labelsHCL := ""
	for k, v := range labels {
		labelsHCL += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_filter_set" "test" {
  name        = %[1]q
  namespace   = "system"
  context_key = "dashboard"

  labels = {
%[2]s  }

  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }
}
`, fsName, labelsHCL)
}

func testAccFilterSetResource_withDescriptionSystem(fsName, description string) string {
	return fmt.Sprintf(`
resource "f5xc_filter_set" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
  context_key = "dashboard"

  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }
}
`, fsName, description)
}

func testAccFilterSetResource_withAnnotationsSystem(fsName string, annotations map[string]string) string {
	annotationsHCL := ""
	for k, v := range annotations {
		annotationsHCL += fmt.Sprintf("    %s = %q\n", k, v)
	}

	return fmt.Sprintf(`
resource "f5xc_filter_set" "test" {
  name        = %[1]q
  namespace   = "system"
  context_key = "dashboard"

  annotations = {
%[2]s  }

  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }
}
`, fsName, annotationsHCL)
}

func testAccFilterSetResource_filterFieldsSystem(fsName string) string {
	return fmt.Sprintf(`
resource "f5xc_filter_set" "test" {
  name        = %[1]q
  namespace   = "system"
  context_key = "dashboard"
  description = "Filter set with filter fields"

  filter_fields {
    field_id = "test-field"
    string_field {
      field_values = ["test-value"]
    }
  }
}
`, fsName)
}

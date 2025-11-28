// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// SAFE STARTING TEST: Namespace Resource
//
// This is the recommended first acceptance test to run when setting up the
// test harness. The namespace resource is the safest choice because:
//
// 1. Minimal required fields (just name)
// 2. No impact on traffic or security
// 3. Fast create/delete operations
// 4. Completely isolated from other resources
// 5. Easy to clean up if something goes wrong
//
// Run with:
//   TF_ACC=1 F5XC_API_URL="..." F5XC_API_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//   go test -v ./internal/provider/ -run TestAccNamespaceResource_basic -timeout 10m
// =============================================================================

// TestAccNamespaceResource_basic tests basic namespace CRUD operations
func TestAccNamespaceResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_namespace"),
		Steps: []resource.TestStep{
			// Step 1: Create namespace with minimal configuration
			{
				Config: testAccNamespaceResourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Step 2: Import state verification
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// Timeouts are not imported, so ignore them
				ImportStateVerifyIgnore: []string{"timeouts"},
			},
		},
	})
}

// TestAccNamespaceResource_withLabels tests namespace with labels
func TestAccNamespaceResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_namespace"),
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig_withLabels(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
				),
			},
		},
	})
}

// TestAccNamespaceResource_update tests namespace update operations
func TestAccNamespaceResource_update(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_namespace"),
		Steps: []resource.TestStep{
			// Step 1: Create with initial labels
			{
				Config: testAccNamespaceResourceConfig_withLabels(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
				),
			},
			// Step 2: Update labels
			{
				Config: testAccNamespaceResourceConfig_updated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.updated", "true"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated namespace for acceptance testing"),
				),
			},
		},
	})
}

// TestAccNamespaceResource_withDescription tests namespace with description
func TestAccNamespaceResource_withDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-ns")
	resourceName := "f5xc_namespace.test"
	description := "Acceptance test namespace - safe to delete"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_namespace"),
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig_withDescription(rName, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", description),
				),
			},
		},
	})
}

// =============================================================================
// Test Configuration Functions
// =============================================================================

func testAccNamespaceResourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"
}
`, name))
}

func testAccNamespaceResourceConfig_withLabels(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }
}
`, name))
}

func testAccNamespaceResourceConfig_updated(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name      = %[1]q
  namespace = "system"

  description = "Updated namespace for acceptance testing"

  labels = {
    environment = "staging"
    managed_by  = "terraform-acceptance-test"
    updated     = "true"
  }
}
`, name))
}

func testAccNamespaceResourceConfig_withDescription(name, description string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q
}
`, name, description))
}

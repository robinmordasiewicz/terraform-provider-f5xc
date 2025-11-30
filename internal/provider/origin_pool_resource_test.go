// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// TEST: Basic origin_pool creation with public_name origin
// =============================================================================
func TestAccOriginPoolResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_origin_pool"),
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "port", "443"),
				),
			},
			// Import test
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccOriginPoolImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: Origin pool with labels and description
// =============================================================================
func TestAccOriginPoolResource_withLabels(t *testing.T) {
	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_origin_pool"),
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_withLabels(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "description", "Test origin pool"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update labels
// =============================================================================
func TestAccOriginPoolResource_updateLabels(t *testing.T) {
	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_origin_pool"),
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_labelsUpdate(nsName, rName, "dev"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccOriginPoolConfig_labelsUpdate(nsName, rName, "prod"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// =============================================================================
func TestAccOriginPoolResource_disappears(t *testing.T) {
	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_origin_pool"),
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					acctest.CheckOriginPoolDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// =============================================================================
func TestAccOriginPoolResource_emptyPlan(t *testing.T) {
	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")
	nsName := acctest.RandomName("tf-test-ns")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_origin_pool"),
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			{
				Config:             testAccOriginPoolConfig_basic(nsName, rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccOriginPoolImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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
// CONFIG HELPERS
// =============================================================================

// testAccOriginPoolConfig_namespaceBase returns the namespace configuration
func testAccOriginPoolConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

# Wait for namespace to be ready before creating origin_pool
resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccOriginPoolConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccOriginPoolConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  port = 443

  origin_servers {
    labels {}  # API returns this even if not set
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}
`, name))
}

func testAccOriginPoolConfig_withLabels(nsName, name string) string {
	return acctest.ConfigCompose(
		testAccOriginPoolConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = %[1]q
  namespace   = f5xc_namespace.test.name
  description = "Test origin pool"

  port = 443

  labels = {
    environment = "test"
    team        = "platform"
  }

  origin_servers {
    labels {}  # API returns this even if not set
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}
`, name))
}

func testAccOriginPoolConfig_labelsUpdate(nsName, name, env string) string {
	return acctest.ConfigCompose(
		testAccOriginPoolConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  port = 443

  labels = {
    environment = %[2]q
  }

  origin_servers {
    labels {}  # API returns this even if not set
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}
`, name, env))
}

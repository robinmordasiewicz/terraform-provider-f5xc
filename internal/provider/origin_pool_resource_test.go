// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccOriginPoolResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckOriginPoolDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccOriginPoolResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckOriginPoolDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_withLabelsSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
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
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccOriginPoolResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckOriginPoolDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_labelsUpdateSystem(rName, "dev"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccOriginPoolConfig_labelsUpdateSystem(rName, "prod"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccOriginPoolResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckOriginPoolDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
					acctest.CheckOriginPoolDisappears(resourceName),
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
func TestAccOriginPoolResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resourceName := "f5xc_origin_pool.test"
	rName := acctest.RandomName("tf-test-op")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckOriginPoolDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
				),
			},
			{
				Config:             testAccOriginPoolConfig_basicSystem(rName),
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
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccOriginPoolConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = %[1]q
  namespace  = "system"

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
`, name)
}

func testAccOriginPoolConfig_withLabelsSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name        = %[1]q
  namespace   = "system"
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
`, name)
}

func testAccOriginPoolConfig_labelsUpdateSystem(name, env string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = %[1]q
  namespace  = "system"

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
`, name, env)
}

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
// TEST: Basic http_loadbalancer creation
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccHTTPLoadBalancerResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccHTTPLoadBalancerConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccHTTPLoadBalancerImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"timeouts",
					"http.dns_volterra_managed",
					"l7_ddos_protection",
				},
			},
		},
	})
}

// =============================================================================
// TEST: HTTP loadbalancer with labels and description
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccHTTPLoadBalancerResource_withLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerConfig_withLabelsSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.team", "platform"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: HTTP loadbalancer with multiple domains
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccHTTPLoadBalancerResource_withDomains(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb-domain")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerConfig_withDomainsSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "2"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update labels
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccHTTPLoadBalancerResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerConfig_labelsUpdateSystem(rName, "dev"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "dev"),
				),
			},
			{
				Config: testAccHTTPLoadBalancerConfig_labelsUpdateSystem(rName, "prod"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "prod"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// =============================================================================
func TestAccHTTPLoadBalancerResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			{
				Config:             testAccHTTPLoadBalancerConfig_basicSystem(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// HELPER: Import state ID function
// =============================================================================
func testAccHTTPLoadBalancerImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccHTTPLoadBalancerConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}
}
`, name)
}

func testAccHTTPLoadBalancerConfig_withLabelsSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    team        = "platform"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}
}
`, name)
}

func testAccHTTPLoadBalancerConfig_withDomainsSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
  }

  domains = [
    "app.example.com",
    "api.example.com"
  ]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}
}
`, name)
}

func testAccHTTPLoadBalancerConfig_labelsUpdateSystem(name, env string) string {
	return fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = %[2]q
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}
}
`, name, env)
}

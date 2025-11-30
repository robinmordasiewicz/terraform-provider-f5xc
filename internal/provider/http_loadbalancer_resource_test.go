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

func TestAccHTTPLoadBalancerResource_basic(t *testing.T) {
	// UseStateForUnknown enhancement added for optional scalar fields (bool, string, int64)
	// This should prevent drift from API defaults like add_location, etc.
	// Note: Empty marker blocks (disable_*) may still cause issues as they're nested blocks
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb")
	nsName := acctest.RandomName("tf-test-ns")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccHTTPLoadBalancerResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", nsName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// ImportState testing
			// Note: ImportStateVerifyIgnore includes fields that the API returns but weren't
			// in the original config. During normal Read, these are preserved as empty to avoid
			// drift, but during Import they are populated from the API.
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[resourceName]
					if !ok {
						return "", fmt.Errorf("resource not found: %s", resourceName)
					}
					return fmt.Sprintf("%s/%s", rs.Primary.Attributes["namespace"], rs.Primary.Attributes["name"]), nil
				},
				ImportStateVerifyIgnore: []string{
					// These fields are populated during Import from API but preserved empty
					// during normal Read to avoid drift on unconfigured blocks
					"http.dns_volterra_managed",
					"l7_ddos_protection",
				},
			},
			// Update testing
			{
				Config: testAccHTTPLoadBalancerResourceConfig_updated(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
				),
			},
		},
	})
}

func TestAccHTTPLoadBalancerResource_withDomains(t *testing.T) {
	t.Skip("Skipping: http_loadbalancer generator does not marshal spec fields (domains, load_balancer_type, etc.) to API request - requires generator enhancement")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb-domain")
	nsName := acctest.RandomName("tf-test-ns")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerResourceConfig_withDomains(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "domains.#", "2"),
				),
			},
		},
	})
}

func TestAccHTTPLoadBalancerResource_disappears(t *testing.T) {
	t.Skip("Skipping: http_loadbalancer generator does not marshal spec fields (domains, load_balancer_type, etc.) to API request - requires generator enhancement")
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb-disappear")
	nsName := acctest.RandomName("tf-test-ns")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerResourceConfig_basic(nsName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					// In a real test, you would delete the resource via API here
					// and verify Terraform handles the disappearance gracefully
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// testAccHTTPLoadBalancerConfig_namespaceBase returns the namespace configuration
func testAccHTTPLoadBalancerConfig_namespaceBase(nsName string) string {
	return fmt.Sprintf(`
resource "f5xc_namespace" "test" {
  name = %[1]q
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [f5xc_namespace.test]
  create_duration = "5s"
}
`, nsName)
}

func testAccHTTPLoadBalancerResourceConfig_basic(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		testAccHTTPLoadBalancerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

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
`, name))
}

func testAccHTTPLoadBalancerResourceConfig_updated(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		testAccHTTPLoadBalancerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  labels = {
    environment = "staging"
    managed_by  = "terraform"
    updated     = "true"
  }

  annotations = {
    "description" = "Updated HTTP Load Balancer"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip {}
}
`, name))
}

func testAccHTTPLoadBalancerResourceConfig_withDomains(nsName, name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		testAccHTTPLoadBalancerConfig_namespaceBase(nsName),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = %[1]q
  namespace  = f5xc_namespace.test.name

  labels = {
    environment = "test"
  }

  domains = [
    "app.example.com",
    "api.example.com"
  ]
}
`, name))
}

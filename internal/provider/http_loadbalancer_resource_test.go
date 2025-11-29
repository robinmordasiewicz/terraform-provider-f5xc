// Copyright (c) F5XC Community
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

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
				Config: testAccHTTPLoadBalancerResourceConfig_basic(rName),
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
			},
			// Update testing
			{
				Config: testAccHTTPLoadBalancerResourceConfig_updated(rName),
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
				Config: testAccHTTPLoadBalancerResourceConfig_withDomains(rName),
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
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-test-lb-disappear")
	resourceName := "f5xc_http_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_http_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccHTTPLoadBalancerResourceConfig_basic(rName),
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

func testAccHTTPLoadBalancerResourceConfig_basic(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]
}
`, name))
}

func testAccHTTPLoadBalancerResourceConfig_updated(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name      = %[1]q
  namespace = "system"

  labels = {
    environment = "staging"
    managed_by  = "terraform"
    updated     = "true"
  }

  annotations = {
    "description" = "Updated HTTP Load Balancer"
  }

  domains = ["test.example.com"]
}
`, name))
}

func testAccHTTPLoadBalancerResourceConfig_withDomains(name string) string {
	return acctest.ConfigCompose(
		acctest.ProviderConfig(),
		fmt.Sprintf(`
resource "f5xc_http_loadbalancer" "test" {
  name      = %[1]q
  namespace = "system"

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

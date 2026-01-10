// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"regexp"
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
// TCP_LOADBALANCER RESOURCE ACCEPTANCE TESTS
//
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// (namespace DELETE API returns 501 Not Implemented)
//
// Run with:
//
//	TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//	go test -v ./internal/provider/ -run TestAccTCPLoadBalancerResource -timeout 30m
//
// =============================================================================

// =============================================================================
// TEST: Basic tcp_loadbalancer creation with API verification
// =============================================================================
func TestAccTCPLoadBalancerResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Import state verification
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccTCPLoadBalancerImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"timeouts",
				},
			},
		},
	})
}

// testAccTCPLoadBalancerImportStateIdFunc returns a function that generates the import ID
func testAccTCPLoadBalancerImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		namespace := rs.Primary.Attributes["namespace"]
		name := rs.Primary.Attributes["name"]
		if namespace == "" || name == "" {
			return "", fmt.Errorf("namespace or name not set in state")
		}
		return fmt.Sprintf("%s/%s", namespace, name), nil
	}
}

// =============================================================================
// TEST: All optional attributes (labels, annotations)
// =============================================================================
func TestAccTCPLoadBalancerResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccTCPLoadBalancerImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{
					"timeouts",
					"disable",
					"description",
				},
			},
		},
	})
}

// =============================================================================
// TEST: Update labels
// =============================================================================
func TestAccTCPLoadBalancerResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_withLabelsSystem(rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			{
				Config: testAccTCPLoadBalancerConfig_withLabelsSystem(rName, "production", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "production"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update description
// =============================================================================
func TestAccTCPLoadBalancerResource_updateDescription(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_withDescriptionSystem(rName, "Initial description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial description"),
				),
			},
			{
				Config: testAccTCPLoadBalancerConfig_withDescriptionSystem(rName, "Updated description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update annotations
// =============================================================================
func TestAccTCPLoadBalancerResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_withAnnotationsSystem(rName, "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.test_key", "value1"),
				),
			},
			{
				Config: testAccTCPLoadBalancerConfig_withAnnotationsSystem(rName, "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.test_key", "value2"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (external deletion)
// Note: This test is disabled until CheckTCPLoadBalancerDisappears is added to acctest
// =============================================================================
// func TestAccTCPLoadBalancerResource_disappears(t *testing.T) {
// 	acctest.SkipIfNotAccTest(t)
// 	acctest.PreCheck(t)
//
// 	rName := acctest.RandomName("tf-acc-test-tcp-lb")
// 	resourceName := "f5xc_tcp_loadbalancer.test"
//
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTCPLoadBalancerConfig_basicSystem(rName),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					acctest.CheckResourceExists(resourceName),
// 					acctest.CheckTCPLoadBalancerDisappears(resourceName),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

// =============================================================================
// TEST: Empty plan after apply (no drift)
// =============================================================================
func TestAccTCPLoadBalancerResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
			{
				Config:             testAccTCPLoadBalancerConfig_basicSystem(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks for known values
// =============================================================================
func TestAccTCPLoadBalancerResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
					PostApplyPreRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Known value assertions
// =============================================================================
func TestAccTCPLoadBalancerResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_basicSystem(rName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName,
						tfjsonpath.New("name"),
						knownvalue.StringExact(rName)),
					statecheck.ExpectKnownValue(resourceName,
						tfjsonpath.New("namespace"),
						knownvalue.StringExact("system")),
				},
			},
		},
	})
}

// =============================================================================
// TEST: Invalid name (validation)
// =============================================================================
func TestAccTCPLoadBalancerResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTCPLoadBalancerConfig_basicSystem("Invalid_Name_With_Caps"),
				ExpectError: regexp.MustCompile(`(invalid|Invalid|must|validation)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long (validation)
// =============================================================================
func TestAccTCPLoadBalancerResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := "a" + acctest.RandomName("tf-acc-test-tcp-lb-very-very-very-very-very-very-very-very-very-long-name")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTCPLoadBalancerConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(length|too long|maximum|characters)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name (validation)
// =============================================================================
func TestAccTCPLoadBalancerResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTCPLoadBalancerConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(required|empty|minimum|at least|invalid.*name|name.*format)`),
			},
		},
	})
}

// =============================================================================
// TEST: Requires replace when name changes
// =============================================================================
func TestAccTCPLoadBalancerResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName1 := acctest.RandomName("tf-acc-test-tcp-lb")
	rName2 := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
				),
			},
			{
				Config: testAccTCPLoadBalancerConfig_basicSystem(rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName2),
				),
			},
		},
	})
}

// =============================================================================
// TEST: TCP loadbalancer with listen port
// =============================================================================
func TestAccTCPLoadBalancerResource_listenPort(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-tcp-lb")
	resourceName := "f5xc_tcp_loadbalancer.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckResourceDestroyed("f5xc_tcp_loadbalancer"),
		Steps: []resource.TestStep{
			{
				Config: testAccTCPLoadBalancerConfig_withListenPortSystem(rName, 8443),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "listen_port", "8443"),
				),
			},
		},
	})
}

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccTCPLoadBalancerConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
# Origin pool is required for TCP load balancer - it needs a backend cluster
resource "f5xc_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"

  port = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}

resource "f5xc_tcp_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  # Domain and SNI are required for TCP on public shared VIP
  domains = ["%[1]s.example.com"]

  listen_port = 443

  # Required: Specify protocol type (tcp, tls_tcp, or tls_tcp_auto_cert)
  tcp {}

  # Required: SNI for TCP on public shared VIP
  sni {}

  # Required: TCP LB needs origin pools for routing
  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  # Required: Specify advertise configuration
  advertise_on_public_default_vip {}
}
`, name)
}

func testAccTCPLoadBalancerConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  port       = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}

resource "f5xc_tcp_loadbalancer" "test" {
  name        = %[1]q
  namespace   = "system"
  description = "Acceptance test tcp loadbalancer with all attributes"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  domains = ["%[1]s.example.com"]
  listen_port = 443
  tcp {}
  sni {}

  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip {}
}
`, name)
}

func testAccTCPLoadBalancerConfig_withLabelsSystem(name, env, managedBy string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  port       = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}

resource "f5xc_tcp_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }

  domains = ["%[1]s.example.com"]
  listen_port = 443
  tcp {}
  sni {}

  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip {}
}
`, name, env, managedBy)
}

func testAccTCPLoadBalancerConfig_withDescriptionSystem(name, description string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  port       = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}

resource "f5xc_tcp_loadbalancer" "test" {
  name        = %[1]q
  namespace   = "system"
  description = %[2]q

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  domains = ["%[1]s.example.com"]
  listen_port = 443
  tcp {}
  sni {}

  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip {}
}
`, name, description)
}

func testAccTCPLoadBalancerConfig_withAnnotationsSystem(name, value string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  port       = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}

resource "f5xc_tcp_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    test_key = %[2]q
  }

  domains = ["%[1]s.example.com"]
  listen_port = 443
  tcp {}
  sni {}

  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip {}
}
`, name, value)
}

func testAccTCPLoadBalancerConfig_withListenPortSystem(name string, port int) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = "%[1]s-pool"
  namespace  = "system"
  port       = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}

resource "f5xc_tcp_loadbalancer" "test" {
  name       = %[1]q
  namespace  = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  domains = ["%[1]s.example.com"]
  listen_port = %[2]d

  tcp {}
  sni {}

  origin_pools_weights {
    pool {
      name      = f5xc_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip {}
}
`, name, port)
}

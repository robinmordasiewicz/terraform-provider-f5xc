// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/f5xc/terraform-provider-f5xc/internal/acctest"
)

// =============================================================================
// HEALTHCHECK RESOURCE ACCEPTANCE TESTS
//
// Uses "system" namespace to avoid creating test namespaces that can't be deleted
// (namespace DELETE API returns 501 Not Implemented)
//
// Run with:
//
//	TF_ACC=1 F5XC_API_URL="..." F5XC_P12_FILE="..." F5XC_P12_PASSWORD="..." \
//	go test -v ./internal/provider/ -run TestAccHealthcheckResource -timeout 30m
//
// =============================================================================

// =============================================================================
// TEST: Basic healthcheck creation with API verification
// =============================================================================
func TestAccHealthcheckResource_basic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckResourceExists(resourceName),
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Import state verification
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// testAccHealthcheckImportStateIdFunc returns a function that generates the import ID
func testAccHealthcheckImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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
func TestAccHealthcheckResource_allAttributes(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "namespace", "system"),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-acceptance-test"),
					resource.TestCheckResourceAttr(resourceName, "annotations.purpose", "acceptance-testing"),
					resource.TestCheckResourceAttr(resourceName, "annotations.owner", "ci-cd"),
					acctest.CheckHealthcheckAttributes(resourceName,
						map[string]string{
							"environment": "test",
							"managed_by":  "terraform-acceptance-test",
						},
						map[string]string{
							"purpose": "acceptance-testing",
							"owner":   "ci-cd",
						},
					),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "disable", "description"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// TEST: Update labels
// =============================================================================
func TestAccHealthcheckResource_updateLabels(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_withLabelsSystem(rName, "test", "terraform"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform"),
				),
			},
			{
				Config: testAccHealthcheckConfig_withLabelsSystem(rName, "staging", "terraform-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "labels.environment", "staging"),
					resource.TestCheckResourceAttr(resourceName, "labels.managed_by", "terraform-updated"),
					acctest.CheckHealthcheckAttributes(resourceName,
						map[string]string{
							"environment": "staging",
							"managed_by":  "terraform-updated",
						},
						nil,
					),
				),
			},
			{
				Config: testAccHealthcheckConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckNoResourceAttr(resourceName, "labels.environment"),
					resource.TestCheckNoResourceAttr(resourceName, "labels.managed_by"),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Update annotations
// =============================================================================
func TestAccHealthcheckResource_updateAnnotations(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_withAnnotationsSystem(rName, "value1", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "value2"),
					acctest.CheckHealthcheckAttributes(resourceName, nil,
						map[string]string{
							"key1": "value1",
							"key2": "value2",
						},
					),
				),
			},
			{
				Config: testAccHealthcheckConfig_withAnnotationsSystem(rName, "updated1", "updated2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "annotations.key1", "updated1"),
					resource.TestCheckResourceAttr(resourceName, "annotations.key2", "updated2"),
				),
			},
			{
				Config: testAccHealthcheckConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
				),
			},
		},
	})
}

// =============================================================================
// TEST: Resource disappears (deleted outside Terraform)
// =============================================================================
func TestAccHealthcheckResource_disappears(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_basicSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					acctest.CheckHealthcheckDisappears(resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// =============================================================================
// TEST: Empty plan after apply (no drift)
// =============================================================================
func TestAccHealthcheckResource_emptyPlan(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_allAttributesSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
				),
			},
			{
				Config: testAccHealthcheckConfig_allAttributesSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: Plan checks (create, update, no-op)
// =============================================================================
func TestAccHealthcheckResource_planChecks(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
				),
			},
			{
				Config: testAccHealthcheckConfig_withLabelsSystem(rName, "test", "terraform"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				Config: testAccHealthcheckConfig_withLabelsSystem(rName, "test", "terraform"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: Known values plan check
// =============================================================================
func TestAccHealthcheckResource_knownValues(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_basicSystem(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(rName),
						),
						plancheck.ExpectKnownValue(resourceName,
							tfjsonpath.New("namespace"),
							knownvalue.StringExact("system"),
						),
					},
				},
			},
		},
	})
}

// =============================================================================
// TEST: Invalid name error
// =============================================================================
func TestAccHealthcheckResource_invalidName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccHealthcheckConfig_basicSystem("Invalid-NAME-Test"),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|must)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name too long error
// =============================================================================
func TestAccHealthcheckResource_nameTooLong(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	longName := "tf-acc-test-this-name-is-way-too-long-and-should-fail-validation-check"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccHealthcheckConfig_basicSystem(longName),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|length|long|exceed|character)`),
			},
		},
	})
}

// =============================================================================
// TEST: Empty name error
// =============================================================================
func TestAccHealthcheckResource_emptyName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccHealthcheckConfig_basicSystem(""),
				ExpectError: regexp.MustCompile(`(?i)(invalid|name|empty|required|blank)`),
			},
		},
	})
}

// =============================================================================
// TEST: Name change requires replacement
// =============================================================================
func TestAccHealthcheckResource_requiresReplace(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName1 := acctest.RandomName("tf-acc-test-hc")
	rName2 := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_basicSystem(rName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName1),
				),
			},
			{
				Config: testAccHealthcheckConfig_basicSystem(rName2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName2),
				),
			},
		},
	})
}

// =============================================================================
// TEST: HTTP health check nested block
// =============================================================================
func TestAccHealthcheckResource_httpHealthCheck(t *testing.T) {
	// Testing actual error to debug schema type mismatch
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_httpHealthCheckSystem(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.path", "/health"),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "example.com"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "http_health_check.use_http2"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// CONFIG HELPERS - Use "system" namespace
// =============================================================================

func testAccHealthcheckConfig_basicSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  tcp_health_check {}
}
`, name)
}

func testAccHealthcheckConfig_allAttributesSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  tcp_health_check {}
}
`, name)
}

func testAccHealthcheckConfig_withLabelsSystem(name, environment, managedBy string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = %[2]q
    managed_by  = %[3]q
  }

  tcp_health_check {}
}
`, name, environment, managedBy)
}

func testAccHealthcheckConfig_withAnnotationsSystem(name, value1, value2 string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  annotations = {
    key1 = %[2]q
    key2 = %[3]q
  }

  tcp_health_check {}
}
`, name, value1, value2)
}

func testAccHealthcheckConfig_httpHealthCheckSystem(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path        = "/health"
    host_header = "example.com"
  }
}
`, name)
}

// =============================================================================
// ConflictsWith TEST CONFIGURATION HELPERS
// =============================================================================

// testAccHealthcheckConfig_httpWithHostHeaderOnly creates HTTP healthcheck with only host_header set (valid)
func testAccHealthcheckConfig_httpWithHostHeaderOnly(name, hostHeader string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path        = "/health"
    host_header = %[2]q
  }
}
`, name, hostHeader)
}

// testAccHealthcheckConfig_httpWithOriginServerName creates HTTP healthcheck with only use_origin_server_name set (valid)
func testAccHealthcheckConfig_httpWithOriginServerName(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path = "/health"
    use_origin_server_name {}
  }
}
`, name)
}

// testAccHealthcheckConfig_httpWithBothConflict creates HTTP healthcheck with BOTH host_header AND use_origin_server_name (INVALID - should fail ConflictsWith)
func testAccHealthcheckConfig_httpWithBothConflict(name, hostHeader string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path                   = "/health"
    host_header            = %[2]q
    use_origin_server_name {}
  }
}
`, name, hostHeader)
}

// testAccHealthcheckConfig_httpWithNeither creates HTTP healthcheck with neither host_header nor use_origin_server_name (valid - defaults)
func testAccHealthcheckConfig_httpWithNeither(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path = "/health"
  }
}
`, name)
}

// testAccHealthcheckConfig_tcpWithCustomPayload creates TCP healthcheck with custom send_payload and expected_response values
func testAccHealthcheckConfig_tcpWithCustomPayload(name, sendPayload, expectedResponse string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  tcp_health_check {
    send_payload      = %[2]q
    expected_response = %[3]q
  }
}
`, name, sendPayload, expectedResponse)
}

// testAccHealthcheckConfig_udpIcmpBasic creates basic UDP/ICMP healthcheck (empty block)
func testAccHealthcheckConfig_udpIcmpBasic(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  udp_icmp_health_check {}
}
`, name)
}

// =============================================================================
// ConflictsWith VALID PATTERN TESTS
// =============================================================================

// TestAccHealthcheckResource_httpWithHostHeaderOnly tests HTTP healthcheck with only host_header (valid pattern)
func TestAccHealthcheckResource_httpWithHostHeaderOnly(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_httpWithHostHeaderOnly(rName, "example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.path", "/health"),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "example.com"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "http_health_check.use_http2"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccHealthcheckResource_httpWithOriginServerName tests HTTP healthcheck with only use_origin_server_name (valid pattern)
func TestAccHealthcheckResource_httpWithOriginServerName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_httpWithOriginServerName(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.path", "/health"),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.use_origin_server_name.%", "0"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts", "http_health_check.use_http2"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccHealthcheckResource_httpWithNeither tests HTTP healthcheck with neither host_header nor use_origin_server_name (valid - uses defaults)
func TestAccHealthcheckResource_httpWithNeither(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_httpWithNeither(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.path", "/health"),
				),
			},
		},
	})
}

// TestAccHealthcheckResource_tcpWithCustomPayload tests TCP healthcheck with custom send_payload and expected_response values
func TestAccHealthcheckResource_tcpWithCustomPayload(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_tcpWithCustomPayload(rName, "48454c4c4f", "574f524c44"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "tcp_health_check.send_payload", "48454c4c4f"),
					resource.TestCheckResourceAttr(resourceName, "tcp_health_check.expected_response", "574f524c44"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccHealthcheckResource_udpIcmpHealthCheck tests UDP/ICMP healthcheck (empty block)
func TestAccHealthcheckResource_udpIcmpHealthCheck(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_udpIcmpBasic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "udp_icmp_health_check.%", "0"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"timeouts"},
				ImportStateIdFunc:       testAccHealthcheckImportStateIdFunc(resourceName),
			},
		},
	})
}

// =============================================================================
// ConflictsWith ANTI-PATTERN TESTS (Should Fail Validation)
// =============================================================================

// TestAccHealthcheckResource_httpHostHeaderConflictsWithOriginServerName tests that using BOTH host_header AND use_origin_server_name fails with ConflictsWith error
func TestAccHealthcheckResource_httpHostHeaderConflictsWithOriginServerName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-test-hc")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccHealthcheckConfig_httpWithBothConflict(rName, "example.com"),
				ExpectError: regexp.MustCompile(`(?i)(conflict|mutually exclusive|cannot be specified together|host_header|use_origin_server_name)`),
			},
		},
	})
}

// =============================================================================
// ConflictsWith CRUD LIFECYCLE TESTS
// =============================================================================

// TestAccHealthcheckResource_switchHostHeaderToOriginServerName tests switching from host_header to use_origin_server_name
func TestAccHealthcheckResource_switchHostHeaderToOriginServerName(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with host_header
				Config: testAccHealthcheckConfig_httpWithHostHeaderOnly(rName, "example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "example.com"),
				),
			},
			{
				// Step 2: Update to use_origin_server_name (switching from host_header)
				Config: testAccHealthcheckConfig_httpWithOriginServerName(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.use_origin_server_name.%", "0"),
					resource.TestCheckNoResourceAttr(resourceName, "http_health_check.host_header"),
				),
			},
			{
				// Step 3: Switch back to host_header
				Config: testAccHealthcheckConfig_httpWithHostHeaderOnly(rName, "updated.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "updated.example.com"),
				),
			},
		},
	})
}

// TestAccHealthcheckResource_switchOriginServerNameToHostHeader tests switching from use_origin_server_name to host_header
func TestAccHealthcheckResource_switchOriginServerNameToHostHeader(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)
	acctest.SkipIfQuotaExhausted(t, "system", "healthcheck")

	rName := acctest.RandomName("tf-acc-test-hc")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with use_origin_server_name
				Config: testAccHealthcheckConfig_httpWithOriginServerName(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.use_origin_server_name.%", "0"),
				),
			},
			{
				// Step 2: Update to host_header (switching from use_origin_server_name)
				Config: testAccHealthcheckConfig_httpWithHostHeaderOnly(rName, "myhost.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "myhost.example.com"),
				),
			},
		},
	})
}

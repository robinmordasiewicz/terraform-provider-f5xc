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
// MATRIX TEST: Healthcheck HTTP with all options
// =============================================================================
func TestAccHealthcheckResource_httpFullOptions(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-hc-http")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_httpFullOptions(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.path", "/api/health"),
					resource.TestCheckResourceAttr(resourceName, "http_health_check.host_header", "api.example.com"),
					resource.TestCheckResourceAttr(resourceName, "healthy_threshold", "3"),
					resource.TestCheckResourceAttr(resourceName, "unhealthy_threshold", "1"),
					resource.TestCheckResourceAttr(resourceName, "interval", "15"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "3"),
					resource.TestCheckResourceAttr(resourceName, "jitter_percent", "30"),
				),
			},
		},
	})
}

// =============================================================================
// MATRIX TEST: Healthcheck TCP with payload
// =============================================================================
func TestAccHealthcheckResource_tcpWithPayload(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-hc-tcp")
	resourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckHealthcheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcheckConfig_tcpWithPayload(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckHealthcheckExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					// Values are hex-encoded as per F5 XC API requirements
					resource.TestCheckResourceAttr(resourceName, "tcp_health_check.send_payload", "48454c4c4f"),
					resource.TestCheckResourceAttr(resourceName, "tcp_health_check.expected_response", "4f4b"),
				),
			},
		},
	})
}

// =============================================================================
// MATRIX TEST: Origin pool with healthcheck reference
// =============================================================================
func TestAccOriginPoolResource_withHealthcheck(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-op-hc")
	opResourceName := "f5xc_origin_pool.test"
	hcResourceName := "f5xc_healthcheck.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy: func(s *terraform.State) error {
			if err := acctest.CheckOriginPoolDestroyed(s); err != nil {
				return err
			}
			return acctest.CheckHealthcheckDestroyed(s)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_withHealthcheck(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(opResourceName),
					acctest.CheckHealthcheckExists(hcResourceName),
					resource.TestCheckResourceAttr(opResourceName, "name", rName),
					resource.TestCheckResourceAttrPair(opResourceName, "healthcheck.0.name", hcResourceName, "name"),
				),
			},
		},
	})
}

// =============================================================================
// MATRIX TEST: Origin pool with custom port
// =============================================================================
func TestAccOriginPoolResource_customPort(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.PreCheck(t)

	rName := acctest.RandomName("tf-acc-op-port")
	resourceName := "f5xc_origin_pool.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             acctest.CheckOriginPoolDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccOriginPoolConfig_tlsWithCustomPort(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					acctest.CheckOriginPoolExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "port", "8443"),
				),
			},
		},
	})
}

// =============================================================================
// CONFIG HELPERS
// =============================================================================

func testAccHealthcheckConfig_httpFullOptions(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  # Using F5 XC UI recommended values
  healthy_threshold   = 3   # UI default: 3 consecutive successes
  unhealthy_threshold = 1   # UI default: 1 consecutive failure
  timeout             = 3   # UI default: 3 seconds
  interval            = 15  # UI default: 15 seconds
  jitter_percent      = 30  # Recommended for production

  http_health_check {
    path        = "/api/health"
    host_header = "api.example.com"
  }
}
`, name)
}

func testAccHealthcheckConfig_tcpWithPayload(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  # Using F5 XC UI recommended values
  healthy_threshold   = 3   # UI default: 3 consecutive successes
  unhealthy_threshold = 1   # UI default: 1 consecutive failure
  timeout             = 3   # UI default: 3 seconds
  interval            = 15  # UI default: 15 seconds
  jitter_percent      = 30  # Recommended for production

  # TCP health check with hex-encoded payloads
  tcp_health_check {
    send_payload      = "48454c4c4f"  # "HELLO" in hex
    expected_response = "4f4b"        # "OK" in hex
  }
}
`, name)
}

func testAccOriginPoolConfig_withHealthcheck(name string) string {
	return fmt.Sprintf(`
resource "f5xc_healthcheck" "test" {
  name      = %[1]q
  namespace = "system"

  # Using F5 XC UI recommended values
  healthy_threshold   = 3   # UI default: 3 consecutive successes
  unhealthy_threshold = 1   # UI default: 1 consecutive failure
  timeout             = 3   # UI default: 3 seconds
  interval            = 15  # UI default: 15 seconds
  jitter_percent      = 30  # Recommended for production

  http_health_check {
    path        = "/health"
    host_header = "example.com"
  }
}

resource "f5xc_origin_pool" "test" {
  name       = %[1]q
  namespace  = "system"

  port = 443

  origin_servers {
    labels {}
    public_name {
      dns_name = "example.com"
    }
  }

  healthcheck {
    name      = f5xc_healthcheck.test.name
    namespace = f5xc_healthcheck.test.namespace
  }

  no_tls {}
  same_as_endpoint_port {}
}
`, name)
}

func testAccOriginPoolConfig_tlsWithCustomPort(name string) string {
	return fmt.Sprintf(`
resource "f5xc_origin_pool" "test" {
  name       = %[1]q
  namespace  = "system"

  port = 8443

  origin_servers {
    labels {}
    public_name {
      dns_name = "secure.example.com"
    }
  }

  # Using no_tls for now since TLS configuration requires additional setup
  no_tls {}
  same_as_endpoint_port {}
}
`, name)
}

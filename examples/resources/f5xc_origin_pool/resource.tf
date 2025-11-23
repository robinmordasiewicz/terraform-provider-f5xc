# Origin Pool Resource Example
# Manages a OriginPool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

# Basic Origin Pool configuration
resource "f5xc_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Origin Pool specific configuration
  origin_servers {
    public_ip {
      ip = "203.0.113.10"
    }
  }

  port = 443

  # Use TLS to origin
  use_tls {
    use_host_header_as_sni {}
    tls_config {
      default_security {}
    }
    skip_server_verification {}
  }

  # Endpoint selection
  endpoint_selection     = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "LB_OVERRIDE"
}

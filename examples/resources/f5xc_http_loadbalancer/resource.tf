# Http Loadbalancer Resource Example
# Manages a HTTPLoadBalancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with advanced routing and security.

# Basic Http Loadbalancer configuration
resource "f5xc_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # HTTP Load Balancer specific configuration
  domains = ["app.example.com"]

  # Advertise on public internet
  advertise_on_internet {
    default_vip {}
  }

  # Default origin server
  default_route_pools {
    pool {
      name      = "example-origin-pool"
      namespace = "system"
    }
    weight   = 1
    priority = 1
  }

  # Enable HTTP to HTTPS redirect
  http {
    dns_volterra_managed = true
  }

  # Disable rate limiting by default
  disable_rate_limit {}

  # No WAF by default
  disable_waf {}
}

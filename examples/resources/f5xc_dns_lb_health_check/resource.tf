# Dns Lb Health Check Resource Example
# Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a error. in F5 Distributed Cloud.

# Basic Dns Lb Health Check configuration
resource "f5xc_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: http_health_check, https_health_check, icmp_healt...
  http_health_check {
    # Configure http_health_check settings
  }
  # HTTP Health Check.
  https_health_check {
    # Configure https_health_check settings
  }
  # Empty. This can be used for messages where no values are ...
  icmp_health_check {
    # Configure icmp_health_check settings
  }
}

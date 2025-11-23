# Dns Lb Pool Resource Example
# Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error. in F5 Distributed Cloud.

# Basic Dns Lb Pool configuration
resource "f5xc_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # [OneOf: a_pool, aaaa_pool, cname_pool, mx_pool, srv_pool]...
    a_pool {
      # Configure a_pool settings
    }
    # Empty. This can be used for messages where no values are ...
    disable_health_check {
      # Configure disable_health_check settings
    }
    # Object reference. This type establishes a direct referenc...
    health_check {
      # Configure health_check settings
    }
}

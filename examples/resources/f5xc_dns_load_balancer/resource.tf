# DNS Load Balancer Resource Example
# [Category: Load Balancing] [Namespace: required] Manages DNS Load Balancer in a given namespace. If one already exist it will give a error. in F5 Distributed Cloud.

# Basic DNS Load Balancer configuration
resource "f5xc_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # DNS Load Balancer configuration
  record_type = "A"

  # DNS zone reference
  dns_zone {
    name      = "example-dns-zone"
    namespace = "staging"
  }

  # Rule-based load balancing
  rule_list {
    rules {
      geo_location_set {
        name      = "us-geo"
        namespace = "shared"
      }
      pool {
        name      = "us-pool"
        namespace = "staging"
      }
      score = 100
    }
  }
}

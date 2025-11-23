# Dns Load Balancer Resource Example
# Manages DNS Load Balancer in a given namespace. If one already exist it will give a error. in F5 Distributed Cloud.

# Basic Dns Load Balancer configuration
resource "f5xc_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"

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
    namespace = "system"
  }

  # Rule-based load balancing
  rule_list {
    rules {
      geo_location_set {
        name      = "us-geo"
        namespace = "system"
      }
      pool {
        name      = "us-pool"
        namespace = "system"
      }
      score = 100
    }
  }
}

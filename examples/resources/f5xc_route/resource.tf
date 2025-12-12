# Route Resource Example
# [Category: Load Balancing] [Namespace: required] [DependsOn: namespace, http_loadbalancer] Manages a Route resource in F5 Distributed Cloud for defining traffic routing rules for load balancers.

# Basic Route configuration
resource "f5xc_route" "example" {
  name      = "example-route"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Route configuration
  routes {
    match {
      path {
        prefix = "/api/"
      }
    }
    route_destination {
      destinations {
        cluster {
          name      = "api-cluster"
          namespace = "system"
        }
        weight = 100
      }
    }
  }
}

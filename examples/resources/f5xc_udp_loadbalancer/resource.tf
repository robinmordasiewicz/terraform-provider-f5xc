# UDP Loadbalancer Resource Example
# Manages a UDPLoadBalancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.

# Basic UDP Loadbalancer configuration
resource "f5xc_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # UDP Load Balancer configuration
  listen_port = 53

  # Advertise on public internet
  advertise_on_internet {
    default_vip {}
  }

  # Origin pools
  origin_pools_weights {
    pool {
      name      = "dns-pool"
      namespace = "staging"
    }
    weight = 1
  }

  # DNS for UDP load balancer
  dns_volterra_managed = true
}

# Udp Loadbalancer Resource Example
# Manages a UDPLoadBalancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.

# Basic Udp Loadbalancer configuration
resource "f5xc_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "system"

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
      namespace = "system"
    }
    weight = 1
  }

  # DNS for UDP load balancer
  dns_volterra_managed = true
}

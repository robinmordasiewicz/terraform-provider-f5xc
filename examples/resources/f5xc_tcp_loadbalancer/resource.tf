# Tcp Loadbalancer Resource Example
# Manages a TCPLoadBalancer resource in F5 Distributed Cloud for load balancing TCP traffic across origin pools.

# Basic Tcp Loadbalancer configuration
resource "f5xc_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # TCP Load Balancer specific configuration
  listen_port = 8443

  # Advertise on public internet
  advertise_on_internet {
    default_vip {}
  }

  # Origin pools
  origin_pools_weights {
    pool {
      name      = "example-tcp-pool"
      namespace = "system"
    }
    weight = 1
  }

  # DNS for TCP load balancer
  dns_volterra_managed = true

  # No retract cluster by default
  retract_cluster {}
}

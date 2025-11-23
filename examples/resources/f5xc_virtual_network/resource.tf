# Virtual Network Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

# Basic Virtual Network configuration
resource "f5xc_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Virtual Network configuration
  site_local_network {}

  # DHCP range for the network
  srv6_network {
    enterprise_network {
      srv6_network_ns_params {
        namespace = "system"
      }
    }
  }
}

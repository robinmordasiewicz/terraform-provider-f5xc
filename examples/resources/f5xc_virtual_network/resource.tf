# Virtual Network Resource Example
# Manages virtual network in given namespace. in F5 Distributed Cloud.

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

  // One of the arguments from this list "global_network legacy_type site_local_inside_network site_local_network srv6_network" must be set

  site_local_network {}

  // Static routes configuration (optional)
  static_routes {
    ip_prefixes = ["10.0.0.0/8"]

    // One of the arguments from this list "default_gateway ip_address node_interface" must be set

    default_gateway {}

    attrs = ["ROUTE_ATTR_INSTALL_FORWARDING"]
  }
}

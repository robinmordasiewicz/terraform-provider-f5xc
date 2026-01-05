# Network Connector Resource Example
# Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by users in system namespace. configuration.

# Basic Network Connector configuration
resource "f5xc_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Network Connector configuration
  # Direct connection
  sli_to_global_dr {
    global_vn {
      name      = "global-network"
      namespace = "system"
    }
  }

  # Disable forward proxy
  disable_forward_proxy {}
}

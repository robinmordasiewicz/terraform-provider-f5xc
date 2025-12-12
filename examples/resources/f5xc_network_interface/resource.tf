# Network Interface Resource Example
# [Category: Networking] [Namespace: required] Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device. it is created by users in system namespace. configuration.

# Basic Network Interface configuration
resource "f5xc_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: dedicated_interface, dedicated_management_interfa...
  dedicated_interface {
    # Configure dedicated_interface settings
  }
  # Empty. This can be used for messages where no values are ...
  cluster {
    # Configure cluster settings
  }
  # Empty. This can be used for messages where no values are ...
  is_primary {
    # Configure is_primary settings
  }
}

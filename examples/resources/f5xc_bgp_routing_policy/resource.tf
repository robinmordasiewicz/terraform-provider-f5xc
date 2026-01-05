# BGP Routing Policy Resource Example
# Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied. these rules help contol routes which are imported or exported to bgp peers. configuration.

# Basic BGP Routing Policy configuration
resource "f5xc_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Rules. A BGP Routing policy is composed of one or more ru...
  rules {
    # Configure rules settings
  }
  # BGP Route Action. Action to be enforced if the BGP route ...
  action {
    # Configure action settings
  }
  # Enable this option
  aggregate {
    # Configure aggregate settings
  }
}

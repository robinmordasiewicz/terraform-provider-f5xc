# Policy Based Routing Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification. configuration.

# Basic Policy Based Routing configuration
resource "f5xc_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: forward_proxy_pbr, network_pbr] L3/L4 routing rul...
  forward_proxy_pbr {
    # Configure forward_proxy_pbr settings
  }
  # L3/L4 routing rules. Network(L3/L4) routing policy rules.
  forward_proxy_pbr_rules {
    # Configure forward_proxy_pbr_rules settings
  }
  # Enable this option
  all_destinations {
    # Configure all_destinations settings
  }
}

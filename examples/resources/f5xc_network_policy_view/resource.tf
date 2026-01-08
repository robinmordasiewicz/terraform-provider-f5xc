# Network Policy View Resource Example
# Manages a Network Policy View resource in F5 Distributed Cloud for network policy view specification. configuration.

# Basic Network Policy View configuration
resource "f5xc_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Ordered list of rules applied to connections from policy ...
  egress_rules {
    # Configure egress_rules settings
  }
  # Network Policy Rule Advanced Action provides additional O...
  adv_action {
    # Configure adv_action settings
  }
  # Can be used for messages where no values are needed.
  all_tcp_traffic {
    # Configure all_tcp_traffic settings
  }
}

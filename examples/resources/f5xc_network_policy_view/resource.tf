# Network Policy View Resource Example
# Manages a NetworkPolicyView resource in F5 Distributed Cloud for network policy view configuration.

# Basic Network Policy View configuration
resource "f5xc_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Egress Rules. Ordered list of rules applied to connection...
  egress_rules {
    # Configure egress_rules settings
  }
  # Network Policy Rule Advanced Action. Network Policy Rule ...
  adv_action {
    # Configure adv_action settings
  }
  # Empty. This can be used for messages where no values are ...
  all_tcp_traffic {
    # Configure all_tcp_traffic settings
  }
}

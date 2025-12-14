# Infraprotect Firewall Rule Resource Example
# [Namespace: required] Manages DDoS transit Firewall Rule in F5 Distributed Cloud.

# Basic Infraprotect Firewall Rule configuration
resource "f5xc_infraprotect_firewall_rule" "example" {
  name      = "example-infraprotect-firewall-rule"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: action_allow, action_deny] Enable this option
  action_allow {
    # Configure action_allow settings
  }
  # Enable this option
  action_deny {
    # Configure action_deny settings
  }
  # [OneOf: destination_prefix_all, destination_prefix_single...
  destination_prefix_all {
    # Configure destination_prefix_all settings
  }
}

# Infraprotect Firewall Rule Resource Example
# Manages DDoS transit Firewall Rule in F5 Distributed Cloud.

# Basic Infraprotect Firewall Rule configuration
resource "f5xc_infraprotect_firewall_rule" "example" {
  name      = "example-infraprotect-firewall-rule"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # [OneOf: action_allow, action_deny] Empty. This can be use...
    action_allow {
      # Configure action_allow settings
    }
    # Empty. This can be used for messages where no values are ...
    action_deny {
      # Configure action_deny settings
    }
    # [OneOf: destination_prefix_all, destination_prefix_single...
    destination_prefix_all {
      # Configure destination_prefix_all settings
    }
}

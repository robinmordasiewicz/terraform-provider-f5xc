# Network Policy Rule Resource Example
# Manages network policy rule with configured parameters in specified namespace. in F5 Distributed Cloud.

# Basic Network Policy Rule configuration
resource "f5xc_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Network Policy Rule Advanced Action provides additional O...
  advanced_action {
    # Configure advanced_action settings
  }
  # [OneOf: ip_prefix_set, prefix, prefix_selector] List of r...
  ip_prefix_set {
    # Configure ip_prefix_set settings
  }
  # List of references to ip_prefix_set objects.
  ref {
    # Configure ref settings
  }
}

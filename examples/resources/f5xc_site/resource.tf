# Site Resource Example
# Manages a Site resource in F5 Distributed Cloud for secure mesh site specification. configuration.

# Basic Site configuration
resource "f5xc_site" "example" {
  name      = "example-site"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: active_enhanced_firewall_policies, no_network_pol...
  active_enhanced_firewall_policies {
    # Configure active_enhanced_firewall_policies settings
  }
  # Enhanced Firewall Policy. Ordered List of Enhanced Firewa...
  enhanced_firewall_policies {
    # Configure enhanced_firewall_policies settings
  }
  # [OneOf: active_forward_proxy_policies, no_forward_proxy; ...
  active_forward_proxy_policies {
    # Configure active_forward_proxy_policies settings
  }
}

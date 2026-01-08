# Network Firewall Resource Example
# Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace. configuration.

# Basic Network Firewall configuration
resource "f5xc_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: active_enhanced_firewall_policies, active_network...
  active_enhanced_firewall_policies {
    # Configure active_enhanced_firewall_policies settings
  }
  # Ordered List of Enhanced Firewall Policies active .
  enhanced_firewall_policies {
    # Configure enhanced_firewall_policies settings
  }
  # [OneOf: active_fast_acls, disable_fast_acl; Default: disa...
  active_fast_acls {
    # Configure active_fast_acls settings
  }
}

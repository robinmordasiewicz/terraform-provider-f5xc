# Securemesh Site V2 Resource Example
# Manages a SecuremeshSiteV2 resource in F5 Distributed Cloud for deploying secure mesh edge sites with enhanced security and networking features.

# Basic Securemesh Site V2 configuration
resource "f5xc_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"

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
  # [OneOf: active_forward_proxy_policies, no_forward_proxy] ...
  active_forward_proxy_policies {
    # Configure active_forward_proxy_policies settings
  }
}

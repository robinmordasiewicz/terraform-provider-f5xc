# Enhanced Firewall Policy Resource Example
# Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification. configuration.

# Basic Enhanced Firewall Policy configuration
resource "f5xc_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Enhanced Firewall Policy configuration
  rule_list {
    rules {
      metadata {
        name = "allow-web-traffic"
      }
      allow {}
      advanced_action {
        action = "LOG"
      }
      source_prefix_list {
        ip_prefix_set {
          name      = "trusted-ips"
          namespace = "shared"
        }
      }
      all_traffic {}
    }
  }
}

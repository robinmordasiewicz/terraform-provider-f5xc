# Infraprotect Firewall Rule Group Resource Example
# Manages a InfraprotectFirewallRuleGroup resource in F5 Distributed Cloud for amends a ddos transit firewall rule group configuration.

# Basic Infraprotect Firewall Rule Group configuration
resource "f5xc_infraprotect_firewall_rule_group" "example" {
  name      = "example-infraprotect-firewall-rule-group"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

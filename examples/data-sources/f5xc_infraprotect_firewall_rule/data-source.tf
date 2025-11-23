# Infraprotect Firewall Rule Data Source Example
# Retrieves information about an existing Infraprotect Firewall Rule

# Look up an existing Infraprotect Firewall Rule by name
data "f5xc_infraprotect_firewall_rule" "example" {
  name      = "example-infraprotect-firewall-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "infraprotect_firewall_rule_id" {
#   value = data.f5xc_infraprotect_firewall_rule.example.id
# }

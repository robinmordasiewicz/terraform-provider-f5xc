# Infraprotect Firewall Rule Group Data Source Example
# Retrieves information about an existing Infraprotect Firewall Rule Group

# Look up an existing Infraprotect Firewall Rule Group by name
data "f5xc_infraprotect_firewall_rule_group" "example" {
  name      = "example-infraprotect-firewall-rule-group"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "infraprotect_firewall_rule_group_id" {
#   value = data.f5xc_infraprotect_firewall_rule_group.example.id
# }

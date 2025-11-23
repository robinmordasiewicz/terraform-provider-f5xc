# Network Policy Rule Data Source Example
# Retrieves information about an existing Network Policy Rule

# Look up an existing Network Policy Rule by name
data "f5xc_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_policy_rule_id" {
#   value = data.f5xc_network_policy_rule.example.id
# }

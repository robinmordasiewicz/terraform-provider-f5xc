# Service Policy Rule Data Source Example
# Retrieves information about an existing Service Policy Rule

# Look up an existing Service Policy Rule by name
data "f5xc_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "service_policy_rule_id" {
#   value = data.f5xc_service_policy_rule.example.id
# }

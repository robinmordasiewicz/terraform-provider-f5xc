# Secret Policy Rule Data Source Example
# Retrieves information about an existing Secret Policy Rule

# Look up an existing Secret Policy Rule by name
data "f5xc_secret_policy_rule" "example" {
  name      = "example-secret-policy-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "secret_policy_rule_id" {
#   value = data.f5xc_secret_policy_rule.example.id
# }

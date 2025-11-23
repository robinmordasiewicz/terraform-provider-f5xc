# Infraprotect Deny List Rule Data Source Example
# Retrieves information about an existing Infraprotect Deny List Rule

# Look up an existing Infraprotect Deny List Rule by name
data "f5xc_infraprotect_deny_list_rule" "example" {
  name      = "example-infraprotect-deny-list-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "infraprotect_deny_list_rule_id" {
#   value = data.f5xc_infraprotect_deny_list_rule.example.id
# }

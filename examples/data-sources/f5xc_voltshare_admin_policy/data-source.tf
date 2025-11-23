# Voltshare Admin Policy Data Source Example
# Retrieves information about an existing Voltshare Admin Policy

# Look up an existing Voltshare Admin Policy by name
data "f5xc_voltshare_admin_policy" "example" {
  name      = "example-voltshare-admin-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "voltshare_admin_policy_id" {
#   value = data.f5xc_voltshare_admin_policy.example.id
# }

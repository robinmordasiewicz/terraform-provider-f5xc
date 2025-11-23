# Network Policy Data Source Example
# Retrieves information about an existing Network Policy

# Look up an existing Network Policy by name
data "f5xc_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_policy_id" {
#   value = data.f5xc_network_policy.example.id
# }

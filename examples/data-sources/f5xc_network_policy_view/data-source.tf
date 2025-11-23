# Network Policy View Data Source Example
# Retrieves information about an existing Network Policy View

# Look up an existing Network Policy View by name
data "f5xc_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_policy_view_id" {
#   value = data.f5xc_network_policy_view.example.id
# }

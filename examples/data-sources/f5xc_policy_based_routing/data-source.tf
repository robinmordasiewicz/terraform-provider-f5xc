# Policy Based Routing Data Source Example
# Retrieves information about an existing Policy Based Routing

# Look up an existing Policy Based Routing by name
data "f5xc_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "policy_based_routing_id" {
#   value = data.f5xc_policy_based_routing.example.id
# }

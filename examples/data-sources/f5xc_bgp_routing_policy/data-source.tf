# Bgp Routing Policy Data Source Example
# Retrieves information about an existing Bgp Routing Policy

# Look up an existing Bgp Routing Policy by name
data "f5xc_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "bgp_routing_policy_id" {
#   value = data.f5xc_bgp_routing_policy.example.id
# }

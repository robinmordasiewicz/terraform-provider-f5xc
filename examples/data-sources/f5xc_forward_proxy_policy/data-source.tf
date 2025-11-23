# Forward Proxy Policy Data Source Example
# Retrieves information about an existing Forward Proxy Policy

# Look up an existing Forward Proxy Policy by name
data "f5xc_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "forward_proxy_policy_id" {
#   value = data.f5xc_forward_proxy_policy.example.id
# }

# BGP Data Source Example
# Retrieves information about an existing BGP

# Look up an existing BGP by name
data "f5xc_bgp" "example" {
  name      = "example-bgp"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "bgp_id" {
#   value = data.f5xc_bgp.example.id
# }

# Bgp Data Source Example
# Retrieves information about an existing Bgp

# Look up an existing Bgp by name
data "f5xc_bgp" "example" {
  name      = "example-bgp"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "bgp_id" {
#   value = data.f5xc_bgp.example.id
# }

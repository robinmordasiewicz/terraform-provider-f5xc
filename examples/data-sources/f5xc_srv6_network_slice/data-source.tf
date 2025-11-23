# Srv6 Network Slice Data Source Example
# Retrieves information about an existing Srv6 Network Slice

# Look up an existing Srv6 Network Slice by name
data "f5xc_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "srv6_network_slice_id" {
#   value = data.f5xc_srv6_network_slice.example.id
# }

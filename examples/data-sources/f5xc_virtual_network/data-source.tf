# Virtual Network Data Source Example
# Retrieves information about an existing Virtual Network

# Look up an existing Virtual Network by name
data "f5xc_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "virtual_network_id" {
#   value = data.f5xc_virtual_network.example.id
# }

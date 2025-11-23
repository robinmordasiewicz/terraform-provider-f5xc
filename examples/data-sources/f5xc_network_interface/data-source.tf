# Network Interface Data Source Example
# Retrieves information about an existing Network Interface

# Look up an existing Network Interface by name
data "f5xc_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_interface_id" {
#   value = data.f5xc_network_interface.example.id
# }

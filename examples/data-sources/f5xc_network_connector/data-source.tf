# Network Connector Data Source Example
# Retrieves information about an existing Network Connector

# Look up an existing Network Connector by name
data "f5xc_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_connector_id" {
#   value = data.f5xc_network_connector.example.id
# }

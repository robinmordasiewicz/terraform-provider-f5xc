# External Connector Data Source Example
# Retrieves information about an existing External Connector

# Look up an existing External Connector by name
data "f5xc_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "external_connector_id" {
#   value = data.f5xc_external_connector.example.id
# }

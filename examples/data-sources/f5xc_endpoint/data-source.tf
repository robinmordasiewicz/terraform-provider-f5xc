# Endpoint Data Source Example
# Retrieves information about an existing Endpoint

# Look up an existing Endpoint by name
data "f5xc_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "endpoint_id" {
#   value = data.f5xc_endpoint.example.id
# }

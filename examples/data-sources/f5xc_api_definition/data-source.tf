# Api Definition Data Source Example
# Retrieves information about an existing Api Definition

# Look up an existing Api Definition by name
data "f5xc_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_definition_id" {
#   value = data.f5xc_api_definition.example.id
# }

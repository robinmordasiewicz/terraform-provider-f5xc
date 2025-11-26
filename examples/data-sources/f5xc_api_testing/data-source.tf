# API Testing Data Source Example
# Retrieves information about an existing API Testing

# Look up an existing API Testing by name
data "f5xc_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_testing_id" {
#   value = data.f5xc_api_testing.example.id
# }

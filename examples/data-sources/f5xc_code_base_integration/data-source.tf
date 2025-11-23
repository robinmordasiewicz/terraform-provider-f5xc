# Code Base Integration Data Source Example
# Retrieves information about an existing Code Base Integration

# Look up an existing Code Base Integration by name
data "f5xc_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "code_base_integration_id" {
#   value = data.f5xc_code_base_integration.example.id
# }

# API Credential Data Source Example
# Retrieves information about an existing API Credential

# Look up an existing API Credential by name
data "f5xc_api_credential" "example" {
  name      = "example-api-credential"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_credential_id" {
#   value = data.f5xc_api_credential.example.id
# }

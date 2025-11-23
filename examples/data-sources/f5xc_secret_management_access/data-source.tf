# Secret Management Access Data Source Example
# Retrieves information about an existing Secret Management Access

# Look up an existing Secret Management Access by name
data "f5xc_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "secret_management_access_id" {
#   value = data.f5xc_secret_management_access.example.id
# }

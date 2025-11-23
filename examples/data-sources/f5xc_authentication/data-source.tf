# Authentication Data Source Example
# Retrieves information about an existing Authentication

# Look up an existing Authentication by name
data "f5xc_authentication" "example" {
  name      = "example-authentication"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "authentication_id" {
#   value = data.f5xc_authentication.example.id
# }

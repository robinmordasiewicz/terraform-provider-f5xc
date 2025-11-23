# Oidc Provider Data Source Example
# Retrieves information about an existing Oidc Provider

# Look up an existing Oidc Provider by name
data "f5xc_oidc_provider" "example" {
  name      = "example-oidc-provider"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "oidc_provider_id" {
#   value = data.f5xc_oidc_provider.example.id
# }

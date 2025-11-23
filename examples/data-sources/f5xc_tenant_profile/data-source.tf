# Tenant Profile Data Source Example
# Retrieves information about an existing Tenant Profile

# Look up an existing Tenant Profile by name
data "f5xc_tenant_profile" "example" {
  name      = "example-tenant-profile"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tenant_profile_id" {
#   value = data.f5xc_tenant_profile.example.id
# }

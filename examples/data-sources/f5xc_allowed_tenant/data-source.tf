# Allowed Tenant Data Source Example
# Retrieves information about an existing Allowed Tenant

# Look up an existing Allowed Tenant by name
data "f5xc_allowed_tenant" "example" {
  name      = "example-allowed-tenant"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "allowed_tenant_id" {
#   value = data.f5xc_allowed_tenant.example.id
# }

# Child Tenant Data Source Example
# Retrieves information about an existing Child Tenant

# Look up an existing Child Tenant by name
data "f5xc_child_tenant" "example" {
  name      = "example-child-tenant"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "child_tenant_id" {
#   value = data.f5xc_child_tenant.example.id
# }

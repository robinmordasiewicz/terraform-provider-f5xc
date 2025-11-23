# Child Tenant Manager Data Source Example
# Retrieves information about an existing Child Tenant Manager

# Look up an existing Child Tenant Manager by name
data "f5xc_child_tenant_manager" "example" {
  name      = "example-child-tenant-manager"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "child_tenant_manager_id" {
#   value = data.f5xc_child_tenant_manager.example.id
# }

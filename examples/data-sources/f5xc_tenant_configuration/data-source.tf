# Tenant Configuration Data Source Example
# Retrieves information about an existing Tenant Configuration

# Look up an existing Tenant Configuration by name
data "f5xc_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tenant_configuration_id" {
#   value = data.f5xc_tenant_configuration.example.id
# }

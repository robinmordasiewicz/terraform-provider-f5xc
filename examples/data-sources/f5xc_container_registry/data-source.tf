# Container Registry Data Source Example
# Retrieves information about an existing Container Registry

# Look up an existing Container Registry by name
data "f5xc_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "container_registry_id" {
#   value = data.f5xc_container_registry.example.id
# }

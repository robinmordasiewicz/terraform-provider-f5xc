# Namespace Data Source Example
# Retrieves information about an existing Namespace

# Look up an existing Namespace by name
data "f5xc_namespace" "example" {
  name      = "example-namespace"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "namespace_id" {
#   value = data.f5xc_namespace.example.id
# }

# Example: Create resources in a namespace discovered via data source
# resource "f5xc_origin_pool" "example" {
#   name      = "example-pool"
#   namespace = data.f5xc_namespace.example.name
#   # ... other configuration
# }

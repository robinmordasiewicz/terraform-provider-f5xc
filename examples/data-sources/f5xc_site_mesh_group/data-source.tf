# Site Mesh Group Data Source Example
# Retrieves information about an existing Site Mesh Group

# Look up an existing Site Mesh Group by name
data "f5xc_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "site_mesh_group_id" {
#   value = data.f5xc_site_mesh_group.example.id
# }

# Securemesh Site Data Source Example
# Retrieves information about an existing Securemesh Site

# Look up an existing Securemesh Site by name
data "f5xc_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "securemesh_site_id" {
#   value = data.f5xc_securemesh_site.example.id
# }

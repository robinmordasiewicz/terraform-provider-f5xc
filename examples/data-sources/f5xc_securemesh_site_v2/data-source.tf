# Securemesh Site V2 Data Source Example
# Retrieves information about an existing Securemesh Site V2

# Look up an existing Securemesh Site V2 by name
data "f5xc_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "securemesh_site_v2_id" {
#   value = data.f5xc_securemesh_site_v2.example.id
# }

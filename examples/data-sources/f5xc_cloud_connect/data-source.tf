# Cloud Connect Data Source Example
# Retrieves information about an existing Cloud Connect

# Look up an existing Cloud Connect by name
data "f5xc_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_connect_id" {
#   value = data.f5xc_cloud_connect.example.id
# }

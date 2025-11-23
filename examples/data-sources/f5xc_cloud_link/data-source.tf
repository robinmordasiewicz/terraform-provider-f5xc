# Cloud Link Data Source Example
# Retrieves information about an existing Cloud Link

# Look up an existing Cloud Link by name
data "f5xc_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_link_id" {
#   value = data.f5xc_cloud_link.example.id
# }

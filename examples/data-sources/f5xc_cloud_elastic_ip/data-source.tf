# Cloud Elastic IP Data Source Example
# Retrieves information about an existing Cloud Elastic IP

# Look up an existing Cloud Elastic IP by name
data "f5xc_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_elastic_ip_id" {
#   value = data.f5xc_cloud_elastic_ip.example.id
# }

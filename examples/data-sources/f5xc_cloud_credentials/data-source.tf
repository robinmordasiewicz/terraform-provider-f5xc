# Cloud Credentials Data Source Example
# Retrieves information about an existing Cloud Credentials

# Look up an existing Cloud Credentials by name
data "f5xc_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_credentials_id" {
#   value = data.f5xc_cloud_credentials.example.id
# }

# Example: Reference cloud credentials in site configuration
# resource "f5xc_aws_vpc_site" "example" {
#   name      = "example-aws-site"
#   namespace = "system"
#
#   aws_cred {
#     name      = data.f5xc_cloud_credentials.example.name
#     namespace = data.f5xc_cloud_credentials.example.namespace
#   }
# }

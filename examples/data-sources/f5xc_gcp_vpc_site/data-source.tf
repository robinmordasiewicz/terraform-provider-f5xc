# GCP VPC Site Data Source Example
# Retrieves information about an existing GCP VPC Site

# Look up an existing GCP VPC Site by name
data "f5xc_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "gcp_vpc_site_id" {
#   value = data.f5xc_gcp_vpc_site.example.id
# }

# Example: Reference cloud site for advertising load balancer
# resource "f5xc_http_loadbalancer" "example" {
#   name      = "site-advertised-lb"
#   namespace = "system"
#
#   advertise_custom {
#     advertise_where {
#       site {
#         site {
#           name      = data.f5xc_gcp_vpc_site.example.name
#           namespace = data.f5xc_gcp_vpc_site.example.namespace
#         }
#       }
#     }
#   }
# }

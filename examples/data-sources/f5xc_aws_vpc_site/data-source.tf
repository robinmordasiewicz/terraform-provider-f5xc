# AWS VPC Site Data Source Example
# Retrieves information about an existing AWS VPC Site

# Look up an existing AWS VPC Site by name
data "f5xc_aws_vpc_site" "example" {
  name      = "example-aws-vpc-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "aws_vpc_site_id" {
#   value = data.f5xc_aws_vpc_site.example.id
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
#           name      = data.f5xc_aws_vpc_site.example.name
#           namespace = data.f5xc_aws_vpc_site.example.namespace
#         }
#       }
#     }
#   }
# }

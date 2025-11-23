# Certificate Data Source Example
# Retrieves information about an existing Certificate

# Look up an existing Certificate by name
data "f5xc_certificate" "example" {
  name      = "example-certificate"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "certificate_id" {
#   value = data.f5xc_certificate.example.id
# }

# Example: Reference certificate in HTTPS configuration
# resource "f5xc_http_loadbalancer" "example" {
#   name      = "https-lb"
#   namespace = "system"
#
#   https {
#     tls_cert_params {
#       certificates {
#         name      = data.f5xc_certificate.example.name
#         namespace = data.f5xc_certificate.example.namespace
#       }
#     }
#   }
# }

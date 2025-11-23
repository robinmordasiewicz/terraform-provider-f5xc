# Azure Vnet Site Data Source Example
# Retrieves information about an existing Azure Vnet Site

# Look up an existing Azure Vnet Site by name
data "f5xc_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "azure_vnet_site_id" {
#   value = data.f5xc_azure_vnet_site.example.id
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
#           name      = data.f5xc_azure_vnet_site.example.name
#           namespace = data.f5xc_azure_vnet_site.example.namespace
#         }
#       }
#     }
#   }
# }

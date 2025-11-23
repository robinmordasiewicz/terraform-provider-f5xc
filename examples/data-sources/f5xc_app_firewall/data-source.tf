# App Firewall Data Source Example
# Retrieves information about an existing App Firewall

# Look up an existing App Firewall by name
data "f5xc_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "app_firewall_id" {
#   value = data.f5xc_app_firewall.example.id
# }

# Example: Reference WAF in HTTP load balancer
# resource "f5xc_http_loadbalancer" "example" {
#   name      = "protected-lb"
#   namespace = "system"
#
#   app_firewall {
#     name      = data.f5xc_app_firewall.example.name
#     namespace = data.f5xc_app_firewall.example.namespace
#   }
# }

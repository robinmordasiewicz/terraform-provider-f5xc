# Http Loadbalancer Data Source Example
# Retrieves information about an existing Http Loadbalancer

# Look up an existing Http Loadbalancer by name
data "f5xc_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "http_loadbalancer_id" {
#   value = data.f5xc_http_loadbalancer.example.id
# }

# Example: Reference in another load balancer configuration
# resource "f5xc_service_policy" "example" {
#   name      = "policy-for-lb"
#   namespace = "system"
#
#   # Use the load balancer's domains
#   # domain = data.f5xc_http_loadbalancer.example.domains[0]
# }

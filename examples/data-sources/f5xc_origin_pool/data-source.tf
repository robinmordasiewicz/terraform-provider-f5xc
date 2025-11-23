# Origin Pool Data Source Example
# Retrieves information about an existing Origin Pool

# Look up an existing Origin Pool by name
data "f5xc_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "origin_pool_id" {
#   value = data.f5xc_origin_pool.example.id
# }

# Example: Use origin pool data in HTTP load balancer
# resource "f5xc_http_loadbalancer" "example" {
#   name      = "example-lb"
#   namespace = "system"
#   domains   = ["app.example.com"]
#
#   default_route_pools {
#     pool {
#       name      = data.f5xc_origin_pool.example.name
#       namespace = data.f5xc_origin_pool.example.namespace
#     }
#   }
# }

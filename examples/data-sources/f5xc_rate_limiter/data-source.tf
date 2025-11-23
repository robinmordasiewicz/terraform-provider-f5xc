# Rate Limiter Data Source Example
# Retrieves information about an existing Rate Limiter

# Look up an existing Rate Limiter by name
data "f5xc_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "rate_limiter_id" {
#   value = data.f5xc_rate_limiter.example.id
# }

# Example: Reference rate limiter in HTTP load balancer
# resource "f5xc_http_loadbalancer" "example" {
#   name      = "rate-limited-lb"
#   namespace = "system"
#
#   rate_limit {
#     rate_limiter {
#       name      = data.f5xc_rate_limiter.example.name
#       namespace = data.f5xc_rate_limiter.example.namespace
#     }
#   }
# }

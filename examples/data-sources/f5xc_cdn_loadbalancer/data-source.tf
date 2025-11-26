# CDN Loadbalancer Data Source Example
# Retrieves information about an existing CDN Loadbalancer

# Look up an existing CDN Loadbalancer by name
data "f5xc_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cdn_loadbalancer_id" {
#   value = data.f5xc_cdn_loadbalancer.example.id
# }

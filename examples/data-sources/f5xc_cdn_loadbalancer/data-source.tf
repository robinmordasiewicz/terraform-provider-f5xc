# Cdn Loadbalancer Data Source Example
# Retrieves information about an existing Cdn Loadbalancer

# Look up an existing Cdn Loadbalancer by name
data "f5xc_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cdn_loadbalancer_id" {
#   value = data.f5xc_cdn_loadbalancer.example.id
# }

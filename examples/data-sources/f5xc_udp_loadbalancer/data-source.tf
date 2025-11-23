# Udp Loadbalancer Data Source Example
# Retrieves information about an existing Udp Loadbalancer

# Look up an existing Udp Loadbalancer by name
data "f5xc_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "udp_loadbalancer_id" {
#   value = data.f5xc_udp_loadbalancer.example.id
# }

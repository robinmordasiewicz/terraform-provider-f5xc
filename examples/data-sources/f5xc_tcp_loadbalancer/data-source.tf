# Tcp Loadbalancer Data Source Example
# Retrieves information about an existing Tcp Loadbalancer

# Look up an existing Tcp Loadbalancer by name
data "f5xc_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tcp_loadbalancer_id" {
#   value = data.f5xc_tcp_loadbalancer.example.id
# }

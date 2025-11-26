# TCP Loadbalancer Data Source Example
# Retrieves information about an existing TCP Loadbalancer

# Look up an existing TCP Loadbalancer by name
data "f5xc_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tcp_loadbalancer_id" {
#   value = data.f5xc_tcp_loadbalancer.example.id
# }

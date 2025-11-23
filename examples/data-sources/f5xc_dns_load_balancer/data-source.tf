# Dns Load Balancer Data Source Example
# Retrieves information about an existing Dns Load Balancer

# Look up an existing Dns Load Balancer by name
data "f5xc_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_load_balancer_id" {
#   value = data.f5xc_dns_load_balancer.example.id
# }

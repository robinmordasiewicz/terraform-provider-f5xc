# DNS LB Pool Data Source Example
# Retrieves information about an existing DNS LB Pool

# Look up an existing DNS LB Pool by name
data "f5xc_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_lb_pool_id" {
#   value = data.f5xc_dns_lb_pool.example.id
# }

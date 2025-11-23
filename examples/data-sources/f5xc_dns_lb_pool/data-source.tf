# Dns Lb Pool Data Source Example
# Retrieves information about an existing Dns Lb Pool

# Look up an existing Dns Lb Pool by name
data "f5xc_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_lb_pool_id" {
#   value = data.f5xc_dns_lb_pool.example.id
# }

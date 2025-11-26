# DNS LB Health Check Data Source Example
# Retrieves information about an existing DNS LB Health Check

# Look up an existing DNS LB Health Check by name
data "f5xc_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_lb_health_check_id" {
#   value = data.f5xc_dns_lb_health_check.example.id
# }

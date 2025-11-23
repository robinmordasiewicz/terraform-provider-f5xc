# Dns Lb Health Check Data Source Example
# Retrieves information about an existing Dns Lb Health Check

# Look up an existing Dns Lb Health Check by name
data "f5xc_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_lb_health_check_id" {
#   value = data.f5xc_dns_lb_health_check.example.id
# }

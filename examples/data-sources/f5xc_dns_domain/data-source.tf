# DNS Domain Data Source Example
# Retrieves information about an existing DNS Domain

# Look up an existing DNS Domain by name
data "f5xc_dns_domain" "example" {
  name      = "example-dns-domain"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_domain_id" {
#   value = data.f5xc_dns_domain.example.id
# }

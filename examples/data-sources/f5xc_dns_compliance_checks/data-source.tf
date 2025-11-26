# DNS Compliance Checks Data Source Example
# Retrieves information about an existing DNS Compliance Checks

# Look up an existing DNS Compliance Checks by name
data "f5xc_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_compliance_checks_id" {
#   value = data.f5xc_dns_compliance_checks.example.id
# }

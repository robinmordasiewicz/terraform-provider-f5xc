# Certificate Chain Data Source Example
# Retrieves information about an existing Certificate Chain

# Look up an existing Certificate Chain by name
data "f5xc_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "certificate_chain_id" {
#   value = data.f5xc_certificate_chain.example.id
# }

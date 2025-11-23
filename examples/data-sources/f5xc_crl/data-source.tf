# Crl Data Source Example
# Retrieves information about an existing Crl

# Look up an existing Crl by name
data "f5xc_crl" "example" {
  name      = "example-crl"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "crl_id" {
#   value = data.f5xc_crl.example.id
# }

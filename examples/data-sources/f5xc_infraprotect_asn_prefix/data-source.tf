# Infraprotect Asn Prefix Data Source Example
# Retrieves information about an existing Infraprotect Asn Prefix

# Look up an existing Infraprotect Asn Prefix by name
data "f5xc_infraprotect_asn_prefix" "example" {
  name      = "example-infraprotect-asn-prefix"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "infraprotect_asn_prefix_id" {
#   value = data.f5xc_infraprotect_asn_prefix.example.id
# }

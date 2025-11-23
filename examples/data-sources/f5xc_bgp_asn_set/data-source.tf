# Bgp Asn Set Data Source Example
# Retrieves information about an existing Bgp Asn Set

# Look up an existing Bgp Asn Set by name
data "f5xc_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "bgp_asn_set_id" {
#   value = data.f5xc_bgp_asn_set.example.id
# }

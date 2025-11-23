# Bgp Asn Set Resource Example
# Manages bgp_asn_set creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Bgp Asn Set configuration
resource "f5xc_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied. these rules help contol routes which are imported or exported to bgp peers. configuration.

# Basic BGP configuration
resource "f5xc_bgp" "example" {
  name      = "example-bgp"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # BGP configuration
  bgp_router_id = "192.168.1.1"

  bgp_peers {
    metadata {
      name = "upstream-peer"
    }
    spec {
      peer_asn     = 65000
      peer_address = "192.168.1.2"
    }
  }

  local_asn = 65001

  # Site reference
  site {
    name      = "example-site"
    namespace = "system"
  }
}

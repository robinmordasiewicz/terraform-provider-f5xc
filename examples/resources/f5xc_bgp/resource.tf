# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers. it is created by users in system namespace. configuration.

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

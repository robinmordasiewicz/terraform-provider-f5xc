# Infraprotect Tunnel Resource Example
# Manages DDoS transit tunnel in F5 Distributed Cloud.

# Basic Infraprotect Tunnel configuration
resource "f5xc_infraprotect_tunnel" "example" {
  name      = "example-infraprotect-tunnel"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Bandwidth Speed Configuration. Bandwidth max allowed
  bandwidth {
    # Configure bandwidth settings
  }
  # BGP. BGP information associated with a DDoS transit tunnel.
  bgp_information {
    # Configure bgp_information settings
  }
  # Object reference. This type establishes a direct referenc...
  asn {
    # Configure asn settings
  }
}

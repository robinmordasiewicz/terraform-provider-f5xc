# Infraprotect Asn Resource Example
# [Namespace: required] Manages DDoS transit ASN in F5 Distributed Cloud.

# Basic Infraprotect Asn configuration
resource "f5xc_infraprotect_asn" "example" {
  name      = "example-infraprotect-asn"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: bgp_session_disabled, bgp_session_enabled] Enable...
  bgp_session_disabled {
    # Configure bgp_session_disabled settings
  }
  # Enable this option
  bgp_session_enabled {
    # Configure bgp_session_enabled settings
  }
}

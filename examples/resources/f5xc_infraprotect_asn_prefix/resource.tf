# Infraprotect Asn Prefix Resource Example
# Manages DDoS transit Prefix in F5 Distributed Cloud.

# Basic Infraprotect Asn Prefix configuration
resource "f5xc_infraprotect_asn_prefix" "example" {
  name      = "example-infraprotect-asn-prefix"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # Object reference. This type establishes a direct referenc...
    asn {
      # Configure asn settings
    }
}

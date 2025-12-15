# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl configuration.

# Basic CRL configuration
resource "f5xc_crl" "example" {
  name      = "example-crl"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # HTTPAccessInfo.
  http_access {
    # Configure http_access settings
  }
}

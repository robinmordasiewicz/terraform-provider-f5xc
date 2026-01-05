# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

# Basic Certificate configuration
resource "f5xc_certificate" "example" {
  name      = "example-certificate"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Certificate configuration
  certificate_url = "string:///LS0tLS1CRUdJTi..."
  private_key {
    clear_secret_info {
      url = "string:///LS0tLS1CRUdJTi..."
    }
  }
}

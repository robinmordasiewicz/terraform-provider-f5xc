# Certificate Chain Resource Example
# Manages a CertificateChain resource in F5 Distributed Cloud for certificate chain configuration for TLS.

# Basic Certificate Chain configuration
resource "f5xc_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

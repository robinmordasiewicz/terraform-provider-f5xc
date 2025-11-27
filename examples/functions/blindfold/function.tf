# Encrypt a secret string using F5XC blindfold
#
# The blindfold function encrypts base64-encoded plaintext using F5 Distributed
# Cloud Secret Management. The encryption happens locally - your secret is never
# transmitted to F5XC during encryption.

# Example: Encrypt a password for use in origin pool authentication
locals {
  encrypted_password = provider::f5xc::blindfold(
    base64encode("my-secret-password"),
    "production-secrets-policy",
    "shared"
  )
}

# Example: Encrypt a TLS private key from a file
locals {
  encrypted_key = provider::f5xc::blindfold(
    base64encode(file("${path.module}/certs/private.key")),
    "tls-secrets-policy",
    "shared"
  )
}

# Example: Using the encrypted secret in a resource
resource "f5xc_http_loadbalancer" "example" {
  name      = "secure-lb"
  namespace = "production"

  domains = ["example.com"]

  https_auto_cert {
    tls_config {
      custom_security {
        private_key {
          blindfold_secret_info {
            location = provider::f5xc::blindfold(
              base64encode(file("${path.module}/certs/server.key")),
              "tls-secrets-policy",
              "shared"
            )
          }
        }
      }
    }
  }
}

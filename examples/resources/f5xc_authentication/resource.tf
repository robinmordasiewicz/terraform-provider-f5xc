# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

# Basic Authentication configuration
resource "f5xc_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Cookie Parameters. Specifies different cookie related con...
  cookie_params {
    # Configure cookie_params settings
  }
  # HMAC Key Pair. HMAC primary and secondary keys to be used...
  auth_hmac {
    # Configure auth_hmac settings
  }
  # Secret. SecretType is used in an object to indicate a sen...
  prim_key {
    # Configure prim_key settings
  }
}

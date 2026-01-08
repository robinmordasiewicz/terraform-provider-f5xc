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
  # Specifies different cookie related config parameters for ...
  cookie_params {
    # Configure cookie_params settings
  }
  # HMAC primary and secondary keys to be used for hashing th...
  auth_hmac {
    # Configure auth_hmac settings
  }
  # SecretType is used in an object to indicate a sensitive/c...
  prim_key {
    # Configure prim_key settings
  }
}

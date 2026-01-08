# Site Resource Example
# Manages a Site resource in F5 Distributed Cloud for gcp vpc site specification. configuration.

# Basic Site configuration
resource "f5xc_site" "example" {
  name      = "example-site"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Secret. SecretType is used in an object to indicate a sen...
  admin_password {
    # Configure admin_password settings
  }
  # Blindfold Secret. BlindfoldSecretInfoType specifies infor...
  blindfold_secret_info {
    # Configure blindfold_secret_info settings
  }
  # In-Clear Secret. ClearSecretInfoType specifies informatio...
  clear_secret_info {
    # Configure clear_secret_info settings
  }
}

# Site Resource Example
# Manages a Site resource in F5 Distributed Cloud for azure vnet site specification. configuration.

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
  # SecretType is used in an object to indicate a sensitive/c...
  admin_password {
    # Configure admin_password settings
  }
  # BlindfoldSecretInfoType specifies information about the S...
  blindfold_secret_info {
    # Configure blindfold_secret_info settings
  }
  # ClearSecretInfoType specifies information about the Secre...
  clear_secret_info {
    # Configure clear_secret_info settings
  }
}

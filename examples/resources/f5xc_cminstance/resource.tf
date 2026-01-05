# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace. in F5 Distributed Cloud.

# Basic Cminstance configuration
resource "f5xc_cminstance" "example" {
  name      = "example-cminstance"
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
  api_token {
    # Configure api_token settings
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

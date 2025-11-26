# Container Registry Resource Example
# Manages a ContainerRegistry resource in F5 Distributed Cloud for container image registry configuration.

# Basic Container Registry configuration
resource "f5xc_container_registry" "example" {
  name      = "example-container-registry"
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
  password {
    # Configure password settings
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

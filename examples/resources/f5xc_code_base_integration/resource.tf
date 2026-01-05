# Code Base Integration Resource Example
# Manages integration details. in F5 Distributed Cloud.

# Basic Code Base Integration configuration
resource "f5xc_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Integration Data. Choose your code base (e.g. GitHub, Git...
  code_base_integration {
    # Configure code_base_integration settings
  }
  # Azure Repos Integration.
  azure_repos {
    # Configure azure_repos settings
  }
  # Secret. SecretType is used in an object to indicate a sen...
  access_token {
    # Configure access_token settings
  }
}

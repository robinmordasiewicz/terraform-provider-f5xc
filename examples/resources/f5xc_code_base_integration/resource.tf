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
  # Choose your code base (e.g. GitHub, GitLab, Bitbucket, Az...
  code_base_integration {
    # Configure code_base_integration settings
  }
  # Azure Repos Integration.
  azure_repos {
    # Configure azure_repos settings
  }
  # SecretType is used in an object to indicate a sensitive/c...
  access_token {
    # Configure access_token settings
  }
}

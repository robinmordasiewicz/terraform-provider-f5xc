# Tpm Api Key Resource Example
# Manages a TpmAPIKey resource in F5 Distributed Cloud for apikey object when successfully created returns actual apikey bytes which is used by the users to call in to tpm provisioning api. configuration.

# Basic Tpm Api Key configuration
resource "f5xc_tpm_api_key" "example" {
  name      = "example-tpm-api-key"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # TPM Category. APIKey needs a reference to an existing TPM...
  category_ref {
    # Configure category_ref settings
  }
}

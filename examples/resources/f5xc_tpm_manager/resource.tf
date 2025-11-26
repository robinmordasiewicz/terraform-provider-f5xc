# Tpm Manager Resource Example
# Manages a TpmManager resource in F5 Distributed Cloud for create a tpm manager configuration.

# Basic Tpm Manager configuration
resource "f5xc_tpm_manager" "example" {
  name      = "example-tpm-manager"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

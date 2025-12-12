# Tpm Category Resource Example
# [Namespace: required] Manages Category object, which is a grouping of APIKeys used for TPM provisioning in F5 Distributed Cloud.

# Basic Tpm Category configuration
resource "f5xc_tpm_category" "example" {
  name      = "example-tpm-category"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # TPM Manager reference. Reference to TPM Manager
  tpm_manager_ref {
    # Configure tpm_manager_ref settings
  }
}

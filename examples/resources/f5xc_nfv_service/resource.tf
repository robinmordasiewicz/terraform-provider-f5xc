# Nfv Service Resource Example
# Manages new NFV service with configured parameters. in F5 Distributed Cloud.

# Basic Nfv Service configuration
resource "f5xc_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: disable_https_management, https_management; Defau...
  disable_https_management {
    # Configure disable_https_management settings
  }
  # [OneOf: disable_ssh_access, enabled_ssh_access; Default: ...
  disable_ssh_access {
    # Configure disable_ssh_access settings
  }
  # SSH based management. SSH based configuration.
  enabled_ssh_access {
    # Configure enabled_ssh_access settings
  }
}

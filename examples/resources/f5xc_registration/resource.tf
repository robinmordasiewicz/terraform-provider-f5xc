# Registration Resource Example
# [Namespace: required] Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users. configuration.

# Basic Registration configuration
resource "f5xc_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Infra Information. InfraMetadata stores information about...
  infra {
    # Configure infra settings
  }
  # Os Info. OsInfo holds information about host OS and HW
  hw_info {
    # Configure hw_info settings
  }
  # Bios Data. BIOS information.
  bios {
    # Configure bios settings
  }
}

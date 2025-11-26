# Protocol Inspection Resource Example
# Manages Protocol Inspection Specification in a given namespace. If one already exists it will give an error. in F5 Distributed Cloud.

# Basic Protocol Inspection configuration
resource "f5xc_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Enable/Disable Compliance Checks. x-required Enable Disab...
  enable_disable_compliance_checks {
    # Configure enable_disable_compliance_checks settings
  }
  # Empty. This can be used for messages where no values are ...
  disable_compliance_checks {
    # Configure disable_compliance_checks settings
  }
  # Object reference. This type establishes a direct referenc...
  enable_compliance_checks {
    # Configure enable_compliance_checks settings
  }
}

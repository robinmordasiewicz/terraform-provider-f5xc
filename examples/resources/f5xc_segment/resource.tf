# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment configuration.

# Basic Segment configuration
resource "f5xc_segment" "example" {
  name      = "example-segment"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: disable, enable] Enable this option
  disable {
    # Configure disable settings
  }
  # Enable this option
  enable {
    # Configure enable settings
  }
}

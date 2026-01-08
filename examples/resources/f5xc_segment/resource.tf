# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

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
  # [OneOf: disable, enable] Can be used for messages where n...
  disable {
    # Configure disable settings
  }
  # Can be used for messages where no values are needed.
  enable {
    # Configure enable settings
  }
}

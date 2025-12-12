# Segment Resource Example
# [Namespace: required] Manages a Segment resource in F5 Distributed Cloud for segment configuration.

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
  # [OneOf: disable, enable] Empty. This can be used for mess...
  disable {
    # Configure disable settings
  }
  # Empty. This can be used for messages where no values are ...
  enable {
    # Configure enable settings
  }
}

# Usb Policy Resource Example
# Manages new USB policy object. in F5 Distributed Cloud.

# Basic Usb Policy configuration
resource "f5xc_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # List of allowed USB devices .
  allowed_devices {
    # Configure allowed_devices settings
  }
}

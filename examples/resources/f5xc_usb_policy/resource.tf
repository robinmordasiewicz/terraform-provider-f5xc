# Usb Policy Resource Example
# [Namespace: required] Manages a Usb Policy resource in F5 Distributed Cloud for creates a new usb policy configuration.

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
  # Allowed USB devices. List of allowed USB devices
  allowed_devices {
    # Configure allowed_devices settings
  }
}

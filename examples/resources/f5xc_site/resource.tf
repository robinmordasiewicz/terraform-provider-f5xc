# Site Resource Example
# Manages a Site resource in F5 Distributed Cloud for app stack site specification. configuration.

# Basic Site configuration
resource "f5xc_site" "example" {
  name      = "example-site"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: allow_all_usb, deny_all_usb, usb_policy] Enable t...
  allow_all_usb {
    # Configure allow_all_usb settings
  }
  # [OneOf: blocked_services, default_blocked_services; Defau...
  blocked_services {
    # Configure blocked_services settings
  }
  # Disable Node Local Services. Blocking or denial configura...
  blocked_sevice {
    # Configure blocked_sevice settings
  }
}

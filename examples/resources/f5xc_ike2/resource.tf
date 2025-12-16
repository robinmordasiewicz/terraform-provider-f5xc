# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile configuration.

# Basic Ike2 configuration
resource "f5xc_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: dh_group_set, disable_pfs; Default: disable_pfs] ...
  dh_group_set {
    # Configure dh_group_set settings
  }
  # Enable this option
  disable_pfs {
    # Configure disable_pfs settings
  }
  # [OneOf: ike_keylifetime_hours, ike_keylifetime_minutes, u...
  ike_keylifetime_hours {
    # Configure ike_keylifetime_hours settings
  }
}

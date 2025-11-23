# Ike Phase2 Profile Resource Example
# Manages a IKEPhase2Profile resource in F5 Distributed Cloud for ike phase2 profile configuration.

# Basic Ike Phase2 Profile configuration
resource "f5xc_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: dh_group_set, disable_pfs] Diffie Hellman Groups....
  dh_group_set {
    # Configure dh_group_set settings
  }
  # Empty. This can be used for messages where no values are ...
  disable_pfs {
    # Configure disable_pfs settings
  }
  # [OneOf: ike_keylifetime_hours, ike_keylifetime_minutes, u...
  ike_keylifetime_hours {
    # Configure ike_keylifetime_hours settings
  }
}

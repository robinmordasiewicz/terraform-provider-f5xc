# Ike1 Resource Example
# [Namespace: required] Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile configuration.

# Basic Ike1 configuration
resource "f5xc_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: ike_keylifetime_hours, ike_keylifetime_minutes, u...
  ike_keylifetime_hours {
    # Configure ike_keylifetime_hours settings
  }
  # Minutes. Set IKE Key Lifetime in minutes
  ike_keylifetime_minutes {
    # Configure ike_keylifetime_minutes settings
  }
  # [OneOf: reauth_disabled, reauth_timeout_days, reauth_time...
  reauth_disabled {
    # Configure reauth_disabled settings
  }
}

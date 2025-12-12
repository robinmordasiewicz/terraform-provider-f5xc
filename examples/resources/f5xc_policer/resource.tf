# Policer Resource Example
# [Namespace: required] Manages new policer with traffic rate limits in F5 Distributed Cloud.

# Basic Policer configuration
resource "f5xc_policer" "example" {
  name      = "example-policer"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # API-discovered default values (shown for reference)
  # These values are applied by the API if not specified
  # policer_mode = "POLICER_MODE_NOT_SHARED"  # API default
  # policer_type = "POLICER_SINGLE_RATE_TWO_COLOR"  # API default
}

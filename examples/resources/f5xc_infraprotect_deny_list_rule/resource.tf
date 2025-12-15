# Infraprotect Deny List Rule Resource Example
# Manages DDoS transit Deny List Rule in F5 Distributed Cloud.

# Basic Infraprotect Deny List Rule configuration
resource "f5xc_infraprotect_deny_list_rule" "example" {
  name      = "example-infraprotect-deny-list-rule"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: expiration_never, expiration_timestamp, one_day, ...
  expiration_never {
    # Configure expiration_never settings
  }
  # Enable this option
  one_day {
    # Configure one_day settings
  }
  # Enable this option
  one_hour {
    # Configure one_hour settings
  }
}

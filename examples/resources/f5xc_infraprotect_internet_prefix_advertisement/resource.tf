# Infraprotect Internet Prefix Advertisement Resource Example
# [Namespace: required] Manages DDoS transit Internet Prefix in F5 Distributed Cloud.

# Basic Infraprotect Internet Prefix Advertisement configuration
resource "f5xc_infraprotect_internet_prefix_advertisement" "example" {
  name      = "example-infraprotect-internet-prefix-advertisement"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: activation_announce, activation_withdraw] Enable ...
  activation_announce {
    # Configure activation_announce settings
  }
  # Enable this option
  activation_withdraw {
    # Configure activation_withdraw settings
  }
  # [OneOf: expiration_never, expiration_timestamp] Enable th...
  expiration_never {
    # Configure expiration_never settings
  }
}

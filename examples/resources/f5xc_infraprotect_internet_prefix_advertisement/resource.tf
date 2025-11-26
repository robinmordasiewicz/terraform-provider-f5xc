# Infraprotect Internet Prefix Advertisement Resource Example
# Manages DDoS transit Internet Prefix in F5 Distributed Cloud.

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
  # [OneOf: activation_announce, activation_withdraw] Empty. ...
  activation_announce {
    # Configure activation_announce settings
  }
  # Empty. This can be used for messages where no values are ...
  activation_withdraw {
    # Configure activation_withdraw settings
  }
  # [OneOf: expiration_never, expiration_timestamp] Empty. Th...
  expiration_never {
    # Configure expiration_never settings
  }
}

# Address Allocator Resource Example
# Manages Address Allocator will create an address allocator object in 'system' namespace of the user. in F5 Distributed Cloud.

# Basic Address Allocator configuration
resource "f5xc_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Decides the scheme to be used to allocate addresses from ...
  address_allocation_scheme {
    # Configure address_allocation_scheme settings
  }
}

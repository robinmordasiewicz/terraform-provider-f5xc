# Protocol Policer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits. in F5 Distributed Cloud.

# Basic Protocol Policer configuration
resource "f5xc_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # List of L4 protocol match condition and associated traffi...
  protocol_policer {
    # Configure protocol_policer settings
  }
  # Reference to policer object to apply traffic rate limits .
  policer {
    # Configure policer settings
  }
  # Protocol and protocol specific flags to be matched in pac...
  protocol {
    # Configure protocol settings
  }
}

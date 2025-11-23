# Protocol Policer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

# Basic Protocol Policer configuration
resource "f5xc_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Protocol Policer. List of L4 protocol match condition and...
  protocol_policer {
    # Configure protocol_policer settings
  }
  # Policer. Reference to policer object to apply traffic rat...
  policer {
    # Configure policer settings
  }
  # Protocol Type. Protocol and protocol specific flags to be...
  protocol {
    # Configure protocol settings
  }
}

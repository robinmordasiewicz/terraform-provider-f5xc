# NAT Policy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,. configuration.

# Basic NAT Policy configuration
resource "f5xc_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # List of rules to apply under the NAT Policy. Rule that ma...
  rules {
    # Configure rules settings
  }
  # Action to apply on the packet if the NAT rule is applied.
  action {
    # Configure action settings
  }
  # Dynamic Pool. Dynamic Pool Configuration.
  dynamic {
    # Configure dynamic settings
  }
}

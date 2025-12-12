# Customer Support Resource Example
# [Namespace: required] Manages new customer support ticket in our customer support provider system. in F5 Distributed Cloud.

# Basic Customer Support configuration
resource "f5xc_customer_support" "example" {
  name      = "example-customer-support"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Comments. Comments are all public comments on an issue. T...
  comments {
    # Configure comments settings
  }
  # Attachments details. Information about any attachments (s...
  attachments_info {
    # Configure attachments_info settings
  }
  # Ticket which this one relates to. Optional reference to a...
  relates_to {
    # Configure relates_to settings
  }
}

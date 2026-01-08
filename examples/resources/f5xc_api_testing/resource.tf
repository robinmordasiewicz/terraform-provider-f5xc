# API Testing Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

# Basic API Testing configuration
resource "f5xc_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Add and configure testing domains and credentials .
  domains {
    # Configure domains settings
  }
  # Add credentials for API testing to use in the selected en...
  credentials {
    # Configure credentials settings
  }
  # Enable this option
  admin {
    # Configure admin settings
  }
}

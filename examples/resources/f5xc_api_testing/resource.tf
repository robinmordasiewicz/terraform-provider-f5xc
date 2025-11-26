# Api Testing Resource Example
# Manages a APITesting resource in F5 Distributed Cloud.

# Basic Api Testing configuration
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
  # Testing Environments. Add and configure testing domains a...
  domains {
    # Configure domains settings
  }
  # Credentials. Add credentials for API testing to use in th...
  credentials {
    # Configure credentials settings
  }
  # Empty. This can be used for messages where no values are ...
  admin {
    # Configure admin settings
  }
}

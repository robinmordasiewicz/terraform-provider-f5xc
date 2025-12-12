# API Crawler Resource Example
# [Category: API Security] [Namespace: required] Manages a API Crawler resource in F5 Distributed Cloud.

# Basic API Crawler configuration
resource "f5xc_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # API Crawler. API Crawler Configuration
  domains {
    # Configure domains settings
  }
  # Simple Login.
  simple_login {
    # Configure simple_login settings
  }
  # Secret. SecretType is used in an object to indicate a sen...
  password {
    # Configure password settings
  }
}

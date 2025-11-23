# Api Discovery Resource Example
# Manages api discovery creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Api Discovery configuration
resource "f5xc_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # Custom Authentication Types. Select your custom authentic...
    custom_auth_types {
      # Configure custom_auth_types settings
    }
}

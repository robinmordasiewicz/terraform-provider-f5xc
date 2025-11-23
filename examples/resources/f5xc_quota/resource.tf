# Quota Resource Example
# Manages quota creates a given object from storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Quota configuration
resource "f5xc_quota" "example" {
  name      = "example-quota"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # API Limits. API Limits defines ratelimit parameters for a...
    api_limits {
      # Configure api_limits settings
    }
    # Object Limits. Object Limits define maximum number of ins...
    object_limits {
      # Configure object_limits settings
    }
    # Resource Limits. Resource Limits define maximum value of ...
    resource_limits {
      # Configure resource_limits settings
    }
}

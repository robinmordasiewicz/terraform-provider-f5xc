# Service Policy Resource Example
# Manages service_policy creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Service Policy configuration
resource "f5xc_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Service Policy configuration
  algo = "FIRST_MATCH"

  # Allow specific paths
  rules {
    metadata {
      name = "allow-api"
    }
    spec {
      action = "ALLOW"
      path {
        prefix = "/api/"
      }
    }
  }
}

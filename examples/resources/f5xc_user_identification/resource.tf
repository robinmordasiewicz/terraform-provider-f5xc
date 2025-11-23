# User Identification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic User Identification configuration
resource "f5xc_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # User Identification configuration
  rules {
    identifier_type = "CLIENT_IP"
    any_client {}
  }
}

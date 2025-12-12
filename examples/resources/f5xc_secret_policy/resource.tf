# Secret Policy Resource Example
# [Category: Authentication] [Namespace: required] Manages secret_policy creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Secret Policy configuration
resource "f5xc_secret_policy" "example" {
  name      = "example-secret-policy"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Secret Policy configuration
  algo = "DENY_OVERRIDES"

  rules {
    metadata {
      name = "allow-team-secrets"
    }
    spec {
      action = "ALLOW"
      secret_match {
        regex {
          regex = "team-.*"
        }
      }
    }
  }
}

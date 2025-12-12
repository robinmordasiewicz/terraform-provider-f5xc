# Malicious User Mitigation Resource Example
# [Category: Security] [Namespace: required] Manages malicious_user_mitigation creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Malicious User Mitigation configuration
resource "f5xc_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Malicious User Mitigation configuration
  # Detection rules
  rules {
    threat_level = "HIGH"
    mitigation_action {
      block {
        body   = "Access denied"
        status = "403"
      }
    }
  }
}

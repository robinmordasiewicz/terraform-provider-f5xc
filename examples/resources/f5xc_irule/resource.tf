# Irule Resource Example
# [Namespace: required] Manages iRule in a given namespace. If one already exists it will give an error. in F5 Distributed Cloud.

# Basic Irule configuration
resource "f5xc_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

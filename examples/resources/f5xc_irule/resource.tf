# Irule Resource Example
# Manages a Irule resource in F5 Distributed Cloud for desired state for big-ip irule service. configuration.

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

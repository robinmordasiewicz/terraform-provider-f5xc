# Bigip Irule Resource Example
# Manages a BigIPIrule resource in F5 Distributed Cloud for desired state for big-ip irule service configuration.

# Basic Bigip Irule configuration
resource "f5xc_bigip_irule" "example" {
  name      = "example-bigip-irule"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

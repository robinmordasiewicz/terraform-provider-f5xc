# Fleet Resource Example
# [Namespace: required] Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

# Basic Fleet configuration
resource "f5xc_fleet" "example" {
  name      = "example-fleet"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Fleet configuration
  fleet_label = "env=production"

  # Network connectors
  inside_virtual_network {
    name      = "inside-network"
    namespace = "system"
  }

  outside_virtual_network {
    name      = "outside-network"
    namespace = "system"
  }

  # Default config
  default_config {}
}

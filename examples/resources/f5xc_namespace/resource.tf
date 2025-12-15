# Namespace Resource Example
# Manages new namespace. Name of the object is name of the name space. in F5 Distributed Cloud.

# Basic Namespace configuration
resource "f5xc_namespace" "example" {
  name      = "example-namespace"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Namespace configuration
  description = "Example namespace for application workloads"
}

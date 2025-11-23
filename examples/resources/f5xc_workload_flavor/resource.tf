# Workload Flavor Resource Example
# Manages workload_flavor in F5 Distributed Cloud.

# Basic Workload Flavor configuration
resource "f5xc_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

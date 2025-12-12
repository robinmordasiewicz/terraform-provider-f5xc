# Workload Flavor Resource Example
# [Namespace: required] Manages workload_flavor in F5 Distributed Cloud.

# Basic Workload Flavor configuration
resource "f5xc_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

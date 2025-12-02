# Virtual K8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

# Basic Virtual K8S configuration
resource "f5xc_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  // Virtual site selection for workload deployment
  vsite_refs {
    name      = "example-virtual-site"
    namespace = "system"
  }

  // One of the arguments from this list "disabled isolated" must be set

  isolated {}

  // Default workload flavor reference
  default_flavor_ref {
    name      = "example-workload-flavor"
    namespace = "staging"
  }
}

# Virtual K8s Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

# Basic Virtual K8s configuration
resource "f5xc_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Virtual Kubernetes configuration
  # Virtual site selection
  vsite_refs {
    name      = "example-virtual-site"
    namespace = "system"
  }

  # Disable cluster global access
  disabled {}
}

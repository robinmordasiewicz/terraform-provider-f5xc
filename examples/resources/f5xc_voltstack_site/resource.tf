# Voltstack Site Resource Example
# [Category: Sites] [Namespace: required] Manages a Voltstack Site resource in F5 Distributed Cloud for deploying Volterra stack sites for edge computing.

# Basic Voltstack Site configuration
resource "f5xc_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Voltstack Site configuration
  # Kubernetes configuration
  k8s_cluster {
    name      = "example-k8s-cluster"
    namespace = "system"
  }

  # Master nodes configuration
  master_nodes = ["master1.example.com"]

  # Default fleet configuration
  default_fleet_config {
    no_bond_devices {}
    no_dc_cluster_group {}
    default_storage_config {}
    no_gpu {}
  }

  # Disable HA by default
  disable_ha {}

  # No worker nodes
  no_worker_nodes {}
}

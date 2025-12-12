# Securemesh Site Resource Example
# [Category: Sites] [Namespace: required] Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security capabilities.

# Basic Securemesh Site configuration
resource "f5xc_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Secure Mesh Site configuration
  # Generic provider
  generic {
    not_managed {
      node_list {
        hostname  = "node1.example.com"
        public_ip = "203.0.113.10"
        type      = "Control"
      }
    }
  }

  # Master nodes
  master_nodes_count = 1

  # Default fleet config
  default_fleet_config {}

  # Disable HA
  disable_ha {}
}

# Dc Cluster Group Resource Example
# Manages DC Cluster group in given namespace. in F5 Distributed Cloud.

# Basic Dc Cluster Group configuration
resource "f5xc_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # DC Cluster Group Mesh Type. Details of DC Cluster Group M...
  type {
    # Configure type settings
  }
  # Can be used for messages where no values are needed.
  control_and_data_plane_mesh {
    # Configure control_and_data_plane_mesh settings
  }
  # Can be used for messages where no values are needed.
  data_plane_mesh {
    # Configure data_plane_mesh settings
  }
}

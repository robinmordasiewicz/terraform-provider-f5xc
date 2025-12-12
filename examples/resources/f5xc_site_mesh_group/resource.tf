# Site Mesh Group Resource Example
# [Namespace: required] Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

# Basic Site Mesh Group configuration
resource "f5xc_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Site Mesh Group configuration
  type = "SITE_MESH_GROUP_TYPE_FULL_MESH"

  # Control and data plane settings
  full_mesh {
    control_and_data_plane_mesh {}
  }

  # Hub status
  hub {}

  # Virtual site reference
  virtual_site {
    name      = "example-virtual-site"
    namespace = "system"
  }
}

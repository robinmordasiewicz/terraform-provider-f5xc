# Srv6 Network Slice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Srv6 Network Slice configuration
resource "f5xc_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }
}

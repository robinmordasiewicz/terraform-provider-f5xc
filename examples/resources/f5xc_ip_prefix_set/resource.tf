# IP Prefix Set Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic IP Prefix Set configuration
resource "f5xc_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # IP Prefix Set configuration
  prefix = ["192.168.1.0/24", "10.0.0.0/8"]
}

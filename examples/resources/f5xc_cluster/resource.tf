# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace. in F5 Distributed Cloud.

# Basic Cluster configuration
resource "f5xc_cluster" "example" {
  name      = "example-cluster"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: auto_http_config, http1_config, http2_options] Ca...
  auto_http_config {
    # Configure auto_http_config settings
  }
  # CircuitBreaker provides a mechanism for watching failures...
  circuit_breaker {
    # Configure circuit_breaker settings
  }
  # List of key-value pairs that define default subset. This ...
  default_subset {
    # Configure default_subset settings
  }
}

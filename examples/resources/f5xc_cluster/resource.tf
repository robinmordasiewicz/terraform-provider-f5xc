# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

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
  # [OneOf: auto_http_config, http1_config, http2_options] Em...
  auto_http_config {
    # Configure auto_http_config settings
  }
  # Circuit Breaker. CircuitBreaker provides a mechanism for ...
  circuit_breaker {
    # Configure circuit_breaker settings
  }
  # Default Subset. List of key-value pairs that define defau...
  default_subset {
    # Configure default_subset settings
  }
}

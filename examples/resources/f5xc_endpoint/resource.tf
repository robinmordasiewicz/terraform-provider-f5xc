# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace. in F5 Distributed Cloud.

# Basic Endpoint configuration
resource "f5xc_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Specifies name and TTL used for DNS resolution.
  dns_name_advanced {
    # Configure dns_name_advanced settings
  }
  # Specifies whether endpoint service is discovered by name ...
  service_info {
    # Configure service_info settings
  }
  # Type can be used to establish a 'selector reference' from...
  service_selector {
    # Configure service_selector settings
  }
}

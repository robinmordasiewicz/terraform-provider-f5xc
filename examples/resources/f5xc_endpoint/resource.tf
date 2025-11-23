# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

# Basic Endpoint configuration
resource "f5xc_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # DNS Name Advanced Type. Specifies name and TTL used for D...
    dns_name_advanced {
      # Configure dns_name_advanced settings
    }
    # Service Info Type. Specifies whether endpoint service is ...
    service_info {
      # Configure service_info settings
    }
    # Label Selector. This type can be used to establish a 'sel...
    service_selector {
      # Configure service_selector settings
    }
}

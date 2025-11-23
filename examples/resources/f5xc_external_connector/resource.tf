# External Connector Resource Example
# Manages a ExternalConnector resource in F5 Distributed Cloud for external_connector configuration configuration.

# Basic External Connector configuration
resource "f5xc_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
    # Object reference. This type establishes a direct referenc...
    ce_site_reference {
      # Configure ce_site_reference settings
    }
    # IPSec. External Connector with IPSec tunnel
    ipsec {
      # Configure ipsec settings
    }
    # IKE Parameters. IKE configuration parameters required for...
    ike_parameters {
      # Configure ike_parameters settings
    }
}

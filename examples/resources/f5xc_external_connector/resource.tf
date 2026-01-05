# External Connector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification. configuration.

# Basic External Connector configuration
resource "f5xc_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"

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
  # IPsec. External Connector with IPsec tunnel.
  ipsec {
    # Configure ipsec settings
  }
  # IKE Parameters. IKE configuration parameters required for...
  ike_parameters {
    # Configure ike_parameters settings
  }
}

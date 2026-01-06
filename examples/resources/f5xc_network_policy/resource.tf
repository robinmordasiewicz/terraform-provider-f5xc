# Network Policy Resource Example
# Manages a Network Policy resource in F5 Distributed Cloud for network policy view specification. configuration.

# Basic Network Policy configuration
resource "f5xc_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Network Policy configuration
  endpoint {
    any {}
  }

  ingress_rules {
    metadata {
      name = "allow-http"
    }
    spec {
      action = "ALLOW"
      any    = {}
    }
  }

  egress_rules {
    metadata {
      name = "allow-all-egress"
    }
    spec {
      action = "ALLOW"
      any    = {}
    }
  }
}

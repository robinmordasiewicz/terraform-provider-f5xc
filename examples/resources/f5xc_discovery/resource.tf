# Discovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Discovery configuration
resource "f5xc_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Discovery configuration
  discovery_k8s {
    access_info {
      kubeconfig_url {
        clear_secret_info {
          url = "string:///base64-kubeconfig"
        }
      }
      isolated {}
    }
    publish_info {
      disable {}
    }
  }

  # Site selection
  where {
    site {
      ref {
        name      = "example-site"
        namespace = "system"
      }
      network_type = "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"
    }
  }
}

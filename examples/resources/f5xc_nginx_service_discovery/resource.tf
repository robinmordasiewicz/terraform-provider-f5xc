# Nginx Service Discovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace. configuration.

# Basic Nginx Service Discovery configuration
resource "f5xc_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Discovery Target.
  discovery_target {
    # Configure discovery_target settings
  }
  # ConfigSyncGroup Reference. Select new ConfigSyncGroup.
  config_sync_group {
    # Configure config_sync_group settings
  }
  # NGINXInstance Reference. Select new NGINX Instance.
  nginx_instance {
    # Configure nginx_instance settings
  }
}

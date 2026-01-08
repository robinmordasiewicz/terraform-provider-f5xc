# App API Group Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic App API Group configuration
resource "f5xc_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: bigip_virtual_server, cdn_loadbalancer, http_load...
  bigip_virtual_server {
    # Configure bigip_virtual_server settings
  }
  # Set the scope of the API Group to a specific CDN Loadbala...
  cdn_loadbalancer {
    # Configure cdn_loadbalancer settings
  }
  # List of API group elements with methods and path regex fo...
  elements {
    # Configure elements settings
  }
}

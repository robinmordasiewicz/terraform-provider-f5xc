# Secret Management Access Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Secret Management Access configuration
resource "f5xc_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # HostAccessInfoType contains the information about how to ...
  access_info {
    # Configure access_info settings
  }
  # Authentication parameters for REST based hosts.
  rest_auth_info {
    # Configure rest_auth_info settings
  }
  # AuthnTypeBasicAuth is used for using basic_auth mode of H...
  basic_auth {
    # Configure basic_auth settings
  }
}

# Secret Management Access Resource Example
# [Namespace: required] Manages secret_management_access creates a new object in storage backend for metadata.namespace. in F5 Distributed Cloud.

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
  # Host Access Information. HostAccessInfoType contains the ...
  access_info {
    # Configure access_info settings
  }
  # Rest Authentication Parameters. Authentication parameters...
  rest_auth_info {
    # Configure rest_auth_info settings
  }
  # BasicAuth Authentication Parameters. AuthnTypeBasicAuth i...
  basic_auth {
    # Configure basic_auth settings
  }
}

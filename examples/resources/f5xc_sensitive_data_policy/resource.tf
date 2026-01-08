# Sensitive Data Policy Resource Example
# Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace. in F5 Distributed Cloud.

# Basic Sensitive Data Policy configuration
resource "f5xc_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "shared"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # Select your custom data types to be monitored in the API ...
  custom_data_types {
    # Configure custom_data_types settings
  }
  # Type establishes a direct reference from one object(the r...
  custom_data_type_ref {
    # Configure custom_data_type_ref settings
  }
}

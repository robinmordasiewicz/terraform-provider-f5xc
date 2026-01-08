# App Setting Resource Example
# Manages App setting configuration in namespace metadata.namespace. in F5 Distributed Cloud.

# Basic App Setting configuration
resource "f5xc_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # List of settings to enable for each AppType, given instan...
  app_type_settings {
    # Configure app_type_settings settings
  }
  # The AppType of App instance in current Namespace. Associa...
  app_type_ref {
    # Configure app_type_ref settings
  }
  # Settings specifying how API Discovery will be performed.
  business_logic_markup_setting {
    # Configure business_logic_markup_setting settings
  }
}
